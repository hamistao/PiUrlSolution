import math

C = 640320
C3_OVER_24 = C**3 // 24

def checkWord(word):
    if isPalindrome(word) and isPrime(int(word)):
        return True

def isPrime(n):
    if n == 2 or n == 3: return True
    if n < 2 or not (n&1): return False
    if n < 9: return True
    if n%3 == 0: return False
    r = int(n**0.5) + 1
    f = 5
    while f <= r:
        if n % f == 0: return False
        if n % (f+2) == 0: return False
        f += 6
    return True    

def isPalindrome(s):
    for i in range(len(s)//2 + 1):
        if s[i] != s[len(s)-1-i]:
            return False
    return True

def sqrt(n, one):
    floating_point_precision = 10**16
    n_float = float((n * floating_point_precision) // one) / floating_point_precision
    x = (int(floating_point_precision * math.sqrt(n_float)) * one) // floating_point_precision
    n_one = n * one
    while 1:
        x_old = x
        x = (x + n_one // x) // 2
        if x == x_old:
            break
    return x

def bs(a, b):
        if b - a == 1:
            if a == 0:
                Pab = Qab = 1
            else:
                Pab = (6*a-5)*(2*a-1)*(6*a-1)
                Qab = a*a*a*C3_OVER_24
            Tab = Pab * (13591409 + 545140134*a)
            if a & 1:
                Tab = -Tab
        else:
            m = (a + b) // 2
            Pam, Qam, Tam = bs(a, m)
            Pmb, Qmb, Tmb = bs(m, b)
            Pab = Pam * Pmb
            Qab = Qam * Qmb
            Tab = Qmb * Tam + Pam * Tmb
        return Pab, Qab, Tab

def chudnovsky_bs(digits):
    N = int(digits/14.181647462725477) + 1
    _, Q, T = bs(0, N)
    one = 10**digits
    sqrtC = sqrt(10005*one, one)
    return str((Q*426880*sqrtC) // T)

def sweepPi(initialExp, finalExp):
    index=23
    current="314159265358979323846"
    for log10_digits in range(initialExp, finalExp):
        digits = 10**log10_digits
        pi = chudnovsky_bs(digits)
        while index < len(pi):
            if checkWord(current): return current
            current = current[1:] + pi[index]
            index+=1
    return "not in the fist 10**8 numbers"

def test():
    print("moio")
    return True

if __name__ == "__main__":
    resp = sweepPi(9, 10)
    print(resp)
