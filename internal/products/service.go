package products

import "context"

type Service interface {
	ListProducts(ctx context.Context) ([]string, error)
}

type svc struct {
	//repo
}

func NewService() Service {
	return &svc{}
}

func (s *svc) ListProducts(ctx context.Context) ([]string, error) {
	return []string{"product1", "product2", "product3"}, nil
}
