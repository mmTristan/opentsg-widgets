package zoneplate

import (
	_ "embed"

	"github.com/mrmxf/opentsg-core/config"
)

/*
// zoneplate definitions
const wName = "zone plate"
const wType = "zoneplate"
const wLibrary = "builtin"
const hooks = ""*/

type zoneplateJSON struct {
	Platetype   string       `json:"platetype" yaml:"platetype"`
	Startcolour string       `json:"startcolor" yaml:"startcolor"`
	Angle       interface{}  `json:"angle" yaml:"angle"`
	Mask        string       `json:"mask" yaml:"mask"`
	GridLoc     *config.Grid `json:"grid,omitempty" yaml:"grid,omitempty"`
}

//go:embed jsonschema/zoneplateschema.json
var schemaInit []byte

func (z zoneplateJSON) Alias() string {
	return z.GridLoc.Alias
}

func (z zoneplateJSON) Location() string {
	return z.GridLoc.Location
}
