package api

import (
	"net/http"

	"github.com/citymall/geo-loc/helper"

	"github.com/spf13/cast"
)

//URLMapping ...
func Geolochandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		helper.URLReturnResponseJson(w, getDeliveryBoyNearBy(w, r))
	} else {
		helper.URLReturnResponseJson(w, "error")
	}
}

func getDeliveryBoyNearBy(w http.ResponseWriter, r *http.Request) (returnData helper.ResponseJSON) {

	longi := r.URL.Query().Get("longitude")
	lati := r.URL.Query().Get("latitude")
	rad := r.URL.Query().Get("rad")
	limit := r.URL.Query().Get("limit")

	if helper.IsEmpty(longi) || helper.IsEmpty(lati) || helper.IsEmpty(rad) {
		helper.ErrorResponse(&returnData, "Error: longitude or latitude is empty")
		return
	}

	longitude := cast.ToFloat64(longi)
	latitude := cast.ToFloat64(lati)
	radi := cast.ToFloat64(rad)
	lim := 5
	if !helper.IsEmpty(rad) {
		lim = cast.ToInt(limit)
	}

	returnData = GetLocationFunction(lim, longitude, latitude, radi)
	return

}
