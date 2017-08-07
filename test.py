way=input("Input way to file:")
f = open (way)
counter=0
for line in f:
    for letter in line:
        if letter =='\n':
            counter+=1
print(counter)
                              
