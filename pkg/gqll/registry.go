package gqll

import (
	"reflect"
	"strings"
)

var registryNameMap = make(map[string]*registryEntry)
var registryTypeMap = make(map[reflect.Type]*registryEntry)

func registryType(types ...interface{}) {
	for _, v := range types {
		entry := &registryEntry{}
		typ := reflect.TypeOf(v).Elem()
		entry.name = typ.Name()
		entry.typ = typ

		registryNameMap[entry.name] = entry
		registryTypeMap[entry.typ] = entry
	}
}

type registryEntry struct {
	name string
	typ  reflect.Type
}

func (self *registryEntry) Name() string {
	return self.name
}

func (self *registryEntry) Type() reflect.Type {
	return self.typ
}

func TypeOfName(name string) TypeEntry {
	return registryNameMap[name]
}
func TypeOf(v interface{}) TypeEntry {
	t := reflect.TypeOf(v)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return registryTypeMap[t]
}

type TypeEntry interface {
	// Name of this type
	Name() string
	// Struct type
	Type() reflect.Type

	IsTypeDefinition() bool
	IsTypeExtension() bool
}

func (self *registryEntry) IsTypeDefinition() bool {
	return strings.HasSuffix(self.name, "TypeDefinition")
}
func (self *registryEntry) IsTypeExtension() bool {
	return strings.HasSuffix(self.name, "TypeExtension")
}

func init() {
	registryType(
		// TypeDefinition
		(*SchemaDefinition)(nil),
		(*ScalarTypeDefinition)(nil),
		(*InterfaceTypeDefinition)(nil),
		(*UnionTypeDefinition)(nil),
		(*ObjectTypeDefinition)(nil),
		(*InputObjectTypeDefinition)(nil),
		(*EnumTypeDefinition)(nil),

		// TypeExtension
		(*ScalarTypeExtension)(nil),
		(*InterfaceTypeExtension)(nil),
		(*UnionTypeExtension)(nil),
		(*ObjectTypeExtension)(nil),
		(*InputObjectTypeExtension)(nil),
		(*EnumTypeExtension)(nil),
		(*SchemaTypeExtension)(nil),

		// Directive
		(*Directive)(nil),
		(*DirectiveDefinition)(nil),

		// OperationDefinition
		(*OperationDefinition)(nil),
		(*FragmentDefinition)(nil),
		(*SelectionSet)(nil),
		(*Field)(nil),
		(*VariableDefinition)(nil),
		(*FragmentSpread)(nil),
		(*InlineFragment)(nil),

		// Type
		(*NonNullType)(nil),
		(*ListType)(nil),
		(*NamedType)(nil),

		// Value
		(*IntValue)(nil),
		(*FloatValue)(nil),
		(*StringValue)(nil),
		(*BooleanValue)(nil),
		(*NullValue)(nil),
		(*EnumValue)(nil),
		(*ListValue)(nil),
		(*ObjectValue)(nil),
		(*Variable)(nil),

		// Misc
		(*Comment)(nil),
		(*Document)(nil),
		(*Argument)(nil),
		(*Directive)(nil),
		(*FieldDefinition)(nil),
		(*InputValueDefinition)(nil),
		(*EnumValueDefinition)(nil),
	)
}
