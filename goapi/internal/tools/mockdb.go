package tools 

import (
	"time"	
)

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails{
	"adam": {
		AuthToken: "1234",
		Username: "adam",
	},
	"bob": {
		AuthToken: "5678",
		Username: "bob",
	},
	"charlie": {
		AuthToken: "91011",
		Username: "charlie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"adam": {
		Coins: 100,
		Username: "adam",
	},
	"bob": {
		Coins: 200,
		Username: "bob",
	},
	"charlie": {	
		Coins: 300,
		Username: "charlie",
	},
}

func (db *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate a database query
	time.Sleep(1 * time.Second)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate a database query
	time.Sleep(1 * time.Second)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) SetupDatabase() error {
	return nil
}