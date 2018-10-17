package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// ActivityIdentity represents the Identity of the object.
var ActivityIdentity = elemental.Identity{
	Name:     "activity",
	Category: "activities",
	Package:  "hojo",
	Private:  false,
}

// ActivitiesList represents a list of Activities
type ActivitiesList []*Activity

// Identity returns the identity of the objects in the list.
func (o ActivitiesList) Identity() elemental.Identity {

	return ActivityIdentity
}

// Copy returns a pointer to a copy the ActivitiesList.
func (o ActivitiesList) Copy() elemental.Identifiables {

	copy := append(ActivitiesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ActivitiesList.
func (o ActivitiesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ActivitiesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Activity))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ActivitiesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ActivitiesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ActivitiesList converted to SparseActivitiesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ActivitiesList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o ActivitiesList) Version() int {

	return 1
}

// Activity represents the model of a activity
type Activity struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Claims of the user who performed the operation.
	Claims interface{} `json:"claims" bson:"claims" mapstructure:"claims,omitempty"`

	// Data of the notification.
	Data interface{} `json:"data" bson:"data" mapstructure:"data,omitempty"`

	// Date of the notification.
	Date time.Time `json:"date" bson:"date" mapstructure:"date,omitempty"`

	// Error contains the eventual error.
	Error interface{} `json:"error" bson:"error" mapstructure:"error,omitempty"`

	// Message of the notification.
	Message string `json:"message" bson:"message" mapstructure:"message,omitempty"`

	// Namespace of the notification.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Operation describe what kind of operation the notification represents.
	Operation string `json:"operation" bson:"operation" mapstructure:"operation,omitempty"`

	// OriginalData contains the eventual original data of the object that has been
	// modified.
	OriginalData interface{} `json:"originalData" bson:"originaldata" mapstructure:"originalData,omitempty"`

	// Source contains meta information about the source.
	Source string `json:"source" bson:"source" mapstructure:"source,omitempty"`

	// TargetIdentity is the Identity of the related object.
	TargetIdentity string `json:"targetIdentity" bson:"targetidentity" mapstructure:"targetIdentity,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewActivity returns a new *Activity
func NewActivity() *Activity {

	return &Activity{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Activity) Identity() elemental.Identity {

	return ActivityIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Activity) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Activity) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *Activity) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Activity) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Activity) Doc() string {
	return `Contains all the activity log that happened in a namespace. All successful or
failed actions will be available, and eventual errors as well as the claims of
the user who triggered the actiions. This log is capped and only keeps the last
50k entries by default.`
}

func (o *Activity) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Activity) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseActivity{
			ID:             &o.ID,
			Claims:         &o.Claims,
			Data:           &o.Data,
			Date:           &o.Date,
			Error:          &o.Error,
			Message:        &o.Message,
			Namespace:      &o.Namespace,
			Operation:      &o.Operation,
			OriginalData:   &o.OriginalData,
			Source:         &o.Source,
			TargetIdentity: &o.TargetIdentity,
		}
	}

	sp := &SparseActivity{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "claims":
			sp.Claims = &(o.Claims)
		case "data":
			sp.Data = &(o.Data)
		case "date":
			sp.Date = &(o.Date)
		case "error":
			sp.Error = &(o.Error)
		case "message":
			sp.Message = &(o.Message)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "operation":
			sp.Operation = &(o.Operation)
		case "originalData":
			sp.OriginalData = &(o.OriginalData)
		case "source":
			sp.Source = &(o.Source)
		case "targetIdentity":
			sp.TargetIdentity = &(o.TargetIdentity)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseActivity to the object.
func (o *Activity) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseActivity)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Claims != nil {
		o.Claims = *so.Claims
	}
	if so.Data != nil {
		o.Data = *so.Data
	}
	if so.Date != nil {
		o.Date = *so.Date
	}
	if so.Error != nil {
		o.Error = *so.Error
	}
	if so.Message != nil {
		o.Message = *so.Message
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Operation != nil {
		o.Operation = *so.Operation
	}
	if so.OriginalData != nil {
		o.OriginalData = *so.OriginalData
	}
	if so.Source != nil {
		o.Source = *so.Source
	}
	if so.TargetIdentity != nil {
		o.TargetIdentity = *so.TargetIdentity
	}
}

// Validate valides the current information stored into the structure.
func (o *Activity) Validate() error {

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

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*Activity) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ActivityAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ActivityLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Activity) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ActivityAttributesMap
}

// ActivityAttributesMap represents the map of attribute for Activity.
var ActivityAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Claims",
		Description:    `Claims of the user who performed the operation.`,
		Exposed:        true,
		Name:           "claims",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		SubType:        "raw_data",
		Type:           "external",
	},
	"Data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Data",
		Description:    `Data of the notification.`,
		Exposed:        true,
		Name:           "data",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		SubType:        "raw_data",
		Type:           "external",
	},
	"Date": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Date",
		Description:    `Date of the notification.`,
		Exposed:        true,
		Name:           "date",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"Error": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Error",
		Description:    `Error contains the eventual error.`,
		Exposed:        true,
		Name:           "error",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		SubType:        "raw_data",
		Type:           "external",
	},
	"Message": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Message",
		Description:    `Message of the notification.`,
		Exposed:        true,
		Name:           "message",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Namespace of the notification.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Operation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Operation",
		Description:    `Operation describe what kind of operation the notification represents.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operation",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"OriginalData": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "OriginalData",
		Description: `OriginalData contains the eventual original data of the object that has been
modified.`,
		Exposed:   true,
		Name:      "originalData",
		Orderable: true,
		ReadOnly:  true,
		Stored:    true,
		SubType:   "raw_data",
		Type:      "external",
	},
	"Source": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Source",
		Description:    `Source contains meta information about the source.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "source",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"TargetIdentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "TargetIdentity",
		Description:    `TargetIdentity is the Identity of the related object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "targetIdentity",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
}

// ActivityLowerCaseAttributesMap represents the map of attribute for Activity.
var ActivityLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Claims",
		Description:    `Claims of the user who performed the operation.`,
		Exposed:        true,
		Name:           "claims",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		SubType:        "raw_data",
		Type:           "external",
	},
	"data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Data",
		Description:    `Data of the notification.`,
		Exposed:        true,
		Name:           "data",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		SubType:        "raw_data",
		Type:           "external",
	},
	"date": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Date",
		Description:    `Date of the notification.`,
		Exposed:        true,
		Name:           "date",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"error": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Error",
		Description:    `Error contains the eventual error.`,
		Exposed:        true,
		Name:           "error",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		SubType:        "raw_data",
		Type:           "external",
	},
	"message": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Message",
		Description:    `Message of the notification.`,
		Exposed:        true,
		Name:           "message",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Namespace of the notification.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"operation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Operation",
		Description:    `Operation describe what kind of operation the notification represents.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operation",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"originaldata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "OriginalData",
		Description: `OriginalData contains the eventual original data of the object that has been
modified.`,
		Exposed:   true,
		Name:      "originalData",
		Orderable: true,
		ReadOnly:  true,
		Stored:    true,
		SubType:   "raw_data",
		Type:      "external",
	},
	"source": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Source",
		Description:    `Source contains meta information about the source.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "source",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"targetidentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "TargetIdentity",
		Description:    `TargetIdentity is the Identity of the related object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "targetIdentity",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
}

// SparseActivitiesList represents a list of SparseActivities
type SparseActivitiesList []*SparseActivity

// Identity returns the identity of the objects in the list.
func (o SparseActivitiesList) Identity() elemental.Identity {

	return ActivityIdentity
}

// Copy returns a pointer to a copy the SparseActivitiesList.
func (o SparseActivitiesList) Copy() elemental.Identifiables {

	copy := append(SparseActivitiesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseActivitiesList.
func (o SparseActivitiesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseActivitiesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseActivity))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseActivitiesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseActivitiesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseActivitiesList converted to ActivitiesList.
func (o SparseActivitiesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseActivitiesList) Version() int {

	return 1
}

// SparseActivity represents the sparse version of a activity.
type SparseActivity struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Claims of the user who performed the operation.
	Claims *interface{} `json:"claims,omitempty" bson:"claims" mapstructure:"claims,omitempty"`

	// Data of the notification.
	Data *interface{} `json:"data,omitempty" bson:"data" mapstructure:"data,omitempty"`

	// Date of the notification.
	Date *time.Time `json:"date,omitempty" bson:"date" mapstructure:"date,omitempty"`

	// Error contains the eventual error.
	Error *interface{} `json:"error,omitempty" bson:"error" mapstructure:"error,omitempty"`

	// Message of the notification.
	Message *string `json:"message,omitempty" bson:"message" mapstructure:"message,omitempty"`

	// Namespace of the notification.
	Namespace *string `json:"namespace,omitempty" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Operation describe what kind of operation the notification represents.
	Operation *string `json:"operation,omitempty" bson:"operation" mapstructure:"operation,omitempty"`

	// OriginalData contains the eventual original data of the object that has been
	// modified.
	OriginalData *interface{} `json:"originalData,omitempty" bson:"originaldata" mapstructure:"originalData,omitempty"`

	// Source contains meta information about the source.
	Source *string `json:"source,omitempty" bson:"source" mapstructure:"source,omitempty"`

	// TargetIdentity is the Identity of the related object.
	TargetIdentity *string `json:"targetIdentity,omitempty" bson:"targetidentity" mapstructure:"targetIdentity,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseActivity returns a new  SparseActivity.
func NewSparseActivity() *SparseActivity {
	return &SparseActivity{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseActivity) Identity() elemental.Identity {

	return ActivityIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseActivity) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseActivity) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseActivity) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseActivity) ToPlain() elemental.PlainIdentifiable {

	out := NewActivity()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Claims != nil {
		out.Claims = *o.Claims
	}
	if o.Data != nil {
		out.Data = *o.Data
	}
	if o.Date != nil {
		out.Date = *o.Date
	}
	if o.Error != nil {
		out.Error = *o.Error
	}
	if o.Message != nil {
		out.Message = *o.Message
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Operation != nil {
		out.Operation = *o.Operation
	}
	if o.OriginalData != nil {
		out.OriginalData = *o.OriginalData
	}
	if o.Source != nil {
		out.Source = *o.Source
	}
	if o.TargetIdentity != nil {
		out.TargetIdentity = *o.TargetIdentity
	}

	return out
}
