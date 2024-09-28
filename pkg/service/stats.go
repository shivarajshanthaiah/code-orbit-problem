package service

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/proto"
)

func (pr *ProblemService) GetUserStats(ctx context.Context, p *pb.UserID) (*pb.StatsResponse, error) {
	stats, err := pr.Repo.GetUserStats(p.ID)
	if err != nil {
		return nil, err
	}

	return &pb.StatsResponse{
		UserId:                     p.ID,
		TotalAttempts:              int32(stats["total_attempts"]),
		TotalSuccessfulSubmissions: int32(stats["total_passed"]),
		TotalFailedSubmissions:     int32(stats["total_failed"]),
		EasyAttempts:               int32(stats["easy_attempts"]),
		MediumAttempts:             int32(stats["medium_attempts"]),
		HardAttempts:               int32(stats["hard_attempts"]),
	}, nil
}
