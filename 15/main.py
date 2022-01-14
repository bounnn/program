import random
randomNumber = []
duplicateNumber = []
x = 0
# this loop random 6 number
for i in range(0,6):
    n = random.randint(0,9)#random number
    randomNumber.append(n)#insert number in randomNum
temp = randomNumber.copy()#ao vai kep khar default khrng randomNum
int_to_str = [str(int) for int in randomNumber]#convert int in list to string
resulte = "".join(int_to_str)
print("\nRandom number = ",resulte)
# this loop find duplicate number
for j in range(0,len(randomNumber)):
    if(randomNumber[j] == " "): continue
    for k in range(j+1,len(randomNumber)):
        if(randomNumber[k] == " "): continue
        if(randomNumber[j] == randomNumber[k]):#check duplicate number
            randomNumber[k] = " "#pen karn pien khar pen empty phuea br hai mun ma loop check ik hrp
            # break#vela jer to sum kr hai break lery
    duplicateNumber.append(randomNumber[j])
for j in duplicateNumber:
    x += duplicateNumber.count(j)
if x == 6: print("\tNo duplicate numbers.")
else: 
    for i in duplicateNumber:
        if temp.count(i) == 1: continue
        print("Number ",i," duplicates ",temp.count(i)," values.")
print("\n")