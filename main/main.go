package main

import (
	"bufio"
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

		for {
			fmt.Print("cipher> ")

			if res := scan.Scan(); !res {
				break
			}

			line := scan.Text()
			pieces := strings.Split(line, " ")

			switch pieces[0] {
			// TODO add cases for each supported cipher
			default:
				fmt.Println("Error. Unknown cipher specified!")
				continue
			}
		}
	}
}
