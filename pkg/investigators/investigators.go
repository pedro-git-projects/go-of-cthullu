package investigators

import (
	"fmt"

	"github.com/pedro-git-projects/go-of-cthullu/pkg/dice"
)

type Investigator struct {
	Name         string
	Age          int
	Residence    string
	Birthplace   string
	Str          int
	Con          int
	Pow          int
	Dex          int
	App          int
	Siz          int
	Int          int
	Edu          int
	Luck         int
	Mp           int
	Db           int
	Hp           int
	San          int
	Occupation   string
	CreditRating int
}

func (i *Investigator) DetermineCharacteristics() {
	i.Str = rollStr()
	i.Con = rollCon()
	i.Dex = rollDex()
	i.App = rollApp()
	i.Pow = rollPow()
	i.setSanity()
	i.setMp()
	i.Siz = rollSiz()
	i.setHp()
	i.Int = rollInt()
	i.Edu = rollEdu()
	i.Luck = rollLuck()
}

func (i *Investigator) SetName(name string) {
	i.Name = name
}

func (i *Investigator) SetOccupation(occupation string) {
	i.Occupation = occupation
}

func (i *Investigator) SetBirthplace(birthplace string) {
	i.Birthplace = birthplace
}

func (i *Investigator) SetResidence(residence string) {
	i.Residence = residence
}

func (i *Investigator) SetAge(age int) {
	i.Age = age
}

func (i *Investigator) setSanity() {
	i.San = i.Pow
}

func (i *Investigator) setMp() {
	i.Mp = (i.Pow / 5)
}

func (i *Investigator) setHp() {
	i.Hp = (i.Siz + i.Con) / 10
}

func rollStr() int {
	d6 := dice.Dice{dice.D6}
	str := (3 * d6.Roll()) * 5
	return str
}

func rollCon() int {
	d6 := dice.Dice{dice.D6}
	con := (3 * d6.Roll()) * 5
	return con
}

func rollDex() int {
	d6 := dice.Dice{dice.D6}
	dex := (3 * d6.Roll()) * 5
	return dex
}

func rollApp() int {
	d6 := dice.Dice{dice.D6}
	app := (3 * d6.Roll()) * 5
	return app
}

func rollPow() int {
	d6 := dice.Dice{dice.D6}
	pow := (3 * d6.Roll()) * 5
	return pow
}

func rollSiz() int {
	d6 := dice.Dice{dice.D6}
	siz := ((2*d6.Roll() + 6) * 5)
	return siz
}

func rollInt() int {
	d6 := dice.Dice{dice.D6}
	intelligence := ((2*d6.Roll() + 6) * 5)
	return intelligence
}

func rollEdu() int {
	d6 := dice.Dice{dice.D6}
	edu := ((2*d6.Roll() + 6) * 5)
	return edu
}

func rollLuck() int {
	d6 := dice.Dice{dice.D6}
	luck := (3 * d6.Roll()) * 5
	return luck
}

func PrintInvestigator(i Investigator) {
	fmt.Printf("Your name is %s\n", i.Name)
	fmt.Printf("Your age is %d\n", i.Age)
	fmt.Printf("Your birthplace is %s\n", i.Birthplace)
	fmt.Printf("Your residence  is %s\n", i.Residence)
	fmt.Printf("Your strengh is %d\n", i.Str)
	fmt.Printf("Your constitution is %d\n", i.Con)
	fmt.Printf("Your power is %d\n", i.Pow)
	fmt.Printf("Your dexterity is %d\n", i.Dex)
	fmt.Printf("Your appearence is %d\n", i.App)
	fmt.Printf("Your size is %d\n", i.Siz)
	fmt.Printf("Your intelligence is %d\n", i.Int)
	fmt.Printf("Your education %d\n", i.Edu)
	fmt.Printf("Your luck is %d\n", i.Luck)
	fmt.Printf("Your magic points are %d\n", i.Mp)
	fmt.Printf("Your damage bonus is %d\n", i.Db)
	fmt.Printf("Your hit points are %d\n", i.Hp)
	fmt.Printf("Your sanity is %d\n", i.San)
	fmt.Printf("Your occupation  is %s\n", i.Occupation)
	fmt.Printf("Your credit rating is %d\n", i.CreditRating)
}
