package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
	"strings"
)

// NameGender is a gender and name type
type NameGender struct {
	Name   string
	Gender string
}

// aRgx all regex, match all what we need from the line
var aRgx *regexp.Regexp

// NGData is a slice of NameGender
var NGData []NameGender

func loadNGData() {
	file, err := os.Open("./names_genders.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if !strings.ContainsAny(s, "#") {
			result := aRgx.FindStringSubmatch(s)
			// log.Println(result)
			var t NameGender
			if len(result) > 2 {
				t.Gender = strings.ToLower(result[1])
				t.Name = strings.ToLower(result[3])
			}
			if len(result) > 3 {
				if result[4] != "" {
					t.Name = t.Name + " " + strings.ToLower(result[4])
				}
			}
			NGData = append(NGData, t)
		}
	}
	// log.Println("DEBUG: loaded " + strconv.Itoa(len(NGData)) + " names with gender")
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// checkGender returns the gender of the name
func checkGender(n string) (v NameGender, err error) {
	name := strings.ToLower(n)
	for _, v := range NGData {
		if v.Name == name {
			return v, nil
		}
	}
	return NameGender{}, errors.New("Name not found")
}

func main() {
	aRgx, _ = regexp.Compile("^([=\\?]?[1F]?[1M]?[=F]?[\\?M]?)([\\p{L}]+)?\\s+?([\\p{L}]+)\\s?([\\p{L}]+)?")
	loadNGData()
	r, err := checkGender("Kazik Kazimierz")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(r.Gender)
	}
}
