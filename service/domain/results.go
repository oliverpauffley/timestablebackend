package domain

import (
	"context"
)

type QuestionInfo struct {
	A       int64
	B       int64
	Correct bool
	Time    int64
}

type Result struct {
	Name string
	Score int64
	Answers []QuestionInfo
}

type ResultsController struct {
	Dependencies
}

func (rc ResultsController) AddResult (ctx context.Context, result Result) (int64, error) {
	rank, err := rc.insertResult(result.Name, result.Score)

	return rank, err
}