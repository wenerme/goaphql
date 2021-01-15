// Code generated from GraphQL.g4 by ANTLR 4.9.1. DO NOT EDIT.

package gqlp // GraphQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseGraphQLVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGraphQLVisitor) VisitGraphql(ctx *GraphqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitDocument(ctx *DocumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitDefinition(ctx *DefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitExecutableDefinition(ctx *ExecutableDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitOperationDefinition(ctx *OperationDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitOperationType(ctx *OperationTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitSelectionSet(ctx *SelectionSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitSelection(ctx *SelectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitAlias(ctx *AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitFragmentSpread(ctx *FragmentSpreadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitFragmentDefinition(ctx *FragmentDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitFragmentName(ctx *FragmentNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitTypeCondition(ctx *TypeConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitInlineFragment(ctx *InlineFragmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitIntValue(ctx *IntValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitFloatValue(ctx *FloatValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitBooleanValue(ctx *BooleanValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitStringValue(ctx *StringValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitNullValue(ctx *NullValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitEnumValue(ctx *EnumValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitListValue(ctx *ListValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitObjectValue(ctx *ObjectValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitObjectField(ctx *ObjectFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitVariableDefinitions(ctx *VariableDefinitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitVariableDefinition(ctx *VariableDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitDefaultValue(ctx *DefaultValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitTypeSpec(ctx *TypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitNamedType(ctx *NamedTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitListType(ctx *ListTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitNonNullType(ctx *NonNullTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitDirectives(ctx *DirectivesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitDirective(ctx *DirectiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitTypeSystemDefinition(ctx *TypeSystemDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitSchemaDefinition(ctx *SchemaDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitOperationTypeDefinition(ctx *OperationTypeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitDescription(ctx *DescriptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitTypeDefinition(ctx *TypeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitTypeExtension(ctx *TypeExtensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitScalarTypeDefinition(ctx *ScalarTypeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitScalarTypeExtension(ctx *ScalarTypeExtensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitObjectTypeDefinition(ctx *ObjectTypeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitObjectTypeExtension(ctx *ObjectTypeExtensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitImplementsInterfaces(ctx *ImplementsInterfacesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitFieldsDefinition(ctx *FieldsDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitFieldDefinition(ctx *FieldDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitArgumentsDefinition(ctx *ArgumentsDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitInputValueDefinition(ctx *InputValueDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitInterfaceTypeDefinition(ctx *InterfaceTypeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitInterfaceTypeExtension(ctx *InterfaceTypeExtensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitUnionTypeDefinition(ctx *UnionTypeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitUnionMemberTypes(ctx *UnionMemberTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitUnionTypeExtension(ctx *UnionTypeExtensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitEnumTypeDefinition(ctx *EnumTypeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitEnumValuesDefinition(ctx *EnumValuesDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitEnumValueDefinition(ctx *EnumValueDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitEnumTypeExtension(ctx *EnumTypeExtensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitInputObjectTypeDefinition(ctx *InputObjectTypeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitInputFieldsDefinition(ctx *InputFieldsDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitInputObjectTypeExtension(ctx *InputObjectTypeExtensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitDirectiveDefinition(ctx *DirectiveDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphQLVisitor) VisitDirectiveLocations(ctx *DirectiveLocationsContext) interface{} {
	return v.VisitChildren(ctx)
}
