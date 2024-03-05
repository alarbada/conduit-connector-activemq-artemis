// Code generated by paramgen. DO NOT EDIT.
// Source: github.com/conduitio/conduit-connector-sdk/cmd/paramgen

package activemq

import (
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func (DestinationConfig) Parameters() map[string]sdk.Parameter {
	return map[string]sdk.Parameter{
		"destinationConfigParam": {
			Default:     "yes",
			Description: "destinationConfigParam must be either yes or no (defaults to yes).",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationInclusion{List: []string{"yes", "no"}},
			},
		},
		"global_config_param_name": {
			Default:     "",
			Description: "global_config_param_name is named global_config_param_name and needs to be provided by the user.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationRequired{},
			},
		},
	}
}
