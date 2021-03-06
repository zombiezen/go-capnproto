package rpc_test

import (
	"sync"
	"testing"

	"golang.org/x/net/context"
	"zombiezen.com/go/capnproto"
	"zombiezen.com/go/capnproto/rpc"
	"zombiezen.com/go/capnproto/rpc/internal/logtransport"
	"zombiezen.com/go/capnproto/rpc/internal/pipetransport"
	"zombiezen.com/go/capnproto/rpc/internal/testcapnp"
	"zombiezen.com/go/capnproto/server"
)

func TestRelease(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	p, q := pipetransport.New()
	if *logMessages {
		p = logtransport.New(nil, p)
	}
	c := rpc.NewConn(p)
	hf := new(HandleFactory)
	d := rpc.NewConn(q, rpc.MainInterface(testcapnp.HandleFactory_ServerToClient(hf).Client))
	defer d.Wait()
	defer c.Close()
	client := testcapnp.HandleFactory{Client: c.Bootstrap(ctx)}
	r, err := client.NewHandle(ctx, func(r testcapnp.HandleFactory_newHandle_Params) error { return nil }).Struct()
	if err != nil {
		t.Fatal("NewHandle:", err)
	}
	handle := r.Handle()
	if n := hf.numHandles(); n != 1 {
		t.Fatalf("numHandles = %d; want 1", n)
	}

	if err := handle.Client.Close(); err != nil {
		t.Error("handle.Client.Close():", err)
	}
	flushConn(ctx, c)

	if n := hf.numHandles(); n != 0 {
		t.Errorf("numHandles = %d; want 0", n)
	}
}

func TestReleaseAlias(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	p, q := pipetransport.New()
	if *logMessages {
		p = logtransport.New(nil, p)
	}
	c := rpc.NewConn(p)
	hf := singletonHandleFactory()
	d := rpc.NewConn(q, rpc.MainInterface(testcapnp.HandleFactory_ServerToClient(hf).Client))
	defer d.Wait()
	defer c.Close()
	client := testcapnp.HandleFactory{Client: c.Bootstrap(ctx)}
	r1, err := client.NewHandle(ctx, func(r testcapnp.HandleFactory_newHandle_Params) error { return nil }).Struct()
	if err != nil {
		t.Fatal("NewHandle #1:", err)
	}
	handle1 := r1.Handle()
	r2, err := client.NewHandle(ctx, func(r testcapnp.HandleFactory_newHandle_Params) error { return nil }).Struct()
	if err != nil {
		t.Fatal("NewHandle #2:", err)
	}
	handle2 := r2.Handle()
	if n := hf.numHandles(); n != 1 {
		t.Fatalf("after creation, numHandles = %d; want 1", n)
	}

	if err := handle1.Client.Close(); err != nil {
		t.Error("handle1.Client.Close():", err)
	}
	flushConn(ctx, c)
	if n := hf.numHandles(); n != 1 {
		t.Errorf("after handle1.Client.Close(), numHandles = %d; want 1", n)
	}
	if err := handle2.Client.Close(); err != nil {
		t.Error("handle2.Client.Close():", err)
	}
	flushConn(ctx, c)
	if n := hf.numHandles(); n != 0 {
		t.Errorf("after handle1.Close() and handle2.Close(), numHandles = %d; want 0", n)
	}
}

func flushConn(ctx context.Context, c *rpc.Conn) {
	// discard result
	c.Bootstrap(ctx).Call(&capnp.Call{
		Ctx:        ctx,
		Method:     capnp.Method{InterfaceID: 0xdeadbeef, MethodID: 42},
		ParamsFunc: func(capnp.Struct) error { return nil },
		ParamsSize: capnp.ObjectSize{},
	}).Struct()
}

type Handle struct {
	f *HandleFactory
}

func (h Handle) Close() error {
	h.f.mu.Lock()
	h.f.n--
	h.f.mu.Unlock()
	return nil
}

type HandleFactory struct {
	n         int
	mu        sync.Mutex
	singleton testcapnp.Handle
}

func singletonHandleFactory() *HandleFactory {
	hf := new(HandleFactory)
	hf.singleton = testcapnp.Handle_ServerToClient(&Handle{f: hf})
	return hf
}

func (hf *HandleFactory) NewHandle(call testcapnp.HandleFactory_newHandle) error {
	server.Ack(call.Options)
	if hf.singleton.Client == nil {
		hf.mu.Lock()
		hf.n++
		hf.mu.Unlock()
		call.Results.SetHandle(testcapnp.Handle_ServerToClient(&Handle{f: hf}))
	} else {
		hf.mu.Lock()
		hf.n = 1
		hf.mu.Unlock()
		call.Results.SetHandle(hf.singleton)
	}
	return nil
}

func (hf *HandleFactory) numHandles() int {
	hf.mu.Lock()
	n := hf.n
	hf.mu.Unlock()
	return n
}
