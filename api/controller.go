package api

import (
	"github.com/citymall/geo-loc/helper"
	"github.com/citymall/geo-loc/redisdb"
)

func GetLocationFunction(limit int, longi float64, lati float64, rad float64) (returnData helper.ResponseJSON) {

	redisClient := redisdb.RedisConnection
	redisClient.ConnectRedis()

	response := redisClient.GetDeliveryBoysWithinRadSearchDrivers(limit, lati, longi, rad)
	helper.SuccessResponse(&returnData, response)
	return
}
