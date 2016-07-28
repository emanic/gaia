package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

const (
	// MapEdgeAttributeNameID represents the attribute ID.
	MapEdgeAttributeNameID elemental.AttributeSpecificationNameKey = "mapedge/ID"

	// MapEdgeAttributeNameDestinationID represents the attribute destinationID.
	MapEdgeAttributeNameDestinationID elemental.AttributeSpecificationNameKey = "mapedge/destinationID"

	// MapEdgeAttributeNameLabels represents the attribute labels.
	MapEdgeAttributeNameLabels elemental.AttributeSpecificationNameKey = "mapedge/labels"

	// MapEdgeAttributeNameName represents the attribute name.
	MapEdgeAttributeNameName elemental.AttributeSpecificationNameKey = "mapedge/name"

	// MapEdgeAttributeNameSourceID represents the attribute sourceID.
	MapEdgeAttributeNameSourceID elemental.AttributeSpecificationNameKey = "mapedge/sourceID"
)

// MapEdgeIdentity represents the Identity of the object
var MapEdgeIdentity = elemental.Identity{
	Name:     "mapedge",
	Category: "mapedges",
}

// MapEdgesList represents a list of MapEdges
type MapEdgesList []*MapEdge

// MapEdge represents the model of a mapedge
type MapEdge struct {
	// ID is the identifier of the object.
	ID string `json:"ID,omitempty" cql:"-"`

	// ID of the destination resource
	DestinationID string `json:"destinationID,omitempty" cql:"-"`

	// Labels provide grouping parameters
	Labels []string `json:"labels,omitempty" cql:"-"`

	// Name is the name of the entity
	Name string `json:"name,omitempty" cql:"name,omitempty"`

	// ID of the source resource
	SourceID string `json:"sourceID,omitempty" cql:"-"`
}

// NewMapEdge returns a new *MapEdge
func NewMapEdge() *MapEdge {

	return &MapEdge{}
}

// Identity returns the Identity of the object.
func (o *MapEdge) Identity() elemental.Identity {

	return MapEdgeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *MapEdge) Identifier() string {

	return o.ID
}

func (o *MapEdge) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *MapEdge) SetIdentifier(ID string) {

	o.ID = ID
}

// GetName returns the name of the receiver
func (o *MapEdge) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *MapEdge) SetName(name string) {
	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *MapEdge) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("destinationID", o.DestinationID); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("sourceID", o.SourceID); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o MapEdge) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return MapEdgeAttributesMap[name]
}

// MapEdgeAttributesMap represents the map of attribute for MapEdge.
var MapEdgeAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	MapEdgeAttributeNameID: elemental.AttributeSpecification{
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
	MapEdgeAttributeNameDestinationID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Name:           "destinationID",
		Required:       true,
		Type:           "string",
	},
	MapEdgeAttributeNameLabels: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Name:           "labels",
		Required:       true,
		SubType:        "tags_list",
		Type:           "external",
	},
	MapEdgeAttributeNameName: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	MapEdgeAttributeNameSourceID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Name:           "sourceID",
		Required:       true,
		Type:           "string",
	},
}
