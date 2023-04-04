package matchMaking

import (
	"math"
	"math/rand"
)

const K float64 = 20

var userHistoryData map[string][]int = make(map[string][]int, NUMBER_OF_USERS)
var userHistoryCursor map[string]int = make(map[string]int, NUMBER_OF_USERS)

func InitUserHistoryData() {
	for _, user := range GetAllUserData() {
		userHistoryData[user.Id] = make([]int, user.GamePlayNumber+5)
	}
}

func GetUserHistoryData() map[string][]int {
	return userHistoryData
}
func CalcExpectedWinRate(myRating int, yourRating int) float64 {
	result := 1 / (math.Pow(10, float64(yourRating-myRating)/400) + 1)
	return result
}

func CalcScoreAfterGame(weight float64, scoreBefore int, resultPoint float64, expectedWinRate float64) int {
	Sa := scoreBefore + int(weight*(resultPoint-expectedWinRate))
	return Sa
}

func ProceedMatch(users []string) {
	userDatas := GetUserDatas(users)
	n := len(users)
	totalScoreA := 0
	totalScoreB := 0
	totalAbilityScoreA := 0
	totalAbilityScoreB := 0
	aMembers := userDatas[:n/2]
	bMembers := userDatas[n/2:]

	for _, user := range aMembers {
		totalScoreA += user.Score
		totalAbilityScoreA += user.AbilityScore
	}
	for _, user := range bMembers {
		totalScoreB += user.Score
		totalAbilityScoreB += user.AbilityScore
	}

	totalScoreA /= 5
	totalScoreB /= 5
	totalAbilityScoreA /= 5
	totalAbilityScoreB /= 5

	abilityWinRateA := CalcExpectedWinRate(totalAbilityScoreA, totalAbilityScoreB)
	expectedWinRateA := CalcExpectedWinRate(totalScoreA, totalScoreB)

	//handle result updateUserData
	resultPointA := 0.0
	resultPointB := 0.0

	if rand.Float64() >= abilityWinRateA {
		//Bwins
		resultPointB = 1.0
	} else {
		//Awins
		resultPointA = 1.0
	}

	for _, user := range aMembers {
		cursor := userHistoryCursor[user.Id]
		weight := func(cursor int) float64 {
			if cursor <= 10 {
				return K * 3
			}
			return K
		}(cursor)

		scoreAfter := CalcScoreAfterGame(weight, user.Score, resultPointA, expectedWinRateA)
		userHistoryData[user.Id][cursor] = scoreAfter
		userHistoryCursor[user.Id] += 1
		UpdateUserScore(user.Id, scoreAfter)
	}
	for _, user := range bMembers {
		cursor := userHistoryCursor[user.Id]

		weight := func(cursor int) float64 {
			if cursor <= 10 {
				return K * 3
			}
			return K
		}(cursor)

		scoreAfter := CalcScoreAfterGame(weight, user.Score, resultPointB, 1-expectedWinRateA)
		userHistoryData[user.Id][cursor] = scoreAfter
		userHistoryCursor[user.Id] += 1
		UpdateUserScore(user.Id, scoreAfter)
	}
}
