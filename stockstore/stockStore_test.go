package stockstore_test

import (
	"reflect"
	"stock-service/config"
	"stock-service/stockstore"
	"testing"
)

func ShareStoreSuite(t *testing.T, impl config.StockStore) {
	t.Run("test that endpoints return empty slices with no data when initialized", func(t *testing.T) {
		gs := impl.Gainers()
		ls := impl.Losers()
		mas := impl.Losers()

		if len(gs) > 0 {
			t.Fatal("The gainers slice already had values in it when it shouldn't have had")
		}

		if len(ls) > 0 {
			t.Fatal("The losers slice already had values in it when it shouldn't have had")
		}

		if len(mas) > 0 {
			t.Fatal("The most-active slice already had values in it when it shouldn't have had")
		}
	})

	t.Run("test endpoint data replacement", func(t *testing.T) {
		info := []config.StockInfo{
			config.StockInfo{Symbol: "AAPl", LatestPrice: 300.02, Change: 0.02},
			config.StockInfo{Symbol: "MSFT", LatestPrice: 250.22, Change: 1.45},
			config.StockInfo{Symbol: "BMW", LatestPrice: 30.54, Change: -1.54},
			config.StockInfo{Symbol: "VW", LatestPrice: 45.32, Change: -0.24},
			config.StockInfo{Symbol: "FB", LatestPrice: 230.00, Change: -1.00},
			config.StockInfo{Symbol: "SNAP", LatestPrice: 23.45, Change: 1.02},
		}

		impl.SetOpt(config.Gainers, info[:2])

		if !reflect.DeepEqual(impl.Gainers(), info[:2]) {
			t.Error("The data returned by the gainers endpoint doesn't match the specified data")
		}

		impl.SetOpt(config.Losers, info[2:4])

		if !reflect.DeepEqual(impl.Losers(), info[2:4]) {
			t.Error("The data returned by the gainers endpoint doesn't match the specified data")
		}

		impl.SetOpt(config.MostActive, info[4:6])

		if !reflect.DeepEqual(impl.MostActive(), info[4:6]) {
			t.Error("The data returned by the gainers endpoint doesn't match the specified data")
		}
	})
}

func TestDefaultImpl(t *testing.T) {
	ss, err := stockstore.New(10)
	if err != nil {
		t.Fatal(err)
	}

	ShareStoreSuite(t, ss)
}
