package currencymula

import (
	"github.com/profe-ajedrez/decimal_mula/pkg/mathematics"
)

type StorerFabric func(value interface{}) (mathematics.Arithmetizable, error)

type currency struct {
	isoName      string
	isoDecimals  int
	storerFabric StorerFabric
	amount       mathematics.Arithmetizable
}

func NewCurrency(value interface{}, isoName string, isoDecimals int, storerFabric StorerFabric) (currency, error) {
	arithm, err := storerFabric(value)
	c := currency{
		isoName:      isoName,
		isoDecimals:  isoDecimals,
		storerFabric: storerFabric,
		amount:       arithm,
	}
	return c, err
}

func currencyFromCurrency(c currency, amount mathematics.Arithmetizable) currency {
	return currency{
		isoName:      c.isoName,
		isoDecimals:  c.isoDecimals,
		storerFabric: c.storerFabric,
		amount:       amount,
	}
}

func (c *currency) Add(addings ...currency) currency {
	amount := c.amount
	for _, adding := range addings {
		amount = amount.Add(adding.amount)
	}
	return currencyFromCurrency(*c, amount)
}

func (c *currency) Substract(substractings ...currency) currency {
	amount := c.amount
	for _, substracting := range substractings {
		amount = amount.Add(substracting.amount)
	}
	return currencyFromCurrency(*c, amount)
}

func (c *currency) Multiply(factors ...currency) currency {
	amount := c.amount
	for _, factor := range factors {
		amount = amount.Add(factor.amount)
	}
	return currencyFromCurrency(*c, amount)
}

func (c *currency) Divide(divisors ...currency) currency {
	amount := c.amount
	for _, divisor := range divisors {
		amount = amount.Add(divisor.amount)
	}
	return currencyFromCurrency(*c, amount)
}

func (c *currency) Equals(cur currency) bool {
	return c.amount.Equals(cur.amount)
}

func (c *currency) Gt(cur currency) bool {
	return c.amount.Gt(cur.amount)
}

func (c *currency) Lt(cur currency) bool {
	return c.amount.Lt(cur.amount)
}

func (c *currency) Gte(cur currency) bool {
	return c.amount.Gte(cur.amount)
}

func (c *currency) Lte(cur currency) bool {
	return c.amount.Lte(cur.amount)
}

func (c *currency) Abs() currency {
	amountAbs := c.amount.Abs()
	return currencyFromCurrency(*c, amountAbs)
}

func (c *currency) String() string {
	return c.amount.String()
}

func (c *currency) PartsAsString() (string, string) {
	return c.amount.PartsAsString()
}

func (c *currency) GetClient() interface{} {
	amount, err := c.amount.Clone()
	if err == nil {
		return amount
	}
	panic(err)
}

func (c *currency) Clone() (currency, error) {
	amount, err := c.amount.Clone()
	return currencyFromCurrency(*c, amount), err
}
