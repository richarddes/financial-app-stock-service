// Package config defines globally used interfaces and structs.
package config

import (
	"context"
	"time"
)

// ShareOpt represents the different routes that can be retrieved.
type ShareOpt int

const (
	Gainers ShareOpt = iota
	Losers
	MostActive
)

const (
	// MaxRouteStocks defines the maximum amount of stocks to be fetchched per route.
	MaxRouteStocks = 40
)

type (
	// Env represents a collection of interfaces required for the handlers.
	Env struct {
		StockStore  StockStore
		StockClient Fetcher
	}

	// StockInfo represents the data fetched for each stocks.
	StockInfo struct {
		Symbol      string  `json:"symbol"`
		Change      float32 `json:"change"`
		LatestPrice float32 `json:"latestPrice"`
	}

	// StockChartNode represents the data for a node of a chart for a stock.
	StockChartNode struct {
		Date       string  `json:"date"`
		ClosePrice float32 `json:"close"`
	}
)

type (
	// StockStore defines the function needed for a datastore for stock/chart information.
	StockStore interface {
		SetOpt(opt ShareOpt, shs []StockInfo)
		SetOptChart(opt ShareOpt, symbol string, sis []StockChartNode)
		Gainers() []StockInfo
		Losers() []StockInfo
		MostActive() []StockInfo
		GainersChart() map[string][]StockChartNode
		LosersChart() map[string][]StockChartNode
		MostActiveChart() map[string][]StockChartNode
	}

	// Fetcher defines functions a stock fetcher has to implement.
	Fetcher interface {
		FetchAndSave(ctx context.Context, env *Env, interval time.Duration)
	}
)
