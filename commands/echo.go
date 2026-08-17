package commands

import "fmt"
import "strings"

func EchoCmd(args []string) {
    if len(args) == 0 {
        fmt.Println("no args")
        return
    }

    text := args[0]
	text = strings.ReplaceAll(text, "_", " ")
    if len(text) > 0 && (text[0] == '"' || text[0] == '\'') {
        fmt.Println(text[1:])
    } else if (text[0] == '~') {
		if text[1:] == "true" {
			fmt.Println("✅")
		} else {
			fmt.Println("❌")
		}
	} else {
        fmt.Println("please put a string into echo :). [" + args[0] + "]")
    }
}