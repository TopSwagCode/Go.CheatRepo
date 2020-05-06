# Basics

There is a bunch of ways to work with variables in Go. Like:

* ```var i int```
* ```var i int = 1```
* ```var i = 1```
* ```i := 1``` (_:= syntax is shorthand for declaring and initializing a variable_)
* ```const i int = 1``` (_const is the keyword for a constant that cannot change_)
* ```const i = 1```

Similar to these you are also able to declare multiple variables in one line like:

* ```var i,j int``` (_Multiple of same type_)
* ```var i, isGoAwesome, myName = 1, true, "Joshua"``` (_Multiple different types_)

You can also use the package ```fmt``` to format and/or print text.

```Go

fmt.Println("Hello World")

```
