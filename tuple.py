# tuple is immutable unlike list 
tp=(1,"balu","sagute",[2,5,8,74],1,1,1,1,1,1)
print(tp[-2])
print(tp[-3])
# print(tp[-1][-1])
#IF TUPLE HAS LIST , THEN IT CAN BE CHANGED 
tp[3][-1]=333
print(tp.index("sagute"))
print(tp.count(1))
print(1 in tp)
print(2 in tp)
print("balu" in tp)

