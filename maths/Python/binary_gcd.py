def binary_gcd(n1, n2):
    if n1 == n2: return n1
    if not n1 & 1 and not n2 & 1: return binary_gcd(n1 >> 1, n2 >> 1) << 1
    if not n1 & 1 and n2 & 1: return binary_gcd(n1 >> 1, n2)
    if n1 & 1 and not n2 & 1: return binary_gcd(n1, n2 >> 1)
    if n1 & 1 and n2 & 1:
        if n1 >= n2: return binary_gcd((n1 - n2) >> 1, n2)
        return binary_gcd((n2 - n1) >> 1, n1)