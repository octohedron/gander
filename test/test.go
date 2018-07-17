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
	r, err = gander.CheckGender("Abdel Karim")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(r.Gender)
	}
	r, err = gander.CheckGender("Adi Adolf")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(r.Gender)
	}
	r, err = gander.CheckGender("Adrian")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(r.Gender)
	}
	r, err = gander.CheckGender("Ágúst")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(r.Gender)
	}
}
