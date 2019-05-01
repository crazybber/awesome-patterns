package main

type Operator interface {
	Apply(int, int) int
}

// 通过定义内部的Operator实现不同策略的切换
type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}

func main(){
	a ,b := 2 ,4
	op := &Operation{
		Addition{},
	}
	println(op.Operate(a ,b))// 6
	op.Operator = Multiplication{}
	println(op.Operate(a ,b))// 8
}
