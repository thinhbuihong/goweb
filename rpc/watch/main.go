package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"sync"
	"time"
)

type KVStoreService struct {
	//store key-value pairs
	m map[string]string
	//list filter function is called when change key value pairs
	filer map[string]func(key string)
	mu    sync.Mutex
}

func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:     make(map[string]string),
		filer: make(map[string]func(key string)),
	}
}

func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}

	return fmt.Errorf("not found")
}

func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	key, value := kv[0], kv[1]

	if oldValue := p.m[key]; oldValue != value {
		//cal filter funcs
		for _, fn := range p.filer {
			fn(key)
			// fn(key,val)
		}
	}

	p.m[key] = value
	return nil
}

func (p *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())

	ch := make(chan string, 10)

	p.mu.Lock()
	p.filer[id] = func(key string) { ch <- key }
	p.mu.Unlock()

	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	}
}

func main() {
	rpc.RegisterName("KVStoreService", NewKVStoreService())
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal("listenTCP error: ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept error: ", err)
		}

		go rpc.ServeConn(conn)
	}
}
