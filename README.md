# golang-kraken
 
Golang client for the Kraken API.

## Getting Started

### Installing

`go get github.com/astaluego/golang-kraken`

### Quick start

```go
package main

import (
    "fmt"
    kraken "github.com/astaluego/golang-kraken"
)

func main() {
    client := kraken.New()
    time, err := client.ServerTime()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(time)
}
```

## Supported calls

### Public market data

- [x] Get server time
- [x] Get system status
- [ ] Get asset info
- [ ] Get tradable asset pairs
- [ ] Get ticker information
- [ ] Get OHLC data
- [ ] Get order book
- [ ] Get recent trades
- [ ] Get recent spread data

### Private user data

- [ ] Get account balance
- [ ] Get trade balance
- [ ] Get open orders
- [ ] Get closed orders
- [ ] Query orders info
- [ ] Get trades history
- [ ] Query trades info
- [ ] Get open positions
- [ ] Get ledgers info
- [ ] Query ledgers
- [ ] Get trade volume
- [ ] Request export report
- [ ] Get export statuses
- [ ] Get export report
- [ ] Remove export report

### Private user trading

- [ ] Add order
- [ ] Cancel order
- [ ] Cancel all orders
- [ ] Cancel all orders after X

### Private user funding

- [ ] Get deposit methods
- [ ] Get deposit addresses
- [ ] Get status of recent deposits
- [ ] Get withdrawal information
- [ ] Withdraw funds
- [ ] Get status of recent withdrawals
- [ ] Request withdrawal cancelation
- [ ] Request wallet transfer

### Private user staking

- [ ] Stake asset
- [ ] Unstake asset
- [ ] List of stakeable assets
- [ ] Get pending staking transactions
- [ ] List of staking transactions

## Generated code

In the `generate` folder, you will find the source code to update `assets.go` and `asset_pairs.go`. Two calls on the Kraken API are made in order to get the list of the assets and asset pairs available on the plateform. Then the code is generated through the text/template feature of Golang.

The aim is to have a list of assets and asset pairs in Golang constants to simplify the usage of the library.

NB: `assets.go` and `asset_pairs.go` should not be manually edited. To update these files, run the `make generate` command.

## References

 - [Kraken's API documentation](https://docs.kraken.com/rest/)

