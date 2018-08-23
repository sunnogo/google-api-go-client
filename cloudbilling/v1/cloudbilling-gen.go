// Package cloudbilling provides access to the Cloud Billing API.
//
// See https://cloud.google.com/billing/
//
// Usage example:
//
//   import "github.com/sunnogo/google-api-go-client/cloudbilling/v1"
//   ...
//   cloudbillingService, err := cloudbilling.New(oauthHttpClient)
package cloudbilling // import "github.com/sunnogo/google-api-go-client/cloudbilling/v1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "github.com/sunnogo/net/context"
	ctxhttp "github.com/sunnogo/net/context/ctxhttp"
	gensupport "github.com/sunnogo/google-api-go-client/gensupport"
	googleapi "github.com/sunnogo/google-api-go-client/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "cloudbilling:v1"
const apiName = "cloudbilling"
const apiVersion = "v1"
const basePath = "https://cloudbilling.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

func New(client *http.Client) (*APIService, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &APIService{client: client, BasePath: basePath}
	s.BillingAccounts = NewBillingAccountsService(s)
	s.Projects = NewProjectsService(s)
	s.Services = NewServicesService(s)
	return s, nil
}

type APIService struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	BillingAccounts *BillingAccountsService

	Projects *ProjectsService

	Services *ServicesService
}

func (s *APIService) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewBillingAccountsService(s *APIService) *BillingAccountsService {
	rs := &BillingAccountsService{s: s}
	rs.Projects = NewBillingAccountsProjectsService(s)
	return rs
}

type BillingAccountsService struct {
	s *APIService

	Projects *BillingAccountsProjectsService
}

func NewBillingAccountsProjectsService(s *APIService) *BillingAccountsProjectsService {
	rs := &BillingAccountsProjectsService{s: s}
	return rs
}

type BillingAccountsProjectsService struct {
	s *APIService
}

func NewProjectsService(s *APIService) *ProjectsService {
	rs := &ProjectsService{s: s}
	return rs
}

type ProjectsService struct {
	s *APIService
}

func NewServicesService(s *APIService) *ServicesService {
	rs := &ServicesService{s: s}
	rs.Skus = NewServicesSkusService(s)
	return rs
}

type ServicesService struct {
	s *APIService

	Skus *ServicesSkusService
}

func NewServicesSkusService(s *APIService) *ServicesSkusService {
	rs := &ServicesSkusService{s: s}
	return rs
}

type ServicesSkusService struct {
	s *APIService
}

// AggregationInfo: Represents the aggregation level and interval for
// pricing of a single SKU.
type AggregationInfo struct {
	// AggregationCount: The number of intervals to aggregate over.
	// Example: If aggregation_level is "DAILY" and aggregation_count is
	// 14,
	// aggregation will be over 14 days.
	AggregationCount int64 `json:"aggregationCount,omitempty"`

	// Possible values:
	//   "AGGREGATION_INTERVAL_UNSPECIFIED"
	//   "DAILY"
	//   "MONTHLY"
	AggregationInterval string `json:"aggregationInterval,omitempty"`

	// Possible values:
	//   "AGGREGATION_LEVEL_UNSPECIFIED"
	//   "ACCOUNT"
	//   "PROJECT"
	AggregationLevel string `json:"aggregationLevel,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AggregationCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AggregationCount") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AggregationInfo) MarshalJSON() ([]byte, error) {
	type NoMethod AggregationInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AuditConfig: Specifies the audit configuration for a service.
// The configuration determines which permission types are logged, and
// what
// identities, if any, are exempted from logging.
// An AuditConfig must have one or more AuditLogConfigs.
//
// If there are AuditConfigs for both `allServices` and a specific
// service,
// the union of the two AuditConfigs is used for that service: the
// log_types
// specified in each AuditConfig are enabled, and the exempted_members
// in each
// AuditLogConfig are exempted.
//
// Example Policy with multiple AuditConfigs:
//
//     {
//       "audit_configs": [
//         {
//           "service": "allServices"
//           "audit_log_configs": [
//             {
//               "log_type": "DATA_READ",
//               "exempted_members": [
//                 "user:foo@gmail.com"
//               ]
//             },
//             {
//               "log_type": "DATA_WRITE",
//             },
//             {
//               "log_type": "ADMIN_READ",
//             }
//           ]
//         },
//         {
//           "service": "fooservice.googleapis.com"
//           "audit_log_configs": [
//             {
//               "log_type": "DATA_READ",
//             },
//             {
//               "log_type": "DATA_WRITE",
//               "exempted_members": [
//                 "user:bar@gmail.com"
//               ]
//             }
//           ]
//         }
//       ]
//     }
//
// For fooservice, this policy enables DATA_READ, DATA_WRITE and
// ADMIN_READ
// logging. It also exempts foo@gmail.com from DATA_READ logging,
// and
// bar@gmail.com from DATA_WRITE logging.
type AuditConfig struct {
	// AuditLogConfigs: The configuration for logging of each type of
	// permission.
	AuditLogConfigs []*AuditLogConfig `json:"auditLogConfigs,omitempty"`

	// Service: Specifies a service that will be enabled for audit
	// logging.
	// For example, `storage.googleapis.com`,
	// `cloudsql.googleapis.com`.
	// `allServices` is a special value that covers all services.
	Service string `json:"service,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AuditLogConfigs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AuditLogConfigs") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AuditConfig) MarshalJSON() ([]byte, error) {
	type NoMethod AuditConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AuditLogConfig: Provides the configuration for logging a type of
// permissions.
// Example:
//
//     {
//       "audit_log_configs": [
//         {
//           "log_type": "DATA_READ",
//           "exempted_members": [
//             "user:foo@gmail.com"
//           ]
//         },
//         {
//           "log_type": "DATA_WRITE",
//         }
//       ]
//     }
//
// This enables 'DATA_READ' and 'DATA_WRITE' logging, while
// exempting
// foo@gmail.com from DATA_READ logging.
type AuditLogConfig struct {
	// ExemptedMembers: Specifies the identities that do not cause logging
	// for this type of
	// permission.
	// Follows the same format of Binding.members.
	ExemptedMembers []string `json:"exemptedMembers,omitempty"`

	// LogType: The log type that this config enables.
	//
	// Possible values:
	//   "LOG_TYPE_UNSPECIFIED" - Default case. Should never be this.
	//   "ADMIN_READ" - Admin reads. Example: CloudIAM getIamPolicy
	//   "DATA_WRITE" - Data writes. Example: CloudSQL Users create
	//   "DATA_READ" - Data reads. Example: CloudSQL Users list
	LogType string `json:"logType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExemptedMembers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExemptedMembers") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AuditLogConfig) MarshalJSON() ([]byte, error) {
	type NoMethod AuditLogConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BillingAccount: A billing account in [GCP
// Console](https://console.cloud.google.com/).
// You can assign a billing account to one or more projects.
type BillingAccount struct {
	// DisplayName: The display name given to the billing account, such as
	// `My Billing
	// Account`. This name is displayed in the GCP Console.
	DisplayName string `json:"displayName,omitempty"`

	// MasterBillingAccount: If this account is
	// a
	// [subaccount](https://cloud.google.com/billing/docs/concepts), then
	// this
	// will be the resource name of the master billing account that it is
	// being
	// resold through.
	// Otherwise this will be empty.
	MasterBillingAccount string `json:"masterBillingAccount,omitempty"`

	// Name: The resource name of the billing account. The resource name has
	// the form
	// `billingAccounts/{billing_account_id}`. For
	// example,
	// `billingAccounts/012345-567890-ABCDEF` would be the resource name
	// for
	// billing account `012345-567890-ABCDEF`.
	Name string `json:"name,omitempty"`

	// Open: True if the billing account is open, and will therefore be
	// charged for any
	// usage on associated projects. False if the billing account is closed,
	// and
	// therefore projects associated with it will be unable to use paid
	// services.
	Open bool `json:"open,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisplayName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BillingAccount) MarshalJSON() ([]byte, error) {
	type NoMethod BillingAccount
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Binding: Associates `members` with a `role`.
type Binding struct {
	// Condition: Unimplemented. The condition that is associated with this
	// binding.
	// NOTE: an unsatisfied condition will not allow user access via
	// current
	// binding. Different bindings, including their conditions, are
	// examined
	// independently.
	Condition *Expr `json:"condition,omitempty"`

	// Members: Specifies the identities requesting access for a Cloud
	// Platform resource.
	// `members` can have the following values:
	//
	// * `allUsers`: A special identifier that represents anyone who is
	//    on the internet; with or without a Google account.
	//
	// * `allAuthenticatedUsers`: A special identifier that represents
	// anyone
	//    who is authenticated with a Google account or a service
	// account.
	//
	// * `user:{emailid}`: An email address that represents a specific
	// Google
	//    account. For example, `alice@gmail.com` .
	//
	//
	// * `serviceAccount:{emailid}`: An email address that represents a
	// service
	//    account. For example,
	// `my-other-app@appspot.gserviceaccount.com`.
	//
	// * `group:{emailid}`: An email address that represents a Google
	// group.
	//    For example, `admins@example.com`.
	//
	//
	// * `domain:{domain}`: A Google Apps domain name that represents all
	// the
	//    users of that domain. For example, `google.com` or
	// `example.com`.
	//
	//
	Members []string `json:"members,omitempty"`

	// Role: Role that is assigned to `members`.
	// For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
	Role string `json:"role,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Condition") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Condition") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Binding) MarshalJSON() ([]byte, error) {
	type NoMethod Binding
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Category: Represents the category hierarchy of a SKU.
type Category struct {
	// ResourceFamily: The type of product the SKU refers to.
	// Example: "Compute", "Storage", "Network", "ApplicationServices" etc.
	ResourceFamily string `json:"resourceFamily,omitempty"`

	// ResourceGroup: A group classification for related SKUs.
	// Example: "RAM", "GPU", "Prediction", "Ops", "GoogleEgress" etc.
	ResourceGroup string `json:"resourceGroup,omitempty"`

	// ServiceDisplayName: The display name of the service this SKU belongs
	// to.
	ServiceDisplayName string `json:"serviceDisplayName,omitempty"`

	// UsageType: Represents how the SKU is consumed.
	// Example: "OnDemand", "Preemptible", "Commit1Mo", "Commit1Yr" etc.
	UsageType string `json:"usageType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ResourceFamily") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ResourceFamily") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Category) MarshalJSON() ([]byte, error) {
	type NoMethod Category
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Expr: Represents an expression text. Example:
//
//     title: "User account presence"
//     description: "Determines whether the request has a user account"
//     expression: "size(request.user) > 0"
type Expr struct {
	// Description: An optional description of the expression. This is a
	// longer text which
	// describes the expression, e.g. when hovered over it in a UI.
	Description string `json:"description,omitempty"`

	// Expression: Textual representation of an expression in
	// Common Expression Language syntax.
	//
	// The application context of the containing message determines
	// which
	// well-known feature set of CEL is supported.
	Expression string `json:"expression,omitempty"`

	// Location: An optional string indicating the location of the
	// expression for error
	// reporting, e.g. a file name and a position in the file.
	Location string `json:"location,omitempty"`

	// Title: An optional title for the expression, i.e. a short string
	// describing
	// its purpose. This can be used e.g. in UIs which allow to enter
	// the
	// expression.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Expr) MarshalJSON() ([]byte, error) {
	type NoMethod Expr
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListBillingAccountsResponse: Response message for
// `ListBillingAccounts`.
type ListBillingAccountsResponse struct {
	// BillingAccounts: A list of billing accounts.
	BillingAccounts []*BillingAccount `json:"billingAccounts,omitempty"`

	// NextPageToken: A token to retrieve the next page of results. To
	// retrieve the next page,
	// call `ListBillingAccounts` again with the `page_token` field set to
	// this
	// value. This field is empty if there are no more results to retrieve.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "BillingAccounts") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BillingAccounts") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ListBillingAccountsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListBillingAccountsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListProjectBillingInfoResponse: Request message for
// `ListProjectBillingInfoResponse`.
type ListProjectBillingInfoResponse struct {
	// NextPageToken: A token to retrieve the next page of results. To
	// retrieve the next page,
	// call `ListProjectBillingInfo` again with the `page_token` field set
	// to this
	// value. This field is empty if there are no more results to retrieve.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ProjectBillingInfo: A list of `ProjectBillingInfo` resources
	// representing the projects
	// associated with the billing account.
	ProjectBillingInfo []*ProjectBillingInfo `json:"projectBillingInfo,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListProjectBillingInfoResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListProjectBillingInfoResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListServicesResponse: Response message for `ListServices`.
type ListServicesResponse struct {
	// NextPageToken: A token to retrieve the next page of results. To
	// retrieve the next page,
	// call `ListServices` again with the `page_token` field set to
	// this
	// value. This field is empty if there are no more results to retrieve.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Services: A list of services.
	Services []*Service `json:"services,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListServicesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListServicesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListSkusResponse: Response message for `ListSkus`.
type ListSkusResponse struct {
	// NextPageToken: A token to retrieve the next page of results. To
	// retrieve the next page,
	// call `ListSkus` again with the `page_token` field set to this
	// value. This field is empty if there are no more results to retrieve.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Skus: The list of public SKUs of the given service.
	Skus []*Sku `json:"skus,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListSkusResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListSkusResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Money: Represents an amount of money with its currency type.
type Money struct {
	// CurrencyCode: The 3-letter currency code defined in ISO 4217.
	CurrencyCode string `json:"currencyCode,omitempty"`

	// Nanos: Number of nano (10^-9) units of the amount.
	// The value must be between -999,999,999 and +999,999,999 inclusive.
	// If `units` is positive, `nanos` must be positive or zero.
	// If `units` is zero, `nanos` can be positive, zero, or negative.
	// If `units` is negative, `nanos` must be negative or zero.
	// For example $-1.75 is represented as `units`=-1 and
	// `nanos`=-750,000,000.
	Nanos int64 `json:"nanos,omitempty"`

	// Units: The whole units of the amount.
	// For example if `currencyCode` is "USD", then 1 unit is one US
	// dollar.
	Units int64 `json:"units,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "CurrencyCode") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CurrencyCode") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Money) MarshalJSON() ([]byte, error) {
	type NoMethod Money
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Policy: Defines an Identity and Access Management (IAM) policy. It is
// used to
// specify access control policies for Cloud Platform resources.
//
//
// A `Policy` consists of a list of `bindings`. A `binding` binds a list
// of
// `members` to a `role`, where the members can be user accounts, Google
// groups,
// Google domains, and service accounts. A `role` is a named list of
// permissions
// defined by IAM.
//
// **JSON Example**
//
//     {
//       "bindings": [
//         {
//           "role": "roles/owner",
//           "members": [
//             "user:mike@example.com",
//             "group:admins@example.com",
//             "domain:google.com",
//
// "serviceAccount:my-other-app@appspot.gserviceaccount.com"
//           ]
//         },
//         {
//           "role": "roles/viewer",
//           "members": ["user:sean@example.com"]
//         }
//       ]
//     }
//
// **YAML Example**
//
//     bindings:
//     - members:
//       - user:mike@example.com
//       - group:admins@example.com
//       - domain:google.com
//       - serviceAccount:my-other-app@appspot.gserviceaccount.com
//       role: roles/owner
//     - members:
//       - user:sean@example.com
//       role: roles/viewer
//
//
// For a description of IAM and its features, see the
// [IAM developer's guide](https://cloud.google.com/iam/docs).
type Policy struct {
	// AuditConfigs: Specifies cloud audit logging configuration for this
	// policy.
	AuditConfigs []*AuditConfig `json:"auditConfigs,omitempty"`

	// Bindings: Associates a list of `members` to a `role`.
	// `bindings` with no members will result in an error.
	Bindings []*Binding `json:"bindings,omitempty"`

	// Etag: `etag` is used for optimistic concurrency control as a way to
	// help
	// prevent simultaneous updates of a policy from overwriting each
	// other.
	// It is strongly suggested that systems make use of the `etag` in
	// the
	// read-modify-write cycle to perform policy updates in order to avoid
	// race
	// conditions: An `etag` is returned in the response to `getIamPolicy`,
	// and
	// systems are expected to put that etag in the request to
	// `setIamPolicy` to
	// ensure that their change will be applied to the same version of the
	// policy.
	//
	// If no `etag` is provided in the call to `setIamPolicy`, then the
	// existing
	// policy is overwritten blindly.
	Etag string `json:"etag,omitempty"`

	// Version: Deprecated.
	Version int64 `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AuditConfigs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AuditConfigs") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Policy) MarshalJSON() ([]byte, error) {
	type NoMethod Policy
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PricingExpression: Expresses a mathematical pricing formula. For
// Example:-
//
// `usage_unit: GBy`
// `tiered_rates:`
//    `[start_usage_amount: 20, unit_price: $10]`
//    `[start_usage_amount: 100, unit_price: $5]`
//
// The above expresses a pricing formula where the first 20GB is free,
// the
// next 80GB is priced at $10 per GB followed by $5 per GB for
// additional
// usage.
type PricingExpression struct {
	// BaseUnit: The base unit for the SKU which is the unit used in usage
	// exports.
	// Example: "By"
	BaseUnit string `json:"baseUnit,omitempty"`

	// BaseUnitConversionFactor: Conversion factor for converting from price
	// per usage_unit to price per
	// base_unit, and start_usage_amount to start_usage_amount in
	// base_unit.
	// unit_price / base_unit_conversion_factor = price per
	// base_unit.
	// start_usage_amount * base_unit_conversion_factor = start_usage_amount
	// in
	// base_unit.
	BaseUnitConversionFactor float64 `json:"baseUnitConversionFactor,omitempty"`

	// BaseUnitDescription: The base unit in human readable form.
	// Example: "byte".
	BaseUnitDescription string `json:"baseUnitDescription,omitempty"`

	// DisplayQuantity: The recommended quantity of units for displaying
	// pricing info. When
	// displaying pricing info it is recommended to display:
	// (unit_price * display_quantity) per display_quantity usage_unit.
	// This field does not affect the pricing formula and is for display
	// purposes
	// only.
	// Example: If the unit_price is "0.0001 USD", the usage_unit is "GB"
	// and
	// the display_quantity is "1000" then the recommended way of displaying
	// the
	// pricing info is "0.10 USD per 1000 GB"
	DisplayQuantity float64 `json:"displayQuantity,omitempty"`

	// TieredRates: The list of tiered rates for this pricing. The total
	// cost is computed by
	// applying each of the tiered rates on usage. This repeated list is
	// sorted
	// by ascending order of start_usage_amount.
	TieredRates []*TierRate `json:"tieredRates,omitempty"`

	// UsageUnit: The short hand for unit of usage this pricing is specified
	// in.
	// Example: usage_unit of "GiBy" means that usage is specified in "Gibi
	// Byte".
	UsageUnit string `json:"usageUnit,omitempty"`

	// UsageUnitDescription: The unit of usage in human readable
	// form.
	// Example: "gibi byte".
	UsageUnitDescription string `json:"usageUnitDescription,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BaseUnit") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BaseUnit") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PricingExpression) MarshalJSON() ([]byte, error) {
	type NoMethod PricingExpression
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *PricingExpression) UnmarshalJSON(data []byte) error {
	type NoMethod PricingExpression
	var s1 struct {
		BaseUnitConversionFactor gensupport.JSONFloat64 `json:"baseUnitConversionFactor"`
		DisplayQuantity          gensupport.JSONFloat64 `json:"displayQuantity"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.BaseUnitConversionFactor = float64(s1.BaseUnitConversionFactor)
	s.DisplayQuantity = float64(s1.DisplayQuantity)
	return nil
}

// PricingInfo: Represents the pricing information for a SKU at a single
// point of time.
type PricingInfo struct {
	// AggregationInfo: Aggregation Info. This can be left unspecified if
	// the pricing expression
	// doesn't require aggregation.
	AggregationInfo *AggregationInfo `json:"aggregationInfo,omitempty"`

	// CurrencyConversionRate: Conversion rate used for currency conversion,
	// from USD to the currency
	// specified in the request. This includes any surcharge collected for
	// billing
	// in non USD currency. If a currency is not specified in the request
	// this
	// defaults to 1.0.
	// Example: USD * currency_conversion_rate = JPY
	CurrencyConversionRate float64 `json:"currencyConversionRate,omitempty"`

	// EffectiveTime: The timestamp from which this pricing was effective
	// within the requested
	// time range. This is guaranteed to be greater than or equal to
	// the
	// start_time field in the request and less than the end_time field in
	// the
	// request. If a time range was not specified in the request this field
	// will
	// be equivalent to a time within the last 12 hours, indicating the
	// latest
	// pricing info.
	EffectiveTime string `json:"effectiveTime,omitempty"`

	// PricingExpression: Expresses the pricing formula. See
	// `PricingExpression` for an example.
	PricingExpression *PricingExpression `json:"pricingExpression,omitempty"`

	// Summary: An optional human readable summary of the pricing
	// information, has a
	// maximum length of 256 characters.
	Summary string `json:"summary,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AggregationInfo") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AggregationInfo") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *PricingInfo) MarshalJSON() ([]byte, error) {
	type NoMethod PricingInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *PricingInfo) UnmarshalJSON(data []byte) error {
	type NoMethod PricingInfo
	var s1 struct {
		CurrencyConversionRate gensupport.JSONFloat64 `json:"currencyConversionRate"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.CurrencyConversionRate = float64(s1.CurrencyConversionRate)
	return nil
}

// ProjectBillingInfo: Encapsulation of billing information for a GCP
// Console project. A project
// has at most one associated billing account at a time (but a billing
// account
// can be assigned to multiple projects).
type ProjectBillingInfo struct {
	// BillingAccountName: The resource name of the billing account
	// associated with the project, if
	// any. For example, `billingAccounts/012345-567890-ABCDEF`.
	BillingAccountName string `json:"billingAccountName,omitempty"`

	// BillingEnabled: True if the project is associated with an open
	// billing account, to which
	// usage on the project is charged. False if the project is associated
	// with a
	// closed billing account, or no billing account at all, and therefore
	// cannot
	// use paid services. This field is read-only.
	BillingEnabled bool `json:"billingEnabled,omitempty"`

	// Name: The resource name for the `ProjectBillingInfo`; has the
	// form
	// `projects/{project_id}/billingInfo`. For example, the resource name
	// for the
	// billing information for project `tokyo-rain-123` would
	// be
	// `projects/tokyo-rain-123/billingInfo`. This field is read-only.
	Name string `json:"name,omitempty"`

	// ProjectId: The ID of the project that this `ProjectBillingInfo`
	// represents, such as
	// `tokyo-rain-123`. This is a convenience field so that you don't need
	// to
	// parse the `name` field to obtain a project ID. This field is
	// read-only.
	ProjectId string `json:"projectId,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "BillingAccountName")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BillingAccountName") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ProjectBillingInfo) MarshalJSON() ([]byte, error) {
	type NoMethod ProjectBillingInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Service: Encapsulates a single service in Google Cloud Platform.
type Service struct {
	// DisplayName: A human readable display name for this service.
	DisplayName string `json:"displayName,omitempty"`

	// Name: The resource name for the service.
	// Example: "services/DA34-426B-A397"
	Name string `json:"name,omitempty"`

	// ServiceId: The identifier for the service.
	// Example: "DA34-426B-A397"
	ServiceId string `json:"serviceId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisplayName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Service) MarshalJSON() ([]byte, error) {
	type NoMethod Service
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SetIamPolicyRequest: Request message for `SetIamPolicy` method.
type SetIamPolicyRequest struct {
	// Policy: REQUIRED: The complete policy to be applied to the
	// `resource`. The size of
	// the policy is limited to a few 10s of KB. An empty policy is a
	// valid policy but certain Cloud Platform services (such as
	// Projects)
	// might reject them.
	Policy *Policy `json:"policy,omitempty"`

	// UpdateMask: OPTIONAL: A FieldMask specifying which fields of the
	// policy to modify. Only
	// the fields in the mask will be modified. If no mask is provided,
	// the
	// following default mask is used:
	// paths: "bindings, etag"
	// This field is only used by Cloud IAM.
	UpdateMask string `json:"updateMask,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Policy") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Policy") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SetIamPolicyRequest) MarshalJSON() ([]byte, error) {
	type NoMethod SetIamPolicyRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Sku: Encapsulates a single SKU in Google Cloud Platform
type Sku struct {
	// Category: The category hierarchy of this SKU, purely for
	// organizational purpose.
	Category *Category `json:"category,omitempty"`

	// Description: A human readable description of the SKU, has a maximum
	// length of 256
	// characters.
	Description string `json:"description,omitempty"`

	// Name: The resource name for the SKU.
	// Example: "services/DA34-426B-A397/skus/AA95-CD31-42FE"
	Name string `json:"name,omitempty"`

	// PricingInfo: A timeline of pricing info for this SKU in chronological
	// order.
	PricingInfo []*PricingInfo `json:"pricingInfo,omitempty"`

	// ServiceProviderName: Identifies the service provider.
	// This is 'Google' for first party services in Google Cloud Platform.
	ServiceProviderName string `json:"serviceProviderName,omitempty"`

	// ServiceRegions: List of service regions this SKU is offered
	// at.
	// Example: "asia-east1"
	// Service regions can be found at
	// https://cloud.google.com/about/locations/
	ServiceRegions []string `json:"serviceRegions,omitempty"`

	// SkuId: The identifier for the SKU.
	// Example: "AA95-CD31-42FE"
	SkuId string `json:"skuId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Category") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Category") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Sku) MarshalJSON() ([]byte, error) {
	type NoMethod Sku
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TestIamPermissionsRequest: Request message for `TestIamPermissions`
// method.
type TestIamPermissionsRequest struct {
	// Permissions: The set of permissions to check for the `resource`.
	// Permissions with
	// wildcards (such as '*' or 'storage.*') are not allowed. For
	// more
	// information see
	// [IAM
	// Overview](https://cloud.google.com/iam/docs/overview#permissions).
	Permissions []string `json:"permissions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Permissions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Permissions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TestIamPermissionsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod TestIamPermissionsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TestIamPermissionsResponse: Response message for `TestIamPermissions`
// method.
type TestIamPermissionsResponse struct {
	// Permissions: A subset of `TestPermissionsRequest.permissions` that
	// the caller is
	// allowed.
	Permissions []string `json:"permissions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Permissions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Permissions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TestIamPermissionsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod TestIamPermissionsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TierRate: The price rate indicating starting usage and its
// corresponding price.
type TierRate struct {
	// StartUsageAmount: Usage is priced at this rate only after this
	// amount.
	// Example: start_usage_amount of 10 indicates that the usage will be
	// priced
	// at the unit_price after the first 10 usage_units.
	StartUsageAmount float64 `json:"startUsageAmount,omitempty"`

	// UnitPrice: The price per unit of usage.
	// Example: unit_price of amount $10 indicates that each unit will cost
	// $10.
	UnitPrice *Money `json:"unitPrice,omitempty"`

	// ForceSendFields is a list of field names (e.g. "StartUsageAmount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "StartUsageAmount") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *TierRate) MarshalJSON() ([]byte, error) {
	type NoMethod TierRate
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *TierRate) UnmarshalJSON(data []byte) error {
	type NoMethod TierRate
	var s1 struct {
		StartUsageAmount gensupport.JSONFloat64 `json:"startUsageAmount"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.StartUsageAmount = float64(s1.StartUsageAmount)
	return nil
}

// method id "cloudbilling.billingAccounts.create":

type BillingAccountsCreateCall struct {
	s              *APIService
	billingaccount *BillingAccount
	urlParams_     gensupport.URLParams
	ctx_           context.Context
	header_        http.Header
}

// Create: Creates a billing account.
// This method can only be used to create
// [billing
// subaccounts](https://cloud.google.com/billing/docs/concepts)
// by GCP resellers.
// When creating a subaccount, the current authenticated user must have
// the
// `billing.accounts.update` IAM permission on the master account, which
// is
// typically given to billing
// account
// [administrators](https://cloud.google.com/billing/docs/how-to/
// billing-access).
// This method will return an error if the master account has not
// been
// provisioned as a reseller account.
func (r *BillingAccountsService) Create(billingaccount *BillingAccount) *BillingAccountsCreateCall {
	c := &BillingAccountsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.billingaccount = billingaccount
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BillingAccountsCreateCall) Fields(s ...googleapi.Field) *BillingAccountsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BillingAccountsCreateCall) Context(ctx context.Context) *BillingAccountsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BillingAccountsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BillingAccountsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.billingaccount)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/billingAccounts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbilling.billingAccounts.create" call.
// Exactly one of *BillingAccount or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *BillingAccount.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BillingAccountsCreateCall) Do(opts ...googleapi.CallOption) (*BillingAccount, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &BillingAccount{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a billing account.\nThis method can only be used to create\n[billing subaccounts](https://cloud.google.com/billing/docs/concepts)\nby GCP resellers.\nWhen creating a subaccount, the current authenticated user must have the\n`billing.accounts.update` IAM permission on the master account, which is\ntypically given to billing account\n[administrators](https://cloud.google.com/billing/docs/how-to/billing-access).\nThis method will return an error if the master account has not been\nprovisioned as a reseller account.",
	//   "flatPath": "v1/billingAccounts",
	//   "httpMethod": "POST",
	//   "id": "cloudbilling.billingAccounts.create",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1/billingAccounts",
	//   "request": {
	//     "$ref": "BillingAccount"
	//   },
	//   "response": {
	//     "$ref": "BillingAccount"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbilling.billingAccounts.get":

type BillingAccountsGetCall struct {
	s            *APIService
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets information about a billing account. The current
// authenticated user
// must be a [viewer of the
// billing
// account](https://cloud.google.com/billing/docs/how-to/billing-
// access).
func (r *BillingAccountsService) Get(name string) *BillingAccountsGetCall {
	c := &BillingAccountsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BillingAccountsGetCall) Fields(s ...googleapi.Field) *BillingAccountsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BillingAccountsGetCall) IfNoneMatch(entityTag string) *BillingAccountsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BillingAccountsGetCall) Context(ctx context.Context) *BillingAccountsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BillingAccountsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BillingAccountsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbilling.billingAccounts.get" call.
// Exactly one of *BillingAccount or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *BillingAccount.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BillingAccountsGetCall) Do(opts ...googleapi.CallOption) (*BillingAccount, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &BillingAccount{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets information about a billing account. The current authenticated user\nmust be a [viewer of the billing\naccount](https://cloud.google.com/billing/docs/how-to/billing-access).",
	//   "flatPath": "v1/billingAccounts/{billingAccountsId}",
	//   "httpMethod": "GET",
	//   "id": "cloudbilling.billingAccounts.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The resource name of the billing account to retrieve. For example,\n`billingAccounts/012345-567890-ABCDEF`.",
	//       "location": "path",
	//       "pattern": "^billingAccounts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "BillingAccount"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbilling.billingAccounts.getIamPolicy":

type BillingAccountsGetIamPolicyCall struct {
	s            *APIService
	resource     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetIamPolicy: Gets the access control policy for a billing
// account.
// The caller must have the `billing.accounts.getIamPolicy` permission
// on the
// account, which is often given to billing
// account
// [viewers](https://cloud.google.com/billing/docs/how-to/billing
// -access).
func (r *BillingAccountsService) GetIamPolicy(resource string) *BillingAccountsGetIamPolicyCall {
	c := &BillingAccountsGetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BillingAccountsGetIamPolicyCall) Fields(s ...googleapi.Field) *BillingAccountsGetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BillingAccountsGetIamPolicyCall) IfNoneMatch(entityTag string) *BillingAccountsGetIamPolicyCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BillingAccountsGetIamPolicyCall) Context(ctx context.Context) *BillingAccountsGetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BillingAccountsGetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BillingAccountsGetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:getIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbilling.billingAccounts.getIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *BillingAccountsGetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the access control policy for a billing account.\nThe caller must have the `billing.accounts.getIamPolicy` permission on the\naccount, which is often given to billing account\n[viewers](https://cloud.google.com/billing/docs/how-to/billing-access).",
	//   "flatPath": "v1/billingAccounts/{billingAccountsId}:getIamPolicy",
	//   "httpMethod": "GET",
	//   "id": "cloudbilling.billingAccounts.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^billingAccounts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:getIamPolicy",
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbilling.billingAccounts.list":

type BillingAccountsListCall struct {
	s            *APIService
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the billing accounts that the current authenticated user
// has
// permission to
// [view](https://cloud.google.com/billing/docs/how-to/billing-access).
func (r *BillingAccountsService) List() *BillingAccountsListCall {
	c := &BillingAccountsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Filter sets the optional parameter "filter": Options for how to
// filter the returned billing accounts.
// Currently this only supports filtering
// for
// [subaccounts](https://cloud.google.com/billing/docs/concepts) under
// a
// single provided reseller billing account.
// (e.g.
// "master_billing_account=billingAccounts/012345-678901-ABCDEF").
// Boolea
// n algebra and other fields are not currently supported.
func (c *BillingAccountsListCall) Filter(filter string) *BillingAccountsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The maximum page size is 100; this is also the
// default.
func (c *BillingAccountsListCall) PageSize(pageSize int64) *BillingAccountsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results to return. This should be
// a
// `next_page_token` value returned from a previous
// `ListBillingAccounts`
// call. If unspecified, the first page of results is returned.
func (c *BillingAccountsListCall) PageToken(pageToken string) *BillingAccountsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BillingAccountsListCall) Fields(s ...googleapi.Field) *BillingAccountsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BillingAccountsListCall) IfNoneMatch(entityTag string) *BillingAccountsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BillingAccountsListCall) Context(ctx context.Context) *BillingAccountsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BillingAccountsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BillingAccountsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/billingAccounts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbilling.billingAccounts.list" call.
// Exactly one of *ListBillingAccountsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListBillingAccountsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BillingAccountsListCall) Do(opts ...googleapi.CallOption) (*ListBillingAccountsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListBillingAccountsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the billing accounts that the current authenticated user has\npermission to [view](https://cloud.google.com/billing/docs/how-to/billing-access).",
	//   "flatPath": "v1/billingAccounts",
	//   "httpMethod": "GET",
	//   "id": "cloudbilling.billingAccounts.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "filter": {
	//       "description": "Options for how to filter the returned billing accounts.\nCurrently this only supports filtering for\n[subaccounts](https://cloud.google.com/billing/docs/concepts) under a\nsingle provided reseller billing account.\n(e.g. \"master_billing_account=billingAccounts/012345-678901-ABCDEF\").\nBoolean algebra and other fields are not currently supported.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The maximum page size is 100; this is also the\ndefault.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results to return. This should be a\n`next_page_token` value returned from a previous `ListBillingAccounts`\ncall. If unspecified, the first page of results is returned.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/billingAccounts",
	//   "response": {
	//     "$ref": "ListBillingAccountsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BillingAccountsListCall) Pages(ctx context.Context, f func(*ListBillingAccountsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "cloudbilling.billingAccounts.patch":

type BillingAccountsPatchCall struct {
	s              *APIService
	name           string
	billingaccount *BillingAccount
	urlParams_     gensupport.URLParams
	ctx_           context.Context
	header_        http.Header
}

// Patch: Updates a billing account's fields.
// Currently the only field that can be edited is `display_name`.
// The current authenticated user must have the
// `billing.accounts.update`
// IAM permission, which is typically given to
// the
// [administrator](https://cloud.google.com/billing/docs/how-to/billi
// ng-access)
// of the billing account.
func (r *BillingAccountsService) Patch(name string, billingaccount *BillingAccount) *BillingAccountsPatchCall {
	c := &BillingAccountsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.billingaccount = billingaccount
	return c
}

// UpdateMask sets the optional parameter "updateMask": The update mask
// applied to the resource.
// Only "display_name" is currently supported.
func (c *BillingAccountsPatchCall) UpdateMask(updateMask string) *BillingAccountsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BillingAccountsPatchCall) Fields(s ...googleapi.Field) *BillingAccountsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BillingAccountsPatchCall) Context(ctx context.Context) *BillingAccountsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BillingAccountsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BillingAccountsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.billingaccount)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbilling.billingAccounts.patch" call.
// Exactly one of *BillingAccount or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *BillingAccount.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BillingAccountsPatchCall) Do(opts ...googleapi.CallOption) (*BillingAccount, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &BillingAccount{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a billing account's fields.\nCurrently the only field that can be edited is `display_name`.\nThe current authenticated user must have the `billing.accounts.update`\nIAM permission, which is typically given to the\n[administrator](https://cloud.google.com/billing/docs/how-to/billing-access)\nof the billing account.",
	//   "flatPath": "v1/billingAccounts/{billingAccountsId}",
	//   "httpMethod": "PATCH",
	//   "id": "cloudbilling.billingAccounts.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the billing account resource to be updated.",
	//       "location": "path",
	//       "pattern": "^billingAccounts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "The update mask applied to the resource.\nOnly \"display_name\" is currently supported.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "request": {
	//     "$ref": "BillingAccount"
	//   },
	//   "response": {
	//     "$ref": "BillingAccount"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbilling.billingAccounts.setIamPolicy":

type BillingAccountsSetIamPolicyCall struct {
	s                   *APIService
	resource            string
	setiampolicyrequest *SetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// SetIamPolicy: Sets the access control policy for a billing account.
// Replaces any existing
// policy.
// The caller must have the `billing.accounts.setIamPolicy` permission
// on the
// account, which is often given to billing
// account
// [administrators](https://cloud.google.com/billing/docs/how-to/
// billing-access).
func (r *BillingAccountsService) SetIamPolicy(resource string, setiampolicyrequest *SetIamPolicyRequest) *BillingAccountsSetIamPolicyCall {
	c := &BillingAccountsSetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.setiampolicyrequest = setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BillingAccountsSetIamPolicyCall) Fields(s ...googleapi.Field) *BillingAccountsSetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BillingAccountsSetIamPolicyCall) Context(ctx context.Context) *BillingAccountsSetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BillingAccountsSetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BillingAccountsSetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.setiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:setIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbilling.billingAccounts.setIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *BillingAccountsSetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the access control policy for a billing account. Replaces any existing\npolicy.\nThe caller must have the `billing.accounts.setIamPolicy` permission on the\naccount, which is often given to billing account\n[administrators](https://cloud.google.com/billing/docs/how-to/billing-access).",
	//   "flatPath": "v1/billingAccounts/{billingAccountsId}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "cloudbilling.billingAccounts.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being specified.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^billingAccounts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:setIamPolicy",
	//   "request": {
	//     "$ref": "SetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbilling.billingAccounts.testIamPermissions":

type BillingAccountsTestIamPermissionsCall struct {
	s                         *APIService
	resource                  string
	testiampermissionsrequest *TestIamPermissionsRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
	header_                   http.Header
}

// TestIamPermissions: Tests the access control policy for a billing
// account. This method takes
// the resource and a set of permissions as input and returns the subset
// of
// the input permissions that the caller is allowed for that resource.
func (r *BillingAccountsService) TestIamPermissions(resource string, testiampermissionsrequest *TestIamPermissionsRequest) *BillingAccountsTestIamPermissionsCall {
	c := &BillingAccountsTestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.testiampermissionsrequest = testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BillingAccountsTestIamPermissionsCall) Fields(s ...googleapi.Field) *BillingAccountsTestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BillingAccountsTestIamPermissionsCall) Context(ctx context.Context) *BillingAccountsTestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BillingAccountsTestIamPermissionsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BillingAccountsTestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testiampermissionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:testIamPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbilling.billingAccounts.testIamPermissions" call.
// Exactly one of *TestIamPermissionsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *TestIamPermissionsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BillingAccountsTestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*TestIamPermissionsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TestIamPermissionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Tests the access control policy for a billing account. This method takes\nthe resource and a set of permissions as input and returns the subset of\nthe input permissions that the caller is allowed for that resource.",
	//   "flatPath": "v1/billingAccounts/{billingAccountsId}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "cloudbilling.billingAccounts.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy detail is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^billingAccounts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:testIamPermissions",
	//   "request": {
	//     "$ref": "TestIamPermissionsRequest"
	//   },
	//   "response": {
	//     "$ref": "TestIamPermissionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbilling.billingAccounts.projects.list":

type BillingAccountsProjectsListCall struct {
	s            *APIService
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the projects associated with a billing account. The
// current
// authenticated user must have the `billing.resourceAssociations.list`
// IAM
// permission, which is often given to billing
// account
// [viewers](https://cloud.google.com/billing/docs/how-to/billing
// -access).
func (r *BillingAccountsProjectsService) List(name string) *BillingAccountsProjectsListCall {
	c := &BillingAccountsProjectsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The maximum page size is 100; this is also the
// default.
func (c *BillingAccountsProjectsListCall) PageSize(pageSize int64) *BillingAccountsProjectsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results to be returned. This should be
// a
// `next_page_token` value returned from a previous
// `ListProjectBillingInfo`
// call. If unspecified, the first page of results is returned.
func (c *BillingAccountsProjectsListCall) PageToken(pageToken string) *BillingAccountsProjectsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BillingAccountsProjectsListCall) Fields(s ...googleapi.Field) *BillingAccountsProjectsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BillingAccountsProjectsListCall) IfNoneMatch(entityTag string) *BillingAccountsProjectsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BillingAccountsProjectsListCall) Context(ctx context.Context) *BillingAccountsProjectsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BillingAccountsProjectsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BillingAccountsProjectsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}/projects")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbilling.billingAccounts.projects.list" call.
// Exactly one of *ListProjectBillingInfoResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListProjectBillingInfoResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BillingAccountsProjectsListCall) Do(opts ...googleapi.CallOption) (*ListProjectBillingInfoResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListProjectBillingInfoResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the projects associated with a billing account. The current\nauthenticated user must have the `billing.resourceAssociations.list` IAM\npermission, which is often given to billing account\n[viewers](https://cloud.google.com/billing/docs/how-to/billing-access).",
	//   "flatPath": "v1/billingAccounts/{billingAccountsId}/projects",
	//   "httpMethod": "GET",
	//   "id": "cloudbilling.billingAccounts.projects.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The resource name of the billing account associated with the projects that\nyou want to list. For example, `billingAccounts/012345-567890-ABCDEF`.",
	//       "location": "path",
	//       "pattern": "^billingAccounts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The maximum page size is 100; this is also the\ndefault.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results to be returned. This should be a\n`next_page_token` value returned from a previous `ListProjectBillingInfo`\ncall. If unspecified, the first page of results is returned.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}/projects",
	//   "response": {
	//     "$ref": "ListProjectBillingInfoResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BillingAccountsProjectsListCall) Pages(ctx context.Context, f func(*ListProjectBillingInfoResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "cloudbilling.projects.getBillingInfo":

type ProjectsGetBillingInfoCall struct {
	s            *APIService
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetBillingInfo: Gets the billing information for a project. The
// current authenticated user
// must have [permission to view
// the
// project](https://cloud.google.com/docs/permissions-overview#h.bgs0
// oxofvnoo
// ).
func (r *ProjectsService) GetBillingInfo(name string) *ProjectsGetBillingInfoCall {
	c := &ProjectsGetBillingInfoCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGetBillingInfoCall) Fields(s ...googleapi.Field) *ProjectsGetBillingInfoCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsGetBillingInfoCall) IfNoneMatch(entityTag string) *ProjectsGetBillingInfoCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsGetBillingInfoCall) Context(ctx context.Context) *ProjectsGetBillingInfoCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsGetBillingInfoCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsGetBillingInfoCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}/billingInfo")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbilling.projects.getBillingInfo" call.
// Exactly one of *ProjectBillingInfo or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ProjectBillingInfo.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsGetBillingInfoCall) Do(opts ...googleapi.CallOption) (*ProjectBillingInfo, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ProjectBillingInfo{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the billing information for a project. The current authenticated user\nmust have [permission to view the\nproject](https://cloud.google.com/docs/permissions-overview#h.bgs0oxofvnoo\n).",
	//   "flatPath": "v1/projects/{projectsId}/billingInfo",
	//   "httpMethod": "GET",
	//   "id": "cloudbilling.projects.getBillingInfo",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The resource name of the project for which billing information is\nretrieved. For example, `projects/tokyo-rain-123`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}/billingInfo",
	//   "response": {
	//     "$ref": "ProjectBillingInfo"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbilling.projects.updateBillingInfo":

type ProjectsUpdateBillingInfoCall struct {
	s                  *APIService
	name               string
	projectbillinginfo *ProjectBillingInfo
	urlParams_         gensupport.URLParams
	ctx_               context.Context
	header_            http.Header
}

// UpdateBillingInfo: Sets or updates the billing account associated
// with a project. You specify
// the new billing account by setting the `billing_account_name` in
// the
// `ProjectBillingInfo` resource to the resource name of a billing
// account.
// Associating a project with an open billing account enables billing on
// the
// project and allows charges for resource usage. If the project already
// had a
// billing account, this method changes the billing account used for
// resource
// usage charges.
//
// *Note:* Incurred charges that have not yet been reported in the
// transaction
// history of the GCP Console might be billed to the new
// billing
// account, even if the charge occurred before the new billing account
// was
// assigned to the project.
//
// The current authenticated user must have ownership privileges for
// both
// the
// [project](https://cloud.google.com/docs/permissions-overview#h.bgs
// 0oxofvnoo
// ) and the
// [billing
// account](https://cloud.google.com/billing/docs/how-to/billing
// -access).
//
// You can disable billing on the project by setting
// the
// `billing_account_name` field to empty. This action disassociates
// the
// current billing account from the project. Any billable activity of
// your
// in-use services will stop, and your application could stop
// functioning as
// expected. Any unbilled charges to date will be billed to the
// previously
// associated account. The current authenticated user must be either an
// owner
// of the project or an owner of the billing account for the
// project.
//
// Note that associating a project with a *closed* billing account will
// have
// much the same effect as disabling billing on the project: any
// paid
// resources used by the project will be shut down. Thus, unless you
// wish to
// disable billing, you should always call this method with the name of
// an
// *open* billing account.
func (r *ProjectsService) UpdateBillingInfo(name string, projectbillinginfo *ProjectBillingInfo) *ProjectsUpdateBillingInfoCall {
	c := &ProjectsUpdateBillingInfoCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.projectbillinginfo = projectbillinginfo
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsUpdateBillingInfoCall) Fields(s ...googleapi.Field) *ProjectsUpdateBillingInfoCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsUpdateBillingInfoCall) Context(ctx context.Context) *ProjectsUpdateBillingInfoCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsUpdateBillingInfoCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsUpdateBillingInfoCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.projectbillinginfo)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}/billingInfo")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbilling.projects.updateBillingInfo" call.
// Exactly one of *ProjectBillingInfo or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ProjectBillingInfo.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsUpdateBillingInfoCall) Do(opts ...googleapi.CallOption) (*ProjectBillingInfo, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ProjectBillingInfo{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets or updates the billing account associated with a project. You specify\nthe new billing account by setting the `billing_account_name` in the\n`ProjectBillingInfo` resource to the resource name of a billing account.\nAssociating a project with an open billing account enables billing on the\nproject and allows charges for resource usage. If the project already had a\nbilling account, this method changes the billing account used for resource\nusage charges.\n\n*Note:* Incurred charges that have not yet been reported in the transaction\nhistory of the GCP Console might be billed to the new billing\naccount, even if the charge occurred before the new billing account was\nassigned to the project.\n\nThe current authenticated user must have ownership privileges for both the\n[project](https://cloud.google.com/docs/permissions-overview#h.bgs0oxofvnoo\n) and the [billing\naccount](https://cloud.google.com/billing/docs/how-to/billing-access).\n\nYou can disable billing on the project by setting the\n`billing_account_name` field to empty. This action disassociates the\ncurrent billing account from the project. Any billable activity of your\nin-use services will stop, and your application could stop functioning as\nexpected. Any unbilled charges to date will be billed to the previously\nassociated account. The current authenticated user must be either an owner\nof the project or an owner of the billing account for the project.\n\nNote that associating a project with a *closed* billing account will have\nmuch the same effect as disabling billing on the project: any paid\nresources used by the project will be shut down. Thus, unless you wish to\ndisable billing, you should always call this method with the name of an\n*open* billing account.",
	//   "flatPath": "v1/projects/{projectsId}/billingInfo",
	//   "httpMethod": "PUT",
	//   "id": "cloudbilling.projects.updateBillingInfo",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The resource name of the project associated with the billing information\nthat you want to update. For example, `projects/tokyo-rain-123`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}/billingInfo",
	//   "request": {
	//     "$ref": "ProjectBillingInfo"
	//   },
	//   "response": {
	//     "$ref": "ProjectBillingInfo"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbilling.services.list":

type ServicesListCall struct {
	s            *APIService
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all public cloud services.
func (r *ServicesService) List() *ServicesListCall {
	c := &ServicesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// Defaults to 5000.
func (c *ServicesListCall) PageSize(pageSize int64) *ServicesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results to return. This should be
// a
// `next_page_token` value returned from a previous `ListServices`
// call. If unspecified, the first page of results is returned.
func (c *ServicesListCall) PageToken(pageToken string) *ServicesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ServicesListCall) Fields(s ...googleapi.Field) *ServicesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ServicesListCall) IfNoneMatch(entityTag string) *ServicesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ServicesListCall) Context(ctx context.Context) *ServicesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ServicesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ServicesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/services")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbilling.services.list" call.
// Exactly one of *ListServicesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListServicesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ServicesListCall) Do(opts ...googleapi.CallOption) (*ListServicesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListServicesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all public cloud services.",
	//   "flatPath": "v1/services",
	//   "httpMethod": "GET",
	//   "id": "cloudbilling.services.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Requested page size. Defaults to 5000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results to return. This should be a\n`next_page_token` value returned from a previous `ListServices`\ncall. If unspecified, the first page of results is returned.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/services",
	//   "response": {
	//     "$ref": "ListServicesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ServicesListCall) Pages(ctx context.Context, f func(*ListServicesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "cloudbilling.services.skus.list":

type ServicesSkusListCall struct {
	s            *APIService
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all publicly available SKUs for a given cloud service.
func (r *ServicesSkusService) List(parent string) *ServicesSkusListCall {
	c := &ServicesSkusListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// CurrencyCode sets the optional parameter "currencyCode": The ISO 4217
// currency code for the pricing info in the response proto.
// Will use the conversion rate as of start_time.
//  If not specified USD will be used.
func (c *ServicesSkusListCall) CurrencyCode(currencyCode string) *ServicesSkusListCall {
	c.urlParams_.Set("currencyCode", currencyCode)
	return c
}

// EndTime sets the optional parameter "endTime": Optional exclusive end
// time of the time range for which the pricing
// versions will be returned. Timestamps in the future are not
// allowed.
// The time range has to be within a single calendar month
// in
// America/Los_Angeles timezone. Time range as a whole is optional. If
// not
// specified, the latest pricing will be returned (up to 12 hours old
// at
// most).
func (c *ServicesSkusListCall) EndTime(endTime string) *ServicesSkusListCall {
	c.urlParams_.Set("endTime", endTime)
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// Defaults to 5000.
func (c *ServicesSkusListCall) PageSize(pageSize int64) *ServicesSkusListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results to return. This should be
// a
// `next_page_token` value returned from a previous `ListSkus`
// call. If unspecified, the first page of results is returned.
func (c *ServicesSkusListCall) PageToken(pageToken string) *ServicesSkusListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// StartTime sets the optional parameter "startTime": Optional inclusive
// start time of the time range for which the pricing
// versions will be returned. Timestamps in the future are not
// allowed.
// The time range has to be within a single calendar month
// in
// America/Los_Angeles timezone. Time range as a whole is optional. If
// not
// specified, the latest pricing will be returned (up to 12 hours old
// at
// most).
func (c *ServicesSkusListCall) StartTime(startTime string) *ServicesSkusListCall {
	c.urlParams_.Set("startTime", startTime)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ServicesSkusListCall) Fields(s ...googleapi.Field) *ServicesSkusListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ServicesSkusListCall) IfNoneMatch(entityTag string) *ServicesSkusListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ServicesSkusListCall) Context(ctx context.Context) *ServicesSkusListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ServicesSkusListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ServicesSkusListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/skus")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbilling.services.skus.list" call.
// Exactly one of *ListSkusResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListSkusResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ServicesSkusListCall) Do(opts ...googleapi.CallOption) (*ListSkusResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListSkusResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all publicly available SKUs for a given cloud service.",
	//   "flatPath": "v1/services/{servicesId}/skus",
	//   "httpMethod": "GET",
	//   "id": "cloudbilling.services.skus.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "currencyCode": {
	//       "description": "The ISO 4217 currency code for the pricing info in the response proto.\nWill use the conversion rate as of start_time.\nOptional. If not specified USD will be used.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "endTime": {
	//       "description": "Optional exclusive end time of the time range for which the pricing\nversions will be returned. Timestamps in the future are not allowed.\nThe time range has to be within a single calendar month in\nAmerica/Los_Angeles timezone. Time range as a whole is optional. If not\nspecified, the latest pricing will be returned (up to 12 hours old at\nmost).",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. Defaults to 5000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results to return. This should be a\n`next_page_token` value returned from a previous `ListSkus`\ncall. If unspecified, the first page of results is returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The name of the service.\nExample: \"services/DA34-426B-A397\"",
	//       "location": "path",
	//       "pattern": "^services/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startTime": {
	//       "description": "Optional inclusive start time of the time range for which the pricing\nversions will be returned. Timestamps in the future are not allowed.\nThe time range has to be within a single calendar month in\nAmerica/Los_Angeles timezone. Time range as a whole is optional. If not\nspecified, the latest pricing will be returned (up to 12 hours old at\nmost).",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/skus",
	//   "response": {
	//     "$ref": "ListSkusResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ServicesSkusListCall) Pages(ctx context.Context, f func(*ListSkusResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}
