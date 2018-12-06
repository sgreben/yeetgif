package gifmath

import (
	"math"
	"math/big"
)

func LCM(a, b int) int {
	var an, bn, z big.Int
	an.SetInt64(int64(a))
	bn.SetInt64(int64(b))
	return int(z.Mul(z.Div(&bn, z.GCD(nil, nil, &an, &bn)), &an).Int64())
}

func RandomFloat32Signed(seed *int) float32 {
	*seed *= 16807
	return math.Float32frombits((uint32(*seed)>>9)|0x40000000) - 3.0
}

func RandomFloat32Unsigned(seed *int) float32 {
	*seed *= 16807
	return math.Float32frombits((uint32(*seed)>>9)|0x3f800000) - 1.0
}
