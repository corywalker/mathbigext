package mathbigext

import (
	"math"
	"math/big"
)

func Pow(b *big.Float, x *big.Float) *big.Float {
	// Needs special cases!
	if b.Cmp(big.NewFloat(0)) == -1 {
		if x.IsInt() {
			absb := big.NewFloat(0)
			absb.Set(b)
			absb.Abs(absb)
			exp := Log(absb)
			exp.Mul(exp, x)
			xint, _ := x.Int64()
			if xint % 2 == 0 {
				return Exp(exp)
			} else {
				return exp.Mul(Exp(exp, big.NewFloat(-1))
			}
		}
	}
	exp := Log(b)
	exp.Mul(exp, x)
	return Exp(exp)
}

func Exp(x *big.Float) *big.Float {
	result := big.NewFloat(1)
	cmpresult := x.Cmp(big.NewFloat(0))
	if cmpresult == 0 {
		return result
	} else if cmpresult == 1 {
		// Positive x
		var step int = 0
		for x.Cmp(big.NewFloat(float64(step+1))) == 1 {
			result.Mul(result, big.NewFloat(math.E))
			step = step + 1
		}
		xAsFloat, _ := x.Float64()
		residual := xAsFloat - float64(step)
		result.Mul(result, big.NewFloat(math.Exp(residual)))
		return result
	} else {
		// Negative x
		var step int = 0
		for x.Cmp(big.NewFloat(float64(step-1))) == -1 {
			result.Quo(result, big.NewFloat(math.E))
			step = step - 1
		}
		xAsFloat, _ := x.Float64()
		residual := float64(step) - xAsFloat
		result.Quo(result, big.NewFloat(math.Exp(residual)))
		return result
	}
}

func Log(x *big.Float) *big.Float {
	mant := big.NewFloat(0)
	exp := x.MantExp(mant)
	mant64, _ := mant.Float64()
	ret := math.Log(mant64) + float64(exp) * math.Ln2
	return big.NewFloat(ret)
}
