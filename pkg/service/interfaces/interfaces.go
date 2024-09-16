package interfaces

import pb "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/proto"

type ProblemServiceInter interface{
	InsertProblemService(p *pb.Problem)(*pb.ProblemResponse, error)
}