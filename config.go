package enum

// Options defines the specification of a enum-codegen file which is used
// as data input to generated the types and type tests during template execution
type Options struct {
	PackageName string `yaml:"package"`
	Types       []Type `yaml:"types"`
}

// Type defines the data fields of a type used when executing the template codegen
type Type struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Enums       []string `yaml:"enums"`
}
