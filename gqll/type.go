package gqll

type baseType struct {
	baseNode
}

func (self *baseType) IsNonNull() bool {
	return false
}

func (self *baseType) IsNamed() bool {
	return false
}

func (self *baseType) IsList() bool {
	return false
}

func (self *baseType) GetType() TypeNode {
	return nil
}

func (self *baseType) GetName() string {
	return ""
}

func (self *baseType) IsType() bool {
	return true
}

type NonNullType struct {
	baseType
	Type TypeNode
}

func (self *NonNullType) SetType(node TypeNode) {
	self.Type = node
}
func (self *NonNullType) GetType() TypeNode {
	return self.Type
}
func (self *NonNullType) IsNonNull() bool {
	return true
}

type ListType struct {
	baseType
	Type TypeNode
}

func (self *ListType) IsList() bool {
	return true
}
func (self *ListType) SetType(node TypeNode) {
	self.Type = node
}
func (self *ListType) GetType() TypeNode {
	return self.Type
}

type NamedType struct {
	baseType
	Name string
}

func (self *NamedType) IsNamed() bool {
	return true
}
func (self *NamedType) SetName(name string) {
	self.Name = name
}
func (self *NamedType) GetName() string {
	return self.Name
}
