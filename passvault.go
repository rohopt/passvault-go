/*
Password Vault Implementation in Golang.

*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is the Password Vault")
	/*
		1. Get user inputs - help, add, get, update, list, init.
		2. If DB is not created, ask user to use init.
		3. init will create a new db entry. Create master user/password. and init the db
		4. add will allow user to add new entries. site, username, password, url.
		5. get will take input as site, and copy passwd to clipboard.
		6. list wil list all entries with site and usernames.
		additional notes -
		- user boltdb initially and store master creds and other creds in different buckets
		- future - store passwords in encrypted format in future
		- future - create ability to have expiring sessions so that user does not need to enter master password again and again.
	*/
}
