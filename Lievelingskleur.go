package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)
//Dit definieert de "Colors" structuur met één veld "Colors" dat een slice is van anonieme structuren die elk één veld hebben, genaamd "Name".
type Colors struct {
	Colors []struct {
		Name string `json:"name"`
	} `json:"colors"`
}
//readJSONFile retourneert Colors-struct en error. bij fout lezen van JSON-bestand, error geretourneerd.
func readJSONFile(filePath string) (Colors, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Colors{}, err
	}

	var colors Colors
	err = json.Unmarshal(file, &colors)
	if err != nil {
		return Colors{}, err
	}

	return colors, nil
}
//In de main-functie, worden de waarden geretourneerd door de functie readJSONFile opgeslagen in de variabelen colors en err.
//fout bij het lezen van het JSON-bestand, wordt het bestand "error.txt" weergegeven en wordt het programma afgesloten.
func main() {
	colors, err := readJSONFile("colors.json")
	if err != nil {
		content, err := ioutil.ReadFile("error.txt")
		if err != nil {
			fmt.Println("error reading error.txt", err)
			return
		}
		fmt.Println(string(content))
		return
	}

	fmt.Println("Kies een van de onderstaande kleuren (nummer):")
	for i, color := range colors.Colors {
		fmt.Printf("%d. %s\n", i+1, color.Name)
	}

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

	fmt.Printf("Je hebt gekozen voor kleur %s\n", chosenColor)
}
