package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sourcegraph/sourcegraph/internal/actor"
	"github.com/sourcegraph/sourcegraph/internal/database"
	"github.com/sourcegraph/sourcegraph/internal/database/dbtest"
	"github.com/sourcegraph/sourcegraph/internal/types"
)

func TestCodeMonitorStoreLastSearched(t *testing.T) {
	t.Parallel()

	type testFixtures struct {
		User    *types.User
		Monitor *Monitor
	}
	populateFixtures := func(db EnterpriseDB) testFixtures {
		ctx := context.Background()
		u, err := db.Users().Create(ctx, database.NewUser{Email: "test", Username: "test", EmailVerificationCode: "test"})
		require.NoError(t, err)
		ctx = actor.WithActor(ctx, actor.FromUser(u.ID))
		m, err := db.CodeMonitors().CreateMonitor(ctx, MonitorArgs{NamespaceUserID: &u.ID})
		require.NoError(t, err)
		return testFixtures{User: u, Monitor: m}
	}

	t.Run("insert get upsert get", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		db := NewEnterpriseDB(database.NewDB(dbtest.NewDB(t)))
		fixtures := populateFixtures(db)
		cm := db.CodeMonitors()

		// Insert
		insertLastSearched := []string{"commit1", "commit2"}
		err := cm.UpsertLastSearched(ctx, fixtures.Monitor.ID, 3851, insertLastSearched)
		require.NoError(t, err)

		// Get
		lastSearched, err := cm.GetLastSearched(ctx, fixtures.Monitor.ID, 3851)
		require.NoError(t, err)
		require.Equal(t, insertLastSearched, lastSearched)

		// Update
		updateLastSearched := []string{"commit3", "commit4"}
		err = cm.UpsertLastSearched(ctx, fixtures.Monitor.ID, 3851, updateLastSearched)
		require.NoError(t, err)

		// Get
		lastSearched, err = cm.GetLastSearched(ctx, fixtures.Monitor.ID, 3851)
		require.NoError(t, err)
		require.Equal(t, updateLastSearched, lastSearched)
	})

	t.Run("no error for missing get", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		db := NewEnterpriseDB(database.NewDB(dbtest.NewDB(t)))
		fixtures := populateFixtures(db)
		cm := db.CodeMonitors()

		// GetLastSearched should not return an error for a monitor that hasn't
		// been run yet. It should just return an empty value for lastSearched
		lastSearched, err := cm.GetLastSearched(ctx, fixtures.Monitor.ID+1, 19793)
		require.NoError(t, err)
		require.Empty(t, lastSearched)
	})

	t.Run("no error for missing get", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		db := NewEnterpriseDB(database.NewDB(dbtest.NewDB(t)))
		fixtures := populateFixtures(db)
		cm := db.CodeMonitors()

		// Insert with nil last searched
		err := cm.UpsertLastSearched(ctx, fixtures.Monitor.ID, 3851, nil)
		require.NoError(t, err)

		// Get nil last searched
		lastSearched, err := cm.GetLastSearched(ctx, fixtures.Monitor.ID, 3851)
		require.NoError(t, err)
		require.Empty(t, lastSearched)

		// Insert with empty last searched
		err = cm.UpsertLastSearched(ctx, fixtures.Monitor.ID, 3852, []string{})
		require.NoError(t, err)

		// Get nil last searched
		lastSearched, err = cm.GetLastSearched(ctx, fixtures.Monitor.ID, 3852)
		require.NoError(t, err)
		require.Empty(t, lastSearched)
	})
}
