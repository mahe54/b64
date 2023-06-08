package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	eFlag := flag.Bool("e", false, "encode")
	dFlag := flag.Bool("d", false, "decode")
	uFlag := flag.Bool("u", false, "url encode / decode")

	flag.Parse()

	var lastArg string
	args := flag.Args()
	if len(args) > 0 {
		lastArg = args[len(args)-1]
	} else {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		lastArg = strings.TrimSuffix(input, "\n")
	}

	if *eFlag {
		var encodedArg string
		if *uFlag {
			encodedArg = base64.RawURLEncoding.EncodeToString([]byte(lastArg))
		} else {
			encodedArg = base64.StdEncoding.EncodeToString([]byte(lastArg))
		}
		fmt.Println(encodedArg)
	} else if *dFlag {
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
	} else {
		fmt.Println("Error: No operation specified.")
		os.Exit(1)
	}
}
