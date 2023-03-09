package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

type Colors struct {
    Colors []struct {
        Name string `json:"name"`
    } `json:"colors"`
}

func main() {
    // Lees het JSON-bestand
    file, err := ioutil.ReadFile("colors.json")
    if err != nil {
        fmt.Println("Error reading JSON file:", err)
        return
    }

    // Decodeer de JSON-data, haalt de data eruit.
    var colors Colors
    err = json.Unmarshal(file, &colors)
    if err != nil {
        fmt.Println("Error decoding JSON file:", err)
        return
    }

    // Toont de kleurnamen in de terminal
    fmt.Println("kies een van de onderstaande kleuren(nummer):")
    for i, color := range colors.Colors {
        fmt.Printf("%d. %s\n", i+1, color.Name)
    }

    // Kies een kleur met een switch-statement
    var choice int
    fmt.Scanln(&choice)

    var chosenColor string
    switch choice {
    case 1:
        chosenColor = colors.Colors[0].Name
    case 2:
        chosenColor = colors.Colors[1].Name
    case 3:
        chosenColor = colors.Colors[2].Name
    case 4:
        chosenColor = colors.Colors[3].Name
    case 5:
        chosenColor = colors.Colors[4].Name   
    default:
        fmt.Println("foute keuze")
        return
    }

    // Toon de gekozen kleur in de terminal
    fmt.Printf("je kiest kleur %s\n", chosenColor)
}