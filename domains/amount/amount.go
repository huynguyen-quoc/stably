package amount

import (
	"math/big"
)

type Amount struct {
	Amount *big.Int
}

func NewAmount(amountStr string) *Amount {
	// Convert the string amount to a big integer.
	amountBigInt, _ := new(big.Float).SetPrec(100).SetString(amountStr)
	coin := new(big.Float)
	// convert to big int and use to calculate to avoid rounding up down decimal
	coin.SetInt(getPow())

	amountBigInt.Mul(amountBigInt, coin)

	result := new(big.Int)
	amountBigInt.Int(result)

	return &Amount{
		Amount: result,
	}
}

func (c *Amount) Add(other *Amount) *Amount {
	sum := new(big.Int).Add(c.Amount, other.Amount)
	return &Amount{Amount: sum}
}

func (c *Amount) Subtract(other *Amount) *Amount {
	difference := new(big.Int).Sub(c.Amount, other.Amount)
	return &Amount{Amount: difference}
}

func (c *Amount) Mul(other *Amount) *Amount {
	difference := new(big.Int).Mul(c.Amount, other.Amount)
	difference.Quo(difference, big.NewInt(1).Exp(big.NewInt(10), big.NewInt(18), nil))
	return &Amount{Amount: difference}
}

func (c *Amount) Div(other *Amount) *Amount {
	num1 := new(big.Float).SetInt(c.Amount)
	num2 := new(big.Float).SetInt(other.Amount)
	difference := big.NewFloat(0).Quo(num1, num2)
	coin := new(big.Float)
	// convert to big int and use to calculate to avoid rounding up down decimal
	coin.SetInt(getPow())
	difference.Mul(difference, coin)
	result := new(big.Int)
	difference.Int(result)
	return &Amount{Amount: result}
}

func (c *Amount) Cmp(other *Amount) (r int) {
	difference := c.Amount.Cmp(other.Amount)
	return difference
}

func (c *Amount) Percent(percent string) *Amount {
	percentAmount := NewAmount(percent)
	multiply := getPow()
	oneHundred := big.NewInt(100)

	percentResult := new(big.Int).Mul(c.Amount, percentAmount.Amount)
	result := new(big.Int).Div(percentResult, new(big.Int).Mul(oneHundred, multiply))

	return &Amount{Amount: result}
}

func (c *Amount) ToDecimal() *big.Float {
	scaleFactor := big.NewInt(1).Exp(big.NewInt(10), big.NewInt(18), nil)
	amountFloat := big.NewFloat(0).SetInt(c.Amount)
	amountFloat.Quo(amountFloat, big.NewFloat(0).SetInt(scaleFactor))
	return amountFloat
}

func getPow() *big.Int {
	return big.NewInt(1).Exp(big.NewInt(10), big.NewInt(18), nil)
}
