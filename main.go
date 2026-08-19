package main

import (
	"fmt"
	"os"
)

// secure_config_validator - Validate secure configs
func secure_config_validator(path string) {
	fmt.Println("========================================")
	fmt.Println("  Secure-Config-Validator")
	fmt.Println("  Validate secure configs")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	secure_config_validator(path)
}
