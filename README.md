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
var x := 6
var y := 6.89

// if-else and logical expressions

if (x > 3) && (y > 6) {
    println("true")
} else {
    println("false")
}

// for cycle

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

var temp := 1

while temp < 10 {
    println("While Cycle, iter:", temp)
    if temp == 6 {
        break
    }
    temp = temp + 1
}

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


```


Result:

```go
true 
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
While Cycle, iter: 1 
While Cycle, iter: 2 
While Cycle, iter: 3 
While Cycle, iter: 4 
While Cycle, iter: 5 
While Cycle, iter: 6 
11 
13  
```

Ast-Tree print:

```go
!--> Declaration Variable (Statement) ['x'] |> Line 1
   !--> Integer Number (Number) Value: '6' |> Line 1
!--> Declaration Variable (Statement) ['y'] |> Line 2
   !--> Float Number (Number) Value: '6.890' |> Line 2
!--> If-Else Block (Block) |> Line 6
   !--> If Condition (Condition) |> Line 6
      !--> Logical Operation (Operation) '&&'' |> Line 6
         !--> Logical Operation (Operation) '>'' |> Line 6
            !--> Using Variable (Value) ['x'] |> Line 6
            !--> Integer Number (Number) Value: '3' |> Line 6
         !--> Logical Operation (Operation) '>'' |> Line 6
            !--> Using Variable (Value) ['y'] |> Line 6
            !--> Integer Number (Number) Value: '6' |> Line 6
   !--> If Case (Case) |> Line 7
      !--> Block |> Line 6
         !--> Operator 'println' (Operator) |> Line 7
            !--> String (String) Value: 'true' |> Line 7
   !--> Else Case (Case) |> Line 6
      !--> Block |> Line 8
         !--> Operator 'println' (Operator) |> Line 9
            !--> String (String) Value: 'false' |> Line 9
!--> For Block [IterVar: 'i'] (Block) |> Line 14
   !--> Start Section (Section) |> Line 15
      !--> Integer Number (Number) Value: '0' |> Line 14
   !--> Stop Section (Section) |> Line 15
      !--> Integer Number (Number) Value: '5' |> Line 14
   !--> Step Section (Section) |> Line 15
      !--> Integer Number (Number) Value: '1' |> Line 14
   !--> Iter Code Line 15
      !--> Block |> Line 14
         !--> Operator 'println' (Operator) |> Line 15
            !--> Binary Operation (Operation) '+' |> Line 15
               !--> String (String) Value: 'Cycle 1 ' |> Line 15
               !--> Operator 'string' (Operator) |> Line 15
                  !--> Using Variable (Value) ['i'] |> Line 15
!--> For Block [IterVar: 'i'] (Block) |> Line 18
   !--> Start Section (Section) |> Line 19
      !--> Integer Number (Number) Value: '0' |> Line 18
   !--> Stop Section (Section) |> Line 19
      !--> Integer Number (Number) Value: '5' |> Line 18
   !--> Step Section (Section) |> Line 19
      !--> Integer Number (Number) Value: '1' |> Line 18
   !--> Iter Code Line 19
      !--> Block |> Line 18
         !--> Operator 'println' (Operator) |> Line 19
            !--> Binary Operation (Operation) '+' |> Line 19
               !--> String (String) Value: 'Cycle 2 ' |> Line 19
               !--> Operator 'string' (Operator) |> Line 19
                  !--> Using Variable (Value) ['i'] |> Line 19
!--> For Block [IterVar: 'i'] (Block) |> Line 22
   !--> Start Section (Section) |> Line 23
      !--> Integer Number (Number) Value: '0' |> Line 22
   !--> Stop Section (Section) |> Line 23
      !--> Integer Number (Number) Value: '10' |> Line 22
   !--> Step Section (Section) |> Line 23
      !--> Integer Number (Number) Value: '2' |> Line 22
   !--> Iter Code Line 23
      !--> Block |> Line 22
         !--> Operator 'println' (Operator) |> Line 23
            !--> Binary Operation (Operation) '+' |> Line 23
               !--> String (String) Value: 'Cycle 3 ' |> Line 23
               !--> Operator 'string' (Operator) |> Line 23
                  !--> Using Variable (Value) ['i'] |> Line 23
!--> Declaration Variable (Statement) ['temp'] |> Line 28
   !--> Integer Number (Number) Value: '1' |> Line 28
!--> While-Block (Block) |> Line 39
!--> Cycle Condition (Condition) |> Line 39
!--> Logical Operation (Operation) '<'' |> Line 30
   !--> Using Variable (Value) ['temp'] |> Line 30
   !--> Integer Number (Number) Value: '10' |> Line 30
!--> Cycle Body (Block) |> Line 39
!--> Block |> Line 30
   !--> Operator 'println' (Operator) |> Line 31
      !--> String (String) Value: 'While Cycle, iter:' |> Line 31
      !--> Using Variable (Value) ['temp'] |> Line 31
   !--> If-Else Block (Block) |> Line 32
      !--> If Condition (Condition) |> Line 32
         !--> Logical Operation (Operation) '=='' |> Line 32
            !--> Using Variable (Value) ['temp'] |> Line 32
            !--> Integer Number (Number) Value: '6' |> Line 32
      !--> If Case (Case) |> Line 33
         !--> Block |> Line 32
            !--> Break (Statement) |> Line 34
   !--> Assign Variable (Statement) ['temp'] |> Line 35
      !--> Binary Operation (Operation) '+' |> Line 35
         !--> Using Variable (Value) ['temp'] |> Line 35
         !--> Integer Number (Number) Value: '1' |> Line 35
!--> Declaration Function ['simpleSum'] (Statement) [annotation int] |> Line 39
   !--> Arg Names: 
      +- x [annotation: int]
      +- y [annotation: float]
   !--> Block |> Line 39
      !--> Return (Statement) |> Line 41
         !--> Operator 'int' (Operator) |> Line 40
            !--> Binary Operation (Operation) '+' |> Line 40
               !--> Using Variable (Value) ['x'] |> Line 40
               !--> Using Variable (Value) ['y'] |> Line 40
!--> Declaration Function ['nilFunction'] (Statement) [annotation nil] |> Line 44
   !--> Arg Names: 
      +- x [annotation: int]
      +- y [annotation: int]
   !--> Block |> Line 44
      !--> Operator 'println' (Operator) |> Line 45
         !--> Binary Operation (Operation) '+' |> Line 45
            !--> Using Variable (Value) ['x'] |> Line 45
            !--> Using Variable (Value) ['y'] |> Line 45
!--> Declaration Variable (Statement) ['z'] |> Line 48
   !--> Operator 'simpleSum' (Operator) |> Line 48
      !--> Integer Number (Number) Value: '5' |> Line 48
      !--> Float Number (Number) Value: '6.700' |> Line 48
!--> Operator 'println' (Operator) |> Line 49
   !--> Using Variable (Value) ['z'] |> Line 49
!--> Operator 'nilFunction' (Operator) |> Line 51
   !--> Integer Number (Number) Value: '6' |> Line 51
   !--> Integer Number (Number) Value: '7' |> Line 51
```

