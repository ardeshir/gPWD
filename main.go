package main

import (
        "fmt"
        "os"
        "path/filepath"
)

func main() {

        args := os.Args
        pwd, err := os.Getwd()

        if err == nil {
                fmt.Println(pwd)
        } else  {
                fmt.Println("Error: ", err)
        }

        if len(args)  == 1 {
                return
        }

        if args[1] != "-P" {
                return
        }

        fileInfo, err := os.Lstat(pwd)
        if fileInfo.Mode()&os.ModeSymlink != 0 {
                realpath, err := filepath.EvalSymlinks(pwd)
                if err == nil  {
                        fmt.Println(realpath)
                }
        }

}
