package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/proto"
)

func (ph *ProblemHandler) GetProblemStats(ctx context.Context, p *pb.ProblemStatsRequest) (*pb.ProblemStatsResponse, error) {
	response, err := ph.SVC.GetProblemStatsService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (ph *ProblemHandler) GetLeaderboard(ctx context.Context, p *pb.LeaderboardRequest) (*pb.LeaderboardResponse, error) {
	response, err := ph.SVC.GetLeaderboardService()
	if err != nil {
		return response, err
	}
	return response, nil
}
