# Stock-Service
This service fetches stock data from the *IEX Cloud* service and saves them in an in-memory datastore. Furhermore, it returns the saved data.
It does NOT handle stock buying/selling.

## Setup 
To run this service independently, you have to install all necessary dependencies, like so:
```sh
go get ./...
```
To run the app, run:
```sh
go run main.go
```

## Inner workings
This service fetches the latest *gainers*, *losers*, and *most active* stocks from the *IEX Cloud* service every 15 minutes. It fetches 40 stocks per route with their corresponding chart data (historical data of the last 5 days) and saves themin an in-memory datastore. All 3 available routes (*/api/stocks/(gainers | stocks | most-active)*) return their data in the following format: 
**{data: [{symbol, change, latestPrice}], chart: []{symbol: [{date, closePrice}]}}**

## Contributing 
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
MIT License. Click [here](https://choosealicense.com/licenses/mit/) or see the LICENSE file for details.