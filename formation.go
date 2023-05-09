package main 

import (
		"fmt"
		"time"
		
)

//Fonctions, 2eme paramettre est la sortie

func addition(a,b int)(resu int){
	resu =  a + b
    return resu
	// return resultat, nil
}

//Structure

type Rectangle struct {
	L float64
	H float64
}

//interface
type Forme interface {
	Aire () float64
	Perimetre() float64
}
func (r Rectangle) Aire() float64 {
	return r.L * r.H
}

func (r Rectangle) Perimetre() float64 {
	return 2 * (r.L + r.H)
}

func afficherMesures(f Forme) {
	fmt.Printf("Aire : %.2f\n", f.Aire())
	fmt.Printf("Perimetre : %.2f\n", f.Perimetre())
}

// func testchan(c)


func main() {
	//Variables
	// nom := "Barry"
	a := 5
	b := 3
      addition(a, b)
	
	// fmt.Printf("a: %d\n", a);

	//Tableau / Slice

	slice := make([]string, 2, 3)
	slice[0] = "test"
	slice[1] = "Voiture"
	slice = append(slice, "hello")
	slice = append(slice, "world") //Quand le slice depasse la capacité prevue au debut alors la capaicte double, ici au lieu de 3 ça pase a 6
	fmt.Printf("Taille = %v\n Capacite %v\n", len(slice), cap(slice))

	//struct
	r := Rectangle{L: 4, H: 6}
	fmt.Printf("-----------------------------------------------------------------------\n")
	fmt.Printf("Largeur = %v | Hauteur = %v\n", r.L, r.H)
	fmt.Printf("------------------------------ MESURE ---------------------------------\n")
	afficherMesures(r)

	//Map

	//Boucle

	c := make(chan string)

	// go testchan(c, "test")
	fmt.Printf("Value chan = %v\n", <-c)

   time.Sleep(2 * time.Second)
}