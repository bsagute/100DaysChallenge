list=[]
list[3:6]=[2,5,6]
# list.append(5)
# del list[0]
# del list[2:5]
# list.extend([5,5,68,8])
print(list[2])
del list[2]
print(list[:])
list.append([22,66,88,77,99])
list.extend([25,65,95,85,75,45])
print(list[:])
# list.insert(5)
#insert take the INDEX location and then ELEMENT 
print(list.index(95))# if not found will return the error 
list.insert(2,252)
print(list)
list.remove(25) # REmove the given element 
print(list)
list.pop()  # pop remove the last element 
list[3].sort()
print(list)
list.reverse()
print(list)
xn=list.copy()
print("sassaa",xn)
# remove pop clear index  count sort  reverse copy  

