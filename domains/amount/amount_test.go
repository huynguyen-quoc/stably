package amount

import (
	"testing"
)

func TestAmountOperations(t *testing.T) {
	amount1 := NewAmount("123.456789") // Example amount
	amount2 := NewAmount("78.901234")  // Another example amount

	// Test Add
	sum := amount1.Add(amount2)
	expectedSumStr := "202.358023"
	if sum.ToDecimal().Text('f', 6) != expectedSumStr {
		t.Errorf("Expected %s, but got %s", expectedSumStr, sum.ToDecimal().Text('f', 6))
	}

	// Test Subtract
	difference := amount1.Subtract(amount2)
	expectedDiffStr := "44.555555"
	if difference.ToDecimal().Text('f', 6) != expectedDiffStr {
		t.Errorf("Expected %s, but got %s", expectedDiffStr, difference.ToDecimal().Text('f', 6))
	}

	// Test Mul
	product := amount1.Mul(amount2)
	expectedProdStr := "9740.892998"
	if product.ToDecimal().Text('f', 6) != expectedProdStr {
		t.Errorf("Expected %s, but got %s", expectedProdStr, product.ToDecimal().Text('f', 6))
	}

	// Test Div
	quotient := amount1.Div(amount2)
	expectedQuotientStr := "1.564700"
	if quotient.ToDecimal().Text('f', 6) != expectedQuotientStr {
		t.Errorf("Expected %s, but got %s", expectedQuotientStr, quotient.ToDecimal().Text('f', 6))
	}

	// Test Cmp
	cmpResult := amount1.Cmp(amount2)
	expectedCmpResult := 1 // amount1 > amount2
	if cmpResult != expectedCmpResult {
		t.Errorf("Expected %d, but got %d", expectedCmpResult, cmpResult)
	}

	// Test Percent
	percentResult := amount1.Percent("10")
	expectedPercentStr := "12.345679"
	if percentResult.ToDecimal().Text('f', 6) != expectedPercentStr {
		t.Errorf("Expected %s, but got %s", expectedPercentStr, percentResult.ToDecimal().Text('f', 6))
	}
}
