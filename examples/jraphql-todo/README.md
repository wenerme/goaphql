# JraphQL ToDo

```bash
# Generate common
goaphql g gj -s common.graphqls -P me.wener.todo.spec.common -N Common  -d todo/src/main/java

# Generate todo
goaphql g gj -s todo.graphqls -P me.wener.todo.spec.todo -N Todo  -d todo/src/main/java -i common.graphqls 
```
