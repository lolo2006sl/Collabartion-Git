package main

import "fmt"

func main() {
    player1 := "X"
    player2 := "O"

    currentPlayer := player1
    tour := 0
    var grille [6][7]string

// Remplir la grille avec des espaces
    for i := 0; i < 6; i++ {
        for j := 0; j < 7; j++ {
            grille[i][j] = " "
        }
    }

    fmt.Println("Début du jeu Puissance 4")
    fmt.Println("Joueur 1 :", player1)
    fmt.Println("Joueur 2 :", player2)
    fmt.Println("Joueur courant :", currentPlayer)
    fmt.Println("Tour numéro :", tour)
    fmt.Println()

    for i := 0; i < 6; i++ {
        for j := 0; j < 7; j++ {
            fmt.Print("[", grille[i][j], "]")
        }
        fmt.Println()
    }
}
