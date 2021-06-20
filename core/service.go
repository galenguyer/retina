package core

type Service struct {
	Name  string `yaml:"name"`
	Group string `yaml:"group,omitempty"`
	URL   string `yaml:"url"`
}
