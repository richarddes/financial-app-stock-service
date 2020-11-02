package handler

import (
	"encoding/json"
	"net/http"
	"stock-service/config"
)

// HandleLosers handles the /losers route.
// It returns a json response with a data field which includes the stock data
// and a chart field which includes the chart data for every stock entry in the data field.
func HandleLosers(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		losers := make(map[string]interface{}, 2)

		losers["data"] = env.StockStore.Losers()
		losers["chart"] = env.StockStore.LosersChart()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(losers)
	}
}
