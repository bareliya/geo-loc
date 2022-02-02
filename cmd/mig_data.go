package cmd

import (
	"fmt"
	"log"

	"github.com/citymall/geo-loc/db"
	"github.com/citymall/geo-loc/helper"
	"github.com/citymall/geo-loc/redisdb"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migdata)
}

var migdata = &cobra.Command{
	Use:   "migdata",
	Short: "Data migration from mysql to mongodb",
	Long:  "This command transfer all the data from sql table to mongodb collection",
	Run: func(cmd *cobra.Command, args []string) {

		//seting mysql connection
		mysqlClient := db.SqlConnection
		mysqlClient.ConnectMySql()

		//and fetching all the data
		allData, err := mysqlClient.GetAllDataFromSql()
		fmt.Println(allData)
		if err != nil {
			log.Fatal(err)
		}

		// seting mongodb connection
		mongodbClient := db.MongoDbConnection
		mongodbClient.ConnectMongodb()

		//and adding all the data
		docs := helper.StructToInterface(allData)
		err = mongodbClient.InserManyDoc(docs)
		if err != nil {
			log.Fatal(err)
		}

		//settign redis connection

		redisClient := redisdb.RedisConnection
		redisClient.ConnectRedis()

		//adding data in redis
		geolocs := helper.ConverintoGeoLoc(allData)

		for _, geoloc := range geolocs {
			res := redisClient.AddGeo(&geoloc)
			log.Println(res)

		}
		//res:=redisClient.AddGeo(geolocs)
		//log.Println(res)

	},
}
