package main

import (
    "fmt"
    "github.com/tokisuno/goltime/cmd/goltime/root"
    )

func main() {
    rootCmd := root.RootCommand()

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}
