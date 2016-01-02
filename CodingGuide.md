**Table of Contents**


## Tabs & Indenting ##
Tab characters (\0x09) should always be used in code to increase the readability. Tabs may be replaced by 4 space characters when tabs are not allowed by the editor.

## Bracing ##
Open braces should always be on the line that contains the statement that begins the block. Contents of the brace should be indented by a tab. For example:
```
func fooBar(a int) int {
    if a != 0 {
        return 1
    }
}
```

"case" statements should be indented from the switch statement like this:
```
switch foobar {
    case 0:
        something()

    case 1:
        somethingElse()
}
```

It is not necesary to include a "break" statement after a case.

Braces are never optional. Even for single statement blocks, braces should be used.

## Parenthesis ##
It is not necesary to add parenthesis to a short if statement. However, for longer if statements (e.g. with several comparisons) it is necesary to add parenthesis. For example:
```
if foo != nil {
}

if (foo != nil) && (bar != nil) {
}
```

## Commenting ##
Comments should be used to describe intention, algorithmic overview, and/or logical flow.  It would be ideal, if from reading the comments alone, someone other than the author could understand a function’s intended behavior and general operation. While there are no minimum comment requirements and certainly some very small routines need no commenting at all, it is hoped that most routines will have comments reflecting the programmer’s intent and approach.

The // (two slashes) style of comment tags should be used in most situations. Where ever possible, place comments above the code instead of beside it.  For example:

```
// Create new instance of FooBar
func NewFooBar(a int) *FooBar {
    foobar := &FooBar{}
    
    // Assign "a" to the new instance
    foobar.a = a
    
    return foobar
}
```

## Spacing ##
Use a single space between function arguments, but not after the last argument or after the opening parenthesis.
```
// Right
fooBar(arg0 int, arg1 int)
// Wrong
fooBar(arg0 int,arg1 int)
// Wrong
fooBar( arg0 int, arg1 int )
```
Use a space between comparison operators
```
// Right
if (a == b) {
// Wrong
if (a==b) {
```
Don't use a space between a function name and parenthesis
```
// Right
fooBar()
// Wrong
fooBar ()
```
Don't use spaces between brackets
```
// Right
array[index]
/// Wrong
array[ index ]
```

## Naming ##
Function arguments need to have an `_` (underscore) in front of them. For example:
```
func foo(_bar int) {
}
```

Global variables need to have g`_` in front of them. For example:
```
g_globalfoo *FooBar = NewFooBar()
```

As specified by Go, functions that are allowed to be called from outside a package should start with a capital letter. Other functions should start with a lowercase letter. For example:
```
func (f *FooBar) Open(_path string) {
    f.internalOpen(_path)
}

func (f *FooBar) internalOpen(_path string) {
}
```

Always use camelCasing for member variables, arguments, parameters and local variables.

Always prefix interfaces with "I". Example:
```
type IFooBar interface {
}
```

Constants should be written in ALL\_CAPS. For example:
```
const (
    FOO = 1
    BAR = 2
)
```

## Declarations ##
Arguments should always specify their type, even when it could be ommited. For example:
```
// Right
func foo(_bar1 int, _bar2 int)
// Wrong
func foo(_bar1, _bar2 int)
```

Member variables do not need to specify their type when it can be ommited. For example:
```
type FooBar struct {
    //Allowed
    bar1, bar2 int
}
```

Local variables may be declared in two ways:

- var statement including type
```
var foo int = 0
```
- :`=` operator
```
foo := 0
```
Using the var statement without specifying the type is not allowed:
```
// Wrong
var foo = 0
```

Functions that return an instance of a type should always have the New prefix followed by the name of the type. For example:
```
type Foo struct {
}

func NewFoo() *Foo {
    return &Foo{}
}
```

## File naming ##
All source (.go)  files should be all lowercase. If the name is very long (try to avoid this!), it is allowed to use `_` (underscore) in the name to increase readability.