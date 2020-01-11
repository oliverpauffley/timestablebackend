package domain

func (lc *LeaderboardController) Size() int64 {
	count := lc.Redis().ZCard("leaderboard")
	return count.Val()
}

func (lc *LeaderboardController) GetAll() []User {

	// Get all elements from redis
	zRangeWithScores := lc.Redis().ZRevRangeWithScores("leaderboard", 0, -1)

	var users []User
	for rank, data := range zRangeWithScores.Val() {
		member := data.Member.(string)
		user := User{Rank: int64(rank + 1), Id: member, Score: data.Score}

		users = append(users, user)
	}
	return users
}
