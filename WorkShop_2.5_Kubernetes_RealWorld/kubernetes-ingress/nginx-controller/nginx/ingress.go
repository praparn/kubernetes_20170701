package nginx

import (
	api_v1 "k8s.io/api/core/v1"
	extensions "k8s.io/api/extensions/v1beta1"
)

// IngressEx holds an Ingress along with Secrets and Endpoints of the services
// that are referenced in this Ingress
type IngressEx struct {
	Ingress    *extensions.Ingress
	TLSSecrets map[string]*api_v1.Secret
	JWTKey     *api_v1.Secret
	Endpoints  map[string][]string
}

type MergeableIngresses struct {
	Master  *IngressEx
	Minions []*IngressEx
}

var masterBlacklist = map[string]bool{
	"nginx.org/rewrites":               true,
	"nginx.org/ssl-services":           true,
	"nginx.org/websocket-services":     true,
	"nginx.com/sticky-cookie-services": true,
}

var minionBlacklist = map[string]bool{
	"nginx.org/proxy-hide-headers":       true,
	"nginx.org/proxy-pass-headers":       true,
	"nginx.org/redirect-to-https":        true,
	"ingress.kubernetes.io/ssl-redirect": true,
	"nginx.org/hsts":                     true,
	"nginx.org/hsts-max-age":             true,
	"nginx.org/hsts-include-subdomains":  true,
	"nginx.org/server-tokens":            true,
	"nginx.org/listen-ports":             true,
	"nginx.org/listen-ports-ssl":         true,
	"nginx.com/jwt-key":                  true,
	"nginx.com/jwt-realm":                true,
	"nginx.com/jwt-token":                true,
	"nginx.com/jwt-login-url":            true,
}
