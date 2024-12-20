/*
Copyright © 2023 netr0m <netr0m@pm.me>
*/
package pim

import "github.com/golang-jwt/jwt/v5"

type AzureUserInfo struct {
	ObjectId string `json:"oid"`
	Email    string `json:"unique_name"`
}

type AzureUserInfoClaims struct {
	*jwt.MapClaims
	*AzureUserInfo
}

type PIMRequest struct {
	Url     string
	Token   string
	Method  string
	Headers map[string][]string
	Payload interface{}
	Params  map[string]string
}

type ResourceExpandedProperty struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	Type        string `json:"type"`
	Email       string `json:"email"`
}

type ResourceExpandedProperties struct {
	Principal      *ResourceExpandedProperty `json:"principal"`
	RoleDefinition *ResourceExpandedProperty `json:"roleDefinition"`
	Scope          *ResourceExpandedProperty `json:"scope"`
}

type ResourceProperties struct {
	RoleEligibilityScheduleId string                      `json:"roleEligibilityScheduleId"`
	Scope                     string                      `json:"scope"`
	RoleDefinitionId          string                      `json:"roleDefinitionId"`
	PrincipalId               string                      `json:"principalId"`
	PrincipalType             string                      `json:"principalType"`
	Status                    string                      `json:"status"`
	StartDateTime             string                      `json:"startDateTime"`
	EndDateTime               string                      `json:"endDateTime"`
	MemberType                string                      `json:"memberType"`
	CreatedOn                 string                      `json:"createdOn"`
	Condition                 string                      `json:"condition"`
	ConditionVersion          string                      `json:"conditionVersion"`
	ExpandedProperties        *ResourceExpandedProperties `json:"expandedProperties"`
}

type ResourceAssignment struct {
	Properties *ResourceProperties `json:"properties"`
	Name       string              `json:"name"`
	Id         string              `json:"id"`
	Type       string              `json:"type"`
}

type ResourceAssignmentResponse struct {
	Value []ResourceAssignment `json:"value"`
}

type GovernanceRoleAssignmentSubject struct {
	Id            string `json:"id"`
	Type          string `json:"type"`
	DisplayName   string `json:"displayName"`
	PrincipalName string `json:"principalName"`
	Email         string `json:"email"`
}

type GovernanceRoleResource struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	DisplayName string `json:"displayName"`
	Status      string `json:"status"`
}

type GovernanceRoleDefinition struct {
	Id          string                  `json:"id"`
	ResourceId  string                  `json:"resourceId"`
	Type        string                  `json:"type"`
	DisplayName string                  `json:"displayName"`
	Resource    *GovernanceRoleResource `json:"resource"`
}

type GovernanceRoleAssignment struct {
	Id               string                           `json:"id"`
	ResourceId       string                           `json:"resourceId"`
	RoleDefinitionId string                           `json:"roleDefinitionId"`
	SubjectId        string                           `json:"subjectId"`
	AssignmentState  string                           `json:"assignmentState"`
	Status           string                           `json:"status"`
	Subject          *GovernanceRoleAssignmentSubject `json:"subject"`
	RoleDefinition   *GovernanceRoleDefinition        `json:"roleDefinition"`
}

type GovernanceRoleAssignmentResponse struct {
	Value []GovernanceRoleAssignment `json:"value"`
}

type TicketInfo struct {
	TicketNumber string `json:"ticketNumber"`
	TicketSystem string `json:"ticketSystem"`
}

type ScheduleInfoExpiration struct {
	Type     string `json:"type"`
	Duration string `json:"duration"`
}

type ScheduleInfo struct {
	StartDateTime interface{}             `json:"startDateTime"`
	Expiration    *ScheduleInfoExpiration `json:"expiration"`
	EndDateTime   interface{}             `json:"endDateTime"`
}

const (
	StatusAccepted                    string = "Accepted"
	StatusAdminApproved               string = "AdminApproved"
	StatusAdminDenied                 string = "AdminDenied"
	StatusCanceled                    string = "Canceled"
	StatusDenied                      string = "Denied"
	StatusFailed                      string = "Failed"
	StatusFailedAsResourceIsLocked    string = "FailedAsResourceIsLocked"
	StatusGranted                     string = "Granted"
	StatusInvalid                     string = "Invalid"
	StatusPendingAdminDecision        string = "PendingAdminDecision"
	StatusPendingApproval             string = "PendingApproval"
	StatusPendingApprovalProvisioning string = "PendingApprovalProvisioning"
	StatusPendingEvaluation           string = "PendingEvaluation"
	StatusPendingExternalProvisioning string = "PendingExternalProvisioning"
	StatusPendingProvisioning         string = "PendingProvisioning"
	StatusPendingRevocation           string = "PendingRevocation"
	StatusPendingScheduleCreation     string = "PendingScheduleCreation"
	StatusProvisioned                 string = "Provisioned"
	StatusProvisioningStarted         string = "ProvisioningStarted"
	StatusRevoked                     string = "Revoked"
	StatusScheduleCreated             string = "ScheduleCreated"
	StatusTimedOut                    string = "TimedOut"
)

type ResourceAssignmentValidationProperties struct {
	LinkedRoleEligibilityScheduleId string                      `json:"linkedRoleEligibilityScheduleId"`
	TargetRoleAssignmentScheduleId  string                      `json:"targetRoleAssignmentScheduleId"`
	Scope                           string                      `json:"scope"`
	RoleDefinitionId                string                      `json:"roleDefinitionId"`
	PrincipalId                     string                      `json:"principalId"`
	PrincipalType                   string                      `json:"principalType"`
	RequestType                     string                      `json:"requestType"`
	Status                          string                      `json:"status"`
	ScheduleInfo                    *ScheduleInfo               `json:"scheduleInfo"`
	TicketInfo                      *TicketInfo                 `json:"ticketInfo"`
	Justification                   string                      `json:"justification"`
	RequestorId                     string                      `json:"requestorId"`
	CreatedOn                       string                      `json:"createdOn"`
	ExpandedProperties              *ResourceExpandedProperties `json:"expandedProperties"`
}

type ResourceAssignmentRequestResponse struct {
	Properties *ResourceAssignmentValidationProperties `json:"properties"`
	Name       string                                  `json:"name"`
	Id         string                                  `json:"id"`
	Type       string                                  `json:"type"`
}

type ResourceAssignmentRequestProperties struct {
	PrincipalId                     string        `json:"PrincipalId"`
	RoleDefinitionId                string        `json:"RoleDefinitionId"`
	RequestType                     string        `json:"RequestType"`
	LinkedRoleEligibilityScheduleId string        `json:"LinkedRoleEligibilityScheduleId"`
	Justification                   string        `json:"Justification"`
	ScheduleInfo                    *ScheduleInfo `json:"ScheduleInfo"`
	TicketInfo                      *TicketInfo   `json:"TicketInfo"`
	IsValidationOnly                bool          `json:"IsValidationOnly"`
	IsActivativation                bool          `json:"IsActivativation"` // yes, this typo is in the API
}

type ResourceAssignmentRequestRequest struct {
	Properties ResourceAssignmentRequestProperties `json:"Properties"`
}

type GovernanceRoleAssignmentSchedule struct {
	Type          string      `json:"type"`
	StartDateTime interface{} `json:"startDateTime"`
	EndDateTime   interface{} `json:"endDateTime"`
	Duration      string      `json:"duration"`
}

type GovernanceRoleAssignmentRequest struct {
	RoleDefinitionId               string                            `json:"roleDefinitionId"`
	ResourceId                     string                            `json:"resourceId"`
	SubjectId                      string                            `json:"subjectId"`
	AssignmentState                string                            `json:"assignmentState"`
	Type                           string                            `json:"type"`
	Reason                         string                            `json:"reason"`
	TicketNumber                   string                            `json:"ticketNumber"`
	TicketSystem                   string                            `json:"ticketSystem"`
	Schedule                       *GovernanceRoleAssignmentSchedule `json:"schedule"`
	LinkedEligibleRoleAssignmentId string                            `json:"linkedEligibleRoleAssignmentId"`
	ScopedResourceId               string                            `json:"scopedResourceId"`
}

type GovernanceRoleAssignmentRequestStatus struct {
	Status        string              `json:"status"`
	SubStatus     string              `json:"subStatus"`
	StatusDetails []map[string]string `json:"statusDetails"`
}

type GovernanceRoleAssignmentRequestResponse struct {
	Id                             string                                 `json:"id"`
	ResourceId                     string                                 `json:"resourceId"`
	RoleDefinitionId               string                                 `json:"roleDefinitionId"`
	SubjectId                      string                                 `json:"subjectId"`
	ScopedResourceId               string                                 `json:"scopedResourceId"`
	LinkedEligibleRoleAssignmentId string                                 `json:"linkedEligibleRoleAssignmentId"`
	Type                           string                                 `json:"type"`
	AssignmentState                string                                 `json:"assignmentState"`
	RequestedDateTime              string                                 `json:"requestedDateTime"`
	RoleAssignmentStartDateTime    string                                 `json:"roleAssignmentStartDateTime"`
	RoleAssignmentEndDateTime      string                                 `json:"roleAssignmentEndDateTime"`
	Reason                         string                                 `json:"reason"`
	TicketNumber                   string                                 `json:"ticketNumber"`
	TicketSystem                   string                                 `json:"ticketSystem"`
	Condition                      string                                 `json:"condition"`
	ConditionVersion               string                                 `json:"conditionVersion"`
	ConditionDescription           string                                 `json:"conditionDescription"`
	Status                         *GovernanceRoleAssignmentRequestStatus `json:"status"`
	Schedule                       *GovernanceRoleAssignmentSchedule      `json:"schedule"`
	Metadata                       map[string]interface{}                 `json:"metadata"`
}
