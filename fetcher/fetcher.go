package fetcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"stock-service/config"
	"strconv"
	"time"
)

// Fetcher represents a config.Fetcher.
type Fetcher struct {
	// APIKey should be an api key from the IEXCloud api
	APIKey string
	// RouteLimit represents the maximum amount of stocks to be fetched per route
	RouteLimit uint
}

// New returns a new instace of the fetcher
func New(apiKey string, routeLimit uint) (*Fetcher, error) {
	if apiKey == "" {
		return nil, errors.New("The apiKey must have a value")
	}

	if routeLimit == 0 {
		return nil, errors.New("The route limit should be bigger than 0 to initiate the fetcher")
	}

	f := new(Fetcher)
	f.APIKey = apiKey
	f.RouteLimit = routeLimit

	return f, nil
}

// FetchAndSave fetches the route data and the corresponding chart data and saves it in the StockStores defined in the env parameter.
func (f *Fetcher) FetchAndSave(ctx context.Context, env *config.Env, interval time.Duration) {
	// initial data fetch
	err := f.fetchAndSaveRoutes(ctx, env)
	if err != nil {
		log.Fatal(err)
	}

	// loop to fetch data every *interval* minutes
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	tc := ticker.C

	for {
		select {
		case <-tc:
			err := f.fetchAndSaveRoutes(ctx, env)
			if err != nil {
				log.Fatal(err)
			}
			break

		case <-ctx.Done():
			fmt.Println("Exiting fetch loop")
			break
		}
	}
}

func (f *Fetcher) fetchAndSaveRoutes(ctx context.Context, env *config.Env) error {
	lsLim := strconv.FormatUint(uint64(f.RouteLimit), 10)

	routes := []struct {
		route config.ShareOpt
		url   string
	}{
		{config.Gainers, "https://cloud.iexapis.com/stable/stock/market/list/gainers?token=" + f.APIKey + "&listLimit=" + lsLim},
		{config.Losers, "https://cloud.iexapis.com/stable/stock/market/list/losers?token=" + f.APIKey + "&listLimit=" + lsLim},
		{config.MostActive, "https://cloud.iexapis.com/stable/stock/market/list/mostactive?token=" + f.APIKey + "&listLimit=" + lsLim},
	}

	for _, r := range routes {
		err := f.fetchAndSaveData(ctx, env, r.route, r.url)
		if err != nil {
			return err
		}

		err = f.fetchAndSaveChartData(ctx, env, r.route)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *Fetcher) fetchAndSaveData(ctx context.Context, env *config.Env, opt config.ShareOpt, url string) error {
	var shares []config.StockInfo

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&shares)
	if err != nil {
		return err
	}

	switch opt {
	case config.Gainers:
		env.StockStore.SetOpt(config.Gainers, shares)

		break
	case config.Losers:
		env.StockStore.SetOpt(config.Losers, shares)

		break
	case config.MostActive:
		env.StockStore.SetOpt(config.MostActive, shares)

		break
	}

	return nil
}

func (f *Fetcher) fetchAndSaveChartData(ctx context.Context, env *config.Env, opt config.ShareOpt) error {
	symbols := make([]string, f.RouteLimit)

	switch opt {
	case config.Gainers:
		for i, body := range env.StockStore.Gainers() {
			symbols[i] = body.Symbol
		}

		break
	case config.Losers:
		for i, body := range env.StockStore.Losers() {
			symbols[i] = body.Symbol
		}

		break
	case config.MostActive:
		for i, body := range env.StockStore.MostActive() {
			symbols[i] = body.Symbol
		}

		break
	}

	for _, symbol := range symbols {
		err := f.fetchSymbolChart(ctx, env, opt, symbol)
		if err != nil {
			return err
		}

		time.Sleep(time.Microsecond)
	}

	return nil
}

func (f *Fetcher) fetchSymbolChart(ctx context.Context, env *config.Env, opt config.ShareOpt, symbol string) error {
	if symbol == "" {
		return nil
	}

	var chartNodes []config.StockChartNode

	url := "https://cloud.iexapis.com/stable/stock/" + symbol + "/chart/5d?token=" + f.APIKey + "&chartCloseOnly=true"

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		f.fetchSymbolChart(ctx, env, opt, symbol)
		time.Sleep(time.Microsecond)
	}

	err = json.NewDecoder(resp.Body).Decode(&chartNodes)
	if err != nil {
		return err
	}

	env.StockStore.SetOptChart(opt, symbol, chartNodes)

	return nil
}
