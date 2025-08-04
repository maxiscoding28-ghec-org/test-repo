package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	// Hardcoded credentials
	password := "supersecret123"
	fmt.Println("Password is:", password)

	// Insecure hashing
	data := []byte("sensitive data")
	md5Sum := md5.Sum(data)
	sha1Sum := sha1.Sum(data)
	fmt.Printf("MD5: %x, SHA1: %x\n", md5Sum, sha1Sum)

	// Command injection
	userInput := "ls -al" // pretend this came from user input
	cmd := exec.Command("sh", "-c", userInput)
	out, _ := cmd.Output()
	fmt.Println(string(out))

	// Writing to world-writable file
	f, _ := os.OpenFile("/tmp/insecure.txt", os.O_CREATE|os.O_WRONLY, 0777)
	defer f.Close()
	io.WriteString(f, "This is unsafe\n")
}