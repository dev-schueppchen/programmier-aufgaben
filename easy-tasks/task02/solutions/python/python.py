# Solution without extra tasks

import random

number = random.randint(1, 101)
attempts = 10
user_input = -1

print("The number is between 1 and 100.")
print("You now have 10 tries to guess the number.")

while(attempts > 0):
    print("Remaining attemps: " + str(attempts))
    user_input = int(input("Please enter the number you would like to guess: "))
    attempts -= 1
    print("")
    if(user_input == number):
        print("You guessed the number right!")
        exit(0)
    elif(user_input < number):
        print("The number you entered was too low.")
    elif(user_input > number):
        print("The number you entered was too high.")

if(attempts <= 0):
    print("You have no more attempts left!")