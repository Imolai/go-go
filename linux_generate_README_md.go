package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	input     = "README.src.md"
	output    = "README.md"
	inclBegin = "[["
	inclEnd   = "]]"
)

func main() {
	inFh, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer inFh.Close()
	outFh, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	defer outFh.Close()

	inScanner := bufio.NewScanner(inFh)
	inScanner.Split(bufio.ScanLines) // read line by line by splitting by line

	fmt.Printf("Generating '%v' from '%v' and sources in 'src/'.\n", output, input)
	for inScanner.Scan() { // get next token
		line := inScanner.Text() // get the most recent token
		if strings.HasPrefix(line, inclBegin) && strings.HasSuffix(line, inclEnd) {
			srcFile := strings.TrimSuffix(strings.TrimPrefix(line, inclBegin), inclEnd)
			if _, err := os.Stat(srcFile); err != nil {
				log.Fatal(err) // file does not exist
			}
			srcContent, err := ioutil.ReadFile(srcFile) // slurp file
			if err != nil {
				log.Fatal(err)
			}
			outFh.WriteString(string(srcContent))
		} else {
			outFh.WriteString(fmt.Sprintf("%v\n", line))
		}
	}

	/*

	  Known bug in Termux's Go bin:

	  fork/exec /data/data/com.termux/files/usr/bin/doctoc: no such file or directory             exit status 1

	*/

	outToc, err := exec.Command("doctoc", "--title", "## Contents", output).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(outToc))

	/*
		outHTML, errH := exec.Command("pandoc", "--metadata", "pagetitle=\"Go, Go!\"", "--metadata", "author=\"https://github.com/Imolai\"", "-s", "-o", "Go_Go.html", output).Output()
		if errH != nil {
			log.Fatal(errH)
		}
		fmt.Println(string(outHTML))

		outHTMLA, errHA := exec.Command("pandoc", "--metadata", "pagetitle=\"Go, Go, Far Away!\"", "--metadata", "author=\"https://github.com/Imolai\"", "-s", "-o", "Go_Go_Far_Away.html", "Go_Go_Far_Away.md").Output()
		if errHA != nil {
			log.Fatal(errHA)
		}
		fmt.Println(string(outHTMLA))
	*/
}
