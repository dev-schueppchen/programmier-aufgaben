def get_smallest_dividend(num):
    for i in range(2, num + 1):
        if num % i == 0:
            return i

def prime_factorizing(num):
    factors = []
    current = num
    while current > 1:
        dividend = get_smallest_dividend(int(current))
        factors.append(dividend)
        current /= dividend
    return factors

input = int(input('Number: '))
print(input, 'as prime factors is ', prime_factorizing(input))