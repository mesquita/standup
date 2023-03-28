package main

import (
    "fmt"
    "os/exec"
)

func main() {
    url := "https://saltpay-co.zoom.us/my/" 
    err := exec.Command("open", url).Run()
    if err != nil {
        fmt.Println("Error opening URL:", err)
        return
    }
    fmt.Println("Successfully opened URL in your default browser!")
}

