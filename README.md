# learning-golang

### Comments
- line comment `//`
- multiple lines comment `/* */`
---
### Data Types
-
  #### boolean
    ```go
    var b1 bool = true
    // or
    b1 := false
    ```
-
  #### string
    ```go
    var str string = "hello world"
    // or
    str := "Hello"
    ```
-
  #### int - Signed Integers (uses first bit to identify either the sign + or -) (0 for + and 1 for -)
    - **int**: Depends on platform ; 32 bits in 32 bit systems and 64 bit in 64 bit systems
      ```go
      var num int = 1234
      ```
    - **int8**: 8 bits: -128 to 127
      ```go
      var num int8 = 127
      ```
    - **int16**	16 bits: -32768 to 32767
      ```go
      var num int16 = 32767
      ```
    - **int32**	32 bits: -2147483648 to 2147483647
      ```go
      var num int32 = -2147483648
      ```
    - **int64**	64 bits: -9223372036854775808 to 9223372036854775807
      ```go
      var num int64 = -9223372036854775808
      ```
-
  #### uint - Unsigned Integers (not use any bit to identify either the sign + or -)
    - **uint**: Depends on platform ; 32 bits in 32 bit systems and 64 bit in 64 bit systems
      ```go
      var num int = 1234
      ```
    - **uint8**: 8 bits: 0 to 255
      ```go
      var num uint8 = 255
      ```
    - **uint16**: 16 bits: 	0 to 65535
      ```go
      var num uint16 = 65535
      ```
    - **uint32**: 32 bits: 	0 to 4294967295
      ```go
      var num uint32 = 4294967295
      ```
    - **uint64**: 64 bits: 	0 to 18446744073709551615
      ```go
      var num uint64 = 18446744073709551615
      ```
- 
  #### float
    - **float32**: 32 bits: -3.4e+38 to 3.4e+38
      ```go
      var num float32 = 3.4e+38
      ```
    - **float64**: 64 bits: -1.7e+308 to +1.7e+308
      ```go
      var num float64 = 1.7976931e+308
      ```
---
### Variables
  - declare a variable with a specific type
    ```go
    var x int = 8
    ```
  - declare a variable and the type will be inferred
    ```go
    x := "Islam"
    ```
  - declare a variable without a value
    ```go
    var x int
    ```
  - declare multiple variables with the same type
    ```go
    var x, y, z int = 1, 2, 3
    ```
  - declare multiple variables with different types
    ```go
    var a, b = 6, "hello"
    ```
  - declare multiple variables with different types and the type will be inferred
    ```go
    c, d := 7, "world"
    ```
  - declare multiple variables in a block
    ```go
    var (
      a int
      b int = 1
      c string = "hello"
    )
    ```
---
### Constants
  - declare a constant variable and the type will be inferred
    ```go
    const PI = 3.14
    ```
  - declare a constant variable with specific type
    ```go
    const PI float32 = 3.14
    ```
---
### Output functions
- **Print** function inserts a space between the arguments if neither are strings
    ```go
    fmt.Print("hello world")
    
    // output
    hello world%
    ```
- **Println** function function prints its arguments and add a whitespace is between them, and a newline at the end
    ```go
    fmt.Println("hello world")

    // output
    hello world

    ```
- **Printf** function formats its argument based on the given formatting verb
  - `%v` is used to print the value of the arguments
    ```go
    var name = "Islam"
    fmt.Printf("hello %v", name)

    // output
    hello Islam%
    ```
  - `%T` is used to print the type of the arguments
    ```go
    var name = "Islam"
    fmt.Printf("hello %T", name)

    // output
    hello string%
    ```
---
### Arrays
  - #### Array declaration
    ```go
    var arr = [4]int{1, 2, 3} // length is defined
    fmt.Println(arr[3])
    // outoput
    0

    // or

    arr := [...]int{1, 2, 3, 4} // length is referred
    fmt.Println(arr)
    // output
    [1 2 3 4]
    ```
  - #### Access array elements
    ```go
    var arr = [3]int{1, 2, 3}
    fmt.Println(arr[0])
    // output
    1
    ```
  - #### initialize specific lements
    ```go
    arr := [5]int{1:10, 2:20}
    fmt.Println(arr)
    // output
    [0 10 20 0 0]
    ```
  - Array length
    ```go
    arr := [5]int{1,2,3,4,5}
    fmt.Println(len(arr))
    // output
    5
    ```
---
### Slices
 - 

