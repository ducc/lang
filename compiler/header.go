package compiler

const header = `// auto generated - changes will not be persisted
package main 

import (
        "fmt"
        "github.com/ducc/lang/util"
        "github.com/ducc/lang/builtins"
)

// incase imports are not used
var _ = fmt.Println
var _ = util.NewStack
var _ = builtins.Add
`
