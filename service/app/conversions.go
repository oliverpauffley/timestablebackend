package app

import (
	"timestables/service/domain"
	proto "timestables/service/protos"
)

func QuestionsToProtos(question domain.Question) *proto.NewQuestionResponse {
	return &proto.NewQuestionResponse{
		A:      question.A,
		B:      question.B,
		Answer: question.Answer,
	}
}

func ResultsToDomain(message *proto.ResultsRequest) domain.Result {

	answers := make([]domain.QuestionInfo, len(message.Answers))
	for _, messageAnswer := range message.Answers {
		answer := domain.QuestionInfo{
			A:       messageAnswer.A,
			B:       messageAnswer.B,
			Correct: messageAnswer.Correct,
			Time:    messageAnswer.Time,
		}
		answers = append(answers, answer)
	}

	result := domain.Result{
		Name: message.Name,
		Score: message.Score,
		Answers:  answers,
	}
	return result
}
