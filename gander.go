package gander

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

// NGMap is a map[string]int
var NGMap map[string]int

func init() {
	aRgx, _ = regexp.Compile("^([=\\?]?[1F]?[1M]?[=F]?[\\?M]?)([\\p{L}]+)?\\s+?([\\p{L}]+?[\\p{L}\\+]+)\\s?([\\p{L}]+)?")
	NGMap = make(map[string]int)
	loadNGData()
}

func getGenderFromString(g string) int {
	if g == "f" {
		return 2
	}
	return 1
}

func getGenderFromInt(g int) string {
	if g == 2 {
		return "f"
	} else if g == 1 {
		return "m"
	}
	return "unknown"
}

func loadNGData() {
	gopath := os.Getenv("GOPATH")
	// log.Println(gopath)
	file, err := os.Open(gopath + "/src/github.com/octohedron/gander/nam_dict.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	pRep := []string{"", " ", "-"}
	for scanner.Scan() {
		s := scanner.Text()
		// don't parse the beginning of the file
		if !strings.ContainsAny(s, "#") {
			// parse current line with the regex
			line := aRgx.FindStringSubmatch(s)
			// log.Println(line)
			var t NameGender
			// if the regex found at least 2 groups
			if len(line) > 2 {
				// if the gender is not unknown
				if line[1] != "?" && line[1] != "=" {
					// assign gender
					t.Gender = strings.ToLower(line[1])
					// Make mostly female female
					if strings.ContainsAny(t.Gender, "f") {
						t.Gender = "f"
					}
					// make mostly male male
					if strings.ContainsAny(t.Gender, "m") {
						t.Gender = "m"
					}
					t.Name = strings.ToLower(line[3])
					// If name is 2 words, add the second name to the line
					if len(line) > 3 {
						if line[4] != "" {
							t.Name = t.Name + " " + strings.ToLower(line[4])
						}
					}
					// if the name has a + on it, replace it with ["", " " and "-"]
					if strings.ContainsAny(t.Name, "+") {
						for _, v := range pRep {
							// load it in the global slice
							NGMap[strings.Replace(t.Name, "+", v, -1)] = getGenderFromString(t.Gender)
							NGData = append(NGData, NameGender{
								Name:   strings.Replace(t.Name, "+", v, -1),
								Gender: t.Gender,
							})
						}
					} else {
						NGMap[t.Name] = getGenderFromString(t.Gender)
						NGData = append(NGData, t)
					}
				}
			}
		}
	}
	// log.Println("DEBUG: loaded " + strconv.Itoa(len(NGData)) + " names with gender")
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// CheckGender returns the gender of the name
func CheckGender(n string) (v NameGender, err error) {
	name := strings.ToLower(n)
	for _, v := range NGData {
		if v.Name == name {
			return v, nil
		}
	}
	return NameGender{Gender: "unknown", Name: name}, errors.New("Name not found")
}

// CheckGenderMap returns the gender of the name
func CheckGenderMap(n string) string {
	return getGenderFromInt(NGMap[strings.ToLower(n)])
}
