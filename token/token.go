package token

import (
	"fmt"
	"os"
	"path/filepath"
)

// Get Token from ./token.txt
func GetToken() string {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// Define the filename
	filename := "token.txt"

	// Construct the file path
	filePath := filepath.Join(dir, filename)

	// Open the file with read permissions
	tokenBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(tokenBytes)
}

// Store token to file
func StoreToken(token string) error {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	// Define the filename
	filename := "token.txt"

	// Construct the file path
	filePath := filepath.Join(dir, filename)

	// Open the file with write permissions, create if it doesn't exist
	file, err := os.Create(filePath)
	if err != nil {
	    return err
	}
	defer file.Close()

	// Write the token to the file
    _, err = file.WriteString(token)
    if err != nil {
        fmt.Println(err)
        return nil
    }

	fmt.Printf("Token stored successfully in file: %s\n", filePath)
	return nil
}
