x=int(input("Enter the Number :- "))

sum=0
t=x
while x>0:
    temp=x
    rem =x%10
    print( rem**3)
    sum+=rem**3
    # print(sum)
    x=int(x/10)
if x==sum:
    print(x," Number is Armstrong number ")
else:
    print(x," Number is NOT Armstrong")