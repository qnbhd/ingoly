# Ingoly

<img src="https://i.postimg.cc/mgnx6Cfn/ingoly.png" alt="drawing" heigth="300" width="300"/>

[![GitHub issues](https://img.shields.io/github/issues/qnbhd/ingoly?style=for-the-badge)](https://github.com/qnbhd/ingoly/issues) [![GitHub forks](https://img.shields.io/github/forks/qnbhd/ingoly?style=for-the-badge)](https://github.com/qnbhd/ingoly/network) [![GitHub stars](https://img.shields.io/github/stars/qnbhd/ingoly?style=for-the-badge)](https://github.com/qnbhd/ingoly/stargazers) [![GitHub license](https://img.shields.io/github/license/qnbhd/ingoly?style=for-the-badge)](https://github.com/qnbhd/loripy/blob/master/LICENSE)

### Features

- Support nums, vars, strings, logical condition, if-else, for-while, functions, declaring functions

## Run

Run ingoly with manage script `.\ingoly [Target]` for Linux and `ingoly [Target]` 
`[Target]` - Source file

## Examples

Input INLY-code 

```go
require @builtins.math

var x := 6
var y := 6.89

// if-else and logical expressions

println("If-else test")

if (x > 3) && (y > 6) {
    println("true")
} else {
    println("false")
}

// for cycle

println("For test")

for i in [0; 5) {
    println("Cycle 1 " + string(i))
}

for i in [0; 5] {
    println("Cycle 2 " + string(i))
}

for i in [0; 10; 2] {
    println("Cycle 3 " + string(i))
}

// while cycle
println("While test")

var temp := 1

while temp < 10 {
    println("While Cycle, iter:", temp)
    if temp == 6 {
        break
    }
    temp = temp + 1
}

println("Function declaring test")

// function's declaring with type annotations
declare simpleSum(x int, y float) -> int {
    return int(x + y)
}

// function's returns nil
declare nilFunction(x int, y int) {
    println(x + y)
}

var z := simpleSum(5, 6.7)
println(z)

nilFunction(6, 7)


// singe-comment

/*
multi-line comment
*/

// math functions from builtins.math

println("Requiring math from builtins test")

for x in [0.0; 0.5; 0.1] {
    println(sin(x), cos(x), exp(x), sin(x)*sin(x) + cos(x)*cos(x))
}

// classes: fields, methods

class Vector {

    x float
    y float

    // methods

    declare get_length() -> float {
        return sqrt(this.x * this.x + this.y * this.y)
    }

}

declare add_vectors(a Vector, b Vector) -> Vector {
    return Vector(a.x + b.x, a.y + b.y)
}


println("Class method call test")

var xy := Vector(10, 5)
println(xy.get_length())

// arrays

println("Array test")

var arr := int -> [4, 5, 6]
var str := "abc"

for i in [0; |arr|) {
    println(arr[i], ":", str[i])
}
```


Result:

```go
If-else test 
true 
For test 
Cycle 1 0 
Cycle 1 1 
Cycle 1 2 
Cycle 1 3 
Cycle 1 4 
Cycle 2 0 
Cycle 2 1 
Cycle 2 2 
Cycle 2 3 
Cycle 2 4 
Cycle 2 5 
Cycle 3 0 
Cycle 3 2 
Cycle 3 4 
Cycle 3 6 
Cycle 3 8 
Cycle 3 10 
While test 
While Cycle, iter: 1 
While Cycle, iter: 2 
While Cycle, iter: 3 
While Cycle, iter: 4 
While Cycle, iter: 5 
While Cycle, iter: 6 
Function declaring test 
11 
13 
Requiring math from builtins test 
0 1 1 1 
0.09983341664682815 0.9950041652780257 1.1051709180756477 0.9999999999999999 
0.19866933079506122 0.9800665778412416 1.2214027581601699 1 
0.2955202066613396 0.955336489125606 1.3498588075760032 1 
0.3894183423086505 0.921060994002885 1.4918246976412703 0.9999999999999998 
0.479425538604203 0.8775825618903728 1.6487212707001282 1 
Class method call test 
11.180339887498949 
Array test 
4 : a 
5 : b 
6 : c  
```

Ast-Tree print:

```go
!--> Declaration Function ['sin'] (Statement) [annotation float] |> Line 3
   !--> Arg Names: 
      +- x [annotation: float]
   !--> Block |> Line 3
      !--> Return (Statement) |> Line 5
         !--> Operator '__builtin__sin' (Operator) |> Line 4
            !--> Using Variable (Value) ['x'] |> Line 4
!--> Declaration Function ['cos'] (Statement) [annotation float] |> Line 7
   !--> Arg Names: 
      +- x [annotation: float]
   !--> Block |> Line 7
      !--> Return (Statement) |> Line 9
         !--> Operator '__builtin__cos' (Operator) |> Line 8
            !--> Using Variable (Value) ['x'] |> Line 8
!--> Declaration Function ['tan'] (Statement) [annotation float] |> Line 11
   !--> Arg Names: 
      +- x [annotation: float]
   !--> Block |> Line 11
      !--> Return (Statement) |> Line 13
         !--> Binary Operation (Operation) '/' |> Line 12
            !--> Operator 'sin' (Operator) |> Line 12
               !--> Using Variable (Value) ['x'] |> Line 12
            !--> Operator 'cos' (Operator) |> Line 12
               !--> Using Variable (Value) ['x'] |> Line 12
!--> Declaration Function ['cot'] (Statement) [annotation float] |> Line 15
   !--> Arg Names: 
      +- x [annotation: float]
   !--> Block |> Line 15
      !--> Return (Statement) |> Line 17
         !--> Binary Operation (Operation) '/' |> Line 16
            !--> Integer Number (Number) Value: '1' |> Line 16
            !--> Operator 'tan' (Operator) |> Line 16
               !--> Using Variable (Value) ['x'] |> Line 16
!--> Declaration Function ['sqrt'] (Statement) [annotation float] |> Line 19
   !--> Arg Names: 
      +- x [annotation: float]
   !--> Block |> Line 19
      !--> Return (Statement) |> Line 21
         !--> Operator '__builtin__sqrt' (Operator) |> Line 20
            !--> Using Variable (Value) ['x'] |> Line 20
!--> Declaration Function ['abs'] (Statement) [annotation float] |> Line 23
   !--> Arg Names: 
      +- x [annotation: float]
   !--> Block |> Line 23
      !--> Return (Statement) |> Line 25
         !--> Operator '__builtin__abs' (Operator) |> Line 24
            !--> Using Variable (Value) ['x'] |> Line 24
!--> Declaration Function ['exp'] (Statement) [annotation float] |> Line 27
   !--> Arg Names: 
      +- x [annotation: float]
   !--> Block |> Line 27
      !--> Return (Statement) |> Line 29
         !--> Operator '__builtin__exp' (Operator) |> Line 28
            !--> Using Variable (Value) ['x'] |> Line 28
!--> Declaration Variable (Statement) ['x'] |> Line 31
   !--> Integer Number (Number) Value: '6' |> Line 31
!--> Declaration Variable (Statement) ['y'] |> Line 32
   !--> Float Number (Number) Value: '6.890' |> Line 32
!--> Operator 'println' (Operator) |> Line 36
   !--> String (String) Value: 'If-else test' |> Line 36
!--> If-Else Block (Block) |> Line 38
   !--> If Condition (Condition) |> Line 38
      !--> Logical Operation (Operation) '&&'' |> Line 38
         !--> Logical Operation (Operation) '>'' |> Line 38
            !--> Using Variable (Value) ['x'] |> Line 38
            !--> Integer Number (Number) Value: '3' |> Line 38
         !--> Logical Operation (Operation) '>'' |> Line 38
            !--> Using Variable (Value) ['y'] |> Line 38
            !--> Integer Number (Number) Value: '6' |> Line 38
   !--> If Case (Case) |> Line 39
      !--> Block |> Line 38
         !--> Operator 'println' (Operator) |> Line 39
            !--> String (String) Value: 'true' |> Line 39
   !--> Else Case (Case) |> Line 38
      !--> Block |> Line 40
         !--> Operator 'println' (Operator) |> Line 41
            !--> String (String) Value: 'false' |> Line 41
!--> Operator 'println' (Operator) |> Line 46
   !--> String (String) Value: 'For test' |> Line 46
!--> For Block [IterVar: 'i'] (Block) |> Line 48
   !--> Start Section (Section) |> Line 49
      !--> Integer Number (Number) Value: '0' |> Line 48
   !--> Stop Section (Section) |> Line 49
      !--> Integer Number (Number) Value: '5' |> Line 48
   !--> Step Section (Section) |> Line 49
      !--> Integer Number (Number) Value: '1' |> Line 48
   !--> Iter Code Line 49
      !--> Block |> Line 48
         !--> Operator 'println' (Operator) |> Line 49
            !--> Binary Operation (Operation) '+' |> Line 49
               !--> String (String) Value: 'Cycle 1 ' |> Line 49
               !--> Operator 'string' (Operator) |> Line 49
                  !--> Using Variable (Value) ['i'] |> Line 49
!--> For Block [IterVar: 'i'] (Block) |> Line 52
   !--> Start Section (Section) |> Line 53
      !--> Integer Number (Number) Value: '0' |> Line 52
   !--> Stop Section (Section) |> Line 53
      !--> Integer Number (Number) Value: '5' |> Line 52
   !--> Step Section (Section) |> Line 53
      !--> Integer Number (Number) Value: '1' |> Line 52
   !--> Iter Code Line 53
      !--> Block |> Line 52
         !--> Operator 'println' (Operator) |> Line 53
            !--> Binary Operation (Operation) '+' |> Line 53
               !--> String (String) Value: 'Cycle 2 ' |> Line 53
               !--> Operator 'string' (Operator) |> Line 53
                  !--> Using Variable (Value) ['i'] |> Line 53
!--> For Block [IterVar: 'i'] (Block) |> Line 56
   !--> Start Section (Section) |> Line 57
      !--> Integer Number (Number) Value: '0' |> Line 56
   !--> Stop Section (Section) |> Line 57
      !--> Integer Number (Number) Value: '10' |> Line 56
   !--> Step Section (Section) |> Line 57
      !--> Integer Number (Number) Value: '2' |> Line 56
   !--> Iter Code Line 57
      !--> Block |> Line 56
         !--> Operator 'println' (Operator) |> Line 57
            !--> Binary Operation (Operation) '+' |> Line 57
               !--> String (String) Value: 'Cycle 3 ' |> Line 57
               !--> Operator 'string' (Operator) |> Line 57
                  !--> Using Variable (Value) ['i'] |> Line 57
!--> Operator 'println' (Operator) |> Line 61
   !--> String (String) Value: 'While test' |> Line 61
!--> Declaration Variable (Statement) ['temp'] |> Line 63
   !--> Integer Number (Number) Value: '1' |> Line 63
!--> While-Block (Block) |> Line 73
!--> Cycle Condition (Condition) |> Line 73
!--> Logical Operation (Operation) '<'' |> Line 65
   !--> Using Variable (Value) ['temp'] |> Line 65
   !--> Integer Number (Number) Value: '10' |> Line 65
!--> Cycle Body (Block) |> Line 73
!--> Block |> Line 65
   !--> Operator 'println' (Operator) |> Line 66
      !--> String (String) Value: 'While Cycle, iter:' |> Line 66
      !--> Using Variable (Value) ['temp'] |> Line 66
   !--> If-Else Block (Block) |> Line 67
      !--> If Condition (Condition) |> Line 67
         !--> Logical Operation (Operation) '=='' |> Line 67
            !--> Using Variable (Value) ['temp'] |> Line 67
            !--> Integer Number (Number) Value: '6' |> Line 67
      !--> If Case (Case) |> Line 68
         !--> Block |> Line 67
            !--> Break (Statement) |> Line 69
   !--> Assign Variable (Statement) ['temp'] |> Line 70
      !--> Binary Operation (Operation) '+' |> Line 70
         !--> Using Variable (Value) ['temp'] |> Line 70
         !--> Integer Number (Number) Value: '1' |> Line 70
!--> Operator 'println' (Operator) |> Line 73
   !--> String (String) Value: 'Function declaring test' |> Line 73
!--> Declaration Function ['simpleSum'] (Statement) [annotation int] |> Line 76
   !--> Arg Names: 
      +- x [annotation: int]
      +- y [annotation: float]
   !--> Block |> Line 76
      !--> Return (Statement) |> Line 78
         !--> Operator 'int' (Operator) |> Line 77
            !--> Binary Operation (Operation) '+' |> Line 77
               !--> Using Variable (Value) ['x'] |> Line 77
               !--> Using Variable (Value) ['y'] |> Line 77
!--> Declaration Function ['nilFunction'] (Statement) [annotation nil] |> Line 81
   !--> Arg Names: 
      +- x [annotation: int]
      +- y [annotation: int]
   !--> Block |> Line 81
      !--> Operator 'println' (Operator) |> Line 82
         !--> Binary Operation (Operation) '+' |> Line 82
            !--> Using Variable (Value) ['x'] |> Line 82
            !--> Using Variable (Value) ['y'] |> Line 82
!--> Declaration Variable (Statement) ['z'] |> Line 85
   !--> Operator 'simpleSum' (Operator) |> Line 85
      !--> Integer Number (Number) Value: '5' |> Line 85
      !--> Float Number (Number) Value: '6.700' |> Line 85
!--> Operator 'println' (Operator) |> Line 86
   !--> Using Variable (Value) ['z'] |> Line 86
!--> Operator 'nilFunction' (Operator) |> Line 88
   !--> Integer Number (Number) Value: '6' |> Line 88
   !--> Integer Number (Number) Value: '7' |> Line 88
!--> Operator 'println' (Operator) |> Line 97
   !--> String (String) Value: 'Requiring math from builtins test' |> Line 97
!--> For Block [IterVar: 'x'] (Block) |> Line 99
   !--> Start Section (Section) |> Line 100
      !--> Float Number (Number) Value: '0.000' |> Line 99
   !--> Stop Section (Section) |> Line 100
      !--> Float Number (Number) Value: '0.500' |> Line 99
   !--> Step Section (Section) |> Line 100
      !--> Float Number (Number) Value: '0.100' |> Line 99
   !--> Iter Code Line 100
      !--> Block |> Line 99
         !--> Operator 'println' (Operator) |> Line 100
            !--> Operator 'sin' (Operator) |> Line 100
               !--> Using Variable (Value) ['x'] |> Line 100
            !--> Operator 'cos' (Operator) |> Line 100
               !--> Using Variable (Value) ['x'] |> Line 100
            !--> Operator 'exp' (Operator) |> Line 100
               !--> Using Variable (Value) ['x'] |> Line 100
            !--> Binary Operation (Operation) '+' |> Line 100
               !--> Binary Operation (Operation) '*' |> Line 100
                  !--> Operator 'sin' (Operator) |> Line 100
                     !--> Using Variable (Value) ['x'] |> Line 100
                  !--> Operator 'sin' (Operator) |> Line 100
                     !--> Using Variable (Value) ['x'] |> Line 100
               !--> Binary Operation (Operation) '*' |> Line 100
                  !--> Operator 'cos' (Operator) |> Line 100
                     !--> Using Variable (Value) ['x'] |> Line 100
                  !--> Operator 'cos' (Operator) |> Line 100
                     !--> Using Variable (Value) ['x'] |> Line 100
!--> Class Declaring Declare Vector |> Line 118
   +- Field: x [annotation: float]
   +- Field: y [annotation: float]
   !--> Declaration Function ['get_length'] (Statement) [annotation float] |> Line 112
      !--> Block |> Line 112
         !--> Return (Statement) |> Line 114
            !--> Operator 'sqrt' (Operator) |> Line 113
               !--> Binary Operation (Operation) '+' |> Line 113
                  !--> Binary Operation (Operation) '*' |> Line 113
                     !--> Class RHS Access to 'this' [field: x] |> Line 113
                     !--> Class RHS Access to 'this' [field: x] |> Line 113
                  !--> Binary Operation (Operation) '*' |> Line 113
                     !--> Class RHS Access to 'this' [field: y] |> Line 113
                     !--> Class RHS Access to 'this' [field: y] |> Line 113
!--> Declaration Function ['add_vectors'] (Statement) [annotation Vector] |> Line 118
   !--> Arg Names: 
      +- a [annotation: Vector]
      +- b [annotation: Vector]
   !--> Block |> Line 118
      !--> Return (Statement) |> Line 120
         !--> Operator 'Vector' (Operator) |> Line 119
            !--> Binary Operation (Operation) '+' |> Line 119
               !--> Class RHS Access to 'a' [field: x] |> Line 119
               !--> Class RHS Access to 'b' [field: x] |> Line 119
            !--> Binary Operation (Operation) '+' |> Line 119
               !--> Class RHS Access to 'a' [field: y] |> Line 119
               !--> Class RHS Access to 'b' [field: y] |> Line 119
!--> Operator 'println' (Operator) |> Line 123
   !--> String (String) Value: 'Class method call test' |> Line 123
!--> Declaration Variable (Statement) ['xy'] |> Line 125
   !--> Operator 'Vector' (Operator) |> Line 125
      !--> Integer Number (Number) Value: '10' |> Line 125
      !--> Integer Number (Number) Value: '5' |> Line 125
!--> Operator 'println' (Operator) |> Line 126
   !--> Var 'xy' Executing Method 'get_length' |> Line 126
!--> Operator 'println' (Operator) |> Line 130
   !--> String (String) Value: 'Array test' |> Line 130
!--> Declaration Variable (Statement) ['arr'] |> Line 132
   !--> Array [el annotation: int] |> Line 132
      !--> Integer Number (Number) Value: '4' |> Line 132
      !--> Integer Number (Number) Value: '5' |> Line 132
      !--> Integer Number (Number) Value: '6' |> Line 132
!--> Declaration Variable (Statement) ['str'] |> Line 133
   !--> String (String) Value: 'abc' |> Line 133
!--> For Block [IterVar: 'i'] (Block) |> Line 135
   !--> Start Section (Section) |> Line 136
      !--> Integer Number (Number) Value: '0' |> Line 135
   !--> Stop Section (Section) |> Line 136
      !--> Operator 'len' (Operator) |> Line 135
         !--> Using Variable (Value) ['arr'] |> Line 135
   !--> Step Section (Section) |> Line 136
      !--> Integer Number (Number) Value: '1' |> Line 135
   !--> Iter Code Line 136
      !--> Block |> Line 135
         !--> Operator 'println' (Operator) |> Line 136
            !--> Array Access to 'arr' array |> Line 136
               +- Index
                  !--> Using Variable (Value) ['i'] |> Line 136
            !--> String (String) Value: ':' |> Line 136
            !--> Array Access to 'str' array |> Line 136
               +- Index
                  !--> Using Variable (Value) ['i'] |> Line 136

```

