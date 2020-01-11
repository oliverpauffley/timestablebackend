package app

import (
	"context"
	"github.com/LUSHDigital/core/workers/grpcsrv"
	"github.com/go-redis/redis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"timestables/service/domain"
	proto "timestables/service/protos"
)

type Dependencies interface {
	Redis() redis.Cmdable
}

// RegisterHandlers mounts the handlers to the grpc connection
func RegisterHandlers(server *grpcsrv.Server, deps Dependencies) {
	handlers := &Handlers{deps}
	proto.RegisterTimestablesServer(server.Connection, handlers)
}

// Handlers represent a collection of GRPC method handlers.
type Handlers struct{
	Dependencies
}

// When a new question is requested the server returns a question and answer
func (h *Handlers) NewQuestion (ctx context.Context, req *proto.NewQuestionRequest) (*proto.NewQuestionResponse, error) {
	controller := domain.QuestionController{h}

	question := controller.Get(ctx)

	return QuestionsToProtos(question), nil
}

// Results collects the users results and time and adds to redis cache
func (h *Handlers) Results(ctx context.Context, req *proto.ResultsRequest) (*proto.ResultsResponse, error) {
	controller := domain.ResultsController{h}

	// Collect results from proto
	result := ResultsToDomain(req)

	rank, err := controller.AddResult(ctx, result)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.ResultsResponse{Rank:rank}, nil
}

func (h *Handlers) Leaderboard(ctx context.Context, req *proto.LeaderboardRequest) (*proto.LeaderboardResponse, error) {
	controller := domain.LeaderboardController{h}

	controller.Get(ctx)
}

