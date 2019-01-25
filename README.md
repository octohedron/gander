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
	g, err := gander.CheckGender("Aad")
	if err == nil {
		// prints 'm'
		log.Println(g.Gender)
	} else {
		// It would print 'Aad' if there was an error
		log.Println(g.Name)
    }
}
```

## Performance

```go
func TestAllLoadedNames(t *testing.T) {
	var males int
	var females int
	for _, n := range NGData {
		g, err := CheckGender(n.Name)
		if err == nil {
			if g.Gender == "f" {
				females++
			} else {
				males++
			}
		}
	}
	t.Logf("In %d we found %d females and %d males", len(NGData), females, males)
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
