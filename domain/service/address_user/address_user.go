package address_user

import (
	"context"
	"m1-article-service/domain/entity"
	"m1-article-service/domain/repository/address"
	"m1-article-service/domain/repository/user"
	loggerInfra "m1-article-service/infrastructure/log"
)

const maxWorkerCount = 10
const queueLength = 1000

type job struct {
	address       *entity.Address
	user          *entity.User
	addressIdChan chan<- int64
	userIDChan    chan<- int64
	//can hold task type , but we don't need it now
}

type Service struct {
	articleRepository address.Address
	userRepository    user.User
	logger            loggerInfra.Logger
	queue             chan job
}

func NewService(logger loggerInfra.Logger, articleRepo address.Address) *Service {
	return &Service{
		articleRepository: articleRepo,
		logger:            logger,
		queue:             make(chan job, queueLength),
	}
}

func (s Service) startWorkers() {
	for i := 0; i < maxWorkerCount; i++ {
		go s.worker()
	}
}

func (s Service) worker() {
	for job := range s.queue {
		addrId, err := s.createAddress(context.Background(), job.address)
		if err == nil {
			job.addressIdChan <- addrId
		}
		userID, err := s.createUser(context.Background(), job.user)
		if err == nil {
			job.userIDChan <- userID
		}
	}
}

// I just implemented worker pool design for createAddress because lack of time
func (s Service) Create(address *entity.Address, user *entity.User) {
	s.queue <- job{
		address, user, make(chan<- int64), make(chan<- int64),
	}
}

func (s Service) createAddress(ctx context.Context, address *entity.Address) (int64, error) {
	id, err := s.articleRepository.Create(ctx, address)
	if err != nil {
		s.logger.Error(err)
		return 0, err
	}
	return id, err
}

func (s Service) createUser(ctx context.Context, user *entity.User) (int64, error) {
	id, err := s.userRepository.Create(ctx, user)
	if err != nil {
		s.logger.Error(err)
		return 0, err
	}
	return id, err
}

func (s Service) Update(ctx context.Context, article *entity.Address) error {
	if err := s.articleRepository.Update(ctx, article); err != nil {
		s.logger.Error(err)
		return err
	}
	return nil
}

func (s Service) Delete(ctx context.Context, id int64) error {
	if err := s.articleRepository.Delete(ctx, id); err != nil {
		s.logger.Error(err)
		return err
	}
	return nil
}

func (s Service) Detail(ctx context.Context, id int64) (*entity.Address, error) {
	article, err := s.articleRepository.Detail(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return article, nil
}

func (s Service) List(ctx context.Context, page uint16) ([]*entity.Address, error) {
	articles, err := s.articleRepository.List(ctx, page)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return articles, err
}
