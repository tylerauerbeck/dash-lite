package v3

// ResourceGroup is a logical collection of items that will be applied together
type ResourceGroup struct {
	Name      string `yaml:"name,omitempty"`
	Namespace string `yaml:"namespace,omitempty"`
	Context   string `yaml:"context,omitempty"`
	Resources []ResourceGroupItem
}

// ResourceGroupItem is a single item within a ResourceGroup
type ResourceGroupItem struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace,omitempty"`
	Context   string `yaml:"context,omitempty"`
	Type      string `yaml:"type"`
	Template  string `yaml:"template,omitempty"`
	Params    string `yaml:"params,omitempty"`
	File      string `yaml:"file,omitempty"`
}
