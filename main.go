package main

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Insecure TLS: skip certificate verification
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, _ := client.Get("https://expired.badssl.com")
	fmt.Println("Response:", resp.Status)

	// SQL injection
	userInput := os.Args[1]
	db, _ := sql.Open("mysql", "root:password@tcp(localhost:3306)/testdb")
	query := "SELECT * FROM users WHERE name = '" + userInput + "'"
	rows, _ := db.Query(query)
	defer rows.Close()
	fmt.Println("Executed insecure query")

	// Command injection
	cmd := exec.Command("sh", "-c", "ping -c 1 "+userInput)
	output, _ := cmd.CombinedOutput()
	fmt.Println(string(output))

	// Hardcoded secret
	secret := "aws_access_key=AKIAIOSFODNN7EXAMPLE"
	fmt.Println("Leaked secret:", secret)
}
