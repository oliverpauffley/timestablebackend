package domain

import "context"

type LeaderboardController struct {
	Dependencies
}

type User struct {
	Id    string
	Rank  int64
	Score float64
}

type Leaderboard struct {
	Count int64
	Users []User
}

func (lc *LeaderboardController) Get(ctx context.Context) Leaderboard {

	// Get the size of the leaderboard
	count := lc.Size()

	// Get the full leaderboard of users
	users := lc.GetAll()

	// form leaderboard
	leaderboard := Leaderboard{
		Count: count,
		Users: users,
	}

	return leaderboard
}