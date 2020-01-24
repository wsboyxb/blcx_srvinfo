package blcx_srvinfo

type SeverInfo struct {
	MachineId int32  `bson:"servermachineId"`
	ServerId  int32  `bson:"server_id"`
	Name      string `bson:"name"`
	HttpUrl   string
}
