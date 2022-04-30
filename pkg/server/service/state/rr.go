package state

import (
	"github.com/traefik/traefik/v2/pkg/healthcheck"
)

var LoadBalancer map[string]*healthcheck.LbStatusUpdater

func init() {
	LoadBalancer = make(map[string]*healthcheck.LbStatusUpdater)
}
