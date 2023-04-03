package matchMaking

import (
	"fmt"
	"math/rand"
)

type User struct {
	Id             string
	Score          int
	AbilityScore   int
	GamePlayNumber int
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func InitMatchMaking() {
	initUserData()
	InitMatchingQueue()
	InitUserHistoryData()

	fmt.Println("Start simulation")

	Simulate()

	WriteFileToLocal(GetUserHistoryData(), "userHistories.json")
	fmt.Println("Finish simulation")
}

func Simulate() {
	type CountData struct {
		userId    string
		gamesLeft int
	}
	countTable := make([]CountData, NUMBER_OF_USERS)
	currentTableSize := NUMBER_OF_USERS

	for i := 0; i < NUMBER_OF_USERS; i++ {
		countTable[i] = CountData{
			userId:    tempUserData[i].Id,
			gamesLeft: tempUserData[i].GamePlayNumber,
		}
	}

	for {
		if currentTableSize <= 1 {
			break
		}

		randomIndex := rand.Intn(currentTableSize - 1)

		if countTable[randomIndex].gamesLeft <= 0 {
			countTable[randomIndex] = countTable[currentTableSize-1]
			currentTableSize -= 1
		}
		countTable[randomIndex].gamesLeft -= 1
		AddToQueue(countTable[randomIndex].userId)

	}
}
