# by Roberto Tatasciore (@robdiminished)

# Implementation of a fast exponentiation algorithm
# Source => http://www.johndcook.com/blog/2008/12/10/fast-exponentiation/

def tobinary(n):
    binary = ''
    while n > 0:
        binary = str(n % 2) + binary
        n = n // 2
    return binary

def pow(base, exp=2):
    if not exp: return 1
    power = base
    binary = tobinary(exp)
    for i in range(1, len(binary)):
        if binary[i] == '0': power *= power
        else: power = (power * power) * base
    return power
