package gqll

type TypeDefinition interface {
	GetName() string
	GetDescription() string
	GetDirectives() []*Directive
}
type TypeExtension interface {
	GetExtendTypeName() string
	GetDirectives() []*Directive
}

// region schema
type SchemaDefinition struct {
	baseNode
	hasDirectives
	hasName
	QueryTypeName        string
	MutationTypeName     string
	SubscriptionTypeName string
}

// endregion

// region type

type baseTypeDefinition struct {
	baseNode
	hasName
	hasDescription
	hasDirectives
}

type ScalarTypeDefinition struct {
	baseTypeDefinition
}
type InterfaceTypeDefinition struct {
	baseTypeDefinition
	hasFieldDefinitions
}
type UnionTypeDefinition struct {
	baseTypeDefinition
	hasUnionMemberTypes
}
type ObjectTypeDefinition struct {
	baseTypeDefinition
	hasInterfaces
	hasFieldDefinitions
}
type InputObjectTypeDefinition struct {
	baseTypeDefinition
}
type EnumTypeDefinition struct {
	baseTypeDefinition
	hasEnumValueDefinitions
}

type DirectiveDefinition struct {
	baseTypeDefinition
	hasArgumentDefinitions
	Locations []string
}

// endregion

// region extension

type baseTypeExtension struct {
	baseNode
	hasName
	hasDirectives
	ExtendTypeName string
}

func (self *baseTypeExtension) GetExtendTypeName() string {
	return self.ExtendTypeName
}
func (self *baseTypeExtension) SetExtendTypeName(name string) {
	self.ExtendTypeName = name
}

type ScalarTypeExtension struct {
	baseTypeExtension
}

type InterfaceTypeExtension struct {
	baseTypeExtension
	hasFieldDefinitions
}

type UnionTypeExtension struct {
	baseTypeExtension
	hasUnionMemberTypes
}

type ObjectTypeExtension struct {
	baseTypeExtension
	hasInterfaces
	hasFieldDefinitions
}

type InputObjectTypeExtension struct {
	baseTypeExtension
}

type EnumTypeExtension struct {
	baseTypeExtension
}

// endregion
