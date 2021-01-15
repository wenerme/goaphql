// Code generated from GraphQL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package gqlp // GraphQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by GraphQLParser.
type GraphQLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GraphQLParser#graphql.
	VisitGraphql(ctx *GraphqlContext) interface{}

	// Visit a parse tree produced by GraphQLParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by GraphQLParser#document.
	VisitDocument(ctx *DocumentContext) interface{}

	// Visit a parse tree produced by GraphQLParser#definition.
	VisitDefinition(ctx *DefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#executableDefinition.
	VisitExecutableDefinition(ctx *ExecutableDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#operationDefinition.
	VisitOperationDefinition(ctx *OperationDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#operationType.
	VisitOperationType(ctx *OperationTypeContext) interface{}

	// Visit a parse tree produced by GraphQLParser#selectionSet.
	VisitSelectionSet(ctx *SelectionSetContext) interface{}

	// Visit a parse tree produced by GraphQLParser#selection.
	VisitSelection(ctx *SelectionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by GraphQLParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by GraphQLParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by GraphQLParser#alias.
	VisitAlias(ctx *AliasContext) interface{}

	// Visit a parse tree produced by GraphQLParser#fragmentSpread.
	VisitFragmentSpread(ctx *FragmentSpreadContext) interface{}

	// Visit a parse tree produced by GraphQLParser#fragmentDefinition.
	VisitFragmentDefinition(ctx *FragmentDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#fragmentName.
	VisitFragmentName(ctx *FragmentNameContext) interface{}

	// Visit a parse tree produced by GraphQLParser#typeCondition.
	VisitTypeCondition(ctx *TypeConditionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#inlineFragment.
	VisitInlineFragment(ctx *InlineFragmentContext) interface{}

	// Visit a parse tree produced by GraphQLParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by GraphQLParser#intValue.
	VisitIntValue(ctx *IntValueContext) interface{}

	// Visit a parse tree produced by GraphQLParser#floatValue.
	VisitFloatValue(ctx *FloatValueContext) interface{}

	// Visit a parse tree produced by GraphQLParser#booleanValue.
	VisitBooleanValue(ctx *BooleanValueContext) interface{}

	// Visit a parse tree produced by GraphQLParser#stringValue.
	VisitStringValue(ctx *StringValueContext) interface{}

	// Visit a parse tree produced by GraphQLParser#nullValue.
	VisitNullValue(ctx *NullValueContext) interface{}

	// Visit a parse tree produced by GraphQLParser#enumValue.
	VisitEnumValue(ctx *EnumValueContext) interface{}

	// Visit a parse tree produced by GraphQLParser#listValue.
	VisitListValue(ctx *ListValueContext) interface{}

	// Visit a parse tree produced by GraphQLParser#objectValue.
	VisitObjectValue(ctx *ObjectValueContext) interface{}

	// Visit a parse tree produced by GraphQLParser#objectField.
	VisitObjectField(ctx *ObjectFieldContext) interface{}

	// Visit a parse tree produced by GraphQLParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by GraphQLParser#variableDefinitions.
	VisitVariableDefinitions(ctx *VariableDefinitionsContext) interface{}

	// Visit a parse tree produced by GraphQLParser#variableDefinition.
	VisitVariableDefinition(ctx *VariableDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#defaultValue.
	VisitDefaultValue(ctx *DefaultValueContext) interface{}

	// Visit a parse tree produced by GraphQLParser#typeSpec.
	VisitTypeSpec(ctx *TypeSpecContext) interface{}

	// Visit a parse tree produced by GraphQLParser#namedType.
	VisitNamedType(ctx *NamedTypeContext) interface{}

	// Visit a parse tree produced by GraphQLParser#listType.
	VisitListType(ctx *ListTypeContext) interface{}

	// Visit a parse tree produced by GraphQLParser#nonNullType.
	VisitNonNullType(ctx *NonNullTypeContext) interface{}

	// Visit a parse tree produced by GraphQLParser#directives.
	VisitDirectives(ctx *DirectivesContext) interface{}

	// Visit a parse tree produced by GraphQLParser#directive.
	VisitDirective(ctx *DirectiveContext) interface{}

	// Visit a parse tree produced by GraphQLParser#typeSystemDefinition.
	VisitTypeSystemDefinition(ctx *TypeSystemDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#schemaDefinition.
	VisitSchemaDefinition(ctx *SchemaDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#operationTypeDefinition.
	VisitOperationTypeDefinition(ctx *OperationTypeDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#description.
	VisitDescription(ctx *DescriptionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#typeDefinition.
	VisitTypeDefinition(ctx *TypeDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#typeExtension.
	VisitTypeExtension(ctx *TypeExtensionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#scalarTypeDefinition.
	VisitScalarTypeDefinition(ctx *ScalarTypeDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#scalarTypeExtension.
	VisitScalarTypeExtension(ctx *ScalarTypeExtensionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#objectTypeDefinition.
	VisitObjectTypeDefinition(ctx *ObjectTypeDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#objectTypeExtension.
	VisitObjectTypeExtension(ctx *ObjectTypeExtensionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#implementsInterfaces.
	VisitImplementsInterfaces(ctx *ImplementsInterfacesContext) interface{}

	// Visit a parse tree produced by GraphQLParser#fieldsDefinition.
	VisitFieldsDefinition(ctx *FieldsDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#fieldDefinition.
	VisitFieldDefinition(ctx *FieldDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#argumentsDefinition.
	VisitArgumentsDefinition(ctx *ArgumentsDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#inputValueDefinition.
	VisitInputValueDefinition(ctx *InputValueDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#interfaceTypeDefinition.
	VisitInterfaceTypeDefinition(ctx *InterfaceTypeDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#interfaceTypeExtension.
	VisitInterfaceTypeExtension(ctx *InterfaceTypeExtensionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#unionTypeDefinition.
	VisitUnionTypeDefinition(ctx *UnionTypeDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#unionMemberTypes.
	VisitUnionMemberTypes(ctx *UnionMemberTypesContext) interface{}

	// Visit a parse tree produced by GraphQLParser#unionTypeExtension.
	VisitUnionTypeExtension(ctx *UnionTypeExtensionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#enumTypeDefinition.
	VisitEnumTypeDefinition(ctx *EnumTypeDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#enumValuesDefinition.
	VisitEnumValuesDefinition(ctx *EnumValuesDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#enumValueDefinition.
	VisitEnumValueDefinition(ctx *EnumValueDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#enumTypeExtension.
	VisitEnumTypeExtension(ctx *EnumTypeExtensionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#inputObjectTypeDefinition.
	VisitInputObjectTypeDefinition(ctx *InputObjectTypeDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#inputFieldsDefinition.
	VisitInputFieldsDefinition(ctx *InputFieldsDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#inputObjectTypeExtension.
	VisitInputObjectTypeExtension(ctx *InputObjectTypeExtensionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#directiveDefinition.
	VisitDirectiveDefinition(ctx *DirectiveDefinitionContext) interface{}

	// Visit a parse tree produced by GraphQLParser#directiveLocations.
	VisitDirectiveLocations(ctx *DirectiveLocationsContext) interface{}
}
