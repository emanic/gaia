package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

// MapNodeTypeValue represents the possible values for attribute "type".
type MapNodeTypeValue string

const (
	// MapNodeTypeContainer represents the value Container.
	MapNodeTypeContainer MapNodeTypeValue = "Container"

	// MapNodeTypeVolume represents the value Volume.
	MapNodeTypeVolume MapNodeTypeValue = "Volume"
)

// MapNodeIdentity represents the Identity of the object
var MapNodeIdentity = elemental.Identity{
	Name:     "mapnode",
	Category: "mapnodes",
}

// MapNodesList represents a list of MapNodes
type MapNodesList []*MapNode

// ContentIdentity returns the identity of the objects in the list.
func (o MapNodesList) ContentIdentity() elemental.Identity {
	return MapNodeIdentity
}

// List convert the object to and elemental.IdentifiablesList.
func (o MapNodesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// MapNode represents the model of a mapnode
type MapNode struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description"`

	// Groups for organizing resources
	Groups []string `json:"groups" bson:"-"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name"`

	// Namespace contains the node namespace.
	Namespace string `json:"namespace" bson:"-"`

	// Status tells the status of the node
	Status string `json:"status" bson:"-"`

	// Tags contains the tags associated to the node.
	Tags []string `json:"tags" bson:"-"`

	// Type of the resource represented in the map
	Type MapNodeTypeValue `json:"type" bson:"-"`

	// VulnerabilityLevel tells the current vulnerability of the node
	VulnerabilityLevel string `json:"vulnerabilityLevel" bson:"-"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewMapNode returns a new *MapNode
func NewMapNode() *MapNode {

	return &MapNode{
		ModelVersion: 1.0,
		Groups:       []string{},
		Tags:         []string{},
	}
}

// Identity returns the Identity of the object.
func (o *MapNode) Identity() elemental.Identity {

	return MapNodeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *MapNode) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *MapNode) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *MapNode) Version() float64 {

	return 1.0
}

// Doc returns the documentation for the object
func (o *MapNode) Doc() string {
	return nodocString
}

func (o *MapNode) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetName returns the name of the receiver
func (o *MapNode) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *MapNode) SetName(name string) {
	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *MapNode) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Container", "Volume"}, true); err != nil {
		errors = append(errors, err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*MapNode) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return MapNodeAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*MapNode) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return MapNodeAttributesMap
}

// MapNodeAttributesMap represents the map of attribute for MapNode.
var MapNodeAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
		Unique:         true,
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Groups": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Groups for organizing resources`,
		Exposed:        true,
		Name:           "groups",
		ReadOnly:       true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Name is the name of the entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Namespace contains the node namespace.`,
		Exposed:        true,
		Format:         "free",
		Name:           "namespace",
		ReadOnly:       true,
		Type:           "string",
	},
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Status tells the status of the node`,
		Exposed:        true,
		Format:         "free",
		Name:           "status",
		ReadOnly:       true,
		Type:           "string",
	},
	"Tags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Tags contains the tags associated to the node.`,
		Exposed:        true,
		Name:           "tags",
		ReadOnly:       true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"Container", "Volume"},
		Autogenerated:  true,
		Description:    `Type of the resource represented in the map`,
		Exposed:        true,
		Name:           "type",
		ReadOnly:       true,
		Type:           "enum",
	},
	"VulnerabilityLevel": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `VulnerabilityLevel tells the current vulnerability of the node`,
		Exposed:        true,
		Format:         "free",
		Name:           "vulnerabilityLevel",
		ReadOnly:       true,
		Type:           "string",
	},
}
