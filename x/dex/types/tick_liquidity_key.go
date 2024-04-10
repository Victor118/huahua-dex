package types

import math_utils "github.com/ChihuahuaChain/chihuahua/utils/math"

type TickLiquidityKey interface {
	KeyMarshal() []byte
	PriceTakerToMaker() (priceTakerToMaker math_utils.PrecDec, err error)
}
