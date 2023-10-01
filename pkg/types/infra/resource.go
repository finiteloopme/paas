package infra

type ResourceType struct {
	Parent      ParentType `yaml:"parent,omitempty" json:"parent,omitempty"`
	Id          string     `yaml:"id" json:"id"`
	Kind        string     `yaml:"kind,omitempty" json:"kind,omitempty"`
	Description string     `yaml:"description,omitempty" json:"description,omitempty"`
}

type ParentType struct {
	RefId string `yaml:"refId" json:"id"`
	Kind  string `yaml:"kind,omitempty" json:"kind,omitempty"`
}
