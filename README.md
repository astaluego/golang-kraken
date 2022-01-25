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
- [ ] Get system status
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

- [ ] Add standard order
- [ ] Cancel open order
- [ ] Cancel all open orders
- [ ] Cancel all orders after

### Private user funding

- [ ] Get deposit methods
- [ ] Get deposit addresses
- [ ] Get status of recent deposits
- [ ] Get withdrawal information
- [ ] Withdraw funds
- [ ] Get status of recent withdrawals
- [ ] Request withdrawal cancelation
- [ ] Wallet Transfer

## References

 - [Kraken's API documentation](https://docs.kraken.com/rest/)

