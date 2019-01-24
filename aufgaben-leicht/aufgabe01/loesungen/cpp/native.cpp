/*
	This example describes the 'native' way of solving this
	problem. So it does not use any external libraries (except
	iostream and string).
*/

#include <iostream>
#include <string>

struct User {
	std::string username;
	std::string password;
};

int main() {

	using namespace std;

	User users[] = {
		User{"root", "rootpw"},
	};

	string username, password;

	cout << "Enter username: ";
	cin >> username;

	cout << "Enter password: ";
	cin >> password;
 
	bool pass = false;
	for (int i = 0; i < sizeof(users) / sizeof(User); i++) {
		if (users[i].username == username && users[i].password == password) {
			pass = true;
		}
	}

	if (pass) {
		cout << "You are logged in.";
		return 0;
	}

	cout << "Wrong password";
	return 1;
}
