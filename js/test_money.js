const assert = require('assert');

class Money {

    constructor(amount, currency) {
        this.amount = amount
        this.currency = currency
    }

    times(multiplier) {
        return new Money(this.amount * multiplier, this.currency)
    }
}

let fiver = new Money(5, "USD");
let tenner = fiver.times(2);

assert.strictEqual(tenner.amount, 10);
assert.strictEqual(tenner.currency, "USD");

let tenEuro = new Money(10, "EUR");
let twentyEuro = tenEuro.times(2)
assert.strictEqual(twentyEuro.amount, 20);
assert.strictEqual(twentyEuro.currency, "EUR");