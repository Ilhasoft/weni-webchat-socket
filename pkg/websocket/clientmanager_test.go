package websocket

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/ilhasoft/wwcs/config"
	"github.com/stretchr/testify/assert"
)

func TestClientManager(t *testing.T) {
	rdbOptions, err := redis.ParseURL(config.Get().RedisQueue.URL)
	assert.NoError(t, err)
	rdb := redis.NewClient(rdbOptions)
	cm := NewClientManager(rdb)

	newClientID := "foo_id_123"
	newClient := ConnectedClient{ID: newClientID}

	client, err := cm.GetConnectedClient(newClient.ID)
	assert.NoError(t, err)
	assert.Nil(t, client)

	err = cm.AddConnectedClient(newClient)
	assert.NoError(t, err)

	connClients, err := cm.GetConnectedClients()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(connClients))

	client, err = cm.GetConnectedClient(newClient.ID)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	err = cm.RemoveConnectedClient(newClient.ID)
	assert.NoError(t, err)

	err = cm.RemoveConnectedClient(newClient.ID)
	assert.NoError(t, err)

	client, err = cm.GetConnectedClient(newClient.ID)
	assert.NoError(t, err)
	assert.Nil(t, client)

	err = cm.AddConnectedClient(newClient)
	assert.NoError(t, err)
	time.Sleep(time.Second * ClientTTL)

	client, err = cm.GetConnectedClient(newClient.ID)
	assert.NoError(t, err)
	assert.Nil(t, client)

	rdb.Del(context.TODO(), "connected_clients")
}
