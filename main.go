package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/* bugs
ex Welcome to the Brooklyn bridge (cap)1E (hex) files were added
this results in : strconv.ParseInt: parsing "(cap)1E": invalid syntax exit status 1
*/

func main() {

	args := os.Args[1:]
	if len(args) == 2 {
		input := args[0]
		// output args[1]

		// read input files
		content, err := os.ReadFile(input)
		if err != nil {
			log.Fatal(err)
		}

		words := strings.Fields(string(content))
		var results []string
		fmt.Println(words)
		for i := len(words) - 1; i >= 0; i-- {
			numberMod, _ := strconv.Atoi(words[i])
			switch words[i] {
			case "(hex)":
				i--
				results = append(results, hexaNumberToInteger(words[i]))
			case "(bin)":
				i--
				results = append(results, binToDec(words[i]))
			case "(up)":
				if numberMod != 0 {

				}
				i--
				results = append(results, strings.ToUpper(words[i]))
			case "(low)":
				i--
				results = append(results, strings.ToLower(words[i]))
			case "(cap)":
				i--
				results = append(results, strings.Title(words[i]))
			default:
				results = append(results, words[i])
			}
		}
		for i, j := 0, len(results)-1; i < j; i, j = i+1, j-1 {
			results[i], results[j] = results[j], results[i]
		}
		fmt.Println(results)
	}
}

// numbers, err := strconv.Atoi(words[i])
// switch words[i] {
// case "(hex)":
// 	results += words[i]
// case numbers > 0
// case "(hex)":
// 	if len(words[i]) > 1 {
// 		results += hexaNumberToInteger(words[i]) + " "
// 	}
// case "(bin)":
// 	results += binToDec(previousWord) + " "
// 	i++
// case "(up)":
// 	for i1 := intModifier; i > 0; i-- {
// 	}
// 	results += strings.ToUpper(previousWord) + " "
// 	i++
// case "(low)":
// 	results += strings.ToLower(previousWord) + " "
// 	i++
// case "(cap)":
// 	results += strings.Title(strings.ToLower(previousWord)) + " "
// 	i++
// 	if i == len(words)-1 {
// 		results += words[i]
// 	}
// default:
// results += words[i] + " "
// if i == len(words) {
// 	results += words[i]
//}
// 	for i := 1; i < len(words); i++ {
// 		previousWord := words[i-1]
// 		intModifier, _ := strconv.Atoi(words[i+1])
// 		switch words[i] {
// 		case "(hex)":
// 			results += hexaNumberToInteger(previousWord) + " "
// 			i++
// 			if i == len(words)-1 {
// 				results += words[i]
// 			}
// 		case "(bin)":
// 			results += binToDec(previousWord) + " "
// 			i++
// 			if i == len(words)-1 {
// 				results += words[i]
// 			}
// 		case "(up)":
// 			for i1 := intModifier; i > 0; i--{

// 			}

// 			results += strings.ToUpper(previousWord) + " "
// 			i++
// 			if i == len(words)-1 {
// 				results += words[i]
// 			}
// 		case "(low)":
// 			results += strings.ToLower(previousWord) + " "
// 			i++
// 			if i == len(words)-1 {
// 				results += words[i]
// 			}
// 		case "(cap)":
// 			results += strings.Title(strings.ToLower(previousWord)) + " "
// 			i++
// 			if i == len(words)-1 {
// 				results += words[i]
// 			}
// 		default:
// 			results += previousWord + " "
// 			if i == len(words)-1 {
// 				results += words[i]
// 			}
// 		}
// 	}
// 	fmt.Println(results)
// }

// This function converts hexadecimal numbers to decimal
func hexaNumberToInteger(hexaString string) string {
	hexaString = strings.Replace(hexaString, "0x", "", -1)
	hexaString = strings.Replace(hexaString, "0X,", "", -1)
	output, err := strconv.ParseInt(hexaString, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return strconv.Itoa(int(output))
}

// This function converts binary to decimal
func binToDec(binString string) string {
	output, err := strconv.ParseInt(binString, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return strconv.Itoa(int(output))
}

func toUpper(s string) string {
	return strings.ToUpper(s)
}
