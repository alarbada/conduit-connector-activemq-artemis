// Code generated by paramgen. DO NOT EDIT.
// Source: github.com/ConduitIO/conduit-connector-sdk/tree/main/cmd/paramgen

package activemq

import (
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func (SourceConfig) Parameters() map[string]sdk.Parameter {
	return map[string]sdk.Parameter{
		"consumerWindowSize": {
			Default:     "-1",
			Description: "consumerWindowSize is the size of the consumer window. It maps to the \"consumer-window-size\" header in the STOMP SUBSCRIBE frame.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"destination": {
			Default:     "",
			Description: "destination is the name of the STOMP destination.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationRequired{},
			},
		},
		"password": {
			Default:     "",
			Description: "password is the password to use when connecting to the broker.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationRequired{},
			},
		},
		"recvTimeoutHeartbeat": {
			Default:     "2s",
			Description: "recvTimeoutHeartbeat specifies the minimum amount of time between the client expecting to receive heartbeat notifications from the server",
			Type:        sdk.ParameterTypeDuration,
			Validations: []sdk.Validation{},
		},
		"sendTimeoutHeartbeat": {
			Default:     "2s",
			Description: "sendTimeoutHeartbeat specifies the maximum amount of time between the client sending heartbeat notifications from the server",
			Type:        sdk.ParameterTypeDuration,
			Validations: []sdk.Validation{},
		},
		"subscriptionType": {
			Default:     "ANYCAST",
			Description: "subscriptionType is the subscription type. It can be either ANYCAST or MULTICAST, with ANYCAST being the default. Maps to the \"subscription-type\" header in the STOMP SUBSCRIBE frame.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"tls.caCertPath": {
			Default:     "",
			Description: "caCertPath is the path to the CA certificate file.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"tls.clientCertPath": {
			Default:     "",
			Description: "clientCertPath is the path to the client certificate file.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"tls.clientKeyPath": {
			Default:     "",
			Description: "clientKeyPath is the path to the client key file.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"tls.enabled": {
			Default:     "false",
			Description: "enabled is a flag to enable or disable TLS.",
			Type:        sdk.ParameterTypeBool,
			Validations: []sdk.Validation{},
		},
		"tls.insecureSkipVerify": {
			Default:     "false",
			Description: "insecureSkipVerify is a flag to disable server certificate verification.",
			Type:        sdk.ParameterTypeBool,
			Validations: []sdk.Validation{},
		},
		"url": {
			Default:     "",
			Description: "url is the url of the ActiveMQ Artemis broker.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationRequired{},
			},
		},
		"user": {
			Default:     "",
			Description: "user is the username to use when connecting to the broker.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationRequired{},
			},
		},
	}
}
