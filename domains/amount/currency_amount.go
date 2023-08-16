package amount

import "math/big"

type CurrencyAmount struct {
	Amount   *Amount
	Currency string
}

var currencyScaleFactors = map[string]int{
	"USD": 2,
	"EUR": 2,
	// Add more currencies and their scale factors as needed
}

func NewCurrencyAmount(amountStr string, currency string) *CurrencyAmount {
	amount := NewAmount(amountStr)
	return &CurrencyAmount{
		Amount:   amount,
		Currency: currency,
	}
}

func (c *CurrencyAmount) Add(other *CurrencyAmount) *CurrencyAmount {
	result := c.Amount.Add(other.Amount)
	return &CurrencyAmount{Amount: result, Currency: c.Currency}
}

func (c *CurrencyAmount) Subtract(other *CurrencyAmount) *CurrencyAmount {
	result := c.Amount.Subtract(other.Amount)
	return &CurrencyAmount{Amount: result, Currency: c.Currency}
}

func (c *CurrencyAmount) Mul(other *CurrencyAmount) *CurrencyAmount {
	result := c.Amount.Mul(other.Amount)
	return &CurrencyAmount{Amount: result, Currency: c.Currency}
}

func (c *CurrencyAmount) Div(other *CurrencyAmount) *CurrencyAmount {
	result := c.Amount.Div(other.Amount)
	return &CurrencyAmount{Amount: result, Currency: c.Currency}
}

func (c *CurrencyAmount) Cmp(other *CurrencyAmount) (r int) {
	difference := c.Amount.Cmp(other.Amount)
	return difference
}

func (c *CurrencyAmount) Percent(percent string) *CurrencyAmount {
	result := c.Amount.Percent(percent)
	return &CurrencyAmount{Amount: result, Currency: c.Currency}
}

func (c *CurrencyAmount) ToAmountDecimal() *big.Float {
	return c.Amount.ToDecimal()
}

func (c *CurrencyAmount) ToAmountString() string {
	return c.Amount.ToDecimal().Text('f', currencyScaleFactors[c.Currency])
}
