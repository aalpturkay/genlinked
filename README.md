# GenLinked

Golang generic linked-list implementation using mutexes for thread-safety.

## Usage

go get the package

```bash
go get github.com/aalpturkay/genlinked
```

importing

```go
import "github.com/aalpturkay/genlinked"
```

create a linked-list with multiple items

```go
categories := NewLinkedListWithItems([]string{"burger", "pizza", "wrap", "icecream"})
```
