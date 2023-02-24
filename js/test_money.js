const assert = require('assert');

class Money {

    constructor(amount, currency) {
        this.amount = amount
        this.currency = currency
    }

    times(multiplier) {
        return new Money(this.amount * multiplier, this.currency)
    }
    divide(divisor){
        return new Money(this.amount / divisor, this.currency)
    }
}

let fiveDollar = new Money(5, "USD");
let tenDollar = new Money(10, "USD");
assert.deepStrictEqual(fiveDollar.times(2), tenDollar);

let tenEuro = new Money(10, "EUR");
let twentyEuro = new Money(20, "EUR")
assert.deepStrictEqual(tenEuro.times(2), twentyEuro);

let originalMoney = new Money(4002, "KRW");
let actualMoneyAfterDivision = originalMoney.divide(4);
let expectedMoneyAfterDivision = new Money(1000.5, "KRW");
assert.deepStrictEqual(actualMoneyAfterDivision, expectedMoneyAfterDivision);
