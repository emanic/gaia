package gaia

import (
	"fmt"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TraceMode represents the model of a tracemode
type TraceMode struct {
	// Instructs the enforcers to provide an iptables trace for a processing unit.
	IPTables bool `json:"IPTables" msgpack:"IPTables" bson:"iptables" mapstructure:"IPTables,omitempty"`

	// Instructs the enforcer to send records for all
	// application-initiated connections.
	ApplicationConnections bool `json:"applicationConnections" msgpack:"applicationConnections" bson:"applicationconnections" mapstructure:"applicationConnections,omitempty"`

	// Determines the length of the time interval that the trace must be
	// enabled, using [Golang duration syntax](https://golang.org/pkg/time/#example_Duration).
	Interval string `json:"interval" msgpack:"interval" bson:"interval" mapstructure:"interval,omitempty"`

	// Instructs the enforcer to send records for all network
	// initiated connections.
	NetworkConnections bool `json:"networkConnections" msgpack:"networkConnections" bson:"networkconnections" mapstructure:"networkConnections,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTraceMode returns a new *TraceMode
func NewTraceMode() *TraceMode {

	return &TraceMode{
		ModelVersion: 1,
		Interval:     "10s",
	}
}

// BleveType implements the bleve.Classifier Interface.
func (o *TraceMode) BleveType() string {

	return "tracemode"
}

// DeepCopy returns a deep copy if the TraceMode.
func (o *TraceMode) DeepCopy() *TraceMode {

	if o == nil {
		return nil
	}

	out := &TraceMode{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TraceMode.
func (o *TraceMode) DeepCopyInto(out *TraceMode) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TraceMode: %s", err))
	}

	*out = *target.(*TraceMode)
}

// Validate valides the current information stored into the structure.
func (o *TraceMode) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
