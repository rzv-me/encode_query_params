package encode_query_params

import (
//    "fmt"
//    "io"
    "net/http"
//    "os"
    "strings"
    "net/url"

    "github.com/caddyserver/caddy/v2"
    "github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
    "github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
    "github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
    caddy.RegisterModule(Middleware{})
    httpcaddyfile.RegisterHandlerDirective("encode_query_params", parseCaddyfile)
}

// Middleware implements an HTTP handler that writes the
// visitor's IP address to a file or stream.
type Middleware struct {}

// CaddyModule returns the Caddy module information.
func (Middleware) CaddyModule() caddy.ModuleInfo {
    return caddy.ModuleInfo{
	ID:  "http.handlers.encode_query_params",
	New: func() caddy.Module { return new(Middleware) },
    }
}

// Provision implements caddy.Provisioner.
func (m *Middleware) Provision(ctx caddy.Context) error {
    return nil
}

// Validate implements caddy.Validator.
func (m *Middleware) Validate() error {
    return nil
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
/*
func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
    // Parse the existing query parameters
    query := r.URL.Query()
//    fmt.Println(query)
    for key, values := range query {
        for index := range values {
            // Replace the specific characters in each value
            values[index] = strings.ReplaceAll(values[index], "[", "%5B")
            values[index] = strings.ReplaceAll(values[index], "]", "%5D")
            values[index] = strings.ReplaceAll(values[index], "|", "%7C")
        }
        query[key] = values
    }
    // Encode the query parameters back into the URL
    r.URL.RawQuery = query.Encode()
    return next.ServeHTTP(w, r)
}
*/
func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
    // Parse the existing query parameters
    query := r.URL.Query()
    encodedQuery := url.Values{}

    for key, values := range query {
        // Encode specific characters in the key
	encodedKey := key
    // fix spaces in key
    // fix tilda ~ in key
    // fix dots . in key
//        encodedKey := strings.ReplaceAll(encodedKey, "[", "%5B")
//        encodedKey = strings.ReplaceAll(encodedKey, "]", "%5D")


        for index := range values {
            // Replace the specific character '|' in each value
            values[index] = strings.ReplaceAll(values[index], "|", "%7B")
            values[index] = strings.ReplaceAll(values[index], ";", "%3B")

        }

        encodedKey = strings.ReplaceAll(encodedKey, "|", "%7B")
        encodedKey = strings.ReplaceAll(encodedKey, ";", "%3B")
        encodedKey = strings.ReplaceAll(encodedKey, "~", "%7E")
        // Add the modified key and values to the new query object
        encodedQuery[encodedKey] = values
    }

    // Encode the modified query parameters back into the URL
    r.URL.RawQuery = encodedQuery.Encode()

//    fmt.Fprintln(os.Stderr, query)
//    fmt.Fprintln(os.Stderr, encodedQuery)
    return next.ServeHTTP(w, r)
}

// UnmarshalCaddyfile implements caddyfile.Unmarshaler.
func (m *Middleware) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
    return nil
}

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
    var m Middleware
    err := m.UnmarshalCaddyfile(h.Dispenser)
    return m, err
}

// Interface guards
var (
    _ caddy.Provisioner           = (*Middleware)(nil)
    _ caddy.Validator             = (*Middleware)(nil)
    _ caddyhttp.MiddlewareHandler = (*Middleware)(nil)
    _ caddyfile.Unmarshaler       = (*Middleware)(nil)
)