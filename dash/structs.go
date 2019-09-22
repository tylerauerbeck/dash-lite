package dash

import (
	v2 "github.com/tylerauerbeck/dash/dash/v2"
	v3 "github.com/tylerauerbeck/dash/dash/v3"
)

type Inventory struct {
	Version	string	`yaml:"version,omitempty"`
	Context	string	`yaml:"context,omitempty"`
	Namespace	string	`yaml:"namespace,omitempty"`
	ClusterContentList	v2.ClusterContentList	`yaml:",omitempty"`
	Resources	[]v3.ResourceGroup	`yaml:"resource_groups,omitempty"`	
}
