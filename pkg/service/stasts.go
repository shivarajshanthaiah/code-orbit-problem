package service

import (
	"fmt"

	pb "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/proto"
)

func (p *ProblemService) GetProblemStatsService(req *pb.ProblemStatsRequest) (*pb.ProblemStatsResponse, error) {
	var totalProblems, easyProblems, mediumProblems, hardProblems, premiumProblems, nonPremiumProblems int64
	typeProblemCount := make(map[string]int32)

	totalProblems = p.Repo.CountByCondition("")

	easyProblems = p.Repo.CountByDifficulty("Easy")
	mediumProblems = p.Repo.CountByDifficulty("Medium")
	hardProblems = p.Repo.CountByDifficulty("Hard")

	types, err := p.Repo.GetDistinctProblemTypes()
	if err != nil {
		return nil, err
	}
	for _, problemType := range types {
		typeCount := p.Repo.CountByCondition(fmt.Sprintf("type = '%s'", problemType))
		typeProblemCount[problemType] = int32(typeCount)
	}

	premiumProblems = p.Repo.CountByPremiumStatus(true)
	nonPremiumProblems = p.Repo.CountByPremiumStatus(false)

	return &pb.ProblemStatsResponse{
		TotalProblems:      int32(totalProblems),
		EasyProblems:       int32(easyProblems),
		MediumProblems:     int32(mediumProblems),
		HardProblems:       int32(hardProblems),
		TypeProblemCount:   typeProblemCount,
		PremiumProblems:    int32(premiumProblems),
		NonPremiumProblems: int32(nonPremiumProblems),
	}, nil
}

func (p *ProblemService) GetLeaderboardService() (*pb.LeaderboardResponse, error) {
	leaderboardEntries, err := p.Repo.GetLeaderboard()
	if err != nil {
		return nil, err
	}

	var leaderboard []*pb.LeaderboardEntry
	rank := 1
	for _, entry := range leaderboardEntries {
		leaderboard = append(leaderboard, &pb.LeaderboardEntry{
			UserId:      entry.UserID,
			SolvedCount: entry.SolvedCount,
			Rank:        int32(rank),  // Assign rank based on the loop order
		})
		rank++  // Increment rank for the next user
	}

	return &pb.LeaderboardResponse{
		Leaderboard: leaderboard,
	}, nil
}
