package terraform

import (
	"log"

	"github.com/hashicorp/terraform/config"
	"github.com/hashicorp/terraform/config/configschema"
)

// EvalInterpolate is an EvalNode implementation that takes a raw
// configuration and interpolates it.
type EvalInterpolate struct {
	Config        *config.RawConfig
	Schema        *configschema.Block
	Resource      *Resource
	Output        **ResourceConfig
	ContinueOnErr bool
}

func (n *EvalInterpolate) Eval(ctx EvalContext) (interface{}, error) {
	rc, err := ctx.Interpolate(n.Config, n.Resource, n.Schema)
	if err != nil {
		if n.ContinueOnErr {
			log.Printf("[WARN] Interpolation %q failed: %s", n.Config.Key, err)
			return nil, EvalEarlyExitError{}
		}
		return nil, err
	}

	if n.Output != nil {
		*n.Output = rc
	}

	return nil, nil
}
