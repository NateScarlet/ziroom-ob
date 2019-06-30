package watch

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/NateScarlet/ziroom-ob/pkg/db"

	"github.com/NateScarlet/ziroom-ob/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestCheckRoomStatus(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "ziroom-test-")
	assert.NoError(t, err)
	defer os.Remove(tempDir)
	db := db.Connect(tempDir)
	defer db.Close()

	newValue, oldValue, err := CheckRoomStatus(db, &api.RoomID{CityCode: "110000", ID: "62337045"})
	assert.NoError(t, err)
	assert.NotNil(t, newValue)
	assert.Nil(t, oldValue)

	newValue.Status = "test"
	err = db.WriteRoomData(newValue)
	assert.NoError(t, err)
	newValue, oldValue, err = CheckRoomStatus(db, &api.RoomID{CityCode: "110000", ID: "62337045"})
	assert.NoError(t, err)
	assert.NotNil(t, newValue)
	assert.NotNil(t, oldValue)
	assert.Equal(t, "test", oldValue.Status)
}

func TestStart(t *testing.T) {
	os.Setenv("POLL_INTERVAL", "3s")
	tempDir, err := ioutil.TempDir("", "ziroom-test-")
	assert.NoError(t, err)
	defer os.Remove(tempDir)
	db := db.Connect(tempDir)
	defer db.Close()

	done := Start(db, &api.RoomID{CityCode: "110000", ID: "62337045"})
	time.Sleep(10 * time.Second)
	done <- true
}
