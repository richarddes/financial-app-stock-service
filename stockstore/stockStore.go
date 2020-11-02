// Package stockstore defines a datastore for storing stock(-chart) data.
package stockstore

import (
	"errors"
	"stock-service/config"
)

// StockStore saves the stock and chart data fir each route,
type StockStore struct {
	config.StockStore
	gainers, losers, mostActive                []config.StockInfo
	gainersChart, losersChart, mostActiveChart map[string][]config.StockChartNode
}

// New returns a newinstace of the StockStore with the cao representing the
// maximum amount of stocks that can be saved for each route.
func New(cap int) (*StockStore, error) {
	if cap < 1 {
		return nil, errors.New("The capacity cannot be smaller than 1")
	}

	ss := new(StockStore)
	ss.gainers = make([]config.StockInfo, 0, cap)
	ss.losers = make([]config.StockInfo, 0, cap)
	ss.mostActive = make([]config.StockInfo, 0, cap)

	ss.gainersChart = make(map[string][]config.StockChartNode, cap)
	ss.losersChart = make(map[string][]config.StockChartNode, cap)
	ss.mostActiveChart = make(map[string][]config.StockChartNode, cap)

	return ss, nil
}

// SetOpt replaces the data for the route specified in opt to values in shs.
func (ss *StockStore) SetOpt(opt config.ShareOpt, shs []config.StockInfo) {
	switch opt {
	case config.Gainers:
		ss.gainers = shs
		break
	case config.Losers:
		ss.losers = shs
		break
	case config.MostActive:
		ss.mostActive = shs
		break
	}
}

// SetOptChart replaces the chart data for a symol in the route specified in opt to the values in scns.
func (ss *StockStore) SetOptChart(opt config.ShareOpt, symbol string, scns []config.StockChartNode) {
	switch opt {
	case config.Gainers:
		ss.gainersChart[symbol] = scns
		break
	case config.Losers:
		ss.losersChart[symbol] = scns
		break
	case config.MostActive:
		ss.mostActiveChart[symbol] = scns
		break
	}
}

// Gainers returns the gainers saved in the stock store
func (ss *StockStore) Gainers() []config.StockInfo {
	return ss.gainers
}

// Losers returns the losers saved in the stock store
func (ss *StockStore) Losers() []config.StockInfo {
	return ss.losers
}

// MostActive returns the most active stocks saved in the stock store
func (ss *StockStore) MostActive() []config.StockInfo {
	return ss.mostActive
}

// GainersChart returns the chart data belonging to the gainers in the stock store
func (ss *StockStore) GainersChart() map[string][]config.StockChartNode {
	return ss.gainersChart
}

// LosersChart returns the chart data belonging to the losers in the stock store
func (ss *StockStore) LosersChart() map[string][]config.StockChartNode {
	return ss.losersChart
}

// MostActiveChart returns the chart data belonging to the most active stocks in the stock store
func (ss *StockStore) MostActiveChart() map[string][]config.StockChartNode {
	return ss.mostActiveChart
}
