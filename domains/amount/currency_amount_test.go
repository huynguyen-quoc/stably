package amount

import (
	"testing"
)

func TestCurrencyAmountOperations(t *testing.T) {
	amount1 := NewCurrencyAmount("123.456789", "USD") // Example amount
	amount2 := NewCurrencyAmount("78.901234", "USD")  // Another example amount

	// Test Add
	sum := amount1.Add(amount2)
	expectedSumStr := "202.36"
	if sum.ToAmountString() != expectedSumStr {
		t.Errorf("Expected %s, but got %s", expectedSumStr, sum.ToAmountString())
	}

	// Test Subtract
	difference := amount1.Subtract(amount2)
	expectedDiffStr := "44.56"
	if difference.ToAmountString() != expectedDiffStr {
		t.Errorf("Expected %s, but got %s", expectedDiffStr, difference.ToAmountString())
	}

	// Test Mul
	product := amount1.Mul(amount2)
	expectedProdStr := "9740.89"
	if product.ToAmountString() != expectedProdStr {
		t.Errorf("Expected %s, but got %s", expectedProdStr, product.ToAmountString())
	}

	// Test Div
	quotient := amount1.Div(amount2)
	expectedQuotientStr := "1.56"
	if quotient.ToAmountString() != expectedQuotientStr {
		t.Errorf("Expected %s, but got %s", expectedQuotientStr, quotient.ToAmountString())
	}

	// Test Cmp
	cmpResult := amount1.Cmp(amount2)
	expectedCmpResult := 1 // amount1 > amount2
	if cmpResult != expectedCmpResult {
		t.Errorf("Expected %d, but got %d", expectedCmpResult, cmpResult)
	}

	// Test Percent
	percentResult := amount1.Percent("10")
	expectedPercentStr := "12.35"
	if percentResult.ToAmountString() != expectedPercentStr {
		t.Errorf("Expected %s, but got %s", expectedPercentStr, percentResult.ToAmountString())
	}
}
