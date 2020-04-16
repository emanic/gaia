package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CounterReportIdentity represents the Identity of the object.
var CounterReportIdentity = elemental.Identity{
	Name:     "counterreport",
	Category: "counterreports",
	Package:  "zack",
	Private:  false,
}

// CounterReportsList represents a list of CounterReports
type CounterReportsList []*CounterReport

// Identity returns the identity of the objects in the list.
func (o CounterReportsList) Identity() elemental.Identity {

	return CounterReportIdentity
}

// Copy returns a pointer to a copy the CounterReportsList.
func (o CounterReportsList) Copy() elemental.Identifiables {

	copy := append(CounterReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CounterReportsList.
func (o CounterReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CounterReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CounterReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CounterReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CounterReportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CounterReportsList converted to SparseCounterReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CounterReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCounterReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCounterReport)
	}

	return out
}

// Version returns the version of the content.
func (o CounterReportsList) Version() int {

	return 1
}

// CounterReport represents the model of a counterreport
type CounterReport struct {
	// Counter for sending FIN ACK received in unknown connection state.
	AckInUnknownState int `json:"AckInUnknownState" msgpack:"AckInUnknownState" bson:"-" mapstructure:"AckInUnknownState,omitempty"`

	// Counter for ACK packet dropped because of invalid format.
	AckInvalidFormat int `json:"AckInvalidFormat" msgpack:"AckInvalidFormat" bson:"-" mapstructure:"AckInvalidFormat,omitempty"`

	// Counter for ACK packets rejected as per policy.
	AckRejected int `json:"AckRejected" msgpack:"AckRejected" bson:"-" mapstructure:"AckRejected,omitempty"`

	// Counter for ACK packet dropped because signature validation failed.
	AckSigValidationFailed int `json:"AckSigValidationFailed" msgpack:"AckSigValidationFailed" bson:"-" mapstructure:"AckSigValidationFailed,omitempty"`

	// Counter for TCP authentication option not found.
	AckTCPNoTCPAuthOption int `json:"AckTCPNoTCPAuthOption" msgpack:"AckTCPNoTCPAuthOption" bson:"-" mapstructure:"AckTCPNoTCPAuthOption,omitempty"`

	// Counter for connections processed.
	ConnectionsProcessed int `json:"ConnectionsProcessed" msgpack:"ConnectionsProcessed" bson:"-" mapstructure:"ConnectionsProcessed,omitempty"`

	// Counter for unable to find ContextID.
	ContextIDNotFound int `json:"ContextIDNotFound" msgpack:"ContextIDNotFound" bson:"-" mapstructure:"ContextIDNotFound,omitempty"`

	// Counter for no ACLs found for external services. Dropping application SYN
	// packet.
	DroppedExternalService int `json:"DroppedExternalService" msgpack:"DroppedExternalService" bson:"-" mapstructure:"DroppedExternalService,omitempty"`

	// Counter for invalid connection state.
	InvalidConnState int `json:"InvalidConnState" msgpack:"InvalidConnState" bson:"-" mapstructure:"InvalidConnState,omitempty"`

	// Counter for invalid net state.
	InvalidNetState int `json:"InvalidNetState" msgpack:"InvalidNetState" bson:"-" mapstructure:"InvalidNetState,omitempty"`

	// Counter for invalid protocol.
	InvalidProtocol int `json:"InvalidProtocol" msgpack:"InvalidProtocol" bson:"-" mapstructure:"InvalidProtocol,omitempty"`

	// Counter for processing unit is already dead - drop SYN ACK packet.
	InvalidSynAck int `json:"InvalidSynAck" msgpack:"InvalidSynAck" bson:"-" mapstructure:"InvalidSynAck,omitempty"`

	// Counter for processing unit mark not found.
	MarkNotFound int `json:"MarkNotFound" msgpack:"MarkNotFound" bson:"-" mapstructure:"MarkNotFound,omitempty"`

	// Counter for network SYN packet was not seen.
	NetSynNotSeen int `json:"NetSynNotSeen" msgpack:"NetSynNotSeen" bson:"-" mapstructure:"NetSynNotSeen,omitempty"`

	// Counter for no context or connection found.
	NoConnFound int `json:"NoConnFound" msgpack:"NoConnFound" bson:"-" mapstructure:"NoConnFound,omitempty"`

	// Counter for traffic that belongs to a non-processing unit process.
	NonPUTraffic int `json:"NonPUTraffic" msgpack:"NonPUTraffic" bson:"-" mapstructure:"NonPUTraffic,omitempty"`

	// Counter for SYN ACK for flow with processed FIN ACK.
	OutOfOrderSynAck int `json:"OutOfOrderSynAck" msgpack:"OutOfOrderSynAck" bson:"-" mapstructure:"OutOfOrderSynAck,omitempty"`

	// Counter for port not found.
	PortNotFound int `json:"PortNotFound" msgpack:"PortNotFound" bson:"-" mapstructure:"PortNotFound,omitempty"`

	// Counter for reject the packet as per policy.
	RejectPacket int `json:"RejectPacket" msgpack:"RejectPacket" bson:"-" mapstructure:"RejectPacket,omitempty"`

	// Counter for post service processing failed for network packet.
	ServicePostprocessorFailed int `json:"ServicePostprocessorFailed" msgpack:"ServicePostprocessorFailed" bson:"-" mapstructure:"ServicePostprocessorFailed,omitempty"`

	// Counter for network packets that failed preprocessing.
	ServicePreprocessorFailed int `json:"ServicePreprocessorFailed" msgpack:"ServicePreprocessorFailed" bson:"-" mapstructure:"ServicePreprocessorFailed,omitempty"`

	// Counter for SYN ACK packet dropped because of bad claims.
	SynAckBadClaims int `json:"SynAckBadClaims" msgpack:"SynAckBadClaims" bson:"-" mapstructure:"SynAckBadClaims,omitempty"`

	// Counter for SYN ACK packet dropped because of encryption mismatch.
	SynAckClaimsMisMatch int `json:"SynAckClaimsMisMatch" msgpack:"SynAckClaimsMisMatch" bson:"-" mapstructure:"SynAckClaimsMisMatch,omitempty"`

	// Counter for SYN ACK from external service dropped.
	SynAckDroppedExternalService int `json:"SynAckDroppedExternalService" msgpack:"SynAckDroppedExternalService" bson:"-" mapstructure:"SynAckDroppedExternalService,omitempty"`

	// Counter for SYN ACK packet dropped because of invalid format.
	SynAckInvalidFormat int `json:"SynAckInvalidFormat" msgpack:"SynAckInvalidFormat" bson:"-" mapstructure:"SynAckInvalidFormat,omitempty"`

	// Counter for SYN ACK packet dropped because of no claims.
	SynAckMissingClaims int `json:"SynAckMissingClaims" msgpack:"SynAckMissingClaims" bson:"-" mapstructure:"SynAckMissingClaims,omitempty"`

	// Counter for SYN ACK packet dropped because of missing token.
	SynAckMissingToken int `json:"SynAckMissingToken" msgpack:"SynAckMissingToken" bson:"-" mapstructure:"SynAckMissingToken,omitempty"`

	// Counter for TCP authentication option not found.
	SynAckNoTCPAuthOption int `json:"SynAckNoTCPAuthOption" msgpack:"SynAckNoTCPAuthOption" bson:"-" mapstructure:"SynAckNoTCPAuthOption,omitempty"`

	// Counter for dropping because of reject rule on transmitter.
	SynAckRejected int `json:"SynAckRejected" msgpack:"SynAckRejected" bson:"-" mapstructure:"SynAckRejected,omitempty"`

	// Counter for SYN packet dropped because of invalid format.
	SynDroppedInvalidFormat int `json:"SynDroppedInvalidFormat" msgpack:"SynDroppedInvalidFormat" bson:"-" mapstructure:"SynDroppedInvalidFormat,omitempty"`

	// Counter for SYN packet dropped because of invalid token.
	SynDroppedInvalidToken int `json:"SynDroppedInvalidToken" msgpack:"SynDroppedInvalidToken" bson:"-" mapstructure:"SynDroppedInvalidToken,omitempty"`

	// Counter for SYN packet dropped because of no claims.
	SynDroppedNoClaims int `json:"SynDroppedNoClaims" msgpack:"SynDroppedNoClaims" bson:"-" mapstructure:"SynDroppedNoClaims,omitempty"`

	// Counter for TCP authentication option not found.
	SynDroppedTCPOption int `json:"SynDroppedTCPOption" msgpack:"SynDroppedTCPOption" bson:"-" mapstructure:"SynDroppedTCPOption,omitempty"`

	// Counter for SYN packet dropped due to policy.
	SynRejectPacket int `json:"SynRejectPacket" msgpack:"SynRejectPacket" bson:"-" mapstructure:"SynRejectPacket,omitempty"`

	// Counter for received SYN packet from unknown processing unit.
	SynUnexpectedPacket int `json:"SynUnexpectedPacket" msgpack:"SynUnexpectedPacket" bson:"-" mapstructure:"SynUnexpectedPacket,omitempty"`

	// Counter for TCP authentication option not found.
	TCPAuthNotFound int `json:"TCPAuthNotFound" msgpack:"TCPAuthNotFound" bson:"-" mapstructure:"TCPAuthNotFound,omitempty"`

	// Counter for UDP ACK packet dropped due to an invalid signature.
	UDPAckInvalidSignature int `json:"UDPAckInvalidSignature" msgpack:"UDPAckInvalidSignature" bson:"-" mapstructure:"UDPAckInvalidSignature,omitempty"`

	// Counter for number of processed UDP connections.
	UDPConnectionsProcessed int `json:"UDPConnectionsProcessed" msgpack:"UDPConnectionsProcessed" bson:"-" mapstructure:"UDPConnectionsProcessed,omitempty"`

	// Counter for dropped UDP data packets with no context.
	UDPDropContextNotFound int `json:"UDPDropContextNotFound" msgpack:"UDPDropContextNotFound" bson:"-" mapstructure:"UDPDropContextNotFound,omitempty"`

	// Counter for dropped UDP FIN handshake packets.
	UDPDropFin int `json:"UDPDropFin" msgpack:"UDPDropFin" bson:"-" mapstructure:"UDPDropFin,omitempty"`

	// Counter for dropped UDP in NfQueue.
	UDPDropInNfQueue int `json:"UDPDropInNfQueue" msgpack:"UDPDropInNfQueue" bson:"-" mapstructure:"UDPDropInNfQueue,omitempty"`

	// Counter for dropped UDP data packets with no connection.
	UDPDropNoConnection int `json:"UDPDropNoConnection" msgpack:"UDPDropNoConnection" bson:"-" mapstructure:"UDPDropNoConnection,omitempty"`

	// Counter for dropped UDP data packets.
	UDPDropPacket int `json:"UDPDropPacket" msgpack:"UDPDropPacket" bson:"-" mapstructure:"UDPDropPacket,omitempty"`

	// Counter for dropped UDP Queue Full.
	UDPDropQueueFull int `json:"UDPDropQueueFull" msgpack:"UDPDropQueueFull" bson:"-" mapstructure:"UDPDropQueueFull,omitempty"`

	// Counter for dropped UDP SYN ACK handshake packets.
	UDPDropSynAck int `json:"UDPDropSynAck" msgpack:"UDPDropSynAck" bson:"-" mapstructure:"UDPDropSynAck,omitempty"`

	// Counter for UDP packets received in invalid network state.
	UDPInvalidNetState int `json:"UDPInvalidNetState" msgpack:"UDPInvalidNetState" bson:"-" mapstructure:"UDPInvalidNetState,omitempty"`

	// Counter for UDP packets failing postprocessing.
	UDPPostProcessingFailed int `json:"UDPPostProcessingFailed" msgpack:"UDPPostProcessingFailed" bson:"-" mapstructure:"UDPPostProcessingFailed,omitempty"`

	// Counter for UDP packets failing preprocessing.
	UDPPreProcessingFailed int `json:"UDPPreProcessingFailed" msgpack:"UDPPreProcessingFailed" bson:"-" mapstructure:"UDPPreProcessingFailed,omitempty"`

	// Counter for UDP packets dropped due to policy.
	UDPRejected int `json:"UDPRejected" msgpack:"UDPRejected" bson:"-" mapstructure:"UDPRejected,omitempty"`

	// Counter for UDP SYN ACK packets dropped due to bad claims.
	UDPSynAckDropBadClaims int `json:"UDPSynAckDropBadClaims" msgpack:"UDPSynAckDropBadClaims" bson:"-" mapstructure:"UDPSynAckDropBadClaims,omitempty"`

	// Counter for UDP SYN ACK packets dropped due to missing claims.
	UDPSynAckMissingClaims int `json:"UDPSynAckMissingClaims" msgpack:"UDPSynAckMissingClaims" bson:"-" mapstructure:"UDPSynAckMissingClaims,omitempty"`

	// Counter for UDP SYN ACK packets dropped due to bad claims.
	UDPSynAckPolicy int `json:"UDPSynAckPolicy" msgpack:"UDPSynAckPolicy" bson:"-" mapstructure:"UDPSynAckPolicy,omitempty"`

	// Counter for dropped UDP SYN transmits.
	UDPSynDrop int `json:"UDPSynDrop" msgpack:"UDPSynDrop" bson:"-" mapstructure:"UDPSynDrop,omitempty"`

	// Counter for dropped UDP SYN policy.
	UDPSynDropPolicy int `json:"UDPSynDropPolicy" msgpack:"UDPSynDropPolicy" bson:"-" mapstructure:"UDPSynDropPolicy,omitempty"`

	// Counter for dropped UDP FIN handshake packets.
	UDPSynInvalidToken int `json:"UDPSynInvalidToken" msgpack:"UDPSynInvalidToken" bson:"-" mapstructure:"UDPSynInvalidToken,omitempty"`

	// Counter for UDP SYN packet dropped due to missing claims.
	UDPSynMissingClaims int `json:"UDPSynMissingClaims" msgpack:"UDPSynMissingClaims" bson:"-" mapstructure:"UDPSynMissingClaims,omitempty"`

	// Counter for unknown error.
	UnknownError int `json:"UnknownError" msgpack:"UnknownError" bson:"-" mapstructure:"UnknownError,omitempty"`

	// Non-zero counter indicates analyzed connections for unencrypted, encrypted, and
	// packets from endpoint applications with the TCP Fast Open option set. These are
	// not dropped counter.
	ConnectionsAnalyzed int `json:"connectionsAnalyzed" msgpack:"connectionsAnalyzed" bson:"-" mapstructure:"connectionsAnalyzed,omitempty"`

	// Non-zero counter indicates dropped connections because of invalid state,
	// non-processing unit traffic, or out of order packets.
	ConnectionsDropped int `json:"connectionsDropped" msgpack:"connectionsDropped" bson:"-" mapstructure:"connectionsDropped,omitempty"`

	// Non-zero counter indicates expired connections because of response not being
	// received within a certain amount of time after the request is made.
	ConnectionsExpired int `json:"connectionsExpired" msgpack:"connectionsExpired" bson:"-" mapstructure:"connectionsExpired,omitempty"`

	// Non-zero counter indicates dropped packets that did not hit any of our iptables
	// rules and queue drops.
	DroppedPackets int `json:"droppedPackets" msgpack:"droppedPackets" bson:"-" mapstructure:"droppedPackets,omitempty"`

	// Non-zero counter indicates encryption processing failures of data packets.
	EncryptionFailures int `json:"encryptionFailures" msgpack:"encryptionFailures" bson:"-" mapstructure:"encryptionFailures,omitempty"`

	// Identifier of the enforcer sending the report.
	EnforcerID string `json:"enforcerID" msgpack:"enforcerID" bson:"enforcerid" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer sending the report.
	EnforcerNamespace string `json:"enforcerNamespace" msgpack:"enforcerNamespace" bson:"enforcernamespace" mapstructure:"enforcerNamespace,omitempty"`

	// Non-zero counter indicates connections going to and from external networks.
	// These may be drops or allowed counters.
	ExternalNetworkConnections int `json:"externalNetworkConnections" msgpack:"externalNetworkConnections" bson:"-" mapstructure:"externalNetworkConnections,omitempty"`

	// Non-zero counter indicates packets dropped due to a reject policy.
	PolicyDrops int `json:"policyDrops" msgpack:"policyDrops" bson:"-" mapstructure:"policyDrops,omitempty"`

	// PUID is the ID of the processing unit reporting the counter.
	ProcessingUnitID string `json:"processingUnitID" msgpack:"processingUnitID" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the processing unit reporting the counter.
	ProcessingUnitNamespace string `json:"processingUnitNamespace" msgpack:"processingUnitNamespace" bson:"-" mapstructure:"processingUnitNamespace,omitempty"`

	// Timestamp is the date of the report.
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	// Non-zero counter indicates packets rejected due to anything related to token
	// creation/parsing failures.
	TokenDrops int `json:"tokenDrops" msgpack:"tokenDrops" bson:"-" mapstructure:"tokenDrops,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCounterReport returns a new *CounterReport
func NewCounterReport() *CounterReport {

	return &CounterReport{
		ModelVersion:                 1,
		TCPAuthNotFound:              0,
		InvalidConnState:             0,
		UDPAckInvalidSignature:       0,
		InvalidProtocol:              0,
		UDPConnectionsProcessed:      0,
		ConnectionsAnalyzed:          0,
		ConnectionsExpired:           0,
		ConnectionsDropped:           0,
		MarkNotFound:                 0,
		AckInUnknownState:            0,
		ExternalNetworkConnections:   0,
		EncryptionFailures:           0,
		DroppedPackets:               0,
		ContextIDNotFound:            0,
		DroppedExternalService:       0,
		UDPDropFin:                   0,
		TokenDrops:                   0,
		PolicyDrops:                  0,
		AckRejected:                  0,
		AckSigValidationFailed:       0,
		AckTCPNoTCPAuthOption:        0,
		ConnectionsProcessed:         0,
		NoConnFound:                  0,
		NonPUTraffic:                 0,
		AckInvalidFormat:             0,
		UDPDropSynAck:                0,
		SynDroppedInvalidFormat:      0,
		ServicePostprocessorFailed:   0,
		ServicePreprocessorFailed:    0,
		SynAckBadClaims:              0,
		SynAckClaimsMisMatch:         0,
		SynAckDroppedExternalService: 0,
		RejectPacket:                 0,
		SynAckMissingClaims:          0,
		SynAckMissingToken:           0,
		SynAckNoTCPAuthOption:        0,
		SynAckRejected:               0,
		UDPSynAckMissingClaims:       0,
		UDPDropNoConnection:          0,
		SynDroppedNoClaims:           0,
		SynDroppedTCPOption:          0,
		SynRejectPacket:              0,
		SynUnexpectedPacket:          0,
		InvalidNetState:              0,
		InvalidSynAck:                0,
		NetSynNotSeen:                0,
		UDPDropContextNotFound:       0,
		OutOfOrderSynAck:             0,
		UDPDropInNfQueue:             0,
		UDPDropPacket:                0,
		SynDroppedInvalidToken:       0,
		UDPSynInvalidToken:           0,
		UDPSynAckDropBadClaims:       0,
		UDPSynDropPolicy:             0,
		UDPSynDrop:                   0,
		UDPSynAckPolicy:              0,
		UDPInvalidNetState:           0,
		SynAckInvalidFormat:          0,
		UDPRejected:                  0,
		UDPPreProcessingFailed:       0,
		UDPPostProcessingFailed:      0,
		UnknownError:                 0,
		UDPDropQueueFull:             0,
		PortNotFound:                 0,
		UDPSynMissingClaims:          0,
	}
}

// Identity returns the Identity of the object.
func (o *CounterReport) Identity() elemental.Identity {

	return CounterReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CounterReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CounterReport) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CounterReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCounterReport{}

	s.EnforcerID = o.EnforcerID
	s.EnforcerNamespace = o.EnforcerNamespace

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CounterReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCounterReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.EnforcerID = s.EnforcerID
	o.EnforcerNamespace = s.EnforcerNamespace

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CounterReport) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CounterReport) BleveType() string {

	return "counterreport"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CounterReport) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CounterReport) Doc() string {

	return `Post a new counter tracing report.`
}

func (o *CounterReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CounterReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCounterReport{
			AckInUnknownState:            &o.AckInUnknownState,
			AckInvalidFormat:             &o.AckInvalidFormat,
			AckRejected:                  &o.AckRejected,
			AckSigValidationFailed:       &o.AckSigValidationFailed,
			AckTCPNoTCPAuthOption:        &o.AckTCPNoTCPAuthOption,
			ConnectionsProcessed:         &o.ConnectionsProcessed,
			ContextIDNotFound:            &o.ContextIDNotFound,
			DroppedExternalService:       &o.DroppedExternalService,
			InvalidConnState:             &o.InvalidConnState,
			InvalidNetState:              &o.InvalidNetState,
			InvalidProtocol:              &o.InvalidProtocol,
			InvalidSynAck:                &o.InvalidSynAck,
			MarkNotFound:                 &o.MarkNotFound,
			NetSynNotSeen:                &o.NetSynNotSeen,
			NoConnFound:                  &o.NoConnFound,
			NonPUTraffic:                 &o.NonPUTraffic,
			OutOfOrderSynAck:             &o.OutOfOrderSynAck,
			PortNotFound:                 &o.PortNotFound,
			RejectPacket:                 &o.RejectPacket,
			ServicePostprocessorFailed:   &o.ServicePostprocessorFailed,
			ServicePreprocessorFailed:    &o.ServicePreprocessorFailed,
			SynAckBadClaims:              &o.SynAckBadClaims,
			SynAckClaimsMisMatch:         &o.SynAckClaimsMisMatch,
			SynAckDroppedExternalService: &o.SynAckDroppedExternalService,
			SynAckInvalidFormat:          &o.SynAckInvalidFormat,
			SynAckMissingClaims:          &o.SynAckMissingClaims,
			SynAckMissingToken:           &o.SynAckMissingToken,
			SynAckNoTCPAuthOption:        &o.SynAckNoTCPAuthOption,
			SynAckRejected:               &o.SynAckRejected,
			SynDroppedInvalidFormat:      &o.SynDroppedInvalidFormat,
			SynDroppedInvalidToken:       &o.SynDroppedInvalidToken,
			SynDroppedNoClaims:           &o.SynDroppedNoClaims,
			SynDroppedTCPOption:          &o.SynDroppedTCPOption,
			SynRejectPacket:              &o.SynRejectPacket,
			SynUnexpectedPacket:          &o.SynUnexpectedPacket,
			TCPAuthNotFound:              &o.TCPAuthNotFound,
			UDPAckInvalidSignature:       &o.UDPAckInvalidSignature,
			UDPConnectionsProcessed:      &o.UDPConnectionsProcessed,
			UDPDropContextNotFound:       &o.UDPDropContextNotFound,
			UDPDropFin:                   &o.UDPDropFin,
			UDPDropInNfQueue:             &o.UDPDropInNfQueue,
			UDPDropNoConnection:          &o.UDPDropNoConnection,
			UDPDropPacket:                &o.UDPDropPacket,
			UDPDropQueueFull:             &o.UDPDropQueueFull,
			UDPDropSynAck:                &o.UDPDropSynAck,
			UDPInvalidNetState:           &o.UDPInvalidNetState,
			UDPPostProcessingFailed:      &o.UDPPostProcessingFailed,
			UDPPreProcessingFailed:       &o.UDPPreProcessingFailed,
			UDPRejected:                  &o.UDPRejected,
			UDPSynAckDropBadClaims:       &o.UDPSynAckDropBadClaims,
			UDPSynAckMissingClaims:       &o.UDPSynAckMissingClaims,
			UDPSynAckPolicy:              &o.UDPSynAckPolicy,
			UDPSynDrop:                   &o.UDPSynDrop,
			UDPSynDropPolicy:             &o.UDPSynDropPolicy,
			UDPSynInvalidToken:           &o.UDPSynInvalidToken,
			UDPSynMissingClaims:          &o.UDPSynMissingClaims,
			UnknownError:                 &o.UnknownError,
			ConnectionsAnalyzed:          &o.ConnectionsAnalyzed,
			ConnectionsDropped:           &o.ConnectionsDropped,
			ConnectionsExpired:           &o.ConnectionsExpired,
			DroppedPackets:               &o.DroppedPackets,
			EncryptionFailures:           &o.EncryptionFailures,
			EnforcerID:                   &o.EnforcerID,
			EnforcerNamespace:            &o.EnforcerNamespace,
			ExternalNetworkConnections:   &o.ExternalNetworkConnections,
			PolicyDrops:                  &o.PolicyDrops,
			ProcessingUnitID:             &o.ProcessingUnitID,
			ProcessingUnitNamespace:      &o.ProcessingUnitNamespace,
			Timestamp:                    &o.Timestamp,
			TokenDrops:                   &o.TokenDrops,
		}
	}

	sp := &SparseCounterReport{}
	for _, f := range fields {
		switch f {
		case "AckInUnknownState":
			sp.AckInUnknownState = &(o.AckInUnknownState)
		case "AckInvalidFormat":
			sp.AckInvalidFormat = &(o.AckInvalidFormat)
		case "AckRejected":
			sp.AckRejected = &(o.AckRejected)
		case "AckSigValidationFailed":
			sp.AckSigValidationFailed = &(o.AckSigValidationFailed)
		case "AckTCPNoTCPAuthOption":
			sp.AckTCPNoTCPAuthOption = &(o.AckTCPNoTCPAuthOption)
		case "ConnectionsProcessed":
			sp.ConnectionsProcessed = &(o.ConnectionsProcessed)
		case "ContextIDNotFound":
			sp.ContextIDNotFound = &(o.ContextIDNotFound)
		case "DroppedExternalService":
			sp.DroppedExternalService = &(o.DroppedExternalService)
		case "InvalidConnState":
			sp.InvalidConnState = &(o.InvalidConnState)
		case "InvalidNetState":
			sp.InvalidNetState = &(o.InvalidNetState)
		case "InvalidProtocol":
			sp.InvalidProtocol = &(o.InvalidProtocol)
		case "InvalidSynAck":
			sp.InvalidSynAck = &(o.InvalidSynAck)
		case "MarkNotFound":
			sp.MarkNotFound = &(o.MarkNotFound)
		case "NetSynNotSeen":
			sp.NetSynNotSeen = &(o.NetSynNotSeen)
		case "NoConnFound":
			sp.NoConnFound = &(o.NoConnFound)
		case "NonPUTraffic":
			sp.NonPUTraffic = &(o.NonPUTraffic)
		case "OutOfOrderSynAck":
			sp.OutOfOrderSynAck = &(o.OutOfOrderSynAck)
		case "PortNotFound":
			sp.PortNotFound = &(o.PortNotFound)
		case "RejectPacket":
			sp.RejectPacket = &(o.RejectPacket)
		case "ServicePostprocessorFailed":
			sp.ServicePostprocessorFailed = &(o.ServicePostprocessorFailed)
		case "ServicePreprocessorFailed":
			sp.ServicePreprocessorFailed = &(o.ServicePreprocessorFailed)
		case "SynAckBadClaims":
			sp.SynAckBadClaims = &(o.SynAckBadClaims)
		case "SynAckClaimsMisMatch":
			sp.SynAckClaimsMisMatch = &(o.SynAckClaimsMisMatch)
		case "SynAckDroppedExternalService":
			sp.SynAckDroppedExternalService = &(o.SynAckDroppedExternalService)
		case "SynAckInvalidFormat":
			sp.SynAckInvalidFormat = &(o.SynAckInvalidFormat)
		case "SynAckMissingClaims":
			sp.SynAckMissingClaims = &(o.SynAckMissingClaims)
		case "SynAckMissingToken":
			sp.SynAckMissingToken = &(o.SynAckMissingToken)
		case "SynAckNoTCPAuthOption":
			sp.SynAckNoTCPAuthOption = &(o.SynAckNoTCPAuthOption)
		case "SynAckRejected":
			sp.SynAckRejected = &(o.SynAckRejected)
		case "SynDroppedInvalidFormat":
			sp.SynDroppedInvalidFormat = &(o.SynDroppedInvalidFormat)
		case "SynDroppedInvalidToken":
			sp.SynDroppedInvalidToken = &(o.SynDroppedInvalidToken)
		case "SynDroppedNoClaims":
			sp.SynDroppedNoClaims = &(o.SynDroppedNoClaims)
		case "SynDroppedTCPOption":
			sp.SynDroppedTCPOption = &(o.SynDroppedTCPOption)
		case "SynRejectPacket":
			sp.SynRejectPacket = &(o.SynRejectPacket)
		case "SynUnexpectedPacket":
			sp.SynUnexpectedPacket = &(o.SynUnexpectedPacket)
		case "TCPAuthNotFound":
			sp.TCPAuthNotFound = &(o.TCPAuthNotFound)
		case "UDPAckInvalidSignature":
			sp.UDPAckInvalidSignature = &(o.UDPAckInvalidSignature)
		case "UDPConnectionsProcessed":
			sp.UDPConnectionsProcessed = &(o.UDPConnectionsProcessed)
		case "UDPDropContextNotFound":
			sp.UDPDropContextNotFound = &(o.UDPDropContextNotFound)
		case "UDPDropFin":
			sp.UDPDropFin = &(o.UDPDropFin)
		case "UDPDropInNfQueue":
			sp.UDPDropInNfQueue = &(o.UDPDropInNfQueue)
		case "UDPDropNoConnection":
			sp.UDPDropNoConnection = &(o.UDPDropNoConnection)
		case "UDPDropPacket":
			sp.UDPDropPacket = &(o.UDPDropPacket)
		case "UDPDropQueueFull":
			sp.UDPDropQueueFull = &(o.UDPDropQueueFull)
		case "UDPDropSynAck":
			sp.UDPDropSynAck = &(o.UDPDropSynAck)
		case "UDPInvalidNetState":
			sp.UDPInvalidNetState = &(o.UDPInvalidNetState)
		case "UDPPostProcessingFailed":
			sp.UDPPostProcessingFailed = &(o.UDPPostProcessingFailed)
		case "UDPPreProcessingFailed":
			sp.UDPPreProcessingFailed = &(o.UDPPreProcessingFailed)
		case "UDPRejected":
			sp.UDPRejected = &(o.UDPRejected)
		case "UDPSynAckDropBadClaims":
			sp.UDPSynAckDropBadClaims = &(o.UDPSynAckDropBadClaims)
		case "UDPSynAckMissingClaims":
			sp.UDPSynAckMissingClaims = &(o.UDPSynAckMissingClaims)
		case "UDPSynAckPolicy":
			sp.UDPSynAckPolicy = &(o.UDPSynAckPolicy)
		case "UDPSynDrop":
			sp.UDPSynDrop = &(o.UDPSynDrop)
		case "UDPSynDropPolicy":
			sp.UDPSynDropPolicy = &(o.UDPSynDropPolicy)
		case "UDPSynInvalidToken":
			sp.UDPSynInvalidToken = &(o.UDPSynInvalidToken)
		case "UDPSynMissingClaims":
			sp.UDPSynMissingClaims = &(o.UDPSynMissingClaims)
		case "UnknownError":
			sp.UnknownError = &(o.UnknownError)
		case "connectionsAnalyzed":
			sp.ConnectionsAnalyzed = &(o.ConnectionsAnalyzed)
		case "connectionsDropped":
			sp.ConnectionsDropped = &(o.ConnectionsDropped)
		case "connectionsExpired":
			sp.ConnectionsExpired = &(o.ConnectionsExpired)
		case "droppedPackets":
			sp.DroppedPackets = &(o.DroppedPackets)
		case "encryptionFailures":
			sp.EncryptionFailures = &(o.EncryptionFailures)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "enforcerNamespace":
			sp.EnforcerNamespace = &(o.EnforcerNamespace)
		case "externalNetworkConnections":
			sp.ExternalNetworkConnections = &(o.ExternalNetworkConnections)
		case "policyDrops":
			sp.PolicyDrops = &(o.PolicyDrops)
		case "processingUnitID":
			sp.ProcessingUnitID = &(o.ProcessingUnitID)
		case "processingUnitNamespace":
			sp.ProcessingUnitNamespace = &(o.ProcessingUnitNamespace)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		case "tokenDrops":
			sp.TokenDrops = &(o.TokenDrops)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCounterReport to the object.
func (o *CounterReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCounterReport)
	if so.AckInUnknownState != nil {
		o.AckInUnknownState = *so.AckInUnknownState
	}
	if so.AckInvalidFormat != nil {
		o.AckInvalidFormat = *so.AckInvalidFormat
	}
	if so.AckRejected != nil {
		o.AckRejected = *so.AckRejected
	}
	if so.AckSigValidationFailed != nil {
		o.AckSigValidationFailed = *so.AckSigValidationFailed
	}
	if so.AckTCPNoTCPAuthOption != nil {
		o.AckTCPNoTCPAuthOption = *so.AckTCPNoTCPAuthOption
	}
	if so.ConnectionsProcessed != nil {
		o.ConnectionsProcessed = *so.ConnectionsProcessed
	}
	if so.ContextIDNotFound != nil {
		o.ContextIDNotFound = *so.ContextIDNotFound
	}
	if so.DroppedExternalService != nil {
		o.DroppedExternalService = *so.DroppedExternalService
	}
	if so.InvalidConnState != nil {
		o.InvalidConnState = *so.InvalidConnState
	}
	if so.InvalidNetState != nil {
		o.InvalidNetState = *so.InvalidNetState
	}
	if so.InvalidProtocol != nil {
		o.InvalidProtocol = *so.InvalidProtocol
	}
	if so.InvalidSynAck != nil {
		o.InvalidSynAck = *so.InvalidSynAck
	}
	if so.MarkNotFound != nil {
		o.MarkNotFound = *so.MarkNotFound
	}
	if so.NetSynNotSeen != nil {
		o.NetSynNotSeen = *so.NetSynNotSeen
	}
	if so.NoConnFound != nil {
		o.NoConnFound = *so.NoConnFound
	}
	if so.NonPUTraffic != nil {
		o.NonPUTraffic = *so.NonPUTraffic
	}
	if so.OutOfOrderSynAck != nil {
		o.OutOfOrderSynAck = *so.OutOfOrderSynAck
	}
	if so.PortNotFound != nil {
		o.PortNotFound = *so.PortNotFound
	}
	if so.RejectPacket != nil {
		o.RejectPacket = *so.RejectPacket
	}
	if so.ServicePostprocessorFailed != nil {
		o.ServicePostprocessorFailed = *so.ServicePostprocessorFailed
	}
	if so.ServicePreprocessorFailed != nil {
		o.ServicePreprocessorFailed = *so.ServicePreprocessorFailed
	}
	if so.SynAckBadClaims != nil {
		o.SynAckBadClaims = *so.SynAckBadClaims
	}
	if so.SynAckClaimsMisMatch != nil {
		o.SynAckClaimsMisMatch = *so.SynAckClaimsMisMatch
	}
	if so.SynAckDroppedExternalService != nil {
		o.SynAckDroppedExternalService = *so.SynAckDroppedExternalService
	}
	if so.SynAckInvalidFormat != nil {
		o.SynAckInvalidFormat = *so.SynAckInvalidFormat
	}
	if so.SynAckMissingClaims != nil {
		o.SynAckMissingClaims = *so.SynAckMissingClaims
	}
	if so.SynAckMissingToken != nil {
		o.SynAckMissingToken = *so.SynAckMissingToken
	}
	if so.SynAckNoTCPAuthOption != nil {
		o.SynAckNoTCPAuthOption = *so.SynAckNoTCPAuthOption
	}
	if so.SynAckRejected != nil {
		o.SynAckRejected = *so.SynAckRejected
	}
	if so.SynDroppedInvalidFormat != nil {
		o.SynDroppedInvalidFormat = *so.SynDroppedInvalidFormat
	}
	if so.SynDroppedInvalidToken != nil {
		o.SynDroppedInvalidToken = *so.SynDroppedInvalidToken
	}
	if so.SynDroppedNoClaims != nil {
		o.SynDroppedNoClaims = *so.SynDroppedNoClaims
	}
	if so.SynDroppedTCPOption != nil {
		o.SynDroppedTCPOption = *so.SynDroppedTCPOption
	}
	if so.SynRejectPacket != nil {
		o.SynRejectPacket = *so.SynRejectPacket
	}
	if so.SynUnexpectedPacket != nil {
		o.SynUnexpectedPacket = *so.SynUnexpectedPacket
	}
	if so.TCPAuthNotFound != nil {
		o.TCPAuthNotFound = *so.TCPAuthNotFound
	}
	if so.UDPAckInvalidSignature != nil {
		o.UDPAckInvalidSignature = *so.UDPAckInvalidSignature
	}
	if so.UDPConnectionsProcessed != nil {
		o.UDPConnectionsProcessed = *so.UDPConnectionsProcessed
	}
	if so.UDPDropContextNotFound != nil {
		o.UDPDropContextNotFound = *so.UDPDropContextNotFound
	}
	if so.UDPDropFin != nil {
		o.UDPDropFin = *so.UDPDropFin
	}
	if so.UDPDropInNfQueue != nil {
		o.UDPDropInNfQueue = *so.UDPDropInNfQueue
	}
	if so.UDPDropNoConnection != nil {
		o.UDPDropNoConnection = *so.UDPDropNoConnection
	}
	if so.UDPDropPacket != nil {
		o.UDPDropPacket = *so.UDPDropPacket
	}
	if so.UDPDropQueueFull != nil {
		o.UDPDropQueueFull = *so.UDPDropQueueFull
	}
	if so.UDPDropSynAck != nil {
		o.UDPDropSynAck = *so.UDPDropSynAck
	}
	if so.UDPInvalidNetState != nil {
		o.UDPInvalidNetState = *so.UDPInvalidNetState
	}
	if so.UDPPostProcessingFailed != nil {
		o.UDPPostProcessingFailed = *so.UDPPostProcessingFailed
	}
	if so.UDPPreProcessingFailed != nil {
		o.UDPPreProcessingFailed = *so.UDPPreProcessingFailed
	}
	if so.UDPRejected != nil {
		o.UDPRejected = *so.UDPRejected
	}
	if so.UDPSynAckDropBadClaims != nil {
		o.UDPSynAckDropBadClaims = *so.UDPSynAckDropBadClaims
	}
	if so.UDPSynAckMissingClaims != nil {
		o.UDPSynAckMissingClaims = *so.UDPSynAckMissingClaims
	}
	if so.UDPSynAckPolicy != nil {
		o.UDPSynAckPolicy = *so.UDPSynAckPolicy
	}
	if so.UDPSynDrop != nil {
		o.UDPSynDrop = *so.UDPSynDrop
	}
	if so.UDPSynDropPolicy != nil {
		o.UDPSynDropPolicy = *so.UDPSynDropPolicy
	}
	if so.UDPSynInvalidToken != nil {
		o.UDPSynInvalidToken = *so.UDPSynInvalidToken
	}
	if so.UDPSynMissingClaims != nil {
		o.UDPSynMissingClaims = *so.UDPSynMissingClaims
	}
	if so.UnknownError != nil {
		o.UnknownError = *so.UnknownError
	}
	if so.ConnectionsAnalyzed != nil {
		o.ConnectionsAnalyzed = *so.ConnectionsAnalyzed
	}
	if so.ConnectionsDropped != nil {
		o.ConnectionsDropped = *so.ConnectionsDropped
	}
	if so.ConnectionsExpired != nil {
		o.ConnectionsExpired = *so.ConnectionsExpired
	}
	if so.DroppedPackets != nil {
		o.DroppedPackets = *so.DroppedPackets
	}
	if so.EncryptionFailures != nil {
		o.EncryptionFailures = *so.EncryptionFailures
	}
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.EnforcerNamespace != nil {
		o.EnforcerNamespace = *so.EnforcerNamespace
	}
	if so.ExternalNetworkConnections != nil {
		o.ExternalNetworkConnections = *so.ExternalNetworkConnections
	}
	if so.PolicyDrops != nil {
		o.PolicyDrops = *so.PolicyDrops
	}
	if so.ProcessingUnitID != nil {
		o.ProcessingUnitID = *so.ProcessingUnitID
	}
	if so.ProcessingUnitNamespace != nil {
		o.ProcessingUnitNamespace = *so.ProcessingUnitNamespace
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
	if so.TokenDrops != nil {
		o.TokenDrops = *so.TokenDrops
	}
}

// DeepCopy returns a deep copy if the CounterReport.
func (o *CounterReport) DeepCopy() *CounterReport {

	if o == nil {
		return nil
	}

	out := &CounterReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CounterReport.
func (o *CounterReport) DeepCopyInto(out *CounterReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CounterReport: %s", err))
	}

	*out = *target.(*CounterReport)
}

// Validate valides the current information stored into the structure.
func (o *CounterReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("enforcerID", o.EnforcerID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerNamespace", o.EnforcerNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
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
func (*CounterReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CounterReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CounterReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CounterReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CounterReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CounterReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "AckInUnknownState":
		return o.AckInUnknownState
	case "AckInvalidFormat":
		return o.AckInvalidFormat
	case "AckRejected":
		return o.AckRejected
	case "AckSigValidationFailed":
		return o.AckSigValidationFailed
	case "AckTCPNoTCPAuthOption":
		return o.AckTCPNoTCPAuthOption
	case "ConnectionsProcessed":
		return o.ConnectionsProcessed
	case "ContextIDNotFound":
		return o.ContextIDNotFound
	case "DroppedExternalService":
		return o.DroppedExternalService
	case "InvalidConnState":
		return o.InvalidConnState
	case "InvalidNetState":
		return o.InvalidNetState
	case "InvalidProtocol":
		return o.InvalidProtocol
	case "InvalidSynAck":
		return o.InvalidSynAck
	case "MarkNotFound":
		return o.MarkNotFound
	case "NetSynNotSeen":
		return o.NetSynNotSeen
	case "NoConnFound":
		return o.NoConnFound
	case "NonPUTraffic":
		return o.NonPUTraffic
	case "OutOfOrderSynAck":
		return o.OutOfOrderSynAck
	case "PortNotFound":
		return o.PortNotFound
	case "RejectPacket":
		return o.RejectPacket
	case "ServicePostprocessorFailed":
		return o.ServicePostprocessorFailed
	case "ServicePreprocessorFailed":
		return o.ServicePreprocessorFailed
	case "SynAckBadClaims":
		return o.SynAckBadClaims
	case "SynAckClaimsMisMatch":
		return o.SynAckClaimsMisMatch
	case "SynAckDroppedExternalService":
		return o.SynAckDroppedExternalService
	case "SynAckInvalidFormat":
		return o.SynAckInvalidFormat
	case "SynAckMissingClaims":
		return o.SynAckMissingClaims
	case "SynAckMissingToken":
		return o.SynAckMissingToken
	case "SynAckNoTCPAuthOption":
		return o.SynAckNoTCPAuthOption
	case "SynAckRejected":
		return o.SynAckRejected
	case "SynDroppedInvalidFormat":
		return o.SynDroppedInvalidFormat
	case "SynDroppedInvalidToken":
		return o.SynDroppedInvalidToken
	case "SynDroppedNoClaims":
		return o.SynDroppedNoClaims
	case "SynDroppedTCPOption":
		return o.SynDroppedTCPOption
	case "SynRejectPacket":
		return o.SynRejectPacket
	case "SynUnexpectedPacket":
		return o.SynUnexpectedPacket
	case "TCPAuthNotFound":
		return o.TCPAuthNotFound
	case "UDPAckInvalidSignature":
		return o.UDPAckInvalidSignature
	case "UDPConnectionsProcessed":
		return o.UDPConnectionsProcessed
	case "UDPDropContextNotFound":
		return o.UDPDropContextNotFound
	case "UDPDropFin":
		return o.UDPDropFin
	case "UDPDropInNfQueue":
		return o.UDPDropInNfQueue
	case "UDPDropNoConnection":
		return o.UDPDropNoConnection
	case "UDPDropPacket":
		return o.UDPDropPacket
	case "UDPDropQueueFull":
		return o.UDPDropQueueFull
	case "UDPDropSynAck":
		return o.UDPDropSynAck
	case "UDPInvalidNetState":
		return o.UDPInvalidNetState
	case "UDPPostProcessingFailed":
		return o.UDPPostProcessingFailed
	case "UDPPreProcessingFailed":
		return o.UDPPreProcessingFailed
	case "UDPRejected":
		return o.UDPRejected
	case "UDPSynAckDropBadClaims":
		return o.UDPSynAckDropBadClaims
	case "UDPSynAckMissingClaims":
		return o.UDPSynAckMissingClaims
	case "UDPSynAckPolicy":
		return o.UDPSynAckPolicy
	case "UDPSynDrop":
		return o.UDPSynDrop
	case "UDPSynDropPolicy":
		return o.UDPSynDropPolicy
	case "UDPSynInvalidToken":
		return o.UDPSynInvalidToken
	case "UDPSynMissingClaims":
		return o.UDPSynMissingClaims
	case "UnknownError":
		return o.UnknownError
	case "connectionsAnalyzed":
		return o.ConnectionsAnalyzed
	case "connectionsDropped":
		return o.ConnectionsDropped
	case "connectionsExpired":
		return o.ConnectionsExpired
	case "droppedPackets":
		return o.DroppedPackets
	case "encryptionFailures":
		return o.EncryptionFailures
	case "enforcerID":
		return o.EnforcerID
	case "enforcerNamespace":
		return o.EnforcerNamespace
	case "externalNetworkConnections":
		return o.ExternalNetworkConnections
	case "policyDrops":
		return o.PolicyDrops
	case "processingUnitID":
		return o.ProcessingUnitID
	case "processingUnitNamespace":
		return o.ProcessingUnitNamespace
	case "timestamp":
		return o.Timestamp
	case "tokenDrops":
		return o.TokenDrops
	}

	return nil
}

// CounterReportAttributesMap represents the map of attribute for CounterReport.
var CounterReportAttributesMap = map[string]elemental.AttributeSpecification{
	"AckInUnknownState": {
		AllowedChoices: []string{},
		ConvertedName:  "AckInUnknownState",
		Description:    `Counter for sending FIN ACK received in unknown connection state.`,
		Exposed:        true,
		Name:           "AckInUnknownState",
		Type:           "integer",
	},
	"AckInvalidFormat": {
		AllowedChoices: []string{},
		ConvertedName:  "AckInvalidFormat",
		Description:    `Counter for ACK packet dropped because of invalid format.`,
		Exposed:        true,
		Name:           "AckInvalidFormat",
		Type:           "integer",
	},
	"AckRejected": {
		AllowedChoices: []string{},
		ConvertedName:  "AckRejected",
		Description:    `Counter for ACK packets rejected as per policy.`,
		Exposed:        true,
		Name:           "AckRejected",
		Type:           "integer",
	},
	"AckSigValidationFailed": {
		AllowedChoices: []string{},
		ConvertedName:  "AckSigValidationFailed",
		Description:    `Counter for ACK packet dropped because signature validation failed.`,
		Exposed:        true,
		Name:           "AckSigValidationFailed",
		Type:           "integer",
	},
	"AckTCPNoTCPAuthOption": {
		AllowedChoices: []string{},
		ConvertedName:  "AckTCPNoTCPAuthOption",
		Description:    `Counter for TCP authentication option not found.`,
		Exposed:        true,
		Name:           "AckTCPNoTCPAuthOption",
		Type:           "integer",
	},
	"ConnectionsProcessed": {
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsProcessed",
		Description:    `Counter for connections processed.`,
		Exposed:        true,
		Name:           "ConnectionsProcessed",
		Type:           "integer",
	},
	"ContextIDNotFound": {
		AllowedChoices: []string{},
		ConvertedName:  "ContextIDNotFound",
		Description:    `Counter for unable to find ContextID.`,
		Exposed:        true,
		Name:           "ContextIDNotFound",
		Type:           "integer",
	},
	"DroppedExternalService": {
		AllowedChoices: []string{},
		ConvertedName:  "DroppedExternalService",
		Description: `Counter for no ACLs found for external services. Dropping application SYN
packet.`,
		Exposed: true,
		Name:    "DroppedExternalService",
		Type:    "integer",
	},
	"InvalidConnState": {
		AllowedChoices: []string{},
		ConvertedName:  "InvalidConnState",
		Description:    `Counter for invalid connection state.`,
		Exposed:        true,
		Name:           "InvalidConnState",
		Type:           "integer",
	},
	"InvalidNetState": {
		AllowedChoices: []string{},
		ConvertedName:  "InvalidNetState",
		Description:    `Counter for invalid net state.`,
		Exposed:        true,
		Name:           "InvalidNetState",
		Type:           "integer",
	},
	"InvalidProtocol": {
		AllowedChoices: []string{},
		ConvertedName:  "InvalidProtocol",
		Description:    `Counter for invalid protocol.`,
		Exposed:        true,
		Name:           "InvalidProtocol",
		Type:           "integer",
	},
	"InvalidSynAck": {
		AllowedChoices: []string{},
		ConvertedName:  "InvalidSynAck",
		Description:    `Counter for processing unit is already dead - drop SYN ACK packet.`,
		Exposed:        true,
		Name:           "InvalidSynAck",
		Type:           "integer",
	},
	"MarkNotFound": {
		AllowedChoices: []string{},
		ConvertedName:  "MarkNotFound",
		Description:    `Counter for processing unit mark not found.`,
		Exposed:        true,
		Name:           "MarkNotFound",
		Type:           "integer",
	},
	"NetSynNotSeen": {
		AllowedChoices: []string{},
		ConvertedName:  "NetSynNotSeen",
		Description:    `Counter for network SYN packet was not seen.`,
		Exposed:        true,
		Name:           "NetSynNotSeen",
		Type:           "integer",
	},
	"NoConnFound": {
		AllowedChoices: []string{},
		ConvertedName:  "NoConnFound",
		Description:    `Counter for no context or connection found.`,
		Exposed:        true,
		Name:           "NoConnFound",
		Type:           "integer",
	},
	"NonPUTraffic": {
		AllowedChoices: []string{},
		ConvertedName:  "NonPUTraffic",
		Description:    `Counter for traffic that belongs to a non-processing unit process.`,
		Exposed:        true,
		Name:           "NonPUTraffic",
		Type:           "integer",
	},
	"OutOfOrderSynAck": {
		AllowedChoices: []string{},
		ConvertedName:  "OutOfOrderSynAck",
		Description:    `Counter for SYN ACK for flow with processed FIN ACK.`,
		Exposed:        true,
		Name:           "OutOfOrderSynAck",
		Type:           "integer",
	},
	"PortNotFound": {
		AllowedChoices: []string{},
		ConvertedName:  "PortNotFound",
		Description:    `Counter for port not found.`,
		Exposed:        true,
		Name:           "PortNotFound",
		Type:           "integer",
	},
	"RejectPacket": {
		AllowedChoices: []string{},
		ConvertedName:  "RejectPacket",
		Description:    `Counter for reject the packet as per policy.`,
		Exposed:        true,
		Name:           "RejectPacket",
		Type:           "integer",
	},
	"ServicePostprocessorFailed": {
		AllowedChoices: []string{},
		ConvertedName:  "ServicePostprocessorFailed",
		Description:    `Counter for post service processing failed for network packet.`,
		Exposed:        true,
		Name:           "ServicePostprocessorFailed",
		Type:           "integer",
	},
	"ServicePreprocessorFailed": {
		AllowedChoices: []string{},
		ConvertedName:  "ServicePreprocessorFailed",
		Description:    `Counter for network packets that failed preprocessing.`,
		Exposed:        true,
		Name:           "ServicePreprocessorFailed",
		Type:           "integer",
	},
	"SynAckBadClaims": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckBadClaims",
		Description:    `Counter for SYN ACK packet dropped because of bad claims.`,
		Exposed:        true,
		Name:           "SynAckBadClaims",
		Type:           "integer",
	},
	"SynAckClaimsMisMatch": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckClaimsMisMatch",
		Description:    `Counter for SYN ACK packet dropped because of encryption mismatch.`,
		Exposed:        true,
		Name:           "SynAckClaimsMisMatch",
		Type:           "integer",
	},
	"SynAckDroppedExternalService": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckDroppedExternalService",
		Description:    `Counter for SYN ACK from external service dropped.`,
		Exposed:        true,
		Name:           "SynAckDroppedExternalService",
		Type:           "integer",
	},
	"SynAckInvalidFormat": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckInvalidFormat",
		Description:    `Counter for SYN ACK packet dropped because of invalid format.`,
		Exposed:        true,
		Name:           "SynAckInvalidFormat",
		Type:           "integer",
	},
	"SynAckMissingClaims": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckMissingClaims",
		Description:    `Counter for SYN ACK packet dropped because of no claims.`,
		Exposed:        true,
		Name:           "SynAckMissingClaims",
		Type:           "integer",
	},
	"SynAckMissingToken": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckMissingToken",
		Description:    `Counter for SYN ACK packet dropped because of missing token.`,
		Exposed:        true,
		Name:           "SynAckMissingToken",
		Type:           "integer",
	},
	"SynAckNoTCPAuthOption": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckNoTCPAuthOption",
		Description:    `Counter for TCP authentication option not found.`,
		Exposed:        true,
		Name:           "SynAckNoTCPAuthOption",
		Type:           "integer",
	},
	"SynAckRejected": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckRejected",
		Description:    `Counter for dropping because of reject rule on transmitter.`,
		Exposed:        true,
		Name:           "SynAckRejected",
		Type:           "integer",
	},
	"SynDroppedInvalidFormat": {
		AllowedChoices: []string{},
		ConvertedName:  "SynDroppedInvalidFormat",
		Description:    `Counter for SYN packet dropped because of invalid format.`,
		Exposed:        true,
		Name:           "SynDroppedInvalidFormat",
		Type:           "integer",
	},
	"SynDroppedInvalidToken": {
		AllowedChoices: []string{},
		ConvertedName:  "SynDroppedInvalidToken",
		Description:    `Counter for SYN packet dropped because of invalid token.`,
		Exposed:        true,
		Name:           "SynDroppedInvalidToken",
		Type:           "integer",
	},
	"SynDroppedNoClaims": {
		AllowedChoices: []string{},
		ConvertedName:  "SynDroppedNoClaims",
		Description:    `Counter for SYN packet dropped because of no claims.`,
		Exposed:        true,
		Name:           "SynDroppedNoClaims",
		Type:           "integer",
	},
	"SynDroppedTCPOption": {
		AllowedChoices: []string{},
		ConvertedName:  "SynDroppedTCPOption",
		Description:    `Counter for TCP authentication option not found.`,
		Exposed:        true,
		Name:           "SynDroppedTCPOption",
		Type:           "integer",
	},
	"SynRejectPacket": {
		AllowedChoices: []string{},
		ConvertedName:  "SynRejectPacket",
		Description:    `Counter for SYN packet dropped due to policy.`,
		Exposed:        true,
		Name:           "SynRejectPacket",
		Type:           "integer",
	},
	"SynUnexpectedPacket": {
		AllowedChoices: []string{},
		ConvertedName:  "SynUnexpectedPacket",
		Description:    `Counter for received SYN packet from unknown processing unit.`,
		Exposed:        true,
		Name:           "SynUnexpectedPacket",
		Type:           "integer",
	},
	"TCPAuthNotFound": {
		AllowedChoices: []string{},
		ConvertedName:  "TCPAuthNotFound",
		Description:    `Counter for TCP authentication option not found.`,
		Exposed:        true,
		Name:           "TCPAuthNotFound",
		Type:           "integer",
	},
	"UDPAckInvalidSignature": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPAckInvalidSignature",
		Description:    `Counter for UDP ACK packet dropped due to an invalid signature.`,
		Exposed:        true,
		Name:           "UDPAckInvalidSignature",
		Type:           "integer",
	},
	"UDPConnectionsProcessed": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPConnectionsProcessed",
		Description:    `Counter for number of processed UDP connections.`,
		Exposed:        true,
		Name:           "UDPConnectionsProcessed",
		Type:           "integer",
	},
	"UDPDropContextNotFound": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropContextNotFound",
		Description:    `Counter for dropped UDP data packets with no context.`,
		Exposed:        true,
		Name:           "UDPDropContextNotFound",
		Type:           "integer",
	},
	"UDPDropFin": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropFin",
		Description:    `Counter for dropped UDP FIN handshake packets.`,
		Exposed:        true,
		Name:           "UDPDropFin",
		Type:           "integer",
	},
	"UDPDropInNfQueue": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropInNfQueue",
		Description:    `Counter for dropped UDP in NfQueue.`,
		Exposed:        true,
		Name:           "UDPDropInNfQueue",
		Type:           "integer",
	},
	"UDPDropNoConnection": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropNoConnection",
		Description:    `Counter for dropped UDP data packets with no connection.`,
		Exposed:        true,
		Name:           "UDPDropNoConnection",
		Type:           "integer",
	},
	"UDPDropPacket": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropPacket",
		Description:    `Counter for dropped UDP data packets.`,
		Exposed:        true,
		Name:           "UDPDropPacket",
		Type:           "integer",
	},
	"UDPDropQueueFull": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropQueueFull",
		Description:    `Counter for dropped UDP Queue Full.`,
		Exposed:        true,
		Name:           "UDPDropQueueFull",
		Type:           "integer",
	},
	"UDPDropSynAck": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropSynAck",
		Description:    `Counter for dropped UDP SYN ACK handshake packets.`,
		Exposed:        true,
		Name:           "UDPDropSynAck",
		Type:           "integer",
	},
	"UDPInvalidNetState": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPInvalidNetState",
		Description:    `Counter for UDP packets received in invalid network state.`,
		Exposed:        true,
		Name:           "UDPInvalidNetState",
		Type:           "integer",
	},
	"UDPPostProcessingFailed": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPPostProcessingFailed",
		Description:    `Counter for UDP packets failing postprocessing.`,
		Exposed:        true,
		Name:           "UDPPostProcessingFailed",
		Type:           "integer",
	},
	"UDPPreProcessingFailed": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPPreProcessingFailed",
		Description:    `Counter for UDP packets failing preprocessing.`,
		Exposed:        true,
		Name:           "UDPPreProcessingFailed",
		Type:           "integer",
	},
	"UDPRejected": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPRejected",
		Description:    `Counter for UDP packets dropped due to policy.`,
		Exposed:        true,
		Name:           "UDPRejected",
		Type:           "integer",
	},
	"UDPSynAckDropBadClaims": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynAckDropBadClaims",
		Description:    `Counter for UDP SYN ACK packets dropped due to bad claims.`,
		Exposed:        true,
		Name:           "UDPSynAckDropBadClaims",
		Type:           "integer",
	},
	"UDPSynAckMissingClaims": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynAckMissingClaims",
		Description:    `Counter for UDP SYN ACK packets dropped due to missing claims.`,
		Exposed:        true,
		Name:           "UDPSynAckMissingClaims",
		Type:           "integer",
	},
	"UDPSynAckPolicy": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynAckPolicy",
		Description:    `Counter for UDP SYN ACK packets dropped due to bad claims.`,
		Exposed:        true,
		Name:           "UDPSynAckPolicy",
		Type:           "integer",
	},
	"UDPSynDrop": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynDrop",
		Description:    `Counter for dropped UDP SYN transmits.`,
		Exposed:        true,
		Name:           "UDPSynDrop",
		Type:           "integer",
	},
	"UDPSynDropPolicy": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynDropPolicy",
		Description:    `Counter for dropped UDP SYN policy.`,
		Exposed:        true,
		Name:           "UDPSynDropPolicy",
		Type:           "integer",
	},
	"UDPSynInvalidToken": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynInvalidToken",
		Description:    `Counter for dropped UDP FIN handshake packets.`,
		Exposed:        true,
		Name:           "UDPSynInvalidToken",
		Type:           "integer",
	},
	"UDPSynMissingClaims": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynMissingClaims",
		Description:    `Counter for UDP SYN packet dropped due to missing claims.`,
		Exposed:        true,
		Name:           "UDPSynMissingClaims",
		Type:           "integer",
	},
	"UnknownError": {
		AllowedChoices: []string{},
		ConvertedName:  "UnknownError",
		Description:    `Counter for unknown error.`,
		Exposed:        true,
		Name:           "UnknownError",
		Type:           "integer",
	},
	"ConnectionsAnalyzed": {
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsAnalyzed",
		Description: `Non-zero counter indicates analyzed connections for unencrypted, encrypted, and
packets from endpoint applications with the TCP Fast Open option set. These are 
not dropped counter.`,
		Exposed: true,
		Name:    "connectionsAnalyzed",
		Type:    "integer",
	},
	"ConnectionsDropped": {
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsDropped",
		Description: `Non-zero counter indicates dropped connections because of invalid state, 
non-processing unit traffic, or out of order packets.`,
		Exposed: true,
		Name:    "connectionsDropped",
		Type:    "integer",
	},
	"ConnectionsExpired": {
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsExpired",
		Description: `Non-zero counter indicates expired connections because of response not being
received within a certain amount of time after the request is made.`,
		Exposed: true,
		Name:    "connectionsExpired",
		Type:    "integer",
	},
	"DroppedPackets": {
		AllowedChoices: []string{},
		ConvertedName:  "DroppedPackets",
		Description: `Non-zero counter indicates dropped packets that did not hit any of our iptables
rules and queue drops.`,
		Exposed: true,
		Name:    "droppedPackets",
		Type:    "integer",
	},
	"EncryptionFailures": {
		AllowedChoices: []string{},
		ConvertedName:  "EncryptionFailures",
		Description:    `Non-zero counter indicates encryption processing failures of data packets.`,
		Exposed:        true,
		Name:           "encryptionFailures",
		Type:           "integer",
	},
	"EnforcerID": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `Identifier of the enforcer sending the report.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"EnforcerNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the enforcer sending the report.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"ExternalNetworkConnections": {
		AllowedChoices: []string{},
		ConvertedName:  "ExternalNetworkConnections",
		Description: `Non-zero counter indicates connections going to and from external networks.
These may be drops or allowed counters.`,
		Exposed: true,
		Name:    "externalNetworkConnections",
		Type:    "integer",
	},
	"PolicyDrops": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyDrops",
		Description:    `Non-zero counter indicates packets dropped due to a reject policy.`,
		Exposed:        true,
		Name:           "policyDrops",
		Type:           "integer",
	},
	"ProcessingUnitID": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `PUID is the ID of the processing unit reporting the counter.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "processingUnitID",
		Type:           "string",
	},
	"ProcessingUnitNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the processing unit reporting the counter.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "processingUnitNamespace",
		Type:           "string",
	},
	"Timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Timestamp is the date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
	"TokenDrops": {
		AllowedChoices: []string{},
		ConvertedName:  "TokenDrops",
		Description: `Non-zero counter indicates packets rejected due to anything related to token
creation/parsing failures.`,
		Exposed: true,
		Name:    "tokenDrops",
		Type:    "integer",
	},
}

// CounterReportLowerCaseAttributesMap represents the map of attribute for CounterReport.
var CounterReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"ackinunknownstate": {
		AllowedChoices: []string{},
		ConvertedName:  "AckInUnknownState",
		Description:    `Counter for sending FIN ACK received in unknown connection state.`,
		Exposed:        true,
		Name:           "AckInUnknownState",
		Type:           "integer",
	},
	"ackinvalidformat": {
		AllowedChoices: []string{},
		ConvertedName:  "AckInvalidFormat",
		Description:    `Counter for ACK packet dropped because of invalid format.`,
		Exposed:        true,
		Name:           "AckInvalidFormat",
		Type:           "integer",
	},
	"ackrejected": {
		AllowedChoices: []string{},
		ConvertedName:  "AckRejected",
		Description:    `Counter for ACK packets rejected as per policy.`,
		Exposed:        true,
		Name:           "AckRejected",
		Type:           "integer",
	},
	"acksigvalidationfailed": {
		AllowedChoices: []string{},
		ConvertedName:  "AckSigValidationFailed",
		Description:    `Counter for ACK packet dropped because signature validation failed.`,
		Exposed:        true,
		Name:           "AckSigValidationFailed",
		Type:           "integer",
	},
	"acktcpnotcpauthoption": {
		AllowedChoices: []string{},
		ConvertedName:  "AckTCPNoTCPAuthOption",
		Description:    `Counter for TCP authentication option not found.`,
		Exposed:        true,
		Name:           "AckTCPNoTCPAuthOption",
		Type:           "integer",
	},
	"connectionsprocessed": {
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsProcessed",
		Description:    `Counter for connections processed.`,
		Exposed:        true,
		Name:           "ConnectionsProcessed",
		Type:           "integer",
	},
	"contextidnotfound": {
		AllowedChoices: []string{},
		ConvertedName:  "ContextIDNotFound",
		Description:    `Counter for unable to find ContextID.`,
		Exposed:        true,
		Name:           "ContextIDNotFound",
		Type:           "integer",
	},
	"droppedexternalservice": {
		AllowedChoices: []string{},
		ConvertedName:  "DroppedExternalService",
		Description: `Counter for no ACLs found for external services. Dropping application SYN
packet.`,
		Exposed: true,
		Name:    "DroppedExternalService",
		Type:    "integer",
	},
	"invalidconnstate": {
		AllowedChoices: []string{},
		ConvertedName:  "InvalidConnState",
		Description:    `Counter for invalid connection state.`,
		Exposed:        true,
		Name:           "InvalidConnState",
		Type:           "integer",
	},
	"invalidnetstate": {
		AllowedChoices: []string{},
		ConvertedName:  "InvalidNetState",
		Description:    `Counter for invalid net state.`,
		Exposed:        true,
		Name:           "InvalidNetState",
		Type:           "integer",
	},
	"invalidprotocol": {
		AllowedChoices: []string{},
		ConvertedName:  "InvalidProtocol",
		Description:    `Counter for invalid protocol.`,
		Exposed:        true,
		Name:           "InvalidProtocol",
		Type:           "integer",
	},
	"invalidsynack": {
		AllowedChoices: []string{},
		ConvertedName:  "InvalidSynAck",
		Description:    `Counter for processing unit is already dead - drop SYN ACK packet.`,
		Exposed:        true,
		Name:           "InvalidSynAck",
		Type:           "integer",
	},
	"marknotfound": {
		AllowedChoices: []string{},
		ConvertedName:  "MarkNotFound",
		Description:    `Counter for processing unit mark not found.`,
		Exposed:        true,
		Name:           "MarkNotFound",
		Type:           "integer",
	},
	"netsynnotseen": {
		AllowedChoices: []string{},
		ConvertedName:  "NetSynNotSeen",
		Description:    `Counter for network SYN packet was not seen.`,
		Exposed:        true,
		Name:           "NetSynNotSeen",
		Type:           "integer",
	},
	"noconnfound": {
		AllowedChoices: []string{},
		ConvertedName:  "NoConnFound",
		Description:    `Counter for no context or connection found.`,
		Exposed:        true,
		Name:           "NoConnFound",
		Type:           "integer",
	},
	"nonputraffic": {
		AllowedChoices: []string{},
		ConvertedName:  "NonPUTraffic",
		Description:    `Counter for traffic that belongs to a non-processing unit process.`,
		Exposed:        true,
		Name:           "NonPUTraffic",
		Type:           "integer",
	},
	"outofordersynack": {
		AllowedChoices: []string{},
		ConvertedName:  "OutOfOrderSynAck",
		Description:    `Counter for SYN ACK for flow with processed FIN ACK.`,
		Exposed:        true,
		Name:           "OutOfOrderSynAck",
		Type:           "integer",
	},
	"portnotfound": {
		AllowedChoices: []string{},
		ConvertedName:  "PortNotFound",
		Description:    `Counter for port not found.`,
		Exposed:        true,
		Name:           "PortNotFound",
		Type:           "integer",
	},
	"rejectpacket": {
		AllowedChoices: []string{},
		ConvertedName:  "RejectPacket",
		Description:    `Counter for reject the packet as per policy.`,
		Exposed:        true,
		Name:           "RejectPacket",
		Type:           "integer",
	},
	"servicepostprocessorfailed": {
		AllowedChoices: []string{},
		ConvertedName:  "ServicePostprocessorFailed",
		Description:    `Counter for post service processing failed for network packet.`,
		Exposed:        true,
		Name:           "ServicePostprocessorFailed",
		Type:           "integer",
	},
	"servicepreprocessorfailed": {
		AllowedChoices: []string{},
		ConvertedName:  "ServicePreprocessorFailed",
		Description:    `Counter for network packets that failed preprocessing.`,
		Exposed:        true,
		Name:           "ServicePreprocessorFailed",
		Type:           "integer",
	},
	"synackbadclaims": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckBadClaims",
		Description:    `Counter for SYN ACK packet dropped because of bad claims.`,
		Exposed:        true,
		Name:           "SynAckBadClaims",
		Type:           "integer",
	},
	"synackclaimsmismatch": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckClaimsMisMatch",
		Description:    `Counter for SYN ACK packet dropped because of encryption mismatch.`,
		Exposed:        true,
		Name:           "SynAckClaimsMisMatch",
		Type:           "integer",
	},
	"synackdroppedexternalservice": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckDroppedExternalService",
		Description:    `Counter for SYN ACK from external service dropped.`,
		Exposed:        true,
		Name:           "SynAckDroppedExternalService",
		Type:           "integer",
	},
	"synackinvalidformat": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckInvalidFormat",
		Description:    `Counter for SYN ACK packet dropped because of invalid format.`,
		Exposed:        true,
		Name:           "SynAckInvalidFormat",
		Type:           "integer",
	},
	"synackmissingclaims": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckMissingClaims",
		Description:    `Counter for SYN ACK packet dropped because of no claims.`,
		Exposed:        true,
		Name:           "SynAckMissingClaims",
		Type:           "integer",
	},
	"synackmissingtoken": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckMissingToken",
		Description:    `Counter for SYN ACK packet dropped because of missing token.`,
		Exposed:        true,
		Name:           "SynAckMissingToken",
		Type:           "integer",
	},
	"synacknotcpauthoption": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckNoTCPAuthOption",
		Description:    `Counter for TCP authentication option not found.`,
		Exposed:        true,
		Name:           "SynAckNoTCPAuthOption",
		Type:           "integer",
	},
	"synackrejected": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckRejected",
		Description:    `Counter for dropping because of reject rule on transmitter.`,
		Exposed:        true,
		Name:           "SynAckRejected",
		Type:           "integer",
	},
	"syndroppedinvalidformat": {
		AllowedChoices: []string{},
		ConvertedName:  "SynDroppedInvalidFormat",
		Description:    `Counter for SYN packet dropped because of invalid format.`,
		Exposed:        true,
		Name:           "SynDroppedInvalidFormat",
		Type:           "integer",
	},
	"syndroppedinvalidtoken": {
		AllowedChoices: []string{},
		ConvertedName:  "SynDroppedInvalidToken",
		Description:    `Counter for SYN packet dropped because of invalid token.`,
		Exposed:        true,
		Name:           "SynDroppedInvalidToken",
		Type:           "integer",
	},
	"syndroppednoclaims": {
		AllowedChoices: []string{},
		ConvertedName:  "SynDroppedNoClaims",
		Description:    `Counter for SYN packet dropped because of no claims.`,
		Exposed:        true,
		Name:           "SynDroppedNoClaims",
		Type:           "integer",
	},
	"syndroppedtcpoption": {
		AllowedChoices: []string{},
		ConvertedName:  "SynDroppedTCPOption",
		Description:    `Counter for TCP authentication option not found.`,
		Exposed:        true,
		Name:           "SynDroppedTCPOption",
		Type:           "integer",
	},
	"synrejectpacket": {
		AllowedChoices: []string{},
		ConvertedName:  "SynRejectPacket",
		Description:    `Counter for SYN packet dropped due to policy.`,
		Exposed:        true,
		Name:           "SynRejectPacket",
		Type:           "integer",
	},
	"synunexpectedpacket": {
		AllowedChoices: []string{},
		ConvertedName:  "SynUnexpectedPacket",
		Description:    `Counter for received SYN packet from unknown processing unit.`,
		Exposed:        true,
		Name:           "SynUnexpectedPacket",
		Type:           "integer",
	},
	"tcpauthnotfound": {
		AllowedChoices: []string{},
		ConvertedName:  "TCPAuthNotFound",
		Description:    `Counter for TCP authentication option not found.`,
		Exposed:        true,
		Name:           "TCPAuthNotFound",
		Type:           "integer",
	},
	"udpackinvalidsignature": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPAckInvalidSignature",
		Description:    `Counter for UDP ACK packet dropped due to an invalid signature.`,
		Exposed:        true,
		Name:           "UDPAckInvalidSignature",
		Type:           "integer",
	},
	"udpconnectionsprocessed": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPConnectionsProcessed",
		Description:    `Counter for number of processed UDP connections.`,
		Exposed:        true,
		Name:           "UDPConnectionsProcessed",
		Type:           "integer",
	},
	"udpdropcontextnotfound": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropContextNotFound",
		Description:    `Counter for dropped UDP data packets with no context.`,
		Exposed:        true,
		Name:           "UDPDropContextNotFound",
		Type:           "integer",
	},
	"udpdropfin": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropFin",
		Description:    `Counter for dropped UDP FIN handshake packets.`,
		Exposed:        true,
		Name:           "UDPDropFin",
		Type:           "integer",
	},
	"udpdropinnfqueue": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropInNfQueue",
		Description:    `Counter for dropped UDP in NfQueue.`,
		Exposed:        true,
		Name:           "UDPDropInNfQueue",
		Type:           "integer",
	},
	"udpdropnoconnection": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropNoConnection",
		Description:    `Counter for dropped UDP data packets with no connection.`,
		Exposed:        true,
		Name:           "UDPDropNoConnection",
		Type:           "integer",
	},
	"udpdroppacket": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropPacket",
		Description:    `Counter for dropped UDP data packets.`,
		Exposed:        true,
		Name:           "UDPDropPacket",
		Type:           "integer",
	},
	"udpdropqueuefull": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropQueueFull",
		Description:    `Counter for dropped UDP Queue Full.`,
		Exposed:        true,
		Name:           "UDPDropQueueFull",
		Type:           "integer",
	},
	"udpdropsynack": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropSynAck",
		Description:    `Counter for dropped UDP SYN ACK handshake packets.`,
		Exposed:        true,
		Name:           "UDPDropSynAck",
		Type:           "integer",
	},
	"udpinvalidnetstate": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPInvalidNetState",
		Description:    `Counter for UDP packets received in invalid network state.`,
		Exposed:        true,
		Name:           "UDPInvalidNetState",
		Type:           "integer",
	},
	"udppostprocessingfailed": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPPostProcessingFailed",
		Description:    `Counter for UDP packets failing postprocessing.`,
		Exposed:        true,
		Name:           "UDPPostProcessingFailed",
		Type:           "integer",
	},
	"udppreprocessingfailed": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPPreProcessingFailed",
		Description:    `Counter for UDP packets failing preprocessing.`,
		Exposed:        true,
		Name:           "UDPPreProcessingFailed",
		Type:           "integer",
	},
	"udprejected": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPRejected",
		Description:    `Counter for UDP packets dropped due to policy.`,
		Exposed:        true,
		Name:           "UDPRejected",
		Type:           "integer",
	},
	"udpsynackdropbadclaims": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynAckDropBadClaims",
		Description:    `Counter for UDP SYN ACK packets dropped due to bad claims.`,
		Exposed:        true,
		Name:           "UDPSynAckDropBadClaims",
		Type:           "integer",
	},
	"udpsynackmissingclaims": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynAckMissingClaims",
		Description:    `Counter for UDP SYN ACK packets dropped due to missing claims.`,
		Exposed:        true,
		Name:           "UDPSynAckMissingClaims",
		Type:           "integer",
	},
	"udpsynackpolicy": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynAckPolicy",
		Description:    `Counter for UDP SYN ACK packets dropped due to bad claims.`,
		Exposed:        true,
		Name:           "UDPSynAckPolicy",
		Type:           "integer",
	},
	"udpsyndrop": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynDrop",
		Description:    `Counter for dropped UDP SYN transmits.`,
		Exposed:        true,
		Name:           "UDPSynDrop",
		Type:           "integer",
	},
	"udpsyndroppolicy": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynDropPolicy",
		Description:    `Counter for dropped UDP SYN policy.`,
		Exposed:        true,
		Name:           "UDPSynDropPolicy",
		Type:           "integer",
	},
	"udpsyninvalidtoken": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynInvalidToken",
		Description:    `Counter for dropped UDP FIN handshake packets.`,
		Exposed:        true,
		Name:           "UDPSynInvalidToken",
		Type:           "integer",
	},
	"udpsynmissingclaims": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynMissingClaims",
		Description:    `Counter for UDP SYN packet dropped due to missing claims.`,
		Exposed:        true,
		Name:           "UDPSynMissingClaims",
		Type:           "integer",
	},
	"unknownerror": {
		AllowedChoices: []string{},
		ConvertedName:  "UnknownError",
		Description:    `Counter for unknown error.`,
		Exposed:        true,
		Name:           "UnknownError",
		Type:           "integer",
	},
	"connectionsanalyzed": {
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsAnalyzed",
		Description: `Non-zero counter indicates analyzed connections for unencrypted, encrypted, and
packets from endpoint applications with the TCP Fast Open option set. These are 
not dropped counter.`,
		Exposed: true,
		Name:    "connectionsAnalyzed",
		Type:    "integer",
	},
	"connectionsdropped": {
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsDropped",
		Description: `Non-zero counter indicates dropped connections because of invalid state, 
non-processing unit traffic, or out of order packets.`,
		Exposed: true,
		Name:    "connectionsDropped",
		Type:    "integer",
	},
	"connectionsexpired": {
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsExpired",
		Description: `Non-zero counter indicates expired connections because of response not being
received within a certain amount of time after the request is made.`,
		Exposed: true,
		Name:    "connectionsExpired",
		Type:    "integer",
	},
	"droppedpackets": {
		AllowedChoices: []string{},
		ConvertedName:  "DroppedPackets",
		Description: `Non-zero counter indicates dropped packets that did not hit any of our iptables
rules and queue drops.`,
		Exposed: true,
		Name:    "droppedPackets",
		Type:    "integer",
	},
	"encryptionfailures": {
		AllowedChoices: []string{},
		ConvertedName:  "EncryptionFailures",
		Description:    `Non-zero counter indicates encryption processing failures of data packets.`,
		Exposed:        true,
		Name:           "encryptionFailures",
		Type:           "integer",
	},
	"enforcerid": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `Identifier of the enforcer sending the report.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"enforcernamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the enforcer sending the report.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"externalnetworkconnections": {
		AllowedChoices: []string{},
		ConvertedName:  "ExternalNetworkConnections",
		Description: `Non-zero counter indicates connections going to and from external networks.
These may be drops or allowed counters.`,
		Exposed: true,
		Name:    "externalNetworkConnections",
		Type:    "integer",
	},
	"policydrops": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyDrops",
		Description:    `Non-zero counter indicates packets dropped due to a reject policy.`,
		Exposed:        true,
		Name:           "policyDrops",
		Type:           "integer",
	},
	"processingunitid": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `PUID is the ID of the processing unit reporting the counter.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "processingUnitID",
		Type:           "string",
	},
	"processingunitnamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the processing unit reporting the counter.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "processingUnitNamespace",
		Type:           "string",
	},
	"timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Timestamp is the date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
	"tokendrops": {
		AllowedChoices: []string{},
		ConvertedName:  "TokenDrops",
		Description: `Non-zero counter indicates packets rejected due to anything related to token
creation/parsing failures.`,
		Exposed: true,
		Name:    "tokenDrops",
		Type:    "integer",
	},
}

// SparseCounterReportsList represents a list of SparseCounterReports
type SparseCounterReportsList []*SparseCounterReport

// Identity returns the identity of the objects in the list.
func (o SparseCounterReportsList) Identity() elemental.Identity {

	return CounterReportIdentity
}

// Copy returns a pointer to a copy the SparseCounterReportsList.
func (o SparseCounterReportsList) Copy() elemental.Identifiables {

	copy := append(SparseCounterReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCounterReportsList.
func (o SparseCounterReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCounterReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCounterReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCounterReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCounterReportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCounterReportsList converted to CounterReportsList.
func (o SparseCounterReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCounterReportsList) Version() int {

	return 1
}

// SparseCounterReport represents the sparse version of a counterreport.
type SparseCounterReport struct {
	// Counter for sending FIN ACK received in unknown connection state.
	AckInUnknownState *int `json:"AckInUnknownState,omitempty" msgpack:"AckInUnknownState,omitempty" bson:"-" mapstructure:"AckInUnknownState,omitempty"`

	// Counter for ACK packet dropped because of invalid format.
	AckInvalidFormat *int `json:"AckInvalidFormat,omitempty" msgpack:"AckInvalidFormat,omitempty" bson:"-" mapstructure:"AckInvalidFormat,omitempty"`

	// Counter for ACK packets rejected as per policy.
	AckRejected *int `json:"AckRejected,omitempty" msgpack:"AckRejected,omitempty" bson:"-" mapstructure:"AckRejected,omitempty"`

	// Counter for ACK packet dropped because signature validation failed.
	AckSigValidationFailed *int `json:"AckSigValidationFailed,omitempty" msgpack:"AckSigValidationFailed,omitempty" bson:"-" mapstructure:"AckSigValidationFailed,omitempty"`

	// Counter for TCP authentication option not found.
	AckTCPNoTCPAuthOption *int `json:"AckTCPNoTCPAuthOption,omitempty" msgpack:"AckTCPNoTCPAuthOption,omitempty" bson:"-" mapstructure:"AckTCPNoTCPAuthOption,omitempty"`

	// Counter for connections processed.
	ConnectionsProcessed *int `json:"ConnectionsProcessed,omitempty" msgpack:"ConnectionsProcessed,omitempty" bson:"-" mapstructure:"ConnectionsProcessed,omitempty"`

	// Counter for unable to find ContextID.
	ContextIDNotFound *int `json:"ContextIDNotFound,omitempty" msgpack:"ContextIDNotFound,omitempty" bson:"-" mapstructure:"ContextIDNotFound,omitempty"`

	// Counter for no ACLs found for external services. Dropping application SYN
	// packet.
	DroppedExternalService *int `json:"DroppedExternalService,omitempty" msgpack:"DroppedExternalService,omitempty" bson:"-" mapstructure:"DroppedExternalService,omitempty"`

	// Counter for invalid connection state.
	InvalidConnState *int `json:"InvalidConnState,omitempty" msgpack:"InvalidConnState,omitempty" bson:"-" mapstructure:"InvalidConnState,omitempty"`

	// Counter for invalid net state.
	InvalidNetState *int `json:"InvalidNetState,omitempty" msgpack:"InvalidNetState,omitempty" bson:"-" mapstructure:"InvalidNetState,omitempty"`

	// Counter for invalid protocol.
	InvalidProtocol *int `json:"InvalidProtocol,omitempty" msgpack:"InvalidProtocol,omitempty" bson:"-" mapstructure:"InvalidProtocol,omitempty"`

	// Counter for processing unit is already dead - drop SYN ACK packet.
	InvalidSynAck *int `json:"InvalidSynAck,omitempty" msgpack:"InvalidSynAck,omitempty" bson:"-" mapstructure:"InvalidSynAck,omitempty"`

	// Counter for processing unit mark not found.
	MarkNotFound *int `json:"MarkNotFound,omitempty" msgpack:"MarkNotFound,omitempty" bson:"-" mapstructure:"MarkNotFound,omitempty"`

	// Counter for network SYN packet was not seen.
	NetSynNotSeen *int `json:"NetSynNotSeen,omitempty" msgpack:"NetSynNotSeen,omitempty" bson:"-" mapstructure:"NetSynNotSeen,omitempty"`

	// Counter for no context or connection found.
	NoConnFound *int `json:"NoConnFound,omitempty" msgpack:"NoConnFound,omitempty" bson:"-" mapstructure:"NoConnFound,omitempty"`

	// Counter for traffic that belongs to a non-processing unit process.
	NonPUTraffic *int `json:"NonPUTraffic,omitempty" msgpack:"NonPUTraffic,omitempty" bson:"-" mapstructure:"NonPUTraffic,omitempty"`

	// Counter for SYN ACK for flow with processed FIN ACK.
	OutOfOrderSynAck *int `json:"OutOfOrderSynAck,omitempty" msgpack:"OutOfOrderSynAck,omitempty" bson:"-" mapstructure:"OutOfOrderSynAck,omitempty"`

	// Counter for port not found.
	PortNotFound *int `json:"PortNotFound,omitempty" msgpack:"PortNotFound,omitempty" bson:"-" mapstructure:"PortNotFound,omitempty"`

	// Counter for reject the packet as per policy.
	RejectPacket *int `json:"RejectPacket,omitempty" msgpack:"RejectPacket,omitempty" bson:"-" mapstructure:"RejectPacket,omitempty"`

	// Counter for post service processing failed for network packet.
	ServicePostprocessorFailed *int `json:"ServicePostprocessorFailed,omitempty" msgpack:"ServicePostprocessorFailed,omitempty" bson:"-" mapstructure:"ServicePostprocessorFailed,omitempty"`

	// Counter for network packets that failed preprocessing.
	ServicePreprocessorFailed *int `json:"ServicePreprocessorFailed,omitempty" msgpack:"ServicePreprocessorFailed,omitempty" bson:"-" mapstructure:"ServicePreprocessorFailed,omitempty"`

	// Counter for SYN ACK packet dropped because of bad claims.
	SynAckBadClaims *int `json:"SynAckBadClaims,omitempty" msgpack:"SynAckBadClaims,omitempty" bson:"-" mapstructure:"SynAckBadClaims,omitempty"`

	// Counter for SYN ACK packet dropped because of encryption mismatch.
	SynAckClaimsMisMatch *int `json:"SynAckClaimsMisMatch,omitempty" msgpack:"SynAckClaimsMisMatch,omitempty" bson:"-" mapstructure:"SynAckClaimsMisMatch,omitempty"`

	// Counter for SYN ACK from external service dropped.
	SynAckDroppedExternalService *int `json:"SynAckDroppedExternalService,omitempty" msgpack:"SynAckDroppedExternalService,omitempty" bson:"-" mapstructure:"SynAckDroppedExternalService,omitempty"`

	// Counter for SYN ACK packet dropped because of invalid format.
	SynAckInvalidFormat *int `json:"SynAckInvalidFormat,omitempty" msgpack:"SynAckInvalidFormat,omitempty" bson:"-" mapstructure:"SynAckInvalidFormat,omitempty"`

	// Counter for SYN ACK packet dropped because of no claims.
	SynAckMissingClaims *int `json:"SynAckMissingClaims,omitempty" msgpack:"SynAckMissingClaims,omitempty" bson:"-" mapstructure:"SynAckMissingClaims,omitempty"`

	// Counter for SYN ACK packet dropped because of missing token.
	SynAckMissingToken *int `json:"SynAckMissingToken,omitempty" msgpack:"SynAckMissingToken,omitempty" bson:"-" mapstructure:"SynAckMissingToken,omitempty"`

	// Counter for TCP authentication option not found.
	SynAckNoTCPAuthOption *int `json:"SynAckNoTCPAuthOption,omitempty" msgpack:"SynAckNoTCPAuthOption,omitempty" bson:"-" mapstructure:"SynAckNoTCPAuthOption,omitempty"`

	// Counter for dropping because of reject rule on transmitter.
	SynAckRejected *int `json:"SynAckRejected,omitempty" msgpack:"SynAckRejected,omitempty" bson:"-" mapstructure:"SynAckRejected,omitempty"`

	// Counter for SYN packet dropped because of invalid format.
	SynDroppedInvalidFormat *int `json:"SynDroppedInvalidFormat,omitempty" msgpack:"SynDroppedInvalidFormat,omitempty" bson:"-" mapstructure:"SynDroppedInvalidFormat,omitempty"`

	// Counter for SYN packet dropped because of invalid token.
	SynDroppedInvalidToken *int `json:"SynDroppedInvalidToken,omitempty" msgpack:"SynDroppedInvalidToken,omitempty" bson:"-" mapstructure:"SynDroppedInvalidToken,omitempty"`

	// Counter for SYN packet dropped because of no claims.
	SynDroppedNoClaims *int `json:"SynDroppedNoClaims,omitempty" msgpack:"SynDroppedNoClaims,omitempty" bson:"-" mapstructure:"SynDroppedNoClaims,omitempty"`

	// Counter for TCP authentication option not found.
	SynDroppedTCPOption *int `json:"SynDroppedTCPOption,omitempty" msgpack:"SynDroppedTCPOption,omitempty" bson:"-" mapstructure:"SynDroppedTCPOption,omitempty"`

	// Counter for SYN packet dropped due to policy.
	SynRejectPacket *int `json:"SynRejectPacket,omitempty" msgpack:"SynRejectPacket,omitempty" bson:"-" mapstructure:"SynRejectPacket,omitempty"`

	// Counter for received SYN packet from unknown processing unit.
	SynUnexpectedPacket *int `json:"SynUnexpectedPacket,omitempty" msgpack:"SynUnexpectedPacket,omitempty" bson:"-" mapstructure:"SynUnexpectedPacket,omitempty"`

	// Counter for TCP authentication option not found.
	TCPAuthNotFound *int `json:"TCPAuthNotFound,omitempty" msgpack:"TCPAuthNotFound,omitempty" bson:"-" mapstructure:"TCPAuthNotFound,omitempty"`

	// Counter for UDP ACK packet dropped due to an invalid signature.
	UDPAckInvalidSignature *int `json:"UDPAckInvalidSignature,omitempty" msgpack:"UDPAckInvalidSignature,omitempty" bson:"-" mapstructure:"UDPAckInvalidSignature,omitempty"`

	// Counter for number of processed UDP connections.
	UDPConnectionsProcessed *int `json:"UDPConnectionsProcessed,omitempty" msgpack:"UDPConnectionsProcessed,omitempty" bson:"-" mapstructure:"UDPConnectionsProcessed,omitempty"`

	// Counter for dropped UDP data packets with no context.
	UDPDropContextNotFound *int `json:"UDPDropContextNotFound,omitempty" msgpack:"UDPDropContextNotFound,omitempty" bson:"-" mapstructure:"UDPDropContextNotFound,omitempty"`

	// Counter for dropped UDP FIN handshake packets.
	UDPDropFin *int `json:"UDPDropFin,omitempty" msgpack:"UDPDropFin,omitempty" bson:"-" mapstructure:"UDPDropFin,omitempty"`

	// Counter for dropped UDP in NfQueue.
	UDPDropInNfQueue *int `json:"UDPDropInNfQueue,omitempty" msgpack:"UDPDropInNfQueue,omitempty" bson:"-" mapstructure:"UDPDropInNfQueue,omitempty"`

	// Counter for dropped UDP data packets with no connection.
	UDPDropNoConnection *int `json:"UDPDropNoConnection,omitempty" msgpack:"UDPDropNoConnection,omitempty" bson:"-" mapstructure:"UDPDropNoConnection,omitempty"`

	// Counter for dropped UDP data packets.
	UDPDropPacket *int `json:"UDPDropPacket,omitempty" msgpack:"UDPDropPacket,omitempty" bson:"-" mapstructure:"UDPDropPacket,omitempty"`

	// Counter for dropped UDP Queue Full.
	UDPDropQueueFull *int `json:"UDPDropQueueFull,omitempty" msgpack:"UDPDropQueueFull,omitempty" bson:"-" mapstructure:"UDPDropQueueFull,omitempty"`

	// Counter for dropped UDP SYN ACK handshake packets.
	UDPDropSynAck *int `json:"UDPDropSynAck,omitempty" msgpack:"UDPDropSynAck,omitempty" bson:"-" mapstructure:"UDPDropSynAck,omitempty"`

	// Counter for UDP packets received in invalid network state.
	UDPInvalidNetState *int `json:"UDPInvalidNetState,omitempty" msgpack:"UDPInvalidNetState,omitempty" bson:"-" mapstructure:"UDPInvalidNetState,omitempty"`

	// Counter for UDP packets failing postprocessing.
	UDPPostProcessingFailed *int `json:"UDPPostProcessingFailed,omitempty" msgpack:"UDPPostProcessingFailed,omitempty" bson:"-" mapstructure:"UDPPostProcessingFailed,omitempty"`

	// Counter for UDP packets failing preprocessing.
	UDPPreProcessingFailed *int `json:"UDPPreProcessingFailed,omitempty" msgpack:"UDPPreProcessingFailed,omitempty" bson:"-" mapstructure:"UDPPreProcessingFailed,omitempty"`

	// Counter for UDP packets dropped due to policy.
	UDPRejected *int `json:"UDPRejected,omitempty" msgpack:"UDPRejected,omitempty" bson:"-" mapstructure:"UDPRejected,omitempty"`

	// Counter for UDP SYN ACK packets dropped due to bad claims.
	UDPSynAckDropBadClaims *int `json:"UDPSynAckDropBadClaims,omitempty" msgpack:"UDPSynAckDropBadClaims,omitempty" bson:"-" mapstructure:"UDPSynAckDropBadClaims,omitempty"`

	// Counter for UDP SYN ACK packets dropped due to missing claims.
	UDPSynAckMissingClaims *int `json:"UDPSynAckMissingClaims,omitempty" msgpack:"UDPSynAckMissingClaims,omitempty" bson:"-" mapstructure:"UDPSynAckMissingClaims,omitempty"`

	// Counter for UDP SYN ACK packets dropped due to bad claims.
	UDPSynAckPolicy *int `json:"UDPSynAckPolicy,omitempty" msgpack:"UDPSynAckPolicy,omitempty" bson:"-" mapstructure:"UDPSynAckPolicy,omitempty"`

	// Counter for dropped UDP SYN transmits.
	UDPSynDrop *int `json:"UDPSynDrop,omitempty" msgpack:"UDPSynDrop,omitempty" bson:"-" mapstructure:"UDPSynDrop,omitempty"`

	// Counter for dropped UDP SYN policy.
	UDPSynDropPolicy *int `json:"UDPSynDropPolicy,omitempty" msgpack:"UDPSynDropPolicy,omitempty" bson:"-" mapstructure:"UDPSynDropPolicy,omitempty"`

	// Counter for dropped UDP FIN handshake packets.
	UDPSynInvalidToken *int `json:"UDPSynInvalidToken,omitempty" msgpack:"UDPSynInvalidToken,omitempty" bson:"-" mapstructure:"UDPSynInvalidToken,omitempty"`

	// Counter for UDP SYN packet dropped due to missing claims.
	UDPSynMissingClaims *int `json:"UDPSynMissingClaims,omitempty" msgpack:"UDPSynMissingClaims,omitempty" bson:"-" mapstructure:"UDPSynMissingClaims,omitempty"`

	// Counter for unknown error.
	UnknownError *int `json:"UnknownError,omitempty" msgpack:"UnknownError,omitempty" bson:"-" mapstructure:"UnknownError,omitempty"`

	// Non-zero counter indicates analyzed connections for unencrypted, encrypted, and
	// packets from endpoint applications with the TCP Fast Open option set. These are
	// not dropped counter.
	ConnectionsAnalyzed *int `json:"connectionsAnalyzed,omitempty" msgpack:"connectionsAnalyzed,omitempty" bson:"-" mapstructure:"connectionsAnalyzed,omitempty"`

	// Non-zero counter indicates dropped connections because of invalid state,
	// non-processing unit traffic, or out of order packets.
	ConnectionsDropped *int `json:"connectionsDropped,omitempty" msgpack:"connectionsDropped,omitempty" bson:"-" mapstructure:"connectionsDropped,omitempty"`

	// Non-zero counter indicates expired connections because of response not being
	// received within a certain amount of time after the request is made.
	ConnectionsExpired *int `json:"connectionsExpired,omitempty" msgpack:"connectionsExpired,omitempty" bson:"-" mapstructure:"connectionsExpired,omitempty"`

	// Non-zero counter indicates dropped packets that did not hit any of our iptables
	// rules and queue drops.
	DroppedPackets *int `json:"droppedPackets,omitempty" msgpack:"droppedPackets,omitempty" bson:"-" mapstructure:"droppedPackets,omitempty"`

	// Non-zero counter indicates encryption processing failures of data packets.
	EncryptionFailures *int `json:"encryptionFailures,omitempty" msgpack:"encryptionFailures,omitempty" bson:"-" mapstructure:"encryptionFailures,omitempty"`

	// Identifier of the enforcer sending the report.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"enforcerid,omitempty" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer sending the report.
	EnforcerNamespace *string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"enforcernamespace,omitempty" mapstructure:"enforcerNamespace,omitempty"`

	// Non-zero counter indicates connections going to and from external networks.
	// These may be drops or allowed counters.
	ExternalNetworkConnections *int `json:"externalNetworkConnections,omitempty" msgpack:"externalNetworkConnections,omitempty" bson:"-" mapstructure:"externalNetworkConnections,omitempty"`

	// Non-zero counter indicates packets dropped due to a reject policy.
	PolicyDrops *int `json:"policyDrops,omitempty" msgpack:"policyDrops,omitempty" bson:"-" mapstructure:"policyDrops,omitempty"`

	// PUID is the ID of the processing unit reporting the counter.
	ProcessingUnitID *string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the processing unit reporting the counter.
	ProcessingUnitNamespace *string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"-" mapstructure:"processingUnitNamespace,omitempty"`

	// Timestamp is the date of the report.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"-" mapstructure:"timestamp,omitempty"`

	// Non-zero counter indicates packets rejected due to anything related to token
	// creation/parsing failures.
	TokenDrops *int `json:"tokenDrops,omitempty" msgpack:"tokenDrops,omitempty" bson:"-" mapstructure:"tokenDrops,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCounterReport returns a new  SparseCounterReport.
func NewSparseCounterReport() *SparseCounterReport {
	return &SparseCounterReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCounterReport) Identity() elemental.Identity {

	return CounterReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCounterReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCounterReport) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCounterReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCounterReport{}

	if o.EnforcerID != nil {
		s.EnforcerID = o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		s.EnforcerNamespace = o.EnforcerNamespace
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCounterReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCounterReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.EnforcerID != nil {
		o.EnforcerID = s.EnforcerID
	}
	if s.EnforcerNamespace != nil {
		o.EnforcerNamespace = s.EnforcerNamespace
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCounterReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCounterReport) ToPlain() elemental.PlainIdentifiable {

	out := NewCounterReport()
	if o.AckInUnknownState != nil {
		out.AckInUnknownState = *o.AckInUnknownState
	}
	if o.AckInvalidFormat != nil {
		out.AckInvalidFormat = *o.AckInvalidFormat
	}
	if o.AckRejected != nil {
		out.AckRejected = *o.AckRejected
	}
	if o.AckSigValidationFailed != nil {
		out.AckSigValidationFailed = *o.AckSigValidationFailed
	}
	if o.AckTCPNoTCPAuthOption != nil {
		out.AckTCPNoTCPAuthOption = *o.AckTCPNoTCPAuthOption
	}
	if o.ConnectionsProcessed != nil {
		out.ConnectionsProcessed = *o.ConnectionsProcessed
	}
	if o.ContextIDNotFound != nil {
		out.ContextIDNotFound = *o.ContextIDNotFound
	}
	if o.DroppedExternalService != nil {
		out.DroppedExternalService = *o.DroppedExternalService
	}
	if o.InvalidConnState != nil {
		out.InvalidConnState = *o.InvalidConnState
	}
	if o.InvalidNetState != nil {
		out.InvalidNetState = *o.InvalidNetState
	}
	if o.InvalidProtocol != nil {
		out.InvalidProtocol = *o.InvalidProtocol
	}
	if o.InvalidSynAck != nil {
		out.InvalidSynAck = *o.InvalidSynAck
	}
	if o.MarkNotFound != nil {
		out.MarkNotFound = *o.MarkNotFound
	}
	if o.NetSynNotSeen != nil {
		out.NetSynNotSeen = *o.NetSynNotSeen
	}
	if o.NoConnFound != nil {
		out.NoConnFound = *o.NoConnFound
	}
	if o.NonPUTraffic != nil {
		out.NonPUTraffic = *o.NonPUTraffic
	}
	if o.OutOfOrderSynAck != nil {
		out.OutOfOrderSynAck = *o.OutOfOrderSynAck
	}
	if o.PortNotFound != nil {
		out.PortNotFound = *o.PortNotFound
	}
	if o.RejectPacket != nil {
		out.RejectPacket = *o.RejectPacket
	}
	if o.ServicePostprocessorFailed != nil {
		out.ServicePostprocessorFailed = *o.ServicePostprocessorFailed
	}
	if o.ServicePreprocessorFailed != nil {
		out.ServicePreprocessorFailed = *o.ServicePreprocessorFailed
	}
	if o.SynAckBadClaims != nil {
		out.SynAckBadClaims = *o.SynAckBadClaims
	}
	if o.SynAckClaimsMisMatch != nil {
		out.SynAckClaimsMisMatch = *o.SynAckClaimsMisMatch
	}
	if o.SynAckDroppedExternalService != nil {
		out.SynAckDroppedExternalService = *o.SynAckDroppedExternalService
	}
	if o.SynAckInvalidFormat != nil {
		out.SynAckInvalidFormat = *o.SynAckInvalidFormat
	}
	if o.SynAckMissingClaims != nil {
		out.SynAckMissingClaims = *o.SynAckMissingClaims
	}
	if o.SynAckMissingToken != nil {
		out.SynAckMissingToken = *o.SynAckMissingToken
	}
	if o.SynAckNoTCPAuthOption != nil {
		out.SynAckNoTCPAuthOption = *o.SynAckNoTCPAuthOption
	}
	if o.SynAckRejected != nil {
		out.SynAckRejected = *o.SynAckRejected
	}
	if o.SynDroppedInvalidFormat != nil {
		out.SynDroppedInvalidFormat = *o.SynDroppedInvalidFormat
	}
	if o.SynDroppedInvalidToken != nil {
		out.SynDroppedInvalidToken = *o.SynDroppedInvalidToken
	}
	if o.SynDroppedNoClaims != nil {
		out.SynDroppedNoClaims = *o.SynDroppedNoClaims
	}
	if o.SynDroppedTCPOption != nil {
		out.SynDroppedTCPOption = *o.SynDroppedTCPOption
	}
	if o.SynRejectPacket != nil {
		out.SynRejectPacket = *o.SynRejectPacket
	}
	if o.SynUnexpectedPacket != nil {
		out.SynUnexpectedPacket = *o.SynUnexpectedPacket
	}
	if o.TCPAuthNotFound != nil {
		out.TCPAuthNotFound = *o.TCPAuthNotFound
	}
	if o.UDPAckInvalidSignature != nil {
		out.UDPAckInvalidSignature = *o.UDPAckInvalidSignature
	}
	if o.UDPConnectionsProcessed != nil {
		out.UDPConnectionsProcessed = *o.UDPConnectionsProcessed
	}
	if o.UDPDropContextNotFound != nil {
		out.UDPDropContextNotFound = *o.UDPDropContextNotFound
	}
	if o.UDPDropFin != nil {
		out.UDPDropFin = *o.UDPDropFin
	}
	if o.UDPDropInNfQueue != nil {
		out.UDPDropInNfQueue = *o.UDPDropInNfQueue
	}
	if o.UDPDropNoConnection != nil {
		out.UDPDropNoConnection = *o.UDPDropNoConnection
	}
	if o.UDPDropPacket != nil {
		out.UDPDropPacket = *o.UDPDropPacket
	}
	if o.UDPDropQueueFull != nil {
		out.UDPDropQueueFull = *o.UDPDropQueueFull
	}
	if o.UDPDropSynAck != nil {
		out.UDPDropSynAck = *o.UDPDropSynAck
	}
	if o.UDPInvalidNetState != nil {
		out.UDPInvalidNetState = *o.UDPInvalidNetState
	}
	if o.UDPPostProcessingFailed != nil {
		out.UDPPostProcessingFailed = *o.UDPPostProcessingFailed
	}
	if o.UDPPreProcessingFailed != nil {
		out.UDPPreProcessingFailed = *o.UDPPreProcessingFailed
	}
	if o.UDPRejected != nil {
		out.UDPRejected = *o.UDPRejected
	}
	if o.UDPSynAckDropBadClaims != nil {
		out.UDPSynAckDropBadClaims = *o.UDPSynAckDropBadClaims
	}
	if o.UDPSynAckMissingClaims != nil {
		out.UDPSynAckMissingClaims = *o.UDPSynAckMissingClaims
	}
	if o.UDPSynAckPolicy != nil {
		out.UDPSynAckPolicy = *o.UDPSynAckPolicy
	}
	if o.UDPSynDrop != nil {
		out.UDPSynDrop = *o.UDPSynDrop
	}
	if o.UDPSynDropPolicy != nil {
		out.UDPSynDropPolicy = *o.UDPSynDropPolicy
	}
	if o.UDPSynInvalidToken != nil {
		out.UDPSynInvalidToken = *o.UDPSynInvalidToken
	}
	if o.UDPSynMissingClaims != nil {
		out.UDPSynMissingClaims = *o.UDPSynMissingClaims
	}
	if o.UnknownError != nil {
		out.UnknownError = *o.UnknownError
	}
	if o.ConnectionsAnalyzed != nil {
		out.ConnectionsAnalyzed = *o.ConnectionsAnalyzed
	}
	if o.ConnectionsDropped != nil {
		out.ConnectionsDropped = *o.ConnectionsDropped
	}
	if o.ConnectionsExpired != nil {
		out.ConnectionsExpired = *o.ConnectionsExpired
	}
	if o.DroppedPackets != nil {
		out.DroppedPackets = *o.DroppedPackets
	}
	if o.EncryptionFailures != nil {
		out.EncryptionFailures = *o.EncryptionFailures
	}
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		out.EnforcerNamespace = *o.EnforcerNamespace
	}
	if o.ExternalNetworkConnections != nil {
		out.ExternalNetworkConnections = *o.ExternalNetworkConnections
	}
	if o.PolicyDrops != nil {
		out.PolicyDrops = *o.PolicyDrops
	}
	if o.ProcessingUnitID != nil {
		out.ProcessingUnitID = *o.ProcessingUnitID
	}
	if o.ProcessingUnitNamespace != nil {
		out.ProcessingUnitNamespace = *o.ProcessingUnitNamespace
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}
	if o.TokenDrops != nil {
		out.TokenDrops = *o.TokenDrops
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCounterReport.
func (o *SparseCounterReport) DeepCopy() *SparseCounterReport {

	if o == nil {
		return nil
	}

	out := &SparseCounterReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCounterReport.
func (o *SparseCounterReport) DeepCopyInto(out *SparseCounterReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCounterReport: %s", err))
	}

	*out = *target.(*SparseCounterReport)
}

type mongoAttributesCounterReport struct {
	EnforcerID        string `bson:"enforcerid"`
	EnforcerNamespace string `bson:"enforcernamespace"`
}
type mongoAttributesSparseCounterReport struct {
	EnforcerID        *string `bson:"enforcerid,omitempty"`
	EnforcerNamespace *string `bson:"enforcernamespace,omitempty"`
}
