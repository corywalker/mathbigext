package mathbigext

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math/big"
	"math"
)

func TestLog(t *testing.T) {
	// Test big input
	bigflt := big.NewFloat(4.657554276902127e+231)
	assert.Equal(t, "4.657554277e+231", bigflt.String())
	lnBigflt := Log(bigflt)
	assert.Equal(t, "533.435647", lnBigflt.String())

	// Test huge input
	hugeflt := big.NewFloat(0)
	hugeflt.Mul(bigflt, bigflt)
	assert.Equal(t, "2.169281184e+463", hugeflt.String())
	lnHugeflt := Log(hugeflt)
	assert.Equal(t, "1066.871294", lnHugeflt.String())

	// Test small input
	smallflt := big.NewFloat(4.657554276902127e-231)
	assert.Equal(t, "4.657554277e-231", smallflt.String())
	lnSmallflt := Log(smallflt)
	assert.Equal(t, "-530.358666", lnSmallflt.String())

	// Test tiny input
	tinyflt := big.NewFloat(0)
	tinyflt.Mul(smallflt, smallflt)
	assert.Equal(t, "2.169281184e-461", tinyflt.String())
	lnTinyflt := Log(tinyflt)
	assert.Equal(t, "-1060.717332", lnTinyflt.String())

	// Test E input
	lnE := Log(big.NewFloat(math.E))
	assert.Equal(t, "1", lnE.String())

	// Test zero input
	lnZeroflt := Log(big.NewFloat(0))
	assert.Equal(t, "-Inf", lnZeroflt.String())

	// Test invalid domain
	defer func() {
		if p, ok := recover().(big.ErrNaN); !ok {
			t.Errorf("got %v; want ErrNaN panic", p)
		}
	}()
	lnInvalidflt := Log(big.NewFloat(-.0001))
	// Should not reach here
	t.Errorf("got %s; want ErrNaN panic", lnInvalidflt)
}

func TestExp(t *testing.T) {
	// Test positive exponents
	res := Exp(big.NewFloat(2000.5))
	assert.Equal(t, "6.398984342e+868", res.String())
	res2 := Exp(big.NewFloat(0.001))
	assert.Equal(t, "1.0010005", res2.String())
	res3 := Exp(big.NewFloat(3))
	assert.Equal(t, "20.08553692", res3.String())

	// Test negative exponents
	res4 := Exp(big.NewFloat(-0.001))
	assert.Equal(t, "0.9990004998", res4.String())
	res5 := Exp(big.NewFloat(-2000.5))
	assert.Equal(t, "1.562748003e-869", res5.String())
	res6 := Exp(big.NewFloat(-3))
	assert.Equal(t, "0.04978706837", res6.String())
}

func TestPow(t *testing.T) {
	// Test large results
	res := Pow(big.NewFloat(2000.5), big.NewFloat(500))
	assert.Equal(t, "3.709179557e+1650", res.String())

	bigflt := big.NewFloat(4.657554276902127e+231)
	assert.Equal(t, "4.657554277e+231", bigflt.String())
	hugeflt := big.NewFloat(0)
	hugeflt.Mul(bigflt, bigflt)
	assert.Equal(t, "2.169281184e+463", hugeflt.String())
	res2 := Pow(hugeflt, big.NewFloat(2.5))
	assert.Equal(t, "2.191742975e+1158", res2.String())

	// Test small results
	res3:= Pow(big.NewFloat(2000.5), big.NewFloat(-500))
	assert.Equal(t, "2.696013996e-1651", res3.String())

	// Test negative bases
	assert.Equal(t, "-1", Pow(big.NewFloat(-1), big.NewFloat(1)).String())
	assert.Equal(t, "1", Pow(big.NewFloat(-1), big.NewFloat(2)).String())
	assert.Equal(t, "1", Pow(big.NewFloat(-2), big.NewFloat(0)).String())
	assert.Equal(t, "-2", Pow(big.NewFloat(-2), big.NewFloat(1)).String())
	assert.Equal(t, "4", Pow(big.NewFloat(-2), big.NewFloat(2)).String())
}

