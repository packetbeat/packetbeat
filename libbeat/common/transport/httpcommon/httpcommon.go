package httpcommon

import (
	"net/http"
	"time"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/transport"
	"github.com/elastic/beats/v7/libbeat/common/transport/tlscommon"
	"go.elastic.co/apm/module/apmhttp"
)

// HTTPTransportSettings provides common HTTP settings for HTTP clients.
type HTTPTransportSettings struct {
	// TLS provides ssl/tls setup settings
	TLS *tlscommon.TLSConfig `config:"ssl" yaml:"ssl,omitempty"`

	// Timeout configures the `(http.Transport).Timeout`.
	Timeout time.Duration `config:"timeout" yaml:"timeout,omitempty"`

	Proxy HTTPClientProxySettings `config:",inline" yaml:",inline"`

	// TODO: Add more settings:
	//  - DisableKeepAlive
	//  - MaxIdleConns
	//  - IdleConnTimeout
	//  - ResponseHeaderTimeout
	//  - ConnectionTimeout (currently 'Timeout' is used for both)
}

// WithKeepaliveSettings options can be used to modify the Keepalive
type WithKeepaliveSettings struct {
	Disable             bool
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	IdleConnTimeout     time.Duration
}

var _ httpTransportOption = WithKeepaliveSettings{}

type (
	// TransportOption are applied to the http.RoundTripper to be build
	// from HTTPTransportSettings.
	TransportOption interface{ sealTransportOption() }

	dialerOption interface {
		TransportOption
		baseDialer() transport.Dialer
	}
	dialerModOption interface {
		TransportOption
		applyDialer(*HTTPTransportSettings, transport.Dialer) transport.Dialer
	}
	httpTransportOption interface {
		TransportOption
		applyTransport(*HTTPTransportSettings, *http.Transport)
	}
	roundTripperOption interface {
		TransportOption
		applyRoundTripper(*HTTPTransportSettings, http.RoundTripper) http.RoundTripper
	}
)

type baseDialerFunc func() transport.Dialer

var _ dialerOption = baseDialerFunc(nil)

func (baseDialerFunc) sealTransportOption() {}
func (fn baseDialerFunc) baseDialer() transport.Dialer {
	return fn()
}

type dialerOptFunc func(transport.Dialer) transport.Dialer

var _ dialerModOption = dialerOptFunc(nil)

func (dialerOptFunc) sealTransportOption() {}
func (fn dialerOptFunc) applyDialer(_ *HTTPTransportSettings, d transport.Dialer) transport.Dialer {
	return fn(d)

}

type transportOptFunc func(*HTTPTransportSettings, *http.Transport)

var _ httpTransportOption = transportOptFunc(nil)

func (transportOptFunc) sealTransportOption() {}
func (fn transportOptFunc) applyTransport(s *HTTPTransportSettings, t *http.Transport) {
	fn(s, t)
}

type rtOptFunc func(http.RoundTripper) http.RoundTripper

var _ roundTripperOption = rtOptFunc(nil)

func (rtOptFunc) sealTransportOption() {}
func (fn rtOptFunc) applyRoundTripper(_ *HTTPTransportSettings, rt http.RoundTripper) http.RoundTripper {
	return fn(rt)
}

// DefaultHTTPTransportSettings returns the default HTTP transport setting.
func DefaultHTTPTransportSettings() HTTPTransportSettings {
	return HTTPTransportSettings{
		Proxy: DefaultHTTPClientProxySettings(),
	}
}

// Unpack reads a config object into the settings.
func (settings *HTTPTransportSettings) Unpack(cfg *common.Config) error {
	tmp := struct {
		TLS     *tlscommon.Config `config:"ssl"`
		Timeout time.Duration     `config:"timeout"`
	}{Timeout: settings.Timeout}

	if err := cfg.Unpack(&tmp); err != nil {
		return err
	}

	var proxy HTTPClientProxySettings
	if err := cfg.Unpack(&proxy); err != nil {
		return err
	}

	tls, err := tlscommon.LoadTLSConfig(tmp.TLS)
	if err != nil {
		return err
	}

	*settings = HTTPTransportSettings{
		TLS:     tls,
		Timeout: tmp.Timeout,
		Proxy:   proxy,
	}
	return nil
}

// RoundTripper creates a http.RoundTripper for use with http.Client.
//
// The dialers will registers with stats if given. Stats is used to collect metrics for io errors,
// bytes in, and bytes out.
func (settings *HTTPTransportSettings) RoundTripper(opts ...TransportOption) http.RoundTripper {
	var dialer transport.Dialer

	for _, opt := range opts {
		if dialOpt, ok := opt.(dialerOption); ok {
			dialer = dialOpt.baseDialer()
		}
	}

	if dialer == nil {
		dialer = transport.NetDialer(settings.Timeout)
	}
	tlsDialer := transport.TLSDialer(dialer, settings.TLS, settings.Timeout)

	for _, opt := range opts {
		if dialOpt, ok := opt.(dialerModOption); ok {
			dialer = dialOpt.applyDialer(settings, dialer)
			tlsDialer = dialOpt.applyDialer(settings, tlsDialer)
		}
	}

	t := http.DefaultTransport.(*http.Transport).Clone()
	t.DialContext = nil
	t.DialTLSContext = nil
	t.Dial = dialer.Dial
	t.DialTLS = tlsDialer.Dial
	t.TLSClientConfig = settings.TLS.ToConfig()
	t.ForceAttemptHTTP2 = false
	t.Proxy = settings.Proxy.ProxyFunc()

	for _, opt := range opts {
		if transportOpt, ok := opt.(httpTransportOption); ok {
			transportOpt.applyTransport(settings, t)
		}
	}

	rt := http.RoundTripper(t)
	for _, opt := range opts {
		if rtOpt, ok := opt.(roundTripperOption); ok {
			rt = rtOpt.applyRoundTripper(settings, rt)
		}
	}

	return rt
}

// Client creates a new http.Client with configured Transport. The transport is
// instrumented using apmhttp.WrapRoundTripper.
func (settings HTTPTransportSettings) Client(opts ...TransportOption) *http.Client {
	return &http.Client{
		Transport: settings.RoundTripper(opts...),
		Timeout:   settings.Timeout,
	}
}

func (opts WithKeepaliveSettings) sealTransportOption() {}
func (opts WithKeepaliveSettings) applyTransport(_ *HTTPTransportSettings, t *http.Transport) {
	t.DisableKeepAlives = opts.Disable
	if opts.IdleConnTimeout != 0 {
		t.IdleConnTimeout = opts.IdleConnTimeout
	}
	if opts.MaxIdleConns != 0 {
		t.MaxIdleConns = opts.MaxIdleConns
	}
	if opts.MaxIdleConnsPerHost != 0 {
		t.MaxIdleConnsPerHost = opts.MaxIdleConnsPerHost
	}
}

// WithBaseDialer configures the dialer used for TCP and TLS connections.
func WithBaseDialer(d transport.Dialer) TransportOption {
	return baseDialerFunc(func() transport.Dialer {
		return d
	})
}

// WithIOStats instruments the RoundTripper dialers with the given statser, such
// that bytes in, bytes out, and errors can be monitored.
func WithIOStats(stats transport.IOStatser) TransportOption {
	return dialerOptFunc(func(d transport.Dialer) transport.Dialer {
		if stats == nil {
			return d
		}
		return transport.StatsDialer(d, stats)
	})
}

// WithTransportFunc register a custom function that is used to apply
// custom changes to the net.Transport, when the Client is build.
func WithTransportFunc(fn func(*http.Transport)) TransportOption {
	return transportOptFunc(func(_ *HTTPTransportSettings, t *http.Transport) {
		fn(t)
	})
}

// WithForceAttemptHTTP2 sets the `http.Tansport.ForceAttemptHTTP2` field.
func WithForceAttemptHTTP2(b bool) TransportOption {
	return transportOptFunc(func(settings *HTTPTransportSettings, t *http.Transport) {
		t.ForceAttemptHTTP2 = b
	})
}

// WithNOProxy disables the configured proxy. Proxy environment variables
// like HTTP_PROXY and HTTPS_PROXY will have no affect.
func WithNOProxy() TransportOption {
	return transportOptFunc(func(s *HTTPTransportSettings, t *http.Transport) {
		t.Proxy = nil
	})
}

// WithoutProxyEnvironmentVariables disables support for the HTTP_PROXY, HTTPS_PROXY and
// NO_PROXY envionrment variables. Explicitely configured proxy URLs will still applied.
func WithoutProxyEnvironmentVariables() TransportOption {
	return transportOptFunc(func(settings *HTTPTransportSettings, t *http.Transport) {
		if settings.Proxy.Disable || settings.Proxy.URL == nil {
			t.Proxy = nil
		}
	})
}

// WithModRoundtripper allows customization of the roundtipper.
func WithModRoundtripper(w func(http.RoundTripper) http.RoundTripper) TransportOption {
	return rtOptFunc(w)
}

var withAPMHTTPRountTripper = WithModRoundtripper(func(rt http.RoundTripper) http.RoundTripper {
	return apmhttp.WrapRoundTripper(rt)
})

// WithAPMHTTPInstrumentation insruments the HTTP client via apmhttp.WrapRoundTripper.
// Custom APM round tripper wrappers can be configured via WithModRoundtripper.
func WithAPMHTTPInstrumentation() TransportOption {
	return withAPMHTTPRountTripper
}
