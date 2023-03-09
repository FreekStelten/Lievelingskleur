package main

import (
    "fmt"
)

func main() {
    var kleur string

    fmt.Print("voer een kleur in: ")
    fmt.Scanln(&kleur)

    if kleur == "rood" {fmt.Println("rood met passie")  }

    if kleur == "blauw" {fmt.Println("zo blauw als de lucht") }

	if kleur == "geel" {fmt.Println("geel als de stralen van de zon") }

	if kleur == "groen" {fmt.Println("groen als de natuur")}

	else {fmt.Println("deze kleur is niet geregistreerd.")}
    
	fmt.Printf("Je hebt de kleur %s ingevoerd.\n", kleur)
}
