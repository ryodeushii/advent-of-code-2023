package main

import (
	"bufio"
	"fmt"

	"log"
	"os"

	"github.com/ryodeushii/advent-of-code-2023/day-01/lib/trebuchet"
)

func main() {
    // read lines from file
    // create NewTrebuchet
    // call GetCalibrationSum
    // print result
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	trebuchet := trebuchet.NewTrebuchet(lines)
	fmt.Println(trebuchet.GetCalibrationSum())
}
