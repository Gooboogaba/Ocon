package commands

import (
	"fmt"
	"oscarsgoofysite/OCON/state"
	"crypto/rand"
	"encoding/hex"
	"oscarsgoofysite/OCON/returncommands"
	"io"
	"bufio"
	"os/exec"
	"strings"
)

type API struct {
	Stdin  io.WriteCloser
	Stdout io.ReadCloser
}

var ImportedAPIs = map[string]*API{}

func FancyImports(path string) {
	//create the password
	//var ApiPass := "empty"
	if state.ApiPass == "empty" {
		bytes := make([]byte, 32)

		_, err := rand.Read(bytes)
		if err != nil {
			panic(err)
		}

		state.ApiPass = hex.EncodeToString(bytes)
	}
	fmt.Println("API password:", state.ApiPass)
	
	
	//ok realy pipes here
	cmd := exec.Command(path)
	//stdin has NOTHING to do with the other kind of stds
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	if err := cmd.Start(); err != nil {
		panic(err)
	}
	
	// Send to THIS specific process
	fmt.Fprintln(stdin, "AUTH "+state.ApiPass)
	
	// Read from THIS specific process
	scanner := bufio.NewScanner(stdout)
	
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println("API:", message)
		commandHandler(message, stdin, stdout)
	}
}

func commandHandler(cmd string, stdin io.Writer, stdout io.Reader) {
	command := strings.Fields(cmd) // turn into an array
	
	//check if they have the key
	if command[0] != state.ApiPass {
		fmt.Fprintln(stdin, "ERROR wrong or no api key")
		return		
	} else {
		command = command[1:]
	}
	

	switch command[0] {
		case "REGISTER":
			registerCommand(command[1:], stdin ,stdout)
		case "ANSWER":
			return //ANSWER is handeled by another func
		//more commands maybe later
		//i know there is only 2 cases but keep this a switch-case
		default:
			return
			fmt.Fprintln(stdin, "ERROR unrecognized fancy import api command")
	}
}
//hey same as above for params (⌐■_■)
func registerCommand(cmd []string, stdin io.Writer, stdout io.Reader) {
	if cmd[0] == "r" {
		//return command
		returncommands.AddCommand(
			cmd[0],
			RunReturnRegisteredCommand(cmd[0], stdin, stdout),
		)
	} else {
		//reglaur command
		
	}
}

func RunReturnRegisteredCommand(name string, stdin io.Writer, stdout io.Reader) returncommands.Function {
	return func(args []string) []string {//sadly had to use ai for most of this becuse the intenet is unhelpful sorry :(
		fmt.Fprintln(stdin, "RUN "+name)

		scanner := bufio.NewScanner(stdout)

		if !scanner.Scan() {
			return []string{}
		}

		parts := strings.Fields(scanner.Text())

		if len(parts) < 3 {
			return []string{}
		}

		if parts[0] != state.ApiPass {
			fmt.Fprintln(stdin, "ERROR apipass is wrong")
			return []string{}
		}

		if parts[1] != "ANSWER" {
			return []string{}
		}

		return parts[2:]
	}
}