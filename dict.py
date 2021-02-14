my_dict=dict({"name":"bhagwat",2:25,"ID":253})
print(my_dict["name"])
print(my_dict.get("ID"))
dt={}
dt={1:11,2:22,3:33}
del dt[3]
print(dt.get(1))
dt[30]=55
print(dt)
dt[30]=66
print(dt.pop(2),dt)
