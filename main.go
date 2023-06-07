package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

func main() {
	eFlag := flag.Bool("e", false, "description of e flag")
	dFlag := flag.Bool("d", false, "description of d flag")
	uFlag := flag.Bool("u", false, "description of u flag")

	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		// Get the last argument
		lastArg := args[len(args)-1]

		if *eFlag {
			// Base64 (URL) encode the last argument
			var encodedArg string
			if *uFlag {
				encodedArg = base64.RawURLEncoding.EncodeToString([]byte(lastArg))
			} else {
				encodedArg = base64.StdEncoding.EncodeToString([]byte(lastArg))
			}
			fmt.Println(encodedArg)
		} else if *dFlag {
			// Base64 (URL) decode the last argument
			var decodedArg []byte
			var err error
			if *uFlag {
				decodedArg, err = base64.RawURLEncoding.DecodeString(lastArg)
			} else {
				decodedArg, err = base64.StdEncoding.DecodeString(lastArg)
			}
			if err != nil {
				fmt.Println("Error: Invalid base64 string.")
				os.Exit(1)
			}
			fmt.Println(string(decodedArg))
		}
	} else {
		fmt.Println("Error: No string provided.")
		os.Exit(1)
	}
}
