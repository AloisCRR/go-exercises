/*
	Experiment to measure the difference in running time between our potentially inefficient
	versions and the one that uses strings.Join. (Section 1.6 illustrates part of the time package, and Section
	11.4 shows how to write benchmark tests for systematic performance evaluation.)
*/
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	echo2()

	echo3()
}

func echo2() {

	start := time.Now()

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	secs := time.Since(start).Seconds()

	fmt.Printf("ECHO 2 Output : %s\nECHO 2 Time : %.8f\n\n", s, secs)
}

func echo3() {
	start := time.Now()

	s := strings.Join(os.Args[1:], " ")

	secs := time.Since(start).Seconds()

	fmt.Printf("ECHO 3 Output : %s\nECHO 3 Time : %.8f\n\n", s, secs)
}
