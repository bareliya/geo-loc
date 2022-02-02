package helper

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/citymall/geo-loc/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-redis/redis"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
)

func toD(v interface{}) (doc bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	if err != nil {
		log.Fatal(err)

	}
	return
}

func StructToInterface(input []types.SqlDbOutput) []interface{} {
	output := make([]interface{}, len(input))
	for i, s := range input {
		output[i], _ = toD(s)
	}

	return output
}

func ConverintoGeoLoc(input []types.SqlDbOutput) []redis.GeoLocation {
	len := len(input)
	output := make([]redis.GeoLocation, len)

	for i, s := range input {
		output[i].Longitude = cast.ToFloat64(s.Longitude)
		output[i].Latitude = cast.ToFloat64(s.Latitude)
		output[i].Name = s.RoamDbId
	}
	return output
}

type ResponseJSON struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Model interface{} `json:"model"`
}

func UnprocessableResponse(returnData *ResponseJSON) {
	returnData.Code = 422
	returnData.Msg = "Failure: Unprocessable data error"
	returnData.Model = nil
}

func ErrorResponse(returnData *ResponseJSON, msg string) {
	returnData.Msg = msg
	returnData.Code = 400
}

func SuccessResponse(returnData *ResponseJSON, data interface{}) {
	returnData.Msg = "success"
	returnData.Code = 200
	returnData.Model = data
}

func URLReturnResponseJson(w http.ResponseWriter, data interface{}) {
	w.Header().Del("currentUser")
	spew.Dump("================ API RESPONSE ================", data, "======== END RESPONSE ========")
	returnJson, _ := json.Marshal(data)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(returnJson)
}

func IsEmpty(s interface{}) bool {
	if s == nil || cast.ToString(s) == "" {
		return true
	}
	return false
}
