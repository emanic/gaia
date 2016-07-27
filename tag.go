package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

const (
	TagAttributeNameID             elemental.AttributeSpecificationNameKey = "tag/ID"
	TagAttributeNameAnnotation     elemental.AttributeSpecificationNameKey = "tag/annotation"
	TagAttributeNameAssociatedTags elemental.AttributeSpecificationNameKey = "tag/associatedTags"
	TagAttributeNameCreatedAt      elemental.AttributeSpecificationNameKey = "tag/createdAt"
	TagAttributeNameDeleted        elemental.AttributeSpecificationNameKey = "tag/deleted"
	TagAttributeNameDescription    elemental.AttributeSpecificationNameKey = "tag/description"
	TagAttributeNameKey            elemental.AttributeSpecificationNameKey = "tag/key"
	TagAttributeNameKeyValuePair   elemental.AttributeSpecificationNameKey = "tag/keyValuePair"
	TagAttributeNameName           elemental.AttributeSpecificationNameKey = "tag/name"
	TagAttributeNameNamespace      elemental.AttributeSpecificationNameKey = "tag/namespace"
	TagAttributeNameOwner          elemental.AttributeSpecificationNameKey = "tag/owner"
	TagAttributeNameStatus         elemental.AttributeSpecificationNameKey = "tag/status"
	TagAttributeNameUpdatedAt      elemental.AttributeSpecificationNameKey = "tag/updatedAt"
	TagAttributeNameValue          elemental.AttributeSpecificationNameKey = "tag/value"
)

// TagIdentity represents the Identity of the object
var TagIdentity = elemental.Identity{
	Name:     "tag",
	Category: "tags",
}

// TagsList represents a list of Tags
type TagsList []*Tag

// Tag represents the model of a tag
type Tag struct {
	ID             string            `json:"ID,omitempty" cql:"id,omitempty"`
	Annotation     map[string]string `json:"annotation,omitempty" cql:"annotation,omitempty"`
	AssociatedTags []string          `json:"associatedTags,omitempty" cql:"associatedtags,omitempty"`
	CreatedAt      time.Time         `json:"createdAt,omitempty" cql:"createdat,omitempty"`
	Deleted        bool              `json:"-" cql:"deleted,omitempty"`
	Description    string            `json:"description,omitempty" cql:"description,omitempty"`
	Key            string            `json:"key,omitempty" cql:"key,primarykey,omitempty"`
	KeyValuePair   string            `json:"keyValuePair,omitempty" cql:"keyvaluepair,omitempty"`
	Name           string            `json:"name,omitempty" cql:"name,omitempty"`
	Namespace      string            `json:"namespace,omitempty" cql:"namespace,primarykey,omitempty"`
	Owner          []string          `json:"owner,omitempty" cql:"owner,omitempty"`
	Status         enum.EntityStatus `json:"status,omitempty" cql:"status,omitempty"`
	UpdatedAt      time.Time         `json:"updatedAt,omitempty" cql:"updatedat,omitempty"`
	Value          string            `json:"value,omitempty" cql:"value,primarykey,omitempty"`
}

// NewTag returns a new *Tag
func NewTag() *Tag {

	return &Tag{}
}

// Identity returns the Identity of the object.
func (o *Tag) Identity() elemental.Identity {

	return TagIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Tag) Identifier() string {

	return o.ID
}

func (o *Tag) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Tag) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *Tag) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *Tag) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *Tag) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *Tag) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *Tag) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetName returns the name of the receiver
func (o *Tag) GetName() string {
	return o.Name
}

// GetNamespace returns the namespace of the receiver
func (o *Tag) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *Tag) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetStatus returns the status of the receiver
func (o *Tag) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *Tag) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *Tag) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *Tag) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("keyValuePair", o.KeyValuePair); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Active", "Candidate", "Disabled"}); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o Tag) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return TagAttributesMap[name]
}

var TagAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	TagAttributeNameID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	TagAttributeNameAnnotation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	TagAttributeNameAssociatedTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "associatedTags",
		Stored:         true,
		SubType:        "tag_list",
		Type:           "external",
	},
	TagAttributeNameCreatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	TagAttributeNameDeleted: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Filterable:     true,
		ForeignKey:     true,
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	TagAttributeNameDescription: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Format:         "free",
		Name:           "description",
		Stored:         true,
		Type:           "string",
	},
	TagAttributeNameKey: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "key",
		Orderable:      true,
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
	},
	TagAttributeNameKeyValuePair: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Format:         "free",
		Name:           "keyValuePair",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	TagAttributeNameName: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	TagAttributeNameNamespace: elemental.AttributeSpecification{
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
	TagAttributeNameOwner: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "owner",
		Stored:         true,
		SubType:        "tag_list",
		Type:           "external",
	},
	TagAttributeNameStatus: elemental.AttributeSpecification{
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
	TagAttributeNameUpdatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "updatedAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	TagAttributeNameValue: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "value",
		Orderable:      true,
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
	},
}
