package posts

import "github.com/CcyBborg/golik-blog/internal/models"

type store interface {
	GetPosts(opts Opts) (ordersList []models.Post, err error)
}

type Service struct {
	store store
}

func New(store store) *Service {
	return &Service{store: store}
}

func (s *Service) GetPosts(opts Opts) ([]models.Post, error) {
	postList, err := s.store.GetPosts(opts)
	if err != nil {
		return nil, err
	}

	return postList, nil
}
