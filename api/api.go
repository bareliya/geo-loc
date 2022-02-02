package api

import(
	"fmt"
	"net/http"
)


func StartServer() {
	fmt.Println("Start...")
	http.HandleFunc("/citymall/geoloc",Geolochandler)
	http.ListenAndServe(":3000", nil)
}