package main

import (
	"tdd/stocks"
	"testing"
)

func TestMultiplication(t *testing.T) {
	tenEuros := stocks.NewMoney(10, "EUR")
	actualResult := tenEuros.Times(2)
	expectedResult := stocks.NewMoney(20, "EUR")
	assertEqual(t, expectedResult, actualResult)
}

func TestDivision(t *testing.T) {
	originalMoney := stocks.NewMoney(4002, "KRW")
	actualMoneyAfterDivision := originalMoney.Divide(4)
	expectedMoneyAfterDivision := stocks.NewMoney(1000.5, "KRW")
	assertEqual(t, expectedMoneyAfterDivision, actualMoneyAfterDivision)
}

func TestAdditionInUsd(t *testing.T) {
	var portfolio stocks.Portfolio
	var portfolioInDollars stocks.Money

	fiveDollar := stocks.NewMoney(5, "USD")
	tenDollar := stocks.NewMoney(10, "USD")
	fifteenDollar := stocks.NewMoney(15, "USD")

	portfolio = portfolio.Add(fiveDollar)
	portfolio = portfolio.Add(tenDollar)
	portfolioInDollars, _ = portfolio.Evaluate("USD")

	assertEqual(t, fifteenDollar, portfolioInDollars)
}

func TestAdditionOfDollarsAndEuros(t *testing.T) {
	var portfolio stocks.Portfolio

	fiveDollars := stocks.NewMoney(5, "USD")
	tenEuros := stocks.NewMoney(10, "EUR")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenEuros)

	expectedValue := stocks.NewMoney(17, "USD")
	actualValue, _ := portfolio.Evaluate("USD")

	assertEqual(t, expectedValue, actualValue)
}

func TestAdditionOfDollarsAndWons(t *testing.T) {
	var portfolio stocks.Portfolio

	oneDollar := stocks.NewMoney(1, "USD")
	elevenHundretWon := stocks.NewMoney(1100, "KRW")

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(elevenHundretWon)

	expectedValue := stocks.NewMoney(2200, "KRW")
	actualValue, _ := portfolio.Evaluate("KRW")

	assertEqual(t, expectedValue, actualValue)
}

func TestAdditionWithMultipleMissingExchangeRates(t *testing.T) {
	var portfolio stocks.Portfolio

	oneDollar := stocks.NewMoney(1, "USD")
	oneEuro := stocks.NewMoney(1, "EUR")
	oneWon := stocks.NewMoney(1, "KRW")

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(oneEuro)
	portfolio = portfolio.Add(oneWon)

	expectedErrorMessage := "Missing exchange rate(s):[USD->Kalganid,EUR->Kalganid,KRW->Kalganid,]"
	_, actualError := portfolio.Evaluate("Kalganid")

	assertEqual(t, expectedErrorMessage, actualError.Error())
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %+v Got %+v", expected, actual)
	}
}
