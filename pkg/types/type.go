package types

import (
    "strings"
)

// A Package is a buildable unit within the source packages
// collection.
type Package struct {
	Name        string
	Dirty       bool
	Failed      bool
	Version     string `plist:"pkgver"`
	HostDepends map[string]struct{}
	MakeDepends map[string]struct{}
	Depends     map[string]struct{}
	Subpackages map[string]struct{}
}

func (p Package) String() string {
	return p.Name + "-" + p.Version
}

// Atom is a storage struct that contains all the serializable data
// for a single arch graph.
type Atom struct {
	Pkgs    map[string]*Package
	Virtual map[string]string

	// bad returned some errors, so we keep an eye on what the
	// error was and continue.
	Bad map[string]string

	// These keep track of what the archs this graph is rendered
	// from are.
	Spec SpecTuple

	// Rev stores the git revision of the PkgGraph for later so
	// that we can tell if the graph needs to be reloaded.
	Rev string
}

// A SpecTuple is a listing of the host and target arch.
type SpecTuple struct {
	Host   string
	Target string
}

func (st SpecTuple) String() string {
	return st.Host + ":" + st.Target
}

// SpecTupleFromString returns a spec tuple from its string
// representation.
func SpecTupleFromString(s string) SpecTuple {
	p := strings.SplitN(s, ":", 2)
	return SpecTuple{p[0], p[1]}
}
