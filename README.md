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
// variable declarations

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

declare double_print(firstArg, secondArg) {
    println(firstArg, secondArg)
    println(firstArg, secondArg)
}

double_print("Hello", "World")

// singe-comment

/*
multi-line comment
*/

```


Result:

```
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
Hello World 
Hello World 
```

Ast-Tree print:

```
!--> Declaration Variable Parse (Parse) var 'x' Line 3
   !--> Integer (Number) Value: 6, Line: 3
!--> Declaration Variable Parse (Parse) var 'y' Line 4
   !--> Float (Number) Value: 6.890, Line: 4
!--> If Else Block Line 8
!--> Logical Operation (Operation) '&&' Line 8
   !--> Logical Operation (Operation) '>' Line 8
      !--> Using Variable (Value) 'x' Line 8
      !--> Integer (Number) Value: 3, Line: 8
   !--> Logical Operation (Operation) '>' Line 8
      !--> Using Variable (Value) 'y' Line 8
      !--> Integer (Number) Value: 6, Line: 8
!--> If Case Line 9
!--> Block
   !--> println Operator (Keyword) Line 9
      !--> String (String) Value: true, Line: 9
!--> Else Case 
!--> Block
   !--> println Operator (Keyword) Line 11
      !--> String (String) Value: false, Line: 11
!--> For Block [IterVar: 'i'] Line: 16
!--> Start Section Line 17
   !--> Integer (Number) Value: 0, Line: 16
!--> Stop Section Line 17
   !--> Integer (Number) Value: 5, Line: 16
!--> Step Section Line 17
   !--> Integer (Number) Value: 1, Line: 16
!--> Iter Code Line 17
!--> Block
   !--> println Operator (Keyword) Line 17
      !--> Binary Operation (Operation) '+' Line 17
         !--> String (String) Value: Cycle 1 , Line: 17
         !--> string Operator (Keyword) Line 17
            !--> Using Variable (Value) 'i' Line 17
!--> For Block [IterVar: 'i'] Line: 21
!--> Start Section Line 22
   !--> Integer (Number) Value: 0, Line: 21
!--> Stop Section Line 22
   !--> Integer (Number) Value: 5, Line: 21
!--> Step Section Line 22
   !--> Integer (Number) Value: 1, Line: 21
!--> Iter Code Line 22
!--> Block
   !--> println Operator (Keyword) Line 22
      !--> Binary Operation (Operation) '+' Line 22
         !--> String (String) Value: Cycle 2 , Line: 22
         !--> string Operator (Keyword) Line 22
            !--> Using Variable (Value) 'i' Line 22
!--> For Block [IterVar: 'i'] Line: 26
!--> Start Section Line 27
   !--> Integer (Number) Value: 0, Line: 26
!--> Stop Section Line 27
   !--> Integer (Number) Value: 10, Line: 26
!--> Step Section Line 27
   !--> Integer (Number) Value: 2, Line: 26
!--> Iter Code Line 27
!--> Block
   !--> println Operator (Keyword) Line 27
      !--> Binary Operation (Operation) '+' Line 27
         !--> String (String) Value: Cycle 3 , Line: 27
         !--> string Operator (Keyword) Line 27
            !--> Using Variable (Value) 'i' Line 27
!--> Declaration Function (Statement)Line 30
!--> Arg Names: 
   +- firstArg
   +- secondArg
!--> Block
   !--> println Operator (Keyword) Line 31
      !--> Using Variable (Value) 'firstArg' Line 31
      !--> Using Variable (Value) 'secondArg' Line 31
   !--> println Operator (Keyword) Line 32
      !--> Using Variable (Value) 'firstArg' Line 32
      !--> Using Variable (Value) 'secondArg' Line 32
!--> double_print Operator (Keyword) Line 35
   !--> String (String) Value: Hello, Line: 35
   !--> String (String) Value: World, Line: 35

```

