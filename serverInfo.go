package blcx_srvinfo

type SeverInfo struct {
	MachineId int32  `bson:"servermachineId"`
	ServerId  int32  `bson:"server_id"`
	Name      string `bson:"name"`
	DBID      int32  `bson:"db_id"`
	HttpUrl   string
}

type DbMgr struct {
	Start int32  `bson:"startServerId"`
	End   int32  `bson:"endServerId"`
	DbUrl string `bson:"dbURL"`
}
