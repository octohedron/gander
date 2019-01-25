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

### LICENSE
+ The data file nam_dict.txt is released under the GNU Free Documentation License.
+ The project is released under the MIT license
