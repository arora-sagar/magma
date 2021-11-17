// Package TS29122CpProvisioning provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package TS29122CpProvisioning

import (
	externalRef0 "magma/feg/gateway/sbi/specs/TS29122CommonData"
)

const (
	OAuth2ClientCredentialsScopes = "oAuth2ClientCredentials.Scopes"
)

// Possible values are - PERIODICALLY: Identifies the UE communicates periodically - ON_DEMAND: Identifies the UE communicates on demand
type CommunicationIndicator interface{}

// Possible values are - MALFUNCTION: This value indicates that something functions wrongly in CP parameter provisioning or the CP parameter provisioning does not function at all. - SET_ID_DUPLICATED: The received CP set identifier(s) are already provisioned. - OTHER_REASON: Other reason unspecified.
type CpFailureCode interface{}

// CpInfo defines model for CpInfo.
type CpInfo interface{}

// CpParameterSet defines model for CpParameterSet.
type CpParameterSet struct {
	// Unsigned integer identifying a period of time in units of seconds.
	CommunicationDurationTime *externalRef0.DurationSec `json:"communicationDurationTime,omitempty"`

	// Identifies the UE's expected geographical movement. The attribute is only applicable in 5G.
	ExpectedUmts *[]externalRef0.LocationArea5G `json:"expectedUmts,omitempty"`

	// Possible values are - PERIODICALLY: Identifies the UE communicates periodically - ON_DEMAND: Identifies the UE communicates on demand
	PeriodicCommunicationIndicator *CommunicationIndicator `json:"periodicCommunicationIndicator,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	PeriodicTime               *externalRef0.DurationSec   `json:"periodicTime,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self *externalRef0.Link `json:"self,omitempty"`

	// SCS/AS-chosen correlator provided by the SCS/AS in the request to create a resource fo CP parameter set(s).
	SetId string `json:"setId"`

	// Possible values are - STATIONARY: Identifies the UE is stationary - MOBILE: Identifies the UE is mobile
	StationaryIndication *StationaryIndication `json:"stationaryIndication,omitempty"`

	// string with format "date-time" as defined in OpenAPI.
	ValidityTime *externalRef0.DateTime `json:"validityTime,omitempty"`
}

// CpReport defines model for CpReport.
type CpReport struct {
	// Possible values are - MALFUNCTION: This value indicates that something functions wrongly in CP parameter provisioning or the CP parameter provisioning does not function at all. - SET_ID_DUPLICATED: The received CP set identifier(s) are already provisioned. - OTHER_REASON: Other reason unspecified.
	FailureCode CpFailureCode `json:"failureCode"`

	// Identifies the CP set identifier(s) which CP parameter(s) are not added or modified successfully
	SetIds *[]string `json:"setIds,omitempty"`
}

// ScheduledCommunicationTime defines model for ScheduledCommunicationTime.
type ScheduledCommunicationTime struct {
	// Identifies the day(s) of the week. If absent, it indicates every day of the week.
	DaysOfWeek *[]externalRef0.DayOfWeek `json:"daysOfWeek,omitempty"`

	// String with format partial-time or full-time as defined in subclause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC).
	TimeOfDayEnd *externalRef0.TimeOfDay `json:"timeOfDayEnd,omitempty"`

	// String with format partial-time or full-time as defined in subclause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC).
	TimeOfDayStart *externalRef0.TimeOfDay `json:"timeOfDayStart,omitempty"`
}

// Possible values are - STATIONARY: Identifies the UE is stationary - MOBILE: Identifies the UE is mobile
type StationaryIndication interface{}

// PostScsAsIdSubscriptionsJSONBody defines parameters for PostScsAsIdSubscriptions.
type PostScsAsIdSubscriptionsJSONBody CpInfo

// PutScsAsIdSubscriptionsSubscriptionIdJSONBody defines parameters for PutScsAsIdSubscriptionsSubscriptionId.
type PutScsAsIdSubscriptionsSubscriptionIdJSONBody CpInfo

// PutScsAsIdSubscriptionsSubscriptionIdCpSetsSetIdJSONBody defines parameters for PutScsAsIdSubscriptionsSubscriptionIdCpSetsSetId.
type PutScsAsIdSubscriptionsSubscriptionIdCpSetsSetIdJSONBody CpParameterSet

// PostScsAsIdSubscriptionsJSONRequestBody defines body for PostScsAsIdSubscriptions for application/json ContentType.
type PostScsAsIdSubscriptionsJSONRequestBody PostScsAsIdSubscriptionsJSONBody

// PutScsAsIdSubscriptionsSubscriptionIdJSONRequestBody defines body for PutScsAsIdSubscriptionsSubscriptionId for application/json ContentType.
type PutScsAsIdSubscriptionsSubscriptionIdJSONRequestBody PutScsAsIdSubscriptionsSubscriptionIdJSONBody

// PutScsAsIdSubscriptionsSubscriptionIdCpSetsSetIdJSONRequestBody defines body for PutScsAsIdSubscriptionsSubscriptionIdCpSetsSetId for application/json ContentType.
type PutScsAsIdSubscriptionsSubscriptionIdCpSetsSetIdJSONRequestBody PutScsAsIdSubscriptionsSubscriptionIdCpSetsSetIdJSONBody