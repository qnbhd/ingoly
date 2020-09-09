# Ingoly

<img src="https://i.postimg.cc/mgnx6Cfn/ingoly.png" alt="drawing" heigth="300" width="300"/>

[![GitHub issues](https://img.shields.io/github/issues/qnbhd/ingoly?style=for-the-badge)](https://github.com/qnbhd/ingoly/issues) [![GitHub forks](https://img.shields.io/github/forks/qnbhd/ingoly?style=for-the-badge)](https://github.com/qnbhd/ingoly/network) [![GitHub stars](https://img.shields.io/github/stars/qnbhd/ingoly?style=for-the-badge)](https://github.com/qnbhd/ingoly/stargazers) [![GitHub license](https://img.shields.io/github/license/qnbhd/ingoly?style=for-the-badge)](https://github.com/qnbhd/loripy/blob/master/LICENSE)

### Features

- Support nums, vars, strings, logical condition, if-else, print, comments now

## Run

Run ingoly with manage script `.\ingoly [Target]` for Linux and `ingoly [Target]` 
`[Target]` - Source file

## Examples

Input INLY-code 

```html
var x := 6
var y := 7

print(x + y * 10)

if x <= 10:
  print("x <= 10")
else:
  print("x > 10")


print("Кириллица и все остальное")

```


Result:

```
76.000
x <= 10
Кириллица и все остальное
```

Ast-Tree print:

```
!--> Declaration Variable Statement (Statement) var 'x'
   !--> Value Node (Value) '6.000'
!--> Declaration Variable Statement (Statement) var 'y'
   !--> Value Node (Value) '7.000'
!--> Print Operator (Keyword)
   !--> Binary Operation (Operation) '+'
      !--> Using Variable (Value) 'x'
      !--> Binary Operation (Operation) '*'
         !--> Using Variable (Value) 'y'
         !--> Value Node (Value) '10.000'
!--> If Else Block
   !--> Logical Operation (Operation) '<='
      !--> Using Variable (Value) 'x'
      !--> Value Node (Value) '10.000'
!--> If Case
   !--> Print Operator (Keyword)
      !--> Value Node (Value) 'x <= 10'
!--> Else Case
   !--> Print Operator (Keyword)
      !--> Value Node (Value) 'x > 10'
!--> Print Operator (Keyword)
   !--> Value Node (Value) 'Кириллица и все остальное'
```

