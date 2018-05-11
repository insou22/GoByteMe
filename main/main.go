package main

import (
	"os"
	"bufio"
	"fmt"
	"co.insou/gobyteme/vm"
	"strings"
	"co.insou/gobyteme/compile"
	"regexp"
)

var spaceFinder = regexp.MustCompile("[\\n\\r\\s]+")

func main() {
	codeArray := prepareCode()
	code := compile.Compile(codeArray)

	vm.LoadCode(code)
	vm.Begin()

	fmt.Println("Finished with exit code", vm.GetResult())
}

func prepareCode() []string {
	file := loadFile(os.Args[1])

	lines := strings.Builder{}

	forEachLine(file, func(text string) {
		if strings.TrimSpace(text) == "" {
			return
		}

		if strings.Contains(text, ";") {
			text = strings.Split(text, ";")[0]
		}

		lines.WriteString(strings.TrimSpace(text) + "\r\n")
	})

	allLines := lines.String()

	return spaceFinder.Split(strings.TrimSpace(allLines), -1)
}

func loadFile(path string) (file *os.File) {
	file, err := os.Open(path)

	if err != nil {
		panic("File could not be found")
	}

	return
}

func forEachLine(file *os.File, op func(line string)) {
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		op(scanner.Text())
	}
}