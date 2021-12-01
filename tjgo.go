package tjgo

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

type Vec2 struct {
	X int
	Y int
}

func (v *Vec2) Add(i Vec2) Vec2 {
	return Vec2{X: v.X + i.X, Y: v.Y + i.Y}
}

func (v *Vec2) Manhattan(i Vec2) int {
	return int(math.Abs(float64(i.X-v.X)) + math.Abs(float64(i.Y-v.Y)))
}

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
	return "Poop2"
}
