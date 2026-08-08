package jsonstorage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sync"
	"time"

	"github.com/dakotaodev/brag/internal/brag"
)

type Repository struct {
	path string
	mu   sync.Mutex
}

func NewJsonRepository() (*Repository, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("unable to get platform dir for user: %w", err)
	}
	path := filepath.Join(dir, "brag")
	if err := os.MkdirAll(path, 0o700); err != nil {
		return nil, fmt.Errorf("unable create directory at %w", err)
	}
	return &Repository{path: filepath.Join(path, "brag.json")}, nil
}

func (r *Repository) Add(
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

func (r *Repository) List() ([]brag.Entry, error) {
	entries, err := r.readFile()
	if err != nil {
		return make([]brag.Entry, 0), err
	}
	return entries, nil
}

func (r *Repository) Delete(
	id string,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	entries, err := r.readFile()
	if err != nil {
		return err
	}
	found := false
	entries = slices.DeleteFunc(entries, func(entry brag.Entry) bool {

		if entry.ID == id {
			found = true
			return true
		}
		return false

	})

	if !found {
		return fmt.Errorf("Unable to delete record for ID %s as it was not found.", id)
	}

	return r.writeFile(entries)

}

func (r *Repository) Update(
	id string,
	newValue string,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	entries, err := r.readFile()
	if err != nil {
		return err
	}
	for i, e := range entries {
		if e.ID == id {
			e.Value = newValue
			e.ModifiedAt = time.Now()
			entries[i] = e
			return r.writeFile(entries)
		}
	}
	return errors.New("ID specified does not exist. Please validate or use the add command.")

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
