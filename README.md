Elo Match making code
How fast and accurately converges users's rank score and presumed abilityScore?

`go run main.go`

check userHistory.json files
increase or decrease instance numbers based on your computer

# Process 
1. Create mockUsers and save user informations 
2. Randomly pick user and Add to Matching queue based on users currentScore (if user A's score is 1511, will be in Queue of Score range 1450~1550)
3. if one of Matching queue reaches needed people number(10 in this case 5:5) get result of Win or lose of game based on Ability score
4. With that result, update Each user's Elo score
5. record and save the history.