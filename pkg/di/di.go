package di

import (
	"log"

	"github.com/shivaraj-shanthaiah/code_orbit_problem/config"
	"github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/db"
	"github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/handler"
	"github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/repo"
	"github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/server"
	"github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/service"
)

func Init() {
	cnfg := config.LoadConfig()

	db := db.ConnectDB(cnfg)

	problemRepo := repo.NewProblemRepository(db)

	problemService := service.NewProblemService(problemRepo)

	problemHandler := handler.NewProblemHandler(problemService)

	err := server.NewGrpcProblemServer(cnfg.GrpcPort, problemHandler)
	if err != nil {
		log.Fatalf("failed to start gRPC server %v", err.Error())
	}
}
