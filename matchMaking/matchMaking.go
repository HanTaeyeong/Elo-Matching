package matchMaking

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const NUMBER_OF_INSTANCES = 4

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
	Clustering()

	WriteFileToLocal(GetUserHistoryData(), "userHistories.json")
	fmt.Println("Finish simulation")
}

func simulateUser(wg *sync.WaitGroup, number int) {
	defer wg.Done()

	const LIMIT = int(NUMBER_OF_USERS / NUMBER_OF_INSTANCES)
	type CountData struct {
		userId    string
		gamesLeft int
	}
	countTable := make([]CountData, LIMIT)
	currentTableSize := LIMIT

	for i := 0; i < LIMIT; i++ {
		countTable[i] = CountData{
			userId:    tempUserData[LIMIT*number+i].Id,
			gamesLeft: tempUserData[LIMIT*number+i].GamePlayNumber,
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
		time.Sleep(time.Duration(10+rand.Intn(10)) * time.Millisecond)
	}

	fmt.Println(number, "instance finished")
}

func Clustering() {
	var wg sync.WaitGroup

	wg.Add(NUMBER_OF_INSTANCES)

	for i := 0; i < NUMBER_OF_INSTANCES; i++ {
		go simulateUser(&wg, i)
	}

	wg.Wait()
}
