package jsonstorage

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/dakotaodev/brag/internal/brag"
)

type Repository struct {
	path string
	mu   sync.Mutex
}

func NewJsonRepository(path string) *Repository {
	return &Repository{path: path}
}

func (r *Repository) Add(
	ctx context.Context,
	entry brag.Entry,
) (brag.Entry, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	entries, err := r.readFile()
	if err != nil {
		return brag.Entry{}, err
	}

	entries = append(entries, entry)

	err = r.writeFile(entries)
	if err != nil {
		return brag.Entry{}, err
	}
	return entry, nil

}

func (r *Repository) List(
	ctx context.Context,
) ([]brag.Entry, error) {
	entries, err := r.readFile()
	if err != nil {
		return make([]brag.Entry, 0), err
	}
	return entries, nil
}

func (r *Repository) readFile() ([]brag.Entry, error) {

	data, err := os.ReadFile(r.path)
	if errors.Is(err, os.ErrNotExist) {
		return []brag.Entry{}, nil
	}
	if err != nil {
		return nil, err
	}
	var values []brag.Entry
	err = json.Unmarshal(data, &values)
	if err != nil {
		return make([]brag.Entry, 0), err
	}
	return values, nil
}

func (r *Repository) writeFile(
	entries []brag.Entry,
) error {
	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(r.path, data, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
