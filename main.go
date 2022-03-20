package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Stats struct {
	passed int
	failed int
	errors []string
}

func (st *Stats) printSummary() {
	fmt.Println("SUMMARY")
	fmt.Printf("# of tests: %d\n", st.failed+st.passed)
	fmt.Printf("# of passed %d\n", st.passed)
	fmt.Printf("# of failed %d\n", st.failed)
	if len(st.errors) == 0 {
		fmt.Println("No errors...")
	} else {
		fmt.Println("Logged Errors:")
		for _, str := range st.errors {
			fmt.Printf("%s\n", str)
		}
	}
}

func main() {
	workingDirectory := os.Args[1]
	cmd := exec.Command("g++", "-std=gnu++17", "-Wall", "-Wextra", workingDirectory+"/main.cpp")
	_, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	if _, err := os.Stat("a.out"); err != nil {
		panic(err)
	}

	defer func() {
		err = os.Remove("a.out")
		if err != nil {
			panic(err)
		}
	}()

	testDir := workingDirectory + "/tests"
	files, err := ioutil.ReadDir(testDir)

	if err != nil {
		panic(err)
	}
	stats := Stats{}
	for _, file := range files {
		filePath := filepath.Join(testDir, file.Name())
		body, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		text := string(body)
		idxOutput := strings.Index(text, "output")
		input := text[6:idxOutput]
		output := text[idxOutput+6:]

		prog := exec.Command("./a.out")
		stdinBuff := bytes.Buffer{}
		stdinBuff.Write([]byte(input))
		prog.Stdin = &stdinBuff

		stdout, err := prog.Output()
		if err != nil {
			panic(err)
		}

		if strings.TrimSpace(string(stdout)) == strings.TrimSpace(output) {
			stats.passed += 1
			fmt.Printf("Test %s passed\n", file.Name())
		} else {
			fmt.Printf("Test %s failed\n", file.Name())
			stats.failed += 1
			stats.errors = append(stats.errors, string(stdout))
		}
	}

	stats.printSummary()
}
