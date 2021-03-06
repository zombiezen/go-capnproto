package rpc_test

import (
	"testing"

	"golang.org/x/net/context"
	"zombiezen.com/go/capnproto/rpc"
	"zombiezen.com/go/capnproto/rpc/internal/logtransport"
	"zombiezen.com/go/capnproto/rpc/internal/pipetransport"
	"zombiezen.com/go/capnproto/rpc/internal/testcapnp"
)

func TestIssue3(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	p, q := pipetransport.New()
	if *logMessages {
		p = logtransport.New(nil, p)
	}
	c := rpc.NewConn(p)
	echoSrv := testcapnp.Echoer_ServerToClient(new(SideEffectEchoer))
	d := rpc.NewConn(q, rpc.MainInterface(echoSrv.Client))
	defer d.Wait()
	defer c.Close()
	client := testcapnp.Echoer{Client: c.Bootstrap(ctx)}
	localCap := testcapnp.CallOrder_ServerToClient(new(CallOrder))
	echo := client.Echo(ctx, func(p testcapnp.Echoer_echo_Params) error {
		return p.SetCap(localCap)
	})

	// This should not deadlock.
	_, err := echo.Struct()
	if err != nil {
		t.Error("Echo error:", err)
	}
}

type SideEffectEchoer struct {
	CallOrder
}

func (*SideEffectEchoer) Echo(call testcapnp.Echoer_echo) error {
	call.Params.Cap().GetCallSequence(call.Ctx, func(p testcapnp.CallOrder_getCallSequence_Params) error {
		return nil
	})
	call.Results.SetCap(call.Params.Cap())
	return nil
}
