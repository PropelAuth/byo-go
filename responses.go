package byo

import (
	"encoding/json"
	"fmt"
)

// CompleteOidcLoginResponse represents a generated type
type CompleteOidcLoginResponse struct {
	ClientID string `json:"clientId"`
	CustomerID string `json:"customerId"`
	OIDCUserID string `json:"oidcUserId"`
	Email *string `json:"email,omitempty"`
	EmailVerified bool `json:"emailVerified"`
	PreferredUsername *string `json:"preferredUsername,omitempty"`
	DataFromSso json.RawMessage `json:"dataFromSso"`
	ScimUser *CompleteScimUserResponse `json:"scimUser,omitempty"`
	PostLoginRedirectURL *string `json:"postLoginRedirectUrl,omitempty"`
}


// CompleteScimUserResponse represents a generated type
type CompleteScimUserResponse struct {
	ConnectionID string `json:"connectionId"`
	ScimUser json.RawMessage `json:"scimUser"`
	PrimaryEmail *string `json:"primaryEmail,omitempty"`
	ParsedUserData json.RawMessage `json:"parsedUserData"`
	Active bool `json:"active"`
	UserID *string `json:"userId,omitempty"`
}


// CompletedScimRequestResponse represents a generated type
type CompletedScimRequestResponse struct {
	ConnectionID string `json:"connectionId"`
	ResponseData *json.RawMessage `json:"responseData,omitempty"`
	ResponseHTTPCode int `json:"responseHttpCode"`
	AffectedUserIds []string `json:"affectedUserIds,omitempty"`
}


// CreateDeviceChallengeResponse represents a generated type
type CreateDeviceChallengeResponse struct {
	DeviceChallenge string `json:"deviceChallenge"`
	ExpiresAt int64 `json:"expiresAt"`
}


// CreateImpersonationSessionResponse represents a generated type
type CreateImpersonationSessionResponse struct {
	SessionID string `json:"sessionId"`
	ImpersonationSessionToken string `json:"impersonationSessionToken"`
	ExpiresAt int64 `json:"expiresAt"`
}


// CreateOidcClientResponse represents a generated type
type CreateOidcClientResponse struct {
	ClientID string `json:"clientId"`
}


// CreateScimConnectionResponse represents a generated type
type CreateScimConnectionResponse struct {
	ConnectionID string `json:"connectionId"`
	ScimAPIKey string `json:"scimApiKey"`
}


// CreateSessionResponse represents a generated type
type CreateSessionResponse struct {
	SessionID string `json:"sessionId"`
	SessionToken string `json:"sessionToken"`
	ExpiresAt int64 `json:"expiresAt"`
	NewDeviceDetected *bool `json:"newDeviceDetected,omitempty"`
}


// CreateStatelessTokenResponse represents a generated type
type CreateStatelessTokenResponse struct {
	StatelessToken string `json:"statelessToken"`
	ExpiresAt int64 `json:"expiresAt"`
}


// DeleteOidcClientResponse represents a generated type
type DeleteOidcClientResponse struct {
}


// DeleteScimConnectionResponse represents a generated type
type DeleteScimConnectionResponse struct {
}


// DeregisterAllPasskeysForUserResponse represents a generated type
type DeregisterAllPasskeysForUserResponse struct {
}


// DeregisterPasskeyResponse represents a generated type
type DeregisterPasskeyResponse struct {
}


// FetchAllActiveImpersonationSessionsResponse represents a generated type
type FetchAllActiveImpersonationSessionsResponse struct {
	Sessions []ImpersonationSessionInfo `json:"sessions"`
	NextPagingToken *string `json:"nextPagingToken,omitempty"`
	HasMoreResults bool `json:"hasMoreResults"`
}


// FetchAllImpersonationSessionsForEmployeeResponse represents a generated type
type FetchAllImpersonationSessionsForEmployeeResponse struct {
	Sessions []ImpersonationSessionInfo `json:"sessions"`
}


// FetchAllImpersonationSessionsForUserResponse represents a generated type
type FetchAllImpersonationSessionsForUserResponse struct {
	Sessions []ImpersonationSessionInfo `json:"sessions"`
}


// FetchAllPasskeysForUserResponse represents a generated type
type FetchAllPasskeysForUserResponse struct {
	Passkeys []PasskeyInfo `json:"passkeys"`
}


// FetchAllSessionsForUserResponse represents a generated type
type FetchAllSessionsForUserResponse struct {
	Sessions []SessionInfo `json:"sessions"`
}


// FetchAllSessionsResponse represents a generated type
type FetchAllSessionsResponse struct {
	Items []SessionInfo `json:"items"`
	Page int `json:"page"`
	PageSize int `json:"pageSize"`
	TotalCount int `json:"totalCount"`
	HasMoreResults bool `json:"hasMoreResults"`
}


// FetchOidcClientResponse represents a generated type
type FetchOidcClientResponse struct {
	IdpInfoFromCustomer IdpInfoFromCustomerResponse `json:"idpInfoFromCustomer"`
	CustomerID string `json:"customerId"`
	DisplayName *string `json:"displayName,omitempty"`
	RedirectURL string `json:"redirectUrl"`
	EmailDomainAllowlist []string `json:"emailDomainAllowlist"`
	ScimConnection *FetchScimConnectionResponse `json:"scimConnection,omitempty"`
	AdditionalScopes []string `json:"additionalScopes"`
	ScimMatchingDefinition *ScimMatchingDefinition `json:"scimMatchingDefinition,omitempty"`
}


// FetchScimConnectionResponse represents a generated type
type FetchScimConnectionResponse struct {
	ConnectionID string `json:"connectionId"`
	CustomerID string `json:"customerId"`
	DisplayName *string `json:"displayName,omitempty"`
	ScimAPIKeyValidUntil *int `json:"scimApiKeyValidUntil,omitempty"`
	UserMapping ScimUserMappingConfig `json:"userMapping"`
}


// FinishPasskeyAuthenticationResponse represents a generated type
type FinishPasskeyAuthenticationResponse struct {
}


// FinishPasskeyRegistrationResponse represents a generated type
type FinishPasskeyRegistrationResponse struct {
}


// GetJwksResponse represents a generated type
type GetJwksResponse struct {
	Keys []JwkKey `json:"keys"`
}


// GetScimUserResponse represents a generated type
type GetScimUserResponse struct {
	ConnectionID string `json:"connectionId"`
	User CompleteScimUserResponse `json:"user"`
	Groups []ScimGroupMembershipResponse `json:"groups"`
}


// GetScimUsersResponse represents a generated type
type GetScimUsersResponse struct {
	ConnectionID string `json:"connectionId"`
	Users []CompleteScimUserResponse `json:"users"`
	PageNumber int `json:"pageNumber"`
	PageSize int `json:"pageSize"`
	TotalResults int `json:"totalResults"`
}


// IdpInfoFromCustomerResponse is a discriminated union response type
type IdpInfoFromCustomerResponse interface {
	isIdpInfoFromCustomerResponse()
}

// IdpInfoFromCustomerResponseMicrosoftEntra represents the MicrosoftEntra variant
type IdpInfoFromCustomerResponseMicrosoftEntra struct {
	ClientID string `json:"clientId"`
	UsesPkce bool `json:"usesPkce"`
	TenantID string `json:"tenantId"`
}

func (r *IdpInfoFromCustomerResponseMicrosoftEntra) isIdpInfoFromCustomerResponse() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *IdpInfoFromCustomerResponseMicrosoftEntra) MarshalJSON() ([]byte, error) {
	type Alias IdpInfoFromCustomerResponseMicrosoftEntra
	return json.Marshal(&struct {
		IdpType string `json:"idpType"`
		*Alias
	}{
		IdpType: "MicrosoftEntra",
		Alias: (*Alias)(v),
	})
}

// IdpInfoFromCustomerResponseOkta represents the Okta variant
type IdpInfoFromCustomerResponseOkta struct {
	ClientID string `json:"clientId"`
	UsesPkce bool `json:"usesPkce"`
	SsoDomain string `json:"ssoDomain"`
}

func (r *IdpInfoFromCustomerResponseOkta) isIdpInfoFromCustomerResponse() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *IdpInfoFromCustomerResponseOkta) MarshalJSON() ([]byte, error) {
	type Alias IdpInfoFromCustomerResponseOkta
	return json.Marshal(&struct {
		IdpType string `json:"idpType"`
		*Alias
	}{
		IdpType: "Okta",
		Alias: (*Alias)(v),
	})
}

// IdpInfoFromCustomerResponseGeneric represents the Generic variant
type IdpInfoFromCustomerResponseGeneric struct {
	ClientID string `json:"clientId"`
	UsesPkce bool `json:"usesPkce"`
	AuthURL string `json:"authUrl"`
	TokenURL string `json:"tokenUrl"`
	UserinfoURL string `json:"userinfoUrl"`
}

func (r *IdpInfoFromCustomerResponseGeneric) isIdpInfoFromCustomerResponse() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *IdpInfoFromCustomerResponseGeneric) MarshalJSON() ([]byte, error) {
	type Alias IdpInfoFromCustomerResponseGeneric
	return json.Marshal(&struct {
		IdpType string `json:"idpType"`
		*Alias
	}{
		IdpType: "Generic",
		Alias: (*Alias)(v),
	})
}



// InitiateOidcLoginResponse represents a generated type
type InitiateOidcLoginResponse struct {
	SendUserToIdpURL string `json:"sendUserToIdpUrl"`
	StateForCookie string `json:"stateForCookie"`
}


// InvalidateAllImpersonationSessionsForEmployeeResponse represents a generated type
type InvalidateAllImpersonationSessionsForEmployeeResponse struct {
	SessionsInvalidated int `json:"sessionsInvalidated"`
}


// InvalidateAllImpersonationSessionsForUserResponse represents a generated type
type InvalidateAllImpersonationSessionsForUserResponse struct {
	SessionsInvalidated int `json:"sessionsInvalidated"`
}


// InvalidateAllSessionsForUserExceptOneResponse represents a generated type
type InvalidateAllSessionsForUserExceptOneResponse struct {
	SessionsInvalidated int `json:"sessionsInvalidated"`
}


// InvalidateAllSessionsForUserResponse represents a generated type
type InvalidateAllSessionsForUserResponse struct {
	SessionsInvalidated int `json:"sessionsInvalidated"`
}


// InvalidateImpersonationSessionByIdResponse represents a generated type
type InvalidateImpersonationSessionByIdResponse struct {
}


// InvalidateImpersonationSessionByTokenResponse represents a generated type
type InvalidateImpersonationSessionByTokenResponse struct {
}


// InvalidateSessionByIdResponse represents a generated type
type InvalidateSessionByIdResponse struct {
}


// InvalidateSessionByTokenResponse represents a generated type
type InvalidateSessionByTokenResponse struct {
}


// PatchOidcClientResponse represents a generated type
type PatchOidcClientResponse struct {
	ClientID string `json:"clientId"`
}


// PatchScimConnectionResponse represents a generated type
type PatchScimConnectionResponse struct {
}


// PingResponse represents a generated type
type PingResponse struct {
	Timestamp int64 `json:"timestamp"`
}


// RegisterDeviceResponse represents a generated type
type RegisterDeviceResponse struct {
	NewDeviceDetected bool `json:"newDeviceDetected"`
}


// ResetScimApiKeyResponse represents a generated type
type ResetScimApiKeyResponse struct {
	ConnectionID string `json:"connectionId"`
	ScimAPIKey string `json:"scimApiKey"`
}


// RotateStatelessTokenKeyResponse represents a generated type
type RotateStatelessTokenKeyResponse struct {
	NewKeyID string `json:"newKeyId"`
	NewKeyBecomesDefaultAt int64 `json:"newKeyBecomesDefaultAt"`
	ExistingKeysExpireAt int64 `json:"existingKeysExpireAt"`
}


// ScimGroupMembershipResponse represents a generated type
type ScimGroupMembershipResponse struct {
	GroupID string `json:"groupId"`
	DisplayName string `json:"displayName"`
	ExternalID *string `json:"externalId,omitempty"`
}


// ScimResponse is a discriminated union response type
type ScimResponse interface {
	isScimResponse()
}

// ScimResponseCompleted represents the Completed variant
type ScimResponseCompleted struct {
	ConnectionID string `json:"connectionId"`
	ResponseData *json.RawMessage `json:"responseData,omitempty"`
	ResponseHTTPCode int `json:"responseHttpCode"`
	AffectedUserIds []string `json:"affectedUserIds,omitempty"`
}

func (r *ScimResponseCompleted) isScimResponse() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *ScimResponseCompleted) MarshalJSON() ([]byte, error) {
	type Alias ScimResponseCompleted
	return json.Marshal(&struct {
		Status string `json:"status"`
		*Alias
	}{
		Status: "Completed",
		Alias: (*Alias)(v),
	})
}

// ScimResponseActionRequiredLinkUser represents the ActionRequired variant
type ScimResponseActionRequiredLinkUser struct {
	ConnectionID string `json:"connectionId"`
	CommitID string `json:"commitId"`
	PrimaryEmail *string `json:"primaryEmail,omitempty"`
	UserName string `json:"userName"`
	ParsedUserData json.RawMessage `json:"parsedUserData"`
	Active bool `json:"active"`
	SsoUserSubject *string `json:"ssoUserSubject,omitempty"`
}

func (r *ScimResponseActionRequiredLinkUser) isScimResponse() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *ScimResponseActionRequiredLinkUser) MarshalJSON() ([]byte, error) {
	type Alias ScimResponseActionRequiredLinkUser
	return json.Marshal(&struct {
		Status string `json:"status"`
		*Alias
	}{
		Status: "ActionRequired",
		Alias: (*Alias)(v),
	})
}

// ScimResponseActionRequiredDisableUser represents the ActionRequired variant
type ScimResponseActionRequiredDisableUser struct {
	ConnectionID string `json:"connectionId"`
	CommitID string `json:"commitId"`
	UserID string `json:"userId"`
	PrimaryEmail *string `json:"primaryEmail,omitempty"`
	ParsedUserData json.RawMessage `json:"parsedUserData"`
}

func (r *ScimResponseActionRequiredDisableUser) isScimResponse() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *ScimResponseActionRequiredDisableUser) MarshalJSON() ([]byte, error) {
	type Alias ScimResponseActionRequiredDisableUser
	return json.Marshal(&struct {
		Status string `json:"status"`
		*Alias
	}{
		Status: "ActionRequired",
		Alias: (*Alias)(v),
	})
}

// ScimResponseActionRequiredEnableUser represents the ActionRequired variant
type ScimResponseActionRequiredEnableUser struct {
	ConnectionID string `json:"connectionId"`
	CommitID string `json:"commitId"`
	UserID string `json:"userId"`
	PrimaryEmail *string `json:"primaryEmail,omitempty"`
	ParsedUserData json.RawMessage `json:"parsedUserData"`
}

func (r *ScimResponseActionRequiredEnableUser) isScimResponse() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *ScimResponseActionRequiredEnableUser) MarshalJSON() ([]byte, error) {
	type Alias ScimResponseActionRequiredEnableUser
	return json.Marshal(&struct {
		Status string `json:"status"`
		*Alias
	}{
		Status: "ActionRequired",
		Alias: (*Alias)(v),
	})
}

// ScimResponseActionRequiredDeleteUser represents the ActionRequired variant
type ScimResponseActionRequiredDeleteUser struct {
	ConnectionID string `json:"connectionId"`
	CommitID string `json:"commitId"`
	UserID string `json:"userId"`
	PrimaryEmail *string `json:"primaryEmail,omitempty"`
	ParsedUserData json.RawMessage `json:"parsedUserData"`
}

func (r *ScimResponseActionRequiredDeleteUser) isScimResponse() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *ScimResponseActionRequiredDeleteUser) MarshalJSON() ([]byte, error) {
	type Alias ScimResponseActionRequiredDeleteUser
	return json.Marshal(&struct {
		Status string `json:"status"`
		*Alias
	}{
		Status: "ActionRequired",
		Alias: (*Alias)(v),
	})
}



// StartPasskeyAuthenticationResponse represents a generated type
type StartPasskeyAuthenticationResponse struct {
	AuthenticationOptions json.RawMessage `json:"authenticationOptions"`
}


// StartPasskeyRegistrationResponse represents a generated type
type StartPasskeyRegistrationResponse struct {
	RegistrationOptions json.RawMessage `json:"registrationOptions"`
}


// UpdateSessionResponse represents a generated type
type UpdateSessionResponse struct {
}


// UpdateSessionsResponse represents a generated type
type UpdateSessionsResponse struct {
	UpdatedCount int `json:"updatedCount"`
}


// ValidateAndRefreshSessionResponse represents a generated type
type ValidateAndRefreshSessionResponse struct {
	SessionID string `json:"sessionId"`
	UserID string `json:"userId"`
	CreatedAt int64 `json:"createdAt"`
	ExpiresAt int64 `json:"expiresAt"`
	Tags []string `json:"tags,omitempty"`
	Metadata *json.RawMessage `json:"metadata,omitempty"`
	HasDeviceRegistered bool `json:"hasDeviceRegistered"`
	NewSessionToken *string `json:"newSessionToken,omitempty"`
}


// ValidateImpersonationSessionResponse represents a generated type
type ValidateImpersonationSessionResponse struct {
	ImpersonationSessionID string `json:"impersonationSessionId"`
	EmployeeEmail string `json:"employeeEmail"`
	TargetUserID string `json:"targetUserId"`
	CreatedAt int64 `json:"createdAt"`
	ExpiresAt int64 `json:"expiresAt"`
	Metadata *interface{} `json:"metadata,omitempty"`
}


// ValidateSessionResponse represents a generated type
type ValidateSessionResponse struct {
	SessionID string `json:"sessionId"`
	UserID string `json:"userId"`
	CreatedAt int64 `json:"createdAt"`
	ExpiresAt int64 `json:"expiresAt"`
	Tags []string `json:"tags,omitempty"`
	Metadata *json.RawMessage `json:"metadata,omitempty"`
	HasDeviceRegistered bool `json:"hasDeviceRegistered"`
}


// Response Unmarshalers

// UnmarshalIdpInfoFromCustomerResponse unmarshals JSON data into the appropriate IdpInfoFromCustomerResponse variant
func UnmarshalIdpInfoFromCustomerResponse(data []byte) (IdpInfoFromCustomerResponse, error) {
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
		var variant IdpInfoFromCustomerResponseMicrosoftEntra
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "Okta":
		var variant IdpInfoFromCustomerResponseOkta
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	case "Generic":
		var variant IdpInfoFromCustomerResponseGeneric
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	default:
		return nil, fmt.Errorf("unknown idpType: %s", discriminator)
	}
}


// UnmarshalScimResponse unmarshals JSON data into the appropriate ScimResponse variant
func UnmarshalScimResponse(data []byte) (ScimResponse, error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Primary discriminator
	var primary string
	if err := json.Unmarshal(raw["status"], &primary); err != nil {
		return nil, fmt.Errorf("parse status discriminator: %w", err)
	}

	if primary == "Completed" {
		var variant ScimResponseCompleted
		if err := json.Unmarshal(data, &variant); err != nil {
			return nil, err
		}
		return &variant, nil
	} else if primary == "ActionRequired" {
		// Secondary discriminator
		var secondary string
		if err := json.Unmarshal(raw["action"], &secondary); err != nil {
			return nil, fmt.Errorf("parse action discriminator: %w", err)
		}

		switch secondary {
		case "LinkUser":
			var variant ScimResponseActionRequiredLinkUser
			if err := json.Unmarshal(data, &variant); err != nil {
				return nil, err
			}
			return &variant, nil
		case "DisableUser":
			var variant ScimResponseActionRequiredDisableUser
			if err := json.Unmarshal(data, &variant); err != nil {
				return nil, err
			}
			return &variant, nil
		case "EnableUser":
			var variant ScimResponseActionRequiredEnableUser
			if err := json.Unmarshal(data, &variant); err != nil {
				return nil, err
			}
			return &variant, nil
		case "DeleteUser":
			var variant ScimResponseActionRequiredDeleteUser
			if err := json.Unmarshal(data, &variant); err != nil {
				return nil, err
			}
			return &variant, nil
		default:
			return nil, fmt.Errorf("unknown action: %s", secondary)
		}
	}

	return nil, fmt.Errorf("unknown status: %s", primary)
}
