package domain

import "github.com/go-redis/redis"

func (rc ResultsController) insertResult(user string, score int64) (int64, error) {
	Add := rc.Redis().ZAdd("leaderBoard", redis.Z{Score: float64(score), Member: user})
	added, err := Add.Result()
	if err != nil {
		return 0, err
	}
	return added, nil
}