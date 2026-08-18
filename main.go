package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"oscarsgoofysite/OCON/commands"
	"oscarsgoofysite/OCON/state"
	"oscarsgoofysite/OCON/returncommands"

)

func main() {
	fmt.Println("🐟")

	if len(os.Args) < 2 {
		var i string
		fmt.Println("This is OCON a language created by Oscar! (see more at: https://github.com/oscar366/Ocon)")
		fmt.Println(`To execute an .ocon file do: "ocon execute {path}"`)
		fmt.Println("this is a command line tool use it in cmd")
		fmt.Println("")
		fmt.Println("type something and click enter to leave")
		fmt.Scan(&i)
		fmt.Println(i)
		return
	}

	switch os.Args[1] {
	case "execute":
		if len(os.Args) < 3 {
			fmt.Println("missing file path")
			return
		}
		readFile(os.Args[2])
		
	case "🐟":
		fmt.Println("🐟")
		for i := 0; i < 9999; i++ {
			fmt.Println(string(i) + ": 🐟")
		}

	case "help":
		fmt.Println("This is OCON a language created by Oscar!")
		fmt.Println(`To execute an .ocon file do: "ocon execute {path}"`)
	default:
		fmt.Println("Put in an input")
	}
}

func readFile(path string) {
	fmt.Println("executing:" + path)
	
	lines := []string{}
	
	file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	
	scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
		/*if pointer == 1 { 
			fmt.Println(scanner.Text())
		}*/
		lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	
	
	executeall(lines)
}

type Command func([]string)
var commmands = map[string]Command{
	//other
    "echo":      commands.EchoCmd,
	"#":	emptycommand,//comments
	
	//vars
    "var":       commands.VarCmd,
    "increment": commands.IncrementCmd,
	"decremnt": commands.DecremntCmd, 
	//"program": commands.ProgramCmd,
	
	//sections
	"§": emptycommand, //we aculy dont need these commands but to not confuse the inteprted we keep them
	"sec": emptycommand,
	"goto": commands.GotoCmd,
	
	//conditionals
	"if": commands.If,
	"import": emptycommand,
}


func executeall(lines []string) {
	//before runing the program get the positons of all the sections
fmt.Println(" ")
fmt.Println("doing setup")
fmt.Println("===============================")
fmt.Println(" ")

//more complacted for loop :-0
for i, line := range lines {
    parts := strings.Fields(line) // sperate by " "
    if len(parts) >= 2 && (parts[0] == "§" || parts[0] == "sec") { //if there is more then 2 inputs and its a section then run
        state.SectionList[parts[1][1:]] = i //add to the map the name and line number of the section
    }
	
	if len(parts) >= 2 && parts[0] == "import" {
        if parts[1] != "f" {
            importfile := parts[1][1:]

            dat, err := os.ReadFile(importfile)
            if err != nil {
                panic(err)
            }

            newlines := strings.Split(string(dat), "\n")

            // Replace the import line with the imported lines thank you stakoverflow
            lines = append(
                lines[:i],
                append(newlines, lines[i+1:]...)...,
            )

            // Move past the newly inserted lines
            i += len(newlines) - 1
        } else {
			//go commands.FancyImports([2][1:])//2 cuz of the f
			fmt.Println("Not curretly complete oscar is working hard on it though :)")
		}
	}
}

fmt.Println("debug S:")
for key, value := range state.SectionList {
    fmt.Println("Key:", key, "Value:", value)
}
fmt.Println("____________________________")
fmt.Println(" ")
fmt.Println("setup complete")
fmt.Println("====================================")
fmt.Println(" ")

//real execution above is setup
	for state.Pointer < len(lines) {
		//fmt.Printf("%d: %s\n", state.Pointer, lines[state.Pointer]) debug
	
		//rember that 0 is equal to start of a list
		execute(lines[state.Pointer])
		state.Pointer += 1// dude to this after goto is init the pointer adds one so that is why its n-1
	}
	
	
}

func execute(prgmstring string) {
	if strings.TrimSpace(prgmstring) == "" {
		return
	}
	//fmt.Printf("Executing: %q\n", prgmstring) debug
	
	words := strings.Split(prgmstring, " ")
	
//these are vars for in between
inreturn := false
//unset will set later. This is the pos of after the "["
firstpos := 0
	
	//replaces vars with var content and does return functions
	for i, element := range words {
		//this is for the rutrn command
		if inreturn && (element == "]") {
			//get the return of the return func interp
			returncommand := words[firstpos:i]
			value := returncommands.Intrp(returncommand)
			//fmt.Println("debug. Value gotten from interp: ", value[0])
			//rebuild the slaice with value inbtween and remove the ret command
			words = append(
				words[:firstpos - 1],
				append(value, words[i+1:]...)...,
			)
			inreturn = false
		}
		//if there is a "]" part but no starting counterpart
		
		if(element == "[") {
			inreturn = true
			firstpos = i + 1
		}
		
		if(element[0] == '$') {
			//if it is a var then replaces
			varname := element[1:]
			
			val, ok := state.VarStorage[varname]
			//if var does NOT exsit
			if (!ok) {
				fmt.Println("Var does not exist. Or other bug.")
				return
			}
			//replace with var value
			words[i] = val
		}
	}
	
	
	cmd := words[0]
	args := words[1:]
	//get the function(fn) if it exits in the commands list(ok)
	if fn, ok := commmands[cmd]; ok {
			fn(args)
	} else {
			fmt.Println("Unknown command:", cmd)
	}
}

func emptycommand(args []string) {
	return //the sections dont do anything but we requrie this so it does not say "unknown command"
}
