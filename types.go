package byo

import (
	"encoding/json"
	"fmt"
)

// CreateApiKeyErrorInvalidFieldsDetails represents a generated type
type CreateApiKeyErrorInvalidFieldsDetails struct {
	Fields []InvalidField `json:"fields"`
}


// CreateOidcClientErrorInvalidFieldsDetails represents a generated type
type CreateOidcClientErrorInvalidFieldsDetails struct {
	Fields []InvalidField `json:"fields"`
}


// CreateScimConnectionErrorInvalidFieldsDetails represents a generated type
type CreateScimConnectionErrorInvalidFieldsDetails struct {
	Fields []InvalidField `json:"fields"`
}


// CreateSessionErrorSessionLimitExceededDetails represents a generated type
type CreateSessionErrorSessionLimitExceededDetails struct {
	CurrentCount int `json:"currentCount"`
	MaxAllowed int `json:"maxAllowed"`
}


// DeviceInfo represents a generated type
type DeviceInfo struct {
	DisplayName string `json:"displayName"`
	DeviceType DeviceType `json:"deviceType"`
	Browser *string `json:"browser,omitempty"`
	BrowserVersion *string `json:"browserVersion,omitempty"`
	Os *string `json:"os,omitempty"`
	OsVersion *string `json:"osVersion,omitempty"`
}


// DeviceRegistration represents a generated type
type DeviceRegistration struct {
	SignedDeviceChallenge string `json:"signedDeviceChallenge"`
	RememberDevice *bool `json:"rememberDevice,omitempty"`
	RequestURL *string `json:"requestUrl,omitempty"`
	RequestMethod *string `json:"requestMethod,omitempty"`
}


// DeviceVerification represents a generated type
type DeviceVerification struct {
	SignedDeviceChallenge string `json:"signedDeviceChallenge"`
	RequestURL *string `json:"requestUrl,omitempty"`
	RequestMethod *string `json:"requestMethod,omitempty"`
}


// FinishPasskeyAuthenticationErrorOriginNotAllowedDetails represents a generated type
type FinishPasskeyAuthenticationErrorOriginNotAllowedDetails struct {
	Message string `json:"message"`
}


// FinishPasskeyRegistrationErrorOriginNotAllowedDetails represents a generated type
type FinishPasskeyRegistrationErrorOriginNotAllowedDetails struct {
	Message string `json:"message"`
}


// GetApiKeysErrorInvalidQueryFieldDetails represents a generated type
type GetApiKeysErrorInvalidQueryFieldDetails struct {
	Field string `json:"field"`
	Message string `json:"message"`
}


// GetScimUsersErrorInvalidQueryFieldDetails represents a generated type
type GetScimUsersErrorInvalidQueryFieldDetails struct {
	Field string `json:"field"`
	Message string `json:"message"`
}


// ImpersonationSessionInfo represents a generated type
type ImpersonationSessionInfo struct {
	ImpersonationSessionID string `json:"impersonationSessionId"`
	EmployeeEmail string `json:"employeeEmail"`
	TargetUserID string `json:"targetUserId"`
	Metadata *json.RawMessage `json:"metadata,omitempty"`
	CreatedAt int64 `json:"createdAt"`
	ExpiresAt int64 `json:"expiresAt"`
}


// InitiateOidcLoginErrorRedirectUrlInvalidDetails represents a generated type
type InitiateOidcLoginErrorRedirectUrlInvalidDetails struct {
	Message string `json:"message"`
}


// IntegrationKeyErrorCommandNotAllowedDetails represents a generated type
type IntegrationKeyErrorCommandNotAllowedDetails struct {
	CommandName string `json:"command_name"`
}


// InvalidField represents a generated type
type InvalidField struct {
	Field string `json:"field"`
	Code string `json:"code"`
	Message string `json:"message"`
}


// JwkKey represents a generated type
type JwkKey struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	Alg string `json:"alg"`
	N string `json:"n"`
	E string `json:"e"`
}


// OptionalIdpInfoFromCustomer is a discriminated union response type
type OptionalIdpInfoFromCustomer interface {
	isOptionalIdpInfoFromCustomer()
}

// OptionalIdpInfoFromCustomerMicrosoftEntra represents the MicrosoftEntra variant
type OptionalIdpInfoFromCustomerMicrosoftEntra struct {
	ClientSecret *string `json:"clientSecret,omitempty"`
	UsesPkce *bool `json:"usesPkce,omitempty"`
	TenantID *string `json:"tenantId,omitempty"`
}

func (r *OptionalIdpInfoFromCustomerMicrosoftEntra) isOptionalIdpInfoFromCustomer() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *OptionalIdpInfoFromCustomerMicrosoftEntra) MarshalJSON() ([]byte, error) {
	type Alias OptionalIdpInfoFromCustomerMicrosoftEntra
	return json.Marshal(&struct {
		IdpType string `json:"idpType"`
		*Alias
	}{
		IdpType: "MicrosoftEntra",
		Alias: (*Alias)(v),
	})
}

// OptionalIdpInfoFromCustomerOkta represents the Okta variant
type OptionalIdpInfoFromCustomerOkta struct {
	ClientSecret *string `json:"clientSecret,omitempty"`
	UsesPkce *bool `json:"usesPkce,omitempty"`
	SsoDomain *string `json:"ssoDomain,omitempty"`
}

func (r *OptionalIdpInfoFromCustomerOkta) isOptionalIdpInfoFromCustomer() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *OptionalIdpInfoFromCustomerOkta) MarshalJSON() ([]byte, error) {
	type Alias OptionalIdpInfoFromCustomerOkta
	return json.Marshal(&struct {
		IdpType string `json:"idpType"`
		*Alias
	}{
		IdpType: "Okta",
		Alias: (*Alias)(v),
	})
}

// OptionalIdpInfoFromCustomerGeneric represents the Generic variant
type OptionalIdpInfoFromCustomerGeneric struct {
	ClientSecret *string `json:"clientSecret,omitempty"`
	UsesPkce *bool `json:"usesPkce,omitempty"`
	AuthURL *string `json:"authUrl,omitempty"`
	TokenURL *string `json:"tokenUrl,omitempty"`
	UserinfoURL *string `json:"userinfoUrl,omitempty"`
}

func (r *OptionalIdpInfoFromCustomerGeneric) isOptionalIdpInfoFromCustomer() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *OptionalIdpInfoFromCustomerGeneric) MarshalJSON() ([]byte, error) {
	type Alias OptionalIdpInfoFromCustomerGeneric
	return json.Marshal(&struct {
		IdpType string `json:"idpType"`
		*Alias
	}{
		IdpType: "Generic",
		Alias: (*Alias)(v),
	})
}



// UnmarshalOptionalIdpInfoFromCustomer unmarshals JSON data into the appropriate OptionalIdpInfoFromCustomer variant
func UnmarshalOptionalIdpInfoFromCustomer(data []byte) (OptionalIdpInfoFromCustomer, error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Single discriminator
	var discriminator string
	if err := json.Unmarshal(raw["idpType"], &discriminator); err != nil {
		return nil, fmt.Errorf("parse idpType discriminator: %w", err)
	}

	switch discriminator {
	case "MicrosoftEntra":
		var variant OptionalIdpInfoFromCustomerMicrosoftEntra
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "Okta":
		var variant OptionalIdpInfoFromCustomerOkta
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "Generic":
		var variant OptionalIdpInfoFromCustomerGeneric
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	default:
		return nil, fmt.Errorf("unknown idpType: %s", discriminator)
	}
}


// OptionalIdpInfoFromCustomerForEntra represents a generated type
type OptionalIdpInfoFromCustomerForEntra struct {
	ClientSecret *string `json:"clientSecret,omitempty"`
	UsesPkce *bool `json:"usesPkce,omitempty"`
	TenantID *string `json:"tenantId,omitempty"`
}


// OptionalIdpInfoFromCustomerForGeneric represents a generated type
type OptionalIdpInfoFromCustomerForGeneric struct {
	ClientSecret *string `json:"clientSecret,omitempty"`
	UsesPkce *bool `json:"usesPkce,omitempty"`
	AuthURL *string `json:"authUrl,omitempty"`
	TokenURL *string `json:"tokenUrl,omitempty"`
	UserinfoURL *string `json:"userinfoUrl,omitempty"`
}


// OptionalIdpInfoFromCustomerForOkta represents a generated type
type OptionalIdpInfoFromCustomerForOkta struct {
	ClientSecret *string `json:"clientSecret,omitempty"`
	UsesPkce *bool `json:"usesPkce,omitempty"`
	SsoDomain *string `json:"ssoDomain,omitempty"`
}


// PasskeyInfo represents a generated type
type PasskeyInfo struct {
	DisplayName *string `json:"displayName,omitempty"`
	CredentialID string `json:"credentialId"`
}


// PatchApiKeyErrorInvalidFieldsDetails represents a generated type
type PatchApiKeyErrorInvalidFieldsDetails struct {
	Fields []InvalidField `json:"fields"`
}


// PatchOidcClientErrorInvalidFieldsDetails represents a generated type
type PatchOidcClientErrorInvalidFieldsDetails struct {
	Fields []InvalidField `json:"fields"`
}


// PatchScimConnectionErrorDisplayNameInvalidDetails represents a generated type
type PatchScimConnectionErrorDisplayNameInvalidDetails struct {
	Message string `json:"message"`
}


// RegisterDeviceErrorNewDeviceChallengeRequiredDetails represents a generated type
type RegisterDeviceErrorNewDeviceChallengeRequiredDetails struct {
	DeviceChallenge string `json:"deviceChallenge"`
	ExpiresAt int64 `json:"expiresAt"`
}


// ScimActionHook is a discriminated union response type
type ScimActionHook interface {
	isScimActionHook()
}

// ScimActionHookLinkUser represents the LinkUser variant
type ScimActionHookLinkUser struct {
	CommitID string `json:"commitId"`
	PrimaryEmail *string `json:"primaryEmail,omitempty"`
	UserName string `json:"userName"`
	ParsedUserData json.RawMessage `json:"parsedUserData"`
	Active bool `json:"active"`
	SsoUserSubject *string `json:"ssoUserSubject,omitempty"`
}

func (r *ScimActionHookLinkUser) isScimActionHook() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *ScimActionHookLinkUser) MarshalJSON() ([]byte, error) {
	type Alias ScimActionHookLinkUser
	return json.Marshal(&struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "LinkUser",
		Alias: (*Alias)(v),
	})
}

// ScimActionHookDisableUser represents the DisableUser variant
type ScimActionHookDisableUser struct {
	CommitID string `json:"commitId"`
	UserID string `json:"userId"`
	PrimaryEmail *string `json:"primaryEmail,omitempty"`
	ParsedUserData json.RawMessage `json:"parsedUserData"`
}

func (r *ScimActionHookDisableUser) isScimActionHook() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *ScimActionHookDisableUser) MarshalJSON() ([]byte, error) {
	type Alias ScimActionHookDisableUser
	return json.Marshal(&struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "DisableUser",
		Alias: (*Alias)(v),
	})
}

// ScimActionHookEnableUser represents the EnableUser variant
type ScimActionHookEnableUser struct {
	CommitID string `json:"commitId"`
	UserID string `json:"userId"`
	PrimaryEmail *string `json:"primaryEmail,omitempty"`
	ParsedUserData json.RawMessage `json:"parsedUserData"`
}

func (r *ScimActionHookEnableUser) isScimActionHook() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *ScimActionHookEnableUser) MarshalJSON() ([]byte, error) {
	type Alias ScimActionHookEnableUser
	return json.Marshal(&struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "EnableUser",
		Alias: (*Alias)(v),
	})
}

// ScimActionHookDeleteUser represents the DeleteUser variant
type ScimActionHookDeleteUser struct {
	CommitID string `json:"commitId"`
	UserID string `json:"userId"`
	PrimaryEmail *string `json:"primaryEmail,omitempty"`
	ParsedUserData json.RawMessage `json:"parsedUserData"`
}

func (r *ScimActionHookDeleteUser) isScimActionHook() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *ScimActionHookDeleteUser) MarshalJSON() ([]byte, error) {
	type Alias ScimActionHookDeleteUser
	return json.Marshal(&struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "DeleteUser",
		Alias: (*Alias)(v),
	})
}



// UnmarshalScimActionHook unmarshals JSON data into the appropriate ScimActionHook variant
func UnmarshalScimActionHook(data []byte) (ScimActionHook, error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Single discriminator
	var discriminator string
	if err := json.Unmarshal(raw["action"], &discriminator); err != nil {
		return nil, fmt.Errorf("parse action discriminator: %w", err)
	}

	switch discriminator {
	case "LinkUser":
		var variant ScimActionHookLinkUser
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "DisableUser":
		var variant ScimActionHookDisableUser
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "EnableUser":
		var variant ScimActionHookEnableUser
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "DeleteUser":
		var variant ScimActionHookDeleteUser
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	default:
		return nil, fmt.Errorf("unknown action: %s", discriminator)
	}
}


// ScimUnderlyingErrorMissingRequiredFieldDetails represents a generated type
type ScimUnderlyingErrorMissingRequiredFieldDetails struct {
	Field string `json:"field"`
}


// ScimUnderlyingErrorCantRemoveRequiredFieldDetails represents a generated type
type ScimUnderlyingErrorCantRemoveRequiredFieldDetails struct {
	Field string `json:"field"`
}


// ScimUnderlyingErrorScimUserIdAlreadyTakenDetails represents a generated type
type ScimUnderlyingErrorScimUserIdAlreadyTakenDetails struct {
	Field string `json:"field"`
}


// ScimUnderlyingErrorInvalidQueryFieldDetails represents a generated type
type ScimUnderlyingErrorInvalidQueryFieldDetails struct {
	Field string `json:"field"`
	Message string `json:"message"`
}


// ScimUnderlyingErrorInvalidPatchValueDetails represents a generated type
type ScimUnderlyingErrorInvalidPatchValueDetails struct {
	Field string `json:"field"`
}


// ScimUsersPageEqualityFilter is a discriminated union response type
type ScimUsersPageEqualityFilter interface {
	isScimUsersPageEqualityFilter()
}

// ScimUsersPageEqualityFilterUserName represents the userName variant
type ScimUsersPageEqualityFilterUserName struct {
	UserName string `json:"userName"`
}

func (r *ScimUsersPageEqualityFilterUserName) isScimUsersPageEqualityFilter() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *ScimUsersPageEqualityFilterUserName) MarshalJSON() ([]byte, error) {
	type Alias ScimUsersPageEqualityFilterUserName
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "userName",
		Alias: (*Alias)(v),
	})
}

// ScimUsersPageEqualityFilterExternalID represents the externalId variant
type ScimUsersPageEqualityFilterExternalID struct {
	ExternalID string `json:"externalId"`
}

func (r *ScimUsersPageEqualityFilterExternalID) isScimUsersPageEqualityFilter() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *ScimUsersPageEqualityFilterExternalID) MarshalJSON() ([]byte, error) {
	type Alias ScimUsersPageEqualityFilterExternalID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "externalId",
		Alias: (*Alias)(v),
	})
}

// ScimUsersPageEqualityFilterPrimaryEmail represents the primaryEmail variant
type ScimUsersPageEqualityFilterPrimaryEmail struct {
	PrimaryEmail string `json:"primaryEmail"`
}

func (r *ScimUsersPageEqualityFilterPrimaryEmail) isScimUsersPageEqualityFilter() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *ScimUsersPageEqualityFilterPrimaryEmail) MarshalJSON() ([]byte, error) {
	type Alias ScimUsersPageEqualityFilterPrimaryEmail
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "primaryEmail",
		Alias: (*Alias)(v),
	})
}

// ScimUsersPageEqualityFilterUserID represents the userId variant
type ScimUsersPageEqualityFilterUserID struct {
	UserID string `json:"userId"`
}

func (r *ScimUsersPageEqualityFilterUserID) isScimUsersPageEqualityFilter() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *ScimUsersPageEqualityFilterUserID) MarshalJSON() ([]byte, error) {
	type Alias ScimUsersPageEqualityFilterUserID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "userId",
		Alias: (*Alias)(v),
	})
}



// UnmarshalScimUsersPageEqualityFilter unmarshals JSON data into the appropriate ScimUsersPageEqualityFilter variant
func UnmarshalScimUsersPageEqualityFilter(data []byte) (ScimUsersPageEqualityFilter, error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Single discriminator
	var discriminator string
	if err := json.Unmarshal(raw["__fieldPresence"], &discriminator); err != nil {
		return nil, fmt.Errorf("parse __fieldPresence discriminator: %w", err)
	}

	switch discriminator {
	case "userName":
		var variant ScimUsersPageEqualityFilterUserName
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "externalId":
		var variant ScimUsersPageEqualityFilterExternalID
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "primaryEmail":
		var variant ScimUsersPageEqualityFilterPrimaryEmail
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "userId":
		var variant ScimUsersPageEqualityFilterUserID
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	default:
		return nil, fmt.Errorf("unknown __fieldPresence: %s", discriminator)
	}
}


// SessionInfo represents a generated type
type SessionInfo struct {
	SessionID string `json:"sessionId"`
	CreatedAt int64 `json:"createdAt"`
	ExpiresAt int64 `json:"expiresAt"`
	LastActivityAt int64 `json:"lastActivityAt"`
	Device DeviceInfo `json:"device"`
	IPAddress *string `json:"ipAddress,omitempty"`
	SessionTags []string `json:"sessionTags,omitempty"`
	Metadata *json.RawMessage `json:"metadata,omitempty"`
}


// SessionTag represents a generated type
type SessionTag struct {
	TagName string `json:"tagName"`
	TagValue string `json:"tagValue"`
}


// SessionsFilter represents a generated type
type SessionsFilter struct {
	UserID *string `json:"userId,omitempty"`
	SessionTags []string `json:"sessionTags,omitempty"`
}


// StartPasskeyRegistrationErrorTooManyPasskeysDetails represents a generated type
type StartPasskeyRegistrationErrorTooManyPasskeysDetails struct {
	MaxPasskeys int `json:"max_passkeys"`
	CurrentCount int `json:"current_count"`
}


// UpdateSessionErrorCannotModifyOnCreateOnlyTagsDetails represents a generated type
type UpdateSessionErrorCannotModifyOnCreateOnlyTagsDetails struct {
	TagNames []string `json:"tag_names"`
}


// UpdateSessionsErrorCannotModifyOnCreateOnlyTagsDetails represents a generated type
type UpdateSessionsErrorCannotModifyOnCreateOnlyTagsDetails struct {
	TagNames []string `json:"tag_names"`
}


// ValidateAndRefreshSessionErrorNewDeviceChallengeRequiredDetails represents a generated type
type ValidateAndRefreshSessionErrorNewDeviceChallengeRequiredDetails struct {
	DeviceChallenge string `json:"deviceChallenge"`
	ExpiresAt int64 `json:"expiresAt"`
}


// ValidateSessionErrorNewDeviceChallengeRequiredDetails represents a generated type
type ValidateSessionErrorNewDeviceChallengeRequiredDetails struct {
	DeviceChallenge string `json:"deviceChallenge"`
	ExpiresAt int64 `json:"expiresAt"`
}


// JsonValue represents arbitrary JSON values (maps, arrays, primitives, etc.)
type JsonValue = any


// ApiKeyStatus represents an enumeration type
type ApiKeyStatus string

const (
	ApiKeyStatusLive ApiKeyStatus = "Live"
	ApiKeyStatusRevoked ApiKeyStatus = "Revoked"
)


// DeviceType represents an enumeration type
type DeviceType string

const (
	DeviceTypeDesktop DeviceType = "Desktop"
	DeviceTypeMobile DeviceType = "Mobile"
	DeviceTypeTablet DeviceType = "Tablet"
	DeviceTypeUnknown DeviceType = "Unknown"
)


// IdentityProvider is a discriminated union response type
type IdentityProvider interface {
	isIdentityProvider()
}

// IdentityProviderMicrosoftEntra represents the MicrosoftEntra variant
type IdentityProviderMicrosoftEntra struct {
	TenantID string `json:"tenantId"`
}

func (r *IdentityProviderMicrosoftEntra) isIdentityProvider() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *IdentityProviderMicrosoftEntra) MarshalJSON() ([]byte, error) {
	type Alias IdentityProviderMicrosoftEntra
	return json.Marshal(&struct {
		IdpType string `json:"idpType"`
		*Alias
	}{
		IdpType: "MicrosoftEntra",
		Alias: (*Alias)(v),
	})
}

// IdentityProviderOkta represents the Okta variant
type IdentityProviderOkta struct {
	SsoDomain string `json:"ssoDomain"`
}

func (r *IdentityProviderOkta) isIdentityProvider() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *IdentityProviderOkta) MarshalJSON() ([]byte, error) {
	type Alias IdentityProviderOkta
	return json.Marshal(&struct {
		IdpType string `json:"idpType"`
		*Alias
	}{
		IdpType: "Okta",
		Alias: (*Alias)(v),
	})
}

// IdentityProviderGeneric represents the Generic variant
type IdentityProviderGeneric struct {
	AuthURL string `json:"authUrl"`
	TokenURL string `json:"tokenUrl"`
	UserinfoURL string `json:"userinfoUrl"`
}

func (r *IdentityProviderGeneric) isIdentityProvider() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *IdentityProviderGeneric) MarshalJSON() ([]byte, error) {
	type Alias IdentityProviderGeneric
	return json.Marshal(&struct {
		IdpType string `json:"idpType"`
		*Alias
	}{
		IdpType: "Generic",
		Alias: (*Alias)(v),
	})
}



// UnmarshalIdentityProvider unmarshals JSON data into the appropriate IdentityProvider variant
func UnmarshalIdentityProvider(data []byte) (IdentityProvider, error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Single discriminator
	var discriminator string
	if err := json.Unmarshal(raw["idpType"], &discriminator); err != nil {
		return nil, fmt.Errorf("parse idpType discriminator: %w", err)
	}

	switch discriminator {
	case "MicrosoftEntra":
		var variant IdentityProviderMicrosoftEntra
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "Okta":
		var variant IdentityProviderOkta
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "Generic":
		var variant IdentityProviderGeneric
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	default:
		return nil, fmt.Errorf("unknown idpType: %s", discriminator)
	}
}


// IdpInfoFromCustomer is a discriminated union response type
type IdpInfoFromCustomer interface {
	isIdpInfoFromCustomer()
}

// IdpInfoFromCustomerMicrosoftEntra represents the MicrosoftEntra variant
type IdpInfoFromCustomerMicrosoftEntra struct {
	ClientID string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	UsesPkce bool `json:"usesPkce"`
	TenantID string `json:"tenantId"`
}

func (r *IdpInfoFromCustomerMicrosoftEntra) isIdpInfoFromCustomer() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *IdpInfoFromCustomerMicrosoftEntra) MarshalJSON() ([]byte, error) {
	type Alias IdpInfoFromCustomerMicrosoftEntra
	return json.Marshal(&struct {
		IdpType string `json:"idpType"`
		*Alias
	}{
		IdpType: "MicrosoftEntra",
		Alias: (*Alias)(v),
	})
}

// IdpInfoFromCustomerOkta represents the Okta variant
type IdpInfoFromCustomerOkta struct {
	ClientID string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	UsesPkce bool `json:"usesPkce"`
	SsoDomain string `json:"ssoDomain"`
}

func (r *IdpInfoFromCustomerOkta) isIdpInfoFromCustomer() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *IdpInfoFromCustomerOkta) MarshalJSON() ([]byte, error) {
	type Alias IdpInfoFromCustomerOkta
	return json.Marshal(&struct {
		IdpType string `json:"idpType"`
		*Alias
	}{
		IdpType: "Okta",
		Alias: (*Alias)(v),
	})
}

// IdpInfoFromCustomerGeneric represents the Generic variant
type IdpInfoFromCustomerGeneric struct {
	ClientID string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	UsesPkce bool `json:"usesPkce"`
	AuthURL string `json:"authUrl"`
	TokenURL string `json:"tokenUrl"`
	UserinfoURL string `json:"userinfoUrl"`
}

func (r *IdpInfoFromCustomerGeneric) isIdpInfoFromCustomer() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *IdpInfoFromCustomerGeneric) MarshalJSON() ([]byte, error) {
	type Alias IdpInfoFromCustomerGeneric
	return json.Marshal(&struct {
		IdpType string `json:"idpType"`
		*Alias
	}{
		IdpType: "Generic",
		Alias: (*Alias)(v),
	})
}



// UnmarshalIdpInfoFromCustomer unmarshals JSON data into the appropriate IdpInfoFromCustomer variant
func UnmarshalIdpInfoFromCustomer(data []byte) (IdpInfoFromCustomer, error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Single discriminator
	var discriminator string
	if err := json.Unmarshal(raw["idpType"], &discriminator); err != nil {
		return nil, fmt.Errorf("parse idpType discriminator: %w", err)
	}

	switch discriminator {
	case "MicrosoftEntra":
		var variant IdpInfoFromCustomerMicrosoftEntra
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "Okta":
		var variant IdpInfoFromCustomerOkta
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "Generic":
		var variant IdpInfoFromCustomerGeneric
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	default:
		return nil, fmt.Errorf("unknown idpType: %s", discriminator)
	}
}


// MappingErrors represents a generated type
type MappingErrors struct {
	Errors []PropertyParseError `json:"errors"`
}


// OidcClientIdentifier is a discriminated union response type
type OidcClientIdentifier interface {
	isOidcClientIdentifier()
}

// OidcClientIdentifierOIDCClientID represents the oidcClientId variant
type OidcClientIdentifierOIDCClientID struct {
	OIDCClientID string `json:"oidcClientId"`
}

func (r *OidcClientIdentifierOIDCClientID) isOidcClientIdentifier() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *OidcClientIdentifierOIDCClientID) MarshalJSON() ([]byte, error) {
	type Alias OidcClientIdentifierOIDCClientID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "oidcClientId",
		Alias: (*Alias)(v),
	})
}

// OidcClientIdentifierCustomerID represents the customerId variant
type OidcClientIdentifierCustomerID struct {
	CustomerID string `json:"customerId"`
}

func (r *OidcClientIdentifierCustomerID) isOidcClientIdentifier() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *OidcClientIdentifierCustomerID) MarshalJSON() ([]byte, error) {
	type Alias OidcClientIdentifierCustomerID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "customerId",
		Alias: (*Alias)(v),
	})
}



// UnmarshalOidcClientIdentifier unmarshals JSON data into the appropriate OidcClientIdentifier variant
func UnmarshalOidcClientIdentifier(data []byte) (OidcClientIdentifier, error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Single discriminator
	var discriminator string
	if err := json.Unmarshal(raw["__fieldPresence"], &discriminator); err != nil {
		return nil, fmt.Errorf("parse __fieldPresence discriminator: %w", err)
	}

	switch discriminator {
	case "oidcClientId":
		var variant OidcClientIdentifierOIDCClientID
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "customerId":
		var variant OidcClientIdentifierCustomerID
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	default:
		return nil, fmt.Errorf("unknown __fieldPresence: %s", discriminator)
	}
}


// PropertyType is a discriminated union response type
type PropertyType interface {
	isPropertyType()
}

// PropertyTypeString represents the String variant
type PropertyTypeString struct {
}

func (r *PropertyTypeString) isPropertyType() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *PropertyTypeString) MarshalJSON() ([]byte, error) {
	type Alias PropertyTypeString
	return json.Marshal(&struct {
		DataType string `json:"dataType"`
		*Alias
	}{
		DataType: "String",
		Alias: (*Alias)(v),
	})
}

// PropertyTypeInteger represents the Integer variant
type PropertyTypeInteger struct {
}

func (r *PropertyTypeInteger) isPropertyType() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *PropertyTypeInteger) MarshalJSON() ([]byte, error) {
	type Alias PropertyTypeInteger
	return json.Marshal(&struct {
		DataType string `json:"dataType"`
		*Alias
	}{
		DataType: "Integer",
		Alias: (*Alias)(v),
	})
}

// PropertyTypeFloat represents the Float variant
type PropertyTypeFloat struct {
}

func (r *PropertyTypeFloat) isPropertyType() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *PropertyTypeFloat) MarshalJSON() ([]byte, error) {
	type Alias PropertyTypeFloat
	return json.Marshal(&struct {
		DataType string `json:"dataType"`
		*Alias
	}{
		DataType: "Float",
		Alias: (*Alias)(v),
	})
}

// PropertyTypeBoolean represents the Boolean variant
type PropertyTypeBoolean struct {
}

func (r *PropertyTypeBoolean) isPropertyType() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *PropertyTypeBoolean) MarshalJSON() ([]byte, error) {
	type Alias PropertyTypeBoolean
	return json.Marshal(&struct {
		DataType string `json:"dataType"`
		*Alias
	}{
		DataType: "Boolean",
		Alias: (*Alias)(v),
	})
}

// PropertyTypeDate represents the Date variant
type PropertyTypeDate struct {
}

func (r *PropertyTypeDate) isPropertyType() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *PropertyTypeDate) MarshalJSON() ([]byte, error) {
	type Alias PropertyTypeDate
	return json.Marshal(&struct {
		DataType string `json:"dataType"`
		*Alias
	}{
		DataType: "Date",
		Alias: (*Alias)(v),
	})
}

// PropertyTypeDateTime represents the DateTime variant
type PropertyTypeDateTime struct {
}

func (r *PropertyTypeDateTime) isPropertyType() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *PropertyTypeDateTime) MarshalJSON() ([]byte, error) {
	type Alias PropertyTypeDateTime
	return json.Marshal(&struct {
		DataType string `json:"dataType"`
		*Alias
	}{
		DataType: "DateTime",
		Alias: (*Alias)(v),
	})
}

// PropertyTypeEnum represents the Enum variant
type PropertyTypeEnum struct {
	Options []string `json:"options"`
}

func (r *PropertyTypeEnum) isPropertyType() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *PropertyTypeEnum) MarshalJSON() ([]byte, error) {
	type Alias PropertyTypeEnum
	return json.Marshal(&struct {
		DataType string `json:"dataType"`
		*Alias
	}{
		DataType: "Enum",
		Alias: (*Alias)(v),
	})
}

// PropertyTypeList represents the List variant
type PropertyTypeList struct {
	ItemType PropertyType `json:"itemType"`
}

func (r *PropertyTypeList) isPropertyType() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *PropertyTypeList) MarshalJSON() ([]byte, error) {
	type Alias PropertyTypeList
	return json.Marshal(&struct {
		DataType string `json:"dataType"`
		*Alias
	}{
		DataType: "List",
		Alias: (*Alias)(v),
	})
}



// UnmarshalPropertyType unmarshals JSON data into the appropriate PropertyType variant
func UnmarshalPropertyType(data []byte) (PropertyType, error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Single discriminator
	var discriminator string
	if err := json.Unmarshal(raw["dataType"], &discriminator); err != nil {
		return nil, fmt.Errorf("parse dataType discriminator: %w", err)
	}

	switch discriminator {
	case "String":
		var variant PropertyTypeString
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "Integer":
		var variant PropertyTypeInteger
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "Float":
		var variant PropertyTypeFloat
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "Boolean":
		var variant PropertyTypeBoolean
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "Date":
		var variant PropertyTypeDate
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "DateTime":
		var variant PropertyTypeDateTime
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "Enum":
		var variant PropertyTypeEnum
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "List":
		var variant PropertyTypeList
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	default:
		return nil, fmt.Errorf("unknown dataType: %s", discriminator)
	}
}


// ScimMatchingDefinition represents a generated type
type ScimMatchingDefinition struct {
	Strategy ScimUserMatchingStrategy `json:"strategy"`
}


// ScimUserMappingConfig represents a generated type
type ScimUserMappingConfig struct {
	UserSchema []ScimUserMappingFieldDefinition `json:"userSchema"`
}


// ScimUserMappingFieldDefinition represents a generated type
type ScimUserMappingFieldDefinition struct {
	OutputField string `json:"outputField"`
	InputPath string `json:"inputPath"`
	FallbackInputPaths []string `json:"fallbackInputPaths"`
	PropertyType PropertyType `json:"propertyType"`
	DisplayName string `json:"displayName"`
	Description *string `json:"description,omitempty"`
	WarnIfMissing bool `json:"warnIfMissing"`
	DefaultValue *json.RawMessage `json:"defaultValue,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler to decode discriminated-union fields
// whose static types are interfaces. The alias-shadow pattern lets standard
// fields decode normally via the embedded *Alias while routing union fields
// through their dedicated Unmarshal<TypeName> helpers.
func (r *ScimUserMappingFieldDefinition) UnmarshalJSON(data []byte) error {
	type Alias ScimUserMappingFieldDefinition
	aux := &struct {
		PropertyType json.RawMessage `json:"propertyType"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	if len(aux.PropertyType) > 0 && string(aux.PropertyType) != "null" {
		val, err := UnmarshalPropertyType(aux.PropertyType)
		if err != nil {
			return fmt.Errorf("unmarshal propertyType: %w", err)
		}
		r.PropertyType = val
	}
	return nil
}


// ScimUserMatchingStrategy represents an enumeration type
type ScimUserMatchingStrategy string

const (
	ScimUserMatchingStrategyOIDCSubToScimUsername ScimUserMatchingStrategy = "OidcSubToScimUsername"
	ScimUserMatchingStrategyOIDCSubToScimExternalID ScimUserMatchingStrategy = "OidcSubToScimExternalId"
	ScimUserMatchingStrategyOIDCEmailToScimUsername ScimUserMatchingStrategy = "OidcEmailToScimUsername"
	ScimUserMatchingStrategyOIDCEmailUsernameToScimUsername ScimUserMatchingStrategy = "OidcEmailUsernameToScimUsername"
	ScimUserMatchingStrategyOIDCPreferredUsernameToScimUsername ScimUserMatchingStrategy = "OidcPreferredUsernameToScimUsername"
)


// ScimUserWithWarnings represents a generated type
type ScimUserWithWarnings struct {
	ExternalUserID string `json:"externalUserId"`
	UserName string `json:"userName"`
	MissingFields []string `json:"missingFields"`
}


// ScimUsersWithWarningsPage represents a generated type
type ScimUsersWithWarningsPage struct {
	Users []ScimUserWithWarnings `json:"users"`
	TotalCount int `json:"totalCount"`
	Page int `json:"page"`
	PageSize int `json:"pageSize"`
}
