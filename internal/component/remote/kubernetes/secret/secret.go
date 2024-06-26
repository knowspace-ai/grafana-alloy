package secret

import (
	"github.com/grafana/alloy/internal/component"
	"github.com/grafana/alloy/internal/component/remote/kubernetes"
	"github.com/grafana/alloy/internal/featuregate"
)

func init() {
	component.Register(component.Registration{
		Name:      "remote.kubernetes.secret",
		Stability: featuregate.StabilityGenerallyAvailable,
		Args:      kubernetes.Arguments{},
		Exports:   kubernetes.Exports{},
		Build: func(opts component.Options, args component.Arguments) (component.Component, error) {
			return kubernetes.New(opts, args.(kubernetes.Arguments), kubernetes.TypeSecret)
		},
	})
}
