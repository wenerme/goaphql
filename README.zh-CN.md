# GoaphQL

其他语言 | [中文](./README.zh-CN.md) | [English](./README.md)

## 命令行

```bash
# 安装
go get -u -v github.com/wenerme/goaphql/cmd/goaphql
```


### 生成器

```bash
# 生成
# 包名 com.github
# Schema 名 Github
goaphql generate graphql-java -s github.graphqls -p com.github -n GitHub
# 缩写
goaphql g gj -s github.graphqls -p com.github -n GitHub

# 模板定制化
# 导出模板
goaphql g dump -t tmpl/graphql-java --prefix tmpl/graphql-java -d ./tpl
# 修改模板
# 添加 package-info.java
echo -e 'package {{Config.JavaPackage}};' > './tpl/Java#package-info.java.tmpl'
# 使用修改后的目标进行生成
goaphql generate graphql-java -t ./tpl -s github.graphqls -p com.github -n GitHub
# 查看新添加模板生成的文件
cat out/com/github/package-info.java
```

## 内容

包 | 描述
--------|------------
gqlp    | 解析器
gqll    | 语言结构
gqlg    | 代码生成器
cmd/goaphql | 命令行工具

## GraphQL 语言扩展

* Antlr Grammar [GraphQL.g4](https://github.com/wenerme/wener/blob/master/tricks/languages/antlr/GraphQL.g4)

```graphqls
# 1. extend by name 语法

type MyQuery {
  myUser(id:ID!):MyUser
}

extend Query by MyQuery

# 2. 允许在定义指令时添加指令, 添加 DIRECTIVE_DEFINITION 指令位置

directive @JavaType(type:String) on DIRECTIVE_DEFINITION
directive @Auth(value:String) @JavaType(type:"Auth") on FIELD_DEFINITION;

# 3. 允许 schema 指定名字
schema Test {
  query: MyQuery
}
```
