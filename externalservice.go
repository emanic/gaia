package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

const (
	// ExternalServiceAttributeNameID represents the attribute ID.
	ExternalServiceAttributeNameID elemental.AttributeSpecificationNameKey = "externalservice/ID"

	// ExternalServiceAttributeNameAnnotation represents the attribute annotation.
	ExternalServiceAttributeNameAnnotation elemental.AttributeSpecificationNameKey = "externalservice/annotation"

	// ExternalServiceAttributeNameAssociatedTags represents the attribute associatedTags.
	ExternalServiceAttributeNameAssociatedTags elemental.AttributeSpecificationNameKey = "externalservice/associatedTags"

	// ExternalServiceAttributeNameCreatedAt represents the attribute createdAt.
	ExternalServiceAttributeNameCreatedAt elemental.AttributeSpecificationNameKey = "externalservice/createdAt"

	// ExternalServiceAttributeNameDeleted represents the attribute deleted.
	ExternalServiceAttributeNameDeleted elemental.AttributeSpecificationNameKey = "externalservice/deleted"

	// ExternalServiceAttributeNameDescription represents the attribute description.
	ExternalServiceAttributeNameDescription elemental.AttributeSpecificationNameKey = "externalservice/description"

	// ExternalServiceAttributeNameName represents the attribute name.
	ExternalServiceAttributeNameName elemental.AttributeSpecificationNameKey = "externalservice/name"

	// ExternalServiceAttributeNameNamespace represents the attribute namespace.
	ExternalServiceAttributeNameNamespace elemental.AttributeSpecificationNameKey = "externalservice/namespace"

	// ExternalServiceAttributeNameNetwork represents the attribute network.
	ExternalServiceAttributeNameNetwork elemental.AttributeSpecificationNameKey = "externalservice/network"

	// ExternalServiceAttributeNamePort represents the attribute port.
	ExternalServiceAttributeNamePort elemental.AttributeSpecificationNameKey = "externalservice/port"

	// ExternalServiceAttributeNameProtocol represents the attribute protocol.
	ExternalServiceAttributeNameProtocol elemental.AttributeSpecificationNameKey = "externalservice/protocol"

	// ExternalServiceAttributeNameStatus represents the attribute status.
	ExternalServiceAttributeNameStatus elemental.AttributeSpecificationNameKey = "externalservice/status"

	// ExternalServiceAttributeNameUpdatedAt represents the attribute updatedAt.
	ExternalServiceAttributeNameUpdatedAt elemental.AttributeSpecificationNameKey = "externalservice/updatedAt"
)

// ExternalServiceIdentity represents the Identity of the object
var ExternalServiceIdentity = elemental.Identity{
	Name:     "externalservice",
	Category: "externalservices",
}

// ExternalServicesList represents a list of ExternalServices
type ExternalServicesList []*ExternalService

// ExternalService represents the model of a externalservice
type ExternalService struct {
	// ID is the identifier of the object.
	ID string `json:"ID,omitempty" cql:"id,primarykey,omitempty"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation,omitempty" cql:"annotation,omitempty"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags,omitempty" cql:"associatedtags,omitempty"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt,omitempty" cql:"createdat,omitempty"`

	// Deleted marks if the entity has been deleted.
	Deleted bool `json:"-" cql:"deleted,omitempty"`

	// Description is the description of the object.
	Description string `json:"description,omitempty" cql:"description,omitempty"`

	// Name is the name of the entity
	Name string `json:"name,omitempty" cql:"name,omitempty"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace,omitempty" cql:"namespace,primarykey,omitempty"`

	// Network refers to either CIDR or domain name
	Network string `json:"network,omitempty" cql:"network,omitempty"`

	// Port refers to network port which could be 100-2000 or * to represent all ports
	Port string `json:"port,omitempty" cql:"port,omitempty"`

	// Protocol refers to network protocol like TCP/UDP etc or * to represent all protocols
	Protocol string `json:"protocol,omitempty" cql:"protocol,omitempty"`

	// Status of an entity
	Status enum.EntityStatus `json:"status,omitempty" cql:"status,omitempty"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt,omitempty" cql:"updatedat,omitempty"`
}

// NewExternalService returns a new *ExternalService
func NewExternalService() *ExternalService {

	return &ExternalService{
		Status: enum.Active,
	}
}

// Identity returns the Identity of the object.
func (o *ExternalService) Identity() elemental.Identity {

	return ExternalServiceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ExternalService) Identifier() string {

	return o.ID
}

func (o *ExternalService) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ExternalService) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *ExternalService) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *ExternalService) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *ExternalService) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *ExternalService) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *ExternalService) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetName returns the name of the receiver
func (o *ExternalService) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *ExternalService) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *ExternalService) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *ExternalService) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetStatus returns the status of the receiver
func (o *ExternalService) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *ExternalService) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *ExternalService) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *ExternalService) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("network", o.Network); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("port", o.Port); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("protocol", o.Protocol); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Active", "Candidate", "Disabled"}); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o ExternalService) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return ExternalServiceAttributesMap[name]
}

// ExternalServiceAttributesMap represents the map of attribute for ExternalService.
var ExternalServiceAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	ExternalServiceAttributeNameID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	ExternalServiceAttributeNameAnnotation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	ExternalServiceAttributeNameAssociatedTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "associatedTags",
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	ExternalServiceAttributeNameCreatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	ExternalServiceAttributeNameDeleted: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Filterable:     true,
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	ExternalServiceAttributeNameDescription: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	ExternalServiceAttributeNameName: elemental.AttributeSpecification{
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
	ExternalServiceAttributeNameNamespace: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	ExternalServiceAttributeNameNetwork: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "network",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	ExternalServiceAttributeNamePort: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "port",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	ExternalServiceAttributeNameProtocol: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "protocol",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	ExternalServiceAttributeNameStatus: elemental.AttributeSpecification{
		AllowedChoices: []string{"Active", "Candidate", "Disabled"},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		Stored:         true,
		SubType:        "status_enum",
		Type:           "external",
	},
	ExternalServiceAttributeNameUpdatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "updatedAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
}
