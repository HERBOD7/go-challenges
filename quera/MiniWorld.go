package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	splitInput := strings.Split(input, " ")
	people := splitInput[0]
	mailCount := splitInput[1]
	peopleInt, _ := strconv.Atoi(people)
	mailCountInt, _ := strconv.Atoi(mailCount)
	peoplesLocation := make(map[string][]string)
	relations := make(map[string][]string)
	mails := []string{}
	for i := 0; i < peopleInt*2; i++ {
		scanner.Scan()
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		if i < peopleInt {
			name := splitLine[0]
			x := strings.Trim(splitLine[1], ",")
			y := splitLine[2]
			peoplesLocation[name] = append(peoplesLocation[name], x)
			peoplesLocation[name] = append(peoplesLocation[name], y)
		} else {
			name := splitLine[0]
			friends := splitLine[1:]
			relations[name] = append(relations[name], friends...)
		}
	}

	for i := 0; i < mailCountInt; i++ {
		scanner.Scan()
		line := scanner.Text()
		mails = append(mails, line)
	}

	move := 0
	for _, mail := range mails {
		splitMail := strings.Split(mail, " ")
		budget, _ := strconv.Atoi(splitMail[2])
		origin := splitMail[0]
		destination := splitMail[1]
		recursiveCount := 0

		sendCount := sendToNearFriend(peoplesLocation, relations, origin, destination, budget, recursiveCount)
		move = move + sendCount
		//fmt.Println(sendCount)
		fmt.Println(move)
	}
}

func sendToNearFriend(peoplesLocation, relations map[string][]string, origin string, destination string, budget int, recursiveCount int) int {
	distance := map[float64]string{}
	move := 0
	mailBudget := budget
	for _, friend := range relations[origin] {
		personLocation := peoplesLocation[friend]
		destinationLocation := peoplesLocation[destination]
		x1, _ := strconv.ParseFloat(personLocation[0], 64)
		y1, _ := strconv.ParseFloat(personLocation[1], 64)
		x2, _ := strconv.ParseFloat(destinationLocation[0], 64)
		y2, _ := strconv.ParseFloat(destinationLocation[1], 64)
		finalX := math.Abs(math.Pow(x2, 2) - math.Pow(x1, 2))
		finalY := math.Abs(math.Pow(y2, 2) - math.Pow(y1, 2))
		squareRoot := finalX + finalY
		finalDistance := math.Sqrt(squareRoot)
		distance[finalDistance] = friend
	}

	keys := []float64{}
	for key, _ := range distance {
		keys = append(keys, key)
	}
	sort.Float64s(keys)

	currentPerson := distance[keys[0]]
	if origin != destination && recursiveCount != 0 {
		move = move + 1
	}

	if recursiveCount != 0 {
		mailBudget = mailBudget - 1
	}

	if currentPerson != destination {
		if mailBudget > 0 {
			nextMove := sendToNearFriend(peoplesLocation, relations, currentPerson, destination, mailBudget, recursiveCount+1)
			move = nextMove + move
			return move
		} else {
			return 0
		}
	} else {
		return move
	}
}
