const n = 10;

const fibonacci = n => {
    // Wir deklarieren 2 Variablen
    // um eine Art Cache zu haben.
    // Eine Variable speichert die letzte, eine die vorletzte
    // Fibonacci Zahl

    // Wir nutzten BigInt, da Fibonacci Zahlen irgenwann
    // das Limit der von JS erlaubten Größe einer Zahl übersteigt
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