const assert = require('assert');
const Money = require('./money');
const Portfolio = require('./portfolio');

class MoneyTest {
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
        assert.deepStrictEqual(portfolio.evaluate("USD"), fifteenDollar)
    }

    testAdditionOfDollarsAndEuros() {
        let fiveDollars = new Money(5, "USD");
        let tenEuros = new Money(10, "EUR");
        let portfolio = new Portfolio();
        portfolio.add(fiveDollars, tenEuros)
        let expectedValue = new Money(17, "USD");
        assert.deepStrictEqual(portfolio.evaluate("USD"), expectedValue)
    }

    testAdditionOfDollarsAndWon() {
        let oneDollar = new Money(1, "USD");
        let elevenHundredWon = new Money(1100, "KRW");
        let portfolio = new Portfolio();
        portfolio.add(oneDollar, elevenHundredWon)
        let expectedValue = new Money(2200, "KRW");
        assert.deepStrictEqual(portfolio.evaluate("KRW"), expectedValue)
    }

    runAllTests() {
        let testMethods = this.getAllTestMethods()
        testMethods.forEach(m => {
            console.log("Running: %s()", m)
            let method = Reflect.get(this, m);
            try {

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