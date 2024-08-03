package psql

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/emma769/chatty/internal/config"
)

func newTestRepo(t *testing.T) *Repository {
	repo, err := NewRepository(&config.Getter{})
	if err != nil {
		require.FailNow(t, err.Error())
	}

	t.Cleanup(func() { _ = repo.Close() })
	return repo
}
