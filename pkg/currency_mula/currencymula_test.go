package currencymula

import (
	"fmt"
	"testing"

	"github.com/profe-ajedrez/decimal_mula/pkg/mathematics"
)

func TestCurrencyMulaCreate(t *testing.T) {
	var storerFabricString StorerFabric = func(value interface{}) (mathematics.Arithmetizable, error) {
		v := fmt.Sprint(value)
		return mathematics.DMArithmetizableFromString(v)
	}

	var storerFabricInt StorerFabric = func(value interface{}) (mathematics.Arithmetizable, error) {
		v := value.(int)
		return mathematics.DMArithmetizableFromString(fmt.Sprint(v))
	}

	dma1, _ := NewCurrency(1.1234, "CLP", 2, storerFabricString)
	dma1, _ = NewCurrency(1, "CLP", 2, storerFabricInt)
	dma2, _ := dma1.Clone()
	dma1 = dma2
}
