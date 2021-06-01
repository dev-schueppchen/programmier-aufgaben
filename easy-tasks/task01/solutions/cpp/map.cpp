/*
	This example describes a way using the std::map
	library for saving username-password combinations
	and checking them.
*/

#include <iostream>
#include <string>
#include <map>

int main() {

	using namespace std;

	map<string, string> users = {
		{"root", "rootpw"}
	};

	string username, password;

	cout << "Enter username: ";
	cin >> username;

	cout << "Enter password: ";
	cin >> password;
 
	if (users[username] == password) {
		cout << "You are now logged in!";
		return 0;
	}

	cout << "Wrong password";
	return 1;
}
