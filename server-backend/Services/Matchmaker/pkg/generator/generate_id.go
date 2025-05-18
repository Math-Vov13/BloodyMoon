package generator

import "math/rand"

func GenerateID(length int) string {
	// Define the characters to use for the ID
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Create a byte slice to hold the generated ID
	id := make([]byte, length)

	// Generate a random ID
	for i := range id {
		id[i] = charset[rand.Intn(len(charset))]
	}

	return string(id)
}
