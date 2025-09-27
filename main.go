package main

import "fmt"

func main() {
    var grille [6][7]string

    // Remplir la grille avec des espaces  possible de remplacé par X ou O
    for i := 0; i < 6; i++ {
        for j := 0; j < 7; j++ {
            grille[i][j] = " "
        }
    }

    for i := 0; i < 6; i++ {
        for j := 0; j < 7; j++ {
            fmt.Print("[", grille[i][j], "]")
        }
        fmt.Println()
    }
}
