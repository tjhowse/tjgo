package tjgo

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func (v *Vec2) Distance(i Vec2) float64 {
	return math.Abs(float64(i.X-v.X)) + math.Abs(float64(i.Y-v.Y))
}

func (v *Vec2) Scale(s float64) {
	v.X = int(float64(v.X) * s)
	v.Y = int(float64(v.Y) * s)
}

// This splits a file into a slice, one line per element.
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

// This splits a file into a 2D slice of strings based on the provided delimiter.
func FileTo2DSlice(filename string, delimiter rune) [][]string {
	var contents [][]string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), string(delimiter))
		contents = append(contents, s)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return contents
}

// This splits a file into a 2D slice of strings. Each line is processed
// with the provided regex. Each capture group of the regex produces one
// element of each slice in the result.
func FileTo2DSliceRegex(filename string, regex string) [][]string {
	var contents [][]string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	rgx := regexp.MustCompile(regex)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := DoRegexFaster(scanner.Text(), rgx)
		contents = append(contents, s)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return contents
}

func FileTo2DIntSlice(filename string, delimiter rune) [][]int {
	var contents [][]int
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), string(delimiter))
		line := []int{Str2int(s[0]), Str2int(s[1])}
		contents = append(contents, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return contents
}
func FileToIntSlice(filename string) []int {
	var contents []int
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, Str2int(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return contents
}

// See also fields: https://pkg.go.dev/strings#Fields

func Str2int(input string) int {
	output, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Fatal("Couldn't parse string as int", err)
	}
	return int(output)
}

func Int2str(input int) string {
	return strconv.Itoa(input)
}

func Poop() string {
	return "Poop2"
}

func DoRegex(input, regex string) []string {
	rgx := regexp.MustCompile(regex)
	return DoRegexFaster(input, rgx)
}

func DoRegexFaster(input string, regex *regexp.Regexp) []string {
	results := regex.FindStringSubmatch(input)
	if len(results) > 1 {
		return results[1:]
	}
	return []string{}
}

// var rgx = regexp.MustCompile(`\((.*?)\)`)

// func main() {
//     s := `(tag)SomeText`
//     rs := rgx.FindStringSubmatch(s)
//     fmt.Println(rs[1])
// }
