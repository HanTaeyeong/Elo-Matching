package matchMaking

// needed players per game
const QueueSize = 10
const ORDER = 100
const EMPTY_KEY = ""
const MAX_SCORE_RANGE = 100

var matchingQueueMap = make(map[int][]string)

func AddToQueue(userId string) {
	user := GetUserData(userId)

	userScore := int((user.Score+50)/100) * 100
	currentQueue := matchingQueueMap[MAX_SCORE_RANGE]

	if userScore < MAX_SCORE_RANGE {
		currentQueue = matchingQueueMap[userScore]
	}

	storedIndex := 0

	for index, id := range currentQueue {
		if id == EMPTY_KEY {
			currentQueue[index] = userId
			storedIndex = index
			break
		}
	}

	if storedIndex+1 >= QueueSize {
		//matching queue completed
		userIds := make([]string, QueueSize)

		for index, userId := range currentQueue {
			userIds[index] = userId
			currentQueue[index] = EMPTY_KEY
		}
		ProceedMatch(userIds)
	}
}

func InitMatchingQueue() {
	for i := 0; i <= 100; i++ {
		matchingQueueMap[i*ORDER] = make([]string, QueueSize)
	}

	//100
	//cycling queue
	//50,100,150
	//high range ,middle range, low range

	//pending users
	//constraints

	// <-[]<-person

	//first make 10 person size array[]
	// set variation.

	//include or not probability() based on average person already have

}
