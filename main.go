package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/harisahmed05/gosh/cmd"
	"github.com/harisahmed05/gosh/prompt"
)

func main() {

	fmt.Println(`
       _        ____           _           
  ___ | |__    / ___| ___  ___| |__        
 / _ \| '_ \  | |  _ / _ \/ __| '_ \       
| (_) | | | | | |_| | (_) \__ \ | | |_ _ _ 
 \___/|_| |_|  \____|\___/|___/_| |_(_|_|_)

			simple Unix(-like) shell written in Go
 `)

	reader := bufio.NewReader(os.Stdin)
	for {
		user, err := prompt.InitUser()
		prompt.Show(user, err == nil)

		// Reading keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if err = cmd.Execute(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
