package profut

type Promise struct {
	Future  *Future
	Request chan chan interface{}
}

type Future struct {
	Promise *Promise
	Value   interface{}
}

func (p *Promise) GetFuture() *Future {
	if p.Future == nil {
		p.Request = make(chan chan interface{}, 2)
		p.Future = &Future{
			Promise: p,
		}
	}
	return p.Future
}

func (p *Promise) SetValue(v interface{}) {
	res := <-p.Request
	res <- v
}

func (f *Future) Wait() {
	res := make(chan interface{}, 2)
	f.Promise.Request <- res
	f.Value = <-res
}

func (f *Future) Get() interface{} {
	return f.Value
}
