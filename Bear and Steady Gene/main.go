package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'steadyGene' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING gene as parameter.
 */

func balanced(m map[rune]int, n int) bool {
	N := n / 4
	if m['A'] <= N && m['C'] <= N && m['G'] <= N && m['T'] <= N {
		return true
	}
	return false
}

func steadyGene(gene string, n int) int {
	// Write your code here
	answer := n
	geneMap := map[rune]int{}

	geneRune := []rune(gene)
	//Step 1. Find min substring length. It'll help to stop Sliding Window alg
	//if we find balanced part before all steps passed
	//Count number of each letter and if its number>n/4 calculate
	//how much it bigger.
	for _, r := range geneRune {
		geneMap[r]++
	}

	minLength := 0
	for _, val := range geneMap {
		if val > n/4 {
			minLength += val - n/4
		}
	}

	//define head and tail indexes
	var head, tail int = 0, 0

	//Apply Sliding Window
	for head < n && tail < n {
		//If
		if !balanced(geneMap, n) {

			geneMap[geneRune[head]] -= 1
			head++
			//fmt.Println("head go",geneMap, tail, head)
		} else {
			if head-tail == minLength {
				return minLength
			}

			answer = int(math.Min(float64(answer), float64(head-tail)))
			geneMap[geneRune[tail]] += 1
			tail++
			//fmt.Println("tail go",geneMap, tail, head,answer)
		}
	}

	//fmt.Println(geneMap, minLength)
	return answer
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int(nTemp)

	gene := readLine(reader)

	result := steadyGene(gene, n)

	fmt.Fprintf(writer, "%d\n", result)

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
