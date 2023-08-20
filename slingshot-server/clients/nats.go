package clients

import (
	"slingshot-server/slingshot"
	"sync"

	"github.com/nats-io/nats.go"
)

var natsConnection sync.Map

// GetNatsConnection returns a nats connection
func GetNatsConnection(id string) *nats.Conn {
	cli, ok := natsConnection.Load(id)
	if ok {
		return cli.(*nats.Conn)
	} else {
		return nil
	}
}

func CreateOrGetNatsConnection(record slingshot.NatsConfig) (*nats.Conn, error) {
	var natsCli *nats.Conn

	cli, _ := natsConnection.Load(record.Id)
	if cli == nil {
		// TODO: we need a "ParseURL" like for Redis
		natsCli, err := nats.Connect(record.Url)
		if err != nil {
			return nil, err
		}
		natsConnection.Store(record.Id, natsCli)
		//fmt.Println("🟣 NATS Debug [CreateOrGetNatsClient]", natsCli.IsConnected())
		return natsCli, nil
	} else {
		natsCli = cli.(*nats.Conn)
		return natsCli, nil
	}
}

func CreateNatsConnection(config slingshot.NatsConfig) (*nats.Conn, error) {

	natsCli, err := nats.Connect(config.Url)
	if err != nil {
		return nil, err
	}
	//fmt.Println("🟣 NATS Debug [CreateOrGetNatsClient]", natsCli.IsConnected())

	return natsCli, nil
}
