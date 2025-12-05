package byo

import (
	"encoding/json"
	"fmt"
)

// CommitScimUserChangeError is the base error type
type CommitScimUserChangeError struct {
	Type    string
	message string
}

func (e *CommitScimUserChangeError) Error() string {
	return e.message
}

// CommitScimUserChangeErrorScimConnectionNotFound represents the ScimConnectionNotFound error variant
type CommitScimUserChangeErrorScimConnectionNotFound struct {
	*CommitScimUserChangeError
}

func newCommitScimUserChangeErrorScimConnectionNotFound() *CommitScimUserChangeErrorScimConnectionNotFound {
	return &CommitScimUserChangeErrorScimConnectionNotFound{
		CommitScimUserChangeError: &CommitScimUserChangeError{
			Type:    "ScimConnectionNotFound",
			message: "ScimConnectionNotFound",
		},
	}
}

// CommitScimUserChangeErrorStagedChangeNotFound represents the StagedChangeNotFound error variant
type CommitScimUserChangeErrorStagedChangeNotFound struct {
	*CommitScimUserChangeError
}

func newCommitScimUserChangeErrorStagedChangeNotFound() *CommitScimUserChangeErrorStagedChangeNotFound {
	return &CommitScimUserChangeErrorStagedChangeNotFound{
		CommitScimUserChangeError: &CommitScimUserChangeError{
			Type:    "StagedChangeNotFound",
			message: "StagedChangeNotFound",
		},
	}
}

// CommitScimUserChangeErrorUnexpectedError represents the UnexpectedError error variant
type CommitScimUserChangeErrorUnexpectedError struct {
	*CommitScimUserChangeError
	Details string
}

func newCommitScimUserChangeErrorUnexpectedError(details string) *CommitScimUserChangeErrorUnexpectedError {
	return &CommitScimUserChangeErrorUnexpectedError{
		CommitScimUserChangeError: &CommitScimUserChangeError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// CompleteOidcLoginError is the base error type
type CompleteOidcLoginError struct {
	Type    string
	message string
}

func (e *CompleteOidcLoginError) Error() string {
	return e.message
}

// CompleteOidcLoginErrorLoginBlockedByEmailAllowlist represents the LoginBlockedByEmailAllowlist error variant
type CompleteOidcLoginErrorLoginBlockedByEmailAllowlist struct {
	*CompleteOidcLoginError
}

func newCompleteOidcLoginErrorLoginBlockedByEmailAllowlist() *CompleteOidcLoginErrorLoginBlockedByEmailAllowlist {
	return &CompleteOidcLoginErrorLoginBlockedByEmailAllowlist{
		CompleteOidcLoginError: &CompleteOidcLoginError{
			Type:    "LoginBlockedByEmailAllowlist",
			message: "LoginBlockedByEmailAllowlist",
		},
	}
}

// CompleteOidcLoginErrorScimUserNotFoundWhereExpected represents the ScimUserNotFoundWhereExpected error variant
type CompleteOidcLoginErrorScimUserNotFoundWhereExpected struct {
	*CompleteOidcLoginError
}

func newCompleteOidcLoginErrorScimUserNotFoundWhereExpected() *CompleteOidcLoginErrorScimUserNotFoundWhereExpected {
	return &CompleteOidcLoginErrorScimUserNotFoundWhereExpected{
		CompleteOidcLoginError: &CompleteOidcLoginError{
			Type:    "ScimUserNotFoundWhereExpected",
			message: "ScimUserNotFoundWhereExpected",
		},
	}
}

// CompleteOidcLoginErrorScimUserNotActive represents the ScimUserNotActive error variant
type CompleteOidcLoginErrorScimUserNotActive struct {
	*CompleteOidcLoginError
}

func newCompleteOidcLoginErrorScimUserNotActive() *CompleteOidcLoginErrorScimUserNotActive {
	return &CompleteOidcLoginErrorScimUserNotActive{
		CompleteOidcLoginError: &CompleteOidcLoginError{
			Type:    "ScimUserNotActive",
			message: "ScimUserNotActive",
		},
	}
}

// CompleteOidcLoginErrorInvalidLoginRequest represents the InvalidLoginRequest error variant
type CompleteOidcLoginErrorInvalidLoginRequest struct {
	*CompleteOidcLoginError
	Details OidcLoginRequestError
}

func newCompleteOidcLoginErrorInvalidLoginRequest(details OidcLoginRequestError) *CompleteOidcLoginErrorInvalidLoginRequest {
	return &CompleteOidcLoginErrorInvalidLoginRequest{
		CompleteOidcLoginError: &CompleteOidcLoginError{
			Type:    "InvalidLoginRequest",
			message: "InvalidLoginRequest",
		},
		Details: details,
	}
}

// CompleteOidcLoginErrorIdentityProviderError represents the IdentityProviderError error variant
type CompleteOidcLoginErrorIdentityProviderError struct {
	*CompleteOidcLoginError
	Details IdentityProviderError
}

func newCompleteOidcLoginErrorIdentityProviderError(details IdentityProviderError) *CompleteOidcLoginErrorIdentityProviderError {
	return &CompleteOidcLoginErrorIdentityProviderError{
		CompleteOidcLoginError: &CompleteOidcLoginError{
			Type:    "IdentityProviderError",
			message: "IdentityProviderError",
		},
		Details: details,
	}
}

// CompleteOidcLoginErrorUnexpectedError represents the UnexpectedError error variant
type CompleteOidcLoginErrorUnexpectedError struct {
	*CompleteOidcLoginError
	Details string
}

func newCompleteOidcLoginErrorUnexpectedError(details string) *CompleteOidcLoginErrorUnexpectedError {
	return &CompleteOidcLoginErrorUnexpectedError{
		CompleteOidcLoginError: &CompleteOidcLoginError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// CreateDeviceChallengeError represents a generated type
type CreateDeviceChallengeError struct {
	Type string `json:"type"`
	Details string `json:"details"`
}



func (e *CreateDeviceChallengeError) Error() string {
	return fmt.Sprintf("CreateDeviceChallengeError: %v", e)
}


// CreateImpersonationSessionError is the base error type
type CreateImpersonationSessionError struct {
	Type    string
	message string
}

func (e *CreateImpersonationSessionError) Error() string {
	return e.message
}

// CreateImpersonationSessionErrorImpersonationDisabled represents the ImpersonationDisabled error variant
type CreateImpersonationSessionErrorImpersonationDisabled struct {
	*CreateImpersonationSessionError
}

func newCreateImpersonationSessionErrorImpersonationDisabled() *CreateImpersonationSessionErrorImpersonationDisabled {
	return &CreateImpersonationSessionErrorImpersonationDisabled{
		CreateImpersonationSessionError: &CreateImpersonationSessionError{
			Type:    "ImpersonationDisabled",
			message: "ImpersonationDisabled",
		},
	}
}

// CreateImpersonationSessionErrorUnauthorizedEmployee represents the UnauthorizedEmployee error variant
type CreateImpersonationSessionErrorUnauthorizedEmployee struct {
	*CreateImpersonationSessionError
	Details string
}

func newCreateImpersonationSessionErrorUnauthorizedEmployee(details string) *CreateImpersonationSessionErrorUnauthorizedEmployee {
	return &CreateImpersonationSessionErrorUnauthorizedEmployee{
		CreateImpersonationSessionError: &CreateImpersonationSessionError{
			Type:    "UnauthorizedEmployee",
			message: "UnauthorizedEmployee",
		},
		Details: details,
	}
}

// CreateImpersonationSessionErrorUnexpectedError represents the UnexpectedError error variant
type CreateImpersonationSessionErrorUnexpectedError struct {
	*CreateImpersonationSessionError
	Details string
}

func newCreateImpersonationSessionErrorUnexpectedError(details string) *CreateImpersonationSessionErrorUnexpectedError {
	return &CreateImpersonationSessionErrorUnexpectedError{
		CreateImpersonationSessionError: &CreateImpersonationSessionError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// CreateOidcClientError is the base error type
type CreateOidcClientError struct {
	Type    string
	message string
}

func (e *CreateOidcClientError) Error() string {
	return e.message
}

// CreateOidcClientErrorInvalidFields represents the InvalidFields error variant
type CreateOidcClientErrorInvalidFields struct {
	*CreateOidcClientError
	Details CreateOidcClientErrorInvalidFieldsDetails
}

func newCreateOidcClientErrorInvalidFields(details CreateOidcClientErrorInvalidFieldsDetails) *CreateOidcClientErrorInvalidFields {
	return &CreateOidcClientErrorInvalidFields{
		CreateOidcClientError: &CreateOidcClientError{
			Type:    "InvalidFields",
			message: "InvalidFields",
		},
		Details: details,
	}
}

// CreateOidcClientErrorClientIDAlreadyTaken represents the ClientIdAlreadyTaken error variant
type CreateOidcClientErrorClientIDAlreadyTaken struct {
	*CreateOidcClientError
}

func newCreateOidcClientErrorClientIDAlreadyTaken() *CreateOidcClientErrorClientIDAlreadyTaken {
	return &CreateOidcClientErrorClientIDAlreadyTaken{
		CreateOidcClientError: &CreateOidcClientError{
			Type:    "ClientIdAlreadyTaken",
			message: "ClientIdAlreadyTaken",
		},
	}
}

// CreateOidcClientErrorCustomerIDAlreadyTakenForEoidcClient represents the CustomerIdAlreadyTakenForEoidcClient error variant
type CreateOidcClientErrorCustomerIDAlreadyTakenForEoidcClient struct {
	*CreateOidcClientError
}

func newCreateOidcClientErrorCustomerIDAlreadyTakenForEoidcClient() *CreateOidcClientErrorCustomerIDAlreadyTakenForEoidcClient {
	return &CreateOidcClientErrorCustomerIDAlreadyTakenForEoidcClient{
		CreateOidcClientError: &CreateOidcClientError{
			Type:    "CustomerIdAlreadyTakenForEoidcClient",
			message: "CustomerIdAlreadyTakenForEoidcClient",
		},
	}
}

// CreateOidcClientErrorUnexpectedError represents the UnexpectedError error variant
type CreateOidcClientErrorUnexpectedError struct {
	*CreateOidcClientError
	Details string
}

func newCreateOidcClientErrorUnexpectedError(details string) *CreateOidcClientErrorUnexpectedError {
	return &CreateOidcClientErrorUnexpectedError{
		CreateOidcClientError: &CreateOidcClientError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// CreateScimConnectionError is the base error type
type CreateScimConnectionError struct {
	Type    string
	message string
}

func (e *CreateScimConnectionError) Error() string {
	return e.message
}

// CreateScimConnectionErrorInvalidFields represents the InvalidFields error variant
type CreateScimConnectionErrorInvalidFields struct {
	*CreateScimConnectionError
	Details CreateScimConnectionErrorInvalidFieldsDetails
}

func newCreateScimConnectionErrorInvalidFields(details CreateScimConnectionErrorInvalidFieldsDetails) *CreateScimConnectionErrorInvalidFields {
	return &CreateScimConnectionErrorInvalidFields{
		CreateScimConnectionError: &CreateScimConnectionError{
			Type:    "InvalidFields",
			message: "InvalidFields",
		},
		Details: details,
	}
}

// CreateScimConnectionErrorScimConnectionForCustomerIDAlreadyExists represents the ScimConnectionForCustomerIdAlreadyExists error variant
type CreateScimConnectionErrorScimConnectionForCustomerIDAlreadyExists struct {
	*CreateScimConnectionError
}

func newCreateScimConnectionErrorScimConnectionForCustomerIDAlreadyExists() *CreateScimConnectionErrorScimConnectionForCustomerIDAlreadyExists {
	return &CreateScimConnectionErrorScimConnectionForCustomerIDAlreadyExists{
		CreateScimConnectionError: &CreateScimConnectionError{
			Type:    "ScimConnectionForCustomerIdAlreadyExists",
			message: "ScimConnectionForCustomerIdAlreadyExists",
		},
	}
}

// CreateScimConnectionErrorUnexpectedError represents the UnexpectedError error variant
type CreateScimConnectionErrorUnexpectedError struct {
	*CreateScimConnectionError
	Details string
}

func newCreateScimConnectionErrorUnexpectedError(details string) *CreateScimConnectionErrorUnexpectedError {
	return &CreateScimConnectionErrorUnexpectedError{
		CreateScimConnectionError: &CreateScimConnectionError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// CreateSessionError is the base error type
type CreateSessionError struct {
	Type    string
	message string
}

func (e *CreateSessionError) Error() string {
	return e.message
}

// CreateSessionErrorSessionLimitExceeded represents the SessionLimitExceeded error variant
type CreateSessionErrorSessionLimitExceeded struct {
	*CreateSessionError
	Details CreateSessionErrorSessionLimitExceededDetails
}

func newCreateSessionErrorSessionLimitExceeded(details CreateSessionErrorSessionLimitExceededDetails) *CreateSessionErrorSessionLimitExceeded {
	return &CreateSessionErrorSessionLimitExceeded{
		CreateSessionError: &CreateSessionError{
			Type:    "SessionLimitExceeded",
			message: "SessionLimitExceeded",
		},
		Details: details,
	}
}

// CreateSessionErrorIPAddressError represents the IpAddressError error variant
type CreateSessionErrorIPAddressError struct {
	*CreateSessionError
	Details string
}

func newCreateSessionErrorIPAddressError(details string) *CreateSessionErrorIPAddressError {
	return &CreateSessionErrorIPAddressError{
		CreateSessionError: &CreateSessionError{
			Type:    "IpAddressError",
			message: "IpAddressError",
		},
		Details: details,
	}
}

// CreateSessionErrorTagParseError represents the TagParseError error variant
type CreateSessionErrorTagParseError struct {
	*CreateSessionError
	Details InvalidTagError
}

func newCreateSessionErrorTagParseError(details InvalidTagError) *CreateSessionErrorTagParseError {
	return &CreateSessionErrorTagParseError{
		CreateSessionError: &CreateSessionError{
			Type:    "TagParseError",
			message: "TagParseError",
		},
		Details: details,
	}
}

// CreateSessionErrorInvalidDeviceRegistration represents the InvalidDeviceRegistration error variant
type CreateSessionErrorInvalidDeviceRegistration struct {
	*CreateSessionError
}

func newCreateSessionErrorInvalidDeviceRegistration() *CreateSessionErrorInvalidDeviceRegistration {
	return &CreateSessionErrorInvalidDeviceRegistration{
		CreateSessionError: &CreateSessionError{
			Type:    "InvalidDeviceRegistration",
			message: "InvalidDeviceRegistration",
		},
	}
}

// CreateSessionErrorUnexpectedError represents the UnexpectedError error variant
type CreateSessionErrorUnexpectedError struct {
	*CreateSessionError
	Details string
}

func newCreateSessionErrorUnexpectedError(details string) *CreateSessionErrorUnexpectedError {
	return &CreateSessionErrorUnexpectedError{
		CreateSessionError: &CreateSessionError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// CreateStatelessTokenError is the base error type
type CreateStatelessTokenError struct {
	Type    string
	message string
}

func (e *CreateStatelessTokenError) Error() string {
	return e.message
}

// CreateStatelessTokenErrorTokenCreationFailed represents the TokenCreationFailed error variant
type CreateStatelessTokenErrorTokenCreationFailed struct {
	*CreateStatelessTokenError
	Details string
}

func newCreateStatelessTokenErrorTokenCreationFailed(details string) *CreateStatelessTokenErrorTokenCreationFailed {
	return &CreateStatelessTokenErrorTokenCreationFailed{
		CreateStatelessTokenError: &CreateStatelessTokenError{
			Type:    "TokenCreationFailed",
			message: "TokenCreationFailed",
		},
		Details: details,
	}
}

// CreateStatelessTokenErrorUnexpectedError represents the UnexpectedError error variant
type CreateStatelessTokenErrorUnexpectedError struct {
	*CreateStatelessTokenError
	Details string
}

func newCreateStatelessTokenErrorUnexpectedError(details string) *CreateStatelessTokenErrorUnexpectedError {
	return &CreateStatelessTokenErrorUnexpectedError{
		CreateStatelessTokenError: &CreateStatelessTokenError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// DeleteOidcClientError is the base error type
type DeleteOidcClientError struct {
	Type    string
	message string
}

func (e *DeleteOidcClientError) Error() string {
	return e.message
}

// DeleteOidcClientErrorOIDCClientNotFound represents the OidcClientNotFound error variant
type DeleteOidcClientErrorOIDCClientNotFound struct {
	*DeleteOidcClientError
}

func newDeleteOidcClientErrorOIDCClientNotFound() *DeleteOidcClientErrorOIDCClientNotFound {
	return &DeleteOidcClientErrorOIDCClientNotFound{
		DeleteOidcClientError: &DeleteOidcClientError{
			Type:    "OidcClientNotFound",
			message: "OidcClientNotFound",
		},
	}
}

// DeleteOidcClientErrorUnexpectedError represents the UnexpectedError error variant
type DeleteOidcClientErrorUnexpectedError struct {
	*DeleteOidcClientError
	Details string
}

func newDeleteOidcClientErrorUnexpectedError(details string) *DeleteOidcClientErrorUnexpectedError {
	return &DeleteOidcClientErrorUnexpectedError{
		DeleteOidcClientError: &DeleteOidcClientError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// DeleteScimConnectionError is the base error type
type DeleteScimConnectionError struct {
	Type    string
	message string
}

func (e *DeleteScimConnectionError) Error() string {
	return e.message
}

// DeleteScimConnectionErrorScimConnectionNotFound represents the ScimConnectionNotFound error variant
type DeleteScimConnectionErrorScimConnectionNotFound struct {
	*DeleteScimConnectionError
}

func newDeleteScimConnectionErrorScimConnectionNotFound() *DeleteScimConnectionErrorScimConnectionNotFound {
	return &DeleteScimConnectionErrorScimConnectionNotFound{
		DeleteScimConnectionError: &DeleteScimConnectionError{
			Type:    "ScimConnectionNotFound",
			message: "ScimConnectionNotFound",
		},
	}
}

// DeleteScimConnectionErrorUnexpectedError represents the UnexpectedError error variant
type DeleteScimConnectionErrorUnexpectedError struct {
	*DeleteScimConnectionError
	Details string
}

func newDeleteScimConnectionErrorUnexpectedError(details string) *DeleteScimConnectionErrorUnexpectedError {
	return &DeleteScimConnectionErrorUnexpectedError{
		DeleteScimConnectionError: &DeleteScimConnectionError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// DeregisterAllPasskeysForUserError represents a generated type
type DeregisterAllPasskeysForUserError struct {
	Type string `json:"type"`
	Details string `json:"details"`
}



func (e *DeregisterAllPasskeysForUserError) Error() string {
	return fmt.Sprintf("DeregisterAllPasskeysForUserError: %v", e)
}


// DeregisterPasskeyError is the base error type
type DeregisterPasskeyError struct {
	Type    string
	message string
}

func (e *DeregisterPasskeyError) Error() string {
	return e.message
}

// DeregisterPasskeyErrorPasskeyNotFound represents the PasskeyNotFound error variant
type DeregisterPasskeyErrorPasskeyNotFound struct {
	*DeregisterPasskeyError
}

func newDeregisterPasskeyErrorPasskeyNotFound() *DeregisterPasskeyErrorPasskeyNotFound {
	return &DeregisterPasskeyErrorPasskeyNotFound{
		DeregisterPasskeyError: &DeregisterPasskeyError{
			Type:    "PasskeyNotFound",
			message: "PasskeyNotFound",
		},
	}
}

// DeregisterPasskeyErrorUnexpectedError represents the UnexpectedError error variant
type DeregisterPasskeyErrorUnexpectedError struct {
	*DeregisterPasskeyError
	Details string
}

func newDeregisterPasskeyErrorUnexpectedError(details string) *DeregisterPasskeyErrorUnexpectedError {
	return &DeregisterPasskeyErrorUnexpectedError{
		DeregisterPasskeyError: &DeregisterPasskeyError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// FetchAllActiveImpersonationSessionsError is the base error type
type FetchAllActiveImpersonationSessionsError struct {
	Type    string
	message string
}

func (e *FetchAllActiveImpersonationSessionsError) Error() string {
	return e.message
}

// FetchAllActiveImpersonationSessionsErrorInvalidPagingToken represents the InvalidPagingToken error variant
type FetchAllActiveImpersonationSessionsErrorInvalidPagingToken struct {
	*FetchAllActiveImpersonationSessionsError
}

func newFetchAllActiveImpersonationSessionsErrorInvalidPagingToken() *FetchAllActiveImpersonationSessionsErrorInvalidPagingToken {
	return &FetchAllActiveImpersonationSessionsErrorInvalidPagingToken{
		FetchAllActiveImpersonationSessionsError: &FetchAllActiveImpersonationSessionsError{
			Type:    "InvalidPagingToken",
			message: "InvalidPagingToken",
		},
	}
}

// FetchAllActiveImpersonationSessionsErrorUnexpectedError represents the UnexpectedError error variant
type FetchAllActiveImpersonationSessionsErrorUnexpectedError struct {
	*FetchAllActiveImpersonationSessionsError
	Details string
}

func newFetchAllActiveImpersonationSessionsErrorUnexpectedError(details string) *FetchAllActiveImpersonationSessionsErrorUnexpectedError {
	return &FetchAllActiveImpersonationSessionsErrorUnexpectedError{
		FetchAllActiveImpersonationSessionsError: &FetchAllActiveImpersonationSessionsError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// FetchAllImpersonationSessionsForEmployeeError represents a generated type
type FetchAllImpersonationSessionsForEmployeeError struct {
	Type string `json:"type"`
	Details string `json:"details"`
}



func (e *FetchAllImpersonationSessionsForEmployeeError) Error() string {
	return fmt.Sprintf("FetchAllImpersonationSessionsForEmployeeError: %v", e)
}


// FetchAllImpersonationSessionsForUserError represents a generated type
type FetchAllImpersonationSessionsForUserError struct {
	Type string `json:"type"`
	Details string `json:"details"`
}



func (e *FetchAllImpersonationSessionsForUserError) Error() string {
	return fmt.Sprintf("FetchAllImpersonationSessionsForUserError: %v", e)
}


// FetchAllPasskeysForUserError represents a generated type
type FetchAllPasskeysForUserError struct {
	Type string `json:"type"`
	Details string `json:"details"`
}



func (e *FetchAllPasskeysForUserError) Error() string {
	return fmt.Sprintf("FetchAllPasskeysForUserError: %v", e)
}


// FetchAllSessionsError represents a generated type
type FetchAllSessionsError struct {
	Type string `json:"type"`
	Details string `json:"details"`
}



func (e *FetchAllSessionsError) Error() string {
	return fmt.Sprintf("FetchAllSessionsError: %v", e)
}


// FetchAllSessionsForUserError represents a generated type
type FetchAllSessionsForUserError struct {
	Type string `json:"type"`
	Details string `json:"details"`
}



func (e *FetchAllSessionsForUserError) Error() string {
	return fmt.Sprintf("FetchAllSessionsForUserError: %v", e)
}


// FetchImpersonationSessionByIdError is the base error type
type FetchImpersonationSessionByIdError struct {
	Type    string
	message string
}

func (e *FetchImpersonationSessionByIdError) Error() string {
	return e.message
}

// FetchImpersonationSessionByIdErrorSessionNotFound represents the SessionNotFound error variant
type FetchImpersonationSessionByIdErrorSessionNotFound struct {
	*FetchImpersonationSessionByIdError
}

func newFetchImpersonationSessionByIdErrorSessionNotFound() *FetchImpersonationSessionByIdErrorSessionNotFound {
	return &FetchImpersonationSessionByIdErrorSessionNotFound{
		FetchImpersonationSessionByIdError: &FetchImpersonationSessionByIdError{
			Type:    "SessionNotFound",
			message: "SessionNotFound",
		},
	}
}

// FetchImpersonationSessionByIdErrorUnexpectedError represents the UnexpectedError error variant
type FetchImpersonationSessionByIdErrorUnexpectedError struct {
	*FetchImpersonationSessionByIdError
	Details string
}

func newFetchImpersonationSessionByIdErrorUnexpectedError(details string) *FetchImpersonationSessionByIdErrorUnexpectedError {
	return &FetchImpersonationSessionByIdErrorUnexpectedError{
		FetchImpersonationSessionByIdError: &FetchImpersonationSessionByIdError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// FetchOidcClientError is the base error type
type FetchOidcClientError struct {
	Type    string
	message string
}

func (e *FetchOidcClientError) Error() string {
	return e.message
}

// FetchOidcClientErrorOIDCClientNotFound represents the OidcClientNotFound error variant
type FetchOidcClientErrorOIDCClientNotFound struct {
	*FetchOidcClientError
}

func newFetchOidcClientErrorOIDCClientNotFound() *FetchOidcClientErrorOIDCClientNotFound {
	return &FetchOidcClientErrorOIDCClientNotFound{
		FetchOidcClientError: &FetchOidcClientError{
			Type:    "OidcClientNotFound",
			message: "OidcClientNotFound",
		},
	}
}

// FetchOidcClientErrorUnexpectedError represents the UnexpectedError error variant
type FetchOidcClientErrorUnexpectedError struct {
	*FetchOidcClientError
	Details string
}

func newFetchOidcClientErrorUnexpectedError(details string) *FetchOidcClientErrorUnexpectedError {
	return &FetchOidcClientErrorUnexpectedError{
		FetchOidcClientError: &FetchOidcClientError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// FetchScimConnectionError is the base error type
type FetchScimConnectionError struct {
	Type    string
	message string
}

func (e *FetchScimConnectionError) Error() string {
	return e.message
}

// FetchScimConnectionErrorScimConnectionNotFound represents the ScimConnectionNotFound error variant
type FetchScimConnectionErrorScimConnectionNotFound struct {
	*FetchScimConnectionError
}

func newFetchScimConnectionErrorScimConnectionNotFound() *FetchScimConnectionErrorScimConnectionNotFound {
	return &FetchScimConnectionErrorScimConnectionNotFound{
		FetchScimConnectionError: &FetchScimConnectionError{
			Type:    "ScimConnectionNotFound",
			message: "ScimConnectionNotFound",
		},
	}
}

// FetchScimConnectionErrorUnexpectedError represents the UnexpectedError error variant
type FetchScimConnectionErrorUnexpectedError struct {
	*FetchScimConnectionError
	Details string
}

func newFetchScimConnectionErrorUnexpectedError(details string) *FetchScimConnectionErrorUnexpectedError {
	return &FetchScimConnectionErrorUnexpectedError{
		FetchScimConnectionError: &FetchScimConnectionError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// FetchSessionByIdError is the base error type
type FetchSessionByIdError struct {
	Type    string
	message string
}

func (e *FetchSessionByIdError) Error() string {
	return e.message
}

// FetchSessionByIdErrorSessionNotFound represents the SessionNotFound error variant
type FetchSessionByIdErrorSessionNotFound struct {
	*FetchSessionByIdError
}

func newFetchSessionByIdErrorSessionNotFound() *FetchSessionByIdErrorSessionNotFound {
	return &FetchSessionByIdErrorSessionNotFound{
		FetchSessionByIdError: &FetchSessionByIdError{
			Type:    "SessionNotFound",
			message: "SessionNotFound",
		},
	}
}

// FetchSessionByIdErrorUnexpectedError represents the UnexpectedError error variant
type FetchSessionByIdErrorUnexpectedError struct {
	*FetchSessionByIdError
	Details string
}

func newFetchSessionByIdErrorUnexpectedError(details string) *FetchSessionByIdErrorUnexpectedError {
	return &FetchSessionByIdErrorUnexpectedError{
		FetchSessionByIdError: &FetchSessionByIdError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// FinishPasskeyAuthenticationError is the base error type
type FinishPasskeyAuthenticationError struct {
	Type    string
	message string
}

func (e *FinishPasskeyAuthenticationError) Error() string {
	return e.message
}

// FinishPasskeyAuthenticationErrorCannotParseAdditionalAllowedOrigin represents the CannotParseAdditionalAllowedOrigin error variant
type FinishPasskeyAuthenticationErrorCannotParseAdditionalAllowedOrigin struct {
	*FinishPasskeyAuthenticationError
}

func newFinishPasskeyAuthenticationErrorCannotParseAdditionalAllowedOrigin() *FinishPasskeyAuthenticationErrorCannotParseAdditionalAllowedOrigin {
	return &FinishPasskeyAuthenticationErrorCannotParseAdditionalAllowedOrigin{
		FinishPasskeyAuthenticationError: &FinishPasskeyAuthenticationError{
			Type:    "CannotParseAdditionalAllowedOrigin",
			message: "CannotParseAdditionalAllowedOrigin",
		},
	}
}

// FinishPasskeyAuthenticationErrorNoAuthenticationChallengeFound represents the NoAuthenticationChallengeFound error variant
type FinishPasskeyAuthenticationErrorNoAuthenticationChallengeFound struct {
	*FinishPasskeyAuthenticationError
}

func newFinishPasskeyAuthenticationErrorNoAuthenticationChallengeFound() *FinishPasskeyAuthenticationErrorNoAuthenticationChallengeFound {
	return &FinishPasskeyAuthenticationErrorNoAuthenticationChallengeFound{
		FinishPasskeyAuthenticationError: &FinishPasskeyAuthenticationError{
			Type:    "NoAuthenticationChallengeFound",
			message: "NoAuthenticationChallengeFound",
		},
	}
}

// FinishPasskeyAuthenticationErrorOriginNotAllowed represents the OriginNotAllowed error variant
type FinishPasskeyAuthenticationErrorOriginNotAllowed struct {
	*FinishPasskeyAuthenticationError
	Details FinishPasskeyAuthenticationErrorOriginNotAllowedDetails
}

func newFinishPasskeyAuthenticationErrorOriginNotAllowed(details FinishPasskeyAuthenticationErrorOriginNotAllowedDetails) *FinishPasskeyAuthenticationErrorOriginNotAllowed {
	return &FinishPasskeyAuthenticationErrorOriginNotAllowed{
		FinishPasskeyAuthenticationError: &FinishPasskeyAuthenticationError{
			Type:    "OriginNotAllowed",
			message: "OriginNotAllowed",
		},
		Details: details,
	}
}

// FinishPasskeyAuthenticationErrorUnexpectedError represents the UnexpectedError error variant
type FinishPasskeyAuthenticationErrorUnexpectedError struct {
	*FinishPasskeyAuthenticationError
	Details string
}

func newFinishPasskeyAuthenticationErrorUnexpectedError(details string) *FinishPasskeyAuthenticationErrorUnexpectedError {
	return &FinishPasskeyAuthenticationErrorUnexpectedError{
		FinishPasskeyAuthenticationError: &FinishPasskeyAuthenticationError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// FinishPasskeyRegistrationError is the base error type
type FinishPasskeyRegistrationError struct {
	Type    string
	message string
}

func (e *FinishPasskeyRegistrationError) Error() string {
	return e.message
}

// FinishPasskeyRegistrationErrorCannotParseAdditionalAllowedOrigin represents the CannotParseAdditionalAllowedOrigin error variant
type FinishPasskeyRegistrationErrorCannotParseAdditionalAllowedOrigin struct {
	*FinishPasskeyRegistrationError
}

func newFinishPasskeyRegistrationErrorCannotParseAdditionalAllowedOrigin() *FinishPasskeyRegistrationErrorCannotParseAdditionalAllowedOrigin {
	return &FinishPasskeyRegistrationErrorCannotParseAdditionalAllowedOrigin{
		FinishPasskeyRegistrationError: &FinishPasskeyRegistrationError{
			Type:    "CannotParseAdditionalAllowedOrigin",
			message: "CannotParseAdditionalAllowedOrigin",
		},
	}
}

// FinishPasskeyRegistrationErrorNoRegistrationChallengeFound represents the NoRegistrationChallengeFound error variant
type FinishPasskeyRegistrationErrorNoRegistrationChallengeFound struct {
	*FinishPasskeyRegistrationError
}

func newFinishPasskeyRegistrationErrorNoRegistrationChallengeFound() *FinishPasskeyRegistrationErrorNoRegistrationChallengeFound {
	return &FinishPasskeyRegistrationErrorNoRegistrationChallengeFound{
		FinishPasskeyRegistrationError: &FinishPasskeyRegistrationError{
			Type:    "NoRegistrationChallengeFound",
			message: "NoRegistrationChallengeFound",
		},
	}
}

// FinishPasskeyRegistrationErrorOriginNotAllowed represents the OriginNotAllowed error variant
type FinishPasskeyRegistrationErrorOriginNotAllowed struct {
	*FinishPasskeyRegistrationError
	Details FinishPasskeyRegistrationErrorOriginNotAllowedDetails
}

func newFinishPasskeyRegistrationErrorOriginNotAllowed(details FinishPasskeyRegistrationErrorOriginNotAllowedDetails) *FinishPasskeyRegistrationErrorOriginNotAllowed {
	return &FinishPasskeyRegistrationErrorOriginNotAllowed{
		FinishPasskeyRegistrationError: &FinishPasskeyRegistrationError{
			Type:    "OriginNotAllowed",
			message: "OriginNotAllowed",
		},
		Details: details,
	}
}

// FinishPasskeyRegistrationErrorPasskeyForUserAlreadyExists represents the PasskeyForUserAlreadyExists error variant
type FinishPasskeyRegistrationErrorPasskeyForUserAlreadyExists struct {
	*FinishPasskeyRegistrationError
}

func newFinishPasskeyRegistrationErrorPasskeyForUserAlreadyExists() *FinishPasskeyRegistrationErrorPasskeyForUserAlreadyExists {
	return &FinishPasskeyRegistrationErrorPasskeyForUserAlreadyExists{
		FinishPasskeyRegistrationError: &FinishPasskeyRegistrationError{
			Type:    "PasskeyForUserAlreadyExists",
			message: "PasskeyForUserAlreadyExists",
		},
	}
}

// FinishPasskeyRegistrationErrorUnexpectedError represents the UnexpectedError error variant
type FinishPasskeyRegistrationErrorUnexpectedError struct {
	*FinishPasskeyRegistrationError
	Details string
}

func newFinishPasskeyRegistrationErrorUnexpectedError(details string) *FinishPasskeyRegistrationErrorUnexpectedError {
	return &FinishPasskeyRegistrationErrorUnexpectedError{
		FinishPasskeyRegistrationError: &FinishPasskeyRegistrationError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// GetJwksError represents a generated type
type GetJwksError struct {
	Type string `json:"type"`
	Details string `json:"details"`
}



func (e *GetJwksError) Error() string {
	return fmt.Sprintf("GetJwksError: %v", e)
}


// GetScimUserError is the base error type
type GetScimUserError struct {
	Type    string
	message string
}

func (e *GetScimUserError) Error() string {
	return e.message
}

// GetScimUserErrorScimConnectionNotFound represents the ScimConnectionNotFound error variant
type GetScimUserErrorScimConnectionNotFound struct {
	*GetScimUserError
}

func newGetScimUserErrorScimConnectionNotFound() *GetScimUserErrorScimConnectionNotFound {
	return &GetScimUserErrorScimConnectionNotFound{
		GetScimUserError: &GetScimUserError{
			Type:    "ScimConnectionNotFound",
			message: "ScimConnectionNotFound",
		},
	}
}

// GetScimUserErrorUserNotFound represents the UserNotFound error variant
type GetScimUserErrorUserNotFound struct {
	*GetScimUserError
}

func newGetScimUserErrorUserNotFound() *GetScimUserErrorUserNotFound {
	return &GetScimUserErrorUserNotFound{
		GetScimUserError: &GetScimUserError{
			Type:    "UserNotFound",
			message: "UserNotFound",
		},
	}
}

// GetScimUserErrorUnexpectedError represents the UnexpectedError error variant
type GetScimUserErrorUnexpectedError struct {
	*GetScimUserError
	Details string
}

func newGetScimUserErrorUnexpectedError(details string) *GetScimUserErrorUnexpectedError {
	return &GetScimUserErrorUnexpectedError{
		GetScimUserError: &GetScimUserError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// GetScimUsersError is the base error type
type GetScimUsersError struct {
	Type    string
	message string
}

func (e *GetScimUsersError) Error() string {
	return e.message
}

// GetScimUsersErrorScimConnectionNotFound represents the ScimConnectionNotFound error variant
type GetScimUsersErrorScimConnectionNotFound struct {
	*GetScimUsersError
}

func newGetScimUsersErrorScimConnectionNotFound() *GetScimUsersErrorScimConnectionNotFound {
	return &GetScimUsersErrorScimConnectionNotFound{
		GetScimUsersError: &GetScimUsersError{
			Type:    "ScimConnectionNotFound",
			message: "ScimConnectionNotFound",
		},
	}
}

// GetScimUsersErrorInvalidQueryField represents the InvalidQueryField error variant
type GetScimUsersErrorInvalidQueryField struct {
	*GetScimUsersError
	Details GetScimUsersErrorInvalidQueryFieldDetails
}

func newGetScimUsersErrorInvalidQueryField(details GetScimUsersErrorInvalidQueryFieldDetails) *GetScimUsersErrorInvalidQueryField {
	return &GetScimUsersErrorInvalidQueryField{
		GetScimUsersError: &GetScimUsersError{
			Type:    "InvalidQueryField",
			message: "InvalidQueryField",
		},
		Details: details,
	}
}

// GetScimUsersErrorUnexpectedError represents the UnexpectedError error variant
type GetScimUsersErrorUnexpectedError struct {
	*GetScimUsersError
	Details string
}

func newGetScimUsersErrorUnexpectedError(details string) *GetScimUsersErrorUnexpectedError {
	return &GetScimUsersErrorUnexpectedError{
		GetScimUsersError: &GetScimUsersError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// IdentityProviderError represents an enumeration type
type IdentityProviderError string

const (
	IdentityProviderErrorNoCodeProvidedByIdp IdentityProviderError = "NoCodeProvidedByIdp"
	IdentityProviderErrorUnableToExchangeCodeForToken IdentityProviderError = "UnableToExchangeCodeForToken"
	IdentityProviderErrorErrorCommunicatingWithIdp IdentityProviderError = "ErrorCommunicatingWithIdp"
)


// InitiateOidcLoginError is the base error type
type InitiateOidcLoginError struct {
	Type    string
	message string
}

func (e *InitiateOidcLoginError) Error() string {
	return e.message
}

// InitiateOidcLoginErrorClientNotFound represents the ClientNotFound error variant
type InitiateOidcLoginErrorClientNotFound struct {
	*InitiateOidcLoginError
}

func newInitiateOidcLoginErrorClientNotFound() *InitiateOidcLoginErrorClientNotFound {
	return &InitiateOidcLoginErrorClientNotFound{
		InitiateOidcLoginError: &InitiateOidcLoginError{
			Type:    "ClientNotFound",
			message: "ClientNotFound",
		},
	}
}

// InitiateOidcLoginErrorRedirectURLInvalid represents the RedirectUrlInvalid error variant
type InitiateOidcLoginErrorRedirectURLInvalid struct {
	*InitiateOidcLoginError
	Details InitiateOidcLoginErrorRedirectUrlInvalidDetails
}

func newInitiateOidcLoginErrorRedirectURLInvalid(details InitiateOidcLoginErrorRedirectUrlInvalidDetails) *InitiateOidcLoginErrorRedirectURLInvalid {
	return &InitiateOidcLoginErrorRedirectURLInvalid{
		InitiateOidcLoginError: &InitiateOidcLoginError{
			Type:    "RedirectUrlInvalid",
			message: "RedirectUrlInvalid",
		},
		Details: details,
	}
}

// InitiateOidcLoginErrorUnexpectedError represents the UnexpectedError error variant
type InitiateOidcLoginErrorUnexpectedError struct {
	*InitiateOidcLoginError
	Details string
}

func newInitiateOidcLoginErrorUnexpectedError(details string) *InitiateOidcLoginErrorUnexpectedError {
	return &InitiateOidcLoginErrorUnexpectedError{
		InitiateOidcLoginError: &InitiateOidcLoginError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// IntegrationKeyError is the base error type
type IntegrationKeyError struct {
	Type    string
	message string
}

func (e *IntegrationKeyError) Error() string {
	return e.message
}

// IntegrationKeyErrorInvalidPrefix represents the InvalidPrefix error variant
type IntegrationKeyErrorInvalidPrefix struct {
	*IntegrationKeyError
	Details string
}

func newIntegrationKeyErrorInvalidPrefix(details string) *IntegrationKeyErrorInvalidPrefix {
	return &IntegrationKeyErrorInvalidPrefix{
		IntegrationKeyError: &IntegrationKeyError{
			Type:    "InvalidPrefix",
			message: "InvalidPrefix",
		},
		Details: details,
	}
}

// IntegrationKeyErrorIntegrationKeyNotFound represents the IntegrationKeyNotFound error variant
type IntegrationKeyErrorIntegrationKeyNotFound struct {
	*IntegrationKeyError
}

func newIntegrationKeyErrorIntegrationKeyNotFound() *IntegrationKeyErrorIntegrationKeyNotFound {
	return &IntegrationKeyErrorIntegrationKeyNotFound{
		IntegrationKeyError: &IntegrationKeyError{
			Type:    "IntegrationKeyNotFound",
			message: "IntegrationKeyNotFound",
		},
	}
}

// IntegrationKeyErrorNoIntegrationKeyFoundInHeader represents the NoIntegrationKeyFoundInHeader error variant
type IntegrationKeyErrorNoIntegrationKeyFoundInHeader struct {
	*IntegrationKeyError
}

func newIntegrationKeyErrorNoIntegrationKeyFoundInHeader() *IntegrationKeyErrorNoIntegrationKeyFoundInHeader {
	return &IntegrationKeyErrorNoIntegrationKeyFoundInHeader{
		IntegrationKeyError: &IntegrationKeyError{
			Type:    "NoIntegrationKeyFoundInHeader",
			message: "NoIntegrationKeyFoundInHeader",
		},
	}
}

// IntegrationKeyErrorCommandNotAllowed represents the CommandNotAllowed error variant
type IntegrationKeyErrorCommandNotAllowed struct {
	*IntegrationKeyError
	Details IntegrationKeyErrorCommandNotAllowedDetails
}

func newIntegrationKeyErrorCommandNotAllowed(details IntegrationKeyErrorCommandNotAllowedDetails) *IntegrationKeyErrorCommandNotAllowed {
	return &IntegrationKeyErrorCommandNotAllowed{
		IntegrationKeyError: &IntegrationKeyError{
			Type:    "CommandNotAllowed",
			message: "CommandNotAllowed",
		},
		Details: details,
	}
}

// IntegrationKeyErrorUnexpectedError represents the UnexpectedError error variant
type IntegrationKeyErrorUnexpectedError struct {
	*IntegrationKeyError
	Details string
}

func newIntegrationKeyErrorUnexpectedError(details string) *IntegrationKeyErrorUnexpectedError {
	return &IntegrationKeyErrorUnexpectedError{
		IntegrationKeyError: &IntegrationKeyError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// InvalidTagError represents a generated type
type InvalidTagError struct {
	InvalidTagFormat string `json:"InvalidTagFormat"`
}



func (e *InvalidTagError) Error() string {
	return fmt.Sprintf("InvalidTagError: %v", e)
}


// InvalidateAllImpersonationSessionsForEmployeeError represents a generated type
type InvalidateAllImpersonationSessionsForEmployeeError struct {
	Type string `json:"type"`
	Details string `json:"details"`
}



func (e *InvalidateAllImpersonationSessionsForEmployeeError) Error() string {
	return fmt.Sprintf("InvalidateAllImpersonationSessionsForEmployeeError: %v", e)
}


// InvalidateAllImpersonationSessionsForUserError represents a generated type
type InvalidateAllImpersonationSessionsForUserError struct {
	Type string `json:"type"`
	Details string `json:"details"`
}



func (e *InvalidateAllImpersonationSessionsForUserError) Error() string {
	return fmt.Sprintf("InvalidateAllImpersonationSessionsForUserError: %v", e)
}


// InvalidateAllSessionsForUserError represents a generated type
type InvalidateAllSessionsForUserError struct {
	Type string `json:"type"`
	Details string `json:"details"`
}



func (e *InvalidateAllSessionsForUserError) Error() string {
	return fmt.Sprintf("InvalidateAllSessionsForUserError: %v", e)
}


// InvalidateAllSessionsForUserExceptOneError represents a generated type
type InvalidateAllSessionsForUserExceptOneError struct {
	Type string `json:"type"`
	Details string `json:"details"`
}



func (e *InvalidateAllSessionsForUserExceptOneError) Error() string {
	return fmt.Sprintf("InvalidateAllSessionsForUserExceptOneError: %v", e)
}


// InvalidateImpersonationSessionByIdError is the base error type
type InvalidateImpersonationSessionByIdError struct {
	Type    string
	message string
}

func (e *InvalidateImpersonationSessionByIdError) Error() string {
	return e.message
}

// InvalidateImpersonationSessionByIdErrorSessionNotFound represents the SessionNotFound error variant
type InvalidateImpersonationSessionByIdErrorSessionNotFound struct {
	*InvalidateImpersonationSessionByIdError
}

func newInvalidateImpersonationSessionByIdErrorSessionNotFound() *InvalidateImpersonationSessionByIdErrorSessionNotFound {
	return &InvalidateImpersonationSessionByIdErrorSessionNotFound{
		InvalidateImpersonationSessionByIdError: &InvalidateImpersonationSessionByIdError{
			Type:    "SessionNotFound",
			message: "SessionNotFound",
		},
	}
}

// InvalidateImpersonationSessionByIdErrorUnexpectedError represents the UnexpectedError error variant
type InvalidateImpersonationSessionByIdErrorUnexpectedError struct {
	*InvalidateImpersonationSessionByIdError
	Details string
}

func newInvalidateImpersonationSessionByIdErrorUnexpectedError(details string) *InvalidateImpersonationSessionByIdErrorUnexpectedError {
	return &InvalidateImpersonationSessionByIdErrorUnexpectedError{
		InvalidateImpersonationSessionByIdError: &InvalidateImpersonationSessionByIdError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// InvalidateImpersonationSessionByTokenError is the base error type
type InvalidateImpersonationSessionByTokenError struct {
	Type    string
	message string
}

func (e *InvalidateImpersonationSessionByTokenError) Error() string {
	return e.message
}

// InvalidateImpersonationSessionByTokenErrorSessionNotFound represents the SessionNotFound error variant
type InvalidateImpersonationSessionByTokenErrorSessionNotFound struct {
	*InvalidateImpersonationSessionByTokenError
}

func newInvalidateImpersonationSessionByTokenErrorSessionNotFound() *InvalidateImpersonationSessionByTokenErrorSessionNotFound {
	return &InvalidateImpersonationSessionByTokenErrorSessionNotFound{
		InvalidateImpersonationSessionByTokenError: &InvalidateImpersonationSessionByTokenError{
			Type:    "SessionNotFound",
			message: "SessionNotFound",
		},
	}
}

// InvalidateImpersonationSessionByTokenErrorUnexpectedError represents the UnexpectedError error variant
type InvalidateImpersonationSessionByTokenErrorUnexpectedError struct {
	*InvalidateImpersonationSessionByTokenError
	Details string
}

func newInvalidateImpersonationSessionByTokenErrorUnexpectedError(details string) *InvalidateImpersonationSessionByTokenErrorUnexpectedError {
	return &InvalidateImpersonationSessionByTokenErrorUnexpectedError{
		InvalidateImpersonationSessionByTokenError: &InvalidateImpersonationSessionByTokenError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// InvalidateSessionByIdError is the base error type
type InvalidateSessionByIdError struct {
	Type    string
	message string
}

func (e *InvalidateSessionByIdError) Error() string {
	return e.message
}

// InvalidateSessionByIdErrorSessionNotFound represents the SessionNotFound error variant
type InvalidateSessionByIdErrorSessionNotFound struct {
	*InvalidateSessionByIdError
}

func newInvalidateSessionByIdErrorSessionNotFound() *InvalidateSessionByIdErrorSessionNotFound {
	return &InvalidateSessionByIdErrorSessionNotFound{
		InvalidateSessionByIdError: &InvalidateSessionByIdError{
			Type:    "SessionNotFound",
			message: "SessionNotFound",
		},
	}
}

// InvalidateSessionByIdErrorUnexpectedError represents the UnexpectedError error variant
type InvalidateSessionByIdErrorUnexpectedError struct {
	*InvalidateSessionByIdError
	Details string
}

func newInvalidateSessionByIdErrorUnexpectedError(details string) *InvalidateSessionByIdErrorUnexpectedError {
	return &InvalidateSessionByIdErrorUnexpectedError{
		InvalidateSessionByIdError: &InvalidateSessionByIdError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// InvalidateSessionByTokenError represents a generated type
type InvalidateSessionByTokenError struct {
	Type string `json:"type"`
	Details string `json:"details"`
}



func (e *InvalidateSessionByTokenError) Error() string {
	return fmt.Sprintf("InvalidateSessionByTokenError: %v", e)
}


// LicenseCheckError represents an enumeration type
type LicenseCheckError string

const (
	LicenseCheckErrorInvalidLicense LicenseCheckError = "InvalidLicense"
)


// LicenseFeatureLimitError represents an enumeration type
type LicenseFeatureLimitError string

const (
	LicenseFeatureLimitErrorNoAccessToFeature LicenseFeatureLimitError = "NoAccessToFeature"
	LicenseFeatureLimitErrorLimitReached LicenseFeatureLimitError = "LimitReached"
	LicenseFeatureLimitErrorAccessLimitedToLowerTier LicenseFeatureLimitError = "AccessLimitedToLowerTier"
	LicenseFeatureLimitErrorUnexpectedError LicenseFeatureLimitError = "UnexpectedError"
)


// LinkScimUserError is the base error type
type LinkScimUserError struct {
	Type    string
	message string
}

func (e *LinkScimUserError) Error() string {
	return e.message
}

// LinkScimUserErrorScimConnectionNotFound represents the ScimConnectionNotFound error variant
type LinkScimUserErrorScimConnectionNotFound struct {
	*LinkScimUserError
}

func newLinkScimUserErrorScimConnectionNotFound() *LinkScimUserErrorScimConnectionNotFound {
	return &LinkScimUserErrorScimConnectionNotFound{
		LinkScimUserError: &LinkScimUserError{
			Type:    "ScimConnectionNotFound",
			message: "ScimConnectionNotFound",
		},
	}
}

// LinkScimUserErrorUserNotFound represents the UserNotFound error variant
type LinkScimUserErrorUserNotFound struct {
	*LinkScimUserError
}

func newLinkScimUserErrorUserNotFound() *LinkScimUserErrorUserNotFound {
	return &LinkScimUserErrorUserNotFound{
		LinkScimUserError: &LinkScimUserError{
			Type:    "UserNotFound",
			message: "UserNotFound",
		},
	}
}

// LinkScimUserErrorUnexpectedError represents the UnexpectedError error variant
type LinkScimUserErrorUnexpectedError struct {
	*LinkScimUserError
	Details string
}

func newLinkScimUserErrorUnexpectedError(details string) *LinkScimUserErrorUnexpectedError {
	return &LinkScimUserErrorUnexpectedError{
		LinkScimUserError: &LinkScimUserError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// OidcLoginRequestError represents an enumeration type
type OidcLoginRequestError string

const (
	OidcLoginRequestErrorInvalidCsrfState OidcLoginRequestError = "InvalidCsrfState"
)


// PatchOidcClientError is the base error type
type PatchOidcClientError struct {
	Type    string
	message string
}

func (e *PatchOidcClientError) Error() string {
	return e.message
}

// PatchOidcClientErrorOIDCClientNotFound represents the OidcClientNotFound error variant
type PatchOidcClientErrorOIDCClientNotFound struct {
	*PatchOidcClientError
}

func newPatchOidcClientErrorOIDCClientNotFound() *PatchOidcClientErrorOIDCClientNotFound {
	return &PatchOidcClientErrorOIDCClientNotFound{
		PatchOidcClientError: &PatchOidcClientError{
			Type:    "OidcClientNotFound",
			message: "OidcClientNotFound",
		},
	}
}

// PatchOidcClientErrorInvalidFields represents the InvalidFields error variant
type PatchOidcClientErrorInvalidFields struct {
	*PatchOidcClientError
	Details PatchOidcClientErrorInvalidFieldsDetails
}

func newPatchOidcClientErrorInvalidFields(details PatchOidcClientErrorInvalidFieldsDetails) *PatchOidcClientErrorInvalidFields {
	return &PatchOidcClientErrorInvalidFields{
		PatchOidcClientError: &PatchOidcClientError{
			Type:    "InvalidFields",
			message: "InvalidFields",
		},
		Details: details,
	}
}

// PatchOidcClientErrorUnexpectedError represents the UnexpectedError error variant
type PatchOidcClientErrorUnexpectedError struct {
	*PatchOidcClientError
	Details string
}

func newPatchOidcClientErrorUnexpectedError(details string) *PatchOidcClientErrorUnexpectedError {
	return &PatchOidcClientErrorUnexpectedError{
		PatchOidcClientError: &PatchOidcClientError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// PatchScimConnectionError is the base error type
type PatchScimConnectionError struct {
	Type    string
	message string
}

func (e *PatchScimConnectionError) Error() string {
	return e.message
}

// PatchScimConnectionErrorScimConnectionNotFound represents the ScimConnectionNotFound error variant
type PatchScimConnectionErrorScimConnectionNotFound struct {
	*PatchScimConnectionError
}

func newPatchScimConnectionErrorScimConnectionNotFound() *PatchScimConnectionErrorScimConnectionNotFound {
	return &PatchScimConnectionErrorScimConnectionNotFound{
		PatchScimConnectionError: &PatchScimConnectionError{
			Type:    "ScimConnectionNotFound",
			message: "ScimConnectionNotFound",
		},
	}
}

// PatchScimConnectionErrorDisplayNameInvalid represents the DisplayNameInvalid error variant
type PatchScimConnectionErrorDisplayNameInvalid struct {
	*PatchScimConnectionError
	Details PatchScimConnectionErrorDisplayNameInvalidDetails
}

func newPatchScimConnectionErrorDisplayNameInvalid(details PatchScimConnectionErrorDisplayNameInvalidDetails) *PatchScimConnectionErrorDisplayNameInvalid {
	return &PatchScimConnectionErrorDisplayNameInvalid{
		PatchScimConnectionError: &PatchScimConnectionError{
			Type:    "DisplayNameInvalid",
			message: "DisplayNameInvalid",
		},
		Details: details,
	}
}

// PatchScimConnectionErrorUnexpectedError represents the UnexpectedError error variant
type PatchScimConnectionErrorUnexpectedError struct {
	*PatchScimConnectionError
	Details string
}

func newPatchScimConnectionErrorUnexpectedError(details string) *PatchScimConnectionErrorUnexpectedError {
	return &PatchScimConnectionErrorUnexpectedError{
		PatchScimConnectionError: &PatchScimConnectionError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// PingError represents a generated type
type PingError struct {
	Type string `json:"type"`
	Details string `json:"details"`
}



func (e *PingError) Error() string {
	return fmt.Sprintf("PingError: %v", e)
}


// RegisterDeviceError is the base error type
type RegisterDeviceError struct {
	Type    string
	message string
}

func (e *RegisterDeviceError) Error() string {
	return e.message
}

// RegisterDeviceErrorSessionNotFound represents the SessionNotFound error variant
type RegisterDeviceErrorSessionNotFound struct {
	*RegisterDeviceError
}

func newRegisterDeviceErrorSessionNotFound() *RegisterDeviceErrorSessionNotFound {
	return &RegisterDeviceErrorSessionNotFound{
		RegisterDeviceError: &RegisterDeviceError{
			Type:    "SessionNotFound",
			message: "SessionNotFound",
		},
	}
}

// RegisterDeviceErrorNewDeviceChallengeRequired represents the NewDeviceChallengeRequired error variant
type RegisterDeviceErrorNewDeviceChallengeRequired struct {
	*RegisterDeviceError
	Details RegisterDeviceErrorNewDeviceChallengeRequiredDetails
}

func newRegisterDeviceErrorNewDeviceChallengeRequired(details RegisterDeviceErrorNewDeviceChallengeRequiredDetails) *RegisterDeviceErrorNewDeviceChallengeRequired {
	return &RegisterDeviceErrorNewDeviceChallengeRequired{
		RegisterDeviceError: &RegisterDeviceError{
			Type:    "NewDeviceChallengeRequired",
			message: "NewDeviceChallengeRequired",
		},
		Details: details,
	}
}

// RegisterDeviceErrorInvalidDeviceRegistration represents the InvalidDeviceRegistration error variant
type RegisterDeviceErrorInvalidDeviceRegistration struct {
	*RegisterDeviceError
}

func newRegisterDeviceErrorInvalidDeviceRegistration() *RegisterDeviceErrorInvalidDeviceRegistration {
	return &RegisterDeviceErrorInvalidDeviceRegistration{
		RegisterDeviceError: &RegisterDeviceError{
			Type:    "InvalidDeviceRegistration",
			message: "InvalidDeviceRegistration",
		},
	}
}

// RegisterDeviceErrorDeviceAlreadyRegistered represents the DeviceAlreadyRegistered error variant
type RegisterDeviceErrorDeviceAlreadyRegistered struct {
	*RegisterDeviceError
}

func newRegisterDeviceErrorDeviceAlreadyRegistered() *RegisterDeviceErrorDeviceAlreadyRegistered {
	return &RegisterDeviceErrorDeviceAlreadyRegistered{
		RegisterDeviceError: &RegisterDeviceError{
			Type:    "DeviceAlreadyRegistered",
			message: "DeviceAlreadyRegistered",
		},
	}
}

// RegisterDeviceErrorUnexpectedError represents the UnexpectedError error variant
type RegisterDeviceErrorUnexpectedError struct {
	*RegisterDeviceError
	Details string
}

func newRegisterDeviceErrorUnexpectedError(details string) *RegisterDeviceErrorUnexpectedError {
	return &RegisterDeviceErrorUnexpectedError{
		RegisterDeviceError: &RegisterDeviceError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// ResetScimApiKeyError is the base error type
type ResetScimApiKeyError struct {
	Type    string
	message string
}

func (e *ResetScimApiKeyError) Error() string {
	return e.message
}

// ResetScimApiKeyErrorScimConnectionNotFound represents the ScimConnectionNotFound error variant
type ResetScimApiKeyErrorScimConnectionNotFound struct {
	*ResetScimApiKeyError
}

func newResetScimApiKeyErrorScimConnectionNotFound() *ResetScimApiKeyErrorScimConnectionNotFound {
	return &ResetScimApiKeyErrorScimConnectionNotFound{
		ResetScimApiKeyError: &ResetScimApiKeyError{
			Type:    "ScimConnectionNotFound",
			message: "ScimConnectionNotFound",
		},
	}
}

// ResetScimApiKeyErrorUnexpectedError represents the UnexpectedError error variant
type ResetScimApiKeyErrorUnexpectedError struct {
	*ResetScimApiKeyError
	Details string
}

func newResetScimApiKeyErrorUnexpectedError(details string) *ResetScimApiKeyErrorUnexpectedError {
	return &ResetScimApiKeyErrorUnexpectedError{
		ResetScimApiKeyError: &ResetScimApiKeyError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// RotateStatelessTokenKeyError is the base error type
type RotateStatelessTokenKeyError struct {
	Type    string
	message string
}

func (e *RotateStatelessTokenKeyError) Error() string {
	return e.message
}

// RotateStatelessTokenKeyErrorRotationFailed represents the RotationFailed error variant
type RotateStatelessTokenKeyErrorRotationFailed struct {
	*RotateStatelessTokenKeyError
	Details string
}

func newRotateStatelessTokenKeyErrorRotationFailed(details string) *RotateStatelessTokenKeyErrorRotationFailed {
	return &RotateStatelessTokenKeyErrorRotationFailed{
		RotateStatelessTokenKeyError: &RotateStatelessTokenKeyError{
			Type:    "RotationFailed",
			message: "RotationFailed",
		},
		Details: details,
	}
}

// RotateStatelessTokenKeyErrorInvalidParameters represents the InvalidParameters error variant
type RotateStatelessTokenKeyErrorInvalidParameters struct {
	*RotateStatelessTokenKeyError
	Details string
}

func newRotateStatelessTokenKeyErrorInvalidParameters(details string) *RotateStatelessTokenKeyErrorInvalidParameters {
	return &RotateStatelessTokenKeyErrorInvalidParameters{
		RotateStatelessTokenKeyError: &RotateStatelessTokenKeyError{
			Type:    "InvalidParameters",
			message: "InvalidParameters",
		},
		Details: details,
	}
}

// RotateStatelessTokenKeyErrorUnexpectedError represents the UnexpectedError error variant
type RotateStatelessTokenKeyErrorUnexpectedError struct {
	*RotateStatelessTokenKeyError
	Details string
}

func newRotateStatelessTokenKeyErrorUnexpectedError(details string) *RotateStatelessTokenKeyErrorUnexpectedError {
	return &RotateStatelessTokenKeyErrorUnexpectedError{
		RotateStatelessTokenKeyError: &RotateStatelessTokenKeyError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// ScimClientFacingError represents a generated type
type ScimClientFacingError struct {
	StatusToReturn int `json:"statusToReturn"`
	BodyToReturn json.RawMessage `json:"bodyToReturn"`
	UnderlyingError ScimUnderlyingError `json:"underlyingError"`
}



func (e *ScimClientFacingError) Error() string {
	return fmt.Sprintf("ScimClientFacingError: %v", e)
}


// ScimUnderlyingError is the base error type
type ScimUnderlyingError struct {
	Type    string
	message string
}

func (e *ScimUnderlyingError) Error() string {
	return e.message
}

// ScimUnderlyingErrorInvalidAPIKey represents the InvalidApiKey error variant
type ScimUnderlyingErrorInvalidAPIKey struct {
	*ScimUnderlyingError
}

func newScimUnderlyingErrorInvalidAPIKey() *ScimUnderlyingErrorInvalidAPIKey {
	return &ScimUnderlyingErrorInvalidAPIKey{
		ScimUnderlyingError: &ScimUnderlyingError{
			Type:    "InvalidApiKey",
			message: "InvalidApiKey",
		},
	}
}

// ScimUnderlyingErrorInvalidPath represents the InvalidPath error variant
type ScimUnderlyingErrorInvalidPath struct {
	*ScimUnderlyingError
}

func newScimUnderlyingErrorInvalidPath() *ScimUnderlyingErrorInvalidPath {
	return &ScimUnderlyingErrorInvalidPath{
		ScimUnderlyingError: &ScimUnderlyingError{
			Type:    "InvalidPath",
			message: "InvalidPath",
		},
	}
}

// ScimUnderlyingErrorUserNotFound represents the UserNotFound error variant
type ScimUnderlyingErrorUserNotFound struct {
	*ScimUnderlyingError
}

func newScimUnderlyingErrorUserNotFound() *ScimUnderlyingErrorUserNotFound {
	return &ScimUnderlyingErrorUserNotFound{
		ScimUnderlyingError: &ScimUnderlyingError{
			Type:    "UserNotFound",
			message: "UserNotFound",
		},
	}
}

// ScimUnderlyingErrorGroupNotFound represents the GroupNotFound error variant
type ScimUnderlyingErrorGroupNotFound struct {
	*ScimUnderlyingError
}

func newScimUnderlyingErrorGroupNotFound() *ScimUnderlyingErrorGroupNotFound {
	return &ScimUnderlyingErrorGroupNotFound{
		ScimUnderlyingError: &ScimUnderlyingError{
			Type:    "GroupNotFound",
			message: "GroupNotFound",
		},
	}
}

// ScimUnderlyingErrorMissingRequiredField represents the MissingRequiredField error variant
type ScimUnderlyingErrorMissingRequiredField struct {
	*ScimUnderlyingError
	Details ScimUnderlyingErrorMissingRequiredFieldDetails
}

func newScimUnderlyingErrorMissingRequiredField(details ScimUnderlyingErrorMissingRequiredFieldDetails) *ScimUnderlyingErrorMissingRequiredField {
	return &ScimUnderlyingErrorMissingRequiredField{
		ScimUnderlyingError: &ScimUnderlyingError{
			Type:    "MissingRequiredField",
			message: "MissingRequiredField",
		},
		Details: details,
	}
}

// ScimUnderlyingErrorCantRemoveRequiredField represents the CantRemoveRequiredField error variant
type ScimUnderlyingErrorCantRemoveRequiredField struct {
	*ScimUnderlyingError
	Details ScimUnderlyingErrorCantRemoveRequiredFieldDetails
}

func newScimUnderlyingErrorCantRemoveRequiredField(details ScimUnderlyingErrorCantRemoveRequiredFieldDetails) *ScimUnderlyingErrorCantRemoveRequiredField {
	return &ScimUnderlyingErrorCantRemoveRequiredField{
		ScimUnderlyingError: &ScimUnderlyingError{
			Type:    "CantRemoveRequiredField",
			message: "CantRemoveRequiredField",
		},
		Details: details,
	}
}

// ScimUnderlyingErrorScimUserIDAlreadyTaken represents the ScimUserIdAlreadyTaken error variant
type ScimUnderlyingErrorScimUserIDAlreadyTaken struct {
	*ScimUnderlyingError
	Details ScimUnderlyingErrorScimUserIdAlreadyTakenDetails
}

func newScimUnderlyingErrorScimUserIDAlreadyTaken(details ScimUnderlyingErrorScimUserIdAlreadyTakenDetails) *ScimUnderlyingErrorScimUserIDAlreadyTaken {
	return &ScimUnderlyingErrorScimUserIDAlreadyTaken{
		ScimUnderlyingError: &ScimUnderlyingError{
			Type:    "ScimUserIdAlreadyTaken",
			message: "ScimUserIdAlreadyTaken",
		},
		Details: details,
	}
}

// ScimUnderlyingErrorInvalidBody represents the InvalidBody error variant
type ScimUnderlyingErrorInvalidBody struct {
	*ScimUnderlyingError
}

func newScimUnderlyingErrorInvalidBody() *ScimUnderlyingErrorInvalidBody {
	return &ScimUnderlyingErrorInvalidBody{
		ScimUnderlyingError: &ScimUnderlyingError{
			Type:    "InvalidBody",
			message: "InvalidBody",
		},
	}
}

// ScimUnderlyingErrorInvalidQueryField represents the InvalidQueryField error variant
type ScimUnderlyingErrorInvalidQueryField struct {
	*ScimUnderlyingError
	Details ScimUnderlyingErrorInvalidQueryFieldDetails
}

func newScimUnderlyingErrorInvalidQueryField(details ScimUnderlyingErrorInvalidQueryFieldDetails) *ScimUnderlyingErrorInvalidQueryField {
	return &ScimUnderlyingErrorInvalidQueryField{
		ScimUnderlyingError: &ScimUnderlyingError{
			Type:    "InvalidQueryField",
			message: "InvalidQueryField",
		},
		Details: details,
	}
}

// ScimUnderlyingErrorInvalidPatchPath represents the InvalidPatchPath error variant
type ScimUnderlyingErrorInvalidPatchPath struct {
	*ScimUnderlyingError
}

func newScimUnderlyingErrorInvalidPatchPath() *ScimUnderlyingErrorInvalidPatchPath {
	return &ScimUnderlyingErrorInvalidPatchPath{
		ScimUnderlyingError: &ScimUnderlyingError{
			Type:    "InvalidPatchPath",
			message: "InvalidPatchPath",
		},
	}
}

// ScimUnderlyingErrorInvalidPatchOperation represents the InvalidPatchOperation error variant
type ScimUnderlyingErrorInvalidPatchOperation struct {
	*ScimUnderlyingError
}

func newScimUnderlyingErrorInvalidPatchOperation() *ScimUnderlyingErrorInvalidPatchOperation {
	return &ScimUnderlyingErrorInvalidPatchOperation{
		ScimUnderlyingError: &ScimUnderlyingError{
			Type:    "InvalidPatchOperation",
			message: "InvalidPatchOperation",
		},
	}
}

// ScimUnderlyingErrorInvalidPatchValue represents the InvalidPatchValue error variant
type ScimUnderlyingErrorInvalidPatchValue struct {
	*ScimUnderlyingError
	Details ScimUnderlyingErrorInvalidPatchValueDetails
}

func newScimUnderlyingErrorInvalidPatchValue(details ScimUnderlyingErrorInvalidPatchValueDetails) *ScimUnderlyingErrorInvalidPatchValue {
	return &ScimUnderlyingErrorInvalidPatchValue{
		ScimUnderlyingError: &ScimUnderlyingError{
			Type:    "InvalidPatchValue",
			message: "InvalidPatchValue",
		},
		Details: details,
	}
}

// ScimUnderlyingErrorInvalidSchema represents the InvalidSchema error variant
type ScimUnderlyingErrorInvalidSchema struct {
	*ScimUnderlyingError
}

func newScimUnderlyingErrorInvalidSchema() *ScimUnderlyingErrorInvalidSchema {
	return &ScimUnderlyingErrorInvalidSchema{
		ScimUnderlyingError: &ScimUnderlyingError{
			Type:    "InvalidSchema",
			message: "InvalidSchema",
		},
	}
}

// ScimUnderlyingErrorUnexpectedError represents the UnexpectedError error variant
type ScimUnderlyingErrorUnexpectedError struct {
	*ScimUnderlyingError
	Details string
}

func newScimUnderlyingErrorUnexpectedError(details string) *ScimUnderlyingErrorUnexpectedError {
	return &ScimUnderlyingErrorUnexpectedError{
		ScimUnderlyingError: &ScimUnderlyingError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// StartPasskeyAuthenticationError is the base error type
type StartPasskeyAuthenticationError struct {
	Type    string
	message string
}

func (e *StartPasskeyAuthenticationError) Error() string {
	return e.message
}

// StartPasskeyAuthenticationErrorCannotParseAdditionalAllowedOrigin represents the CannotParseAdditionalAllowedOrigin error variant
type StartPasskeyAuthenticationErrorCannotParseAdditionalAllowedOrigin struct {
	*StartPasskeyAuthenticationError
}

func newStartPasskeyAuthenticationErrorCannotParseAdditionalAllowedOrigin() *StartPasskeyAuthenticationErrorCannotParseAdditionalAllowedOrigin {
	return &StartPasskeyAuthenticationErrorCannotParseAdditionalAllowedOrigin{
		StartPasskeyAuthenticationError: &StartPasskeyAuthenticationError{
			Type:    "CannotParseAdditionalAllowedOrigin",
			message: "CannotParseAdditionalAllowedOrigin",
		},
	}
}

// StartPasskeyAuthenticationErrorNoPasskeysRegisteredForUser represents the NoPasskeysRegisteredForUser error variant
type StartPasskeyAuthenticationErrorNoPasskeysRegisteredForUser struct {
	*StartPasskeyAuthenticationError
}

func newStartPasskeyAuthenticationErrorNoPasskeysRegisteredForUser() *StartPasskeyAuthenticationErrorNoPasskeysRegisteredForUser {
	return &StartPasskeyAuthenticationErrorNoPasskeysRegisteredForUser{
		StartPasskeyAuthenticationError: &StartPasskeyAuthenticationError{
			Type:    "NoPasskeysRegisteredForUser",
			message: "NoPasskeysRegisteredForUser",
		},
	}
}

// StartPasskeyAuthenticationErrorUnexpectedError represents the UnexpectedError error variant
type StartPasskeyAuthenticationErrorUnexpectedError struct {
	*StartPasskeyAuthenticationError
	Details string
}

func newStartPasskeyAuthenticationErrorUnexpectedError(details string) *StartPasskeyAuthenticationErrorUnexpectedError {
	return &StartPasskeyAuthenticationErrorUnexpectedError{
		StartPasskeyAuthenticationError: &StartPasskeyAuthenticationError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// StartPasskeyRegistrationError is the base error type
type StartPasskeyRegistrationError struct {
	Type    string
	message string
}

func (e *StartPasskeyRegistrationError) Error() string {
	return e.message
}

// StartPasskeyRegistrationErrorCannotParseAdditionalAllowedOrigin represents the CannotParseAdditionalAllowedOrigin error variant
type StartPasskeyRegistrationErrorCannotParseAdditionalAllowedOrigin struct {
	*StartPasskeyRegistrationError
}

func newStartPasskeyRegistrationErrorCannotParseAdditionalAllowedOrigin() *StartPasskeyRegistrationErrorCannotParseAdditionalAllowedOrigin {
	return &StartPasskeyRegistrationErrorCannotParseAdditionalAllowedOrigin{
		StartPasskeyRegistrationError: &StartPasskeyRegistrationError{
			Type:    "CannotParseAdditionalAllowedOrigin",
			message: "CannotParseAdditionalAllowedOrigin",
		},
	}
}

// StartPasskeyRegistrationErrorTooManyPasskeys represents the TooManyPasskeys error variant
type StartPasskeyRegistrationErrorTooManyPasskeys struct {
	*StartPasskeyRegistrationError
	Details StartPasskeyRegistrationErrorTooManyPasskeysDetails
}

func newStartPasskeyRegistrationErrorTooManyPasskeys(details StartPasskeyRegistrationErrorTooManyPasskeysDetails) *StartPasskeyRegistrationErrorTooManyPasskeys {
	return &StartPasskeyRegistrationErrorTooManyPasskeys{
		StartPasskeyRegistrationError: &StartPasskeyRegistrationError{
			Type:    "TooManyPasskeys",
			message: "TooManyPasskeys",
		},
		Details: details,
	}
}

// StartPasskeyRegistrationErrorUnexpectedError represents the UnexpectedError error variant
type StartPasskeyRegistrationErrorUnexpectedError struct {
	*StartPasskeyRegistrationError
	Details string
}

func newStartPasskeyRegistrationErrorUnexpectedError(details string) *StartPasskeyRegistrationErrorUnexpectedError {
	return &StartPasskeyRegistrationErrorUnexpectedError{
		StartPasskeyRegistrationError: &StartPasskeyRegistrationError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// UpdateSessionError is the base error type
type UpdateSessionError struct {
	Type    string
	message string
}

func (e *UpdateSessionError) Error() string {
	return e.message
}

// UpdateSessionErrorSessionNotFound represents the SessionNotFound error variant
type UpdateSessionErrorSessionNotFound struct {
	*UpdateSessionError
}

func newUpdateSessionErrorSessionNotFound() *UpdateSessionErrorSessionNotFound {
	return &UpdateSessionErrorSessionNotFound{
		UpdateSessionError: &UpdateSessionError{
			Type:    "SessionNotFound",
			message: "SessionNotFound",
		},
	}
}

// UpdateSessionErrorConflictingMetadataOptions represents the ConflictingMetadataOptions error variant
type UpdateSessionErrorConflictingMetadataOptions struct {
	*UpdateSessionError
}

func newUpdateSessionErrorConflictingMetadataOptions() *UpdateSessionErrorConflictingMetadataOptions {
	return &UpdateSessionErrorConflictingMetadataOptions{
		UpdateSessionError: &UpdateSessionError{
			Type:    "ConflictingMetadataOptions",
			message: "ConflictingMetadataOptions",
		},
	}
}

// UpdateSessionErrorInvalidTagFormat represents the InvalidTagFormat error variant
type UpdateSessionErrorInvalidTagFormat struct {
	*UpdateSessionError
	Details InvalidTagError
}

func newUpdateSessionErrorInvalidTagFormat(details InvalidTagError) *UpdateSessionErrorInvalidTagFormat {
	return &UpdateSessionErrorInvalidTagFormat{
		UpdateSessionError: &UpdateSessionError{
			Type:    "InvalidTagFormat",
			message: "InvalidTagFormat",
		},
		Details: details,
	}
}

// UpdateSessionErrorCannotModifyOnCreateOnlyTags represents the CannotModifyOnCreateOnlyTags error variant
type UpdateSessionErrorCannotModifyOnCreateOnlyTags struct {
	*UpdateSessionError
	Details UpdateSessionErrorCannotModifyOnCreateOnlyTagsDetails
}

func newUpdateSessionErrorCannotModifyOnCreateOnlyTags(details UpdateSessionErrorCannotModifyOnCreateOnlyTagsDetails) *UpdateSessionErrorCannotModifyOnCreateOnlyTags {
	return &UpdateSessionErrorCannotModifyOnCreateOnlyTags{
		UpdateSessionError: &UpdateSessionError{
			Type:    "CannotModifyOnCreateOnlyTags",
			message: "CannotModifyOnCreateOnlyTags",
		},
		Details: details,
	}
}

// UpdateSessionErrorUnexpectedError represents the UnexpectedError error variant
type UpdateSessionErrorUnexpectedError struct {
	*UpdateSessionError
	Details string
}

func newUpdateSessionErrorUnexpectedError(details string) *UpdateSessionErrorUnexpectedError {
	return &UpdateSessionErrorUnexpectedError{
		UpdateSessionError: &UpdateSessionError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// UpdateSessionsError is the base error type
type UpdateSessionsError struct {
	Type    string
	message string
}

func (e *UpdateSessionsError) Error() string {
	return e.message
}

// UpdateSessionsErrorConflictingMetadataOptions represents the ConflictingMetadataOptions error variant
type UpdateSessionsErrorConflictingMetadataOptions struct {
	*UpdateSessionsError
}

func newUpdateSessionsErrorConflictingMetadataOptions() *UpdateSessionsErrorConflictingMetadataOptions {
	return &UpdateSessionsErrorConflictingMetadataOptions{
		UpdateSessionsError: &UpdateSessionsError{
			Type:    "ConflictingMetadataOptions",
			message: "ConflictingMetadataOptions",
		},
	}
}

// UpdateSessionsErrorInvalidTagFormat represents the InvalidTagFormat error variant
type UpdateSessionsErrorInvalidTagFormat struct {
	*UpdateSessionsError
	Details InvalidTagError
}

func newUpdateSessionsErrorInvalidTagFormat(details InvalidTagError) *UpdateSessionsErrorInvalidTagFormat {
	return &UpdateSessionsErrorInvalidTagFormat{
		UpdateSessionsError: &UpdateSessionsError{
			Type:    "InvalidTagFormat",
			message: "InvalidTagFormat",
		},
		Details: details,
	}
}

// UpdateSessionsErrorCannotModifyOnCreateOnlyTags represents the CannotModifyOnCreateOnlyTags error variant
type UpdateSessionsErrorCannotModifyOnCreateOnlyTags struct {
	*UpdateSessionsError
	Details UpdateSessionsErrorCannotModifyOnCreateOnlyTagsDetails
}

func newUpdateSessionsErrorCannotModifyOnCreateOnlyTags(details UpdateSessionsErrorCannotModifyOnCreateOnlyTagsDetails) *UpdateSessionsErrorCannotModifyOnCreateOnlyTags {
	return &UpdateSessionsErrorCannotModifyOnCreateOnlyTags{
		UpdateSessionsError: &UpdateSessionsError{
			Type:    "CannotModifyOnCreateOnlyTags",
			message: "CannotModifyOnCreateOnlyTags",
		},
		Details: details,
	}
}

// UpdateSessionsErrorUpdatingTooManySessionsAtOnce represents the UpdatingTooManySessionsAtOnce error variant
type UpdateSessionsErrorUpdatingTooManySessionsAtOnce struct {
	*UpdateSessionsError
}

func newUpdateSessionsErrorUpdatingTooManySessionsAtOnce() *UpdateSessionsErrorUpdatingTooManySessionsAtOnce {
	return &UpdateSessionsErrorUpdatingTooManySessionsAtOnce{
		UpdateSessionsError: &UpdateSessionsError{
			Type:    "UpdatingTooManySessionsAtOnce",
			message: "UpdatingTooManySessionsAtOnce",
		},
	}
}

// UpdateSessionsErrorUnexpectedError represents the UnexpectedError error variant
type UpdateSessionsErrorUnexpectedError struct {
	*UpdateSessionsError
	Details string
}

func newUpdateSessionsErrorUnexpectedError(details string) *UpdateSessionsErrorUnexpectedError {
	return &UpdateSessionsErrorUnexpectedError{
		UpdateSessionsError: &UpdateSessionsError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// ValidateAndRefreshSessionError is the base error type
type ValidateAndRefreshSessionError struct {
	Type    string
	message string
}

func (e *ValidateAndRefreshSessionError) Error() string {
	return e.message
}

// ValidateAndRefreshSessionErrorInvalidSessionToken represents the InvalidSessionToken error variant
type ValidateAndRefreshSessionErrorInvalidSessionToken struct {
	*ValidateAndRefreshSessionError
	Details string
}

func newValidateAndRefreshSessionErrorInvalidSessionToken(details string) *ValidateAndRefreshSessionErrorInvalidSessionToken {
	return &ValidateAndRefreshSessionErrorInvalidSessionToken{
		ValidateAndRefreshSessionError: &ValidateAndRefreshSessionError{
			Type:    "InvalidSessionToken",
			message: "InvalidSessionToken",
		},
		Details: details,
	}
}

// ValidateAndRefreshSessionErrorIPAddressError represents the IpAddressError error variant
type ValidateAndRefreshSessionErrorIPAddressError struct {
	*ValidateAndRefreshSessionError
	Details string
}

func newValidateAndRefreshSessionErrorIPAddressError(details string) *ValidateAndRefreshSessionErrorIPAddressError {
	return &ValidateAndRefreshSessionErrorIPAddressError{
		ValidateAndRefreshSessionError: &ValidateAndRefreshSessionError{
			Type:    "IpAddressError",
			message: "IpAddressError",
		},
		Details: details,
	}
}

// ValidateAndRefreshSessionErrorNewDeviceChallengeRequired represents the NewDeviceChallengeRequired error variant
type ValidateAndRefreshSessionErrorNewDeviceChallengeRequired struct {
	*ValidateAndRefreshSessionError
	Details ValidateAndRefreshSessionErrorNewDeviceChallengeRequiredDetails
}

func newValidateAndRefreshSessionErrorNewDeviceChallengeRequired(details ValidateAndRefreshSessionErrorNewDeviceChallengeRequiredDetails) *ValidateAndRefreshSessionErrorNewDeviceChallengeRequired {
	return &ValidateAndRefreshSessionErrorNewDeviceChallengeRequired{
		ValidateAndRefreshSessionError: &ValidateAndRefreshSessionError{
			Type:    "NewDeviceChallengeRequired",
			message: "NewDeviceChallengeRequired",
		},
		Details: details,
	}
}

// ValidateAndRefreshSessionErrorDeviceVerificationRequired represents the DeviceVerificationRequired error variant
type ValidateAndRefreshSessionErrorDeviceVerificationRequired struct {
	*ValidateAndRefreshSessionError
}

func newValidateAndRefreshSessionErrorDeviceVerificationRequired() *ValidateAndRefreshSessionErrorDeviceVerificationRequired {
	return &ValidateAndRefreshSessionErrorDeviceVerificationRequired{
		ValidateAndRefreshSessionError: &ValidateAndRefreshSessionError{
			Type:    "DeviceVerificationRequired",
			message: "DeviceVerificationRequired",
		},
	}
}

// ValidateAndRefreshSessionErrorDeviceVerificationFailed represents the DeviceVerificationFailed error variant
type ValidateAndRefreshSessionErrorDeviceVerificationFailed struct {
	*ValidateAndRefreshSessionError
}

func newValidateAndRefreshSessionErrorDeviceVerificationFailed() *ValidateAndRefreshSessionErrorDeviceVerificationFailed {
	return &ValidateAndRefreshSessionErrorDeviceVerificationFailed{
		ValidateAndRefreshSessionError: &ValidateAndRefreshSessionError{
			Type:    "DeviceVerificationFailed",
			message: "DeviceVerificationFailed",
		},
	}
}

// ValidateAndRefreshSessionErrorUnexpectedError represents the UnexpectedError error variant
type ValidateAndRefreshSessionErrorUnexpectedError struct {
	*ValidateAndRefreshSessionError
	Details string
}

func newValidateAndRefreshSessionErrorUnexpectedError(details string) *ValidateAndRefreshSessionErrorUnexpectedError {
	return &ValidateAndRefreshSessionErrorUnexpectedError{
		ValidateAndRefreshSessionError: &ValidateAndRefreshSessionError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// ValidateImpersonationSessionError is the base error type
type ValidateImpersonationSessionError struct {
	Type    string
	message string
}

func (e *ValidateImpersonationSessionError) Error() string {
	return e.message
}

// ValidateImpersonationSessionErrorImpersonationNotEnabled represents the ImpersonationNotEnabled error variant
type ValidateImpersonationSessionErrorImpersonationNotEnabled struct {
	*ValidateImpersonationSessionError
}

func newValidateImpersonationSessionErrorImpersonationNotEnabled() *ValidateImpersonationSessionErrorImpersonationNotEnabled {
	return &ValidateImpersonationSessionErrorImpersonationNotEnabled{
		ValidateImpersonationSessionError: &ValidateImpersonationSessionError{
			Type:    "ImpersonationNotEnabled",
			message: "ImpersonationNotEnabled",
		},
	}
}

// ValidateImpersonationSessionErrorInvalidImpersonationToken represents the InvalidImpersonationToken error variant
type ValidateImpersonationSessionErrorInvalidImpersonationToken struct {
	*ValidateImpersonationSessionError
	Details string
}

func newValidateImpersonationSessionErrorInvalidImpersonationToken(details string) *ValidateImpersonationSessionErrorInvalidImpersonationToken {
	return &ValidateImpersonationSessionErrorInvalidImpersonationToken{
		ValidateImpersonationSessionError: &ValidateImpersonationSessionError{
			Type:    "InvalidImpersonationToken",
			message: "InvalidImpersonationToken",
		},
		Details: details,
	}
}

// ValidateImpersonationSessionErrorSessionNotFound represents the SessionNotFound error variant
type ValidateImpersonationSessionErrorSessionNotFound struct {
	*ValidateImpersonationSessionError
}

func newValidateImpersonationSessionErrorSessionNotFound() *ValidateImpersonationSessionErrorSessionNotFound {
	return &ValidateImpersonationSessionErrorSessionNotFound{
		ValidateImpersonationSessionError: &ValidateImpersonationSessionError{
			Type:    "SessionNotFound",
			message: "SessionNotFound",
		},
	}
}

// ValidateImpersonationSessionErrorIPAddressMismatch represents the IpAddressMismatch error variant
type ValidateImpersonationSessionErrorIPAddressMismatch struct {
	*ValidateImpersonationSessionError
}

func newValidateImpersonationSessionErrorIPAddressMismatch() *ValidateImpersonationSessionErrorIPAddressMismatch {
	return &ValidateImpersonationSessionErrorIPAddressMismatch{
		ValidateImpersonationSessionError: &ValidateImpersonationSessionError{
			Type:    "IpAddressMismatch",
			message: "IpAddressMismatch",
		},
	}
}

// ValidateImpersonationSessionErrorUserAgentMismatch represents the UserAgentMismatch error variant
type ValidateImpersonationSessionErrorUserAgentMismatch struct {
	*ValidateImpersonationSessionError
}

func newValidateImpersonationSessionErrorUserAgentMismatch() *ValidateImpersonationSessionErrorUserAgentMismatch {
	return &ValidateImpersonationSessionErrorUserAgentMismatch{
		ValidateImpersonationSessionError: &ValidateImpersonationSessionError{
			Type:    "UserAgentMismatch",
			message: "UserAgentMismatch",
		},
	}
}

// ValidateImpersonationSessionErrorUnexpectedError represents the UnexpectedError error variant
type ValidateImpersonationSessionErrorUnexpectedError struct {
	*ValidateImpersonationSessionError
	Details string
}

func newValidateImpersonationSessionErrorUnexpectedError(details string) *ValidateImpersonationSessionErrorUnexpectedError {
	return &ValidateImpersonationSessionErrorUnexpectedError{
		ValidateImpersonationSessionError: &ValidateImpersonationSessionError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// ValidateSessionError is the base error type
type ValidateSessionError struct {
	Type    string
	message string
}

func (e *ValidateSessionError) Error() string {
	return e.message
}

// ValidateSessionErrorInvalidSessionToken represents the InvalidSessionToken error variant
type ValidateSessionErrorInvalidSessionToken struct {
	*ValidateSessionError
	Details string
}

func newValidateSessionErrorInvalidSessionToken(details string) *ValidateSessionErrorInvalidSessionToken {
	return &ValidateSessionErrorInvalidSessionToken{
		ValidateSessionError: &ValidateSessionError{
			Type:    "InvalidSessionToken",
			message: "InvalidSessionToken",
		},
		Details: details,
	}
}

// ValidateSessionErrorIPAddressError represents the IpAddressError error variant
type ValidateSessionErrorIPAddressError struct {
	*ValidateSessionError
	Details string
}

func newValidateSessionErrorIPAddressError(details string) *ValidateSessionErrorIPAddressError {
	return &ValidateSessionErrorIPAddressError{
		ValidateSessionError: &ValidateSessionError{
			Type:    "IpAddressError",
			message: "IpAddressError",
		},
		Details: details,
	}
}

// ValidateSessionErrorNewDeviceChallengeRequired represents the NewDeviceChallengeRequired error variant
type ValidateSessionErrorNewDeviceChallengeRequired struct {
	*ValidateSessionError
	Details ValidateSessionErrorNewDeviceChallengeRequiredDetails
}

func newValidateSessionErrorNewDeviceChallengeRequired(details ValidateSessionErrorNewDeviceChallengeRequiredDetails) *ValidateSessionErrorNewDeviceChallengeRequired {
	return &ValidateSessionErrorNewDeviceChallengeRequired{
		ValidateSessionError: &ValidateSessionError{
			Type:    "NewDeviceChallengeRequired",
			message: "NewDeviceChallengeRequired",
		},
		Details: details,
	}
}

// ValidateSessionErrorDeviceVerificationRequired represents the DeviceVerificationRequired error variant
type ValidateSessionErrorDeviceVerificationRequired struct {
	*ValidateSessionError
}

func newValidateSessionErrorDeviceVerificationRequired() *ValidateSessionErrorDeviceVerificationRequired {
	return &ValidateSessionErrorDeviceVerificationRequired{
		ValidateSessionError: &ValidateSessionError{
			Type:    "DeviceVerificationRequired",
			message: "DeviceVerificationRequired",
		},
	}
}

// ValidateSessionErrorDeviceVerificationFailed represents the DeviceVerificationFailed error variant
type ValidateSessionErrorDeviceVerificationFailed struct {
	*ValidateSessionError
}

func newValidateSessionErrorDeviceVerificationFailed() *ValidateSessionErrorDeviceVerificationFailed {
	return &ValidateSessionErrorDeviceVerificationFailed{
		ValidateSessionError: &ValidateSessionError{
			Type:    "DeviceVerificationFailed",
			message: "DeviceVerificationFailed",
		},
	}
}

// ValidateSessionErrorUnexpectedError represents the UnexpectedError error variant
type ValidateSessionErrorUnexpectedError struct {
	*ValidateSessionError
	Details string
}

func newValidateSessionErrorUnexpectedError(details string) *ValidateSessionErrorUnexpectedError {
	return &ValidateSessionErrorUnexpectedError{
		ValidateSessionError: &ValidateSessionError{
			Type:    "UnexpectedError",
			message: "UnexpectedError",
		},
		Details: details,
	}
}



// PropertyParseError is the base error type
type PropertyParseError struct {
	Type    string
	message string
}

func (e *PropertyParseError) Error() string {
	return e.message
}

// PropertyParseErrorPropertyMissing represents the property_missing error variant
type PropertyParseErrorPropertyMissing struct {
	*PropertyParseError
	OutputField string
	Path string
}

func newPropertyParseErrorPropertyMissing(outputField string, path string) *PropertyParseErrorPropertyMissing {
	return &PropertyParseErrorPropertyMissing{
		PropertyParseError: &PropertyParseError{
			Type:    "property_missing",
			message: "property_missing",
		},
		OutputField: outputField,
		Path: path,
	}
}

// PropertyParseErrorInvalidType represents the invalid_type error variant
type PropertyParseErrorInvalidType struct {
	*PropertyParseError
	OutputField string
	Path string
	ExpectedType string
	ActualValue json.RawMessage
	Example *string
}

func newPropertyParseErrorInvalidType(outputField string, path string, expectedType string, actualValue json.RawMessage, example *string) *PropertyParseErrorInvalidType {
	return &PropertyParseErrorInvalidType{
		PropertyParseError: &PropertyParseError{
			Type:    "invalid_type",
			message: "invalid_type",
		},
		OutputField: outputField,
		Path: path,
		ExpectedType: expectedType,
		ActualValue: actualValue,
		Example: example,
	}
}

// PropertyParseErrorInvalidDateStructure represents the invalid_date_structure error variant
type PropertyParseErrorInvalidDateStructure struct {
	*PropertyParseError
	OutputField string
	Path string
	Value string
	ExpectedFormat string
	Example string
}

func newPropertyParseErrorInvalidDateStructure(outputField string, path string, value string, expectedFormat string, example string) *PropertyParseErrorInvalidDateStructure {
	return &PropertyParseErrorInvalidDateStructure{
		PropertyParseError: &PropertyParseError{
			Type:    "invalid_date_structure",
			message: "invalid_date_structure",
		},
		OutputField: outputField,
		Path: path,
		Value: value,
		ExpectedFormat: expectedFormat,
		Example: example,
	}
}

// PropertyParseErrorInvalidDateTimeStructure represents the invalid_date_time_structure error variant
type PropertyParseErrorInvalidDateTimeStructure struct {
	*PropertyParseError
	OutputField string
	Path string
	Value string
	ExpectedFormat string
	Example string
}

func newPropertyParseErrorInvalidDateTimeStructure(outputField string, path string, value string, expectedFormat string, example string) *PropertyParseErrorInvalidDateTimeStructure {
	return &PropertyParseErrorInvalidDateTimeStructure{
		PropertyParseError: &PropertyParseError{
			Type:    "invalid_date_time_structure",
			message: "invalid_date_time_structure",
		},
		OutputField: outputField,
		Path: path,
		Value: value,
		ExpectedFormat: expectedFormat,
		Example: example,
	}
}

// PropertyParseErrorInvalidEnumValue represents the invalid_enum_value error variant
type PropertyParseErrorInvalidEnumValue struct {
	*PropertyParseError
	OutputField string
	Path string
	Value string
	AllowedValues []string
}

func newPropertyParseErrorInvalidEnumValue(outputField string, path string, value string, allowedValues []string) *PropertyParseErrorInvalidEnumValue {
	return &PropertyParseErrorInvalidEnumValue{
		PropertyParseError: &PropertyParseError{
			Type:    "invalid_enum_value",
			message: "invalid_enum_value",
		},
		OutputField: outputField,
		Path: path,
		Value: value,
		AllowedValues: allowedValues,
	}
}



// Error Parsers

// parseCommitScimUserChangeError parses error response JSON into the appropriate error variant
func parseCommitScimUserChangeError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newCommitScimUserChangeErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newCommitScimUserChangeErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "ScimConnectionNotFound":
		return newCommitScimUserChangeErrorScimConnectionNotFound()
	case "StagedChangeNotFound":
		return newCommitScimUserChangeErrorStagedChangeNotFound()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newCommitScimUserChangeErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newCommitScimUserChangeErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newCommitScimUserChangeErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseCompleteOidcLoginError parses error response JSON into the appropriate error variant
func parseCompleteOidcLoginError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newCompleteOidcLoginErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newCompleteOidcLoginErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "LoginBlockedByEmailAllowlist":
		return newCompleteOidcLoginErrorLoginBlockedByEmailAllowlist()
	case "ScimUserNotFoundWhereExpected":
		return newCompleteOidcLoginErrorScimUserNotFoundWhereExpected()
	case "ScimUserNotActive":
		return newCompleteOidcLoginErrorScimUserNotActive()
	case "InvalidLoginRequest":
		var details struct {
			Details OidcLoginRequestError `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newCompleteOidcLoginErrorUnexpectedError(fmt.Sprintf("failed to parse InvalidLoginRequest error: %v", err))
		}
		return newCompleteOidcLoginErrorInvalidLoginRequest(details.Details)
	case "IdentityProviderError":
		var details struct {
			Details IdentityProviderError `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newCompleteOidcLoginErrorUnexpectedError(fmt.Sprintf("failed to parse IdentityProviderError error: %v", err))
		}
		return newCompleteOidcLoginErrorIdentityProviderError(details.Details)
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newCompleteOidcLoginErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newCompleteOidcLoginErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newCompleteOidcLoginErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}



// parseCreateDeviceChallengeError parses error response JSON
func parseCreateDeviceChallengeError(data []byte) error {
	var errData CreateDeviceChallengeError
	if err := json.Unmarshal(data, &errData); err != nil {
		return fmt.Errorf("failed to parse CreateDeviceChallengeError: %w", err)
	}
	return &errData
}


// parseCreateImpersonationSessionError parses error response JSON into the appropriate error variant
func parseCreateImpersonationSessionError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newCreateImpersonationSessionErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newCreateImpersonationSessionErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "ImpersonationDisabled":
		return newCreateImpersonationSessionErrorImpersonationDisabled()
	case "UnauthorizedEmployee":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newCreateImpersonationSessionErrorUnexpectedError(fmt.Sprintf("failed to parse UnauthorizedEmployee error: %v", err))
		}
		return newCreateImpersonationSessionErrorUnauthorizedEmployee(details.Details)
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newCreateImpersonationSessionErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newCreateImpersonationSessionErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newCreateImpersonationSessionErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseCreateOidcClientError parses error response JSON into the appropriate error variant
func parseCreateOidcClientError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newCreateOidcClientErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newCreateOidcClientErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "InvalidFields":
		var details struct {
			Details CreateOidcClientErrorInvalidFieldsDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newCreateOidcClientErrorUnexpectedError(fmt.Sprintf("failed to parse InvalidFields error: %v", err))
		}
		return newCreateOidcClientErrorInvalidFields(details.Details)
	case "ClientIdAlreadyTaken":
		return newCreateOidcClientErrorClientIDAlreadyTaken()
	case "CustomerIdAlreadyTakenForEoidcClient":
		return newCreateOidcClientErrorCustomerIDAlreadyTakenForEoidcClient()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newCreateOidcClientErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newCreateOidcClientErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newCreateOidcClientErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseCreateScimConnectionError parses error response JSON into the appropriate error variant
func parseCreateScimConnectionError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newCreateScimConnectionErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newCreateScimConnectionErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "InvalidFields":
		var details struct {
			Details CreateScimConnectionErrorInvalidFieldsDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newCreateScimConnectionErrorUnexpectedError(fmt.Sprintf("failed to parse InvalidFields error: %v", err))
		}
		return newCreateScimConnectionErrorInvalidFields(details.Details)
	case "ScimConnectionForCustomerIdAlreadyExists":
		return newCreateScimConnectionErrorScimConnectionForCustomerIDAlreadyExists()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newCreateScimConnectionErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newCreateScimConnectionErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newCreateScimConnectionErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseCreateSessionError parses error response JSON into the appropriate error variant
func parseCreateSessionError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newCreateSessionErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newCreateSessionErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "SessionLimitExceeded":
		var details struct {
			Details CreateSessionErrorSessionLimitExceededDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newCreateSessionErrorUnexpectedError(fmt.Sprintf("failed to parse SessionLimitExceeded error: %v", err))
		}
		return newCreateSessionErrorSessionLimitExceeded(details.Details)
	case "IpAddressError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newCreateSessionErrorUnexpectedError(fmt.Sprintf("failed to parse IpAddressError error: %v", err))
		}
		return newCreateSessionErrorIPAddressError(details.Details)
	case "TagParseError":
		var details struct {
			Details InvalidTagError `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newCreateSessionErrorUnexpectedError(fmt.Sprintf("failed to parse TagParseError error: %v", err))
		}
		return newCreateSessionErrorTagParseError(details.Details)
	case "InvalidDeviceRegistration":
		return newCreateSessionErrorInvalidDeviceRegistration()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newCreateSessionErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newCreateSessionErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newCreateSessionErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseCreateStatelessTokenError parses error response JSON into the appropriate error variant
func parseCreateStatelessTokenError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newCreateStatelessTokenErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newCreateStatelessTokenErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "TokenCreationFailed":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newCreateStatelessTokenErrorUnexpectedError(fmt.Sprintf("failed to parse TokenCreationFailed error: %v", err))
		}
		return newCreateStatelessTokenErrorTokenCreationFailed(details.Details)
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newCreateStatelessTokenErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newCreateStatelessTokenErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newCreateStatelessTokenErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseDeleteOidcClientError parses error response JSON into the appropriate error variant
func parseDeleteOidcClientError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newDeleteOidcClientErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newDeleteOidcClientErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "OidcClientNotFound":
		return newDeleteOidcClientErrorOIDCClientNotFound()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newDeleteOidcClientErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newDeleteOidcClientErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newDeleteOidcClientErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseDeleteScimConnectionError parses error response JSON into the appropriate error variant
func parseDeleteScimConnectionError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newDeleteScimConnectionErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newDeleteScimConnectionErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "ScimConnectionNotFound":
		return newDeleteScimConnectionErrorScimConnectionNotFound()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newDeleteScimConnectionErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newDeleteScimConnectionErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newDeleteScimConnectionErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}



// parseDeregisterAllPasskeysForUserError parses error response JSON
func parseDeregisterAllPasskeysForUserError(data []byte) error {
	var errData DeregisterAllPasskeysForUserError
	if err := json.Unmarshal(data, &errData); err != nil {
		return fmt.Errorf("failed to parse DeregisterAllPasskeysForUserError: %w", err)
	}
	return &errData
}


// parseDeregisterPasskeyError parses error response JSON into the appropriate error variant
func parseDeregisterPasskeyError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newDeregisterPasskeyErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newDeregisterPasskeyErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "PasskeyNotFound":
		return newDeregisterPasskeyErrorPasskeyNotFound()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newDeregisterPasskeyErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newDeregisterPasskeyErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newDeregisterPasskeyErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseFetchAllActiveImpersonationSessionsError parses error response JSON into the appropriate error variant
func parseFetchAllActiveImpersonationSessionsError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newFetchAllActiveImpersonationSessionsErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newFetchAllActiveImpersonationSessionsErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "InvalidPagingToken":
		return newFetchAllActiveImpersonationSessionsErrorInvalidPagingToken()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newFetchAllActiveImpersonationSessionsErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newFetchAllActiveImpersonationSessionsErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newFetchAllActiveImpersonationSessionsErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}



// parseFetchAllImpersonationSessionsForEmployeeError parses error response JSON
func parseFetchAllImpersonationSessionsForEmployeeError(data []byte) error {
	var errData FetchAllImpersonationSessionsForEmployeeError
	if err := json.Unmarshal(data, &errData); err != nil {
		return fmt.Errorf("failed to parse FetchAllImpersonationSessionsForEmployeeError: %w", err)
	}
	return &errData
}



// parseFetchAllImpersonationSessionsForUserError parses error response JSON
func parseFetchAllImpersonationSessionsForUserError(data []byte) error {
	var errData FetchAllImpersonationSessionsForUserError
	if err := json.Unmarshal(data, &errData); err != nil {
		return fmt.Errorf("failed to parse FetchAllImpersonationSessionsForUserError: %w", err)
	}
	return &errData
}



// parseFetchAllPasskeysForUserError parses error response JSON
func parseFetchAllPasskeysForUserError(data []byte) error {
	var errData FetchAllPasskeysForUserError
	if err := json.Unmarshal(data, &errData); err != nil {
		return fmt.Errorf("failed to parse FetchAllPasskeysForUserError: %w", err)
	}
	return &errData
}



// parseFetchAllSessionsError parses error response JSON
func parseFetchAllSessionsError(data []byte) error {
	var errData FetchAllSessionsError
	if err := json.Unmarshal(data, &errData); err != nil {
		return fmt.Errorf("failed to parse FetchAllSessionsError: %w", err)
	}
	return &errData
}



// parseFetchAllSessionsForUserError parses error response JSON
func parseFetchAllSessionsForUserError(data []byte) error {
	var errData FetchAllSessionsForUserError
	if err := json.Unmarshal(data, &errData); err != nil {
		return fmt.Errorf("failed to parse FetchAllSessionsForUserError: %w", err)
	}
	return &errData
}


// parseFetchImpersonationSessionByIdError parses error response JSON into the appropriate error variant
func parseFetchImpersonationSessionByIdError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newFetchImpersonationSessionByIdErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newFetchImpersonationSessionByIdErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "SessionNotFound":
		return newFetchImpersonationSessionByIdErrorSessionNotFound()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newFetchImpersonationSessionByIdErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newFetchImpersonationSessionByIdErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newFetchImpersonationSessionByIdErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseFetchOidcClientError parses error response JSON into the appropriate error variant
func parseFetchOidcClientError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newFetchOidcClientErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newFetchOidcClientErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "OidcClientNotFound":
		return newFetchOidcClientErrorOIDCClientNotFound()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newFetchOidcClientErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newFetchOidcClientErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newFetchOidcClientErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseFetchScimConnectionError parses error response JSON into the appropriate error variant
func parseFetchScimConnectionError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newFetchScimConnectionErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newFetchScimConnectionErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "ScimConnectionNotFound":
		return newFetchScimConnectionErrorScimConnectionNotFound()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newFetchScimConnectionErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newFetchScimConnectionErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newFetchScimConnectionErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseFetchSessionByIdError parses error response JSON into the appropriate error variant
func parseFetchSessionByIdError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newFetchSessionByIdErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newFetchSessionByIdErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "SessionNotFound":
		return newFetchSessionByIdErrorSessionNotFound()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newFetchSessionByIdErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newFetchSessionByIdErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newFetchSessionByIdErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseFinishPasskeyAuthenticationError parses error response JSON into the appropriate error variant
func parseFinishPasskeyAuthenticationError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newFinishPasskeyAuthenticationErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newFinishPasskeyAuthenticationErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "CannotParseAdditionalAllowedOrigin":
		return newFinishPasskeyAuthenticationErrorCannotParseAdditionalAllowedOrigin()
	case "NoAuthenticationChallengeFound":
		return newFinishPasskeyAuthenticationErrorNoAuthenticationChallengeFound()
	case "OriginNotAllowed":
		var details struct {
			Details FinishPasskeyAuthenticationErrorOriginNotAllowedDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newFinishPasskeyAuthenticationErrorUnexpectedError(fmt.Sprintf("failed to parse OriginNotAllowed error: %v", err))
		}
		return newFinishPasskeyAuthenticationErrorOriginNotAllowed(details.Details)
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newFinishPasskeyAuthenticationErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newFinishPasskeyAuthenticationErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newFinishPasskeyAuthenticationErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseFinishPasskeyRegistrationError parses error response JSON into the appropriate error variant
func parseFinishPasskeyRegistrationError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newFinishPasskeyRegistrationErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newFinishPasskeyRegistrationErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "CannotParseAdditionalAllowedOrigin":
		return newFinishPasskeyRegistrationErrorCannotParseAdditionalAllowedOrigin()
	case "NoRegistrationChallengeFound":
		return newFinishPasskeyRegistrationErrorNoRegistrationChallengeFound()
	case "OriginNotAllowed":
		var details struct {
			Details FinishPasskeyRegistrationErrorOriginNotAllowedDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newFinishPasskeyRegistrationErrorUnexpectedError(fmt.Sprintf("failed to parse OriginNotAllowed error: %v", err))
		}
		return newFinishPasskeyRegistrationErrorOriginNotAllowed(details.Details)
	case "PasskeyForUserAlreadyExists":
		return newFinishPasskeyRegistrationErrorPasskeyForUserAlreadyExists()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newFinishPasskeyRegistrationErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newFinishPasskeyRegistrationErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newFinishPasskeyRegistrationErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}



// parseGetJwksError parses error response JSON
func parseGetJwksError(data []byte) error {
	var errData GetJwksError
	if err := json.Unmarshal(data, &errData); err != nil {
		return fmt.Errorf("failed to parse GetJwksError: %w", err)
	}
	return &errData
}


// parseGetScimUserError parses error response JSON into the appropriate error variant
func parseGetScimUserError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newGetScimUserErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newGetScimUserErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "ScimConnectionNotFound":
		return newGetScimUserErrorScimConnectionNotFound()
	case "UserNotFound":
		return newGetScimUserErrorUserNotFound()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newGetScimUserErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newGetScimUserErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newGetScimUserErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseGetScimUsersError parses error response JSON into the appropriate error variant
func parseGetScimUsersError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newGetScimUsersErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newGetScimUsersErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "ScimConnectionNotFound":
		return newGetScimUsersErrorScimConnectionNotFound()
	case "InvalidQueryField":
		var details struct {
			Details GetScimUsersErrorInvalidQueryFieldDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newGetScimUsersErrorUnexpectedError(fmt.Sprintf("failed to parse InvalidQueryField error: %v", err))
		}
		return newGetScimUsersErrorInvalidQueryField(details.Details)
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newGetScimUsersErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newGetScimUsersErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newGetScimUsersErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseInitiateOidcLoginError parses error response JSON into the appropriate error variant
func parseInitiateOidcLoginError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newInitiateOidcLoginErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newInitiateOidcLoginErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "ClientNotFound":
		return newInitiateOidcLoginErrorClientNotFound()
	case "RedirectUrlInvalid":
		var details struct {
			Details InitiateOidcLoginErrorRedirectUrlInvalidDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newInitiateOidcLoginErrorUnexpectedError(fmt.Sprintf("failed to parse RedirectUrlInvalid error: %v", err))
		}
		return newInitiateOidcLoginErrorRedirectURLInvalid(details.Details)
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newInitiateOidcLoginErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newInitiateOidcLoginErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newInitiateOidcLoginErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseIntegrationKeyError parses error response JSON into the appropriate error variant
func parseIntegrationKeyError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newIntegrationKeyErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newIntegrationKeyErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "InvalidPrefix":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newIntegrationKeyErrorUnexpectedError(fmt.Sprintf("failed to parse InvalidPrefix error: %v", err))
		}
		return newIntegrationKeyErrorInvalidPrefix(details.Details)
	case "IntegrationKeyNotFound":
		return newIntegrationKeyErrorIntegrationKeyNotFound()
	case "NoIntegrationKeyFoundInHeader":
		return newIntegrationKeyErrorNoIntegrationKeyFoundInHeader()
	case "CommandNotAllowed":
		var details struct {
			Details IntegrationKeyErrorCommandNotAllowedDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newIntegrationKeyErrorUnexpectedError(fmt.Sprintf("failed to parse CommandNotAllowed error: %v", err))
		}
		return newIntegrationKeyErrorCommandNotAllowed(details.Details)
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newIntegrationKeyErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newIntegrationKeyErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newIntegrationKeyErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}



// parseInvalidTagError parses error response JSON
func parseInvalidTagError(data []byte) error {
	var errData InvalidTagError
	if err := json.Unmarshal(data, &errData); err != nil {
		return fmt.Errorf("failed to parse InvalidTagError: %w", err)
	}
	return &errData
}



// parseInvalidateAllImpersonationSessionsForEmployeeError parses error response JSON
func parseInvalidateAllImpersonationSessionsForEmployeeError(data []byte) error {
	var errData InvalidateAllImpersonationSessionsForEmployeeError
	if err := json.Unmarshal(data, &errData); err != nil {
		return fmt.Errorf("failed to parse InvalidateAllImpersonationSessionsForEmployeeError: %w", err)
	}
	return &errData
}



// parseInvalidateAllImpersonationSessionsForUserError parses error response JSON
func parseInvalidateAllImpersonationSessionsForUserError(data []byte) error {
	var errData InvalidateAllImpersonationSessionsForUserError
	if err := json.Unmarshal(data, &errData); err != nil {
		return fmt.Errorf("failed to parse InvalidateAllImpersonationSessionsForUserError: %w", err)
	}
	return &errData
}



// parseInvalidateAllSessionsForUserError parses error response JSON
func parseInvalidateAllSessionsForUserError(data []byte) error {
	var errData InvalidateAllSessionsForUserError
	if err := json.Unmarshal(data, &errData); err != nil {
		return fmt.Errorf("failed to parse InvalidateAllSessionsForUserError: %w", err)
	}
	return &errData
}



// parseInvalidateAllSessionsForUserExceptOneError parses error response JSON
func parseInvalidateAllSessionsForUserExceptOneError(data []byte) error {
	var errData InvalidateAllSessionsForUserExceptOneError
	if err := json.Unmarshal(data, &errData); err != nil {
		return fmt.Errorf("failed to parse InvalidateAllSessionsForUserExceptOneError: %w", err)
	}
	return &errData
}


// parseInvalidateImpersonationSessionByIdError parses error response JSON into the appropriate error variant
func parseInvalidateImpersonationSessionByIdError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newInvalidateImpersonationSessionByIdErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newInvalidateImpersonationSessionByIdErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "SessionNotFound":
		return newInvalidateImpersonationSessionByIdErrorSessionNotFound()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newInvalidateImpersonationSessionByIdErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newInvalidateImpersonationSessionByIdErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newInvalidateImpersonationSessionByIdErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseInvalidateImpersonationSessionByTokenError parses error response JSON into the appropriate error variant
func parseInvalidateImpersonationSessionByTokenError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newInvalidateImpersonationSessionByTokenErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newInvalidateImpersonationSessionByTokenErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "SessionNotFound":
		return newInvalidateImpersonationSessionByTokenErrorSessionNotFound()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newInvalidateImpersonationSessionByTokenErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newInvalidateImpersonationSessionByTokenErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newInvalidateImpersonationSessionByTokenErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseInvalidateSessionByIdError parses error response JSON into the appropriate error variant
func parseInvalidateSessionByIdError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newInvalidateSessionByIdErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newInvalidateSessionByIdErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "SessionNotFound":
		return newInvalidateSessionByIdErrorSessionNotFound()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newInvalidateSessionByIdErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newInvalidateSessionByIdErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newInvalidateSessionByIdErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}



// parseInvalidateSessionByTokenError parses error response JSON
func parseInvalidateSessionByTokenError(data []byte) error {
	var errData InvalidateSessionByTokenError
	if err := json.Unmarshal(data, &errData); err != nil {
		return fmt.Errorf("failed to parse InvalidateSessionByTokenError: %w", err)
	}
	return &errData
}


// parseLinkScimUserError parses error response JSON into the appropriate error variant
func parseLinkScimUserError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newLinkScimUserErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newLinkScimUserErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "ScimConnectionNotFound":
		return newLinkScimUserErrorScimConnectionNotFound()
	case "UserNotFound":
		return newLinkScimUserErrorUserNotFound()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newLinkScimUserErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newLinkScimUserErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newLinkScimUserErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parsePatchOidcClientError parses error response JSON into the appropriate error variant
func parsePatchOidcClientError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newPatchOidcClientErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newPatchOidcClientErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "OidcClientNotFound":
		return newPatchOidcClientErrorOIDCClientNotFound()
	case "InvalidFields":
		var details struct {
			Details PatchOidcClientErrorInvalidFieldsDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newPatchOidcClientErrorUnexpectedError(fmt.Sprintf("failed to parse InvalidFields error: %v", err))
		}
		return newPatchOidcClientErrorInvalidFields(details.Details)
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newPatchOidcClientErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newPatchOidcClientErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newPatchOidcClientErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parsePatchScimConnectionError parses error response JSON into the appropriate error variant
func parsePatchScimConnectionError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newPatchScimConnectionErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newPatchScimConnectionErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "ScimConnectionNotFound":
		return newPatchScimConnectionErrorScimConnectionNotFound()
	case "DisplayNameInvalid":
		var details struct {
			Details PatchScimConnectionErrorDisplayNameInvalidDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newPatchScimConnectionErrorUnexpectedError(fmt.Sprintf("failed to parse DisplayNameInvalid error: %v", err))
		}
		return newPatchScimConnectionErrorDisplayNameInvalid(details.Details)
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newPatchScimConnectionErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newPatchScimConnectionErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newPatchScimConnectionErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}



// parsePingError parses error response JSON
func parsePingError(data []byte) error {
	var errData PingError
	if err := json.Unmarshal(data, &errData); err != nil {
		return fmt.Errorf("failed to parse PingError: %w", err)
	}
	return &errData
}


// parseRegisterDeviceError parses error response JSON into the appropriate error variant
func parseRegisterDeviceError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newRegisterDeviceErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newRegisterDeviceErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "SessionNotFound":
		return newRegisterDeviceErrorSessionNotFound()
	case "NewDeviceChallengeRequired":
		var details struct {
			Details RegisterDeviceErrorNewDeviceChallengeRequiredDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newRegisterDeviceErrorUnexpectedError(fmt.Sprintf("failed to parse NewDeviceChallengeRequired error: %v", err))
		}
		return newRegisterDeviceErrorNewDeviceChallengeRequired(details.Details)
	case "InvalidDeviceRegistration":
		return newRegisterDeviceErrorInvalidDeviceRegistration()
	case "DeviceAlreadyRegistered":
		return newRegisterDeviceErrorDeviceAlreadyRegistered()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newRegisterDeviceErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newRegisterDeviceErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newRegisterDeviceErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseResetScimApiKeyError parses error response JSON into the appropriate error variant
func parseResetScimApiKeyError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newResetScimApiKeyErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newResetScimApiKeyErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "ScimConnectionNotFound":
		return newResetScimApiKeyErrorScimConnectionNotFound()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newResetScimApiKeyErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newResetScimApiKeyErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newResetScimApiKeyErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseRotateStatelessTokenKeyError parses error response JSON into the appropriate error variant
func parseRotateStatelessTokenKeyError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newRotateStatelessTokenKeyErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newRotateStatelessTokenKeyErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "RotationFailed":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newRotateStatelessTokenKeyErrorUnexpectedError(fmt.Sprintf("failed to parse RotationFailed error: %v", err))
		}
		return newRotateStatelessTokenKeyErrorRotationFailed(details.Details)
	case "InvalidParameters":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newRotateStatelessTokenKeyErrorUnexpectedError(fmt.Sprintf("failed to parse InvalidParameters error: %v", err))
		}
		return newRotateStatelessTokenKeyErrorInvalidParameters(details.Details)
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newRotateStatelessTokenKeyErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newRotateStatelessTokenKeyErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newRotateStatelessTokenKeyErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}



// parseScimClientFacingError parses error response JSON
func parseScimClientFacingError(data []byte) error {
	var errData ScimClientFacingError
	if err := json.Unmarshal(data, &errData); err != nil {
		return fmt.Errorf("failed to parse ScimClientFacingError: %w", err)
	}
	return &errData
}


// parseScimUnderlyingError parses error response JSON into the appropriate error variant
func parseScimUnderlyingError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newScimUnderlyingErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newScimUnderlyingErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "InvalidApiKey":
		return newScimUnderlyingErrorInvalidAPIKey()
	case "InvalidPath":
		return newScimUnderlyingErrorInvalidPath()
	case "UserNotFound":
		return newScimUnderlyingErrorUserNotFound()
	case "GroupNotFound":
		return newScimUnderlyingErrorGroupNotFound()
	case "MissingRequiredField":
		var details struct {
			Details ScimUnderlyingErrorMissingRequiredFieldDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newScimUnderlyingErrorUnexpectedError(fmt.Sprintf("failed to parse MissingRequiredField error: %v", err))
		}
		return newScimUnderlyingErrorMissingRequiredField(details.Details)
	case "CantRemoveRequiredField":
		var details struct {
			Details ScimUnderlyingErrorCantRemoveRequiredFieldDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newScimUnderlyingErrorUnexpectedError(fmt.Sprintf("failed to parse CantRemoveRequiredField error: %v", err))
		}
		return newScimUnderlyingErrorCantRemoveRequiredField(details.Details)
	case "ScimUserIdAlreadyTaken":
		var details struct {
			Details ScimUnderlyingErrorScimUserIdAlreadyTakenDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newScimUnderlyingErrorUnexpectedError(fmt.Sprintf("failed to parse ScimUserIdAlreadyTaken error: %v", err))
		}
		return newScimUnderlyingErrorScimUserIDAlreadyTaken(details.Details)
	case "InvalidBody":
		return newScimUnderlyingErrorInvalidBody()
	case "InvalidQueryField":
		var details struct {
			Details ScimUnderlyingErrorInvalidQueryFieldDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newScimUnderlyingErrorUnexpectedError(fmt.Sprintf("failed to parse InvalidQueryField error: %v", err))
		}
		return newScimUnderlyingErrorInvalidQueryField(details.Details)
	case "InvalidPatchPath":
		return newScimUnderlyingErrorInvalidPatchPath()
	case "InvalidPatchOperation":
		return newScimUnderlyingErrorInvalidPatchOperation()
	case "InvalidPatchValue":
		var details struct {
			Details ScimUnderlyingErrorInvalidPatchValueDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newScimUnderlyingErrorUnexpectedError(fmt.Sprintf("failed to parse InvalidPatchValue error: %v", err))
		}
		return newScimUnderlyingErrorInvalidPatchValue(details.Details)
	case "InvalidSchema":
		return newScimUnderlyingErrorInvalidSchema()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newScimUnderlyingErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newScimUnderlyingErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newScimUnderlyingErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseStartPasskeyAuthenticationError parses error response JSON into the appropriate error variant
func parseStartPasskeyAuthenticationError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newStartPasskeyAuthenticationErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newStartPasskeyAuthenticationErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "CannotParseAdditionalAllowedOrigin":
		return newStartPasskeyAuthenticationErrorCannotParseAdditionalAllowedOrigin()
	case "NoPasskeysRegisteredForUser":
		return newStartPasskeyAuthenticationErrorNoPasskeysRegisteredForUser()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newStartPasskeyAuthenticationErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newStartPasskeyAuthenticationErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newStartPasskeyAuthenticationErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseStartPasskeyRegistrationError parses error response JSON into the appropriate error variant
func parseStartPasskeyRegistrationError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newStartPasskeyRegistrationErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newStartPasskeyRegistrationErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "CannotParseAdditionalAllowedOrigin":
		return newStartPasskeyRegistrationErrorCannotParseAdditionalAllowedOrigin()
	case "TooManyPasskeys":
		var details struct {
			Details StartPasskeyRegistrationErrorTooManyPasskeysDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newStartPasskeyRegistrationErrorUnexpectedError(fmt.Sprintf("failed to parse TooManyPasskeys error: %v", err))
		}
		return newStartPasskeyRegistrationErrorTooManyPasskeys(details.Details)
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newStartPasskeyRegistrationErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newStartPasskeyRegistrationErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newStartPasskeyRegistrationErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseUpdateSessionError parses error response JSON into the appropriate error variant
func parseUpdateSessionError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newUpdateSessionErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newUpdateSessionErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "SessionNotFound":
		return newUpdateSessionErrorSessionNotFound()
	case "ConflictingMetadataOptions":
		return newUpdateSessionErrorConflictingMetadataOptions()
	case "InvalidTagFormat":
		var details struct {
			Details InvalidTagError `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newUpdateSessionErrorUnexpectedError(fmt.Sprintf("failed to parse InvalidTagFormat error: %v", err))
		}
		return newUpdateSessionErrorInvalidTagFormat(details.Details)
	case "CannotModifyOnCreateOnlyTags":
		var details struct {
			Details UpdateSessionErrorCannotModifyOnCreateOnlyTagsDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newUpdateSessionErrorUnexpectedError(fmt.Sprintf("failed to parse CannotModifyOnCreateOnlyTags error: %v", err))
		}
		return newUpdateSessionErrorCannotModifyOnCreateOnlyTags(details.Details)
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newUpdateSessionErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newUpdateSessionErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newUpdateSessionErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseUpdateSessionsError parses error response JSON into the appropriate error variant
func parseUpdateSessionsError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newUpdateSessionsErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newUpdateSessionsErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "ConflictingMetadataOptions":
		return newUpdateSessionsErrorConflictingMetadataOptions()
	case "InvalidTagFormat":
		var details struct {
			Details InvalidTagError `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newUpdateSessionsErrorUnexpectedError(fmt.Sprintf("failed to parse InvalidTagFormat error: %v", err))
		}
		return newUpdateSessionsErrorInvalidTagFormat(details.Details)
	case "CannotModifyOnCreateOnlyTags":
		var details struct {
			Details UpdateSessionsErrorCannotModifyOnCreateOnlyTagsDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newUpdateSessionsErrorUnexpectedError(fmt.Sprintf("failed to parse CannotModifyOnCreateOnlyTags error: %v", err))
		}
		return newUpdateSessionsErrorCannotModifyOnCreateOnlyTags(details.Details)
	case "UpdatingTooManySessionsAtOnce":
		return newUpdateSessionsErrorUpdatingTooManySessionsAtOnce()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newUpdateSessionsErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newUpdateSessionsErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newUpdateSessionsErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseValidateAndRefreshSessionError parses error response JSON into the appropriate error variant
func parseValidateAndRefreshSessionError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newValidateAndRefreshSessionErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newValidateAndRefreshSessionErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "InvalidSessionToken":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newValidateAndRefreshSessionErrorUnexpectedError(fmt.Sprintf("failed to parse InvalidSessionToken error: %v", err))
		}
		return newValidateAndRefreshSessionErrorInvalidSessionToken(details.Details)
	case "IpAddressError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newValidateAndRefreshSessionErrorUnexpectedError(fmt.Sprintf("failed to parse IpAddressError error: %v", err))
		}
		return newValidateAndRefreshSessionErrorIPAddressError(details.Details)
	case "NewDeviceChallengeRequired":
		var details struct {
			Details ValidateAndRefreshSessionErrorNewDeviceChallengeRequiredDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newValidateAndRefreshSessionErrorUnexpectedError(fmt.Sprintf("failed to parse NewDeviceChallengeRequired error: %v", err))
		}
		return newValidateAndRefreshSessionErrorNewDeviceChallengeRequired(details.Details)
	case "DeviceVerificationRequired":
		return newValidateAndRefreshSessionErrorDeviceVerificationRequired()
	case "DeviceVerificationFailed":
		return newValidateAndRefreshSessionErrorDeviceVerificationFailed()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newValidateAndRefreshSessionErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newValidateAndRefreshSessionErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newValidateAndRefreshSessionErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseValidateImpersonationSessionError parses error response JSON into the appropriate error variant
func parseValidateImpersonationSessionError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newValidateImpersonationSessionErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newValidateImpersonationSessionErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "ImpersonationNotEnabled":
		return newValidateImpersonationSessionErrorImpersonationNotEnabled()
	case "InvalidImpersonationToken":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newValidateImpersonationSessionErrorUnexpectedError(fmt.Sprintf("failed to parse InvalidImpersonationToken error: %v", err))
		}
		return newValidateImpersonationSessionErrorInvalidImpersonationToken(details.Details)
	case "SessionNotFound":
		return newValidateImpersonationSessionErrorSessionNotFound()
	case "IpAddressMismatch":
		return newValidateImpersonationSessionErrorIPAddressMismatch()
	case "UserAgentMismatch":
		return newValidateImpersonationSessionErrorUserAgentMismatch()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newValidateImpersonationSessionErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newValidateImpersonationSessionErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newValidateImpersonationSessionErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parseValidateSessionError parses error response JSON into the appropriate error variant
func parseValidateSessionError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return newValidateSessionErrorUnexpectedError(fmt.Sprintf("malformed error response: %v", err))
	}

	var errorType string
	if err := json.Unmarshal(raw["type"], &errorType); err != nil {
		return newValidateSessionErrorUnexpectedError("missing 'type' field in error response")
	}

	switch errorType {
	case "InvalidSessionToken":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newValidateSessionErrorUnexpectedError(fmt.Sprintf("failed to parse InvalidSessionToken error: %v", err))
		}
		return newValidateSessionErrorInvalidSessionToken(details.Details)
	case "IpAddressError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newValidateSessionErrorUnexpectedError(fmt.Sprintf("failed to parse IpAddressError error: %v", err))
		}
		return newValidateSessionErrorIPAddressError(details.Details)
	case "NewDeviceChallengeRequired":
		var details struct {
			Details ValidateSessionErrorNewDeviceChallengeRequiredDetails `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newValidateSessionErrorUnexpectedError(fmt.Sprintf("failed to parse NewDeviceChallengeRequired error: %v", err))
		}
		return newValidateSessionErrorNewDeviceChallengeRequired(details.Details)
	case "DeviceVerificationRequired":
		return newValidateSessionErrorDeviceVerificationRequired()
	case "DeviceVerificationFailed":
		return newValidateSessionErrorDeviceVerificationFailed()
	case "UnexpectedError":
		var details struct {
			Details string `json:"details"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return newValidateSessionErrorUnexpectedError(fmt.Sprintf("failed to parse UnexpectedError error: %v", err))
		}
		return newValidateSessionErrorUnexpectedError(details.Details)
	default:
		// Unknown error type
		return newValidateSessionErrorUnexpectedError(fmt.Sprintf("unknown error type '%s': %s", errorType, string(data)))
	}
}


// parsePropertyParseError parses error response JSON into the appropriate error variant
func parsePropertyParseError(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("malformed error response: %w", err)
	}

	var errorType string
	if err := json.Unmarshal(raw["error_type"], &errorType); err != nil {
		return fmt.Errorf("missing 'error_type' field in error response")
	}

	switch errorType {
	case "property_missing":
		var details struct {
			OutputField string `json:"output_field"`
			Path string `json:"path"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return fmt.Errorf("failed to parse property_missing error: %w", err)
		}
		return newPropertyParseErrorPropertyMissing(details.OutputField, details.Path)
	case "invalid_type":
		var details struct {
			OutputField string `json:"output_field"`
			Path string `json:"path"`
			ExpectedType string `json:"expected_type"`
			ActualValue json.RawMessage `json:"actual_value"`
			Example *string `json:"example,omitempty"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return fmt.Errorf("failed to parse invalid_type error: %w", err)
		}
		return newPropertyParseErrorInvalidType(details.OutputField, details.Path, details.ExpectedType, details.ActualValue, details.Example)
	case "invalid_date_structure":
		var details struct {
			OutputField string `json:"output_field"`
			Path string `json:"path"`
			Value string `json:"value"`
			ExpectedFormat string `json:"expected_format"`
			Example string `json:"example"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return fmt.Errorf("failed to parse invalid_date_structure error: %w", err)
		}
		return newPropertyParseErrorInvalidDateStructure(details.OutputField, details.Path, details.Value, details.ExpectedFormat, details.Example)
	case "invalid_date_time_structure":
		var details struct {
			OutputField string `json:"output_field"`
			Path string `json:"path"`
			Value string `json:"value"`
			ExpectedFormat string `json:"expected_format"`
			Example string `json:"example"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return fmt.Errorf("failed to parse invalid_date_time_structure error: %w", err)
		}
		return newPropertyParseErrorInvalidDateTimeStructure(details.OutputField, details.Path, details.Value, details.ExpectedFormat, details.Example)
	case "invalid_enum_value":
		var details struct {
			OutputField string `json:"output_field"`
			Path string `json:"path"`
			Value string `json:"value"`
			AllowedValues []string `json:"allowed_values"`
		}
		if err := json.Unmarshal(data, &details); err != nil {
			return fmt.Errorf("failed to parse invalid_enum_value error: %w", err)
		}
		return newPropertyParseErrorInvalidEnumValue(details.OutputField, details.Path, details.Value, details.AllowedValues)
	default:
		// Unknown error type
		return fmt.Errorf("unknown error type '%s': %s", errorType, string(data))
	}
}
