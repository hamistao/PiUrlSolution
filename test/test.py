import decimal
import sys

C = 426880 * decimal.Decimal(10005).sqrt()
k = 6
m = 1
x = 1
l = 13591409
s = l

def compute_pi(i):
    """
    This function calculates the value of pi to 'i'\ 
    number of places using the previously
    defined funtions
    Args:
    i:   precision
    Returns:
    pi:   the value of pi"""
    decimal.getcontext().prec = i+1
    decimal.getcontext().Emax = 999999999
    a = list(map(compute_pi, range(1, i+1)))
    pi = decimal.Decimal(C/a[-1])
    return pi

n = int(input("Please type number: "))
sys.setrecursionlimit(n)
b=compute_pi(n)
print(b)