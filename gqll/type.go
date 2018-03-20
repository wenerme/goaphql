package gqll

type baseTypeNode struct {
	baseNode
}

func (self *baseTypeNode) IsNonNull() bool {
	return false
}

func (self *baseTypeNode) IsNamed() bool {
	return false
}

func (self *baseTypeNode) IsList() bool {
	return false
}

func (self *baseTypeNode) GetType() TypeNode {
	return nil
}

func (self *baseTypeNode) GetName() string {
	return ""
}

type NonNullType struct {
	baseTypeNode
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
	baseTypeNode
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
	baseTypeNode
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
