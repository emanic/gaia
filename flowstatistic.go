package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "github.com/aporeto-inc/go-kairosdb/builder"

const (
	// FlowStatisticAttributeNameID represents the attribute ID.
	FlowStatisticAttributeNameID elemental.AttributeSpecificationNameKey = "flowstatistic/ID"

	// FlowStatisticAttributeNameDatapoints represents the attribute datapoints.
	FlowStatisticAttributeNameDatapoints elemental.AttributeSpecificationNameKey = "flowstatistic/datapoints"
)

// FlowStatisticIdentity represents the Identity of the object
var FlowStatisticIdentity = elemental.Identity{
	Name:     "flowstatistic",
	Category: "flowstatistics",
}

// FlowStatisticsList represents a list of FlowStatistics
type FlowStatisticsList []*FlowStatistic

// FlowStatistic represents the model of a flowstatistic
type FlowStatistic struct {
	// ID is the identifier of the object.
	ID string `json:"ID,omitempty" cql:"-"`

	// Datapoints is a list of time/value pairs that represent the flow events over time.
	Datapoints []builder.DataPoint `json:"datapoints,omitempty" cql:"-"`
}

// NewFlowStatistic returns a new *FlowStatistic
func NewFlowStatistic() *FlowStatistic {

	return &FlowStatistic{}
}

// Identity returns the Identity of the object.
func (o *FlowStatistic) Identity() elemental.Identity {

	return FlowStatisticIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FlowStatistic) Identifier() string {

	return o.ID
}

func (o *FlowStatistic) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FlowStatistic) SetIdentifier(ID string) {

	o.ID = ID
}

// Validate valides the current information stored into the structure.
func (o *FlowStatistic) Validate() elemental.Errors {

	errors := elemental.Errors{}

	return errors
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o FlowStatistic) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return FlowStatisticAttributesMap[name]
}

// FlowStatisticAttributesMap represents the map of attribute for FlowStatistic.
var FlowStatisticAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	FlowStatisticAttributeNameID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		Type:           "string",
		Unique:         true,
	},
	FlowStatisticAttributeNameDatapoints: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Name:           "datapoints",
		SubType:        "datapoints_list",
		Type:           "external",
	},
}
