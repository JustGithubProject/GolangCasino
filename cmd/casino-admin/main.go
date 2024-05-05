package main

import (
    "fmt"
    "github.com/JustGithubProject/GolangCasino/cmd/internal/services"
)

func main() {

    err_2 := services.CheckPasswordHash("12345678", "fef3d83e32b4d981b0c0f75206e891268c6aa8bd8db5a315db7bf24168a4be27")
    fmt.Println(err_2)
}
