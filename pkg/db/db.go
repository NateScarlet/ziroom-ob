package db

import (
	"encoding/json"
	"log"

	"github.com/NateScarlet/ziroom-ob/pkg/api"
	"github.com/dgraph-io/badger"
)

// Database interface
type Database struct {
	conn *badger.DB
}

// Connect to a database
func Connect(path string) *Database {
	opts := badger.DefaultOptions(path)
	opts.Truncate = true
	opts.ValueLogFileSize = 4 << 20
	conn, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	return &Database{
		conn: conn,
	}
}

// Close database
func (db *Database) Close() {
	db.conn.Close()
}

// WriteRoomData to database
func (db *Database) WriteRoomData(room *api.RoomData) error {
	return db.conn.Update(
		func(txn *badger.Txn) error {
			data, err := json.Marshal(room)
			if err != nil {
				return err
			}
			return txn.Set(roomKey(room.Code), data)
		})
}

// ReadRoomData from database
func (db *Database) ReadRoomData(code string) (ret *api.RoomData, err error) {
	err = db.conn.View(
		func(txn *badger.Txn) error {
			ret = &api.RoomData{}
			item, err := txn.Get(roomKey(code))
			if err != nil {
				return err
			}
			return item.Value(func(val []byte) error {
				return json.Unmarshal(val, ret)
			})
		})

	return
}

func roomKey(code string) []byte {
	return []byte("room/" + code)
}
