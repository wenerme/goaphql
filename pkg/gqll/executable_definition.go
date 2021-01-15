package gqll

type OperationDefinition struct {
	baseNode
	hasName
	hasDirectives
	hasSelectionSet
	hasVariableDefinitions
	OperationType string
}

type FragmentDefinition struct {
	baseNode
	hasDirectives
	hasSelectionSet
	FragmentName  string
	ConditionType string
}

type SelectionSet struct {
	Fields          []*Field
	FragmentSpreads []*FragmentSpread
	InlineFragments []*InlineFragment
}

type Field struct {
	baseNode
	hasName
	hasArguments
	hasDirectives
	hasSelectionSet
	hasVariableDefinitions
	Alias string
}

type VariableDefinition struct {
	baseNode
	hasType
	hasDefaultValue
	hasDirectives
	VariableName string
}

type FragmentSpread struct {
	baseNode
	hasDirectives
	FragmentName string
}

type InlineFragment struct {
	baseNode
	hasDirectives
	hasSelectionSet
	ConditionType string
}
