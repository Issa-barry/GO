//Mini projet en go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func creerFichier() {
	var nomFichier string
	fmt.Print("Entrez le nom du fichier à créer : ")
	fmt.Scanln(&nomFichier)

	file, err := os.Create(nomFichier)
	if err != nil {
		log.Fatalf("Erreur : Impossible de créer le fichier. %v", err)
	}
	defer file.Close()

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

func seConnecterFTP() {
	// Code pour se connecter à un serveur FTP
	fmt.Println("Connexion au serveur FTP...")
}

func seConnecterHTTP() {
	// Code pour se connecter à un serveur HTTP
	fmt.Println("Connexion au serveur HTTP...")
}

func seConnecterSSH() {
	// Code pour se connecter à un serveur SSH
	fmt.Println("Connexion au serveur SSH...")
}

func quitterApplication() {
	fmt.Println("L'application a été quittée.")
	os.Exit(0)
}

func main() {
	for i := 0; i<10; i++{
		fmt.Println("\nMenu:")
		fmt.Println("1. Créer un fichier")
		fmt.Println("2. Copier un fichier")
		fmt.Println("3. Lire un fichier")
		fmt.Println("4. Supprimer un fichier")
		fmt.Println("5. Afficher les permissions d'un fichier")
		fmt.Println("6. Se connecter à")
	}
}
