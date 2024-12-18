package models

type Wallet struct {
	ID            uint `gorm:"primarykey"`
	Name          string
	Balance       float64
	Currency_Type CryptoCurrencyAlias
	UserID        User `gorm:"references:UserID"`
}

type CryptoCurrencyAlias = int

const (
	Bitcoin CryptoCurrencyAlias = iota
	USDT
	Ethereum
)
