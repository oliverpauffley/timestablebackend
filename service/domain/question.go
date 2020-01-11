package domain

import (
	"context"
	"math/rand"
)

type Question struct {
	A int64
	B int64
	Answer int64
}

// QuestionController collects actions for generating new questions
type QuestionController struct {
	Dependencies
}

// Get returns a randomly generated question and calculates the answer
func (qc *QuestionController) Get(ctx context.Context) Question {
	a := rand.Int63n(13)
	b := rand.Int63n(13)
	
	answer := a * b
	
	return Question{
		A:      a,
		B:      b,
		Answer: answer,
	}
}