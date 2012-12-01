# pow.py
#
# Powers a base number by a given exponent.
# This implementation uses divide and conquer.
#
# by Roberto Tatasciore @robdiminished
#

def pow(base, exp=2):
    if not exp:
        return 1
    half = pow(base, exp // 2)
    if not exp % 2:
        return half * half
    else:
        return half * half * base
