package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	values [5][5]int
	picked [5][5]bool
}

func (b *Board) setValues(row int, values string) {
	fields := strings.Fields(values)
	if len(fields) != 5 {
		log.Fatal("Wrong num of cols")
	}
	for i, element := range fields {
		b.values[row][i], _ = strconv.Atoi(element)
	}
}

func (b *Board) pick(value int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {

			if b.values[i][j] == value {
				b.picked[i][j] = true
				// println("PICKED", i, j, b.values[i][j], b.picked[i][j])
				return
			}
			// println(i, j, b.values[i][j], b.picked[i][j])
		}
	}
}

func (b *Board) bingo() bool {
	for i := 0; i < 5; i++ {
		bingoH := true
		bingoV := true
		for j := 0; j < 5; j++ {
			if !b.picked[i][j] {
				bingoH = false
			}
			if !b.picked[j][i] {
				bingoV = false
			}
			// println(j, i, b.values[j][i], b.picked[j][i])
		}
		if bingoH || bingoV {
			return true
		}
	}
	return false
}

func (b *Board) sumUnmarked() int {
	result := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.picked[i][j] {
				result = result + b.values[i][j]
			}
		}
	}
	return result
}

func main() {
	file, err := os.Open("data/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	numbers := scanner.Text()
	scanner.Scan()
	if scanner.Text() != "" {
		log.Fatal("Expected empty line")
	}

	boards := make([]Board, 0)
	for scanner.Scan() {
		b := Board{}
		b.setValues(0, scanner.Text())
		for i := 1; i < 5; i++ {
			if !scanner.Scan() {
				log.Fatal("Malformed file")
			}
			b.setValues(i, scanner.Text())
		}
		scanner.Scan()
		if scanner.Text() != "" {
			log.Fatal("Expected empty line")
		}
		boards = append(boards, b)
	}

	for _, num := range strings.Split(numbers, ",") {
		bingo := []int{}
		for i := 0; i < len(boards); i++ {
			pick, _ := strconv.Atoi(num)
			b := boards[i]
			b.pick(pick)
			boards[i] = b
			if b.bingo() {
				println("BINGO! score: ", b.sumUnmarked()*pick)
				// fmt.Printf("%v\n", boards[i])
				// tmp_boards = append(tmp_boards[:i], tmp_boards[i+1:]...)
				bingo = append(bingo, i)
			}
		}
		for i, b := range bingo {
			boards = append(boards[:b-i], boards[b-i+1:]...)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
