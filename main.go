package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pedro-git-projects/go-of-cthullu/pkg/investigators"
)

func main() {
	var character investigators.Investigator
	r := bufio.NewReader(os.Stdout)
	fmt.Println("Please input the name of your investigator:")
	name, _ := r.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	fmt.Println("Please input the age of your investigator:")
	var age int
	_, _ = fmt.Scanf("%d", &age)

	fmt.Println("Please input the birthplace of your investigator:")
	birth, _ := r.ReadString('\n')
	birth = strings.TrimSuffix(birth, "\n")

	fmt.Println("Please input the residence of your investigator:")
	residence, _ := r.ReadString('\n')
	residence = strings.TrimSuffix(residence, "\n")

	fmt.Println("Please input the occupation of your investigator:")
	occupation, _ := r.ReadString('\n')
	occupation = strings.TrimSuffix(occupation, "\n")

	investigators.CreateInvestigator(&character, name, age, birth, residence, occupation)

	fmt.Println()
	investigators.PrintInvestigator(character)

}
