package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	//"fmt"
	"log"
	"github.com/citymall/geo-loc/redisdb"
	"github.com/go-redis/redis"
)

func init() {
	rootCmd.AddCommand(test)
}

var test = &cobra.Command{
	Use:   "test",
	Short: "test Start",
	Long:  "test Start",
	Run: func(cmd *cobra.Command, args []string) {
		//settign redis connection
		redisClient := redisdb.RedisConnection
		err :=redisClient.ConnectRedis()
		if err!=nil{
			log.Fatal(err)
            
		}
		fmt.Println("connected")

		geoloc:=redis.GeoLocation{}
		geoloc.Longitude= 77.389322
		geoloc.Latitude= 28.6703103
		geoloc.Name= "6151430d06f8a05418f0878cvcvc"

	 res := redisClient.AddGeo( &geoloc)
	 fmt.Println(res)



		// //adding data in redis
		// geolocs := helper.ConverintoGeoLoc(allData)

		// for _, geoloc := range geolocs {
		// 	res := redisClient.AddGeo(&geoloc)

		// 	fmt.Println(res)

		// }
	},
}
