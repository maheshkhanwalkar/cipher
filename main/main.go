package main

import (
	"bufio"
	"cipher/types"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// General command format:
	//  ./cipher -t caesar -e -k key-value -i plaintext
	//  ./cipher -t caesar -d -k key-value -i ciphertext
	//  ./cipher -i (interactive mode)

	if len(os.Args) == 1 {
		fmt.Println("Error. No arguments specified. Use -i for interactive mode")
		os.Exit(1)
	}

	itv := flag.Bool("i", false, "interactive mode")
	flag.Parse()

	if *itv {
		fmt.Println("Entering interactive mode. [Use ^C or ^D to break loop]")

		scan := bufio.NewScanner(os.Stdin)
		types.SetupBuiltinCiphers()

		for {
			fmt.Print("cipher> ")

			if res := scan.Scan(); !res {
				break
			}

			line := scan.Text()
			pieces := strings.Split(line, " ")

			if len(pieces) != 4 {
				fmt.Println("Bad input -- expecting 4 components")
				fmt.Println("cipher-name action key data")

				continue
			}

			cipher := types.GetCipherByName(pieces[0])

			if cipher == nil {
				fmt.Println("Error. No cipher by that name supported")
				continue
			}

			if pieces[1] == "encrypt" {
				res := cipher.Encrypt(pieces[3], pieces[2])
				fmt.Println("Encrypted result: ", res)

			} else if pieces[1] == "decrypt" {
				res := cipher.Decrypt(pieces[3], pieces[2])
				fmt.Println("Decrypted result: ", res)
			} else {
				fmt.Println("Error. Unknown action specified")
				continue
			}
		}
	}
}
