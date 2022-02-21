package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/pedro-git-projects/go-of-cthullu/pkg/investigators"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("public/*"))
}

func getInvestigatorData(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	age := r.FormValue("age")
	residence := r.FormValue("residence")
	birth := r.FormValue("birthplace")
	occupation := r.FormValue("occupation")
	var character investigators.Investigator
	ageNumber, _ := strconv.Atoi(age)
	investigators.CreateInvestigator(&character, name, ageNumber, birth, residence, occupation)

	err := tpl.ExecuteTemplate(w, "index.tmpl", character)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatal(err)
	}
}

func main() {

	http.HandleFunc("/", getInvestigatorData)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

	/*	var character investigators.Investigator
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
	*/
}
