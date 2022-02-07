package api

import (
	"github.com/citymall/geo-loc/helper"
	"github.com/citymall/geo-loc/redisdb"
)

func GetLocationFunction(limit int, longi float64, lati float64, rad float64) (returnData helper.ResponseJSON) {

	redisClient := redisdb.RedisConnection
	err:=redisClient.ConnectRedis()
	if err!=nil{
		helper.ErrorResponse(&returnData,"redis is not connected")
		return 
	}

	response ,err := redisClient.GetDeliveryBoysWithinRadSearchDrivers(limit, lati, longi, rad)
	if err!=nil{
		helper.ErrorResponse(&returnData,err.Error())
		return 
	}
	
	helper.SuccessResponse(&returnData, response)
	return
}
