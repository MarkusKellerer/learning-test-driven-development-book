import unittest

from money import Money
from portfolio import Portfolio


class TestMoney(unittest.TestCase):

    def testMultiplication(self):
        tenEuros = Money(10, "EUR")
        twentyEuros = Money(20, "EUR")
        self.assertEqual(twentyEuros, tenEuros.times(2))

    def testDivision(self):
        originalMoney = Money(4002, "KRW")
        expectedMoneyAfterDivision = Money(1000.5, "KRW")
        self.assertEqual(expectedMoneyAfterDivision, originalMoney.divide(4))

    def testAddition(self):
        fiveDollars = Money(5, "USD")
        tenDollars = Money(10, "USD")
        fifteenDollars = Money(15, "USD")
        portfolio = Portfolio()
        portfolio.add(fiveDollars, tenDollars)
        self.assertEqual(fifteenDollars, portfolio.evaluate("USD"))

    def testAdditionDollarsAndEuros(self):
        fiveDollars = Money(5, "USD")
        tenEuros = Money(10, "EUR")
        expectedValue = Money(17, "USD")
        portfolio = Portfolio()
        portfolio.add(fiveDollars, tenEuros)
        actualValue = portfolio.evaluate("USD")
        self.assertEqual(expectedValue, actualValue, "%s != %s" % (expectedValue, actualValue))

    def testAdditonDollarsAndWon(self):
        oneDollar = Money(1, "USD")
        elevenHundretWon = Money(1100, "KRW")
        portfolio = Portfolio()
        portfolio.add(oneDollar, elevenHundretWon)
        expectedValue = Money(2200, "KRW")
        actualValue = portfolio.evaluate("KRW")
        self.assertEqual(expectedValue, actualValue, "%s != %s" % (expectedValue, actualValue))


if __name__ == "__main__":
    unittest.main()
