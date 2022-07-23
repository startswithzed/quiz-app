package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'") // flag name, default value, desc
	flag.Parse()

	file, err := os.Open(*csvFileName) // csvFileName gives us a pointer to a string
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}
	// now we read the csv file
	r := csv.NewReader(file)  // this func takes an io.Reader
	lines, err := r.ReadAll() // we are reading all the lines in the csv at once
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	problems := parseLines(lines)

	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		// scanf removes leading and trailing spaces so it will only catch 1 word
		// store the value in answer
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			fmt.Println("Correct!")
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d\n", correct, len(problems))

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1) // status code 1 means that something is wrong
}

// func returns a slice of problems after parsing the lines from our csv
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines)) // the outer slice in lines is the number of rows in the csv file
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
			// a: strings.Trim(line[1]) if the csv file has spaces
		}
	}
	return ret
}
