# Model
model:
  rest_name: issue
  resource_name: issue
  entity_name: Issue
  package: midgard
  group: core/authentication
  description: Issues a new Aporeto token according to given data.

# Attributes
attributes:
  v1:
  - name: audience
    description: |-
      If given, the issued token will only be valid for the specified namespace. 
      Refer to [JSON Web Token (JWT) RFC 7519](https://tools.ietf.org/html/rfc7519#section-4.1.3) 
      for further information. Example: `aud:*:*:/namespace`. 
    type: string
    exposed: true
    validations:
    - $audience

  - name: data
    description: Contains additional data. The value depends on the issuer type.
    type: string
    exposed: true
    deprecated: true
    orderable: true

  - name: metadata
    description: Contains various additional information. Meaning depends
      on the `realm`.
    type: external
    exposed: true
    subtype: map[string]interface{}
    example_value:
      vinceAccount: acme
      vinceOTP: 665435
      vincePassword: s3cr3t
    orderable: true

  - name: opaque
    description: Opaque data that will be included in the issued token.
    type: external
    exposed: true
    subtype: map[string]string

  - name: quota
    description: Restricts the number of times the issued token can be used.
    type: integer
    exposed: true

  - name: realm
    description: |-
      The authentication realm. `AWSIdentityDocument`, `AWSSecurityToken`, `Certificate`, 
      `Google`, `LDAP`, `Vince`, `GCPIdentityToken`, `AzureIdentityToken`, or `OIDC`.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - AWSIdentityDocument
    - AWSSecurityToken
    - Certificate
    - Google
    - LDAP
    - Vince
    - GCPIdentityToken
    - AzureIdentityToken
    - OIDC
    example_value: Vince

  - name: token
    description: The token to use for the registration.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: validity
    description: |-
      Configures the maximum length of validity for a token, using 
      [Golang duration syntax](https://golang.org/pkg/time/#example_Duration). If it is 
      bigger than the configured max validity, it will be capped. Default: `24h`.
    type: string
    exposed: true
    default_value: 24h
    orderable: true
    validations:
    - $timeDuration
