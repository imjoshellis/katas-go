package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	var err error
	if len(os.Args) == 1 {
		log.Fatal("You must include a kata name")
		return
	}

	kata := os.Args[1]

	dir := fmt.Sprintf("katas/%v", kata) + "/"
	err = os.Mkdir(dir, 0755)
	check(err)

	copyMain(kata, dir)
	copyTest(kata, dir)
}

func copyMain(k string, dir string) {
	old := "lib/template/template.go"
	main := fmt.Sprintf("%v.go", k)
	new := dir + main
	copyFile(old, new)
}

func copyTest(k string, dir string) {
	old := "lib/template/template_test.go"
	test := fmt.Sprintf("%v_test.go", k)
	new := dir + test
	copyFile(old, new)
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
