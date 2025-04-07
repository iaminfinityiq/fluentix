package value_types

type Object interface {
	Name_() string
	Properties_() map[string]Object
	Constants_() map[string]bool
}

type DefinedObject struct {
	Name string
	Properties map[string]Object
}

func (d DefinedObject) Name_() string {
	return d.Name
}

func (d DefinedObject) Properties_() map[string]Object {
	return d.Properties
}

func (d DefinedObject) Constants_() map[string]bool {
	return make(map[string]bool)
}

type ValuedObject struct {
	Name string
	Properties map[string]Object
	Constants map[string]bool
	Value any
}

func (v ValuedObject) Name_() string {
	return v.Name
}

func (v ValuedObject) Properties_() map[string]Object {
	return v.Properties
}

func (v ValuedObject) Constants_() map[string]bool {
	return v.Constants
}