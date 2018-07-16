package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// "strings"
// "sync"

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
				t.Gender = result[1]
				t.Name = result[3]
			} else {
				log.Println(s)
				log.Println(result)
			}
			if len(result) > 3 {
				if result[4] != "" {
					t.Name = t.Name + result[4]
				}
			}
			NGData = append(NGData, t)
		}
	}
	log.Println("DEBUG: loaded " + strconv.Itoa(len(NGData)) + " names with gender")
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	aRgx, _ = regexp.Compile("^([=\\?]?[1F]?[1M]?[=F]?[\\?M]?)([\\p{L}]+)?\\s+?([\\p{L}]+)\\s?([\\p{L}]+)?")
	loadNGData()
}
