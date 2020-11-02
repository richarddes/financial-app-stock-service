package handler

import (
	"encoding/json"
	"net/http"
	"stock-service/config"
)

// HandleMostActive handles the /most-active route.
// It returns a json response with a data field which includes the stock data
// and a chart fieldwich incudes the chart data for every stock entry in the data field.
func HandleMostActive(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mostActive := make(map[string]interface{}, 2)

		mostActive["data"] = env.StockStore.MostActive()
		mostActive["chart"] = env.StockStore.MostActiveChart()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mostActive)
	}
}
