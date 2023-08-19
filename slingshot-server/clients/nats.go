package clients

import (
	"slingshot-server/slingshot"
	"sync"

	"github.com/nats-io/nats.go"
)

var natsClients sync.Map

// GetNatsClient returns a nats connection
func GetNatsClient(id string) *nats.Conn {
	cli, ok := natsClients.Load(id)
	if ok {
		return cli.(*nats.Conn)
	} else {
		return nil
	}
}

func CreateOrGetNatsClient(record slingshot.NatsClientConfig) (*nats.Conn, error) {
	var natsCli *nats.Conn

	cli, _ := natsClients.Load(record.Id)
	if cli == nil {
		// TODO: we need a "ParseURL" like for Redis
		natsCli, err := nats.Connect(record.Url)
		if err != nil {
			return nil, err
		}
		natsClients.Store(record.Id, natsCli)
	} else {
		natsCli = cli.(*nats.Conn)
		return natsCli, nil
	}
	return natsCli, nil
}
