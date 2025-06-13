package config

import "github.com/kelseyhightower/envconfig"

type EventbusConfig struct {
	TopicRegistration string   `envconfig:"TOPIK_NAME_REGISTRATION"`
	BrokerAddresses   []string `envconfig:"BROKER_ADDRESSES"`
}

func MustConfigEventbus() (*EventbusConfig, error) {
	var dc EventbusConfig
	err := envconfig.Process("kafka", &dc)
	if err != nil {
		return nil, err
	}
	return &dc, nil
}
