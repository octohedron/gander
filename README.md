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
$ go get -u github.com/octohedron/gander/gander
```

+ Example

```go
package main

import "log"
import "github.com/octohedron/gander/gander"

func main() {
	r, err := gander.CheckGender("Aad")
	if err == nil {
		// prints 'm'
		log.Println(r.Gender)
	} else {
		// It could print 'unknown'
		log.Println(r.Name)
    }
}
```

## Benchmark

With MongoDB and a collection of around ~2M documents, it updates around 1.000 documents per second

## Contributing

Yes, there's performance improvements and features to add, so feel free to send PRs