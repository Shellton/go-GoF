package adapter

type TargetInterface interface {
	MethodA()
	MethodB()
}

type Adaptee struct{}

func (*Adaptee) MethodA() {

}
func (*Adaptee) MethodB() {

}

type Adaptor struct {
	adaptee Adaptee
}

func (a *Adaptor) MethodA() {
	a.adaptee.MethodA()
}

func (a *Adaptor) MethodB() {
	// do my stuff
}
