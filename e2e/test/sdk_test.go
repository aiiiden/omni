package e2e_test

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/e2e/types"

	"github.com/stretchr/testify/require"
)

// disableSDKTest allows devs to skip SDK tests if local npm isn't setup correctly
// by setting the environment variable DISABLE_SDK_TEST=true to skip the test.
const disableSDKTest = "DISABLE_SDK_TEST"

func TestSDK(t *testing.T) {
	t.Parallel()
	skipFunc := func(manifest types.Manifest) bool {
		return !manifest.AllE2ETests || os.Getenv(disableSDKTest) == "true"
	}
	maybeTestNetwork(t, skipFunc, func(ctx context.Context, t *testing.T, _ NetworkDeps) {
		t.Helper()

		sdkPath, err := filepath.Abs("../../sdk")
		require.NoError(t, err)

		exe := func(ctx context.Context, dir string, name string, args ...string) {
			cmd := exec.CommandContext(ctx, name, args...)
			cmd.Dir = dir

			out, err := cmd.CombinedOutput()
			t.Logf("command: %s %s\n\n%s", name, args, out)
			require.NoError(t, err)
		}

		exe(ctx, sdkPath, "pnpm", "install")
		exe(ctx, sdkPath, "pnpm", "run", "build")
		exe(ctx, sdkPath, "pnpm", "run", "test:integration")
	})
}
