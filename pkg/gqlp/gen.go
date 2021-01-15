// +build generate

//go:generate antlr4 -Dlanguage=Go -visitor -package gqlp GraphQL.g4
package gqlp
