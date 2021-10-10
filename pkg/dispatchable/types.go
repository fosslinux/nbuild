package dispatchable

import (
	"github.com/hashicorp/go-hclog"

	"github.com/the-maldridge/nbuild/pkg/types"
)

// DispatchFinder
type DispatchFinder struct {
	l hclog.Logger

	atoms map[types.SpecTuple]types.Atom
}
