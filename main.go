package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("people.json")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	icerik, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	var people []Person
	err = json.Unmarshal(icerik, &people)
	if err != nil {
		fmt.Println(err)
	}

	// They are already in the awesome superpeer repo
	excludeList := []string{
		"yakuter",
		"selcukusta",
		"f",
		"suadev",
		"eser",
		"edisdev",
		"onuraslan",
		"azmimengu",
		"yukselis",
		"dorian",
		"serhat",
		"omerfarukucar",
		"mribrahim"}

	for _, single := range people {
		if Contains(excludeList, single.Username) {
			continue
		}
		line := fmt.Sprintf("* [%s %s](https://superpeer.com/%s) - %s",
			single.FirstName,
			single.LastName,
			single.Username,
			single.ShortDescription)
		fmt.Println(line)
	}

}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

type Person struct {
	ShortDescription string  `json:"shortDescription"`
	DonatingTo       string  `json:"donatingTo"`
	FirstName        string  `json:"firstName"`
	LastName         string  `json:"lastName"`
	AvatarURL        string  `json:"avatarUrl"`
	VideoURL         string  `json:"videoUrl"`
	Username         string  `json:"username"`
	BookingCount     int     `json:"bookingCount"`
	BookingTotal     float64 `json:"bookingTotal"`
	TwitterURL       string  `json:"twitterUrl"`
}
