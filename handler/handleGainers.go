// Package handler implements all http handlers.
package handler

import (
	"encoding/json"
	"net/http"
	"stock-service/config"
)

// HandleGainers handles the /gainers route.
// It returns a json response with a data field which includes the stock data
// and a chart fieldwich incudes the chart data for every stock entry in the data field.
func HandleGainers(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gainers := make(map[string]interface{}, 2)

		gainers["data"] = env.StockStore.Gainers()
		gainers["chart"] = env.StockStore.GainersChart()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gainers)
	}
}
