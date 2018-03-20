package gqll

var _ ValueNode = (*baseValueNode)(nil)

type baseValueNode struct {
	baseNode
	//Value ValueNode
}

func (self *baseValueNode) GetValue() interface{} {
	panic("No value")
}
func (self *baseValueNode) IsVariable() bool {
	return false
}

type IntValue struct {
	baseValueNode
	Value int64
}

func (self *IntValue) GetValue() interface{} {
	return self.Value
}

type FloatValue struct {
	baseValueNode
	Value float64
}

func (self *FloatValue) GetValue() interface{} {
	return self.Value
}

type StringValue struct {
	baseValueNode
	Value string
}

func (self *StringValue) GetValue() interface{} {
	return self.Value
}

type BooleanValue struct {
	baseValueNode
	Value bool
}

func (self *BooleanValue) GetValue() interface{} {
	return self.Value
}

type NullValue struct {
	baseValueNode
}

func (self *NullValue) GetValue() interface{} {
	return nil
}

type EnumValue struct {
	baseValueNode
	Value string
}

func (self *EnumValue) GetValue() interface{} {
	return self.Value
}

type ListValue struct {
	baseValueNode
	Value []interface{}
}

func (self *ListValue) GetValue() interface{} {
	return self.Value
}

type ObjectValue struct {
	baseValueNode
	Value map[string]interface{}
}

func (self *ObjectValue) GetValue() interface{} {
	return self.Value
}

type Variable struct {
	baseValueNode
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

func IsValueResolved(node ValueNode) bool {
	switch node := node.(type) {
	case *Variable:
		return node.resolved
	case *ListValue:
		for _, v := range node.Value {
			if v, ok := v.(ValueNode); ok {
				if !IsValueResolved(v) {
					return false
				}
			}
		}
	case *ObjectValue:
		for _, v := range node.Value {
			if v, ok := v.(ValueNode); ok {
				if !IsValueResolved(v) {
					return false
				}
			}
		}

	}
	return true
}

func CollectVariables(node ValueNode, vars *[]*Variable) {
	switch node := node.(type) {
	case *Variable:
		*vars = append(*vars, node)
	case *ListValue:
		for _, v := range node.Value {
			if v, ok := v.(ValueNode); ok {
				CollectVariables(v, vars)
			}
		}
	case *ObjectValue:
		for _, v := range node.Value {
			if v, ok := v.(ValueNode); ok {
				CollectVariables(v, vars)
			}
		}

	}
}
