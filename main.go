package main

import "fmt"

//ajout de pion dans grille
func ajouterPion(grille *[6][7]string, colonne int, symbole string) bool {
	if colonne < 0 || colonne >= 7 {
		fmt.Println("Colonne invalide. Choisissez une colonne entre 0 et 6.")
		return false
	}

	for i := 5; i >= 0; i-- {
		if grille[i][colonne] == " " {
			grille[i][colonne] = symbole
			return true
		}
	}

	fmt.Println(" colonne pleine. Choisissez une autre colonne.")
	return false
}

//----------------------------------------------------------------

// Programme principal
func main() {
	player1 := "X"
	player2 := "O"

	currentPlayer := player1
	tour := 0
	var grille [6][7]string

	// initialisation grille (vide)
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			grille[i][j] = " "
		}
	}

	fmt.Println("Début du jeu Puissance 4")
	fmt.Println("Joueur 1 :", player1)
	fmt.Println("Joueur 2 :", player2)
	fmt.Println()

	// Tour du jeu
	for {
		fmt.Println("Tour numéro :", tour)
		fmt.Println("Joueur courant :", currentPlayer)

		// affichage de la grille
		for i := 0; i < 6; i++ {
			for j := 0; j < 7; j++ {
				fmt.Print("[", grille[i][j], "]")
			}
			fmt.Println()
		}

		// Demande de la colonne
		var colonne int
		fmt.Printf("Joueur %s, choisissez une colonne (0 à 6) : ", currentPlayer)
		fmt.Scan(&colonne)

		// Ajout du pion dans la grille
		if ajouterPion(&grille, colonne, currentPlayer) {
			tour++

			// Changement de joueur
			if currentPlayer == player1 {
				currentPlayer = player2
			} else { // FT3
				currentPlayer = player1
			}
		}
	}
}

//----------------------------------------------------------------
