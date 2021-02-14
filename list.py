lista = []
lista = ["A", 25, 2.25, ["AA", "CC", "DD"]]
print(lista)
for x in lista:
    if type(x) == list:
         print(x)
         if type(x)==list:
             for y in x:
                 print("ll",y)

print(type("AA"))