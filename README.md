# TODO CLI APP

Simple Todo management app using json file for persistent storage

---

### Supported Commands and Subcommands
1. add
    - -title
2. edit
    - -index
    - -title
3. toggle
    - -index
4. delete
    - -index
5. list

---
## Usage:

```
go build -o todo .

./todo <command> -<subcommand1> val1 -<subcommand2> val2
```

1. To add a new todo.

```
./todo add -title "Todo 1"
```

2. To edit a todo.

```
./todo edit -index 1 -title "Todo 1 updated"
```

3. To toggle completion status of todo.

```
./todo toggle -index 1
```

4. To delete a todo.

```
./todo delete -index 1
```

5. To list all todos.

```
./todo list
```
