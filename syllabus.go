package main

import "fmt"

func main() {
	fmt.Print(`
	
# TSS_GO_programming

Chapter 1: Introduction to Go
Points Covered:

- Overview and history of Go.

- Setting up the Go development environment.

- Writing and executing a basic program.

- Understanding variables, constants, and data types (var, const, int, float64, string).

- Handling user input/output (fmt.Println, fmt.Scan).

Keywords: var, const, int, float64, string, bool, if, else, for, fmt.Println.

---------------------------------------------------------------------
Chapter 2: Functions
Points Covered:

- Function declaration and invocation (func, return).

- Passing arguments and receiving return values.

- Using variadic functions (... for dynamic input).

- Understanding and implementing anonymous functions (func() {}).

- Closures for encapsulating variables.

- Error handling with defer, panic, and recover.

Keywords: func, return, defer, panic, recover, ... (variadic), anonymous function, closures.

--------------------------------------------------------------------------------------
Chapter 3: Working with Data
Points Covered:

- Using arrays ([n]type), slices ([]type), and maps (map[keyType]valueType) for collections.

- Iterating through collections with loops (for, range).

- Understanding and using pointers (*, &).

- String manipulation techniques (string, len, +, fmt.Sprintf).

- JSON/XML parsing and serialization (json.Marshal, json.Unmarshal).

- Reading from and writing to files (os, io/ioutil).

Keywords: array, slice, map, for, range, *, &, string, len, +, json.Marshal.

--------------------------------------------------------------------------------
Chapter 4: Methods and Interfaces
Points Covered:

- Creating methods for custom types (type, method).

- Differentiating between pointer (*) and value receivers.

- Declaring and implementing interfaces (interface, type).

- Using type assertion (.(type)) and type switches.

- Demonstrating polymorphism via interfaces.

- Hands-on practical scenarios.

Keywords: type, method, interface, pointer (*), value receiver, type assertion, .(type).

---------------------------------------------------------------------------------------------------
Chapter 5: Go Routines and Channels
Points Covered:

- Basics of concurrency in Go (go, sync).

- Creating and managing goroutines (go func()).

- Communicating between goroutines using channels (chan, <-).

- Differentiating between buffered and unbuffered channels (make(chan type, capacity)).

- Synchronization with select and wait groups (sync.WaitGroup).


Keywords: go, sync, chan, <-, select, make.

-----------------------------------------------------------------------------------------------
Chapter 6: Packages and Files
Points Covered:

- Creating reusable custom packages (package, import).

- Importing and managing dependencies.

- File I/O operations (os.Create, os.Open, os.ReadFile, os.WriteFile).

- Performing directory-related tasks (os.Mkdir, os.Remove).

- Utilizing external libraries and packages.

- Structuring projects with best practices.

Keywords: package, import, os, io/ioutil, file, directory.
`)
}
