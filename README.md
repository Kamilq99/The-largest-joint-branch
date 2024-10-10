A long break from programming didn't do me any favors, so I warmed up by calculating the greatest common divisor.

The **Greatest Common Divisor** (GCD) algorithm is a method used to find the largest positive integer that divides two or more integers without leaving a remainder. The most common algorithm for calculating the GCD is **Euclid's Algorithm**. It is efficient and works based on the principle that:

gcd(a,b)=gcd(b,amod  b)\\text{gcd}(a, b) = \\text{gcd}(b, a \\mod b)gcd(a,b)=gcd(b,amodb)

where aaa and bbb are two non-negative integers, and amod  ba \\mod bamodb is the remainder when aaa is divided by bbb. This process is repeated until b=0b = 0b=0, at which point the GCD is aaa.

### Steps of Euclid's Algorithm

1.  Given two numbers, aaa and bbb, where a≥ba \\geq ba≥b:
    
    *   If b=0b = 0b=0, then aaa is the GCD.
        
    *   Otherwise, replace aaa with bbb and bbb with amod  ba \\mod bamodb.
        
2.  Repeat step 1 until b=0b = 0b=0.
    
3.  The GCD is the value of aaa when b=0b = 0b=0.
    

### Example

Find the GCD of 56 and 98:

1.  a=98a = 98a=98, b=56b = 56b=56
    
2.  Apply the modulus: 98mod  56=4298 \\mod 56 = 4298mod56=42
    
3.  Now, a=56a = 56a=56, b=42b = 42b=42
    
4.  Apply the modulus: 56mod  42=1456 \\mod 42 = 1456mod42=14
    
5.  Now, a=42a = 42a=42, b=14b = 14b=14
    
6.  Apply the modulus: 42mod  14=042 \\mod 14 = 042mod14=0
    
7.  The GCD is 14, because when b=0b = 0b=0, a=14a = 14a=14.
    

This algorithm can also be extended to find the GCD of more than two numbers by applying it iteratively.

The standard tool to solve the problem will be the Go (Golang) programming language