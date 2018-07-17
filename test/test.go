package main

import "log"
import "github.com/octohedron/gander/gander"

func main() {
	r, err := gander.CheckGender("Er Dong")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(r.Gender)
	}
}
