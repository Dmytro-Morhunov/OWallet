package models

type Wallet struct {
	ID           uint `gorm:"primarykey"`
	Name         string
	Balance      float64
	CurrencyType CryptoCurrencyAlias
}

type CryptoCurrencyAlias = string

type CryptoCurrencyType struct {
	Bitcoin  CryptoCurrencyAlias
	USDT     CryptoCurrencyAlias
	Ethereum CryptoCurrencyAlias
}
