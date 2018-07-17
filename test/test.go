package test

import "log"
import "github.com/octohedron/gander/gander"

func test() {
	r, err := gander.checkGender("Er Dong")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(r.Gender)
	}
}
