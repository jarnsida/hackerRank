package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	var rankWay []int32

	for i := 1; i < len(ranked); i++ {
		if ranked[i] == ranked[i-1] {
			ranked = append(ranked[:i], ranked[i+1:]...)
			i--
		}
	}

	fmt.Println(ranked)

	var currentRank, scores int32 = int32(len(ranked)), 0

	rankedSize := int32(len(ranked))
	//Iter player scores
	for i, val := range player {

		scores = val
		fmt.Println("scores ", scores)
		//Iter rank backward from last position

		for j := rankedSize - 1 - (rankedSize - currentRank); j >= 0; j-- {

			//Check Scores are maximum
			if scores >= ranked[0] {
				rankWay = append(rankWay, 1)
				break

				//Check scores less then minimum
			} else if scores < ranked[rankedSize-1] {
				rankWay = append(rankWay, rankedSize+1)
				fmt.Println("scores less then last", rankedSize, currentRank, scores)
				break

				//If scores more then current player check next

			} else if scores >= ranked[j] {
				currentRank = j + 1
				rankWay = append(rankWay[:i], currentRank)
				fmt.Println("scores more then", rankedSize, currentRank, scores)
				continue

				//If Scores equal current append curent
			} else if scores < ranked[j] {
				break
			}

		}

	}

	return rankWay
}

func main() {

	//Open file readonly  with test case
	f1, err := os.Open("test.txt")
	checkError(err)
	defer f1.Close()

	scanner := bufio.NewScanner(f1)

	//reader := bufio.NewReaderSize(ioutil.ReadFile("test.txt"), 16*1024*1024)

	stdout, err := os.Create("result.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)
	scanner.Scan()
	rankedCount, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)
	checkError(err)
	scanner.Scan()
	rankedTemp := strings.Split(strings.TrimSpace(scanner.Text()), " ")

	var ranked []int32

	for i := 0; i < int(rankedCount); i++ {
		rankedItemTemp, err := strconv.ParseInt(rankedTemp[i], 10, 64)
		checkError(err)
		rankedItem := int32(rankedItemTemp)
		ranked = append(ranked, rankedItem)
	}
	scanner.Scan()
	playerCount, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)
	checkError(err)
	scanner.Scan()
	playerTemp := strings.Split(strings.TrimSpace(scanner.Text()), " ")

	var player []int32

	for i := 0; i < int(playerCount); i++ {
		playerItemTemp, err := strconv.ParseInt(playerTemp[i], 10, 64)
		checkError(err)
		playerItem := int32(playerItemTemp)
		player = append(player, playerItem)
	}

	result := climbingLeaderboard(ranked, player)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
