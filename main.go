package main

import (
	"github.com/pedro-git-projects/go-of-cthullu/pkg/investigators"
)

func main() {
	var pedro investigators.Investigator
	pedro.DetermineCharacteristics()
	pedro.SetName("Pedro")
	pedro.SetAge(24)
	pedro.SetBirthplace("Boston")
	pedro.SetResidence("New York")
	pedro.SetOccupation("Investigator")

	investigators.PrintInvestigator(pedro)
}
