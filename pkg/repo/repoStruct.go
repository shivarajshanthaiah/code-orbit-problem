package repo

import (
	inter "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/repo/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type ProblemRepository struct {
	DB *gorm.DB
}

func NewProblemRepository(db *gorm.DB) inter.ProblemRepoInter {
	return &ProblemRepository{
		DB: db,
	}
}

type MongoRepository struct {
	Collection *mongo.Collection
}

func NewMongoRepository(mongo *mongo.Database) inter.MongoRepoInter {
	return &MongoRepository{
		Collection: mongo.Collection("TestCases"),
	}
}
