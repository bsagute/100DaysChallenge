x=int(input("Enter number one  "))
y=int(input("Enter number two  "))
print("Binary :- ",bin(x).replace("0b",""))
x<<=y

print("Binary :- ",bin(x).replace("0b",""))
print("Lefet <<<", x)
print("POWER  ",x^y)