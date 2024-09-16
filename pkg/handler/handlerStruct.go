package handler

import (
	pb "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/proto"
	inter "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/service/interfaces"
)

type ProblemHandler struct {
	SVC inter.ProblemServiceInter
	pb.ProblemServiceServer
}

func NewProblemHandler(svc inter.ProblemServiceInter) *ProblemHandler {
	return &ProblemHandler{
		SVC: svc,
	}
}
