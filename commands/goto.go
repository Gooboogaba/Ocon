package commands

import (
    "fmt"
    "oscarsgoofysite/OCON/state"
	"strconv"
	"log"
)

func GotoCmd(args []string) {
    //fmt.Println("going to the line at: " + args[0])

    if len(args) == 0 {//if the goto has NO args eg (goto)
		fmt.Println("no Args for goto")
        return
    }

    wereto := args[0]
    if wereto[0] == '\'' {//gotos a line number
		n, err := strconv.Atoi(wereto[1:]) // converts the string into a number were 1: (as strings are arrys of chars) uis everytihng past the "  
		if err != nil {//if it aculy is NOT a number then leave
			fmt.Println("invalid number:" + wereto[1:])
			return
		}
		state.Pointer = n - 1 //-1 becuse after this the poiter adds 1
	} else if wereto[0] == '"' { //goses to a sec, section, or, § (all the same thing but i digress)
		userSec := wereto[1:]//this is everything past the "
		list := state.SectionList
		sec, ok := list[userSec]//this is an int (map[string]int) of were the pointer needs to go
		if !ok {
			log.Fatalf("section %q not found", userSec)
		}
		state.Pointer = sec - 1//-1 becuse after this the poiter adds 1
	}
}