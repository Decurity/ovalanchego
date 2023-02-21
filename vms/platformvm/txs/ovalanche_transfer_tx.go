// 2023, Decurity, special for ETHDubai

package txs

import (
	"errors"
	"fmt"

	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	//"github.com/ava-labs/avalanchego/vms/platformvm/locked"
)

var (
	_ UnsignedTx = (*TransferTx)(nil)

	errProducedNotEqualConsumed = errors.New("produced amount not equal to consumed amount")
	errNotAVAXAsset             = errors.New("input assetID isn't avax assetID")
	errWrongOutType             = errors.New("wrong output type")
)

type TransferTx struct {
	// Metadata, inputs and outputs
	BaseTx `serialize:"true"`
}

// SyntacticVerify returns nil if [tx] is valid
func (tx *TransferTx) SyntacticVerify(ctx *snow.Context) error {
	switch {
	case tx == nil:
		return ErrNilTx
	case tx.SyntacticallyVerified: // already passed syntactic verification
		return nil
	}

	if err := tx.BaseTx.SyntacticVerify(ctx); err != nil {
		return fmt.Errorf("failed to verify BaseTx: %w", err)
	}

	outAmount := uint64(0)
	for _, out := range tx.Outs {
		if outAssetID := out.AssetID(); outAssetID != ctx.AVAXAssetID {
			return fmt.Errorf("output 0 has asset ID %s but expect %s: %w",
				outAssetID, ctx.AVAXAssetID, errNotAVAXAsset)
		}

		_, ok := out.Out.(*secp256k1fx.TransferOutput)
		if !ok {
			return errWrongOutType
		}

		outAmount += out.Out.Amount()
	}

	inAmount := uint64(0)
	for i, in := range tx.Ins {
		if inputAssetID := in.AssetID(); inputAssetID != ctx.AVAXAssetID {
			return fmt.Errorf("input %d has asset ID %s but expect %s: %w",
				i, inputAssetID, ctx.AVAXAssetID, errNotAVAXAsset)
		}

		inAmount += in.In.Amount()
	}

	if outAmount != inAmount {
		return errProducedNotEqualConsumed
	}

	// cache that this is valid
	tx.SyntacticallyVerified = true
	return nil
}

func (tx *TransferTx) Visit(visitor Visitor) error {
	return visitor.TransferTx(tx)
}
