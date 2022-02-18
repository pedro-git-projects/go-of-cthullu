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
	Occupation   string
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
	Mv           int
	CreditRating int
	Description  Description
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
	i.setMoveRate()
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

func (i *Investigator) setMoveRate() {
	if i.Dex < i.Siz && i.Str < i.Siz {
		i.Mv = 7
	} else if i.Str >= i.Siz || i.Dex >= i.Siz {
		i.Mv = 8
	} else if i.Str > i.Siz && i.Dex > i.Siz {
		i.Mv = 9
	}
}

func rollStr() int {
	d6 := dice.Dice{NumberOfSides: dice.D6}
	str := (3 * d6.Roll()) * 5
	return str
}

func rollCon() int {
	d6 := dice.Dice{NumberOfSides: dice.D6}
	con := (3 * d6.Roll()) * 5
	return con
}

func rollDex() int {
	d6 := dice.Dice{NumberOfSides: dice.D6}
	dex := (3 * d6.Roll()) * 5
	return dex
}

func rollApp() int {
	d6 := dice.Dice{NumberOfSides: dice.D6}
	app := (3 * d6.Roll()) * 5
	return app
}

func rollPow() int {
	d6 := dice.Dice{NumberOfSides: dice.D6}
	pow := (3 * d6.Roll()) * 5
	return pow
}

func rollSiz() int {
	d6 := dice.Dice{NumberOfSides: dice.D6}
	siz := ((2*d6.Roll() + 6) * 5)
	return siz
}

func rollInt() int {
	d6 := dice.Dice{NumberOfSides: dice.D6}
	intelligence := ((2*d6.Roll() + 6) * 5)
	return intelligence
}

func rollEdu() int {
	d6 := dice.Dice{NumberOfSides: dice.D6}
	edu := ((2*d6.Roll() + 6) * 5)
	return edu
}

func rollLuck() int {
	d6 := dice.Dice{NumberOfSides: dice.D6}
	luck := (3 * d6.Roll()) * 5
	return luck
}

func (i *Investigator) eduImprovementCheck() {
	d100 := dice.Dice{NumberOfSides: dice.D100}
	roll := d100.Roll()
	tenthOfRoll := (roll / 10)
	if roll > i.Edu && (i.Edu+tenthOfRoll <= 99) {
		i.Edu += tenthOfRoll
	}
}

func isGreaterThanZero(status int, debuff int) bool {
	if status-debuff >= 0 {
		return true
	}
	return false
}

// finds the highest non-negative debuff value within range
func findOptimumDebuff(maxDebuff int, status int) int {
	dynamicDebuff := 0
	for (dynamicDebuff <= maxDebuff) && (status-dynamicDebuff >= 0) {
		dynamicDebuff++
	}
	// by one correction because of loop break
	return dynamicDebuff - 1
}

func (i *Investigator) youngModifier() {
	strDebuff := findOptimumDebuff(2, i.Str)
	i.Str = i.Str - strDebuff

	sizDebuff := findOptimumDebuff(2, i.Siz)
	i.Siz = i.Siz - sizDebuff
}

func (i *Investigator) fortiesModifier() {
	strDebuff := findOptimumDebuff(1, i.Str)
	i.Str = i.Str - strDebuff

	conDebuff := findOptimumDebuff(2, i.Con)
	i.Con = i.Con - conDebuff

	dexDebuff := findOptimumDebuff(2, i.Dex)
	i.Dex = i.Dex - dexDebuff

	mvDebuff := findOptimumDebuff(1, i.Mv)
	i.Mv = i.Mv - mvDebuff

	i.eduImprovementCheck()
	i.eduImprovementCheck()
}

func (i *Investigator) fifthtiesModifier() {
	strDebuff := findOptimumDebuff(3, i.Str)
	i.Str = i.Str - strDebuff

	conDebuff := findOptimumDebuff(3, i.Con)
	i.Con = i.Con - conDebuff

	dexDebuff := findOptimumDebuff(3, i.Dex)
	i.Dex = i.Dex - dexDebuff

	mvDebuff := findOptimumDebuff(2, i.Mv)
	i.Mv = i.Mv - mvDebuff

	appDebuff := findOptimumDebuff(10, i.App)
	i.App = i.App - appDebuff

	i.eduImprovementCheck()
	i.eduImprovementCheck()
	i.eduImprovementCheck()
}

func (i *Investigator) sixtiesModifier() {

	strDebuff := findOptimumDebuff(7, i.Str)
	i.Str = i.Str - strDebuff

	conDebuff := findOptimumDebuff(1, i.Con)
	i.Con = i.Con - conDebuff

	dexDebuff := findOptimumDebuff(7, i.Dex)
	i.Dex = i.Dex - dexDebuff

	mvDebuff := findOptimumDebuff(3, i.Mv)
	i.Mv = i.Mv - mvDebuff

	appDebuff := findOptimumDebuff(15, i.App)
	i.App = i.App - appDebuff

	i.eduImprovementCheck()
	i.eduImprovementCheck()
	i.eduImprovementCheck()
	i.eduImprovementCheck()
}

func (i *Investigator) seventiesModifier() {

	strDebuff := findOptimumDebuff(13, i.Str)
	i.Str = i.Str - strDebuff

	conDebuff := findOptimumDebuff(13, i.Con)
	i.Con = i.Con - conDebuff

	dexDebuff := findOptimumDebuff(13, i.Dex)
	i.Dex = i.Dex - dexDebuff

	mvDebuff := findOptimumDebuff(4, i.Mv)
	i.Mv = i.Mv - mvDebuff

	appDebuff := findOptimumDebuff(20, i.App)
	i.App = i.App - appDebuff

	i.eduImprovementCheck()
	i.eduImprovementCheck()
	i.eduImprovementCheck()
	i.eduImprovementCheck()
}

func (i *Investigator) eightiesModifier() {

	strDebuff := findOptimumDebuff(13, i.Str)
	i.Str = i.Str - strDebuff

	conDebuff := findOptimumDebuff(13, i.Con)
	i.Con = i.Con - conDebuff

	dexDebuff := findOptimumDebuff(13, i.Dex)
	i.Dex = i.Dex - dexDebuff

	mvDebuff := findOptimumDebuff(5, i.Mv)
	i.Mv = i.Mv - mvDebuff

	appDebuff := findOptimumDebuff(25, i.App)
	i.App = i.App - appDebuff

	i.eduImprovementCheck()
	i.eduImprovementCheck()
	i.eduImprovementCheck()
	i.eduImprovementCheck()
}

func (i *Investigator) AccountForAgeModifiers() {
	switch i.Age >= 15 {
	case i.Age >= 15 && i.Age <= 19:
		i.youngModifier()
	case i.Age >= 20 && i.Age <= 39:
		i.eduImprovementCheck()
	case i.Age >= 40 && i.Age <= 49:
		i.fortiesModifier()
	case i.Age >= 50 && i.Age <= 59:
		i.fifthtiesModifier()
	case i.Age >= 60 && i.Age <= 69:
		i.sixtiesModifier()
	case i.Age >= 70 && i.Age <= 79:
		i.seventiesModifier()
	case i.Age >= 80:
		i.eightiesModifier()
	}
}

func PrintInvestigator(i Investigator) {
	fmt.Printf("Your name is %s\n", i.Name)
	fmt.Printf("Your age is %d\n", i.Age)
	fmt.Printf("Your birthplace is %s\n", i.Birthplace)
	fmt.Printf("Your residence  is %s\n", i.Residence)
	fmt.Printf("Your occupation  is %s\n", i.Occupation)
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
	fmt.Printf("Your move rate is %d\n", i.Mv)
	fmt.Printf("Your credit rating is %d\n", i.CreditRating)
	fmt.Printf("With regards to strenght you are %s\n", i.Description.strDescription)
	fmt.Printf("You%s\n", i.Description.conDescription)
	fmt.Printf("With regards to appearence you are %s\n", i.Description.appDescription)
	fmt.Printf("You %s\n", i.Description.intDescription)
	fmt.Printf("You are%s\n", i.Description.sizDescription)
	fmt.Printf("You%s\n", i.Description.powDescription)
	fmt.Printf("You%s\n", i.Description.dexDescription)
	fmt.Printf("You%s\n", i.Description.eduDescription)
}
