package gqlp

import (
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sirupsen/logrus"
	"github.com/wenerme/goaphql/pkg/gqll"
)

var _ GraphQLVisitor = &GraphQLLangVisitor{}

func NewGraphQLLangVisitor() *GraphQLLangVisitor {
	return &GraphQLLangVisitor{
		ParseTreeVisitor: &antlr.BaseParseTreeVisitor{},
	}
}

type GraphQLLangVisitor struct {
	antlr.ParseTreeVisitor
	Source string
}

func (v *GraphQLLangVisitor) VisitGraphql(ctx *GraphqlContext) interface{} {
	return v.Visit(ctx.Document())
}

func (v *GraphQLLangVisitor) VisitName(ctx *NameContext) interface{} {
	return ctx.GetText()
}

func (v *GraphQLLangVisitor) VisitDocument(ctx *DocumentContext) interface{} {
	node := new(gqll.Document)
	v.Extract(ctx.BaseParserRuleContext, node)
	for _, context := range ctx.AllDefinition() {
		node.Definitions = append(node.Definitions, v.Visit(context).(gqll.Definition))
	}
	return node
}

func (v *GraphQLLangVisitor) VisitDefinition(ctx *DefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitExecutableDefinition(ctx *ExecutableDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitOperationDefinition(ctx *OperationDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitOperationType(ctx *OperationTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitSelectionSet(ctx *SelectionSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitSelection(ctx *SelectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	var arr []*gqll.Argument
	for _, context := range ctx.AllArgument() {
		arr = append(arr, v.VisitArgument(context.(*ArgumentContext)).(*gqll.Argument))
	}
	return arr
}

func (v *GraphQLLangVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, new(gqll.Argument))
}

func (v *GraphQLLangVisitor) VisitAlias(ctx *AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitFragmentSpread(ctx *FragmentSpreadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitFragmentDefinition(ctx *FragmentDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitFragmentName(ctx *FragmentNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitTypeCondition(ctx *TypeConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitInlineFragment(ctx *InlineFragmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitIntValue(ctx *IntValueContext) interface{} {
	val, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		panic(err) // FIXME
	}
	return v.Extract(ctx.BaseParserRuleContext, &gqll.IntValue{Value: val})
}

func (v *GraphQLLangVisitor) VisitFloatValue(ctx *FloatValueContext) interface{} {
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		panic(err) // FIXME
	}
	return v.Extract(ctx.BaseParserRuleContext, &gqll.FloatValue{Value: val})
}

func (v *GraphQLLangVisitor) VisitBooleanValue(ctx *BooleanValueContext) interface{} {
	val, err := strconv.ParseBool(ctx.GetText())
	if err != nil {
		panic(err) // FIXME
	}
	return v.Extract(ctx.BaseParserRuleContext, &gqll.BooleanValue{Value: val})
}

func (v *GraphQLLangVisitor) VisitStringValue(ctx *StringValueContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, &gqll.StringValue{Value: v.extractString(ctx)})
}

func (v *GraphQLLangVisitor) VisitNullValue(ctx *NullValueContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, &gqll.NullValue{})
}

func (v *GraphQLLangVisitor) VisitEnumValue(ctx *EnumValueContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, &gqll.EnumValue{Value: v.extractText(ctx.Name())})
}

func (v *GraphQLLangVisitor) VisitListValue(ctx *ListValueContext) interface{} {
	// TODO

	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitObjectValue(ctx *ObjectValueContext) interface{} {
	// TODO
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitObjectField(ctx *ObjectFieldContext) interface{} {
	// TODO
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, &gqll.Variable{Name: ctx.Name().GetText()})
}

func (v *GraphQLLangVisitor) VisitVariableDefinitions(ctx *VariableDefinitionsContext) interface{} {
	var arr []*gqll.VariableDefinition
	for _, context := range ctx.AllVariableDefinition() {
		arr = append(arr, v.VisitVariableDefinition(context.(*VariableDefinitionContext)).(*gqll.VariableDefinition))
	}
	return arr
}

func (v *GraphQLLangVisitor) VisitVariableDefinition(ctx *VariableDefinitionContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, &gqll.VariableDefinition{VariableName: ctx.Variable().GetText()})
}

func (v *GraphQLLangVisitor) VisitDefaultValue(ctx *DefaultValueContext) interface{} {
	return v.Visit(ctx.Value())
}

func (v *GraphQLLangVisitor) VisitTypeSpec(ctx *TypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitDirectives(ctx *DirectivesContext) interface{} {
	var arr []*gqll.Directive
	for _, context := range ctx.AllDirective() {
		arr = append(arr, v.VisitDirective(context.(*DirectiveContext)).(*gqll.Directive))
	}
	return arr
}

func (v *GraphQLLangVisitor) VisitDirective(ctx *DirectiveContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, new(gqll.Directive))
}

func (v *GraphQLLangVisitor) VisitTypeSystemDefinition(ctx *TypeSystemDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitSchemaDefinition(ctx *SchemaDefinitionContext) interface{} {
	node := new(gqll.SchemaDefinition)
	v.Extract(ctx.BaseParserRuleContext, node)
	for _, op := range ctx.AllOperationTypeDefinition() {
		c := op.(*OperationTypeDefinitionContext)
		typeName := c.NamedType().GetText()
		switch c.OperationType().GetText() {
		case "query":
			node.QueryTypeName = typeName
		case "mutation":
			node.MutationTypeName = typeName
		case "subscription":
			node.SubscriptionTypeName = typeName
		default:
			panic("Unexpected schema operation")
		}
	}
	return node
}

func (v *GraphQLLangVisitor) VisitOperationTypeDefinition(ctx *OperationTypeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitDescription(ctx *DescriptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitTypeDefinition(ctx *TypeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitTypeExtension(ctx *TypeExtensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *GraphQLLangVisitor) VisitScalarTypeDefinition(ctx *ScalarTypeDefinitionContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, new(gqll.ScalarTypeDefinition))
}

func (v *GraphQLLangVisitor) VisitScalarTypeExtension(ctx *ScalarTypeExtensionContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, new(gqll.ScalarTypeExtension))
}

func (v *GraphQLLangVisitor) VisitObjectTypeDefinition(ctx *ObjectTypeDefinitionContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, new(gqll.ObjectTypeDefinition))
}

func (v *GraphQLLangVisitor) VisitObjectTypeExtension(ctx *ObjectTypeExtensionContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, new(gqll.ObjectTypeExtension))
}

func (v *GraphQLLangVisitor) VisitImplementsInterfaces(ctx *ImplementsInterfacesContext) interface{} {
	var arr []string
	c := ctx
	for c != nil {
		arr = append(arr, c.NamedType().GetText())

		if c.ImplementsInterfaces() != nil {
			c = c.ImplementsInterfaces().(*ImplementsInterfacesContext)
		} else {
			c = nil
		}
	}
	return arr
}

func (v *GraphQLLangVisitor) VisitFieldsDefinition(ctx *FieldsDefinitionContext) interface{} {
	var arr []*gqll.FieldDefinition
	for _, context := range ctx.AllFieldDefinition() {
		arr = append(arr, v.VisitFieldDefinition(context.(*FieldDefinitionContext)).(*gqll.FieldDefinition))
	}
	return arr
}

func (v *GraphQLLangVisitor) VisitFieldDefinition(ctx *FieldDefinitionContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, new(gqll.FieldDefinition))
}

func (v *GraphQLLangVisitor) VisitArgumentsDefinition(ctx *ArgumentsDefinitionContext) interface{} {
	var arr []*gqll.InputValueDefinition
	for _, context := range ctx.AllInputValueDefinition() {
		arr = append(arr, v.VisitInputValueDefinition(context.(*InputValueDefinitionContext)).(*gqll.InputValueDefinition))
	}
	return arr
}

func (v *GraphQLLangVisitor) VisitInputValueDefinition(ctx *InputValueDefinitionContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, new(gqll.InputValueDefinition))
}

func (v *GraphQLLangVisitor) VisitInterfaceTypeDefinition(ctx *InterfaceTypeDefinitionContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, new(gqll.InterfaceTypeDefinition))
}

func (v *GraphQLLangVisitor) VisitInterfaceTypeExtension(ctx *InterfaceTypeExtensionContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, new(gqll.InterfaceTypeExtension))
}

func (v *GraphQLLangVisitor) VisitUnionTypeDefinition(ctx *UnionTypeDefinitionContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, new(gqll.UnionTypeDefinition))
}

func (v *GraphQLLangVisitor) VisitUnionMemberTypes(ctx *UnionMemberTypesContext) interface{} {
	var arr []string
	c := ctx
	for c != nil {
		arr = append(arr, c.NamedType().GetText())

		if c.UnionMemberTypes() != nil {
			c = c.UnionMemberTypes().(*UnionMemberTypesContext)
		} else {
			c = nil
		}
	}
	return arr
}

func (v *GraphQLLangVisitor) VisitUnionTypeExtension(ctx *UnionTypeExtensionContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, new(gqll.UnionTypeExtension))
}

func (v *GraphQLLangVisitor) VisitEnumTypeDefinition(ctx *EnumTypeDefinitionContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, new(gqll.EnumTypeDefinition))
}

func (v *GraphQLLangVisitor) VisitEnumValuesDefinition(ctx *EnumValuesDefinitionContext) interface{} {
	var arr []*gqll.EnumValueDefinition
	for _, context := range ctx.AllEnumValueDefinition() {
		arr = append(arr, v.VisitEnumValueDefinition(context.(*EnumValueDefinitionContext)).(*gqll.EnumValueDefinition))
	}
	return arr
}

func (v *GraphQLLangVisitor) VisitEnumValueDefinition(ctx *EnumValueDefinitionContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, &gqll.EnumValueDefinition{Value: ctx.EnumValue().GetText()})
}

func (v *GraphQLLangVisitor) VisitEnumTypeExtension(ctx *EnumTypeExtensionContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, new(gqll.EnumTypeExtension))
}

func (v *GraphQLLangVisitor) VisitInputObjectTypeDefinition(ctx *InputObjectTypeDefinitionContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, new(gqll.InputObjectTypeDefinition))
}

func (v *GraphQLLangVisitor) VisitInputFieldsDefinition(ctx *InputFieldsDefinitionContext) interface{} {
	var arr []*gqll.InputValueDefinition
	for _, context := range ctx.AllInputValueDefinition() {
		arr = append(arr, v.VisitInputValueDefinition(context.(*InputValueDefinitionContext)).(*gqll.InputValueDefinition))
	}
	return arr
}

func (v *GraphQLLangVisitor) VisitInputObjectTypeExtension(ctx *InputObjectTypeExtensionContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, new(gqll.InputObjectTypeExtension))
}

func (v *GraphQLLangVisitor) VisitDirectiveDefinition(ctx *DirectiveDefinitionContext) interface{} {
	node := new(gqll.DirectiveDefinition)
	v.Extract(ctx.BaseParserRuleContext, node)
	val := v.VisitDirectiveLocations(ctx.DirectiveLocations().(*DirectiveLocationsContext))
	if val != nil {
		node.Locations = val.([]string)
	}
	return node
}

func (v *GraphQLLangVisitor) VisitDirectiveLocations(ctx *DirectiveLocationsContext) interface{} {
	var arr []string
	c := ctx
	for c != nil {
		arr = append(arr, c.DirectiveLocation().GetText())

		if c.DirectiveLocations() != nil {
			c = c.DirectiveLocations().(*DirectiveLocationsContext)
		} else {
			c = nil
		}
	}
	return arr
}

func (v *GraphQLLangVisitor) Visit(tree antlr.ParseTree) interface{} {
	//logrus.Debug("Visit", reflect.TypeOf(tree))
	if tree == nil {
		return nil
	}
	return tree.Accept(v)
}
func (v *GraphQLLangVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return nil
}

func (v *GraphQLLangVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	//logrus.Debug("VisitChildren", reflect.TypeOf(node))

	var result interface{}
	n := node.GetChildCount()
	for i := 0; i < n; i++ {
		c := node.GetChild(i).(antlr.ParseTree)
		childResult := c.Accept(v)
		//logrus.Debug("VisitChild ", i, reflect.TypeOf(c))

		result = childResult
	}
	return result
}
func (v *GraphQLLangVisitor) VisitNonNullType(ctx *NonNullTypeContext) interface{} {
	node := &gqll.NonNullType{}
	v.Extract(ctx.BaseParserRuleContext, node)
	if ctx.NamedType() != nil {
		node.Type = v.VisitNamedType(ctx.NamedType().(*NamedTypeContext)).(gqll.Type)
	} else {
		node.Type = v.VisitListType(ctx.ListType().(*ListTypeContext)).(gqll.Type)
	}
	return node
}

func (v *GraphQLLangVisitor) VisitListType(ctx *ListTypeContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, &gqll.ListType{})
}
func (v *GraphQLLangVisitor) VisitNamedType(ctx *NamedTypeContext) interface{} {
	return v.Extract(ctx.BaseParserRuleContext, &gqll.NamedType{})
}

func (v *GraphQLLangVisitor) Extract(ctx *antlr.BaseParserRuleContext, node gqll.Node) interface{} {
	if node == nil || ctx == nil {
		return node
	}
	node.SetSourceLocation(&gqll.SourceLocation{
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
		Source: v.Source,
	})
	// TODO comments

	if node, ok := node.(interface {
		SetArguments([]*gqll.Argument)
	}); ok {
		context := ctx.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)
		val := v.Visit(context)
		if val != nil {
			node.SetArguments(val.([]*gqll.Argument))
		}
	}
	if node, ok := node.(interface {
		SetArgumentDefinitions([]*gqll.InputValueDefinition)
	}); ok {
		context := ctx.GetTypedRuleContext(reflect.TypeOf((*IArgumentsDefinitionContext)(nil)).Elem(), 0)
		val := v.Visit(context)
		if val != nil {
			node.SetArgumentDefinitions(val.([]*gqll.InputValueDefinition))
		}
	}
	if node, ok := node.(interface {
		SetInputFieldDefinitions([]*gqll.InputValueDefinition)
	}); ok {
		context := ctx.GetTypedRuleContext(reflect.TypeOf((*IInputFieldsDefinitionContext)(nil)).Elem(), 0)
		val := v.Visit(context)
		if val != nil {
			node.SetInputFieldDefinitions(val.([]*gqll.InputValueDefinition))
		}
	}
	if node, ok := node.(interface {
		SetDefaultValue(gqll.Value)
	}); ok {
		context := ctx.GetTypedRuleContext(reflect.TypeOf((*IDefaultValueContext)(nil)).Elem(), 0)
		val := v.Visit(context)
		if val != nil {
			node.SetDefaultValue(val.(gqll.Value))
		}
	}
	if node, ok := node.(interface {
		SetDescription(string)
	}); ok {
		context := ctx.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)
		node.SetDescription(v.extractString(context))
	}
	if node, ok := node.(interface {
		SetDirectives([]*gqll.Directive)
	}); ok {
		context := ctx.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)
		val := v.Visit(context)
		if val != nil {
			node.SetDirectives(val.([]*gqll.Directive))
		}
	}
	if node, ok := node.(interface {
		SetEnumValueDefinitions(v []*gqll.EnumValueDefinition)
	}); ok {
		context := ctx.GetTypedRuleContext(reflect.TypeOf((*IEnumValuesDefinitionContext)(nil)).Elem(), 0)
		val := v.Visit(context)
		if val != nil {
			node.SetEnumValueDefinitions(val.([]*gqll.EnumValueDefinition))
		}
	}

	if node, ok := node.(interface {
		SetFieldDefinitions([]*gqll.FieldDefinition)
	}); ok {
		context := ctx.GetTypedRuleContext(reflect.TypeOf((*IFieldsDefinitionContext)(nil)).Elem(), 0)
		val := v.Visit(context)
		if val != nil {
			node.SetFieldDefinitions(val.([]*gqll.FieldDefinition))
		}
	}
	if node, ok := node.(interface {
		SetInterfaces([]string)
	}); ok {
		context := ctx.GetTypedRuleContext(reflect.TypeOf((*IImplementsInterfacesContext)(nil)).Elem(), 0)
		val := v.Visit(context)
		if val != nil {
			node.SetInterfaces(val.([]string))
		}
	}
	if node, ok := node.(gqll.HasName); ok {
		context := ctx.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)
		node.SetName(v.extractText(context))
	}
	// Must after name
	if node, ok := node.(gqll.HasExtendTypeName); ok {
		context := ctx.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 1)
		extendName := v.extractText(context)

		// This is the name of the extension
		nameNode := node.(gqll.HasName)
		logrus.Info("Extension ", nameNode.GetName())
		// implicit name of this extension
		if extendName == "" {
			extendName = nameNode.GetName() + "Extension"
		}
		node.SetExtendTypeName(nameNode.GetName())
		nameNode.SetName(extendName)

	}
	if node, ok := node.(interface {
		SetType(gqll.Type)
	}); ok {
		v := v.Visit(ctx.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0))
		if v != nil {
			node.SetType(v.(gqll.Type))
		}
	}
	if node, ok := node.(interface {
		SetUnionMemberTypes(v []string)
	}); ok {
		v := v.Visit(ctx.GetTypedRuleContext(reflect.TypeOf((*IUnionMemberTypesContext)(nil)).Elem(), 0))
		if v != nil {
			node.SetUnionMemberTypes(v.([]string))
		}
	}
	if node, ok := node.(interface {
		SetValue(gqll.Value)
	}); ok {
		v := v.Visit(ctx.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0))
		if v != nil {
			node.SetValue(v.(gqll.Value))
		}
	}
	return node
}
func (v *GraphQLLangVisitor) extractText(iface interface{}) string {
	if iface == nil {
		return ""
	}
	ctx := iface.(antlr.ParseTree)
	return ctx.GetText()
}
func (v *GraphQLLangVisitor) extractString(iface interface{}) string {
	if iface == nil {
		return ""
	}
	ctx := iface.(antlr.ParseTree)
	// TODO Escape
	text := ctx.GetText()
	return text[1 : len(text)-1]
}
