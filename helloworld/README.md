# How to run code
```bash 
go run {{filename}}.go
```

Other commands:
```bash 
go build {{filename}}.go  # compiles a program
go run {{filename}}.go    # compiles and exicutes program
go fmt {{filename}}.go    # formats code in file
go install {{package}}    # compiles and installs a package
go get {{package}}        # downloads raw source code for someone else's package
go test {{filename}}.go   # runs any tests associated with the program
```

# What is package header

package == project == workspace

Package Main
- main.go
    - contains `package main` at top of file
- support.go
    - contains `package main` at top of file
- helper.go
    - contains `package main` at top of file

This is how we group packages. All files above are
wrapped in the Package Main.

Types of packages:
- Executable
- Reusable
    - helper functions or reusable code

`main` is used to make an executable package.
`main` is used for only executable packages.

# What is import
How you import other packages.

"fmt" (format) is a standard library.

other standard libraries:
- debug
- math
- encoding
- crypto
- io

# What is func
func is function

nothing new...

# How is .go file organized
1. package declaration
2. package imports
3. functions