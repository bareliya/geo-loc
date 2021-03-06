package redisdb

import (
	"context"
	//"strconv"
	//"github.com/citymall/geo-loc/types"
	"log"

	"github.com/citymall/geo-loc/helper"
	"github.com/citymall/geo-loc/util"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

type redisConnectionType struct {
	Client    *redis.Client
	Connected bool
}

var RedisConnection = &redisConnectionType{}

func (r *redisConnectionType) ConnectRedis() {
	if !r.Connected {
		log.Println("Connecting redis")

		red := util.GetConfig().Redis
		host := red["host"].(string)
		redis_password := ""
		if red["password"] != nil {
			redis_password = red["password"].(string)
		}
		redis_db := red["database"].(int)

		log.Println("host ", host)

		r.Client = redis.NewClient(&redis.Options{
			Addr:     host,
			Password: redis_password, // no password set
			DB:       redis_db,       // use default DB
			//MaxRetries: -1,
		})

		r.Connected = true

		res, err := r.Client.Ping().Result()
		if err != nil || helper.IsEmpty(res) {
			log.Fatalf("Could not connect to redis %v", err)
		}
		log.Println("ping : ", res)

	} else {

		log.Println("Already Connected")
	}
	return
}

func (r *redisConnectionType) AddGeo(geolocs *redis.GeoLocation) interface{} {
	if !r.Connected {
		r.ConnectRedis()
	}
	red := util.GetConfig().Redis
	key := red["key"].(string)

	reposne := r.Client.GeoAdd(key, geolocs)
	return reposne

}

func (r *redisConnectionType) GetDeliveryBoysWithinRadSearchDrivers(limit int, lat float64, lng float64, rad float64) []redis.GeoLocation {
	if !r.Connected {
		r.ConnectRedis()
	}
	red := util.GetConfig().Redis
	key := red["key"].(string)

	reposne, err := r.Client.GeoRadius(key, lng, lat, &redis.GeoRadiusQuery{
		Radius:      rad,
		Unit:        "km",
		WithCoord:   true,
		WithDist:    true,
		WithGeoHash: true,
		Count:       limit,
		Sort:        "ASC",
		Store:       "",
		StoreDist:   "",
	}).Result()

	if err != nil {
		log.Fatal(err)

	}

	return reposne

}
