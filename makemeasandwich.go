package main

import (
    "gopkg.in/headzoo/surf.v1"
    "fmt"
    "bufio"
    "os"
    "strings"
    "syscall"

    "golang.org/x/crypto/ssh/terminal"
)

func main() {
    // Create a new browser and open reddit.
    bow := surf.NewBrowser()
    err := bow.Open("https://online.jimmyjohns.com/#/login")
    if err != nil {
        panic(err)
    }
    email, password := getUserLoginInfo();

    fmt.Println(email);
    fmt.Println(password);

}

func getUserLoginInfo() (string, string) {
    //TODO Read config file with login credentials and call function if it doesnt work
    return credentials()
}

func credentials() (string, string){
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter Email: ")
    email, _ := reader.ReadString('\n')

    fmt.Print("Enter Password: ")
    bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
    if err == nil {
        fmt.Println("\nPassword typed: " + string(bytePassword))
    }
    password := string(bytePassword)

    return strings.TrimSpace(email), strings.TrimSpace(password)
}
