package gifmath

import "math/big"

func LCM(a, b int) int {
	var an, bn, z big.Int
	an.SetInt64(int64(a))
	bn.SetInt64(int64(b))
	return int(z.Mul(z.Div(&bn, z.GCD(nil, nil, &an, &bn)), &an).Int64())
}
