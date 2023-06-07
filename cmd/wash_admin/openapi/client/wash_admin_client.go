// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"wash_admin/openapi/client/standard"
	"wash_admin/openapi/client/users"
	"wash_admin/openapi/client/wash_servers"
)

// Default wash admin HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new wash admin HTTP client.
func NewHTTPClient(formats strfmt.Registry) *WashAdmin {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new wash admin HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *WashAdmin {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new wash admin client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *WashAdmin {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(WashAdmin)
	cli.Transport = transport
	cli.Standard = standard.New(transport, formats)
	cli.Users = users.New(transport, formats)
	cli.WashServers = wash_servers.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// WashAdmin is a client for wash admin
type WashAdmin struct {
	Standard standard.ClientService

	Users users.ClientService

	WashServers wash_servers.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *WashAdmin) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Standard.SetTransport(transport)
	c.Users.SetTransport(transport)
	c.WashServers.SetTransport(transport)
}
