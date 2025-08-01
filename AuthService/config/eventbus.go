package config

import "github.com/kelseyhightower/envconfig"

type EventbusConfig struct {
	TopicRegistration string   `envconfig:"topik_name_registration"`
	BrokerAddresses   []string `envconfig:"broker_addresses"`
}

func MustConfigEventbus() (*EventbusConfig, error) {
	var dc EventbusConfig
	err := envconfig.Process("kafka", &dc)
	if err != nil {
		return nil, err
	}
	return &dc, nil
}
