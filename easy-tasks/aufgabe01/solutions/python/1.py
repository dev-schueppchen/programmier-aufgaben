userpws = {"root": "rootpw"}

print('User:')
user = input()

print('Password:')
password = input()

if (password == userpws.get(user)):
    print('You are now logged in')
else:
    print('Wrong password')
