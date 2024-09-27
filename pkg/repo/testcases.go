package repo

import (
	"context"

	"github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *MongoRepository) InsertTestCases(ctx context.Context, testCasesDoc model.TestCasesDocument) (string, error) {
	result, err := r.Collection.InsertOne(ctx, testCasesDoc)
	if err != nil {
		return "", err
	}

	// Return the inserted document's ObjectID as a hex string
	testCaseID := result.InsertedID.(primitive.ObjectID).Hex()
	return testCaseID, nil
}

func (r *MongoRepository) UpdateTestCases(ctx context.Context, testCaseID string, testCasesDoc model.TestCasesDocument) error {
	objectID, err := primitive.ObjectIDFromHex(testCaseID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"problem_id": testCasesDoc.ProblemID,
			"test_cases": testCasesDoc.TestCases,
		},
	}

	_, err = r.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *MongoRepository) GetTestCasesByProblemID(problemID int32) (*model.TestCasesDocument, error) {
	var testCasesDoc model.TestCasesDocument
	filter := bson.M{"problem_id": problemID}
	err := r.Collection.FindOne(context.Background(), filter).Decode(&testCasesDoc)
	if err != nil {
		return nil, err
	}

	return &testCasesDoc, nil
}