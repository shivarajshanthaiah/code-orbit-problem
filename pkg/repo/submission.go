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
