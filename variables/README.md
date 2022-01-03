# Summary of this module

Variable declaration:
- var foo int
- var foo int = 42
- foo := 42

Can't redeclare variables, but you can shadow them

All variables must be used

Visibility:
- lower case first letter for package scope
- upper case first letter to export
- no private scope, but can use block scope

Naming conventions:
- Pascal or camelCase - also capitalize acronyms (HTTP, URL, etc)
- As short as reasonable - longer names for longer lives

Type conversions
- destinationType(variable)
- use strconv package for strings