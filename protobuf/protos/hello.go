package main

type HelloService struct{}

func (p *HelloService) Hello(request *StringM, reply *StringM) error {
	reply.Value = "Hello, " + request.GetValue()
	return nil
}
