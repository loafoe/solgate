package solgate

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/loafoe/solgate/storer"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Middleware struct {
	Endpoint string
	logger   *zap.Logger
	store    *storer.Solgate
}

// CaddyModule returns the Caddy module information.
func (m *Middleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.solgate",
		New: func() caddy.Module { return new(Middleware) },
	}
}

func init() {
	caddy.RegisterModule(&Middleware{})
	httpcaddyfile.RegisterHandlerDirective("solgate", parseCaddyfile)
}

// Provision implements caddy.Provisioner.
func (m *Middleware) Provision(ctx caddy.Context) error {
	m.logger = ctx.Logger() // g.logger is a *zap.Logger
	return nil
}

// Validate implements caddy.Validator.
func (m *Middleware) Validate() error {
	return nil
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	// TODO: implement path checking here
	token := r.URL.Path
	found, err := m.store.Token.FindByToken(token)
	if err != nil || found.ExpiresAt.After(time.Now()) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"message": "permission denied or token expired"}`))
		return err
	}
	return next.ServeHTTP(w, r)
}

// UnmarshalCaddyfile sets up Lessor from Caddyfile tokens. Syntax:
// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	solgate [<issuer_url>] {
//	    issuer <issuer_url>
//	}
func (m *Middleware) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			m.Endpoint = d.Val()
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "endpoint":
				if m.Endpoint != "" {
					return d.Err("Endpoint already set")
				}
				m.Endpoint = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	return nil
}

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	m := &Middleware{}
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
