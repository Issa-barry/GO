package main

import (
	"fmt"
	"log"
	"os"
)

/*
On utilise la fonction os.Create pour créer le fichier. En cas d'erreur lors de la création du fichier, 
une erreur sera affichée. Sinon, un message de confirmation sera affiché. La fonction defer est utilisée
pour s'assurer que le fichier est fermé après son utilisation. Dans cet exemple, la fonction creerFichier 
est appelée directement depuis la fonction main.
*/
func creerFichier() {
	var nomFichier string
	fmt.Print("Entrez le nom du fichier à créer : ")
	fmt.Scanln(&nomFichier)

	fichier, err := os.Create(nomFichier)
	if err != nil {
		log.Fatalf("Erreur : Impossible de créer le fichier. %v", err)
	}
	defer fichier.Close()

	fmt.Println("Le fichier", nomFichier, "a été créé.")
}


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
		creerFichier()
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
