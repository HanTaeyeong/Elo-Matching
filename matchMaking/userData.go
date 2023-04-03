package matchMaking

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

const NUMBER_OF_USERS = 100

var userDataCache map[string]User = make(map[string]User)

const MAX_GAME_PLAY = 1000

const DEFAUTL_SCORE = 1500

func createMockUser() User {
	abilityScore := MaxInt(DEFAUTL_SCORE+int(DEFAUTL_SCORE/2*GetNormalFloat()), 0)
	var user = User{
		Id:             fmt.Sprintf("%s-%d", uuid.New().String(), abilityScore),
		Score:          DEFAUTL_SCORE,
		AbilityScore:   abilityScore,
		GamePlayNumber: 1000 + rand.Intn(MAX_GAME_PLAY),
	}
	return user
}

var tempUserData []User = make([]User, NUMBER_OF_USERS)

func initUserData() {

	for i := 0; i < NUMBER_OF_USERS; i++ {
		userData := createMockUser()
		tempUserData[i] = userData
		userDataCache[userData.Id] = userData
	}

	WriteFileToLocal(tempUserData, "users.json")
	fmt.Println("initialize mock users finished")
}
func GetAllUserData() map[string]User {
	return userDataCache
}

func UpdateUserScore(userId string, Score int) {
	if user, ok := userDataCache[userId]; ok {
		user.Score = Score
		userDataCache[userId] = user
	}
}

func GetUserData(userId string) User {
	return userDataCache[userId]
}

func GetUserDatas(userIds []string) []User {
	userDatas := make([]User, len(userIds))
	for i, user := range userIds {
		userDatas[i] = userDataCache[user]
	}
	return userDatas
}
