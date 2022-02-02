package types
import(
	"time"
)

type Config struct {
  Mysql    map[string]interface{}
  Mongodb  map[string]interface{}
  Redis    map[string]interface{}
}


type SqlDbOutput struct {
	Id              int       `bson:"sql_db_id,omitempty";orm:"column(id);auto"`
	RoamDbId        string    `bson:"roam_db_id,omitempty";orm:"column(roam_db_id);auto"`
	RoamDescription string    `bson:"roam_description,omitempty";orm:"column(roam_description);auto"`
	Longitude       string    `bson:"longitude,omitempty";orm:"column(longitude);auto"`
	Latitude        string    `bson:"latitude,omitempty";orm:"column(latitude);auto"`
	Altitude        string    `bson:"altitude,omitempty";orm:"column(altitude)"`
	Accuracy        string    `bson:"accuracy,omitempty";orm:"column(accuracy)"`
	Activity        string    `bson:"activity,omitempty";orm:"column(activity)"`
	Speed           string    `bson:"speed";orm:"column(speed)"`
	RecordedAt      time.Time `bson:"recorded_at,omitempty";orm:"column(recorded_at)"`


}
func (t *SqlDbOutput) TableName() string {
	return "db_roam_location"
}



type AttributeBody struct {
	Attributes []map[string]interface{} `json:"attributes"`
}



type BulkWriteInput struct {
	Key   string
	Value interface{}
}

type BulkRedisInput struct {
	Key   string
	Value string
}



