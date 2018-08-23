// Package cloudtrace provides access to the Stackdriver Trace API.
//
// See https://cloud.google.com/trace
//
// Usage example:
//
//   import "github.com/sunnogo/google-api-go-client/cloudtrace/v2"
//   ...
//   cloudtraceService, err := cloudtrace.New(oauthHttpClient)
package cloudtrace // import "github.com/sunnogo/google-api-go-client/cloudtrace/v2"

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

const apiId = "cloudtrace:v2"
const apiName = "cloudtrace"
const apiVersion = "v2"
const basePath = "https://cloudtrace.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// Write Trace data for a project or application
	TraceAppendScope = "https://www.googleapis.com/auth/trace.append"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Traces = NewProjectsTracesService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Traces *ProjectsTracesService
}

func NewProjectsTracesService(s *Service) *ProjectsTracesService {
	rs := &ProjectsTracesService{s: s}
	rs.Spans = NewProjectsTracesSpansService(s)
	return rs
}

type ProjectsTracesService struct {
	s *Service

	Spans *ProjectsTracesSpansService
}

func NewProjectsTracesSpansService(s *Service) *ProjectsTracesSpansService {
	rs := &ProjectsTracesSpansService{s: s}
	return rs
}

type ProjectsTracesSpansService struct {
	s *Service
}

// Annotation: Text annotation with a set of attributes.
type Annotation struct {
	// Attributes: A set of attributes on the annotation. You can have up to
	// 4 attributes
	// per Annotation.
	Attributes *Attributes `json:"attributes,omitempty"`

	// Description: A user-supplied message describing the event. The
	// maximum length for
	// the description is 256 bytes.
	Description *TruncatableString `json:"description,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Attributes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Attributes") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Annotation) MarshalJSON() ([]byte, error) {
	type NoMethod Annotation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AttributeValue: The allowed types for [VALUE] in a `[KEY]:[VALUE]`
// attribute.
type AttributeValue struct {
	// BoolValue: A Boolean value represented by `true` or `false`.
	BoolValue bool `json:"boolValue,omitempty"`

	// IntValue: A 64-bit signed integer.
	IntValue int64 `json:"intValue,omitempty,string"`

	// StringValue: A string up to 256 bytes long.
	StringValue *TruncatableString `json:"stringValue,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BoolValue") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BoolValue") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AttributeValue) MarshalJSON() ([]byte, error) {
	type NoMethod AttributeValue
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Attributes: A set of attributes, each in the format `[KEY]:[VALUE]`.
type Attributes struct {
	// AttributeMap: The set of attributes. Each attribute's key can be up
	// to 128 bytes
	// long. The value can be a string up to 256 bytes, an integer, or
	// the
	// Boolean values `true` and `false`. For example:
	//
	//     "/instance_id": "my-instance"
	//     "/http/user_agent": ""
	//     "/http/request_bytes": 300
	//     "abc.com/myattribute": true
	AttributeMap map[string]AttributeValue `json:"attributeMap,omitempty"`

	// DroppedAttributesCount: The number of attributes that were discarded.
	// Attributes can be discarded
	// because their keys are too long or because there are too many
	// attributes.
	// If this value is 0 then all attributes are valid.
	DroppedAttributesCount int64 `json:"droppedAttributesCount,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AttributeMap") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AttributeMap") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Attributes) MarshalJSON() ([]byte, error) {
	type NoMethod Attributes
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BatchWriteSpansRequest: The request message for the `BatchWriteSpans`
// method.
type BatchWriteSpansRequest struct {
	// Spans: A list of new spans. The span names must not match
	// existing
	// spans, or the results are undefined.
	Spans []*Span `json:"spans,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Spans") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Spans") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BatchWriteSpansRequest) MarshalJSON() ([]byte, error) {
	type NoMethod BatchWriteSpansRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// Link: A pointer from the current span to another span in the same
// trace or in a
// different trace. For example, this can be used in batching
// operations,
// where a single batch handler processes multiple requests from
// different
// traces or when the handler receives a request from a different
// project.
type Link struct {
	// Attributes: A set of attributes on the link. You have have up to  32
	// attributes per
	// link.
	Attributes *Attributes `json:"attributes,omitempty"`

	// SpanId: The [SPAN_ID] for a span within a trace.
	SpanId string `json:"spanId,omitempty"`

	// TraceId: The [TRACE_ID] for a trace within a project.
	TraceId string `json:"traceId,omitempty"`

	// Type: The relationship of the current span relative to the linked
	// span.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - The relationship of the two spans is unknown.
	//   "CHILD_LINKED_SPAN" - The linked span is a child of the current
	// span.
	//   "PARENT_LINKED_SPAN" - The linked span is a parent of the current
	// span.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Attributes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Attributes") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Link) MarshalJSON() ([]byte, error) {
	type NoMethod Link
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Links: A collection of links, which are references from this span to
// a span
// in the same or different trace.
type Links struct {
	// DroppedLinksCount: The number of dropped links after the maximum size
	// was enforced. If
	// this value is 0, then no links were dropped.
	DroppedLinksCount int64 `json:"droppedLinksCount,omitempty"`

	// Link: A collection of links.
	Link []*Link `json:"link,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DroppedLinksCount")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DroppedLinksCount") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Links) MarshalJSON() ([]byte, error) {
	type NoMethod Links
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// MessageEvent: An event describing a message sent/received between
// Spans.
type MessageEvent struct {
	// CompressedSizeBytes: The number of compressed bytes sent or received.
	// If missing assumed to
	// be the same size as uncompressed.
	CompressedSizeBytes int64 `json:"compressedSizeBytes,omitempty,string"`

	// Id: An identifier for the MessageEvent's message that can be used to
	// match
	// SENT and RECEIVED MessageEvents. It is recommended to be unique
	// within
	// a Span.
	Id int64 `json:"id,omitempty,string"`

	// Type: Type of MessageEvent. Indicates whether the message was sent
	// or
	// received.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - Unknown event type.
	//   "SENT" - Indicates a sent message.
	//   "RECEIVED" - Indicates a received message.
	Type string `json:"type,omitempty"`

	// UncompressedSizeBytes: The number of uncompressed bytes sent or
	// received.
	UncompressedSizeBytes int64 `json:"uncompressedSizeBytes,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "CompressedSizeBytes")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CompressedSizeBytes") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *MessageEvent) MarshalJSON() ([]byte, error) {
	type NoMethod MessageEvent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Module: Binary module.
type Module struct {
	// BuildId: A unique identifier for the module, usually a hash of
	// its
	// contents (up to 128 bytes).
	BuildId *TruncatableString `json:"buildId,omitempty"`

	// Module: For example: main binary, kernel modules, and dynamic
	// libraries
	// such as libc.so, sharedlib.so (up to 256 bytes).
	Module *TruncatableString `json:"module,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BuildId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BuildId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Module) MarshalJSON() ([]byte, error) {
	type NoMethod Module
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Span: A span represents a single operation within a trace. Spans can
// be
// nested to form a trace tree. Often, a trace contains a root span
// that describes the end-to-end latency, and one or more subspans
// for
// its sub-operations. A trace can also contain multiple root spans,
// or none at all. Spans do not need to be contiguous&mdash;there may
// be
// gaps or overlaps between spans in a trace.
type Span struct {
	// Attributes: A set of attributes on the span. You can have up to 32
	// attributes per
	// span.
	Attributes *Attributes `json:"attributes,omitempty"`

	// ChildSpanCount: An optional number of child spans that were generated
	// while this span
	// was active. If set, allows implementation to detect missing child
	// spans.
	ChildSpanCount int64 `json:"childSpanCount,omitempty"`

	// DisplayName: A description of the span's operation (up to 128
	// bytes).
	// Stackdriver Trace displays the description in the
	// Google Cloud Platform Console.
	// For example, the display name can be a qualified method name or a
	// file name
	// and a line number where the operation is called. A best practice is
	// to use
	// the same display name within an application and at the same call
	// point.
	// This makes it easier to correlate spans in different traces.
	DisplayName *TruncatableString `json:"displayName,omitempty"`

	// EndTime: The end time of the span. On the client side, this is the
	// time kept by
	// the local machine where the span execution ends. On the server side,
	// this
	// is the time when the server application handler stops running.
	EndTime string `json:"endTime,omitempty"`

	// Links: Links associated with the span. You can have up to 128 links
	// per Span.
	Links *Links `json:"links,omitempty"`

	// Name: The resource name of the span in the following format:
	//
	//     projects/[PROJECT_ID]/traces/[TRACE_ID]/spans/SPAN_ID is a unique
	// identifier for a trace within a project;
	// it is a 32-character hexadecimal encoding of a 16-byte
	// array.
	//
	// [SPAN_ID] is a unique identifier for a span within a trace; it
	// is a 16-character hexadecimal encoding of an 8-byte array.
	Name string `json:"name,omitempty"`

	// ParentSpanId: The [SPAN_ID] of this span's parent span. If this is a
	// root span,
	// then this field must be empty.
	ParentSpanId string `json:"parentSpanId,omitempty"`

	// SameProcessAsParentSpan: (Optional) Set this parameter to indicate
	// whether this span is in
	// the same process as its parent. If you do not set this
	// parameter,
	// Stackdriver Trace is unable to take advantage of this
	// helpful
	// information.
	SameProcessAsParentSpan bool `json:"sameProcessAsParentSpan,omitempty"`

	// SpanId: The [SPAN_ID] portion of the span's resource name.
	SpanId string `json:"spanId,omitempty"`

	// StackTrace: Stack trace captured at the start of the span.
	StackTrace *StackTrace `json:"stackTrace,omitempty"`

	// StartTime: The start time of the span. On the client side, this is
	// the time kept by
	// the local machine where the span execution starts. On the server
	// side, this
	// is the time when the server's application handler starts running.
	StartTime string `json:"startTime,omitempty"`

	// Status: An optional final status for this span.
	Status *Status `json:"status,omitempty"`

	// TimeEvents: A set of time events. You can have up to 32 annotations
	// and 128 message
	// events per span.
	TimeEvents *TimeEvents `json:"timeEvents,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Attributes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Attributes") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Span) MarshalJSON() ([]byte, error) {
	type NoMethod Span
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StackFrame: Represents a single stack frame in a stack trace.
type StackFrame struct {
	// ColumnNumber: The column number where the function call appears, if
	// available.
	// This is important in JavaScript because of its anonymous functions.
	ColumnNumber int64 `json:"columnNumber,omitempty,string"`

	// FileName: The name of the source file where the function call appears
	// (up to 256
	// bytes).
	FileName *TruncatableString `json:"fileName,omitempty"`

	// FunctionName: The fully-qualified name that uniquely identifies the
	// function or
	// method that is active in this frame (up to 1024 bytes).
	FunctionName *TruncatableString `json:"functionName,omitempty"`

	// LineNumber: The line number in `file_name` where the function call
	// appears.
	LineNumber int64 `json:"lineNumber,omitempty,string"`

	// LoadModule: The binary module from where the code was loaded.
	LoadModule *Module `json:"loadModule,omitempty"`

	// OriginalFunctionName: An un-mangled function name, if `function_name`
	// is
	// [mangled](http://www.avabodh.com/cxxin/namemangling.html). The name
	// can
	// be fully-qualified (up to 1024 bytes).
	OriginalFunctionName *TruncatableString `json:"originalFunctionName,omitempty"`

	// SourceVersion: The version of the deployed source code (up to 128
	// bytes).
	SourceVersion *TruncatableString `json:"sourceVersion,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ColumnNumber") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ColumnNumber") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StackFrame) MarshalJSON() ([]byte, error) {
	type NoMethod StackFrame
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StackFrames: A collection of stack frames, which can be truncated.
type StackFrames struct {
	// DroppedFramesCount: The number of stack frames that were dropped
	// because there
	// were too many stack frames.
	// If this value is 0, then no stack frames were dropped.
	DroppedFramesCount int64 `json:"droppedFramesCount,omitempty"`

	// Frame: Stack frames in this call stack.
	Frame []*StackFrame `json:"frame,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DroppedFramesCount")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DroppedFramesCount") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *StackFrames) MarshalJSON() ([]byte, error) {
	type NoMethod StackFrames
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StackTrace: A call stack appearing in a trace.
type StackTrace struct {
	// StackFrames: Stack frames in this stack trace. A maximum of 128
	// frames are allowed.
	StackFrames *StackFrames `json:"stackFrames,omitempty"`

	// StackTraceHashId: The hash ID is used to conserve network bandwidth
	// for duplicate
	// stack traces within a single trace.
	//
	// Often multiple spans will have identical stack traces.
	// The first occurrence of a stack trace should contain both
	// the
	// `stackFrame` content and a value in `stackTraceHashId`.
	//
	// Subsequent spans within the same request can refer
	// to that stack trace by only setting `stackTraceHashId`.
	StackTraceHashId int64 `json:"stackTraceHashId,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "StackFrames") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "StackFrames") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StackTrace) MarshalJSON() ([]byte, error) {
	type NoMethod StackTrace
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Status: The `Status` type defines a logical error model that is
// suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` that can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There is a
	// common set of
	// message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TimeEvent: A time-stamped annotation or message event in the Span.
type TimeEvent struct {
	// Annotation: Text annotation with a set of attributes.
	Annotation *Annotation `json:"annotation,omitempty"`

	// MessageEvent: An event describing a message sent/received between
	// Spans.
	MessageEvent *MessageEvent `json:"messageEvent,omitempty"`

	// Time: The timestamp indicating the time the event occurred.
	Time string `json:"time,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Annotation") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Annotation") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TimeEvent) MarshalJSON() ([]byte, error) {
	type NoMethod TimeEvent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TimeEvents: A collection of `TimeEvent`s. A `TimeEvent` is a
// time-stamped annotation
// on the span, consisting of either user-supplied key:value pairs,
// or
// details of a message sent/received between Spans.
type TimeEvents struct {
	// DroppedAnnotationsCount: The number of dropped annotations in all the
	// included time events.
	// If the value is 0, then no annotations were dropped.
	DroppedAnnotationsCount int64 `json:"droppedAnnotationsCount,omitempty"`

	// DroppedMessageEventsCount: The number of dropped message events in
	// all the included time events.
	// If the value is 0, then no message events were dropped.
	DroppedMessageEventsCount int64 `json:"droppedMessageEventsCount,omitempty"`

	// TimeEvent: A collection of `TimeEvent`s.
	TimeEvent []*TimeEvent `json:"timeEvent,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "DroppedAnnotationsCount") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DroppedAnnotationsCount")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *TimeEvents) MarshalJSON() ([]byte, error) {
	type NoMethod TimeEvents
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TruncatableString: Represents a string that might be shortened to a
// specified length.
type TruncatableString struct {
	// TruncatedByteCount: The number of bytes removed from the original
	// string. If this
	// value is 0, then the string was not shortened.
	TruncatedByteCount int64 `json:"truncatedByteCount,omitempty"`

	// Value: The shortened string. For example, if the original string is
	// 500
	// bytes long and the limit of the string is 128 bytes, then
	// `value` contains the first 128 bytes of the 500-byte
	// string.
	//
	// Truncation always happens on a UTF8 character boundary. If there
	// are multi-byte characters in the string, then the length of
	// the
	// shortened string might be less than the size limit.
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "TruncatedByteCount")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "TruncatedByteCount") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *TruncatableString) MarshalJSON() ([]byte, error) {
	type NoMethod TruncatableString
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "cloudtrace.projects.traces.batchWrite":

type ProjectsTracesBatchWriteCall struct {
	s                      *Service
	name                   string
	batchwritespansrequest *BatchWriteSpansRequest
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// BatchWrite: Sends new spans to new or existing traces. You cannot
// update
// existing spans.
func (r *ProjectsTracesService) BatchWrite(name string, batchwritespansrequest *BatchWriteSpansRequest) *ProjectsTracesBatchWriteCall {
	c := &ProjectsTracesBatchWriteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.batchwritespansrequest = batchwritespansrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTracesBatchWriteCall) Fields(s ...googleapi.Field) *ProjectsTracesBatchWriteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTracesBatchWriteCall) Context(ctx context.Context) *ProjectsTracesBatchWriteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTracesBatchWriteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTracesBatchWriteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.batchwritespansrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}/traces:batchWrite")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtrace.projects.traces.batchWrite" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsTracesBatchWriteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	ret := &Empty{
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
	//   "description": "Sends new spans to new or existing traces. You cannot update\nexisting spans.",
	//   "flatPath": "v2/projects/{projectsId}/traces:batchWrite",
	//   "httpMethod": "POST",
	//   "id": "cloudtrace.projects.traces.batchWrite",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the project where the spans belong. The format is\n`projects/[PROJECT_ID]`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}/traces:batchWrite",
	//   "request": {
	//     "$ref": "BatchWriteSpansRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/trace.append"
	//   ]
	// }

}

// method id "cloudtrace.projects.traces.spans.createSpan":

type ProjectsTracesSpansCreateSpanCall struct {
	s          *Service
	nameid     string
	span       *Span
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// CreateSpan: Creates a new span.
func (r *ProjectsTracesSpansService) CreateSpan(nameid string, span *Span) *ProjectsTracesSpansCreateSpanCall {
	c := &ProjectsTracesSpansCreateSpanCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.span = span
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTracesSpansCreateSpanCall) Fields(s ...googleapi.Field) *ProjectsTracesSpansCreateSpanCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTracesSpansCreateSpanCall) Context(ctx context.Context) *ProjectsTracesSpansCreateSpanCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTracesSpansCreateSpanCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTracesSpansCreateSpanCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.span)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtrace.projects.traces.spans.createSpan" call.
// Exactly one of *Span or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Span.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsTracesSpansCreateSpanCall) Do(opts ...googleapi.CallOption) (*Span, error) {
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
	ret := &Span{
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
	//   "description": "Creates a new span.",
	//   "flatPath": "v2/projects/{projectsId}/traces/{tracesId}/spans/{spansId}",
	//   "httpMethod": "POST",
	//   "id": "cloudtrace.projects.traces.spans.createSpan",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The resource name of the span in the following format:\n\n    projects/[PROJECT_ID]/traces/[TRACE_ID]/spans/SPAN_ID is a unique identifier for a trace within a project;\nit is a 32-character hexadecimal encoding of a 16-byte array.\n\n[SPAN_ID] is a unique identifier for a span within a trace; it\nis a 16-character hexadecimal encoding of an 8-byte array.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/traces/[^/]+/spans/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}",
	//   "request": {
	//     "$ref": "Span"
	//   },
	//   "response": {
	//     "$ref": "Span"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/trace.append"
	//   ]
	// }

}
