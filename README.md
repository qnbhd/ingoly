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
    print("true")
} else {
    print("false")
}

// for cycle

for i in [0; 5) {
    print("Cycle 1 " + i)
}

print("")

for i in [0; 5] {
    print("Cycle 2 " + i)
}

print("")

for i in [0; 10; 2] {
    print("Cycle 3 " + i)
}

// singe-comment

/*
multi-line comment
*/

```


Result:

```
true
Cycle 1 0.000
Cycle 1 1.000
Cycle 1 2.000
Cycle 1 3.000
Cycle 1 4.000

Cycle 2 0.000
Cycle 2 1.000
Cycle 2 2.000
Cycle 2 3.000
Cycle 2 4.000
Cycle 2 5.000

Cycle 3 0.000
Cycle 3 2.000
Cycle 3 4.000
Cycle 3 6.000
Cycle 3 8.000
Cycle 3 10.000
```

Ast-Tree print:

```
!--> Declaration Variable Parse (Parse) var 'x' Line 3
   !--> Value Node (Value) '6.000' Line 3
!--> Declaration Variable Parse (Parse) var 'y' Line 4
   !--> Value Node (Value) '6.890' Line 4
!--> If Else Block Line 8
!--> Logical Operation (Operation) '&&' Line 8
   !--> Logical Operation (Operation) '>' Line 8
      !--> Using Variable (Value) 'x' Line 8
      !--> Value Node (Value) '3.000' Line 8
   !--> Logical Operation (Operation) '>' Line 8
      !--> Using Variable (Value) 'y' Line 8
      !--> Value Node (Value) '6.000' Line 8
!--> If Case Line 9
   !--> Print Operator (Keyword) Line 9
      !--> Value Node (Value) 'true' Line 9
!--> Else Case 
   !--> Print Operator (Keyword) Line 11
      !--> Value Node (Value) 'false' Line 11
!--> For Block [iterVar 'i'] [0.000; 5.000; 1.000] Line: 16
   !--> Print Operator (Keyword) Line 17
      !--> Binary Operation (Operation) '+' Line 17
         !--> Value Node (Value) 'Cycle 1 ' Line 17
         !--> Using Variable (Value) 'i' Line 17
!--> Print Operator (Keyword) Line 20
   !--> Value Node (Value) '' Line 20
!--> For Block [iterVar 'i'] [0.000; 5.000; 1.000] Line: 22
   !--> Print Operator (Keyword) Line 23
      !--> Binary Operation (Operation) '+' Line 23
         !--> Value Node (Value) 'Cycle 2 ' Line 23
         !--> Using Variable (Value) 'i' Line 23
!--> Print Operator (Keyword) Line 26
   !--> Value Node (Value) '' Line 26
!--> For Block [iterVar 'i'] [0.000; 10.000; 2.000] Line: 28
   !--> Print Operator (Keyword) Line 29
      !--> Binary Operation (Operation) '+' Line 29
         !--> Value Node (Value) 'Cycle 3 ' Line 29
         !--> Using Variable (Value) 'i' Line 29

```

