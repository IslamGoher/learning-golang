# learning-golang

### Comments
- line comment `//`
- multiple lines comment `/* */`
---
### Data Types
-
  #### boolean
    - `var b1 bool = true`
    - `b2 := false`
-
  #### string
    - `var str string = "hello world"`
    - `str := "Hello"`
-
  #### int - Signed Integers (uses first bit to identify either the sign + or -) (0 for + and 1 for -)
    - int: Depends on platform ; 32 bits in 32 bit systems and 64 bit in 64 bit systems
      - `var num int = 1234`
    - int8: 8 bits: -128 to 127
      - `var num int8 = 127`
    - int16	16 bits: -32768 to 32767
      - `var num int16 = 32767`
    - int32	32 bits: -2147483648 to 2147483647
      - `var num int32 = -2147483648`
    - int64	64 bits: -9223372036854775808 to 9223372036854775807
      - `var num int64 = -9223372036854775808`
-
  #### uint - Unsigned Integers (not use any bit to identify either the sign + or -)
    - uint: Depends on platform ; 32 bits in 32 bit systems and 64 bit in 64 bit systems
      - `var num int = 1234`
    - uint8: 8 bits: 0 to 255
      - `var num uint8 = 255`
    - uint16: 16 bits: 	0 to 65535
      - `var num uint16 = 65535`
    - uint32: 32 bits: 	0 to 4294967295
      - `var num uint32 = 4294967295`
    - uint64: 64 bits: 	0 to 18446744073709551615
      - `var num uint64 = 18446744073709551615`
- 
  #### float
    - float32: 32 bits: -3.4e+38 to 3.4e+38
      - `var num float32 = 3.4e+38`
    - float64: 64 bits: -1.7e+308 to +1.7e+308
      - `var num float64 = 1.7976931e+308`
---
### Variables
  - declear a variable with a specific type
    - `var x int = 8`
  - declear a variable and the type will be inferred
    - `x := "Islam"`
  - declear a variable without a value
    - `var x int`
  - declear multiple variables with the same type
    - `var x, y, z int = 1, 2, 3`
  - declear multiple variables with different types
    - `var a,b = 6, "hello"`
  - declear multiple variables with different types and the type will be inferred
    - `c, d := 7, "world"`
  - declear multiple variables in a block
    ```go
    var (
      a int
      b int = 1
      c string = "hello"
    )
    ```
---
### Constants
  - declear a constant variable and the type will be inferred
    - `const PI = 3.14`
  - declear a constant variable with specific type
    - `const PI float32 = 3.14`
---
