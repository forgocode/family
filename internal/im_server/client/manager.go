package client

import (
	"sync"

	"github.com/forgocode/family/internal/pkg/typed"
	"github.com/forgocode/family/internal/webservice/database/redis"
)

type clientManager struct {
	clients sync.Map
}

var manager = &clientManager{
	clients: sync.Map{},
}

func AddClient(uid string, c *typed.WebSocketClient) {

}

func getClientFromRedis(uid string) (*typed.WebSocketClient, error) {
	rs, err := redis.GetRedisClient()
	if err != nil {
		return nil, err
	}
	c := &typed.WebSocketClient{}
	err = rs.Get(uid).Scan(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func addCientToRedis() {

}

func deleteClientFromRedis() {

}

func (m *clientManager) addClient(uid string, c *typed.WebSocketClient) {
	m.clients.Store(uid, c)
}

func (m *clientManager) delClient(uid string) {
	m.clients.Delete(uid)
}

func (m *clientManager) listClient() {

}

func (m *clientManager) getClient(uid string) (*typed.WebSocketClient, error) {
	c, ok := m.clients.Load(uid)
	if ok {
		return c.(*typed.WebSocketClient), nil

	}
	cli, err := getClientFromRedis(uid)
	if err != nil {
		return nil, err
	}
	m.addClient(uid, cli)
	return cli, nil

}
