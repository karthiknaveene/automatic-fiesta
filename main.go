package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os/exec"
)

// G101: hardcoded credentials
const password = "supersecret123"
const apiKey = "hardcoded-api-key-abc123"

func main() {
	// G404: weak random number
	fmt.Println(rand.Int())

	// G204: subprocess launched with variable
	userInput := "ls"
	cmd := exec.Command(userInput)
	cmd.Run()

	// G107: url provided to HTTP request as taint input
	url := "http://example.com"
	http.Get(url) //nolint

	// G501: import of weak crypto
	fmt.Println(password, apiKey)

	// G115: integer overflow (CWE-190)
	var i int32 = 2147483647
	var j uint32 = uint32(i)
	fmt.Println(j)
}

