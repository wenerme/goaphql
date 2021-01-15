// +build generate

//go:generate antlr -Dlanguage=Go -visitor -package gqlp GraphQL.g4
package gqlp
