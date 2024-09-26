package repo

import (
	"github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/model"
	"gorm.io/gorm"
)

func (p *ProblemRepository) SaveSubmission(submission *model.Submission) error {
	return p.DB.Save(submission).Error
}

func (p *ProblemRepository) FetchSubmission(userID string, problemID int) (*model.Submission, error) {
	var submission model.Submission
	err := p.DB.Where("user_id = ? AND problem_id = ?", userID, problemID).First(&submission).Error
	return &submission, err
}

func (p *ProblemRepository) UpdateSubmission(userID string, problemID int, status string) error {
	return p.DB.Model(&model.Submission{}).
		Where("user_id = ? AND problem_id = ?", userID, problemID).
		Updates(map[string]interface{}{
			"Status":       status,
			"AttemptCount": gorm.Expr("attempt_count + 1"),
		}).Error
}
