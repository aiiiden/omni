//nolint:unused // Partially integrated
package app

import (
	"context"

	"github.com/omni-network/omni/anchor/anchorinbox"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/svmutil"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/solver/job"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func svmProcDeps(
	cl *rpc.Client,
	network netconf.ID,
	ethProcDeps procDeps,
) procDeps {
	resp := ethProcDeps
	resp.GetOrder = adaptSVMGetOrder(cl, network)

	return resp
}

func NewSVMStreamCallback(
	cl *rpc.Client,
	chainVer xchain.ChainVersion,
	cursors *cursors,
	jobDB *job.DB,
	asyncWork asyncWorkFunc,
) svmutil.StreamCallback {
	return func(ctx context.Context, sig *rpc.TransactionSignature) error {
		txResp, err := svmutil.AwaitConfirmedTransaction(ctx, cl, sig.Signature)
		if customErr := anchorinbox.DecodeMetaError(txResp); customErr != nil {
			log.Warn(ctx, "AnchorInbox: Ignoring failed tx", customErr, "tx", sig)
			return nil
		} else if err != nil {
			return errors.Wrap(err, "get tx")
		}

		events, err := anchorinbox.DecodeEvents(txResp, anchorinbox.ProgramID, func([]solana.PublicKey) (map[solana.PublicKey]solana.PublicKeySlice, error) {
			return nil, errors.New("address lookup not supported")
		})
		if err != nil {
			log.Warn(ctx, "AnchorInbox: ignoring decode events failure tx", err, "tx", sig)
			return nil
		}

		for i, event := range events {
			if event.Name != anchorinbox.EventNameUpdated {
				return errors.New("unexpected event [BUG]", "event", event.Name)
			}

			data, ok := event.Data.(*anchorinbox.EventUpdatedEventData)
			if !ok {
				return errors.New("unexpected event data [BUG]", "event", event.Name)
			}

			indexU64, err := umath.ToUint64(i)
			if err != nil {
				return err
			}

			statusU64, err := umath.ToUint64(data.Status)
			if err != nil {
				return err
			}

			j, err := jobDB.Insert(
				ctx,
				chainVer.ID,
				txResp.Slot,
				sig.Signature.String(),
				indexU64,
				data.OrderId[:],
				statusU64,
			)
			if err != nil {
				return err
			}

			if err := asyncWork(ctx, j); err != nil {
				return errors.Wrap(err, "async work")
			}

			if err := cursors.SetTxSig(ctx, chainVer, sig.Signature); err != nil {
				return errors.Wrap(err, "update cursor")
			}
		}

		return nil
	}
}

// adaptSVMGetOrder adapts the svmGetOrder function to the procDeps interface.
func adaptSVMGetOrder(cl *rpc.Client, network netconf.ID) func(context.Context, uint64, OrderID) (Order, bool, error) {
	return func(ctx context.Context, _ uint64, id OrderID) (Order, bool, error) {
		return svmGetOrder(ctx, cl, network, id)
	}
}
