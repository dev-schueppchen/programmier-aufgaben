userpws = {"root": "rootpw"}

print('User:')
user = input()

print('Password:')
password = input()

if (password == userpws.get(user)):
    print('Du bist nun eingeloggt')
else:
    print('Falsches Passwort')
