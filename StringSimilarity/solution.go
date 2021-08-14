package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'stringSimilarity' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func stringSimilarity(s string) int64 {
	// Write your code here

	var answer int64 = int64(len(s))
	z := make([]int64, (len(s)))
	var i, l, r int64
	//Implement z-algorithm
	for i, l, r = 1, 0, 0; i < int64(len(s)); i++ {
		if i > r {
			l, r = i, i
			for n := int64(len(s)); r < n && s[r-l] == s[r]; r++ {
			}

			z[i] = r - int64(l)
			r--
			answer += int64(z[i])
			//fmt.Println(answer)

		} else {
			k := i - l
			if z[k] < r-i+1 {
				z[i] = z[k]
				answer += int64(z[i])
			} else {
				l = i

				for n := int64(len(s)); r < n && s[r-l] == s[r]; r++ {
				}
				z[i] = r - l
				answer += int64(z[i])
				r--

			}

		}
	}

	return answer
}

func main() {
	//Open file readonly  with test case
	f1, err := os.Open("test.txt")
	checkError(err)
	defer f1.Close()

	scanner := bufio.NewScanner(f1)
	maxCapacity := 256 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	//reader := bufio.NewReaderSize(ioutil.ReadFile("test.txt"), 16*1024*1024)

	stdout, err := os.Create("result.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 256*1024)
	scanner.Scan()

	tTemp, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)
	checkError(err)
	t := int32(tTemp)
	fmt.Println(t)

	for tItr := 0; tItr < int(t); tItr++ {
		scanner.Scan()
		s := scanner.Text()
		fmt.Println(len(s))
		result := stringSimilarity(s)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

// func readLine(reader *bufio.Reader) string {
//     str, _, err := reader.ReadLine()
//     if err == io.EOF {
//         return ""
//     }

//     return strings.TrimRight(string(str), "\r\n")
// }

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
