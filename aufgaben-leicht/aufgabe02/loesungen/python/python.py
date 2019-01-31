# Lösung ohne Zusatzaufgaben

import random

zahl = random.randint(1, 101)
versuche = 10
eingabe = -1

print("Die Zahl befindet sich zwischen einschließlich 1 und 100.")
print("Du hast nun 10 Versuche, die Zahl zu raten.")

while(versuche > 0):
    print("Verbleibende Versuche: " + str(versuche))
    eingabe = int(input("Bitte gib die Zahl ein, die du erraten möchtest: "))
    versuche -= 1
    print("")
    if(eingabe == zahl):
        print("Du hast die Zahl richtig erraten!")
        exit(0)
    elif(eingabe < zahl):
        print("Deine eingegebene Zahl war zu niedrig.")
    elif(eingabe > zahl):
        print("Deine eingegebene Zahl war zu hoch.")

if(versuche <= 0):
    print("Du hast keine Versuche mehr übrig!")