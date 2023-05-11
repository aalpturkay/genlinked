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
categories := genlinked.NewLinkedListWithItems([]string{"burger", "pizza", "wrap", "icecream"})
```

initialize linked-list with a custom type

```go
func main() {

	type song struct {
		name     string
		duration uint64
	}

	playlist := genlinked.NewLinkedList[song]()

	playlist.Add(song{"Master Of Puppets", 515000})
	playlist.Add(song{"Issızlığın Ortasında", 243000})

	fmt.Println(playlist)

	// {Master Of Puppets 515000} -> {Issızlığın Ortasında 243000} -> nil
}
```
