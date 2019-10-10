package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// SAMLProviderIdentity represents the Identity of the object.
var SAMLProviderIdentity = elemental.Identity{
	Name:     "samlprovider",
	Category: "samlproviders",
	Package:  "cactuar",
	Private:  false,
}

// SAMLProvidersList represents a list of SAMLProviders
type SAMLProvidersList []*SAMLProvider

// Identity returns the identity of the objects in the list.
func (o SAMLProvidersList) Identity() elemental.Identity {

	return SAMLProviderIdentity
}

// Copy returns a pointer to a copy the SAMLProvidersList.
func (o SAMLProvidersList) Copy() elemental.Identifiables {

	copy := append(SAMLProvidersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SAMLProvidersList.
func (o SAMLProvidersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SAMLProvidersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SAMLProvider))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SAMLProvidersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SAMLProvidersList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the SAMLProvidersList converted to SparseSAMLProvidersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o SAMLProvidersList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseSAMLProvidersList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseSAMLProvider)
	}

	return out
}

// Version returns the version of the content.
func (o SAMLProvidersList) Version() int {

	return 1
}

// SAMLProvider represents the model of a samlprovider
type SAMLProvider struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Identity Provider Certificate in PEM format.
	IDPCertificate string `json:"IDPCertificate" msgpack:"IDPCertificate" bson:"idpcertificate" mapstructure:"IDPCertificate,omitempty"`

	// Identity Provider Issuer (also called Entity ID).
	IDPIssuer string `json:"IDPIssuer" msgpack:"IDPIssuer" bson:"idpissuer" mapstructure:"IDPIssuer,omitempty"`

	// Pass some XML data containing the IDP metadata that can be used for automatic
	// configuration. If you pass this attribute, every other one will be overwritten
	// with the data contained in the metadata file.
	IDPMetadata string `json:"IDPMetadata,omitempty" msgpack:"IDPMetadata,omitempty" bson:"-" mapstructure:"IDPMetadata,omitempty"`

	// URL of the identity provider.
	IDPURL string `json:"IDPURL" msgpack:"IDPURL" bson:"idpurl" mapstructure:"IDPURL,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// If set, this will be the default SAML provider. There can be only one default
	// provider in your account. When logging in with SAML, if no provider name is
	// given, the default will be used.
	Default bool `json:"default" msgpack:"default" bson:"default" mapstructure:"default,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// List of claims that will provide the subject.
	Subjects []string `json:"subjects" msgpack:"subjects" bson:"subjects" mapstructure:"subjects,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone int `json:"zone" msgpack:"zone" bson:"zone" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSAMLProvider returns a new *SAMLProvider
func NewSAMLProvider() *SAMLProvider {

	return &SAMLProvider{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		MigrationsLog:  map[string]string{},
		NormalizedTags: []string{},
		Subjects:       []string{},
	}
}

// Identity returns the Identity of the object.
func (o *SAMLProvider) Identity() elemental.Identity {

	return SAMLProviderIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SAMLProvider) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SAMLProvider) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SAMLProvider) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSAMLProvider{}

	s.ID = bson.ObjectIdHex(o.ID)
	s.IDPCertificate = o.IDPCertificate
	s.IDPIssuer = o.IDPIssuer
	s.IDPURL = o.IDPURL
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.Default = o.Default
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.Protected = o.Protected
	s.Subjects = o.Subjects
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SAMLProvider) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSAMLProvider{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.IDPCertificate = s.IDPCertificate
	o.IDPIssuer = s.IDPIssuer
	o.IDPURL = s.IDPURL
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.Default = s.Default
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.Protected = s.Protected
	o.Subjects = s.Subjects
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SAMLProvider) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *SAMLProvider) BleveType() string {

	return "samlprovider"
}

// DefaultOrder returns the list of default ordering fields.
func (o *SAMLProvider) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *SAMLProvider) Doc() string {

	return `Allows to declare a generic SAML provider that can be used in
exchange for a Midgard token.`
}

func (o *SAMLProvider) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SAMLProvider) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *SAMLProvider) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SAMLProvider) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *SAMLProvider) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SAMLProvider) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *SAMLProvider) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SAMLProvider) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *SAMLProvider) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SAMLProvider) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *SAMLProvider) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SAMLProvider) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *SAMLProvider) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SAMLProvider) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *SAMLProvider) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SAMLProvider) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *SAMLProvider) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SAMLProvider) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *SAMLProvider) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SAMLProvider) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *SAMLProvider) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SAMLProvider) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *SAMLProvider) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SAMLProvider) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *SAMLProvider) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *SAMLProvider) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *SAMLProvider) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *SAMLProvider) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseSAMLProvider{
			ID:                   &o.ID,
			IDPCertificate:       &o.IDPCertificate,
			IDPIssuer:            &o.IDPIssuer,
			IDPMetadata:          &o.IDPMetadata,
			IDPURL:               &o.IDPURL,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Default:              &o.Default,
			MigrationsLog:        &o.MigrationsLog,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Protected:            &o.Protected,
			Subjects:             &o.Subjects,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseSAMLProvider{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "IDPCertificate":
			sp.IDPCertificate = &(o.IDPCertificate)
		case "IDPIssuer":
			sp.IDPIssuer = &(o.IDPIssuer)
		case "IDPMetadata":
			sp.IDPMetadata = &(o.IDPMetadata)
		case "IDPURL":
			sp.IDPURL = &(o.IDPURL)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "default":
			sp.Default = &(o.Default)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "protected":
			sp.Protected = &(o.Protected)
		case "subjects":
			sp.Subjects = &(o.Subjects)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseSAMLProvider to the object.
func (o *SAMLProvider) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseSAMLProvider)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.IDPCertificate != nil {
		o.IDPCertificate = *so.IDPCertificate
	}
	if so.IDPIssuer != nil {
		o.IDPIssuer = *so.IDPIssuer
	}
	if so.IDPMetadata != nil {
		o.IDPMetadata = *so.IDPMetadata
	}
	if so.IDPURL != nil {
		o.IDPURL = *so.IDPURL
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Default != nil {
		o.Default = *so.Default
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Subjects != nil {
		o.Subjects = *so.Subjects
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the SAMLProvider.
func (o *SAMLProvider) DeepCopy() *SAMLProvider {

	if o == nil {
		return nil
	}

	out := &SAMLProvider{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SAMLProvider.
func (o *SAMLProvider) DeepCopyInto(out *SAMLProvider) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SAMLProvider: %s", err))
	}

	*out = *target.(*SAMLProvider)
}

// Validate valides the current information stored into the structure.
func (o *SAMLProvider) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
	}

	// Custom object validation.
	if err := ValidateSAMLProvider(o); err != nil {
		errors = errors.Append(err)
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
func (*SAMLProvider) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := SAMLProviderAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return SAMLProviderLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*SAMLProvider) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return SAMLProviderAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *SAMLProvider) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "IDPCertificate":
		return o.IDPCertificate
	case "IDPIssuer":
		return o.IDPIssuer
	case "IDPMetadata":
		return o.IDPMetadata
	case "IDPURL":
		return o.IDPURL
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "default":
		return o.Default
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "protected":
		return o.Protected
	case "subjects":
		return o.Subjects
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// SAMLProviderAttributesMap represents the map of attribute for SAMLProvider.
var SAMLProviderAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"IDPCertificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IDPCertificate",
		Description:    `Identity Provider Certificate in PEM format.`,
		Exposed:        true,
		Name:           "IDPCertificate",
		Stored:         true,
		Type:           "string",
	},
	"IDPIssuer": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IDPIssuer",
		Description:    `Identity Provider Issuer (also called Entity ID).`,
		Exposed:        true,
		Name:           "IDPIssuer",
		Stored:         true,
		Type:           "string",
	},
	"IDPMetadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IDPMetadata",
		Description: `Pass some XML data containing the IDP metadata that can be used for automatic
configuration. If you pass this attribute, every other one will be overwritten
with the data contained in the metadata file.`,
		Exposed: true,
		Name:    "IDPMetadata",
		Type:    "string",
	},
	"IDPURL": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IDPURL",
		Description:    `URL of the identity provider.`,
		Exposed:        true,
		Name:           "IDPURL",
		Stored:         true,
		Type:           "string",
	},
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"CreateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Default": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Default",
		Description: `If set, this will be the default SAML provider. There can be only one default
provider in your account. When logging in with SAML, if no provider name is
given, the default will be used.`,
		Exposed: true,
		Name:    "default",
		Stored:  true,
		Type:    "boolean",
	},
	"MigrationsLog": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Subjects": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subjects",
		Description:    `List of claims that will provide the subject.`,
		Exposed:        true,
		Name:           "subjects",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"UpdateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"ZHash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SAMLProviderLowerCaseAttributesMap represents the map of attribute for SAMLProvider.
var SAMLProviderLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"idpcertificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IDPCertificate",
		Description:    `Identity Provider Certificate in PEM format.`,
		Exposed:        true,
		Name:           "IDPCertificate",
		Stored:         true,
		Type:           "string",
	},
	"idpissuer": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IDPIssuer",
		Description:    `Identity Provider Issuer (also called Entity ID).`,
		Exposed:        true,
		Name:           "IDPIssuer",
		Stored:         true,
		Type:           "string",
	},
	"idpmetadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IDPMetadata",
		Description: `Pass some XML data containing the IDP metadata that can be used for automatic
configuration. If you pass this attribute, every other one will be overwritten
with the data contained in the metadata file.`,
		Exposed: true,
		Name:    "IDPMetadata",
		Type:    "string",
	},
	"idpurl": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IDPURL",
		Description:    `URL of the identity provider.`,
		Exposed:        true,
		Name:           "IDPURL",
		Stored:         true,
		Type:           "string",
	},
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"associatedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"createidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"default": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Default",
		Description: `If set, this will be the default SAML provider. There can be only one default
provider in your account. When logging in with SAML, if no provider name is
given, the default will be used.`,
		Exposed: true,
		Name:    "default",
		Stored:  true,
		Type:    "boolean",
	},
	"migrationslog": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"normalizedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"subjects": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subjects",
		Description:    `List of claims that will provide the subject.`,
		Exposed:        true,
		Name:           "subjects",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"updateidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"zhash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparseSAMLProvidersList represents a list of SparseSAMLProviders
type SparseSAMLProvidersList []*SparseSAMLProvider

// Identity returns the identity of the objects in the list.
func (o SparseSAMLProvidersList) Identity() elemental.Identity {

	return SAMLProviderIdentity
}

// Copy returns a pointer to a copy the SparseSAMLProvidersList.
func (o SparseSAMLProvidersList) Copy() elemental.Identifiables {

	copy := append(SparseSAMLProvidersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseSAMLProvidersList.
func (o SparseSAMLProvidersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseSAMLProvidersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseSAMLProvider))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseSAMLProvidersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseSAMLProvidersList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseSAMLProvidersList converted to SAMLProvidersList.
func (o SparseSAMLProvidersList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseSAMLProvidersList) Version() int {

	return 1
}

// SparseSAMLProvider represents the sparse version of a samlprovider.
type SparseSAMLProvider struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Identity Provider Certificate in PEM format.
	IDPCertificate *string `json:"IDPCertificate,omitempty" msgpack:"IDPCertificate,omitempty" bson:"idpcertificate,omitempty" mapstructure:"IDPCertificate,omitempty"`

	// Identity Provider Issuer (also called Entity ID).
	IDPIssuer *string `json:"IDPIssuer,omitempty" msgpack:"IDPIssuer,omitempty" bson:"idpissuer,omitempty" mapstructure:"IDPIssuer,omitempty"`

	// Pass some XML data containing the IDP metadata that can be used for automatic
	// configuration. If you pass this attribute, every other one will be overwritten
	// with the data contained in the metadata file.
	IDPMetadata *string `json:"IDPMetadata,omitempty" msgpack:"IDPMetadata,omitempty" bson:"-" mapstructure:"IDPMetadata,omitempty"`

	// URL of the identity provider.
	IDPURL *string `json:"IDPURL,omitempty" msgpack:"IDPURL,omitempty" bson:"idpurl,omitempty" mapstructure:"IDPURL,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// If set, this will be the default SAML provider. There can be only one default
	// provider in your account. When logging in with SAML, if no provider name is
	// given, the default will be used.
	Default *bool `json:"default,omitempty" msgpack:"default,omitempty" bson:"default,omitempty" mapstructure:"default,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// List of claims that will provide the subject.
	Subjects *[]string `json:"subjects,omitempty" msgpack:"subjects,omitempty" bson:"subjects,omitempty" mapstructure:"subjects,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone *int `json:"zone,omitempty" msgpack:"zone,omitempty" bson:"zone,omitempty" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseSAMLProvider returns a new  SparseSAMLProvider.
func NewSparseSAMLProvider() *SparseSAMLProvider {
	return &SparseSAMLProvider{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseSAMLProvider) Identity() elemental.Identity {

	return SAMLProviderIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseSAMLProvider) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseSAMLProvider) SetIdentifier(id string) {

	o.ID = &id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseSAMLProvider) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseSAMLProvider{}

	s.ID = bson.ObjectIdHex(*o.ID)
	if o.IDPCertificate != nil {
		s.IDPCertificate = o.IDPCertificate
	}
	if o.IDPIssuer != nil {
		s.IDPIssuer = o.IDPIssuer
	}
	if o.IDPURL != nil {
		s.IDPURL = o.IDPURL
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.Default != nil {
		s.Default = o.Default
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.Subjects != nil {
		s.Subjects = o.Subjects
	}
	if o.UpdateIdempotencyKey != nil {
		s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		s.UpdateTime = o.UpdateTime
	}
	if o.ZHash != nil {
		s.ZHash = o.ZHash
	}
	if o.Zone != nil {
		s.Zone = o.Zone
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseSAMLProvider) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseSAMLProvider{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.IDPCertificate != nil {
		o.IDPCertificate = s.IDPCertificate
	}
	if s.IDPIssuer != nil {
		o.IDPIssuer = s.IDPIssuer
	}
	if s.IDPURL != nil {
		o.IDPURL = s.IDPURL
	}
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.Default != nil {
		o.Default = s.Default
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.Subjects != nil {
		o.Subjects = s.Subjects
	}
	if s.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	}
	if s.UpdateTime != nil {
		o.UpdateTime = s.UpdateTime
	}
	if s.ZHash != nil {
		o.ZHash = s.ZHash
	}
	if s.Zone != nil {
		o.Zone = s.Zone
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseSAMLProvider) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseSAMLProvider) ToPlain() elemental.PlainIdentifiable {

	out := NewSAMLProvider()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.IDPCertificate != nil {
		out.IDPCertificate = *o.IDPCertificate
	}
	if o.IDPIssuer != nil {
		out.IDPIssuer = *o.IDPIssuer
	}
	if o.IDPMetadata != nil {
		out.IDPMetadata = *o.IDPMetadata
	}
	if o.IDPURL != nil {
		out.IDPURL = *o.IDPURL
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Default != nil {
		out.Default = *o.Default
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Subjects != nil {
		out.Subjects = *o.Subjects
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseSAMLProvider) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseSAMLProvider) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseSAMLProvider) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseSAMLProvider) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseSAMLProvider) GetCreateIdempotencyKey() string {

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseSAMLProvider) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseSAMLProvider) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseSAMLProvider) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseSAMLProvider) GetMigrationsLog() map[string]string {

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseSAMLProvider) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseSAMLProvider) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseSAMLProvider) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseSAMLProvider) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseSAMLProvider) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseSAMLProvider) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseSAMLProvider) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseSAMLProvider) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseSAMLProvider) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseSAMLProvider) GetUpdateIdempotencyKey() string {

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseSAMLProvider) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseSAMLProvider) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseSAMLProvider) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseSAMLProvider) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseSAMLProvider) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseSAMLProvider) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseSAMLProvider) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseSAMLProvider.
func (o *SparseSAMLProvider) DeepCopy() *SparseSAMLProvider {

	if o == nil {
		return nil
	}

	out := &SparseSAMLProvider{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseSAMLProvider.
func (o *SparseSAMLProvider) DeepCopyInto(out *SparseSAMLProvider) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseSAMLProvider: %s", err))
	}

	*out = *target.(*SparseSAMLProvider)
}

type mongoAttributesSAMLProvider struct {
	ID                   bson.ObjectId       `bson:"_id"`
	IDPCertificate       string              `bson:"idpcertificate"`
	IDPIssuer            string              `bson:"idpissuer"`
	IDPURL               string              `bson:"idpurl"`
	Annotations          map[string][]string `bson:"annotations"`
	AssociatedTags       []string            `bson:"associatedtags"`
	CreateIdempotencyKey string              `bson:"createidempotencykey"`
	CreateTime           time.Time           `bson:"createtime"`
	Default              bool                `bson:"default"`
	MigrationsLog        map[string]string   `bson:"migrationslog"`
	Name                 string              `bson:"name"`
	Namespace            string              `bson:"namespace"`
	NormalizedTags       []string            `bson:"normalizedtags"`
	Protected            bool                `bson:"protected"`
	Subjects             []string            `bson:"subjects"`
	UpdateIdempotencyKey string              `bson:"updateidempotencykey"`
	UpdateTime           time.Time           `bson:"updatetime"`
	ZHash                int                 `bson:"zhash"`
	Zone                 int                 `bson:"zone"`
}
type mongoAttributesSparseSAMLProvider struct {
	ID                   bson.ObjectId        `bson:"_id"`
	IDPCertificate       *string              `bson:"idpcertificate,omitempty"`
	IDPIssuer            *string              `bson:"idpissuer,omitempty"`
	IDPURL               *string              `bson:"idpurl,omitempty"`
	Annotations          *map[string][]string `bson:"annotations,omitempty"`
	AssociatedTags       *[]string            `bson:"associatedtags,omitempty"`
	CreateIdempotencyKey *string              `bson:"createidempotencykey,omitempty"`
	CreateTime           *time.Time           `bson:"createtime,omitempty"`
	Default              *bool                `bson:"default,omitempty"`
	MigrationsLog        *map[string]string   `bson:"migrationslog,omitempty"`
	Name                 *string              `bson:"name,omitempty"`
	Namespace            *string              `bson:"namespace,omitempty"`
	NormalizedTags       *[]string            `bson:"normalizedtags,omitempty"`
	Protected            *bool                `bson:"protected,omitempty"`
	Subjects             *[]string            `bson:"subjects,omitempty"`
	UpdateIdempotencyKey *string              `bson:"updateidempotencykey,omitempty"`
	UpdateTime           *time.Time           `bson:"updatetime,omitempty"`
	ZHash                *int                 `bson:"zhash,omitempty"`
	Zone                 *int                 `bson:"zone,omitempty"`
}