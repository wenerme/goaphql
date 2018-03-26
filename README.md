# GoaphQL

Read in | [中文](./README.zh-CN.md) | [English](./README.md)


## Command Line

```bash
# Installation
go get -u -v github.com/wenerme/goaphql/cmd/goaphql
```

## graphql-java

```bash
# Generate
# Package com.github
# Name Github
goaphql generate graphql-java -s github.graphqls -P com.github -N GitHub
# same
goaphql g gj -s github.graphqls -P com.github -N GitHub

# Customization
# Dump templates
goaphql dump -t tmpl/graphql-java --prefix tmpl/graphql-java -d ./tpl
# Modify templats
# Add package-info.java
echo -e 'package {{Config.JavaPackage}};' > './tpl/Java#package-info.java.tmpl'
# Generate using modified templates
goaphql generate graphql-java -t ./tpl -s github.graphqls -P com.github -N GitHub
# Check the new file
cat out/me/wener/package-info.java
```

### Dependencies
Generated depends on [wenerme/jraphql](https://github.com/wenerme/jraphql) .

```xml
<dependencies>
  <dependency>
      <groupId>me.wener.jraphql</groupId>
      <artifactId>jraphql-graphql-java-adapter</artifactId>
      <version>0.0.2</version>
  </dependency>
  <dependency>
      <groupId>org.projectlombok</groupId>
      <artifactId>lombok</artifactId>
      <version>1.16.20</version>
      <scope>provided</scope>
  </dependency>
</dependencies>
```

## Inside

Package | Description
--------|------------
gqlp    | GraphQL Parser
gqll    | GraphQL Language representation
gqlg    | GraphQL Code generator
cmd/goaphql | Command line tool

## GraphQL Language Extension

* Antlr Grammar [GraphQL.g4](https://github.com/wenerme/wener/blob/master/tricks/languages/antlr/GraphQL.g4)

```graphqls
# 1. extend by name syntax

type MyQuery {
  myUser(id:ID!):MyUser
}

extend Query by MyQuery

# 2. Allowed directives on directive definition, add DIRECTIVE_DEFINITION location

directive @JavaType(type:String) on DIRECTIVE_DEFINITION
directive @Auth(value:String) @JavaType(type:"Auth") on FIELD_DEFINITION;

# 3. Allowed schema has optional name
schema Test {
  query: MyQuery
}

```

## Roadmap

* [ ] Parse comment
* [ ] Extract java runtime to [jraphql](https://github.com/wenerme/jraphql)
* [ ] Implement golang runtime
* [ ] Implement golang codegen
* [ ] Implement graphql codegen

