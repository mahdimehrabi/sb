package address

import (
	"context"
	"m1-article-service/domain/entity"
	"m1-article-service/domain/repository/address"
	loggerInfra "m1-article-service/infrastructure/log"
)

type Service struct {
	articleRepository address.Address
	logger            loggerInfra.Logger
}

func NewService(logger loggerInfra.Logger, articleRepo address.Address) *Service {
	return &Service{
		articleRepository: articleRepo,
		logger:            logger,
	}
}

func (s Service) Create(ctx context.Context, article *entity.Address) (int64, error) {
	id, err := s.articleRepository.Create(ctx, article)
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
