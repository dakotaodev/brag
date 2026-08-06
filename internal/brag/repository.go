package brag

import "context"

type Repository interface {
	Add(ctx context.Context, entry Entry) (Entry, error)
	List(ctx context.Context) ([]Entry, error)
}
