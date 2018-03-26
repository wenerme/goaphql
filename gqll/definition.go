package gqll

/*
public interface Node {

  List<Node> getChildren();

  SourceLocation getSourceLocation();

  List<Comment> getComments();

  default String getNodeType() {
    return getClass().getSimpleName();
  }
}

*/

type Node interface {
	GetSourceLocation() *SourceLocation
	SetSourceLocation(location *SourceLocation)

	GetComments() []Comment
	SetComments(comments []Comment)

	//IsTypeDefinition() bool
	//IsTypeExtension() bool
	//IsType() bool
	//IsValue() bool
}

type SourceLocation struct {
	Line   int
	Column int
	Source string
}

type Comment struct {
	Content string
	baseNode
}

type Type interface {
	Node
	IsNonNull() bool
	IsNamed() bool
	IsList() bool

	GetType() Type
	GetName() string
}
type Value interface {
	Node
	GetValue() interface{}
	// TODO Add resolved value
	//IsVariable() bool
}

type Document struct {
	baseNode
	Definitions []Definition
}

type Argument struct {
	baseNode
	hasName
	hasValue
}

type Directive struct {
	baseNode
	hasName
	hasArguments
}

type Definition interface {
	Node
}

type FieldDefinition struct {
	baseNode
	hasName
	hasType
	hasDescription
	hasDirectives
	hasArgumentDefinitions
}

type InputValueDefinition struct {
	baseNode
	hasName
	hasDescription
	hasDirectives
	hasType
	hasDefaultValue
}

type EnumValueDefinition struct {
	baseNode
	hasDescription
	hasDirectives
	Value string // Enum name
}
