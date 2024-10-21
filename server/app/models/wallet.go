package models

type Wallet struct {
	ID           uint `gorm:"primarykey"`
	Name         string
	Balance      float64
	CurrencyType CryptoCurrencyAlias
}

type CryptoCurrencyAlias = int

const (
	Bitcoin CryptoCurrencyAlias = iota
	USDT
	Ethereum
)
