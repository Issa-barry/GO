package main

import (
	"fmt"
	// "log"
	"os"
)

// func creerFichier() {
// 	fmt.Println("Fonctionnalité : Créer un fichier")
// 	// Ajoutez ici le code pour créer un fichier
// }

// func copierFichier() {
// 	fmt.Println("Fonctionnalité : Copier un fichier")
// 	// Ajoutez ici le code pour copier un fichier
// }

// func lireFichier() {
// 	fmt.Println("Fonctionnalité : Lire un fichier")
// 	// Ajoutez ici le code pour lire un fichier
// }

// func supprimerFichier() {
// 	fmt.Println("Fonctionnalité : Supprimer un fichier")
// 	// Ajoutez ici le code pour supprimer un fichier
// }

// func afficherPermissions() {
// 	fmt.Println("Fonctionnalité : Afficher les permissions d'un fichier")
// 	// Ajoutez ici le code pour afficher les permissions d'un fichier
// }

// func seConnecterFTP() {
// 	fmt.Println("Fonctionnalité : Se connecter à un serveur FTP")
// 	// Ajoutez ici le code pour se connecter à un serveur FTP
// }

// func seConnecterHTTP() {
// 	fmt.Println("Fonctionnalité : Se connecter à un serveur HTTP")
// 	// Ajoutez ici le code pour se connecter à un serveur HTTP
// }

// func seConnecterSSH() {
// 	fmt.Println("Fonctionnalité : Se connecter à un serveur SSH")
// 	// Ajoutez ici le code pour se connecter à un serveur SSH
// }

func quitterApplication() {
	fmt.Println("L'application a été quittée.")
	os.Exit(0)
}

func afficherMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. Créer un fichier")
	fmt.Println("2. Copier un fichier")
	fmt.Println("3. Lire un fichier")
	fmt.Println("4. Supprimer un fichier")
	fmt.Println("5. Afficher les permissions d'un fichier")
	fmt.Println("6. Se connecter à un serveur FTP")
	fmt.Println("7. Se connecter à un serveur HTTP")
	fmt.Println("8. Se connecter à un serveur SSH")
	fmt.Println("9. Quitter l'application")
}

func lireChoix() int {
	var choix int
	fmt.Print("\n Sélectionnez une option : \n")
	fmt.Scanln(&choix)
	return choix
}

func executerOption(choix int) {
	switch choix {
	case 1:
		fmt.Printf("Creer un fichier\n")
		// creerFichier()
	case 2: fmt.Printf("Copier un fichier\n")
		// copierFichier()
	case 3: fmt.Printf("lire un fichier\n")
		// lireFichier()
	case 4: fmt.Printf("Supprimer un fichier\n")
		// supprimerFichier()
	case 5: fmt.Printf("Afficher permission\n")
		// afficherPermissions()
	case 6: fmt.Printf("Se connecter FTP\n")
		// seConnecterFTP()
	case 7: fmt.Printf("Se connecter HTTP\n")
		// seConnecterHTTP()
	case 8: fmt.Printf("Se connecter SSH\n")
		// seConnecterSSH()
	case 9:
		// quitterApplication()
	default:
		fmt.Println("Option invalide. Veuillez sélectionner une option valide.")
	}
}

func main() {
	fmt.Printf("-------------------------------------------Afficher-------------------------------------\n")
	for {
		afficherMenu()
		choix := lireChoix()
		executerOption(choix)
	}
}
