package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/manifoldco/promptui"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	var err error
	var kata string

	// Open solved katas folder to get names
	f, err := os.Open("./solved")
	check(err)
	defer f.Close()
	solved, err := f.Readdirnames(0)
	check(err)

	if len(os.Args) == 1 {
		// If no args, prompt user with existing katas
		names := filterNames(solved)
		prompt := promptui.Select{
			Label: "Select Existing Kata",
			Items: names,
		}
		_, kata, err = prompt.Run()
		check(err)
	} else {
		// otherwise, use the provided arg
		kata = os.Args[1]
	}

	// Count how many (if any) kata exist with this name
	count := countExisting(solved, kata)

	// Set up the dir
	dir := fmt.Sprintf("katas/%v%03d", kata, count) + "/"
	err = os.Mkdir(dir, 0755)
	check(err)

	// Copy the files
	copyMain(kata, dir)
	copyTest(kata, dir)
}

func countExisting(solved []string, kata string) int {
	count := 1
	re := "^" + kata + "[0-9]+$"
	for _, n := range solved {
		if match, _ := regexp.Match(re, []byte(n)); match {
			count++
		}
	}
	return count
}

func filterNames(solved []string) []string {
	names := []string{}
	for _, name := range solved {
		match, _ := regexp.Match("[a-zA-Z]+001", []byte(name))
		if match {
			names = append(names, name[:len(name)-3])
		}
	}
	return names
}

func copyMain(k string, dir string) {
	old := "lib/template/template.go"
	main := fmt.Sprintf("%v.go", k)
	new := dir + main
	copyFile(old, new)
	fmt.Printf("Created file: %s\n", new)
}

func copyTest(k string, dir string) {
	old := "lib/template/template_test.go"
	test := fmt.Sprintf("%v_test.go", k)
	new := dir + test
	copyFile(old, new)
	fmt.Printf("Created file: %s\n", new)
}

func copyFile(old string, new string) {
	in, err := os.Open(old)
	check(err)
	defer in.Close()

	out, err := os.Create(new)
	check(err)
	defer out.Close()

	_, err = io.Copy(out, in)
	check(err)
}
