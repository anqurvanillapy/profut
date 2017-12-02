<div align="center"><h1><em>
Profut
</em></h1></div>

**TOY ALERT**.

A simple `std::future<T>`-like promise and future library in Go.

## Install

```bash
$ go get -u github.com/anqurvanillapy/profut
```

## Example

```go
import (
    "fmt"
    "time"
    "github.com/anqurvanillapy/profut"
)

func main() {
    p := &profut.Promise{}
    f := p.GetFuture()
    go func(p *profut.Promise) {
        defer p.SetValue(42) /* like std::promise::set_value_at_thread_exit */
        time.Sleep(100 * time.Millisecond)
    }(p)
    f.Wait()
    fmt.Println(f.Get())
    // Output: 42
}
```

## License

MIT
