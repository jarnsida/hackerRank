package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

/*
 * Complete the 'sherlockAndAnagrams' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func sherlockAndAnagrams(s string) int32 {
	// Write your code here
	var answer int32 = 0

	//increase of anagram elements size from 1 to exlusive len(s)
	for anagramSize := 1; anagramSize < len(s); anagramSize++ {

		//Declare slice of maps of runes sets with counter of every rune inclusion
		//
		//Explanation: because we looks for unordered anagrams,
		//the best way to compare them is to count number of its runes.
		//
		//For ex.:
		//   "iffa" = i:1,f:2,a:1
		//   "fifa" =f:2,i:1,a:1 ==> map of iffa runes = map of fifa runes
		//
		//Because only same-sized anagrams can be matched, we should renew
		//slice on every step of anagram incrise
		anagram := []map[string]int{}

		//Iter over S with all-sized anagrams

		for anagramPos := 0; anagramPos <= len(s)-anagramSize; anagramPos++ {
			tempMap := make(map[string]int)

			//Fill map with s[rune] = counter
			for i := 0; i < anagramSize; i++ {
				tempMap[string(s[anagramPos+i])]++

				// fmt.Println("tempMap", tempMap) //for debug
			}

			//Check is there equal maps in our slice
			for _, m := range anagram {

				if reflect.DeepEqual(m, tempMap) {
					answer++

					//  fmt.Println("match found!", answer) //for debug
				}
			}
			anagram = append(anagram, tempMap)

			//  fmt.Println("map slice", anagram) //for debug
		}
	}

	return answer
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
