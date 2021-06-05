const n = 10;

const fibonacci = (n) => {
    // We declare 2 variables
    // so we have some kind of cache.
    // One variable stores the last Fibonacci number,
    //  one stores the second to last Fibonacci number.

    // We use BigInt because Fibonacci numbers at some point
    // exceed the limit of the size of a number allowed by JS.
    let last = BigInt(1);
    let secondLast = BigInt(0);
    for (let _ = 0; _ < n; _++) {
        let temp = last;
        last = secondLast;
        secondLast = temp + last;
    }
    return secondLast;
};

module.exports = { fibonacci };
