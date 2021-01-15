package gqll

// fixme unused - remove later
var _ = hasEnumValue{}
var _ = hasExtendTypeName{}

type baseNode struct {
	SourceLocation *SourceLocation
	Comments       []Comment
}

func (self *baseNode) SetSourceLocation(v *SourceLocation) {
	self.SourceLocation = v
}

func (self *baseNode) SetComments(v []Comment) {
	self.Comments = v
}

func (self *baseNode) GetSourceLocation() *SourceLocation {
	return self.SourceLocation
}

func (self *baseNode) GetComments() []Comment {
	return self.Comments
}

type hasName struct {
	Name string
}

func (self *hasName) SetName(v string) {
	self.Name = v
}
func (self *hasName) GetName() string {
	return self.Name
}

type hasExtendTypeName struct {
	ExtendTypeName string
}

func (self *hasExtendTypeName) SetExtendTypeName(v string) {
	self.ExtendTypeName = v
}
func (self *hasExtendTypeName) GetExtendTypeName() string {
	return self.ExtendTypeName
}

type hasEnumValueDefinitions struct {
	EnumValueDefinitions []*EnumValueDefinition
}

func (self *hasEnumValueDefinitions) SetEnumValueDefinitions(v []*EnumValueDefinition) {
	self.EnumValueDefinitions = v
}
func (self *hasEnumValueDefinitions) GetEnumValueDefinitions() []*EnumValueDefinition {
	return self.EnumValueDefinitions
}

type hasEnumValue struct {
	EnumValue string
}

func (self *hasEnumValue) SetEnumValue(v string) {
	self.EnumValue = v
}
func (self *hasEnumValue) GetEnumValue() string {
	return self.EnumValue
}

type hasDirectives struct {
	Directives []*Directive
}

func (self *hasDirectives) SetDirectives(v []*Directive) {
	self.Directives = v
}
func (self *hasDirectives) GetDirectives() []*Directive {
	return self.Directives
}
func (self *hasDirectives) GetDirective(name string) *Directive {
	for _, v := range self.Directives {
		if v.Name == name {
			return v
		}
	}
	return nil
}

type hasArgumentDefinitions struct {
	ArgumentDefinitions []*InputValueDefinition
}

func (self *hasArgumentDefinitions) SetArgumentDefinitions(v []*InputValueDefinition) {
	self.ArgumentDefinitions = v
}
func (self *hasArgumentDefinitions) GetArgumentDefinitions() []*InputValueDefinition {
	return self.ArgumentDefinitions
}

type hasInputFieldDefinitions struct {
	InputFieldDefinitions []*InputValueDefinition
}

func (self *hasInputFieldDefinitions) SetInputFieldDefinitions(v []*InputValueDefinition) {
	self.InputFieldDefinitions = v
}
func (self *hasInputFieldDefinitions) GetInputFieldDefinitions() []*InputValueDefinition {
	return self.InputFieldDefinitions
}

type hasFieldDefinitions struct {
	FieldDefinitions []*FieldDefinition
}

func (self *hasFieldDefinitions) SetFieldDefinitions(v []*FieldDefinition) {
	self.FieldDefinitions = v
}
func (self *hasFieldDefinitions) GetFieldDefinitions() []*FieldDefinition {
	return self.FieldDefinitions
}

type hasArguments struct {
	Arguments []*Argument
}

func (self *hasArguments) SetArguments(v []*Argument) {
	self.Arguments = v
}
func (self *hasArguments) GetArguments() []*Argument {
	return self.Arguments
}
func (self *hasArguments) GetArgument(name string) *Argument {
	for _, v := range self.Arguments {
		if v.Name == name {
			return v
		}
	}
	return nil
}

type hasValue struct {
	Value Value
}

func (self *hasValue) SetValue(v Value) {
	self.Value = v
}
func (self *hasValue) GetValue() Value {
	return self.Value
}

type hasInterfaces struct {
	Interfaces []string
}

func (self *hasInterfaces) SetInterfaces(v []string) {
	self.Interfaces = v
}
func (self *hasInterfaces) GetInterfaces() []string {
	return self.Interfaces
}

type hasUnionMemberTypes struct {
	UnionMemberTypes []string
}

func (self *hasUnionMemberTypes) SetUnionMemberTypes(v []string) {
	self.UnionMemberTypes = v
}
func (self *hasUnionMemberTypes) GetUnionMemberTypes() []string {
	return self.UnionMemberTypes
}

type hasDefaultValue struct {
	DefaultValue []Value
}

func (self *hasDefaultValue) SetDefaultValue(v []Value) {
	self.DefaultValue = v
}
func (self *hasDefaultValue) GetDefaultValue() []Value {
	return self.DefaultValue
}

type hasDescription struct {
	Description string
}

func (self *hasDescription) SetDescription(v string) {
	self.Description = v
}
func (self *hasDescription) GetDescription() string {
	return self.Description
}

type hasType struct {
	Type Type
}

func (self *hasType) SetType(v Type) {
	self.Type = v
}
func (self *hasType) GetType() Type {
	return self.Type
}

type hasSelectionSet struct {
	SelectionSet *SelectionSet
}

func (self *hasSelectionSet) GetSelectionSet() *SelectionSet {
	return self.SelectionSet
}
func (self *hasSelectionSet) SetSelectionSet(v *SelectionSet) {
	self.SelectionSet = v
}

type hasVariableDefinitions struct {
	VariableDefinitions []*VariableDefinition
}

func (self *hasVariableDefinitions) GetVariableDefinitions() []*VariableDefinition {
	return self.VariableDefinitions
}
func (self *hasVariableDefinitions) SetVariableDefinitions(v []*VariableDefinition) {
	self.VariableDefinitions = v
}

type HasName interface {
	SetName(v string)
	GetName() string
}

// The extended type name
type HasExtendTypeName interface {
	SetExtendTypeName(v string)
	GetExtendTypeName() string
}

type HasEnumValueDefinitions interface {
	SetEnumValueDefinitions(v []*EnumValueDefinition)
	GetEnumValueDefinitions() []*EnumValueDefinition
}

type HasDirectives interface {
	SetDirectives(v []*Directive)
	GetDirectives() []*Directive
	GetDirective(name string) *Directive
}

type HasArguments interface {
	SetArguments(v []*Argument)
	GetArguments() []*Argument
	GetArgument(name string) *Argument
}

type HasArgumentDefinitions interface {
	SetArgumentDefinitions(v []*InputValueDefinition)
	GetArgumentDefinitions() []*InputValueDefinition
}

type HasInputFieldDefinitions interface {
	SetInputFieldDefinitions(v []*InputValueDefinition)
	GetInputFieldDefinitions() []*InputValueDefinition
}

type HasFieldDefinitions interface {
	SetFieldDefinitions(v []*FieldDefinition)
	GetFieldDefinitions() []*FieldDefinition
}

type HasInterfaces interface {
	SetInterfaces(v []string)
	GetInterfaces() []string
}
