package dispatchable

import (
	"github.com/hashicorp/go-hclog"

	"github.com/the-maldridge/nbuild/pkg/types"
)

func NewDispatchFinder(l hclog.Logger, sliceAtoms []types.Atom) *DispatchFinder {
	mAtoms := make(map[types.SpecTuple]types.Atom)
	for _, atom := range sliceAtoms {
		mAtoms[atom.Spec] = atom
	}
	x := DispatchFinder{
		l:     l,
		atoms: mAtoms,
	}
	return &x
}

// IsDispatchable determines whether a specific package could be dispatched
// right now.
func (d *DispatchFinder) IsDispatchable(spec types.SpecTuple, p *types.Package) bool {
	hAtom := d.atoms[types.SpecTuple{spec.Host, spec.Host}]
	for hdep, _ := range p.HostDepends {
		if hAtom.Pkgs[hdep].Dirty {
			return false
		}
	}

	tAtom := d.atoms[spec]
	for dep, _ := range p.MakeDepends {
		if tAtom.Pkgs[dep].Dirty {
			return false
		}
	}
	for dep, _ := range p.Depends {
		if tAtom.Pkgs[dep].Dirty {
			return false
		}
	}
	// If we get this far, all hostdeps, makedeps, deps are clean.
	return true
}

// ImmediatelyDispatchable returns a map of tuples -> packages that can be
// hypothetically dispatched right now.
// *Assumes graph is freshly Cleaned*.
func (d *DispatchFinder) ImmediatelyDispatchable() map[types.SpecTuple][]*types.Package {
	dispatchable := make(map[types.SpecTuple][]*types.Package)
	for spec, atom := range d.atoms {
		// We need the host atom as well to be able to do this
		if _, ok := d.atoms[types.SpecTuple{spec.Host, spec.Host}]; !ok {
			d.l.Warn("Unable to find dispatchable due to lack of host atom", "spec", spec)
			continue
		}
		dispatchable[spec] = make([]*types.Package, 0)
		for _, pkg := range atom.Pkgs {
			if pkg.Dirty && d.IsDispatchable(spec, pkg) {
				dispatchable[spec] = append(dispatchable[spec], pkg)
			}
		}
	}
	return dispatchable
}
