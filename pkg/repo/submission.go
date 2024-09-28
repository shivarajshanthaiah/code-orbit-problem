package repo

import (
	"github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/model"
)

func (p *ProblemRepository) SaveSubmission(submission *model.Submission) error {
	return p.DB.Save(submission).Error
}

func (p *ProblemRepository) FetchSubmission(userID string, problemID int) (*model.Submission, error) {
	var submission model.Submission
	err := p.DB.Where("user_id = ? AND problem_id = ?", userID, problemID).First(&submission).Error
	return &submission, err
}

func (p *ProblemRepository) UpdateSubmission(submission *model.Submission) error {
	return p.DB.Model(&model.Submission{}).
		Where("user_id = ? AND problem_id = ?", submission.UserID, submission.ProblemID).
		Updates(map[string]interface{}{
			"Code":         submission.Code,
			"Status":       submission.Status,
			"AttemptCount": submission.AttemptCount,
		}).Error
}

func (p *ProblemRepository) GetUserStats(userID string) (map[string]int, error) {
	var totalAttempts int64
	var totalPassed int64
	var totalFailed int64
	var easyAttempts int64
	var mediumAttempts int64
	var hardAttempts int64

	p.DB.Model(&model.Submission{}).Where("user_id = ?", userID).Count(&totalAttempts)
	p.DB.Model(&model.Submission{}).Where("user_id = ? AND status = ?", userID, "passed").Count(&totalPassed)
	p.DB.Model(&model.Submission{}).Where("user_id = ? AND status = ?", userID, "failed").Count(&totalFailed)
	p.DB.Model(&model.Submission{}).Joins("join problems on problems.id = submissions.problem_id").Where("user_id = ? AND problems.difficulty = ?", userID, "Easy").Count(&easyAttempts)
	p.DB.Model(&model.Submission{}).Joins("join problems on problems.id = submissions.problem_id").Where("user_id = ? AND problems.difficulty = ?", userID, "Medium").Count(&mediumAttempts)
	p.DB.Model(&model.Submission{}).Joins("join problems on problems.id = submissions.problem_id").Where("user_id = ? AND problems.difficulty = ?", userID, "Hard").Count(&hardAttempts)

	return map[string]int{
		"total_attempts":  int(totalAttempts),
		"total_passed":    int(totalPassed),
		"total_failed":    int(totalFailed),
		"easy_attempts":   int(easyAttempts),
		"medium_attempts": int(mediumAttempts),
		"hard_attempts":   int(hardAttempts),
	}, nil
}
