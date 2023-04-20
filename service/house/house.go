package house

import "context"

type House struct {
	ID       string
	Slug     string
	Title    string
	Price    string
	Location string
	Author   string
}

type Store interface {
	GetHouse(context.Context, string) (House, error)
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetHouse(ctx context.Context, id string) (House, error) {
	house, err := s.Store.GetHouse(ctx, id)
	if err != nil {
		return House{}, err
	}
	return house, nil
}
