package db

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/NateScarlet/ziroom-ob/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "ziroom-test-")
	assert.NoError(t, err)
	defer os.Remove(tempDir)
	db := Connect(tempDir)
	defer db.Close()
	payload := &api.RoomData{
		Code:       "CODE001",
		CityCode:   "110000",
		ID:         "12345678",
		Name:       "测试1",
		NoticeWord: "测试2",
		Status:     "test",
	}
	err = db.WriteRoomData(payload)
	assert.NoError(t, err)
	result, err := db.ReadRoomData(payload.Code)
	assert.NoError(t, err)
	assert.Equal(t, payload, result)
}
