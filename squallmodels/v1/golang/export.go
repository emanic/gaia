package squallmodels

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// ExportIdentity represents the Identity of the object.
var ExportIdentity = elemental.Identity{
	Name:     "export",
	Category: "export",
}

// ExportsList represents a list of Exports
type ExportsList []*Export

// ContentIdentity returns the identity of the objects in the list.
func (o ExportsList) ContentIdentity() elemental.Identity {

	return ExportIdentity
}

// Copy returns a pointer to a copy the ExportsList.
func (o ExportsList) Copy() elemental.ContentIdentifiable {

	copy := append(ExportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ExportsList.
func (o ExportsList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(ExportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Export))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ExportsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ExportsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o ExportsList) Version() int {

	return 1
}

// Export represents the model of a export
type Export struct {
	// APIVersion of the api used for the exported data
	APIVersion int `json:"APIVersion" bson:"-" mapstructure:"APIVersion,omitempty"`

	// List of exported external services.
	ExternalServices []map[string]interface{} `json:"externalServices" bson:"-" mapstructure:"externalServices,omitempty"`

	// List of exported file access policies.
	FileAccessPolicies []map[string]interface{} `json:"fileAccessPolicies" bson:"-" mapstructure:"fileAccessPolicies,omitempty"`

	// List of exported file paths.
	FilePaths []map[string]interface{} `json:"filePaths" bson:"-" mapstructure:"filePaths,omitempty"`

	// Namespace of the exported data.
	Namespace string `json:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// List of exported network policies
	NetworkAccessPolicies []map[string]interface{} `json:"networkAccessPolicies" bson:"-" mapstructure:"networkAccessPolicies,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewExport returns a new *Export
func NewExport() *Export {

	return &Export{
		ModelVersion:          1,
		ExternalServices:      []map[string]interface{}{},
		FileAccessPolicies:    []map[string]interface{}{},
		FilePaths:             []map[string]interface{}{},
		NetworkAccessPolicies: []map[string]interface{}{},
	}
}

// Identity returns the Identity of the object.
func (o *Export) Identity() elemental.Identity {

	return ExportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Export) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Export) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Export) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Export) DefaultOrder() []string {

	return []string{}
}

func (o *Export) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Export) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("fileAccessPolicies", o.FileAccessPolicies); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("fileAccessPolicies", o.FileAccessPolicies); err != nil {
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
func (*Export) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ExportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ExportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Export) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ExportAttributesMap
}

// ExportAttributesMap represents the map of attribute for Export.
var ExportAttributesMap = map[string]elemental.AttributeSpecification{
	"APIVersion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "APIVersion",
		Description:    `APIVersion of the api used for the exported data`,
		Exposed:        true,
		Name:           "APIVersion",
		ReadOnly:       true,
		Type:           "integer",
	},
	"ExternalServices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ExternalServices",
		Description:    `List of exported external services.`,
		Exposed:        true,
		Name:           "externalServices",
		ReadOnly:       true,
		SubType:        "exported_data_content",
		Type:           "external",
	},
	"FileAccessPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "FileAccessPolicies",
		Description:    `List of exported file access policies.`,
		Exposed:        true,
		Name:           "fileAccessPolicies",
		Required:       true,
		SubType:        "exported_data_content",
		Type:           "external",
	},
	"FilePaths": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "FilePaths",
		Description:    `List of exported file paths.`,
		Exposed:        true,
		Name:           "filePaths",
		ReadOnly:       true,
		SubType:        "exported_data_content",
		Type:           "external",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Namespace of the exported data.`,
		Exposed:        true,
		Format:         "free",
		Name:           "namespace",
		ReadOnly:       true,
		Type:           "string",
	},
	"NetworkAccessPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NetworkAccessPolicies",
		Description:    `List of exported network policies`,
		Exposed:        true,
		Name:           "networkAccessPolicies",
		ReadOnly:       true,
		SubType:        "exported_data_content",
		Type:           "external",
	},
}

// ExportLowerCaseAttributesMap represents the map of attribute for Export.
var ExportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"apiversion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "APIVersion",
		Description:    `APIVersion of the api used for the exported data`,
		Exposed:        true,
		Name:           "APIVersion",
		ReadOnly:       true,
		Type:           "integer",
	},
	"externalservices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ExternalServices",
		Description:    `List of exported external services.`,
		Exposed:        true,
		Name:           "externalServices",
		ReadOnly:       true,
		SubType:        "exported_data_content",
		Type:           "external",
	},
	"fileaccesspolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "FileAccessPolicies",
		Description:    `List of exported file access policies.`,
		Exposed:        true,
		Name:           "fileAccessPolicies",
		Required:       true,
		SubType:        "exported_data_content",
		Type:           "external",
	},
	"filepaths": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "FilePaths",
		Description:    `List of exported file paths.`,
		Exposed:        true,
		Name:           "filePaths",
		ReadOnly:       true,
		SubType:        "exported_data_content",
		Type:           "external",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Namespace of the exported data.`,
		Exposed:        true,
		Format:         "free",
		Name:           "namespace",
		ReadOnly:       true,
		Type:           "string",
	},
	"networkaccesspolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NetworkAccessPolicies",
		Description:    `List of exported network policies`,
		Exposed:        true,
		Name:           "networkAccessPolicies",
		ReadOnly:       true,
		SubType:        "exported_data_content",
		Type:           "external",
	},
}