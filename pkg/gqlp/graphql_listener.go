// Code generated from GraphQL.g4 by ANTLR 4.9.1. DO NOT EDIT.

package gqlp // GraphQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// GraphQLListener is a complete listener for a parse tree produced by GraphQLParser.
type GraphQLListener interface {
	antlr.ParseTreeListener

	// EnterGraphql is called when entering the graphql production.
	EnterGraphql(c *GraphqlContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterExecutableDefinition is called when entering the executableDefinition production.
	EnterExecutableDefinition(c *ExecutableDefinitionContext)

	// EnterOperationDefinition is called when entering the operationDefinition production.
	EnterOperationDefinition(c *OperationDefinitionContext)

	// EnterOperationType is called when entering the operationType production.
	EnterOperationType(c *OperationTypeContext)

	// EnterSelectionSet is called when entering the selectionSet production.
	EnterSelectionSet(c *SelectionSetContext)

	// EnterSelection is called when entering the selection production.
	EnterSelection(c *SelectionContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterFragmentSpread is called when entering the fragmentSpread production.
	EnterFragmentSpread(c *FragmentSpreadContext)

	// EnterFragmentDefinition is called when entering the fragmentDefinition production.
	EnterFragmentDefinition(c *FragmentDefinitionContext)

	// EnterFragmentName is called when entering the fragmentName production.
	EnterFragmentName(c *FragmentNameContext)

	// EnterTypeCondition is called when entering the typeCondition production.
	EnterTypeCondition(c *TypeConditionContext)

	// EnterInlineFragment is called when entering the inlineFragment production.
	EnterInlineFragment(c *InlineFragmentContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterIntValue is called when entering the intValue production.
	EnterIntValue(c *IntValueContext)

	// EnterFloatValue is called when entering the floatValue production.
	EnterFloatValue(c *FloatValueContext)

	// EnterBooleanValue is called when entering the booleanValue production.
	EnterBooleanValue(c *BooleanValueContext)

	// EnterStringValue is called when entering the stringValue production.
	EnterStringValue(c *StringValueContext)

	// EnterNullValue is called when entering the nullValue production.
	EnterNullValue(c *NullValueContext)

	// EnterEnumValue is called when entering the enumValue production.
	EnterEnumValue(c *EnumValueContext)

	// EnterListValue is called when entering the listValue production.
	EnterListValue(c *ListValueContext)

	// EnterObjectValue is called when entering the objectValue production.
	EnterObjectValue(c *ObjectValueContext)

	// EnterObjectField is called when entering the objectField production.
	EnterObjectField(c *ObjectFieldContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterVariableDefinitions is called when entering the variableDefinitions production.
	EnterVariableDefinitions(c *VariableDefinitionsContext)

	// EnterVariableDefinition is called when entering the variableDefinition production.
	EnterVariableDefinition(c *VariableDefinitionContext)

	// EnterDefaultValue is called when entering the defaultValue production.
	EnterDefaultValue(c *DefaultValueContext)

	// EnterTypeSpec is called when entering the typeSpec production.
	EnterTypeSpec(c *TypeSpecContext)

	// EnterNamedType is called when entering the namedType production.
	EnterNamedType(c *NamedTypeContext)

	// EnterListType is called when entering the listType production.
	EnterListType(c *ListTypeContext)

	// EnterNonNullType is called when entering the nonNullType production.
	EnterNonNullType(c *NonNullTypeContext)

	// EnterDirectives is called when entering the directives production.
	EnterDirectives(c *DirectivesContext)

	// EnterDirective is called when entering the directive production.
	EnterDirective(c *DirectiveContext)

	// EnterTypeSystemDefinition is called when entering the typeSystemDefinition production.
	EnterTypeSystemDefinition(c *TypeSystemDefinitionContext)

	// EnterSchemaDefinition is called when entering the schemaDefinition production.
	EnterSchemaDefinition(c *SchemaDefinitionContext)

	// EnterOperationTypeDefinition is called when entering the operationTypeDefinition production.
	EnterOperationTypeDefinition(c *OperationTypeDefinitionContext)

	// EnterDescription is called when entering the description production.
	EnterDescription(c *DescriptionContext)

	// EnterTypeDefinition is called when entering the typeDefinition production.
	EnterTypeDefinition(c *TypeDefinitionContext)

	// EnterTypeExtension is called when entering the typeExtension production.
	EnterTypeExtension(c *TypeExtensionContext)

	// EnterScalarTypeDefinition is called when entering the scalarTypeDefinition production.
	EnterScalarTypeDefinition(c *ScalarTypeDefinitionContext)

	// EnterScalarTypeExtension is called when entering the scalarTypeExtension production.
	EnterScalarTypeExtension(c *ScalarTypeExtensionContext)

	// EnterObjectTypeDefinition is called when entering the objectTypeDefinition production.
	EnterObjectTypeDefinition(c *ObjectTypeDefinitionContext)

	// EnterObjectTypeExtension is called when entering the objectTypeExtension production.
	EnterObjectTypeExtension(c *ObjectTypeExtensionContext)

	// EnterImplementsInterfaces is called when entering the implementsInterfaces production.
	EnterImplementsInterfaces(c *ImplementsInterfacesContext)

	// EnterFieldsDefinition is called when entering the fieldsDefinition production.
	EnterFieldsDefinition(c *FieldsDefinitionContext)

	// EnterFieldDefinition is called when entering the fieldDefinition production.
	EnterFieldDefinition(c *FieldDefinitionContext)

	// EnterArgumentsDefinition is called when entering the argumentsDefinition production.
	EnterArgumentsDefinition(c *ArgumentsDefinitionContext)

	// EnterInputValueDefinition is called when entering the inputValueDefinition production.
	EnterInputValueDefinition(c *InputValueDefinitionContext)

	// EnterInterfaceTypeDefinition is called when entering the interfaceTypeDefinition production.
	EnterInterfaceTypeDefinition(c *InterfaceTypeDefinitionContext)

	// EnterInterfaceTypeExtension is called when entering the interfaceTypeExtension production.
	EnterInterfaceTypeExtension(c *InterfaceTypeExtensionContext)

	// EnterUnionTypeDefinition is called when entering the unionTypeDefinition production.
	EnterUnionTypeDefinition(c *UnionTypeDefinitionContext)

	// EnterUnionMemberTypes is called when entering the unionMemberTypes production.
	EnterUnionMemberTypes(c *UnionMemberTypesContext)

	// EnterUnionTypeExtension is called when entering the unionTypeExtension production.
	EnterUnionTypeExtension(c *UnionTypeExtensionContext)

	// EnterEnumTypeDefinition is called when entering the enumTypeDefinition production.
	EnterEnumTypeDefinition(c *EnumTypeDefinitionContext)

	// EnterEnumValuesDefinition is called when entering the enumValuesDefinition production.
	EnterEnumValuesDefinition(c *EnumValuesDefinitionContext)

	// EnterEnumValueDefinition is called when entering the enumValueDefinition production.
	EnterEnumValueDefinition(c *EnumValueDefinitionContext)

	// EnterEnumTypeExtension is called when entering the enumTypeExtension production.
	EnterEnumTypeExtension(c *EnumTypeExtensionContext)

	// EnterInputObjectTypeDefinition is called when entering the inputObjectTypeDefinition production.
	EnterInputObjectTypeDefinition(c *InputObjectTypeDefinitionContext)

	// EnterInputFieldsDefinition is called when entering the inputFieldsDefinition production.
	EnterInputFieldsDefinition(c *InputFieldsDefinitionContext)

	// EnterInputObjectTypeExtension is called when entering the inputObjectTypeExtension production.
	EnterInputObjectTypeExtension(c *InputObjectTypeExtensionContext)

	// EnterDirectiveDefinition is called when entering the directiveDefinition production.
	EnterDirectiveDefinition(c *DirectiveDefinitionContext)

	// EnterDirectiveLocations is called when entering the directiveLocations production.
	EnterDirectiveLocations(c *DirectiveLocationsContext)

	// ExitGraphql is called when exiting the graphql production.
	ExitGraphql(c *GraphqlContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitExecutableDefinition is called when exiting the executableDefinition production.
	ExitExecutableDefinition(c *ExecutableDefinitionContext)

	// ExitOperationDefinition is called when exiting the operationDefinition production.
	ExitOperationDefinition(c *OperationDefinitionContext)

	// ExitOperationType is called when exiting the operationType production.
	ExitOperationType(c *OperationTypeContext)

	// ExitSelectionSet is called when exiting the selectionSet production.
	ExitSelectionSet(c *SelectionSetContext)

	// ExitSelection is called when exiting the selection production.
	ExitSelection(c *SelectionContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitFragmentSpread is called when exiting the fragmentSpread production.
	ExitFragmentSpread(c *FragmentSpreadContext)

	// ExitFragmentDefinition is called when exiting the fragmentDefinition production.
	ExitFragmentDefinition(c *FragmentDefinitionContext)

	// ExitFragmentName is called when exiting the fragmentName production.
	ExitFragmentName(c *FragmentNameContext)

	// ExitTypeCondition is called when exiting the typeCondition production.
	ExitTypeCondition(c *TypeConditionContext)

	// ExitInlineFragment is called when exiting the inlineFragment production.
	ExitInlineFragment(c *InlineFragmentContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitIntValue is called when exiting the intValue production.
	ExitIntValue(c *IntValueContext)

	// ExitFloatValue is called when exiting the floatValue production.
	ExitFloatValue(c *FloatValueContext)

	// ExitBooleanValue is called when exiting the booleanValue production.
	ExitBooleanValue(c *BooleanValueContext)

	// ExitStringValue is called when exiting the stringValue production.
	ExitStringValue(c *StringValueContext)

	// ExitNullValue is called when exiting the nullValue production.
	ExitNullValue(c *NullValueContext)

	// ExitEnumValue is called when exiting the enumValue production.
	ExitEnumValue(c *EnumValueContext)

	// ExitListValue is called when exiting the listValue production.
	ExitListValue(c *ListValueContext)

	// ExitObjectValue is called when exiting the objectValue production.
	ExitObjectValue(c *ObjectValueContext)

	// ExitObjectField is called when exiting the objectField production.
	ExitObjectField(c *ObjectFieldContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitVariableDefinitions is called when exiting the variableDefinitions production.
	ExitVariableDefinitions(c *VariableDefinitionsContext)

	// ExitVariableDefinition is called when exiting the variableDefinition production.
	ExitVariableDefinition(c *VariableDefinitionContext)

	// ExitDefaultValue is called when exiting the defaultValue production.
	ExitDefaultValue(c *DefaultValueContext)

	// ExitTypeSpec is called when exiting the typeSpec production.
	ExitTypeSpec(c *TypeSpecContext)

	// ExitNamedType is called when exiting the namedType production.
	ExitNamedType(c *NamedTypeContext)

	// ExitListType is called when exiting the listType production.
	ExitListType(c *ListTypeContext)

	// ExitNonNullType is called when exiting the nonNullType production.
	ExitNonNullType(c *NonNullTypeContext)

	// ExitDirectives is called when exiting the directives production.
	ExitDirectives(c *DirectivesContext)

	// ExitDirective is called when exiting the directive production.
	ExitDirective(c *DirectiveContext)

	// ExitTypeSystemDefinition is called when exiting the typeSystemDefinition production.
	ExitTypeSystemDefinition(c *TypeSystemDefinitionContext)

	// ExitSchemaDefinition is called when exiting the schemaDefinition production.
	ExitSchemaDefinition(c *SchemaDefinitionContext)

	// ExitOperationTypeDefinition is called when exiting the operationTypeDefinition production.
	ExitOperationTypeDefinition(c *OperationTypeDefinitionContext)

	// ExitDescription is called when exiting the description production.
	ExitDescription(c *DescriptionContext)

	// ExitTypeDefinition is called when exiting the typeDefinition production.
	ExitTypeDefinition(c *TypeDefinitionContext)

	// ExitTypeExtension is called when exiting the typeExtension production.
	ExitTypeExtension(c *TypeExtensionContext)

	// ExitScalarTypeDefinition is called when exiting the scalarTypeDefinition production.
	ExitScalarTypeDefinition(c *ScalarTypeDefinitionContext)

	// ExitScalarTypeExtension is called when exiting the scalarTypeExtension production.
	ExitScalarTypeExtension(c *ScalarTypeExtensionContext)

	// ExitObjectTypeDefinition is called when exiting the objectTypeDefinition production.
	ExitObjectTypeDefinition(c *ObjectTypeDefinitionContext)

	// ExitObjectTypeExtension is called when exiting the objectTypeExtension production.
	ExitObjectTypeExtension(c *ObjectTypeExtensionContext)

	// ExitImplementsInterfaces is called when exiting the implementsInterfaces production.
	ExitImplementsInterfaces(c *ImplementsInterfacesContext)

	// ExitFieldsDefinition is called when exiting the fieldsDefinition production.
	ExitFieldsDefinition(c *FieldsDefinitionContext)

	// ExitFieldDefinition is called when exiting the fieldDefinition production.
	ExitFieldDefinition(c *FieldDefinitionContext)

	// ExitArgumentsDefinition is called when exiting the argumentsDefinition production.
	ExitArgumentsDefinition(c *ArgumentsDefinitionContext)

	// ExitInputValueDefinition is called when exiting the inputValueDefinition production.
	ExitInputValueDefinition(c *InputValueDefinitionContext)

	// ExitInterfaceTypeDefinition is called when exiting the interfaceTypeDefinition production.
	ExitInterfaceTypeDefinition(c *InterfaceTypeDefinitionContext)

	// ExitInterfaceTypeExtension is called when exiting the interfaceTypeExtension production.
	ExitInterfaceTypeExtension(c *InterfaceTypeExtensionContext)

	// ExitUnionTypeDefinition is called when exiting the unionTypeDefinition production.
	ExitUnionTypeDefinition(c *UnionTypeDefinitionContext)

	// ExitUnionMemberTypes is called when exiting the unionMemberTypes production.
	ExitUnionMemberTypes(c *UnionMemberTypesContext)

	// ExitUnionTypeExtension is called when exiting the unionTypeExtension production.
	ExitUnionTypeExtension(c *UnionTypeExtensionContext)

	// ExitEnumTypeDefinition is called when exiting the enumTypeDefinition production.
	ExitEnumTypeDefinition(c *EnumTypeDefinitionContext)

	// ExitEnumValuesDefinition is called when exiting the enumValuesDefinition production.
	ExitEnumValuesDefinition(c *EnumValuesDefinitionContext)

	// ExitEnumValueDefinition is called when exiting the enumValueDefinition production.
	ExitEnumValueDefinition(c *EnumValueDefinitionContext)

	// ExitEnumTypeExtension is called when exiting the enumTypeExtension production.
	ExitEnumTypeExtension(c *EnumTypeExtensionContext)

	// ExitInputObjectTypeDefinition is called when exiting the inputObjectTypeDefinition production.
	ExitInputObjectTypeDefinition(c *InputObjectTypeDefinitionContext)

	// ExitInputFieldsDefinition is called when exiting the inputFieldsDefinition production.
	ExitInputFieldsDefinition(c *InputFieldsDefinitionContext)

	// ExitInputObjectTypeExtension is called when exiting the inputObjectTypeExtension production.
	ExitInputObjectTypeExtension(c *InputObjectTypeExtensionContext)

	// ExitDirectiveDefinition is called when exiting the directiveDefinition production.
	ExitDirectiveDefinition(c *DirectiveDefinitionContext)

	// ExitDirectiveLocations is called when exiting the directiveLocations production.
	ExitDirectiveLocations(c *DirectiveLocationsContext)
}
