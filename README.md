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
		// It would print 'Aad', if there was an error
		log.Println(g.Name)
    }
}
```

## Performance

The fastest is to use the CheckGenderMap func, but it doesn't return an error, only `"unknown"` if it doesn't find the gender.

```golang
func TestBenchmark(t *testing.T) {
	var males int
	var females int
	const total = 100000000 // 100 million
	t.Log("START")
	for i := 0; i < total; i++ {
		g := CheckGenderMap(
			NGData[rand.Intn(len(NGData))].Name)
		if g == "f" {
			females++
		} else {
			males++
		}
	}
	t.Logf("In %d we found %d females and %d males", total, females, males)
}
```
```bash
2019/01/27 20:51:16.996857 START
2019/01/27 20:51:29.781984 In 100000000 we found 47278122 females and 52721878 males
```
That's ~ 7.7M/s 🔥🔥🔥

With error
```go
func TestAllLoadedNames(t *testing.T) {
	var males int
	var females int
	c := make(chan int)
	var complete int
	for _, n := range NGData {
		go func(n NameGender) {
			g, err := CheckGender(n.Name)
			if err == nil {
				if g.Gender == "f" {
					c <- 1
				} else {
					c <- 0
				}
			}
		}(n)
	}
	for p := range c {
		complete++
		if complete == len(NGData) {
			break
		}
		if p == 1 {
			females++
		} else {
			males++
		}
	}
	log.Println("COMPLETED")
	t.Logf("In %d we found %d females and %d males", len(NGData), females, males)
}
```

Prints 

```bash
2019/01/27 13:05:01.639677 LOADED
=== RUN   TestAllLoadedNames
2019/01/27 13:05:02.165385 COMPLETED
--- PASS: TestAllLoadedNames (2.32s)
    ...gander/gander_test.go:50: In 41437 we found 20505 females and 20932 males
PASS
ok  	github.com/octohedron/gander	0.604s
Success: Tests passed.
```

0.525708 for 41.437 names, or ~78.821 names per second in a laptop

### LICENSE
+ The data file nam_dict.txt is released under the GNU Free Documentation License.
+ The project is released under the MIT license
