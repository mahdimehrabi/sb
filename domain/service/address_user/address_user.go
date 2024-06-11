package address_user

import (
	"context"
	"errors"
	"fmt"
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

func NewService(logger loggerInfra.Logger, addressRepo address.Address, userRepo user.User) *Service {
	s := &Service{
		articleRepository: addressRepo,
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
		for repeat := 0; repeat < 10; repeat++ {
			if s.work(job) {
				break
			}
		}
		s.logger.Error(fmt.Errorf("failed to insert %s, cancelling", job))
	}
}

func (s Service) work(job job) (done bool) {
	defer func() {
		if !done {
			//cool down before retry
			time.Sleep(500 * time.Millisecond)
		}
	}()
	if job.userID == 0 {
		userID, err := s.createUser(context.Background(), job.user)
		if err != nil {
			done = false
			return
		}
		job.userID = userID
		for _, addr := range job.addresses {
			addr.UserID = job.userID
		}
	}

	err := s.createBatchAddresses(context.Background(), job.addresses)
	if err != nil {
		done = false
		return
	}
	return true
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

func (s Service) Detail(ctx context.Context, id int64) (*entity.Address, error) {
	article, err := s.articleRepository.Detail(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return article, nil
}
