package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

type Student struct { //struct for json
	Adress     string `json:"a"`
	Bgroup     string `json:"b"`
	Department string `json:"d"`
	Gender     string `json:"g"`
	Hall       string `json:"h"`
	ID         string `json:"i"`
	Name       string `json:"n"`
	r          string `json:"r"`
	Programme  string `json:"p"`
	Useranme   string `json:"u"`
}

func check(s string) bool { //checking valid name
	for _, c := range s {
		if (!unicode.IsLetter(c)) && (c != ' ') {
			fmt.Println(s)
			return false
		}
	}
	if (!strings.ContainsAny(s, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")) && (!strings.ContainsAny(s, " ")) {
		fmt.Println(s)
		return false
	}
	return true

}

func main() {
	jsonFilename := flag.String("json", "student.json", " json filename [string]")
	csvFilename := flag.String("csv", "project.csv", " a csv file with question , answers format")

	flag.Parse()
	jsonFile, err := os.Open(*jsonFilename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var students []Student

	json.Unmarshal(byteValue, &students)

	//fmt.Println(students)

	csvFile, err := os.Open(*csvFilename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := csv.NewReader(csvFile)
	for {
		line, error := r.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if check(line[0]) {
			for _, student := range students {
				if student.Name == line[0] {
					fmt.Printf("Name : %s\n", line[0])
					fmt.Printf("Roll No : %s\n", student.ID)
					fmt.Printf("Branch : %s\n", student.Department)
					fmt.Printf("Oragnisation : %s\n", line[1])
					fmt.Printf("Project : %s\n\n\n", line[2])

				}
			}
		}
	}
}
