package main

import (
    "fmt"
)

func main() {
    var kleur string

    fmt.Print("voer een kleur in: ")
    fmt.Scanln(&kleur)

    switch kleur {
    case "rood":
        fmt.Println("rood met passie")
    case "blauw":
        fmt.Println("zo blauw als de lucht")
    case "geel":
        fmt.Println("geel als de stralen van de zon")
    case "groen":
        fmt.Println("groen als de natuur")
	case "paars":
		fmt.Println("Paarse vrijdag")
	case "oranje": 
		fmt.Println("kleur van Nederland.")
    default:
        fmt.Println("deze kleur is niet geregistreerd.")
    }
}
