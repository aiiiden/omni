package e2e_test

import (
	"context"
	"testing"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/log"

	"github.com/cometbft/cometbft/rpc/client/http"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

// TestPlanCancelUpgrade tests planning and canceling a far-future upgrade.
func TestPlanCancelUpgrade(t *testing.T) {
	t.Parallel()
	skipFunc := func(manifest types.Manifest) bool {
		return !manifest.AllE2ETests
	}
	maybeTestNetwork(t, skipFunc, func(ctx context.Context, t *testing.T, deps NetworkDeps) {
		t.Helper()

		network := deps.Network
		cl, err := http.New(network.ID.Static().ConsensusRPC(), "/websocket")
		require.NoError(t, err)
		cprov := provider.NewABCI(cl, network.ID)

		backend, err := deps.OmniBackend()
		require.NoError(t, err)

		contract, err := bindings.NewUpgrade(common.HexToAddress(predeploys.Upgrade), backend)
		require.NoError(t, err)
		txOpts, err := backend.BindOpts(ctx, eoa.MustAddress(network.ID, eoa.RoleUpgrader))
		require.NoError(t, err)

		assertCurrentPlan := func(t *testing.T, name string) {
			t.Helper()
			require.Eventually(t, func() bool {
				current, _, err := cprov.CurrentPlannedPlan(ctx)
				return err == nil && current.Name == name
			}, time.Second*30, time.Second)
		}

		const upgrade = "far-future-upgrade"
		const farFuture = 1_000_000_000

		// Ensure no upgrade planned
		assertCurrentPlan(t, "")

		// Plan far future upgrade
		tx, err := contract.PlanUpgrade(txOpts, bindings.UpgradePlan{
			Name:   upgrade,
			Height: farFuture,
		})
		require.NoError(t, err)
		rc, err := backend.WaitMined(ctx, tx)
		require.NoError(t, err)
		log.Debug(ctx, "Planned far-future upgrade", "block", rc.BlockNumber)

		// Ensure far-future upgrade planned
		assertCurrentPlan(t, upgrade)

		// Cancel far future upgrade
		tx, err = contract.CancelUpgrade(txOpts)
		require.NoError(t, err)
		rc, err = backend.WaitMined(ctx, tx)
		require.NoError(t, err)
		log.Debug(ctx, "Canceled far-future upgrade", "block", rc.BlockNumber)

		// Ensure no upgrade planned
		assertCurrentPlan(t, "")
	})
}
