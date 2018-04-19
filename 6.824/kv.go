package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

//
// Common RPC message
//

// Constant
const (
	OK       = "OK"
	ErrNoKey = "ErrNoKey"
)

// Err string type
type Err string

// PutArgs structure
type PutArgs struct {
	Key   string
	Value string
}

// PutReply structure
type PutReply struct {
	Err Err
}

// GetArgs structure
type GetArgs struct {
	Key string
}

// GetReply structure
type GetReply struct {
	Err   Err
	Value string
}

//
// Client
//

type Clerk struct {
	me *rpc.Client
}

func (ck *Clerk) Connect() {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	ck.me = client
}

func (ck *Clerk) Get(key string) string {
	args := GetArgs{key}
	reply := GetReply{}

	err := ck.me.Call("KV.Get", &args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}

	// ck.me.Close()
	return reply.Value
}

func (ck *Clerk) Put(key string, val string) {
	args := PutArgs{key, val}
	reply := PutReply{}

	err := ck.me.Call("KV.Put", &args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}

	// ck.me.Close()
}

//
// Server
//

type KV struct {
	mu   sync.Mutex
	data map[string]string
}

func server() {
	kv := new(KV)
	kv.data = map[string]string{}

	rpcs := rpc.NewServer()
	rpcs.Register(kv)

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	go func() {
		for {
			conn, err := l.Accept()
			if err == nil {
				go rpcs.ServeConn(conn)
			} else {
				break
			}
		}
		l.Close()
	}()
}

func (kv *KV) Get(args *GetArgs, reply *GetReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	val, ok := kv.data[args.Key]
	if ok {
		reply.Err = OK
		reply.Value = val
	} else {
		reply.Err = ErrNoKey
		reply.Value = ""
	}
	return nil
}

func (kv *KV) Put(args *PutArgs, reply *PutReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	kv.data[args.Key] = args.Value
	reply.Err = OK
	return nil
}

//
// main
//

func main() {
	server()

	kv1 := PutArgs{Key: "foo", Value: "bar"}

	clerk := new(Clerk)
	clerk.Connect()
	clerk.Put(kv1.Key, kv1.Value)

	fmt.Printf("Put(%v, %v)\n", kv1.Key, kv1.Value)
	fmt.Printf("get(%v) -> %s\n", kv1.Key, clerk.Get(kv1.Key))
}
