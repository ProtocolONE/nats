package nats_manager

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

type MsgHandler stan.MsgHandler
type Subscription stan.Subscription
type ConnectionLostHandler stan.ConnectionLostHandler

type Options struct {
	ServerUrls            string `envconfig:"NATS_SERVER_URLS" default:"127.0.0.1:4222"`
	ClusterId             string `envconfig:"NATS_CLUSTER_ID" default:"test-cluster"`
	ClientId              string `envconfig:"NATS_CLIENT_ID" default:"billing-server-publisher"`
	ClientName            string `envconfig:"NATS_CLIENT_NAME" default:"NATS publisher"`
	Async                 bool   `envconfig:"NATS_ASYNC" default:"false"`
	User                  string `envconfig:"NATS_USER" default:""`
	Password              string `envconfig:"NATS_PASSWORD" default:""`
	ConnectionLostHandler ConnectionLostHandler
}

type Option func(*Options)

func ServerUrls(serverUrls string) Option {
	return func(opts *Options) {
		opts.ServerUrls = serverUrls
	}
}

func ClusterId(clusterId string) Option {
	return func(opts *Options) {
		opts.ClusterId = clusterId
	}
}

func ClientId(clientId string) Option {
	return func(opts *Options) {
		opts.ClientId = clientId
	}
}

func ClientName(clientName string) Option {
	return func(opts *Options) {
		opts.ClientName = clientName
	}
}

func Async(async bool) Option {
	return func(opts *Options) {
		opts.Async = async
	}
}

func User(user string) Option {
	return func(opts *Options) {
		opts.User = user
	}
}

func Password(password string) Option {
	return func(opts *Options) {
		opts.Password = password
	}
}

func SetConnectionLostHandler(handler ConnectionLostHandler) Option {
	return func(opts *Options) {
		opts.ConnectionLostHandler = handler
	}
}

func NewNatsManager(options ...Option) (NatsManagerInterface, error) {
	opts := getConnectionOpts(options...)

	nasOpts := []nats.Option{
		nats.Name(opts.ClientName),
	}

	if opts.User != "" && opts.Password != "" {
		nasOpts = append(nasOpts, nats.UserInfo(opts.User, opts.Password))
	}

	nc, err := nats.Connect(opts.ServerUrls, nasOpts...)
	if err != nil {
		return nil, err
	}

	mb := &NatsManager{options: opts}
	mb.client, err = stan.Connect(
		opts.ClusterId,
		opts.ClientId,
		stan.NatsConn(nc),
		stan.SetConnectionLostHandler(stan.ConnectionLostHandler(opts.ConnectionLostHandler)),
	)

	if err != nil {
		return nil, err
	}

	return mb, nil
}

func (opts *Options) HasEmptySettings() bool {
	return opts.ServerUrls == "" || opts.ClientId == "" || opts.ClusterId == ""
}

func getConnectionOpts(options ...Option) *Options {
	opts := Options{}
	conn := &Options{}

	for _, opt := range options {
		opt(&opts)
	}

	if opts.HasEmptySettings() {
		_ = envconfig.Process("", conn)
	}

	if opts.ServerUrls != "" {
		conn.ServerUrls = opts.ServerUrls
	}

	if opts.ClusterId != "" {
		conn.ClusterId = opts.ClusterId
	}

	if opts.ClientId != "" {
		conn.ClientId = opts.ClientId
	}

	if opts.ClientName != "" {
		conn.ClientName = opts.ClientName
	}

	if opts.User != "" {
		conn.User = opts.User
	}

	if opts.Password != "" {
		conn.Password = opts.Password
	}

	if opts.ConnectionLostHandler != nil {
		conn.ConnectionLostHandler = opts.ConnectionLostHandler
	}

	return conn
}
