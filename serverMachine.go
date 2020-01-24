package blcx_srvinfo

type serverMachine struct {
	Id       int32  `bson:"_id"`
	Host     string `bson:"host"`
	HttpPort int32  `bson:"http_port"`
}
