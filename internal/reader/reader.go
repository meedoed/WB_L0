package reader

import (
	"WB_L0/internal/models"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

var (
	configPath = "configs/reader_config.toml"
	stream     = "WH_ITEMS"
)

type natsReader struct {
	config *Config
	nc     *nats.Conn
	js     nats.JetStream
	sub    *nats.Subscription
	// cli
	//dbHelper
	logger *logrus.Logger
}

func readConfig(prefix string) *Config {
	var config Config
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		panic("Invalid config format.")
	}
	return &config
}

func NewReader() *natsReader {
	return &natsReader{}
}

func (n *natsReader) Init(configPrefix string) {

	config := readConfig(configPrefix)
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	})
	n.logger = logger
	natsconn, err := nats.Connect(config.Address)
	if err != nil {
		n.logger.Error(err)
	}

	jetstream, err := natsconn.JetStream()
	if err != nil {
		n.logger.Error(err)
		return
	}
	_, err = jetstream.AddStream(&nats.StreamConfig{
		Name:     stream,
		Subjects: []string{(fmt.Sprintf("%s.*", stream))},
	})
	if err != nil {
		n.logger.Error(err)
		return
	}

	n.nc = natsconn
	n.js = jetstream

	subscription, err := n.js.Subscribe(stream+".*", n.getMessage)
	if err != nil {
		n.logger.Error(err)
		return
	}

	n.sub = subscription
}

func (n *natsReader) Shutdown() error {
	if err := n.sub.Unsubscribe(); err != nil {
		return err
	}
	n.nc.Close()
	return nil
}

func (n *natsReader) getMessage(msg *nats.Msg) {
	orders := models.Order{}

}
