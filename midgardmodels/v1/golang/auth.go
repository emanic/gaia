package midgardmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "github.com/aporeto-inc/midgard-lib/claims"

// AuthIdentity represents the Identity of the object
var AuthIdentity = elemental.Identity{
	Name:     "auth",
	Category: "auth",
}

// AuthsList represents a list of Auths
type AuthsList []*Auth

// ContentIdentity returns the identity of the objects in the list.
func (o AuthsList) ContentIdentity() elemental.Identity {
	return AuthIdentity
}

// List convert the object to and elemental.IdentifiablesList.
func (o AuthsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// Auth represents the model of a auth
type Auth struct {
	// Claims are the claims.
	Claims *claims.MidgardClaims `json:"claims" bson:"-"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`
}

// NewAuth returns a new *Auth
func NewAuth() *Auth {

	return &Auth{
		ModelVersion: 1.0,
		Claims:       claims.NewMidgardClaims(),
	}
}

// Identity returns the Identity of the object.
func (o *Auth) Identity() elemental.Identity {

	return AuthIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Auth) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Auth) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *Auth) Version() float64 {

	return 1.0
}

func (o *Auth) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Auth) Validate() error {

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
func (Auth) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return AuthAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (Auth) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AuthAttributesMap
}

// AuthAttributesMap represents the map of attribute for Auth.
var AuthAttributesMap = map[string]elemental.AttributeSpecification{
	"Claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Claims are the claims.`,
		Exposed:        true,
		Name:           "claims",
		ReadOnly:       true,
		SubType:        "claims",
		Type:           "external",
	},
}