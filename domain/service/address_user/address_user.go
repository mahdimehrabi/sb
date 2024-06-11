package address_user

import (
	"context"
	"errors"
	"m1-article-service/domain/entity"
	"m1-article-service/domain/repository/address"
	"m1-article-service/domain/repository/user"
	loggerInfra "m1-article-service/infrastructure/log"
	"time"
)

const maxWorkerCount = 10
const queueLength = 1000

var ErrServiceUnavailable = errors.New("service unavailable")

type job struct {
	addresses []*entity.Address
	user      *entity.User
	userID    int64
	//can hold task type , but we don't need it now
}

type Service struct {
	articleRepository address.Address
	userRepository    user.User
	logger            loggerInfra.Logger
	queue             chan job
}

func NewService(logger loggerInfra.Logger, articleRepo address.Address, userRepo user.User) *Service {
	s := &Service{
		articleRepository: articleRepo,
		userRepository:    userRepo,
		logger:            logger,
		queue:             make(chan job, queueLength),
	}
	s.startWorkers()
	return s
}

func (s Service) startWorkers() {
	for i := 0; i < maxWorkerCount; i++ {
		go s.worker()
	}
}

func (s Service) worker() {
	for job := range s.queue {
		if job.userID == 0 {
			userID, err := s.createUser(context.Background(), job.user)
			if err != nil {
				//cool down queue
				time.Sleep(100 * time.Millisecond)
				s.queue <- job
			}
			job.userID = userID
			for _, addr := range job.addresses {
				addr.UserID = job.userID
			}
		}

		err := s.createBatchAddresses(context.Background(), job.addresses)
		if err != nil {
			//cool down queue
			time.Sleep(100 * time.Millisecond)
			s.queue <- job
		}

	}
}

// I just implemented worker pool design for create because lack of time
func (s Service) Create(addresses []*entity.Address, user *entity.User) error {
	if len(s.queue) == queueLength {
		return ErrServiceUnavailable
	}
	s.queue <- job{
		addresses: addresses,
		user:      user,
	}
	return nil
}

func (s Service) createBatchAddresses(ctx context.Context, addresses []*entity.Address) error {
	err := s.articleRepository.BatchCreate(ctx, addresses)
	if err != nil {
		s.logger.Error(err)
		return err
	}
	return err
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
