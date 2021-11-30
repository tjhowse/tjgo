package tjgo

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func FileToSlice(filename string) []string {
	var contents []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return contents
}

func Str2int(input string) int {
	output, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Fatal("Couldn't parse string as int", err)
	}
	return int(output)
}

func Poop() string {
	return "Poop"
}
