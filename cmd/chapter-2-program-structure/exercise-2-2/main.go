package main

import (
	"bufio"
	"fmt"
	tempconv "github.com/AloisCRR/go-exercises/cmd/chapter-2-program-structure/exercise-2-1"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) == 0 {

		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			t := scanner.Text()
			v := conv(t)
			fmt.Printf("Converted value %s :\n %v\n\n", t, v)
		}

		return
	}

	for _, arg := range os.Args[1:] {
		v := conv(arg)
		fmt.Printf("%v\n", v)
	}
}

func conv(value string) string {
	return printTemps(value) + printDist(value) + printWeight(value) + printTime(value)
}

func convFloat(value string) float64 {
	v, err := strconv.Atoi(value)

	if err != nil {
		panic(fmt.Sprintln("You have to provide only numbers"))
	}

	return float64(v)
}

func printTemps(value string) string {
	var v string

	v += fmt.Sprintf("%s ºC = %.2f ºF\n", value, tempconv.CToF(tempconv.Celsius(convFloat(value))))

	v += fmt.Sprintf("%s ºC = %.2f ºK\n", value, tempconv.CToK(tempconv.Celsius(convFloat(value))))

	v += fmt.Sprintf("%s ºF = %.2f ºC\n", value, tempconv.FToC(tempconv.Fahrenheit(convFloat(value))))

	v += fmt.Sprintf("%s ºF = %.2f ºK\n", value, tempconv.FToK(tempconv.Fahrenheit(convFloat(value))))

	v += fmt.Sprintf("%s ºK = %.2f ºC\n", value, tempconv.KToC(tempconv.Kelvin(convFloat(value))))

	v += fmt.Sprintf("%s ºK = %.2f ºF\n", value, tempconv.KToF(tempconv.Kelvin(convFloat(value))))

	return v
}

func printDist(value string) string {
	var v string

	v += fmt.Sprintf("%s m = %.2f ft\n", value, convFloat(value)*3.28084)

	v += fmt.Sprintf("%s m = %.2f Km\n", value, convFloat(value)/1000)

	v += fmt.Sprintf("%s m = %.2f Cm\n", value, convFloat(value)*100)

	v += fmt.Sprintf("%s ft = %.2f m\n", value, convFloat(value)/3.28084)

	v += fmt.Sprintf("%s ft = %.2f Km\n", value, convFloat(value)/(1000*3.28084))

	v += fmt.Sprintf("%s ft = %.2f Cm\n", value, convFloat(value)*30.48)

	v += fmt.Sprintf("%s Km = %.2f m\n", value, convFloat(value)*1000)

	v += fmt.Sprintf("%s Km = %.2f ft\n", value, convFloat(value)*1000*3.28084)

	v += fmt.Sprintf("%s Km = %.2f Cm\n", value, convFloat(value)*100000)

	return v
}

func printWeight(value string) string {
	var v string

	v += fmt.Sprintf("%s Lb = %.2f Kg\n", value, convFloat(value)/2.205)

	v += fmt.Sprintf("%s Lb = %.2f t\n", value, convFloat(value)/(2.205*1000))

	return v
}

func printTime(value string) string {
	var v string

	v += fmt.Sprintf("%s Sec = %.2f Days\n", value, convFloat(value)/60*60*24)

	v += fmt.Sprintf("%s Sec = %.2f Weeks\n", value, convFloat(value)/60*60*24*7)

	v += fmt.Sprintf("%s Sec = %.2f Years\n", value, convFloat(value)/60*60*24*7*4*12)

	v += fmt.Sprintf("%s Days = %.2f Weeks\n", value, convFloat(value)/7)

	v += fmt.Sprintf("%s Days = %.2f Sec\n", value, convFloat(value)*60*60*24)

	v += fmt.Sprintf("%s Days = %.2f Years\n", value, convFloat(value)/365)

	v += fmt.Sprintf("%s Weeks = %.2f Days\n", value, convFloat(value)*7)

	v += fmt.Sprintf("%s Weeks = %.2f Sec\n", value, convFloat(value)*60*60*24*7)

	v += fmt.Sprintf("%s Weeks = %.2f Years\n", value, convFloat(value)/4*12)

	v += fmt.Sprintf("%s Years = %.2f Weeks\n", value, convFloat(value)*4*12)

	v += fmt.Sprintf("%s Years = %.2f Sec\n", value, convFloat(value)*60*60*24*7*4*12)

	v += fmt.Sprintf("%s Years = %.2f Days\n", value, convFloat(value)*365)

	return v
}
