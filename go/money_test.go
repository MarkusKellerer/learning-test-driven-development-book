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
	portfolioInDollars = portfolio.Evaluate("USD")

	assertEqual(t, fifteenDollar, portfolioInDollars)
}

func TestAdditionOfDollarsAndEuros(t *testing.T) {
	var portfolio stocks.Portfolio

	fiveDollars := stocks.NewMoney(5, "USD")
	tenEuros := stocks.NewMoney(10, "EUR")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenEuros)

	expectedValue := stocks.NewMoney(17, "USD")
	actualValue := portfolio.Evaluate("USD")

	assertEqual(t, expectedValue, actualValue)
}

func assertEqual(t *testing.T, expected stocks.Money, actual stocks.Money) {
	if expected != actual {
		t.Errorf("Expected %+v Got %+v", expected, actual)
	}
}
