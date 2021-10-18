package currencymula

import (
	"fmt"
	"testing"

	"github.com/profe-ajedrez/decimal_mula/pkg/mathematics"
)

var storerFabricString StorerFabric = func(value interface{}) (mathematics.Arithmetizable, error) {
	v := fmt.Sprint(value)
	return mathematics.DMArithmetizableFromString(v)
}

var storerFabricInt StorerFabric = func(value interface{}) (mathematics.Arithmetizable, error) {
	v := value.(int)
	return mathematics.DMArithmetizableFromString(fmt.Sprint(v))
}

func TestCurrencyMulaCreate(t *testing.T) {
	mathematics.Precision = 128
	dma1, _ := NewCurrency("54674571.98765432123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890", "CLP", 2, storerFabricString)
	dma2, _ := NewCurrency("1000.5789653789", "CLP", 2, storerFabricString)
	dma2 = dma1.Add(dma2)
	fmt.Println(dma2.String())
	dma1 = dma2
}

func TestCurrencyMulaMulaAdd(t *testing.T) {
	mathematics.Precision = 32
	dm1, _ := NewCurrency("1.123456789012345678901235", "CLP", 2, storerFabricString)
	dm2, _ := NewCurrency("-1.1", "CLP", 2, storerFabricString)
	dm3, _ := NewCurrency("1.202", "CLP", 2, storerFabricString)
	dm4, _ := NewCurrency("1.61", "CLP", 2, storerFabricString)

	sum := dm1.Add(dm2, dm3, dm4)
	want := "2.83545678901234567890123500000000"

	if want != sum.String() {
		t.Errorf("Expected %v, got %v at decimalMulaArithmetizable.Add()", want, sum.String())
	}
}

func TestCurrencyMulaMulaSubstract(t *testing.T) {
	mathematics.Precision = 32
	dm1, _ := NewCurrency("11.123456789012345678901235", "CLP", 2, storerFabricString)
	dm2, _ := NewCurrency("1.1", "CLP", 2, storerFabricString)
	dm3, _ := NewCurrency("-1.202", "CLP", 2, storerFabricString)
	dm4, _ := NewCurrency("1.61", "CLP", 2, storerFabricString)
	dm5, _ := NewCurrency("0.8989", "CLP", 2, storerFabricString)

	sub := dm1.Substract(dm2, dm3, dm4, dm5)
	want := "8.71655678901234567890123500000000"

	if want != sub.String() {
		t.Errorf("Expected %v, got %v at decimalMulaArithmetizable.Substract()", want, sub.String())
	}
}

func TestCurrencyMulaMulaMultiply(t *testing.T) {
	mathematics.Precision = 32
	dm1, _ := NewCurrency("11.123456789012345678901235", "CLP", 2, storerFabricString)
	dm2, _ := NewCurrency("1.1", "CLP", 2, storerFabricString)
	dm3, _ := NewCurrency("-1.202", "CLP", 2, storerFabricString)
	dm4, _ := NewCurrency("1.61", "CLP", 2, storerFabricString)
	dm5, _ := NewCurrency("0.8989", "CLP", 2, storerFabricString)

	mul := dm1.Multiply(dm2, dm3, dm4, dm5)
	want := "-21.28502582014299559803430038665700"

	if want != mul.String() {
		t.Errorf("Expected %v, got %v at decimalMulaArithmetizable.Multiply()", want, mul.String())
	}
}

func TestCurrencyMulaMulaDivide(t *testing.T) {
	mathematics.Precision = 32
	dm1, _ := NewCurrency("11.123456789012345678901235", "CLP", 2, storerFabricString)
	dm2, _ := NewCurrency("1.1", "CLP", 2, storerFabricString)
	dm3, _ := NewCurrency("-1.202", "CLP", 2, storerFabricString)
	dm4, _ := NewCurrency("1.61", "CLP", 2, storerFabricString)
	dm5, _ := NewCurrency("0.8989", "CLP", 2, storerFabricString)

	divi := dm1.Divide(dm2, dm3, dm4, dm5)
	want := "-5.81306745796531999297392368402250"

	if want != divi.String() {
		t.Errorf("Expected %v, got %v at decimalMulaArithmetizable.Multiply()", want, divi.String())
	}
}
