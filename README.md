# Nil Map Access Panic in Go

This repository demonstrates a common error in Go: panicking due to accessing a nil map.  The `bug.go` file contains code that will cause a runtime panic because it attempts to access an element in an uninitialized map.  The `bugSolution.go` file offers a solution to prevent the panic.

## Problem

Maps in Go, if not explicitly initialized, have a nil value. Accessing a key in a nil map will result in a runtime panic.

## Solution

The solution involves checking if the map is nil before attempting to access any elements.  You can do this using a simple `if` statement or by using the comma ok idiom.
