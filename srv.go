package blcx_srvinfo

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func GetServerInfo(dbUri, dbName string) (map[int32]SeverInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))
	if err != nil {
		return nil, errors.Wrap(err, "connect db error")
	}
	defer cancel()

	db := client.Database(dbName)
	projection := bson.D{
		{"host", 1},
		{"http_port", 1},
	}

	opts := options.Find().SetProjection(projection)
	cursor, err := db.Collection("serverMachine").Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, errors.Wrap(err, "get collection serverMachine error")
	}

	var mList []serverMachine
	if err = cursor.All(ctx, &mList); err != nil {
		return nil, errors.Wrap(err, "decode serverMachine error")
	}

	//cursor.Close(ctx)
	mmap := make(map[int32]serverMachine)
	for _, m := range mList {
		mmap[m.Id] = m
	}

	//==============================
	cursor, err = db.Collection("serverInfo").Find(ctx, bson.D{})
	if err != nil {
		return nil, errors.Wrap(err, "get collection serverInfo error")
	}
	var sList []SeverInfo
	if err = cursor.All(ctx, &sList); err != nil {
		return nil, errors.Wrap(err, "decode serverInfo error")
	}

	msrv := make(map[int32]SeverInfo)
	for _, info := range sList {
		if m, isOk := mmap[info.MachineId]; isOk {
			info.HttpUrl = fmt.Sprintf("http://%s:%d/", m.Host, m.HttpPort)
			msrv[info.ServerId] = info
		}
	}

	return msrv, nil
}
