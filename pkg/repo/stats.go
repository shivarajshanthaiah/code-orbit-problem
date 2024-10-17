package repo

import (
	"strconv"

	"github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/model"
)

func (r *ProblemRepository) CountByCondition(condition string) int64 {
	var count int64
	if condition != "" {
		r.DB.Model(&model.Problem{}).Where(condition).Count(&count)
	} else {
		r.DB.Model(&model.Problem{}).Count(&count)
	}
	return count
}

func (r *ProblemRepository) CountByDifficulty(difficulty string) int64 {
	return r.CountByCondition("difficulty = '" + difficulty + "'")
}

func (r *ProblemRepository) GetDistinctProblemTypes() ([]string, error) {
	var types []string
	err := r.DB.Model(&model.Problem{}).Select("DISTINCT type").Find(&types).Error
	return types, err
}

func (r *ProblemRepository) CountByPremiumStatus(isPremium bool) int64 {
	return r.CountByCondition("is_premium = " + strconv.FormatBool(isPremium))
}

func (r *ProblemRepository) GetLeaderboard() ([]model.LeaderboardEntry, error) {
	var leaderboard []model.LeaderboardEntry

	err := r.DB.Model(&model.Submission{}).
		Select("user_id, COUNT(DISTINCT problem_id) as solved_count").
		Where("status = ?", "passed").
		Group("user_id").
		Order("solved_count DESC").
		Scan(&leaderboard).Error

	if err != nil {
		return nil, err
	}
	return leaderboard, nil
}
