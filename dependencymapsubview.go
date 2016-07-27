package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

const (
	DependencyMapSubviewAttributeNameID           elemental.AttributeSpecificationNameKey = "dependencymapsubview/ID"
	DependencyMapSubviewAttributeNameSelector     elemental.AttributeSpecificationNameKey = "dependencymapsubview/selector"
	DependencyMapSubviewAttributeNameSubSelectors elemental.AttributeSpecificationNameKey = "dependencymapsubview/subSelectors"
	DependencyMapSubviewAttributeNameTonality     elemental.AttributeSpecificationNameKey = "dependencymapsubview/tonality"
)

// DependencyMapSubviewIdentity represents the Identity of the object
var DependencyMapSubviewIdentity = elemental.Identity{
	Name:     "dependencymapsubview",
	Category: "dependencymapsubviews",
}

// DependencyMapSubviewsList represents a list of DependencyMapSubviews
type DependencyMapSubviewsList []*DependencyMapSubview

// DependencyMapSubview represents the model of a dependencymapsubview
type DependencyMapSubview struct {
	ID           string              `json:"ID,omitempty" cql:"id,omitempty"`
	Selector     []string            `json:"selector,omitempty" cql:"selector,omitempty"`
	SubSelectors map[string][]string `json:"subSelectors,omitempty" cql:"subselectors,omitempty"`
	Tonality     string              `json:"tonality,omitempty" cql:"tonality,omitempty"`
}

// NewDependencyMapSubview returns a new *DependencyMapSubview
func NewDependencyMapSubview() *DependencyMapSubview {

	return &DependencyMapSubview{
		Selector: []string{},

		SubSelectors: map[string][]string{},
	}
}

// Identity returns the Identity of the object.
func (o *DependencyMapSubview) Identity() elemental.Identity {

	return DependencyMapSubviewIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DependencyMapSubview) Identifier() string {

	return o.ID
}

func (o *DependencyMapSubview) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DependencyMapSubview) SetIdentifier(ID string) {

	o.ID = ID
}

// Validate valides the current information stored into the structure.
func (o *DependencyMapSubview) Validate() elemental.Errors {

	errors := elemental.Errors{}

	return errors
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o DependencyMapSubview) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return DependencyMapSubviewAttributesMap[name]
}

var DependencyMapSubviewAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	DependencyMapSubviewAttributeNameID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	DependencyMapSubviewAttributeNameSelector: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Name:           "selector",
		Orderable:      true,
		Stored:         true,
		SubType:        "dependency_map_view_selector",
		Type:           "external",
	},
	DependencyMapSubviewAttributeNameSubSelectors: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Name:           "subSelectors",
		Orderable:      true,
		Stored:         true,
		SubType:        "dependency_map_view_sub_selector",
		Type:           "external",
	},
	DependencyMapSubviewAttributeNameTonality: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "tonality",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}
