const assert = require('assert');
const Money = require('./money');
const Bank = require('./bank');
const Portfolio = require('./portfolio');

class MoneyTest {
    setup() {
        this.bank = new Bank();
        this.bank.addExchangeRate("EUR", "USD", 1.2)
        this.bank.addExchangeRate("USD", "KRW", 1100)
    }

    testMultiplication() {
        let tenEuro = new Money(10, "EUR");
        let twentyEuro = new Money(20, "EUR")
        assert.deepStrictEqual(tenEuro.times(2), twentyEuro);
    }

    testDivision() {
        let originalMoney = new Money(4002, "KRW");
        let actualMoneyAfterDivision = originalMoney.divide(4);
        let expectedMoneyAfterDivision = new Money(1000.5, "KRW");
        assert.deepStrictEqual(actualMoneyAfterDivision, expectedMoneyAfterDivision);
    }

    testAddition() {
        let fiveDollar = new Money(5, "USD");
        let tenDollar = new Money(10, "USD");
        let fifteenDollar = new Money(15, "USD");
        let portfolio = new Portfolio();
        portfolio.add(fiveDollar, tenDollar);
        assert.deepStrictEqual(portfolio.evaluate(new Bank(), "USD"), fifteenDollar)
    }

    testAdditionOfDollarsAndEuros() {
        let fiveDollars = new Money(5, "USD");
        let tenEuros = new Money(10, "EUR");
        let portfolio = new Portfolio();
        portfolio.add(fiveDollars, tenEuros)
        let expectedValue = new Money(17, "USD");
        assert.deepStrictEqual(portfolio.evaluate(this.bank, "USD"), expectedValue)
    }

    testAdditionOfDollarsAndWon() {
        let oneDollar = new Money(1, "USD");
        let elevenHundredWon = new Money(1100, "KRW");
        let portfolio = new Portfolio();
        portfolio.add(oneDollar, elevenHundredWon)
        let expectedValue = new Money(2200, "KRW");
        assert.deepStrictEqual(portfolio.evaluate(this.bank, "KRW"), expectedValue)
    }

    testAdditionWithMultipleMissingExchangeRates() {
        let oneDollar = new Money(1, "USD")
        let oneEuro = new Money(1, "EUR")
        let oneWon = new Money(1, "KRW")
        let portfolio = new Portfolio()
        portfolio.add(oneDollar, oneEuro, oneWon)
        let expectedError = new Error("Missing exchange rate(s):[USD->Kalganid,EUR->Kalganid,KRW->Kalganid]")
        assert.throws(() => portfolio.evaluate(this.bank, "Kalganid"), expectedError)
    }

    testConversionWithDifferentRatesBetweenTwoCurrencies() {
        let tenEuros = new Money(10, "EUR")
        assert.deepStrictEqual(this.bank.convert(tenEuros, "USD"), new Money(12, "USD"))
        this.bank.addExchangeRate("EUR", "USD", 1.3)
        assert.deepStrictEqual(this.bank.convert(tenEuros, "USD"), new Money(13, "USD"))
    }

    testConversionWithMissingExchangeRates() {
        let bank = new Bank();
        let tenEuros = new Money(10, "EUR");
        let expectedError = new Error("EUR->Kalganid")
        assert.throws(function () {
            bank.convert(tenEuros, "Kalganid")
        }, expectedError)
    }

    runAllTests() {
        let testMethods = this.getAllTestMethods()
        this.randomizeTestOrder(testMethods).forEach(m => {
            console.log("Running: %s()", m)
            let method = Reflect.get(this, m);
            try {
                this.setup()
                Reflect.apply(method, this, []);
            } catch (e) {
                if (e instanceof assert.AssertionError) {
                    console.log(e)
                } else {
                    throw e
                }
            }
        })
    }

    randomizeTestOrder(testMethods) {
        for (let i = testMethods.length - 1; i > 0; i--) {
            const j = Math.floor(Math.random() * (i + 1));
            [testMethods[i], testMethods[j]] = [testMethods[j], testMethods[i]]
        }
        return testMethods
    }

    getAllTestMethods() {
        let moneyPrototype = MoneyTest.prototype
        let allProps = Object.getOwnPropertyNames(moneyPrototype);
        let testMethods = allProps.filter(p => {
            return typeof moneyPrototype[p] === 'function' && p.startsWith('test')
        })
        return testMethods
    }
}

new MoneyTest().runAllTests();