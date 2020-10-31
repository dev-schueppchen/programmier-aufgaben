const num = 10;

const getSmallestDivident = (num) => {
    for (let i = 2; i < num + 1; i++) {
        if (num % i === 0) return i;
    }
};

module.exports = { getSmallestDivident };