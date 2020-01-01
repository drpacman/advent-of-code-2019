package files

import (
	"io/ioutil"
	"strings"
)

func ReadLines(filepath string) []string {
	// read lines
	data, _ := ioutil.ReadFile(filepath)
	contents := string(data)
	return strings.Split(contents, "\n")
}

func ReadSingleLineCSV(filepath string) []string {
	// read lines
	data, _ := ioutil.ReadFile(filepath)
	contents := string(data)
	return strings.Split(contents, ",")
}

func ReadMultiLineCSV(file string) [][]string {
	// read lines
	lines := ReadLines(file)
	entries := make([][]string, len(lines))
	for i := 0; i < len(entries); i++ {
		entries[i] = strings.Split(lines[i], ",")
	}
	return entries
}
