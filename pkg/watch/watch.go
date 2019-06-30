package watch

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/NateScarlet/ziroom-ob/pkg/email"

	"github.com/NateScarlet/ziroom-ob/pkg/api"
	"github.com/NateScarlet/ziroom-ob/pkg/db"
	"github.com/dgraph-io/badger"
)

// CheckRoomStatus provide new and old room status to compare,
// and save new status to database.
func CheckRoomStatus(
	db *db.Database,
	id *api.RoomID,
) (newValue *api.RoomData, oldValue *api.RoomData, err error) {
	newValue, err = id.FetchData()
	if err != nil {
		return
	}
	defer db.WriteRoomData(newValue)
	oldValue, err = db.ReadRoomData(newValue.Code)
	if err == badger.ErrKeyNotFound {
		oldValue = nil
		err = nil
	}
	if err != nil {
		return
	}
	return
}

// Start a routine to watch room status change.
func Start(
	db *db.Database,
	id *api.RoomID,
) chan<- bool {
	done := make(chan bool)

	d, _ := time.ParseDuration(os.Getenv("POLL_INTERVAL"))
	if d == 0 {
		d = 30 * time.Second
	}
	log.Printf("Start watching: %+v", id)
	ticker := time.NewTicker(d)
	watchRoomStatus(db, id, done)
	go func() {
		for {
			select {
			case <-done:
				close(done)
				return
			case <-ticker.C:
				watchRoomStatus(db, id, done)
			}
		}
	}()

	return done
}

func watchRoomStatus(
	db *db.Database,
	id *api.RoomID,
	done chan<- bool,
) {
	u := id.RoomURL()
	link := u.String()
	newValue, oldValue, err := CheckRoomStatus(db, id)
	if err != nil {
		log.Printf("Error during watch: %+v, %s", id, err)
		email.Send(fmt.Sprintf("出错: %s", link), err.Error())
		done <- true
		return
	}
	subject := fmt.Sprintf("[%s]%s", newValue.FormatStatus(), newValue.Name)
	if oldValue == nil {
		log.Printf(subject)
	} else if newValue.Status != oldValue.Status || newValue.NoticeWord != oldValue.NoticeWord {
		log.Printf(subject)
		email.Send(subject, link)
	}
	log.Printf("Updated: %+v", id)
}

// StartAll wathing task
func StartAll(db *db.Database) {
	links := os.Getenv("ROOM_LINKS")
	linkList := regexp.MustCompile(" +").Split(links, -1)
	if len(links) == 0 || len(linkList) == 0 {
		log.Fatal("Not configured environment variable `ROOM_LINKS`, nothing to watch.")
	}

	for _, link := range linkList {
		id, err := api.ParseRoomURLString(link)
		if err != nil {
			log.Fatal(err)
		}
		Start(db, id)
	}

}
