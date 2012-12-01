import math

# by phanghos, following https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
def erath_sieve(n):
    primes = [True for i in xrange(n + 1)]
    primes[0] = primes[1] = False
    limit = int(math.sqrt(n))
    
    for i in xrange(2, limit + 1):
        if primes[i]:
            for j in xrange(i * i, n + 1, i):
                primes[j] = False
    
    return primes

def erath_numbers(n):
    sieve = erath_sieve(n)
    return [i for i in xrange(len(sieve)) if sieve[i]]

def isprime(n):
    if n == 2: return True
    if not n % 2 or n == 1: return False
    limit = int(math.sqrt(n))
    for i in range(3, limit + 1, 2):
        if not n % i: return False
    return True
