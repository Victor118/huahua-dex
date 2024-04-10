package types

import (
	"cosmossdk.io/math"

	math_utils "github.com/ChihuahuaChain/chihuahua/utils/math"
)

type Liquidity interface {
	Swap(maxAmountTakerIn math.Int, maxAmountMakerOut *math.Int) (inAmount, outAmount math.Int)
	Price() math_utils.PrecDec
}
