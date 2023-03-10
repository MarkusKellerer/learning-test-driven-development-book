package main

import (
	"reflect"
	"tdd/stocks"
	"testing"
)

var bank stocks.Bank

func initializeExchangeRates() {
	bank = stocks.NewBank()
	bank.AddExchangeRate("EUR", "USD", 1.2)
	bank.AddExchangeRate("USD", "KRW", 1100)
}

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
	initializeExchangeRates()
	var portfolio stocks.Portfolio

	fiveDollar := stocks.NewMoney(5, "USD")
	tenDollar := stocks.NewMoney(10, "USD")
	fifteenDollar := stocks.NewMoney(15, "USD")

	portfolio = portfolio.Add(fiveDollar)
	portfolio = portfolio.Add(tenDollar)
	portfolioInDollars, err := portfolio.Evaluate(bank, "USD")

	assertNil(t, err)
	assertEqual(t, fifteenDollar, *portfolioInDollars)
}

func TestAdditionOfDollarsAndEuros(t *testing.T) {
	initializeExchangeRates()

	var portfolio stocks.Portfolio

	fiveDollars := stocks.NewMoney(5, "USD")
	tenEuros := stocks.NewMoney(10, "EUR")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenEuros)

	expectedValue := stocks.NewMoney(17, "USD")
	actualValue, err := portfolio.Evaluate(bank, "USD")

	assertNil(t, err)
	assertEqual(t, expectedValue, *actualValue)
}

func TestAdditionOfDollarsAndWons(t *testing.T) {
	initializeExchangeRates()

	var portfolio stocks.Portfolio

	oneDollar := stocks.NewMoney(1, "USD")
	elevenHundretWon := stocks.NewMoney(1100, "KRW")

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(elevenHundretWon)

	expectedValue := stocks.NewMoney(2200, "KRW")
	actualValue, err := portfolio.Evaluate(bank, "KRW")

	assertNil(t, err)
	assertEqual(t, expectedValue, *actualValue)
}

func TestAdditionWithMultipleMissingExchangeRates(t *testing.T) {
	initializeExchangeRates()

	var portfolio stocks.Portfolio

	oneDollar := stocks.NewMoney(1, "USD")
	oneEuro := stocks.NewMoney(1, "EUR")
	oneWon := stocks.NewMoney(1, "KRW")

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(oneEuro)
	portfolio = portfolio.Add(oneWon)

	expectedErrorMessage := "Missing exchange rate(s):[USD->Kalganid,EUR->Kalganid,KRW->Kalganid,]"
	value, actualError := portfolio.Evaluate(bank, "Kalganid")

	assertNil(t, value)
	assertEqual(t, expectedErrorMessage, actualError.Error())
}

func TestConversionWithDifferentRatesBetweenCurrencies(t *testing.T) {
	initializeExchangeRates()

	tenEuros := stocks.NewMoney(10, "EUR")
	actualConvertedMoney, err := bank.Convert(tenEuros, "USD")
	assertNil(t, err)
	assertEqual(t, stocks.NewMoney(12, "USD"), *actualConvertedMoney)
	bank.AddExchangeRate("EUR", "USD", 1.3)
	actualConvertedMoney, err = bank.Convert(tenEuros, "USD")
	assertNil(t, err)
	assertEqual(t, stocks.NewMoney(13, "USD"), *actualConvertedMoney)
}

func TestConversionWithMissingExchangeRate(t *testing.T) {
	bank := stocks.NewBank()
	tenEuros := stocks.NewMoney(10, "EUR")
	actualConvertedMoney, err := bank.Convert(tenEuros, "Kalganid")
	assertNil(t, actualConvertedMoney)
	assertEqual(t, "EUR->Kalganid", err.Error())
}

func assertNil(t *testing.T, actual interface{}) {
	if actual != nil && !reflect.ValueOf(actual).IsNil() {
		t.Errorf("Excpected error to be nil, found [%+v]", actual)
	}
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %+v Got %+v", expected, actual)
	}
}
