# Gender guesser written in Go

Inspired by the [python gender guesser](https://github.com/lead-ratings/gender-guesser), with data from the program "gender" by Jorg Michael

## Motivation

+ Performance
+ Only males or females
+ In a go package

## Usage

+ Make sure you have your `$GOPATH` set, otherwise 

```bash
$ export GOPATH=~/go # Or your gopath location
```

+ Install

```bash
$ go get -u github.com/octohedron/gander
```

+ Example

```go
package main

import "log"
import "github.com/octohedron/gander"

func main() {
	r, err := gander.CheckGender("Aad")
	if err == nil {
		// prints 'm'
		log.Println(r.Gender)
	} else {
		// It would print 'Aad'
		log.Println(r.Name)
    }
}
```

## Performance

```go
func TestAllLoadedNames(t *testing.T) {
	var males int
	var females int
	var total int
	for _, n := range NGData {
		gender, err := CheckGender(n.Name)
		total++
		if err == nil {
			if gender.Gender == "f" {
				females++
			} else if gender.Gender == "m" {
				males++
			}
		}
	}
	t.Logf("In %d we found %d females and %d males", total, females, males)
}
```

Prints 

```bash
=== RUN   TestAllLoadedNames
--- PASS: TestAllLoadedNames (2.32s)
    ...gander/gander_test.go:50: In 41437 we found 20505 females and 20932 males
PASS
ok  	github.com/octohedron/gander	2.393s
Success: Tests passed.
```

2.32s for 41.437 names, or 17.860 names per second in a laptop

### LICENSE
+ The data file nam_dict.txt is released under the GNU Free Documentation License.
+ The project is released under the MIT license
