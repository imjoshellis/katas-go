package main

import (
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	var err error
	file, err := os.Open("./katas")
	check(err)
	defer file.Close()
	names, err := file.Readdirnames(0)
	check(err)

	if len(names) < 1 {
		log.Fatal("There are no katas.")
	}

	kata := names[0]

	if len(names) > 1 {
		prompt := promptui.Select{
			Label: "Select Kata",
			Items: names,
		}

		_, kata, err = prompt.Run()
		check(err)
	}

	old := fmt.Sprintf("katas/%v", kata) + "/"
	new := fmt.Sprintf("solved/%v", kata)
	err = os.Rename(old, new)
	check(err)
}
