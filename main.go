package main

import (
	"fmt"

	"github.com/pedro-git-projects/go-of-cthullu/pkg/investigators"
)

func main() {
	var pedro investigators.Investigator
	pedro.DetermineCharacteristics()
	pedro.SetName("Pedro")
	pedro.SetAge(80)
	pedro.SetBirthplace("Boston")
	pedro.SetResidence("New York")
	pedro.SetOccupation("Investigator")
	investigators.PrintInvestigator(pedro)
	pedro.AccountForAgeModifiers()
	fmt.Println()
	fmt.Println()
	investigators.PrintInvestigator(pedro)

}
