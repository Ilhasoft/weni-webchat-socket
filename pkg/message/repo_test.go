package message_test

import (
	"testing"
	"time"

	"github.com/ilhasoft/wwcs/pkg/db"
	"github.com/ilhasoft/wwcs/pkg/message"
	"github.com/stretchr/testify/assert"
)

var message1 = message.Message{
	Msg:         "Hello",
	ContactURN:  "test:123",
	ChannelUUID: "AbCdEf-123456-123456",
	Timestamp:   time.Now().UnixNano(),
	Direction:   "incoming",
}

var message2 = message.Message{
	Msg:         "World!",
	ContactURN:  "test:123",
	ChannelUUID: "AbCdEf-123456-123456",
	Timestamp:   time.Now().UnixNano(),
	Direction:   "incoming",
}

func TestRepo(t *testing.T) {
	mdb := db.NewDB()
	err := db.Clear(mdb)
	assert.NoError(t, err)

	repo := message.NewRepo(mdb)

	// should get 0 records
	messages, err := repo.Get(message1.ContactURN, message1.ChannelUUID)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(messages))

	// saving the first message
	err = repo.Save(message1)
	assert.NoError(t, err)

	// saving the second message
	err = repo.Save(message2)
	assert.NoError(t, err)

	// shold get the 2 messages saved before
	messages, err = repo.Get(message1.ContactURN, message1.ChannelUUID)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(messages))
}
