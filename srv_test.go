package blcx_srvinfo

import (
	"log"
	"testing"
)

func TestGetServerInfo(t *testing.T) {
	dbUri, dbName := "mongodb://192.168.1.217:27017", "qgjs_loginnew"
	list, err := GetServerInfo(dbUri, dbName)
	if err != nil {
		log.Fatalln(err)
	}

	for i, info := range list {
		log.Printf("%d,%+v\n", i, info)
	}
}
