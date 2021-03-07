/*
	Modify dup2 to print the names of all files in which each duplicated line occurs

	This solution works by getting each line of a txt file, if line exists in map, the array
	containing that word as key will be increment in position 0 to have the count of how many
	coincidences are had been found, and subsequent values are the files where the coincidence has been found.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	counts := make(map[string][]string)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)

	} else {

		for _, arg := range files {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts)

			f.Close()
		}

	}

	for word, data := range counts {

		count, _ := strconv.Atoi(data[0])

		if count > 1 {
			fmt.Printf("Word: %s\nData: %v\n\n", word, data)
		}
	}
}

func countLines(f *os.File, counts map[string][]string) {

	input := bufio.NewScanner(f)

	for input.Scan() {

		txt := input.Text()

		if len(counts[txt]) == 0 {
			counts[txt] = append(counts[txt], "0")
		}

		count, _ := strconv.Atoi(counts[txt][0])

		count++

		counts[txt][0] = strconv.Itoa(count)

		fileName := filepath.Base(f.Name())

		if counts[txt][len(counts[txt])-1] != fileName {
			counts[txt] = append(counts[txt], fileName)
		}
	}

}
