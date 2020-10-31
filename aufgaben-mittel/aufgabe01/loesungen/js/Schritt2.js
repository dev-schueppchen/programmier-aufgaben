const { getSmallestDivident } = require('./Schritt1');

const input = 1000;

const primeFactorizing = (num) => {
    let factors = [];
    let current = parseInt(num);
    let dividend;
    while (current > 1) {
        dividend = getSmallestDivident(parseInt(current));
        factors.push(dividend);
        current /= dividend;
    }
    return factors;
};

console.log(primeFactorizing(input));