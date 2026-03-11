package nablesimulator_test

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tekkamanendless/httperror"
	"github.com/threatmate/nableclient"
	"github.com/threatmate/nableclient/nablesimulator"
)

func TestSimulator(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})))

	ctx := t.Context()

	simulator := nablesimulator.New(ctx)
	defer simulator.Close()

	simulator.Universe().APIUsers = []*nablesimulator.APIUser{
		{
			Username: "admin@example.com",
			APIKey:   "admin-key-1",
		},
	}
	simulator.Universe().ServiceOrgs = []*nableclient.ServiceOrg{
		{
			SOID:   "1",
			SOName: "Service Org 1",
		},
	}

	t.Run("auth", func(t *testing.T) {
		client := nableclient.New(simulator.URL())

		output, err := client.GetAuth(ctx)
		require.NoError(t, err)
		assert.Equal(t, "/api/auth/refresh", output.Refresh)
		assert.Equal(t, "/api/auth/validate", output.Validate)
		assert.Equal(t, "/api/auth/authenticate", output.Authenticate)
	})
	t.Run("auth/authenticate", func(t *testing.T) {
		t.Run("Bogus API key", func(t *testing.T) {
			client := nableclient.New(simulator.URL())

			err := client.Authenticate(ctx, "bogus")
			require.ErrorIs(t, err, httperror.ErrStatusUnauthorized)
		})
		t.Run("Valid API key", func(t *testing.T) {
			client := nableclient.New(simulator.URL())

			err := client.Authenticate(ctx, "admin-key-1")
			require.NoError(t, err)
		})
	})
	t.Run("Authenticated", func(t *testing.T) {
		client := nableclient.New(simulator.URL())
		err := client.Authenticate(ctx, "admin-key-1")
		require.NoError(t, err)

		t.Run("service-orgs", func(t *testing.T) {
			serviceOrgs, err := client.GetServiceOrgs(ctx)
			require.NoError(t, err)
			if assert.Equal(t, 1, len(serviceOrgs)) {
				assert.Equal(t, "1", serviceOrgs[0].SOID)
				assert.Equal(t, "Service Org 1", serviceOrgs[0].SOName)
			}
		})
	})
}
