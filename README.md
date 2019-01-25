# Gender guesser written in Go

### Used data from the [python gender guesser](https://github.com/lead-ratings/gender-guesser)

## Motivation

Currently, the [python gender guesser](https://github.com/lead-ratings/gender-guesser) can take up to 1 second on my laptop per name, which is not a big deal but if you have to process 100M names, it could take 3 years, this should do it in less than 1 hour.

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
+ The project is released under MIT license
