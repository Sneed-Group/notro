package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// Generate random 24-character alphanumeric string
	randStr := generateRandomString(24)

	// Construct Discord gift URL
	discordURL := "https://discord.gift/" + randStr

	// Send HTTP request to the URL
	resp, err := http.Get(discordURL)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Check if response body contains "invalid"
	if !strings.Contains(string(body), "invalid") {
		// Append URL and string to file
		appendToFile("notro.txt", discordURL)
		fmt.Println("URL appended to file successfully!")
	} else {
		fmt.Println("Response body contains 'invalid'.")
	}
}

// Function to generate a random string of given length
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Function to append content to a file
func appendToFile(filename, content string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(content + "\n"); err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}
}
