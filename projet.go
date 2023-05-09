package main

import (
	"fmt"
	"io/ioutil"
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


func copierFichier() {
	var nomFichierSource, nomFichierDestination string
	fmt.Print("Entrez le nom du fichier source : ")
	fmt.Scanln(&nomFichierSource)
	fmt.Print("Entrez le nom du fichier de destination : ")
	fmt.Scanln(&nomFichierDestination)

	content, err := ioutil.ReadFile(nomFichierSource)
	if err != nil {
		log.Fatalf("Erreur : Impossible de lire le fichier source. %v", err)
	}

	err = ioutil.WriteFile(nomFichierDestination, content, 0644)
	if err != nil {
		log.Fatalf("Erreur : Impossible de copier le fichier. %v", err)
	}

	fmt.Println("Le fichier", nomFichierSource, "a été copié vers", nomFichierDestination)
}

func lireFichier() {
	var nomFichier string
	fmt.Print("Entrez le nom du fichier à lire : ")
	fmt.Scanln(&nomFichier)

	content, err := ioutil.ReadFile(nomFichier)
	if err != nil {
		log.Fatalf("Erreur : Impossible de lire le fichier. %v", err)
	}

	fmt.Printf("Contenu du fichier %s :\n%s", nomFichier, content)
}

func supprimerFichier() {
	var nomFichier string
	fmt.Print("Entrez le nom du fichier à supprimer : ")
	fmt.Scanln(&nomFichier)

	err := os.Remove(nomFichier)
	if err != nil {
		log.Fatalf("Erreur : Impossible de supprimer le fichier. %v", err)
	}

	fmt.Println("Le fichier", nomFichier, "a été supprimé.")
}

func afficherPermissions() {
	var nomFichier string
	fmt.Print("Entrez le nom du fichier : ")
	fmt.Scanln(&nomFichier)

	fileInfo, err := os.Stat(nomFichier)
	if err != nil {
		log.Fatalf("Erreur : Impossible d'obtenir les permissions du fichier. %v", err)
	}

	fmt.Println("Permissions du fichier", nomFichier, ":", fileInfo.Mode().Perm())
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


//Gestion du menu

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
	case 2: 
		copierFichier()
	case 3: 
		lireFichier()
	case 4: 
		supprimerFichier()
	case 5: 
		afficherPermissions()
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
