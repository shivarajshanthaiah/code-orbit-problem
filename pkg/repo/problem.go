package repo

import "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/model"

func (p *ProblemRepository) InsertProblem(Problem *model.Problem) error {
	if err := p.DB.Create(&Problem).Error; err != nil {
		return err
	}
	return nil
}

func (p *ProblemRepository) FindProblemByID(ProblemID uint) (*model.Problem, error) {
	var Problem model.Problem
	if err := p.DB.First(&Problem, ProblemID).Error; err != nil {
		return nil, err
	}
	return &Problem, nil
}

func (p *ProblemRepository) GetAllProblems() (*[]model.Problem, error) {
	var problems []model.Problem
	if err := p.DB.Find(&problems).Error; err != nil {
		return nil, err
	}
	return &problems, nil
}

func (p *ProblemRepository) GetProblemByID(problemID int32) (*model.Problem, error) {
    var problem model.Problem

    err := p.DB.Where("id = ?", problemID).First(&problem).Error
    if err != nil {
        return nil, err
    }

    return &problem, nil
}
