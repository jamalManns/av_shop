package models

type Item struct {
    Type     string `json:"type"`
    Quantity int    `json:"quantity"`
}

type CoinHist struct {
    Received []Transaction `json:"received"`
    Sent     []Transaction `json:"sent"`
}

type InfoResponse struct {
    Coins       int      `json:"coins"`
    Inventory   []Item   `json:"inventory"`
    CoinHistory CoinHist `json:"coinHistory"`
}
