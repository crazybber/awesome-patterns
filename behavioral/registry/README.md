# go-registry

## A generic interface for the registry pattern

Because adding a registry pattern to each project can result in a lot of code duplication,
I decided to write a basic registry to slim down other projects.

The registry itself accepts empty interface for the values to be registered, so there's no need to
worry about type here. I do realize that this is not necessarily the best answer to making a
general purpose type, so I plan on adding some kind of type checking if I can find a good
way to do so.

---

## Usage

```go
import (
    "fmt"

    "github.com/kaezon/go-registry"
)

var myReg registry.Registry

myReg = registry.New()

newID := myReg.Register(1)

myReg.RegisterName(2, "second")

fmt.Println(myReg.Get(newID))
fmt.Println(myReg.Get("second"))

myReg.Deregister(newID)
myReg.Deregister("second")
```
