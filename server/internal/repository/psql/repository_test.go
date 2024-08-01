package psql

import (
	"testing"

	"github.com/emma769/chatty/internal/config"
)

func newTestRepo(t *testing.T) *Repository {
	repo, err := NewRepository(&config.Getter{})
	if err != nil {
		t.FailNow()
	}

	t.Cleanup(func() { _ = repo.Close() })
	return repo
}
