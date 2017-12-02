package profut

import (
	"fmt"
	"time"
)

func ExampleBasic() {
	p := &Promise{}
	f := p.GetFuture()
	go func(p *Promise) {
		defer p.SetValue(42) /* std::promise::set_value_at_thread_exit */
		time.Sleep(100 * time.Millisecond)
	}(p)
	f.Wait()
	fmt.Println(f.Get())
	// Output: 42
}
