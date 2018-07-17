package test

import "log"

func test() {
	r, err := checkGender("Er Dong")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(r.Gender)
	}
}
