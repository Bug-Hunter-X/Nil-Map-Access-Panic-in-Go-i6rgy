```go
package main

import "fmt"

func main() {
    var m map[string]int
    if m == nil {
        m = make(map[string]int)
    }
    fmt.Println(m["key"]) // Now this is safe
    fmt.Println(m[0]) // This will still panic because 0 is not a string key
}
```