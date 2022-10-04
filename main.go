package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 2 {
		input := args[0]
		//output args[1]
		// read input files
		content, err := os.ReadFile(input)
		if err != nil {
			log.Fatal(err)
		}
		//write string to output file args[1]
		os.WriteFile(args[1], []byte(finalizeOutput(parser(content))), 0644)
	} else {
		log.Fatal("Enter input and output file")
	}
}

//                                                   functions
//----------------------------------------------------------------------------------------------------------------------------------

//this function is created to be used after the parser function to finalize and correct the output
func finalizeOutput(results []string) string {
	strResult := strings.Join(results, " ")
	trimWhite := regexp.MustCompile(`\s+`)
	strResult = trimWhite.ReplaceAllString(strResult, " ")
	//separate ?,.!:; from words with spaces between them
	symbols := []string{",", ".", ":", ";", "!", "?"}
	for _, v := range symbols {
		//this if statement corrects the apostrophe issue with it being separated from the word
		if !strings.Contains(strResult, v+"'") {
			strResult = strings.ReplaceAll(strResult, " "+v, v+" ")
		}
		//if the last character is a whitespace then remove it
		if strResult[len(strResult)-1:] == " " {
			strResult = strResult[:len(strResult)-1]
		}
	}
	for _, v := range symbols {
		strResult = strings.ReplaceAll(strResult, " "+v, v)
	}
	strResult = trimWhite.ReplaceAllString(strResult, " ")
	return strResult
}

// removeIndex removes an element from a slice which is in this case used for removing the ends in ex "(cap, 2)"" which removes "2)"
func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

// take string and remove all non digits and return int which is used to determine ex how many letters should be capitalized
func removeNonDigits(s string) int {
	re := regexp.MustCompile("[^0-9]+")
	output, err := strconv.Atoi(re.ReplaceAllString(s, ""))
	if err != nil {
		log.Fatal(err)
	}
	return output
}

// This function converts hexadecimal numbers to decimal
func hexNumToInt(hexaString string) string {
	hexaString = strings.ReplaceAll(hexaString, "0x", "")
	hexaString = strings.ReplaceAll(hexaString, "0X,", "")
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

// this function is the heart of the program, it takes the input file and parses it into a slice of strings which will be used in the finalize function
func parser(content []byte) []string {
	words := strings.Fields(string(content))
	var results []string
	var apostropheCount = 0
	for i := len(words) - 1; i >= 0; i-- {
		switch words[i] {
		case "(hex)":
			i--
			results = append(results, hexNumToInt(words[i]))
		case "(bin)":
			i--
			results = append(results, binToDec(words[i]))
		case "(up)":
			i--
			results = append(results, strings.ToUpper(words[i]))
		case "(up,":
			upMod := removeNonDigits(words[i+1])
			// apply toUpper to the next upMod words
			results = removeIndex(results, len(results)-1)
			for j := 0; j < upMod; j++ {
				i--
				results = append(results, strings.ToUpper(words[i]))
				if i == 0 {
					break
				}
			}
		case "(low)":
			i--
			results = append(results, strings.ToLower(words[i]))
		case "(low,":
			lowMod := removeNonDigits(words[i+1])
			// apply toUpper to the next upMod words
			results = removeIndex(results, len(results)-1)
			for j := 0; j < lowMod; j++ {
				if i == 0 {
					break
				}
				i--
				results = append(results, strings.ToLower(words[i]))
			}
		case "(cap,":
			capMod := removeNonDigits(words[i+1])
			// apply toUpper to the next upMod words
			results = removeIndex(results, len(results)-1)
			for j := 0; j < capMod; j++ {
				i--
				results = append(results, strings.Title(strings.ToLower(words[i])))
				if i == 0 {
					break
				}
			}
		case "(cap)":
			i--
			results = append(results, strings.Title(strings.ToLower(words[i])))
		case "a":
			// Turn a into an if next word starts with a vowel
			if strings.ContainsAny(string(words[i+1][0]), "aeiouhAEIOUH") {
				results = append(results, "an")
			} else {
				results = append(results, "a")
			}
		case "A":
			// Turn A into an if next word starts with a vowel
			if strings.ContainsAny(string(words[i+1][0]), "aeiouhAEIOUH") {
				results = append(results, "An")
			} else {
				results = append(results, "A")
			}

		case "'":
			if apostropheCount == 0 {
				words[i-1] = words[i-1] + "'"
				apostropheCount++
			} else {
				results = removeIndex(results, len(results)-1)
				results = append(results, "'"+words[i+1])
				apostropheCount = 0
			}
		default:
			if strings.Contains(words[i], "'") {
				apostropheCount++
			}
			results = append(results, words[i])
		}
	}
	for i, j := 0, len(results)-1; i < j; i, j = i+1, j-1 {
		results[i], results[j] = results[j], results[i]
	}
	return results
}
