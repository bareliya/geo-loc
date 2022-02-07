package redisdb

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type RedisTestSuite struct {
	suite.Suite
}

func (s *RedisTestSuite) SetupTest() {
	err := RedisConnection.connectRedis()
	assert := assert.New(s.T())
	assert.Nil(err)

}

func (s *RedisTestSuite) TearDownSuite() {
}

func (s *RedisTestSuite) TestRedisC() {
	assert := assert.New(s.T())
	r := RedisConnection

	geoloc := redis.GeoLocation{}
	geoloc.Longitude = "77.389322"
	geoloc.Latitude = "28.6703103"
	geoloc.Name = "6151430d06f8a05418f0878cvcvc"


	res := r.AddGeo( &geoloc)

	res_2:= *redis.IntCmd
	//to do
	// initializing res_2
	assert.Equal(*res,*res_2)

	resp, err := r.GetDeliveryBoysWithinRadSearchDrivers(5,77.389322,28.6703103,50000)
	assert.Nil(err)

	resp2 :=[]redis.GeoLocation
	//todo 
	//initializing resp2
	assert.Equal(resp, resp2)

}


func TestRedis(t *testing.T) {
	suite.Run(t, new(RedisTestSuite))
}
