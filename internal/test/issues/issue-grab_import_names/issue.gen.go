// Package grab_import_names provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package grab_import_names

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// GetFooParams defines parameters for GetFoo.
type GetFooParams struct {

	// base64. bytes. chi. context. echo. errors. fmt. gzip. http. io. ioutil. json. openapi3.
	Foo *string `json:"Foo,omitempty"`

	// openapi_types. path. runtime. strings. time.Duration time.Time url. xml. yaml.
	Bar *string `json:"Bar,omitempty"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetFoo request
	GetFoo(ctx context.Context, params *GetFooParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetFoo(ctx context.Context, params *GetFooParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetFooRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetFooRequest generates requests for GetFoo
func NewGetFooRequest(server string, params *GetFooParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/foo")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	if params.Foo != nil {
		var headerParam0 string

		headerParam0, err = runtime.StyleParam("simple", false, "Foo", *params.Foo)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Foo", headerParam0)
	}

	if params.Bar != nil {
		var headerParam1 string

		headerParam1, err = runtime.StyleParam("simple", false, "Bar", *params.Bar)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Bar", headerParam1)
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	req = req.WithContext(ctx)
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetFoo request
	GetFooWithResponse(ctx context.Context, params *GetFooParams) (*GetFooResponse, error)
}

type GetFooResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *string
}

// Status returns HTTPResponse.Status
func (r GetFooResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetFooResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetFooWithResponse request returning *GetFooResponse
func (c *ClientWithResponses) GetFooWithResponse(ctx context.Context, params *GetFooParams) (*GetFooResponse, error) {
	rsp, err := c.GetFoo(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseGetFooResponse(rsp)
}

// ParseGetFooResponse parses an HTTP response from a GetFooWithResponse call
func ParseGetFooResponse(rsp *http.Response) (*GetFooResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetFooResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest string
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /foo)
	GetFoo(ctx echo.Context, params GetFooParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetFoo converts echo context to params.
func (w *ServerInterfaceWrapper) GetFoo(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetFooParams

	headers := ctx.Request().Header
	// ------------- Optional header parameter "Foo" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Foo")]; found {
		var Foo string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for Foo, got %d", n))
		}

		err = runtime.BindStyledParameter("simple", false, "Foo", valueList[0], &Foo)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter Foo: %s", err))
		}

		params.Foo = &Foo
	}
	// ------------- Optional header parameter "Bar" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Bar")]; found {
		var Bar string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for Bar, got %d", n))
		}

		err = runtime.BindStyledParameter("simple", false, "Bar", valueList[0], &Bar)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter Bar: %s", err))
		}

		params.Bar = &Bar
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetFoo(ctx, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/foo", wrapper.GetFoo)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/3yRwWojMQyGX0XobJQhWfYwx2XZpffeSinORMm4jC1jaUrSMO9e7MkpNDUYIWF9/vXr",
	"ioPELImTKfbXxWFIR8H+ihZsYuyRiNDhBxcNkrDHjjrqcHEomZPPAXvcUUdbdJi9jZWCm6M0xomthgPr",
	"UEK2FbACJXPxtfJ0wB7/s/0TaYjiIxsXxf7lvnPvlX//IthfjJVgGAPBIMn4bAQ8jELApUhRgmM0gtNn",
	"yASjWSYIUu9sYSJ4V0kEN/27qiZU/Mj+wAUdJh/r5KsiHUaOvjlyybWsVkI64bK4e3034lt9qATVDoIy",
	"JwuRCdY+JWjp33kdf82eQ2SYy0RwjhPBxcfpoaw/vvwo69VhYc2SlNsytl1XQzMqtX34nKcwtO831Yta",
	"e8xb3DcLXNr5CgAA//+IR/EuPgIAAA==",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}