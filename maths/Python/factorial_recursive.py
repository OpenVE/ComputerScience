# factorial_recursive.py
#
# Factorial is the product of all positive
# integers less than or equal to a given number.
#
# by José Reyna @jobliz
#

def factorial(n):
    if n == 0:
        return 1
    else:
        temp = factorial(n - 1)
        return n * temp
