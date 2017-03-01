# funcgistry


[![Build Status](https://travis-ci.org/nuying117/funcgistry.svg?branch=master)](https://travis-ci.org/nuying117/funcgistry)  [![codecov](https://codecov.io/gh/nuying117/funcgistry/branch/master/graph/badge.svg)](https://codecov.io/gh/nuying117/funcgistry)


funcgistry is a go library to maintain a mapping between a function name(string) and a function, it can be used to invoke a function via its name.

#example
```
package main
import (
  "fmt"
  "github.com/nuying117/funcgistry"
)
func testFunc(str string) {
  fmt.Println(str)
}

func init() {
  funcgistry.AddElement("test", testFunc)
}

func main(){
  funcgistry.Call("test", "hello, funcgistry")
}
```

it also helps when you need to implement a framework to auto invoke the handler of a http request in a http server
