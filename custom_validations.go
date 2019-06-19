package gaia

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"go.aporeto.io/gaia/constants"
	"gopkg.in/yaml.v2"

	"go.aporeto.io/elemental"
	"go.aporeto.io/gaia/portutils"
	"go.aporeto.io/gaia/protocols"
)

// ValidatePortString validates a string represents a port or a range of port.
// valid: 443, 443:555
func ValidatePortString(attribute string, portExp string) error {

	// TODO: Use portutils to validate a port
	ports := strings.Split(portExp, ":")
	if len(ports) == 0 || len(ports) > 2 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be a port (xxx) or port range (xxx:yyy)", attribute))
	}

	p1, err := strconv.Atoi(ports[0])
	if err != nil {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be a port (xxx) or port range (xxx:yyy)", attribute))
	}

	if p1 < 1 || p1 > 65535 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be a between 1 and 65535", attribute))
	}

	if len(ports) == 1 {
		return nil
	}

	p2, err := strconv.Atoi(ports[1])
	if err != nil {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be a port (xxx) or port range (xxx:yyy)", attribute))
	}

	if p2 < 1 || p2 > 65535 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be a between 1 and 65535", attribute))
	}

	if p1 >= p2 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' left bound is greater or equal to left bound ", attribute))
	}

	return nil
}

// ValidatePortStringList validates a list of ports.
func ValidatePortStringList(attribute string, ports []string) error {

	for _, port := range ports {
		if err := ValidatePortString(attribute, port); err != nil {
			return err
		}
	}

	return nil
}

var rxDNSName = regexp.MustCompile(`^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`)

// ValidateNetwork validates a CIDR.
func ValidateNetwork(attribute string, network string) error {

	if _, _, err := net.ParseCIDR(network); err == nil {
		return nil
	}

	if rxDNSName.MatchString(network) {
		return nil
	}

	return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be a CIDR or hostname", attribute))
}

// ValidateNetworkList validates a list of networks.
// The list cannot be empty
func ValidateNetworkList(attribute string, networks []string) error {

	if len(networks) == 0 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must not be empty", attribute))
	}

	return ValidateOptionalNetworkList(attribute, networks)
}

// ValidateOptionalNetworkList validates a list of networks.
// It can be empty/
func ValidateOptionalNetworkList(attribute string, networks []string) error {

	for _, network := range networks {
		if err := ValidateNetwork(attribute, network); err != nil {
			return err
		}
	}

	return nil
}

// ValidateProtocol validates a string represents netwotk a protocol.
func ValidateProtocol(attribute string, proto string) error {

	upperProto := strings.ToUpper(proto)
	if protocols.L4ProtocolNumberFromName(upperProto) != -1 {
		return nil
	}

	p, err := strconv.Atoi(proto)
	if err != nil {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' valid protocol name or number", attribute))
	}

	if p < 0 || p > 255 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' protocol number must be between 0 and 255 included", attribute))
	}

	return nil
}

// ValidateProtocolList validates a list of protocols.
func ValidateProtocolList(attribute string, protocols []string) error {

	if len(protocols) == 0 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must not be empty", attribute))
	}

	for _, proto := range protocols {
		if err := ValidateProtocol(attribute, proto); err != nil {
			return err
		}
	}

	return nil
}

// ValidateServiceEntity validates a Service.
func ValidateServiceEntity(service *Service) error {

	var errs elemental.Errors

	switch service.AuthorizationType {

	case ServiceAuthorizationTypeOIDC:

		if service.OIDCProviderURL == "" {
			errs = errs.Append(makeValidationError("OIDCProviderURL", "`OIDCProviderURL` is required when `authorizationType` is set to `OIDC`"))
		}

		if u, err := url.Parse(service.OIDCProviderURL); err != nil || u == nil || u.Scheme != "https" {
			errs = errs.Append(makeValidationError("OIDCProviderURL", "`OIDCProviderURL` must be a valid HTTPS URL: example https://xxx.yyy"))
		}

		if service.OIDCClientID == "" {
			errs = errs.Append(makeValidationError("OIDCClientID", "`OIDCClientID` is required when `authorizationType` is set to `OIDC`"))
		}

		if service.OIDCClientSecret == "" {
			errs = errs.Append(makeValidationError("OIDCClientSecret", "`OIDCClientSecret` is required when `authorizationType` is set to `OIDC`"))
		}

	case ServiceAuthorizationTypeJWT:

		if service.JWTSigningCertificate == "" {
			errs = errs.Append(makeValidationError("JWTSigningCertificate", "`JWTSigningCertificate` is required when `authorizationType` is set to `JWT`"))
		}
	}

	if service.TLSType == ServiceTLSTypeExternal {

		if service.TLSCertificate == "" {
			errs = errs.Append(makeValidationError("TLSCertificate", "`TLSCertificate` is required when `TLSType` is set to `External`"))
		}

		if service.TLSCertificateKey == "" {
			errs = errs.Append(makeValidationError("TLSCertificateKey", "`TLSCertificateKey` is required when `TLSType` is set to `External`"))
		}
	}

	allSubnets := []*net.IPNet{}
	for _, ip := range service.IPs {
		ipNet, err := ipNetFromString(ip)
		if err != nil {
			errs = errs.Append(err)
			continue
		}
		for j := 0; j < len(allSubnets); j++ {
			if allSubnets[j].Contains(ipNet.IP) || ipNet.Contains(allSubnets[j].IP) {
				errs = errs.Append(makeValidationError("IPs", "subnets cannot overlap"))
			}
		}
		allSubnets = append(allSubnets, ipNet)
	}

	allHosts := map[string]bool{}
	for _, name := range service.Hosts {
		if !isFQDN(name) {
			errs = errs.Append(makeValidationError("Hosts", "`Hosts` must be a valid hostname or FQDN, compliant with RF952"))
		}
		if _, ok := allHosts[name]; ok {
			errs = errs.Append(makeValidationError("Hosts", "`Hosts` must be unique"))
		}
		allHosts[name] = true
	}

	if len(service.Hosts) == 0 && len(service.IPs) == 0 {
		errs = errs.Append(makeValidationError("", "You must set at least one value in `hosts` or `IPs`"))
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func makeValidationError(attribute string, message string) elemental.Error {

	err := elemental.NewError(
		"Validation Error",
		message,
		"gaia",
		http.StatusUnprocessableEntity,
	)

	if attribute != "" {
		err.Data = map[string]interface{}{"attribute": attribute}
	}

	return err
}

var regHostServiceName = regexp.MustCompile(`^[a-zA-Z0-9_]{0,11}$`)

// ValidateEnforcerProfile validates a an enforcer profile.
func ValidateEnforcerProfile(enforcerProfile *EnforcerProfile) error {

	// Validate Target Networks
	for _, tn := range enforcerProfile.TargetNetworks {
		_, _, err := net.ParseCIDR(tn)
		if err != nil {
			return makeValidationError("targetNetworks", fmt.Sprintf("%s is not a valid CIDR", tn))
		}
	}

	// Validate trusted CAs
	for i, ca := range enforcerProfile.TrustedCAs {
		rest := []byte(ca)
		var block *pem.Block

		for {
			block, rest = pem.Decode(rest)

			if block == nil || block.Type != "CERTIFICATE" {
				return makeValidationError("trustedCAs", fmt.Sprintf("The CA %d is not a valid CA", i))
			}

			if _, err := x509.ParseCertificate(block.Bytes); err != nil {
				return makeValidationError("trustedCAs", err.Error())
			}

			if len(rest) == 0 {
				break
			}
		}
	}

	return nil
}

// ValidateProcessingUnitServicesList validates a list of processing unit services.
func ValidateProcessingUnitServicesList(attribute string, svcs []*ProcessingUnitService) error {

	if _, _, err := ValidateProcessingUnitServicesListWithoutOverlap(svcs, map[int]*portutils.PortsList{}, map[int]*portutils.PortsRangeList{}); err != nil {
		return makeValidationError(attribute, err.Error())
	}
	return nil
}

// ValidateProcessingUnitServicesListWithoutOverlap validates a list of processing unit services has no overlap with any given parameter.
func ValidateProcessingUnitServicesListWithoutOverlap(svcs []*ProcessingUnitService, cachePortsList map[int]*portutils.PortsList, cacheRanges map[int]*portutils.PortsRangeList) (map[int]*portutils.PortsList, map[int]*portutils.PortsRangeList, error) {

	for i, svc := range svcs {

		if svc == nil {
			return nil, nil, fmt.Errorf("nil processingunitservice in list at index %d", i)
		}

		var cpl *portutils.PortsList
		var cpr *portutils.PortsRangeList
		var ok bool

		if cpl, ok = cachePortsList[svc.Protocol]; !ok {
			cpl = &portutils.PortsList{}
			cachePortsList[svc.Protocol] = cpl
		}

		if cpr, ok = cacheRanges[svc.Protocol]; !ok {
			cpr = &portutils.PortsRangeList{}
			cacheRanges[svc.Protocol] = cpr
		}

		targetPorts := svc.TargetPorts

		for _, ports := range targetPorts {
			// Range port
			if strings.Contains(ports, ":") {

				pr, err := portutils.ConvertToPortsRange(ports)
				if err != nil {
					return nil, nil, err
				}

				if pr.HasOverlapWithPortsRanges(cpr) {
					return nil, nil, fmt.Errorf("port range overlaps with another range")
				}

				if pr.HasOverlapWithPortsList(cpl) {
					return nil, nil, fmt.Errorf("port range overlaps with another port")
				}

				*cpr = append(*cpr, pr)
				cacheRanges[svc.Protocol] = cpr

				continue
			}

			// Single & Multiple ports
			pl, err := portutils.ConvertToPortsList(ports)
			if err != nil {
				return nil, nil, err
			}

			if pl.HasOverlapWithPortsList(cpl) {
				return nil, nil, fmt.Errorf("port overlaps with another port")
			}

			if pl.HasOverlapWithPortsRanges(cpr) {
				return nil, nil, fmt.Errorf("port overlaps with another port range")
			}

			*cpl = append(*cpl, *pl...)
			cachePortsList[svc.Protocol] = cpl
		}

	}

	return cachePortsList, cacheRanges, nil
}

// ValidateTimeDuration validates that the time duration provided is compliant
// with the go format.
func ValidateTimeDuration(attribute string, duration string) error {
	_, err := time.ParseDuration(duration)
	if err != nil {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be valid duration (examaple: 1h or 30s)", attribute))
	}
	return nil
}

// hostname regex from github.com/go-playground/validator
var hostnameRegexRFC1123 = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9\-\.]+[a-z-Az0-9]$`)

func isFQDN(val string) bool {

	if val == "" {
		return false
	}

	if val[len(val)-1] == '.' {
		val = val[0 : len(val)-1]
	}

	return hostnameRegexRFC1123.MatchString(val)
}

func ipNetFromString(ip string) (*net.IPNet, error) {
	_, ipNet, err := net.ParseCIDR(ip)
	if err != nil {
		parsedIP := net.ParseIP(ip)
		if parsedIP == nil {
			return nil, makeValidationError("IPs", "`IPs` must be a list of valid IPv4 address or CIDR notation")
		}
		ipNet = &net.IPNet{IP: parsedIP, Mask: []byte{0xf, 0xf, 0xf, 0xf}}
	}
	return ipNet, nil
}

// ValidateHTTPMethods validates the attribute methods is a list of HTTP verbs.
func ValidateHTTPMethods(attribute string, methods []string) error {

	for _, m := range methods {
		mu := strings.ToUpper(m)

		if mu != http.MethodPost &&
			mu != http.MethodGet &&
			mu != http.MethodDelete &&
			mu != http.MethodPut &&
			mu != http.MethodHead &&
			mu != http.MethodPatch {

			return makeValidationError(attribute, fmt.Sprintf("invalid HTTP method %s", m))
		}
	}

	return nil
}

// ValidateAutomation validates an automation
func ValidateAutomation(auto *Automation) error {
	switch auto.Trigger {
	case AutomationTriggerRemoteCall:
		if auto.WebhookMode {
			switch len(auto.Actions) {
			case 1:
			case 0:
				return makeValidationError("webhookMode", "An action must be defined if webhook mode is set to true.")
			default:
				return makeValidationError("webhookMode", "Only one action can be defined if webhook mode is set to true.")
			}
		}
	default:
		if auto.WebhookMode {
			return makeValidationError("webhookMode", fmt.Sprintf("webhook mode is only applicable if the automation trigger type is %s", AutomationTriggerRemoteCall))
		}
	}

	return nil
}

// ValidateHostServices validates a host service. Applies to the new API only.
func ValidateHostServices(hs *HostService) error {

	// Constraint on regex is used because the enforcer is using the name as nativeContextID.
	if !regHostServiceName.MatchString(hs.Name) {
		return makeValidationError("name", "Host service name must be less than 12 characters and contains only alphanumeric or _")
	}

	if !hs.HostModeEnabled && len(hs.Services) == 0 {
		return makeValidationError("services", "Host service must have either HostModeEnabled or must declare services")
	}

	if err := ValidateHostServicesNonOverlapPorts(hs.Services); err != nil {
		return makeValidationError("services", err.Error())
	}

	return nil
}

// ValidateProtoPorts validates a list of protocol/ports.
func ValidateProtoPorts(attribute string, services []string) error {

	for _, service := range services {
		if err := ValidateProtoPort(attribute, service); err != nil {
			return err
		}
	}

	return nil
}

// ValidateProtoPort validates protocol/port.
func ValidateProtoPort(attribute string, service string) error {

	portSubString, _, err := portutils.ExtractPortsAndProtocol(service)
	if err != nil {
		return makeValidationError(attribute, fmt.Sprintf("invalid format: %s", err))
	}

	_, err = portutils.ConvertToSinglePort(portSubString)
	if err != nil {
		return makeValidationError(attribute, fmt.Sprintf("invalid port: %s", err))
	}

	return nil
}

// ValidateHostServicesNonOverlapPorts validates a list of processing unit services has no overlap with any given parameter.
func ValidateHostServicesNonOverlapPorts(svcs []string) error {

	udpPorts := portutils.PortsRangeList{}
	tcpPorts := portutils.PortsRangeList{}

	var pr *portutils.PortsRange
	var protocol string
	var err error

	for _, svc := range svcs {

		pr, protocol, err = portutils.ExtractPortsAndProtocolFromHostService(svc)
		if err != nil {
			return err
		}

		switch protocol {
		case protocols.L4ProtocolTCP:
			if pr.HasOverlapWithPortsRanges(&tcpPorts) {
				return fmt.Errorf("host service cannot have overlapping TCP ports")
			}
			tcpPorts = append(tcpPorts, pr)
		case protocols.L4ProtocolUDP:
			if pr.HasOverlapWithPortsRanges(&udpPorts) {
				return fmt.Errorf("host service cannot have overlapping UDP ports")
			}
			udpPorts = append(udpPorts, pr)
		default:
			return fmt.Errorf("host service has invalid format: %s", protocol)
		}

	}

	return nil
}

// ValidateAudience validates an audience string.
func ValidateAudience(attribute string, audience string) error {
	// TODO: not liking the idea of importing addedeffect here
	return nil
}

// ValidatePEM validates a string contains a PEM.
func ValidatePEM(attribute string, pemdata string) error {

	if pemdata == "" {
		return nil
	}

	var i int
	var block *pem.Block
	rest := []byte(pemdata)

	for {
		block, rest = pem.Decode(rest)

		if block == nil {
			return makeValidationError(attribute, fmt.Sprintf("Unable to decode PEM number %d", i))
		}

		if len(rest) == 0 {
			return nil
		}
		i++
	}
}

// Constants to validate tags.
const (
	prefixDynamicTag  = "$"
	prefixExpandedTag = "+"
	prefixMetadata    = "@"
)

// validateTagStrings validates the given string and check if it can be a valid value for a Tag.
func validateTagStrings(attribute string, acceptReservedPrefix bool, strs ...string) error {

	for _, s := range strs {

		if !acceptReservedPrefix && (strings.HasPrefix(s, prefixMetadata) || strings.HasPrefix(s, prefixDynamicTag) || strings.HasPrefix(s, prefixExpandedTag)) {
			return makeValidationError(attribute, fmt.Sprintf("%s starts with an @, a $ or a + that is reserved", s))
		}

		if err := ValidateTag(attribute, s); err != nil {
			return err
		}
	}

	return nil
}

// tagRegex is the regular expression to check the format of a tag.
var tagRegex = regexp.MustCompile(`^[^= ]+=.+`)

// ValidateTag validates a single tag.
func ValidateTag(attribute string, tag string) error {

	if !tagRegex.MatchString(tag) {
		return makeValidationError(attribute, fmt.Sprintf("`%s must contain at least one '=' symbol separating two valid words", tag))
	}

	return nil
}

// ValidateTags validates a list of tags are valid. Accepts those with reserved prefix.
func ValidateTags(attribute string, tags []string) error {
	return validateTagStrings(attribute, true, tags...)
}

// ValidateTagsWithoutReservedPrefixes a list of tags are valid. Refuse those with reserved prefix.
func ValidateTagsWithoutReservedPrefixes(attribute string, tags []string) error {
	return validateTagStrings(attribute, false, tags...)
}

// ValidateTagsExpression validates an [][]string is a valid tag expression.
func ValidateTagsExpression(attribute string, expression [][]string) error {
	for _, tags := range expression {
		if err := ValidateTags(attribute, tags); err != nil {
			return err
		}
	}

	return nil
}

// ValidateMetadata validates an []string is a valid list of metadata.
func ValidateMetadata(attribute string, metadata []string) error {

	for _, m := range metadata {

		if !strings.HasPrefix(m, prefixMetadata) {
			return makeValidationError(attribute, fmt.Sprintf("Metadata %s does not starts with an @, a $ or a + that is reserved", m))
		}

		if strings.HasPrefix(m, constants.AuthKey) {
			return makeValidationError(attribute, fmt.Sprintf("Metadata %s is using @auth: which is reserved", m))
		}

		if err := ValidateTag(attribute, m); err != nil {
			return err
		}
	}

	return nil
}

// ValidateYAMLString validates the given data is a correct YAML string.
func ValidateYAMLString(attribute, data string) error {

	if err := yaml.Unmarshal([]byte(data), &map[string]interface{}{}); err != nil {
		return makeValidationError(attribute, fmt.Sprintf("not a valid yaml: %s", err))
	}
	return nil
}
