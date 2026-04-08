package byo

import (
	"encoding/json"
)

// Command is a discriminated union response type
type Command interface {
	isCommand()
}

// CommandPing represents the Ping variant
type CommandPing struct {
	Data PingCommand `json:"data"`
}

func (r *CommandPing) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandPing) MarshalJSON() ([]byte, error) {
	type Alias CommandPing
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "Ping",
		Alias: (*Alias)(v),
	})
}

// CommandStartPasskeyRegistration represents the StartPasskeyRegistration variant
type CommandStartPasskeyRegistration struct {
	Data StartPasskeyRegistrationCommand `json:"data"`
}

func (r *CommandStartPasskeyRegistration) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandStartPasskeyRegistration) MarshalJSON() ([]byte, error) {
	type Alias CommandStartPasskeyRegistration
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "StartPasskeyRegistration",
		Alias: (*Alias)(v),
	})
}

// CommandFinishPasskeyRegistration represents the FinishPasskeyRegistration variant
type CommandFinishPasskeyRegistration struct {
	Data FinishPasskeyRegistrationCommand `json:"data"`
}

func (r *CommandFinishPasskeyRegistration) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandFinishPasskeyRegistration) MarshalJSON() ([]byte, error) {
	type Alias CommandFinishPasskeyRegistration
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "FinishPasskeyRegistration",
		Alias: (*Alias)(v),
	})
}

// CommandStartPasskeyAuthentication represents the StartPasskeyAuthentication variant
type CommandStartPasskeyAuthentication struct {
	Data StartPasskeyAuthenticationCommand `json:"data"`
}

func (r *CommandStartPasskeyAuthentication) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandStartPasskeyAuthentication) MarshalJSON() ([]byte, error) {
	type Alias CommandStartPasskeyAuthentication
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "StartPasskeyAuthentication",
		Alias: (*Alias)(v),
	})
}

// CommandFinishPasskeyAuthentication represents the FinishPasskeyAuthentication variant
type CommandFinishPasskeyAuthentication struct {
	Data FinishPasskeyAuthenticationCommand `json:"data"`
}

func (r *CommandFinishPasskeyAuthentication) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandFinishPasskeyAuthentication) MarshalJSON() ([]byte, error) {
	type Alias CommandFinishPasskeyAuthentication
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "FinishPasskeyAuthentication",
		Alias: (*Alias)(v),
	})
}

// CommandFetchAllPasskeysForUser represents the FetchAllPasskeysForUser variant
type CommandFetchAllPasskeysForUser struct {
	Data FetchAllPasskeysForUserCommand `json:"data"`
}

func (r *CommandFetchAllPasskeysForUser) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandFetchAllPasskeysForUser) MarshalJSON() ([]byte, error) {
	type Alias CommandFetchAllPasskeysForUser
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "FetchAllPasskeysForUser",
		Alias: (*Alias)(v),
	})
}

// CommandDeregisterPasskey represents the DeregisterPasskey variant
type CommandDeregisterPasskey struct {
	Data DeregisterPasskeyCommand `json:"data"`
}

func (r *CommandDeregisterPasskey) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandDeregisterPasskey) MarshalJSON() ([]byte, error) {
	type Alias CommandDeregisterPasskey
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "DeregisterPasskey",
		Alias: (*Alias)(v),
	})
}

// CommandDeregisterAllPasskeysForUser represents the DeregisterAllPasskeysForUser variant
type CommandDeregisterAllPasskeysForUser struct {
	Data DeregisterAllPasskeysForUserCommand `json:"data"`
}

func (r *CommandDeregisterAllPasskeysForUser) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandDeregisterAllPasskeysForUser) MarshalJSON() ([]byte, error) {
	type Alias CommandDeregisterAllPasskeysForUser
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "DeregisterAllPasskeysForUser",
		Alias: (*Alias)(v),
	})
}

// CommandCreateDeviceChallenge represents the CreateDeviceChallenge variant
type CommandCreateDeviceChallenge struct {
	Data CreateDeviceChallengeCommand `json:"data"`
}

func (r *CommandCreateDeviceChallenge) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandCreateDeviceChallenge) MarshalJSON() ([]byte, error) {
	type Alias CommandCreateDeviceChallenge
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "CreateDeviceChallenge",
		Alias: (*Alias)(v),
	})
}

// CommandCreateSession represents the CreateSession variant
type CommandCreateSession struct {
	Data CreateSessionCommand `json:"data"`
}

func (r *CommandCreateSession) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandCreateSession) MarshalJSON() ([]byte, error) {
	type Alias CommandCreateSession
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "CreateSession",
		Alias: (*Alias)(v),
	})
}

// CommandRegisterDevice represents the RegisterDevice variant
type CommandRegisterDevice struct {
	Data RegisterDeviceCommand `json:"data"`
}

func (r *CommandRegisterDevice) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandRegisterDevice) MarshalJSON() ([]byte, error) {
	type Alias CommandRegisterDevice
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "RegisterDevice",
		Alias: (*Alias)(v),
	})
}

// CommandCreateStatelessToken represents the CreateStatelessToken variant
type CommandCreateStatelessToken struct {
	Data CreateStatelessTokenCommand `json:"data"`
}

func (r *CommandCreateStatelessToken) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandCreateStatelessToken) MarshalJSON() ([]byte, error) {
	type Alias CommandCreateStatelessToken
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "CreateStatelessToken",
		Alias: (*Alias)(v),
	})
}

// CommandGetJwks represents the GetJwks variant
type CommandGetJwks struct {
	Data GetJwksCommand `json:"data"`
}

func (r *CommandGetJwks) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandGetJwks) MarshalJSON() ([]byte, error) {
	type Alias CommandGetJwks
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "GetJwks",
		Alias: (*Alias)(v),
	})
}

// CommandRotateStatelessTokenKey represents the RotateStatelessTokenKey variant
type CommandRotateStatelessTokenKey struct {
	Data RotateStatelessTokenKeyCommand `json:"data"`
}

func (r *CommandRotateStatelessTokenKey) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandRotateStatelessTokenKey) MarshalJSON() ([]byte, error) {
	type Alias CommandRotateStatelessTokenKey
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "RotateStatelessTokenKey",
		Alias: (*Alias)(v),
	})
}

// CommandInvalidateSessionByToken represents the InvalidateSessionByToken variant
type CommandInvalidateSessionByToken struct {
	Data InvalidateSessionByTokenCommand `json:"data"`
}

func (r *CommandInvalidateSessionByToken) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandInvalidateSessionByToken) MarshalJSON() ([]byte, error) {
	type Alias CommandInvalidateSessionByToken
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "InvalidateSessionByToken",
		Alias: (*Alias)(v),
	})
}

// CommandInvalidateSessionByID represents the InvalidateSessionById variant
type CommandInvalidateSessionByID struct {
	Data InvalidateSessionByIdCommand `json:"data"`
}

func (r *CommandInvalidateSessionByID) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandInvalidateSessionByID) MarshalJSON() ([]byte, error) {
	type Alias CommandInvalidateSessionByID
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "InvalidateSessionById",
		Alias: (*Alias)(v),
	})
}

// CommandInvalidateAllSessionsForUser represents the InvalidateAllSessionsForUser variant
type CommandInvalidateAllSessionsForUser struct {
	Data InvalidateAllSessionsForUserCommand `json:"data"`
}

func (r *CommandInvalidateAllSessionsForUser) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandInvalidateAllSessionsForUser) MarshalJSON() ([]byte, error) {
	type Alias CommandInvalidateAllSessionsForUser
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "InvalidateAllSessionsForUser",
		Alias: (*Alias)(v),
	})
}

// CommandInvalidateAllSessionsForUserExceptOne represents the InvalidateAllSessionsForUserExceptOne variant
type CommandInvalidateAllSessionsForUserExceptOne struct {
	Data InvalidateAllSessionsForUserExceptOneCommand `json:"data"`
}

func (r *CommandInvalidateAllSessionsForUserExceptOne) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandInvalidateAllSessionsForUserExceptOne) MarshalJSON() ([]byte, error) {
	type Alias CommandInvalidateAllSessionsForUserExceptOne
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "InvalidateAllSessionsForUserExceptOne",
		Alias: (*Alias)(v),
	})
}

// CommandFetchAllSessionsForUser represents the FetchAllSessionsForUser variant
type CommandFetchAllSessionsForUser struct {
	Data FetchAllSessionsForUserCommand `json:"data"`
}

func (r *CommandFetchAllSessionsForUser) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandFetchAllSessionsForUser) MarshalJSON() ([]byte, error) {
	type Alias CommandFetchAllSessionsForUser
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "FetchAllSessionsForUser",
		Alias: (*Alias)(v),
	})
}

// CommandFetchAllSessions represents the FetchAllSessions variant
type CommandFetchAllSessions struct {
	Data FetchAllSessionsCommand `json:"data"`
}

func (r *CommandFetchAllSessions) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandFetchAllSessions) MarshalJSON() ([]byte, error) {
	type Alias CommandFetchAllSessions
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "FetchAllSessions",
		Alias: (*Alias)(v),
	})
}

// CommandFetchSessionByID represents the FetchSessionById variant
type CommandFetchSessionByID struct {
	Data FetchSessionByIdCommand `json:"data"`
}

func (r *CommandFetchSessionByID) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandFetchSessionByID) MarshalJSON() ([]byte, error) {
	type Alias CommandFetchSessionByID
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "FetchSessionById",
		Alias: (*Alias)(v),
	})
}

// CommandUpdateSession represents the UpdateSession variant
type CommandUpdateSession struct {
	Data UpdateSessionCommand `json:"data"`
}

func (r *CommandUpdateSession) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandUpdateSession) MarshalJSON() ([]byte, error) {
	type Alias CommandUpdateSession
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "UpdateSession",
		Alias: (*Alias)(v),
	})
}

// CommandUpdateSessions represents the UpdateSessions variant
type CommandUpdateSessions struct {
	Data UpdateSessionsCommand `json:"data"`
}

func (r *CommandUpdateSessions) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandUpdateSessions) MarshalJSON() ([]byte, error) {
	type Alias CommandUpdateSessions
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "UpdateSessions",
		Alias: (*Alias)(v),
	})
}

// CommandValidateSession represents the ValidateSession variant
type CommandValidateSession struct {
	Data ValidateSessionCommand `json:"data"`
}

func (r *CommandValidateSession) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandValidateSession) MarshalJSON() ([]byte, error) {
	type Alias CommandValidateSession
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "ValidateSession",
		Alias: (*Alias)(v),
	})
}

// CommandValidateAndRefreshSession represents the ValidateAndRefreshSession variant
type CommandValidateAndRefreshSession struct {
	Data ValidateAndRefreshSessionCommand `json:"data"`
}

func (r *CommandValidateAndRefreshSession) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandValidateAndRefreshSession) MarshalJSON() ([]byte, error) {
	type Alias CommandValidateAndRefreshSession
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "ValidateAndRefreshSession",
		Alias: (*Alias)(v),
	})
}

// CommandCreateImpersonationSession represents the CreateImpersonationSession variant
type CommandCreateImpersonationSession struct {
	Data CreateImpersonationSessionCommand `json:"data"`
}

func (r *CommandCreateImpersonationSession) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandCreateImpersonationSession) MarshalJSON() ([]byte, error) {
	type Alias CommandCreateImpersonationSession
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "CreateImpersonationSession",
		Alias: (*Alias)(v),
	})
}

// CommandValidateImpersonationSession represents the ValidateImpersonationSession variant
type CommandValidateImpersonationSession struct {
	Data ValidateImpersonationSessionCommand `json:"data"`
}

func (r *CommandValidateImpersonationSession) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandValidateImpersonationSession) MarshalJSON() ([]byte, error) {
	type Alias CommandValidateImpersonationSession
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "ValidateImpersonationSession",
		Alias: (*Alias)(v),
	})
}

// CommandFetchImpersonationSessionByID represents the FetchImpersonationSessionById variant
type CommandFetchImpersonationSessionByID struct {
	Data FetchImpersonationSessionByIdCommand `json:"data"`
}

func (r *CommandFetchImpersonationSessionByID) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandFetchImpersonationSessionByID) MarshalJSON() ([]byte, error) {
	type Alias CommandFetchImpersonationSessionByID
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "FetchImpersonationSessionById",
		Alias: (*Alias)(v),
	})
}

// CommandFetchAllImpersonationSessionsForEmployee represents the FetchAllImpersonationSessionsForEmployee variant
type CommandFetchAllImpersonationSessionsForEmployee struct {
	Data FetchAllImpersonationSessionsForEmployeeCommand `json:"data"`
}

func (r *CommandFetchAllImpersonationSessionsForEmployee) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandFetchAllImpersonationSessionsForEmployee) MarshalJSON() ([]byte, error) {
	type Alias CommandFetchAllImpersonationSessionsForEmployee
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "FetchAllImpersonationSessionsForEmployee",
		Alias: (*Alias)(v),
	})
}

// CommandFetchAllImpersonationSessionsForUser represents the FetchAllImpersonationSessionsForUser variant
type CommandFetchAllImpersonationSessionsForUser struct {
	Data FetchAllImpersonationSessionsForUserCommand `json:"data"`
}

func (r *CommandFetchAllImpersonationSessionsForUser) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandFetchAllImpersonationSessionsForUser) MarshalJSON() ([]byte, error) {
	type Alias CommandFetchAllImpersonationSessionsForUser
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "FetchAllImpersonationSessionsForUser",
		Alias: (*Alias)(v),
	})
}

// CommandFetchAllActiveImpersonationSessions represents the FetchAllActiveImpersonationSessions variant
type CommandFetchAllActiveImpersonationSessions struct {
	Data FetchAllActiveImpersonationSessionsCommand `json:"data"`
}

func (r *CommandFetchAllActiveImpersonationSessions) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandFetchAllActiveImpersonationSessions) MarshalJSON() ([]byte, error) {
	type Alias CommandFetchAllActiveImpersonationSessions
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "FetchAllActiveImpersonationSessions",
		Alias: (*Alias)(v),
	})
}

// CommandInvalidateImpersonationSessionByID represents the InvalidateImpersonationSessionById variant
type CommandInvalidateImpersonationSessionByID struct {
	Data InvalidateImpersonationSessionByIdCommand `json:"data"`
}

func (r *CommandInvalidateImpersonationSessionByID) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandInvalidateImpersonationSessionByID) MarshalJSON() ([]byte, error) {
	type Alias CommandInvalidateImpersonationSessionByID
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "InvalidateImpersonationSessionById",
		Alias: (*Alias)(v),
	})
}

// CommandInvalidateImpersonationSessionByToken represents the InvalidateImpersonationSessionByToken variant
type CommandInvalidateImpersonationSessionByToken struct {
	Data InvalidateImpersonationSessionByTokenCommand `json:"data"`
}

func (r *CommandInvalidateImpersonationSessionByToken) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandInvalidateImpersonationSessionByToken) MarshalJSON() ([]byte, error) {
	type Alias CommandInvalidateImpersonationSessionByToken
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "InvalidateImpersonationSessionByToken",
		Alias: (*Alias)(v),
	})
}

// CommandInvalidateAllImpersonationSessionsForEmployee represents the InvalidateAllImpersonationSessionsForEmployee variant
type CommandInvalidateAllImpersonationSessionsForEmployee struct {
	Data InvalidateAllImpersonationSessionsForEmployeeCommand `json:"data"`
}

func (r *CommandInvalidateAllImpersonationSessionsForEmployee) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandInvalidateAllImpersonationSessionsForEmployee) MarshalJSON() ([]byte, error) {
	type Alias CommandInvalidateAllImpersonationSessionsForEmployee
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "InvalidateAllImpersonationSessionsForEmployee",
		Alias: (*Alias)(v),
	})
}

// CommandInvalidateAllImpersonationSessionsForUser represents the InvalidateAllImpersonationSessionsForUser variant
type CommandInvalidateAllImpersonationSessionsForUser struct {
	Data InvalidateAllImpersonationSessionsForUserCommand `json:"data"`
}

func (r *CommandInvalidateAllImpersonationSessionsForUser) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandInvalidateAllImpersonationSessionsForUser) MarshalJSON() ([]byte, error) {
	type Alias CommandInvalidateAllImpersonationSessionsForUser
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "InvalidateAllImpersonationSessionsForUser",
		Alias: (*Alias)(v),
	})
}

// CommandScimRequest represents the ScimRequest variant
type CommandScimRequest struct {
	Data ScimRequestCommand `json:"data"`
}

func (r *CommandScimRequest) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandScimRequest) MarshalJSON() ([]byte, error) {
	type Alias CommandScimRequest
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "ScimRequest",
		Alias: (*Alias)(v),
	})
}

// CommandLinkScimUser represents the LinkScimUser variant
type CommandLinkScimUser struct {
	Data LinkScimUserCommand `json:"data"`
}

func (r *CommandLinkScimUser) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandLinkScimUser) MarshalJSON() ([]byte, error) {
	type Alias CommandLinkScimUser
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "LinkScimUser",
		Alias: (*Alias)(v),
	})
}

// CommandCommitScimUserChange represents the CommitScimUserChange variant
type CommandCommitScimUserChange struct {
	Data CommitScimUserChangeCommand `json:"data"`
}

func (r *CommandCommitScimUserChange) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandCommitScimUserChange) MarshalJSON() ([]byte, error) {
	type Alias CommandCommitScimUserChange
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "CommitScimUserChange",
		Alias: (*Alias)(v),
	})
}

// CommandCreateScimConnection represents the CreateScimConnection variant
type CommandCreateScimConnection struct {
	Data CreateScimConnectionCommand `json:"data"`
}

func (r *CommandCreateScimConnection) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandCreateScimConnection) MarshalJSON() ([]byte, error) {
	type Alias CommandCreateScimConnection
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "CreateScimConnection",
		Alias: (*Alias)(v),
	})
}

// CommandFetchScimConnection represents the FetchScimConnection variant
type CommandFetchScimConnection struct {
	Data FetchScimConnectionCommand `json:"data"`
}

func (r *CommandFetchScimConnection) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandFetchScimConnection) MarshalJSON() ([]byte, error) {
	type Alias CommandFetchScimConnection
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "FetchScimConnection",
		Alias: (*Alias)(v),
	})
}

// CommandPatchScimConnection represents the PatchScimConnection variant
type CommandPatchScimConnection struct {
	Data PatchScimConnectionCommand `json:"data"`
}

func (r *CommandPatchScimConnection) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandPatchScimConnection) MarshalJSON() ([]byte, error) {
	type Alias CommandPatchScimConnection
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "PatchScimConnection",
		Alias: (*Alias)(v),
	})
}

// CommandResetScimAPIKey represents the ResetScimApiKey variant
type CommandResetScimAPIKey struct {
	Data ResetScimApiKeyCommand `json:"data"`
}

func (r *CommandResetScimAPIKey) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandResetScimAPIKey) MarshalJSON() ([]byte, error) {
	type Alias CommandResetScimAPIKey
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "ResetScimApiKey",
		Alias: (*Alias)(v),
	})
}

// CommandDeleteScimConnection represents the DeleteScimConnection variant
type CommandDeleteScimConnection struct {
	Data DeleteScimConnectionCommand `json:"data"`
}

func (r *CommandDeleteScimConnection) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandDeleteScimConnection) MarshalJSON() ([]byte, error) {
	type Alias CommandDeleteScimConnection
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "DeleteScimConnection",
		Alias: (*Alias)(v),
	})
}

// CommandGetScimUsers represents the GetScimUsers variant
type CommandGetScimUsers struct {
	Data GetScimUsersCommand `json:"data"`
}

func (r *CommandGetScimUsers) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandGetScimUsers) MarshalJSON() ([]byte, error) {
	type Alias CommandGetScimUsers
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "GetScimUsers",
		Alias: (*Alias)(v),
	})
}

// CommandGetScimUser represents the GetScimUser variant
type CommandGetScimUser struct {
	Data GetScimUserCommand `json:"data"`
}

func (r *CommandGetScimUser) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandGetScimUser) MarshalJSON() ([]byte, error) {
	type Alias CommandGetScimUser
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "GetScimUser",
		Alias: (*Alias)(v),
	})
}

// CommandCreateOIDCClient represents the CreateOidcClient variant
type CommandCreateOIDCClient struct {
	Data CreateOidcClientCommand `json:"data"`
}

func (r *CommandCreateOIDCClient) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandCreateOIDCClient) MarshalJSON() ([]byte, error) {
	type Alias CommandCreateOIDCClient
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "CreateOidcClient",
		Alias: (*Alias)(v),
	})
}

// CommandFetchOIDCClient represents the FetchOidcClient variant
type CommandFetchOIDCClient struct {
	Data FetchOidcClientCommand `json:"data"`
}

func (r *CommandFetchOIDCClient) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandFetchOIDCClient) MarshalJSON() ([]byte, error) {
	type Alias CommandFetchOIDCClient
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "FetchOidcClient",
		Alias: (*Alias)(v),
	})
}

// CommandPatchOIDCClient represents the PatchOidcClient variant
type CommandPatchOIDCClient struct {
	Data PatchOidcClientCommand `json:"data"`
}

func (r *CommandPatchOIDCClient) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandPatchOIDCClient) MarshalJSON() ([]byte, error) {
	type Alias CommandPatchOIDCClient
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "PatchOidcClient",
		Alias: (*Alias)(v),
	})
}

// CommandDeleteOIDCClient represents the DeleteOidcClient variant
type CommandDeleteOIDCClient struct {
	Data DeleteOidcClientCommand `json:"data"`
}

func (r *CommandDeleteOIDCClient) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandDeleteOIDCClient) MarshalJSON() ([]byte, error) {
	type Alias CommandDeleteOIDCClient
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "DeleteOidcClient",
		Alias: (*Alias)(v),
	})
}

// CommandInitiateOIDCLogin represents the InitiateOidcLogin variant
type CommandInitiateOIDCLogin struct {
	Data InitiateOidcLoginCommand `json:"data"`
}

func (r *CommandInitiateOIDCLogin) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandInitiateOIDCLogin) MarshalJSON() ([]byte, error) {
	type Alias CommandInitiateOIDCLogin
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "InitiateOidcLogin",
		Alias: (*Alias)(v),
	})
}

// CommandCompleteOIDCLogin represents the CompleteOidcLogin variant
type CommandCompleteOIDCLogin struct {
	Data CompleteOidcLoginCommand `json:"data"`
}

func (r *CommandCompleteOIDCLogin) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandCompleteOIDCLogin) MarshalJSON() ([]byte, error) {
	type Alias CommandCompleteOIDCLogin
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "CompleteOidcLogin",
		Alias: (*Alias)(v),
	})
}

// CommandCreateAPIKey represents the CreateApiKey variant
type CommandCreateAPIKey struct {
	Data CreateApiKeyCommand `json:"data"`
}

func (r *CommandCreateAPIKey) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandCreateAPIKey) MarshalJSON() ([]byte, error) {
	type Alias CommandCreateAPIKey
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "CreateApiKey",
		Alias: (*Alias)(v),
	})
}

// CommandValidateAPIKey represents the ValidateApiKey variant
type CommandValidateAPIKey struct {
	Data ValidateApiKeyCommand `json:"data"`
}

func (r *CommandValidateAPIKey) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandValidateAPIKey) MarshalJSON() ([]byte, error) {
	type Alias CommandValidateAPIKey
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "ValidateApiKey",
		Alias: (*Alias)(v),
	})
}

// CommandRevokeAPIKey represents the RevokeApiKey variant
type CommandRevokeAPIKey struct {
	Data RevokeApiKeyCommand `json:"data"`
}

func (r *CommandRevokeAPIKey) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandRevokeAPIKey) MarshalJSON() ([]byte, error) {
	type Alias CommandRevokeAPIKey
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "RevokeApiKey",
		Alias: (*Alias)(v),
	})
}

// CommandPatchAPIKey represents the PatchApiKey variant
type CommandPatchAPIKey struct {
	Data PatchApiKeyCommand `json:"data"`
}

func (r *CommandPatchAPIKey) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandPatchAPIKey) MarshalJSON() ([]byte, error) {
	type Alias CommandPatchAPIKey
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "PatchApiKey",
		Alias: (*Alias)(v),
	})
}

// CommandGetAPIKeys represents the GetApiKeys variant
type CommandGetAPIKeys struct {
	Data GetApiKeysCommand `json:"data"`
}

func (r *CommandGetAPIKeys) isCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *CommandGetAPIKeys) MarshalJSON() ([]byte, error) {
	type Alias CommandGetAPIKeys
	return json.Marshal(&struct {
		Command string `json:"command"`
		*Alias
	}{
		Command: "GetApiKeys",
		Alias: (*Alias)(v),
	})
}



// CommitScimUserChangeCommand represents a generated type
type CommitScimUserChangeCommand struct {
	ConnectionID string `json:"connectionId"`
	CommitID string `json:"commitId"`
}


// CompleteOidcLoginCommand represents a generated type
type CompleteOidcLoginCommand struct {
	StateFromCookie *string `json:"stateFromCookie,omitempty"`
	CallbackPathAndQueryParams string `json:"callbackPathAndQueryParams"`
}


// CreateApiKeyCommand represents a generated type
type CreateApiKeyCommand struct {
	Prefix *string `json:"prefix,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ExpiresAt *int64 `json:"expiresAt,omitempty"`
	Metadata JsonValue `json:"metadata,omitempty"`
	UserID *string `json:"userId,omitempty"`
	OwnerID *string `json:"ownerId,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
}


// CreateDeviceChallengeCommand represents a generated type
type CreateDeviceChallengeCommand struct {
	UserAgent *string `json:"userAgent,omitempty"`
	IPAddress *string `json:"ipAddress,omitempty"`
}


// CreateImpersonationSessionCommand represents a generated type
type CreateImpersonationSessionCommand struct {
	EmployeeEmail string `json:"employeeEmail"`
	TargetUserID string `json:"targetUserId"`
	UserAgent string `json:"userAgent"`
	IPAddress string `json:"ipAddress"`
	Metadata JsonValue `json:"metadata,omitempty"`
}


// CreateOidcClientCommand represents a generated type
type CreateOidcClientCommand struct {
	IdpInfoFromCustomer IdpInfoFromCustomer `json:"idpInfoFromCustomer"`
	CustomerID string `json:"customerId"`
	DisplayName *string `json:"displayName,omitempty"`
	RedirectURL string `json:"redirectUrl"`
	AdditionalScopes []string `json:"additionalScopes,omitempty"`
	ScimMatchingDefinition *ScimMatchingDefinition `json:"scimMatchingDefinition,omitempty"`
	EmailDomainAllowlist []string `json:"emailDomainAllowlist,omitempty"`
}


// CreateScimConnectionCommand represents a generated type
type CreateScimConnectionCommand struct {
	CustomerID string `json:"customerId"`
	DisplayName *string `json:"displayName,omitempty"`
	ScimAPIKeyExpiration *int `json:"scimApiKeyExpiration,omitempty"`
	CustomMapping *ScimUserMappingConfig `json:"customMapping,omitempty"`
}


// CreateSessionCommand represents a generated type
type CreateSessionCommand struct {
	UserID string `json:"userId"`
	Tags []string `json:"tags,omitempty"`
	UserAgent *string `json:"userAgent,omitempty"`
	IPAddress *string `json:"ipAddress,omitempty"`
	Metadata JsonValue `json:"metadata,omitempty"`
	DeviceRegistration *DeviceRegistration `json:"deviceRegistration,omitempty"`
}


// CreateStatelessTokenCommand represents a generated type
type CreateStatelessTokenCommand struct {
	UserID string `json:"userId"`
	SessionID *string `json:"sessionId,omitempty"`
	CustomClaims JsonValue `json:"customClaims,omitempty"`
	Issuer *string `json:"issuer,omitempty"`
	Audience *string `json:"audience,omitempty"`
	NotBeforeUnixtime *int `json:"notBeforeUnixtime,omitempty"`
	LifetimeSecs *int `json:"lifetimeSecs,omitempty"`
}


// DeleteOidcClientCommand is a discriminated union response type
type DeleteOidcClientCommand interface {
	isDeleteOidcClientCommand()
}

// DeleteOidcClientCommandOIDCClientID represents the oidcClientId variant
type DeleteOidcClientCommandOIDCClientID struct {
	OIDCClientID string `json:"oidcClientId"`
}

func (r *DeleteOidcClientCommandOIDCClientID) isDeleteOidcClientCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *DeleteOidcClientCommandOIDCClientID) MarshalJSON() ([]byte, error) {
	type Alias DeleteOidcClientCommandOIDCClientID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "oidcClientId",
		Alias: (*Alias)(v),
	})
}

// DeleteOidcClientCommandCustomerID represents the customerId variant
type DeleteOidcClientCommandCustomerID struct {
	CustomerID string `json:"customerId"`
}

func (r *DeleteOidcClientCommandCustomerID) isDeleteOidcClientCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *DeleteOidcClientCommandCustomerID) MarshalJSON() ([]byte, error) {
	type Alias DeleteOidcClientCommandCustomerID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "customerId",
		Alias: (*Alias)(v),
	})
}



// DeleteScimConnectionCommand is a discriminated union response type
type DeleteScimConnectionCommand interface {
	isDeleteScimConnectionCommand()
}

// DeleteScimConnectionCommandScimConnectionID represents the scimConnectionId variant
type DeleteScimConnectionCommandScimConnectionID struct {
	ScimConnectionID string `json:"scimConnectionId"`
}

func (r *DeleteScimConnectionCommandScimConnectionID) isDeleteScimConnectionCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *DeleteScimConnectionCommandScimConnectionID) MarshalJSON() ([]byte, error) {
	type Alias DeleteScimConnectionCommandScimConnectionID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "scimConnectionId",
		Alias: (*Alias)(v),
	})
}

// DeleteScimConnectionCommandCustomerID represents the customerId variant
type DeleteScimConnectionCommandCustomerID struct {
	CustomerID string `json:"customerId"`
}

func (r *DeleteScimConnectionCommandCustomerID) isDeleteScimConnectionCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *DeleteScimConnectionCommandCustomerID) MarshalJSON() ([]byte, error) {
	type Alias DeleteScimConnectionCommandCustomerID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "customerId",
		Alias: (*Alias)(v),
	})
}



// DeregisterAllPasskeysForUserCommand represents a generated type
type DeregisterAllPasskeysForUserCommand struct {
	UserID string `json:"userId"`
}


// DeregisterPasskeyCommand represents a generated type
type DeregisterPasskeyCommand struct {
	UserID string `json:"userId"`
	CredentialID string `json:"credentialId"`
}


// FetchAllActiveImpersonationSessionsCommand represents a generated type
type FetchAllActiveImpersonationSessionsCommand struct {
	PagingToken *string `json:"pagingToken,omitempty"`
	EmployeeEmail *string `json:"employeeEmail,omitempty"`
	TargetUserID *string `json:"targetUserId,omitempty"`
}


// FetchAllImpersonationSessionsForEmployeeCommand represents a generated type
type FetchAllImpersonationSessionsForEmployeeCommand struct {
	EmployeeEmail string `json:"employeeEmail"`
}


// FetchAllImpersonationSessionsForUserCommand represents a generated type
type FetchAllImpersonationSessionsForUserCommand struct {
	UserID string `json:"userId"`
}


// FetchAllPasskeysForUserCommand represents a generated type
type FetchAllPasskeysForUserCommand struct {
	UserID string `json:"userId"`
}


// FetchAllSessionsCommand represents a generated type
type FetchAllSessionsCommand struct {
	UserID *string `json:"userId,omitempty"`
	SessionTags []string `json:"sessionTags,omitempty"`
	Page *int `json:"page,omitempty"`
}


// FetchAllSessionsForUserCommand represents a generated type
type FetchAllSessionsForUserCommand struct {
	UserID string `json:"userId"`
	SessionTags []string `json:"sessionTags,omitempty"`
}


// FetchImpersonationSessionByIdCommand represents a generated type
type FetchImpersonationSessionByIdCommand struct {
	ImpersonationSessionID string `json:"impersonationSessionId"`
}


// FetchOidcClientCommand is a discriminated union response type
type FetchOidcClientCommand interface {
	isFetchOidcClientCommand()
}

// FetchOidcClientCommandOIDCClientID represents the oidcClientId variant
type FetchOidcClientCommandOIDCClientID struct {
	OIDCClientID string `json:"oidcClientId"`
}

func (r *FetchOidcClientCommandOIDCClientID) isFetchOidcClientCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *FetchOidcClientCommandOIDCClientID) MarshalJSON() ([]byte, error) {
	type Alias FetchOidcClientCommandOIDCClientID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "oidcClientId",
		Alias: (*Alias)(v),
	})
}

// FetchOidcClientCommandCustomerID represents the customerId variant
type FetchOidcClientCommandCustomerID struct {
	CustomerID string `json:"customerId"`
}

func (r *FetchOidcClientCommandCustomerID) isFetchOidcClientCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *FetchOidcClientCommandCustomerID) MarshalJSON() ([]byte, error) {
	type Alias FetchOidcClientCommandCustomerID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "customerId",
		Alias: (*Alias)(v),
	})
}



// FetchScimConnectionCommand is a discriminated union response type
type FetchScimConnectionCommand interface {
	isFetchScimConnectionCommand()
}

// FetchScimConnectionCommandScimConnectionID represents the scimConnectionId variant
type FetchScimConnectionCommandScimConnectionID struct {
	ScimConnectionID string `json:"scimConnectionId"`
}

func (r *FetchScimConnectionCommandScimConnectionID) isFetchScimConnectionCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *FetchScimConnectionCommandScimConnectionID) MarshalJSON() ([]byte, error) {
	type Alias FetchScimConnectionCommandScimConnectionID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "scimConnectionId",
		Alias: (*Alias)(v),
	})
}

// FetchScimConnectionCommandCustomerID represents the customerId variant
type FetchScimConnectionCommandCustomerID struct {
	CustomerID string `json:"customerId"`
}

func (r *FetchScimConnectionCommandCustomerID) isFetchScimConnectionCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *FetchScimConnectionCommandCustomerID) MarshalJSON() ([]byte, error) {
	type Alias FetchScimConnectionCommandCustomerID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "customerId",
		Alias: (*Alias)(v),
	})
}



// FetchSessionByIdCommand represents a generated type
type FetchSessionByIdCommand struct {
	SessionID string `json:"sessionId"`
}


// FinishPasskeyAuthenticationCommand represents a generated type
type FinishPasskeyAuthenticationCommand struct {
	UserID string `json:"userId"`
	PublicKey JsonValue `json:"publicKey"`
	AdditionalAllowedOrigin *string `json:"additionalAllowedOrigin,omitempty"`
}


// FinishPasskeyRegistrationCommand represents a generated type
type FinishPasskeyRegistrationCommand struct {
	UserID string `json:"userId"`
	PublicKey JsonValue `json:"publicKey"`
	AdditionalAllowedOrigin *string `json:"additionalAllowedOrigin,omitempty"`
}


// GetApiKeysCommand represents a generated type
type GetApiKeysCommand struct {
	UserID *string `json:"userId,omitempty"`
	OwnerID *string `json:"ownerId,omitempty"`
	PageNumber *int `json:"pageNumber,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
	Prefix *string `json:"prefix,omitempty"`
	Status *ApiKeyStatus `json:"status,omitempty"`
}


// GetJwksCommand represents a generated type
type GetJwksCommand struct {
}


// GetScimUserCommand is a discriminated union response type
type GetScimUserCommand interface {
	isGetScimUserCommand()
}

// GetScimUserCommandScimConnectionID represents the scimConnectionId variant
type GetScimUserCommandScimConnectionID struct {
	UserID string `json:"userId"`
	ScimConnectionID string `json:"scimConnectionId"`
}

func (r *GetScimUserCommandScimConnectionID) isGetScimUserCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *GetScimUserCommandScimConnectionID) MarshalJSON() ([]byte, error) {
	type Alias GetScimUserCommandScimConnectionID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "scimConnectionId",
		Alias: (*Alias)(v),
	})
}

// GetScimUserCommandCustomerID represents the customerId variant
type GetScimUserCommandCustomerID struct {
	UserID string `json:"userId"`
	CustomerID string `json:"customerId"`
}

func (r *GetScimUserCommandCustomerID) isGetScimUserCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *GetScimUserCommandCustomerID) MarshalJSON() ([]byte, error) {
	type Alias GetScimUserCommandCustomerID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "customerId",
		Alias: (*Alias)(v),
	})
}



// GetScimUsersCommand is a discriminated union response type
type GetScimUsersCommand interface {
	isGetScimUsersCommand()
}

// GetScimUsersCommandScimConnectionID represents the scimConnectionId variant
type GetScimUsersCommandScimConnectionID struct {
	Filter *ScimUsersPageEqualityFilter `json:"filter,omitempty"`
	PageNumber *int `json:"pageNumber,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
	ScimConnectionID string `json:"scimConnectionId"`
}

func (r *GetScimUsersCommandScimConnectionID) isGetScimUsersCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *GetScimUsersCommandScimConnectionID) MarshalJSON() ([]byte, error) {
	type Alias GetScimUsersCommandScimConnectionID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "scimConnectionId",
		Alias: (*Alias)(v),
	})
}

// GetScimUsersCommandCustomerID represents the customerId variant
type GetScimUsersCommandCustomerID struct {
	Filter *ScimUsersPageEqualityFilter `json:"filter,omitempty"`
	PageNumber *int `json:"pageNumber,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
	CustomerID string `json:"customerId"`
}

func (r *GetScimUsersCommandCustomerID) isGetScimUsersCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *GetScimUsersCommandCustomerID) MarshalJSON() ([]byte, error) {
	type Alias GetScimUsersCommandCustomerID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "customerId",
		Alias: (*Alias)(v),
	})
}



// InitiateOidcLoginCommand is a discriminated union response type
type InitiateOidcLoginCommand interface {
	isInitiateOidcLoginCommand()
}

// InitiateOidcLoginCommandOIDCClientID represents the oidcClientId variant
type InitiateOidcLoginCommandOIDCClientID struct {
	PostLoginRedirectURL *string `json:"postLoginRedirectUrl,omitempty"`
	OIDCClientID string `json:"oidcClientId"`
}

func (r *InitiateOidcLoginCommandOIDCClientID) isInitiateOidcLoginCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *InitiateOidcLoginCommandOIDCClientID) MarshalJSON() ([]byte, error) {
	type Alias InitiateOidcLoginCommandOIDCClientID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "oidcClientId",
		Alias: (*Alias)(v),
	})
}

// InitiateOidcLoginCommandCustomerID represents the customerId variant
type InitiateOidcLoginCommandCustomerID struct {
	PostLoginRedirectURL *string `json:"postLoginRedirectUrl,omitempty"`
	CustomerID string `json:"customerId"`
}

func (r *InitiateOidcLoginCommandCustomerID) isInitiateOidcLoginCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *InitiateOidcLoginCommandCustomerID) MarshalJSON() ([]byte, error) {
	type Alias InitiateOidcLoginCommandCustomerID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "customerId",
		Alias: (*Alias)(v),
	})
}



// InvalidateAllImpersonationSessionsForEmployeeCommand represents a generated type
type InvalidateAllImpersonationSessionsForEmployeeCommand struct {
	EmployeeEmail string `json:"employeeEmail"`
}


// InvalidateAllImpersonationSessionsForUserCommand represents a generated type
type InvalidateAllImpersonationSessionsForUserCommand struct {
	UserID string `json:"userId"`
}


// InvalidateAllSessionsForUserCommand represents a generated type
type InvalidateAllSessionsForUserCommand struct {
	UserID string `json:"userId"`
	SessionTags []string `json:"sessionTags,omitempty"`
}


// InvalidateAllSessionsForUserExceptOneCommand represents a generated type
type InvalidateAllSessionsForUserExceptOneCommand struct {
	UserID string `json:"userId"`
	SessionTokenToKeep string `json:"sessionTokenToKeep"`
	SessionTags []string `json:"sessionTags,omitempty"`
}


// InvalidateImpersonationSessionByIdCommand represents a generated type
type InvalidateImpersonationSessionByIdCommand struct {
	ImpersonationSessionID string `json:"impersonationSessionId"`
}


// InvalidateImpersonationSessionByTokenCommand represents a generated type
type InvalidateImpersonationSessionByTokenCommand struct {
	ImpersonationSessionToken string `json:"impersonationSessionToken"`
}


// InvalidateSessionByIdCommand represents a generated type
type InvalidateSessionByIdCommand struct {
	SessionID string `json:"sessionId"`
	UserID *string `json:"userId,omitempty"`
}


// InvalidateSessionByTokenCommand represents a generated type
type InvalidateSessionByTokenCommand struct {
	SessionToken *string `json:"sessionToken,omitempty"`
}


// LinkScimUserCommand represents a generated type
type LinkScimUserCommand struct {
	ConnectionID string `json:"connectionId"`
	CommitID string `json:"commitId"`
	UserID string `json:"userId"`
}


// PatchApiKeyCommand represents a generated type
type PatchApiKeyCommand struct {
	KeyID string `json:"keyId"`
	DisplayName *string `json:"displayName,omitempty"`
	ExpiresAt *int64 `json:"expiresAt,omitempty"`
	SetToNeverExpire *bool `json:"setToNeverExpire,omitempty"`
	Metadata JsonValue `json:"metadata,omitempty"`
}


// PatchOidcClientCommand is a discriminated union response type
type PatchOidcClientCommand interface {
	isPatchOidcClientCommand()
}

// PatchOidcClientCommandOIDCClientID represents the oidcClientId variant
type PatchOidcClientCommandOIDCClientID struct {
	IdpInfoFromCustomer *OptionalIdpInfoFromCustomer `json:"idpInfoFromCustomer,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	EmailDomainAllowlist []string `json:"emailDomainAllowlist,omitempty"`
	RedirectURL *string `json:"redirectUrl,omitempty"`
	AdditionalScopes []string `json:"additionalScopes,omitempty"`
	ScimMatchingDefinition *ScimMatchingDefinition `json:"scimMatchingDefinition,omitempty"`
	OIDCClientID string `json:"oidcClientId"`
}

func (r *PatchOidcClientCommandOIDCClientID) isPatchOidcClientCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *PatchOidcClientCommandOIDCClientID) MarshalJSON() ([]byte, error) {
	type Alias PatchOidcClientCommandOIDCClientID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "oidcClientId",
		Alias: (*Alias)(v),
	})
}

// PatchOidcClientCommandCustomerID represents the customerId variant
type PatchOidcClientCommandCustomerID struct {
	IdpInfoFromCustomer *OptionalIdpInfoFromCustomer `json:"idpInfoFromCustomer,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	EmailDomainAllowlist []string `json:"emailDomainAllowlist,omitempty"`
	RedirectURL *string `json:"redirectUrl,omitempty"`
	AdditionalScopes []string `json:"additionalScopes,omitempty"`
	ScimMatchingDefinition *ScimMatchingDefinition `json:"scimMatchingDefinition,omitempty"`
	CustomerID string `json:"customerId"`
}

func (r *PatchOidcClientCommandCustomerID) isPatchOidcClientCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *PatchOidcClientCommandCustomerID) MarshalJSON() ([]byte, error) {
	type Alias PatchOidcClientCommandCustomerID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "customerId",
		Alias: (*Alias)(v),
	})
}



// PatchScimConnectionCommand is a discriminated union response type
type PatchScimConnectionCommand interface {
	isPatchScimConnectionCommand()
}

// PatchScimConnectionCommandScimConnectionID represents the scimConnectionId variant
type PatchScimConnectionCommandScimConnectionID struct {
	DisplayName *string `json:"displayName,omitempty"`
	ScimAPIKeyExpiration *int `json:"scimApiKeyExpiration,omitempty"`
	CustomMapping *ScimUserMappingConfig `json:"customMapping,omitempty"`
	ScimConnectionID string `json:"scimConnectionId"`
}

func (r *PatchScimConnectionCommandScimConnectionID) isPatchScimConnectionCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *PatchScimConnectionCommandScimConnectionID) MarshalJSON() ([]byte, error) {
	type Alias PatchScimConnectionCommandScimConnectionID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "scimConnectionId",
		Alias: (*Alias)(v),
	})
}

// PatchScimConnectionCommandCustomerID represents the customerId variant
type PatchScimConnectionCommandCustomerID struct {
	DisplayName *string `json:"displayName,omitempty"`
	ScimAPIKeyExpiration *int `json:"scimApiKeyExpiration,omitempty"`
	CustomMapping *ScimUserMappingConfig `json:"customMapping,omitempty"`
	CustomerID string `json:"customerId"`
}

func (r *PatchScimConnectionCommandCustomerID) isPatchScimConnectionCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *PatchScimConnectionCommandCustomerID) MarshalJSON() ([]byte, error) {
	type Alias PatchScimConnectionCommandCustomerID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "customerId",
		Alias: (*Alias)(v),
	})
}



// PingCommand represents a generated type
type PingCommand struct {
}


// RegisterDeviceCommand represents a generated type
type RegisterDeviceCommand struct {
	SessionToken *string `json:"sessionToken,omitempty"`
	SessionID *string `json:"sessionId,omitempty"`
	SignedDeviceChallenge string `json:"signedDeviceChallenge"`
	RememberDevice bool `json:"rememberDevice"`
	RequestURL *string `json:"requestUrl,omitempty"`
	RequestMethod *string `json:"requestMethod,omitempty"`
	UserAgent *string `json:"userAgent,omitempty"`
	IPAddress *string `json:"ipAddress,omitempty"`
}


// ResetScimApiKeyCommand is a discriminated union response type
type ResetScimApiKeyCommand interface {
	isResetScimApiKeyCommand()
}

// ResetScimApiKeyCommandScimConnectionID represents the scimConnectionId variant
type ResetScimApiKeyCommandScimConnectionID struct {
	ScimAPIKeyExpiration *int `json:"scimApiKeyExpiration,omitempty"`
	ScimConnectionID string `json:"scimConnectionId"`
}

func (r *ResetScimApiKeyCommandScimConnectionID) isResetScimApiKeyCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *ResetScimApiKeyCommandScimConnectionID) MarshalJSON() ([]byte, error) {
	type Alias ResetScimApiKeyCommandScimConnectionID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "scimConnectionId",
		Alias: (*Alias)(v),
	})
}

// ResetScimApiKeyCommandCustomerID represents the customerId variant
type ResetScimApiKeyCommandCustomerID struct {
	ScimAPIKeyExpiration *int `json:"scimApiKeyExpiration,omitempty"`
	CustomerID string `json:"customerId"`
}

func (r *ResetScimApiKeyCommandCustomerID) isResetScimApiKeyCommand() {}

// MarshalJSON implements json.Marshaler to include the discriminator field
func (v *ResetScimApiKeyCommandCustomerID) MarshalJSON() ([]byte, error) {
	type Alias ResetScimApiKeyCommandCustomerID
	return json.Marshal(&struct {
		FieldPresence string `json:"__fieldPresence"`
		*Alias
	}{
		FieldPresence: "customerId",
		Alias: (*Alias)(v),
	})
}



// RevokeApiKeyCommand represents a generated type
type RevokeApiKeyCommand struct {
	KeyID string `json:"keyId"`
}


// RotateStatelessTokenKeyCommand represents a generated type
type RotateStatelessTokenKeyCommand struct {
	SecsBeforeNewKeyBecomesDefault int `json:"secsBeforeNewKeyBecomesDefault"`
	SecsBeforeExistingKeysAreDeactivated int `json:"secsBeforeExistingKeysAreDeactivated"`
}


// ScimRequestCommand represents a generated type
type ScimRequestCommand struct {
	Method string `json:"method"`
	PathAndQueryParams string `json:"pathAndQueryParams"`
	Body JsonValue `json:"body,omitempty"`
	ScimAPIKey *string `json:"scimApiKey,omitempty"`
}


// StartPasskeyAuthenticationCommand represents a generated type
type StartPasskeyAuthenticationCommand struct {
	UserID string `json:"userId"`
	AdditionalAllowedOrigin *string `json:"additionalAllowedOrigin,omitempty"`
}


// StartPasskeyRegistrationCommand represents a generated type
type StartPasskeyRegistrationCommand struct {
	UserID string `json:"userId"`
	EmailOrUsername string `json:"emailOrUsername"`
	UserDisplayName *string `json:"userDisplayName,omitempty"`
	PasskeyDisplayName *string `json:"passkeyDisplayName,omitempty"`
	AdditionalAllowedOrigin *string `json:"additionalAllowedOrigin,omitempty"`
}


// UpdateSessionCommand represents a generated type
type UpdateSessionCommand struct {
	SessionID string `json:"sessionId"`
	TagsToRemove []string `json:"tagsToRemove,omitempty"`
	TagsToAdd []string `json:"tagsToAdd,omitempty"`
	NewMetadata JsonValue `json:"newMetadata,omitempty"`
	PatchMetadata JsonValue `json:"patchMetadata,omitempty"`
}


// UpdateSessionsCommand represents a generated type
type UpdateSessionsCommand struct {
	Filter SessionsFilter `json:"filter"`
	TagsToRemove []string `json:"tagsToRemove,omitempty"`
	TagsToAdd []string `json:"tagsToAdd,omitempty"`
	NewMetadata JsonValue `json:"newMetadata,omitempty"`
	PatchMetadata JsonValue `json:"patchMetadata,omitempty"`
}


// ValidateAndRefreshSessionCommand represents a generated type
type ValidateAndRefreshSessionCommand struct {
	SessionToken *string `json:"sessionToken,omitempty"`
	UserAgent *string `json:"userAgent,omitempty"`
	IPAddress *string `json:"ipAddress,omitempty"`
	RequiredTags []string `json:"requiredTags,omitempty"`
	DeviceVerification *DeviceVerification `json:"deviceVerification,omitempty"`
	IgnoreDeviceForVerification *bool `json:"ignoreDeviceForVerification,omitempty"`
}


// ValidateApiKeyCommand represents a generated type
type ValidateApiKeyCommand struct {
	Key string `json:"key"`
}


// ValidateImpersonationSessionCommand represents a generated type
type ValidateImpersonationSessionCommand struct {
	ImpersonationToken string `json:"impersonationToken"`
	UserAgent string `json:"userAgent"`
	IPAddress string `json:"ipAddress"`
}


// ValidateSessionCommand represents a generated type
type ValidateSessionCommand struct {
	SessionToken *string `json:"sessionToken,omitempty"`
	UserAgent *string `json:"userAgent,omitempty"`
	IPAddress *string `json:"ipAddress,omitempty"`
	RequiredTags []string `json:"requiredTags,omitempty"`
	DeviceVerification *DeviceVerification `json:"deviceVerification,omitempty"`
	IgnoreDeviceForVerification *bool `json:"ignoreDeviceForVerification,omitempty"`
}
