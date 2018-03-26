package gqll

var _ Value = (*baseValue)(nil)

type baseValue struct {
	baseNode
}

func (self *baseValue) GetValue() interface{} {
	panic("No value")
}
func (self *baseValue) IsVariable() bool {
	return false
}
func (self *baseValue) IsValue() bool {
	return true
}

type IntValue struct {
	baseValue
	Value int64
}

func (self *IntValue) GetValue() interface{} {
	return self.Value
}

type FloatValue struct {
	baseValue
	Value float64
}

func (self *FloatValue) GetValue() interface{} {
	return self.Value
}

type StringValue struct {
	baseValue
	Value string
}

func (self *StringValue) GetValue() interface{} {
	return self.Value
}

type BooleanValue struct {
	baseValue
	Value bool
}

func (self *BooleanValue) GetValue() interface{} {
	return self.Value
}

type NullValue struct {
	baseValue
}

func (self *NullValue) GetValue() interface{} {
	return nil
}

type EnumValue struct {
	baseValue
	Value string // Name of enum
}

func (self *EnumValue) GetValue() interface{} {
	return self.Value
}

type ListValue struct {
	baseValue
	Value []interface{}
}

func (self *ListValue) GetValue() interface{} {
	return self.Value
}

type ObjectValue struct {
	baseValue
	Value map[string]interface{}
}

func (self *ObjectValue) GetValue() interface{} {
	return self.Value
}

type Variable struct {
	baseValue
	Name  string
	Value interface{}

	resolved bool
}

func (self *Variable) GetValue() interface{} {
	if !self.resolved {
		panic("Variable " + self.Name + " not resolved")
	}
	return self.Value
}
func (self *Variable) Resolve(v interface{}) {
	self.resolved = true
	self.Value = v
}
func (self *Variable) IsVariable() bool {
	return true
}

func IsValueResolved(node Value) bool {
	switch node := node.(type) {
	case *Variable:
		return node.resolved
	case *ListValue:
		for _, v := range node.Value {
			if v, ok := v.(Value); ok {
				if !IsValueResolved(v) {
					return false
				}
			}
		}
	case *ObjectValue:
		for _, v := range node.Value {
			if v, ok := v.(Value); ok {
				if !IsValueResolved(v) {
					return false
				}
			}
		}

	}
	return true
}

func CollectVariables(node Value, vars *[]*Variable) {
	switch node := node.(type) {
	case *Variable:
		*vars = append(*vars, node)
	case *ListValue:
		for _, v := range node.Value {
			if v, ok := v.(Value); ok {
				CollectVariables(v, vars)
			}
		}
	case *ObjectValue:
		for _, v := range node.Value {
			if v, ok := v.(Value); ok {
				CollectVariables(v, vars)
			}
		}

	}
}
