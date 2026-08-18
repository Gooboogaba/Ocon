package commands

import (
    "fmt"
    "io"
    "log"
    "os"
)


func FancyImports(string path) {
    r, w := io.Pipe()

    go func() {
        fmt.Fprint(w, "Hello there\n")
        w.Close()
    }()

    a, err := io.Copy(os.Stdout, r)

    if err != nil {
        log.Fatal(err)
    }
	fmt.Println(a)
}
/*
func main() {
    r, w := io.Pipe()

    go func() {
        fmt.Fprint(w, "Hello there\n")
        w.Close()
    }()

    _, err := io.Copy(os.Stdout, r)

    if err != nil {
        log.Fatal(err)
    }
}
*/