// main_test.go

package main

import (
	"encoding/json"
	model "github.com/huy-nguyenquoc/stably/domains"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetUserIntegration(t *testing.T) {
	router := SetUpRouter() // Initialize your Gin router

	SetupRoutes(router)

	testCases := []struct {
		name             string
		transaction      *model.Transaction
		customerTier     string
		expectedStatus   int
		expectedAsset    string
		expectedProvider string
		expectedFee      string
	}{
		{"TestForACHFixFee", &model.Transaction{
			FromAmount:  "100",
			FromNetwork: "ACH",
			FromAsset:   "USD",
			ToNetwork:   "ethereum",
			ToAsset:     "ETH",
			FeeAsset:    "USD",
		}, "1", http.StatusOK, "USD", "Goose", "12.26"},
		{"TestForACHFixFeeCustomerTier2", &model.Transaction{
			FromAmount:  "100",
			FromNetwork: "ACH",
			FromAsset:   "USD",
			ToNetwork:   "ethereum",
			ToAsset:     "ETH",
			FeeAsset:    "USD",
		}, "2", http.StatusOK, "USD", "Goose", "9.27"},
		{"TestForACHFixFeeCustomerTier3", &model.Transaction{
			FromAmount:  "100",
			FromNetwork: "ACH",
			FromAsset:   "USD",
			ToNetwork:   "ethereum",
			ToAsset:     "ETH",
			FeeAsset:    "USD",
		}, "3", http.StatusOK, "USD", "Goose", "7.78"},
		{"TestForACHPercentFee", &model.Transaction{
			FromAmount:  "320",
			FromNetwork: "ACH",
			FromAsset:   "USD",
			ToNetwork:   "ethereum",
			ToAsset:     "ETH",
			FeeAsset:    "USD",
		}, "1", http.StatusOK, "USD", "Goose", "14.12"},
		{"TestForWire", &model.Transaction{
			FromAmount:  "120",
			FromNetwork: "Wire",
			FromAsset:   "USD",
			ToNetwork:   "ethereum",
			ToAsset:     "ETH",
			FeeAsset:    "USD",
		}, "1", http.StatusOK, "USD", "Goose", "35.26"},
		{"TestForWireCustomerTier3", &model.Transaction{
			FromAmount:  "120",
			FromNetwork: "Wire",
			FromAsset:   "USD",
			ToNetwork:   "ethereum",
			ToAsset:     "ETH",
			FeeAsset:    "USD",
		}, "3", http.StatusOK, "USD", "Goose", "7.84"},
		{"TestForCard", &model.Transaction{
			FromAmount:  "120",
			FromNetwork: "Card",
			FromAsset:   "USD",
			ToNetwork:   "ethereum",
			ToAsset:     "ETH",
			FeeAsset:    "USD",
		}, "1", http.StatusOK, "USD", "Goose", "13.92"},
		{"TestForUnSupported", &model.Transaction{
			FromAmount:  "120",
			FromNetwork: "Card2",
			FromAsset:   "USD",
			ToNetwork:   "ethereum",
			ToAsset:     "ETH",
			FeeAsset:    "USD",
		}, "1", 500, "", "", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			queryParams := url.Values{}
			queryParams.Set("fromAmount", tc.transaction.FromAmount)
			queryParams.Set("fromNetwork", tc.transaction.FromNetwork)
			queryParams.Set("fromAsset", tc.transaction.FromAsset)
			queryParams.Set("toNetwork", tc.transaction.ToNetwork)
			queryParams.Set("toAsset", tc.transaction.ToAsset)
			queryParams.Set("feeAsset", tc.transaction.FeeAsset)

			req, err := http.NewRequest("GET", "v1/fee?"+queryParams.Encode(), nil)
			req.Header.Set("X-Customer-Tier", tc.customerTier)
			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			if rec.Code != tc.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", tc.expectedStatus, rec.Code)
			}
			if rec.Code != http.StatusInternalServerError {
				var responseBody map[string]interface{}
				err = json.Unmarshal(rec.Body.Bytes(), &responseBody)
				if err != nil {
					t.Fatal(err)
				}

				expectedMessageAsset := tc.expectedAsset
				actualMessage, ok := responseBody["asset"].(string)
				if expectedMessageAsset != "" && (!ok || actualMessage != expectedMessageAsset) {
					t.Errorf("Expected asset %s, but got %s", expectedMessageAsset, actualMessage)
				}
				expectedMessageFee := tc.expectedFee
				actualMessageFee, ok := responseBody["fee"].(string)
				if expectedMessageFee != "" && (!ok || actualMessageFee != expectedMessageFee) {
					t.Errorf("Expected fee %s, but got %s", expectedMessageFee, actualMessageFee)

				}
				expectedMessageProvider := tc.expectedProvider
				actualMessageProvider, ok := responseBody["provider"].(string)
				if expectedMessageProvider != "" && (!ok || actualMessageProvider != expectedMessageProvider) {
					t.Errorf("Expected provider %s, but got %s", expectedMessageProvider, actualMessageProvider)
				}
			}
		})
	}

}
