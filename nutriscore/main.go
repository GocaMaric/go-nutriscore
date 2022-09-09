package main

import (
	"fmt"
)

// Calculating nutritional score
// Enter values in brackets and find out nutritional score of your food
// More information about nutritional score on https://en.wikipedia.org/wiki/Nutri-Score
func main() {
	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(125),
		Sugars:              SugarGram(5),
		SaturatedFattyAcids: SaturatedFattyAcidsGram(15),
		Sodium:              SodiumMilligram(15),
		Fruits:              FruitsPercent(10),
		Fibre:               FibreGram(20),
		Protein:             ProteinGram(3),
	}, Food)

	// Prints output of main and nutrition score of the food
	fmt.Printf("Nutritional score: %d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())

}
