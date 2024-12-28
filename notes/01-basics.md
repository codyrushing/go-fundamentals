* Can compile to WASM or machine code.
* We don't create classes. It's not object-oriented
* Functions are first class
* No exceptions, no try/catch. Errors are just values

## Modules
* `go mod init example.com/path` on an existing dir
* creates `go.mod` in that directory, with a module name and specified node version. kinda like package.json
* the entry point for that module is any file that has the following (usually main.go):
```go
package main

func main() {
    print("Hello from module")
}
```

## Workspaces
* A multi-module app concept introduced in v1.18
* Workspace
    -> Module(s)
        -> Package(s)

## Packages
Packages are one or more files in the same folder with the same `package package_name` as the first line. The advantage of using multiple files within the same package is that they implicitly share "globals" and don't need to import them. You can think of it as all files in a package get merged together by the compiler. it feels like it's kind of a bad idea to use multiple files, should probably keep every package to one file ideally

You do not explicitly export variables from a package, that's defined through a naming convention.  Capital-case members get exported (`func MyFunc` will be made available to importers, `func myFunc` will be private to the package)

When you import a package, the last segment of the import path becomes the name of the pseudo-object that gets created for that import by default:

```go
import "example/data"

func main(){
    print(data.SomeValue)
}
```

You can also create an alias for the imported package

```go
import dataLib "example/data"

func main(){
    print(dataLib.SomeValue)
}
```

## Variables
* `var` and `const`. `var` can create variables without values, in which case they will be `nil` by default (unlike undefined in js)
* `const` has to have a value, and can only be bool, string, or number
    * it's not like a JS const, which is just an allocation of memory for a variable, and at runtime it prevents reassignment of that variable
    * `const` in go is like the actual value, not a location in memory. it has to be set at compilation time. you won't use them as often as js const
* Always use double quotes for strings
* var shortcut:
```go
someVar := "hi"
// same as
var text string
text = "hi"
```
* types can be inferred when being assigned in the declaration
* vars are block scoped by default, if you make them inside an if block for example

## data types
* int, uint, float, all with different bits (eg. int32, float32)

## collections
* Arrays are fixed in size
* You declare the length and type when initializing an array:

```go
// declare a variable that is of size 10 and uses strings
var Countries [10]string
```

There are also slices, which can be appended to, but returns a new slice, does not update in-place

```go
names := []string { "Alice", "Bob" }
names := append(names, "Charlie")
```

* key-values types are maps

```go
// declare a map with string key and boolean values
var Codes map[string]bool
```

## Functions
* Functions can return multiple values.
* By default, function args are always values, not references.  So if a caller passes in a variable, but default it will be a clone of that value and the original will not be mutated. It will always be passed a clone of that variable.

```go
// syntax like so
func save(text string) {}

// return types like so
func add(a int, b int) int {
    return a + b
}

// multiple returns like so. like a python tuple
func addAndSub(a int, b int) (int, int) {
    return a + b, a - b
}
```

You can also provide names to multiple return types.  Functions can not be overloaded, except the function name `init` which is special

### Pointers
* `*someRef` gets the value of some reference
* `&someVar` gets a reference for some variable

### Panic/Defer
* `panic(msg)` will close the app, not used for regular error management
* `defer someFunc()` will call someFunc after its containing function finishes execution. Useful for cleaning up a function. Doesn't matter where it is called in the function, it will execute after the function completes

### Error design pattern
A common convention in go is when you have a possibly errorable function, you return two values from the function, the first being the success return value, and the second being an error. You then check the error

### Control structure
```go
// here we are creating a var as part of a conditional, 
// the var message is only available inside both the if and else clauses, but not outside 
if message:"hello"; user <> nil {

} else {

}

// switches don't need breaks, and you can fallthrough by specifically calling for it
switch day {
    case "Monday":
        fmt.Println("It's monday")
    case "Tuesday":
        fallthrough
    case "Wednesday":
        fallthrough
    case "Thursday":
        fmt.Println("Boo middle of the week")
    default:        
}

// switches can also have arbitrary conditions for each case, much better than multiple if/elseif/elseif/else
switch {
    case user = nil:
    case user.active = false:
    case user.deleted = true:
}

// loops
// classic
for i:=0; i<len(some_collection); i++ {

}

// for range on a collection gets the index
for index := range some_collection {
    item := some_collection[index]
}

// for range on a map gets the key and the value, which is nice
for key, value := range map {

}

// for can also accept a boolean exp, which basically a while
endOfGame := false
for !endOfGame {
    // do something    
    endOfGame = true
}

// with no expression, it runs forever
for {
    // this will run for ever
}
```

## Type declarations
You can create new types based on other types
```go
// type alias
type distance = float32

// a new type based on another type
type distance float32

// types have a constructor
d := distance(26.2)
```

Consider this code which creates two new new types from float32, which automatically get constructors, and you can create conversion methods between the two
```go
package main

import "fmt"

type distanceKm float32
type distanceMile float32

// this creates a method on the distanceMile type
// sort of like an instance method, but we don't have classes/instances in go
func (d distanceMile) toKm() distanceKm {
	return distanceKm(d * 1.60934)
}

// this creates a method on the distanceKm type
func (d distanceKm) toMile() distanceMile {
	return distanceMile(d / 1.60934)
}

func test() {
	dMi := distanceMile(26.2)
	dKm := distanceKm(dMi.toKm())
	fmt.Printf("Distance in miles: %v", dKm.toMile())
	fmt.Println("")
	fmt.Printf("Distance in km: %v", dMi.toKm())
	fmt.Println("")
}

func main() {
	test()
}
```

#### Complex types
* Structures
    * A data type with strongly typed properties
    * Sort of like a class, non inheritable though
    * Has default constructor, and can add methods to it just like regular types
* Interfaces
    * A definition of methods
    * Can emulate OOP-style polymorphism
    * Interfaces can be embedded in other interfaces

Struct usage:
```go
package main

import "fmt"

type User struct {
	id   int
	name string
}

func (u User) PrettyPrint() string {
	return fmt.Sprintf("User %v, name: %v", u.id, u.name)
}

func main() {
	var u0 User
	// default constructor usage, can use named properties or indexed properties
	u0 = User{id: 1, name: "Some guy"}
	fmt.Println(u0.PrettyPrint())
	u1 := User{2, "Cool dude"}
	fmt.Println(u1.PrettyPrint())
}
```

Example is at [/example/types](/example/types/main.go)

__Factories__
Factories are just functions that return new instances of a type. You could just do 

```go
User{id: 1, name: "some guy"}
// or use a factory fn
func NewUser(id int, name string) User {
    return User{id: id, name: name}
}
```

__Embedding__
Sort of like extending/polymorphism types in OOP. Not exactly though, the compiler just kind of pastes everything into your new type

__Interfaces__
Just a list of methods, that can then be used as a type.

# Async 
## Goroutines
* Goroutines are how Go uses threads
* Using `go someFunc()` will run someFunc in a thread
* The `main` function of a Go program is run in a goroutine, and it can spawn other goroutines
* Calling a goroutine is like calling an async function in JS, but not awaiting it. If your main function only has gorouting calls, the function will not wait for those to complete.
* Goroutines share memory, so you likely don't want to mutate variables that are shared by other goroutines

## Channel
* The way that goroutines communicate with each other
* A channel contains a value of any type. If goroutines are like promise-returning functions, channels are like the promises themselves
* Listeners can wait on a channel for its value to be present. Like awaiting a promise.
* Channels can also be buffered like a Stream, where multiple values are pushed to a channel piecemeal
* Channels will stay open unless you otherwise close them, so any function listening for a channel will stay open and create a deadlock error