package brag

import (
	"context"
	"time"
)

type Service struct {
	repo Repository
}

type AddInput struct {
	Category Category
	Value string
	Role string
	Source string
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Add(
	ctx context.Context,
	input AddInput, 
) (Entry, error) {
		
	now := time.Now().UTC()

	entry:= Entry{
		Category: input.Category,
		CreatedAt: now,
		ModifiedAt: now,
		Value: input.Value,
		Role: input.Role,
		Source: input.Source,
	}

	e, err := s.repo.Add(ctx, entry)
	if err != nil {
		return Entry{}, nil
	}
	return e, nil
}

func (s *Service) List(
	ctx context.Context,
) ([]Entry, error) {

	entries, err := s.repo.List(ctx)
	if err != nil {
		return make([]Entry, 0), nil
	}
	return entries, nil
}