package db

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type MongoTestSuite struct {
	suite.Suite
}

func (s *MongoTestSuite) SetupTest() {
	err := MongoDbConnection.ConnectMongodb()
	assert := assert.New(s.T())
	assert.Nil(err)

}

func (s *MongoTestSuite) TearDownSuite() {
}

func (s *MongoTestSuite) TestMongoC() {
	assert := assert.New(s.T())
	r := MongoDbConnection

    doc:= [2]interface{}
	//todo 
	// intialize doc
	err := r.InserManyDoc( doc)
	assert.Nil(err)

}


func TestMongo(t *testing.T) {
	suite.Run(t, new(MongoTestSuite))
}
