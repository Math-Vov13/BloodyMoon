package fakedb

var fakedb = make(map[string]string)

// AddUser adds a user to the fake database
func AddUser(username, password string) bool {
	// Check if the user already exists
	if _, exists := fakedb[username]; exists {
		return false
	}
	fakedb[username] = password
	return true
}

// GetUser retrieves a user from the fake database
func GetUser(username, password string) bool {
	password_db := fakedb[username]
	return password_db == password
}
