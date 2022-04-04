package internal

import (
	"context"
	"errors"
)

type Service interface {
	GetAll(ctx context.Context) ([]Product, error)
	GetByID(ctx context.Context, id int) (Product, error)
	Create(ctx context.Context, p Product) (Product, error)
	Update(ctx context.Context, p Product, id int) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAll(ctx context.Context) ([]Product, error) {
	ps, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) GetByID(ctx context.Context, id int) (Product, error) {
	p, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (s *service) Create(ctx context.Context, p Product) (Product, error) {
	exist, _ := s.repo.Exist(ctx, p.ID)
	if exist {
		return Product{}, errors.New("Product already exists")
	}
	id, err := s.repo.Save(ctx, p)
	if err != nil {
		return Product{}, err
	}

	p.ID = id

	return p, nil
}

func (s *service) Update(ctx context.Context, p Product, id int) error {
	p.ID = id
	err := s.repo.Update(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
