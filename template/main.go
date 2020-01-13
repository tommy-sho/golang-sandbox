package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func main() {
	fread, err := ioutil.ReadFile("./template.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
		Hoge       []string
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true, []string{"hoge", "geho"}},
		//{"Uncle John", "moleskin pants", false},
		//{"Cousin Rodney", "", false},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(string(fread)))

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}
