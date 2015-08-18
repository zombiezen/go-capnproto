package testcapnp

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	C "zombiezen.com/go/capnproto"
	server "zombiezen.com/go/capnproto/server"
)

type Handle struct{ C.Client }

type Handle_Server interface {
}

func Handle_ServerToClient(s Handle_Server) Handle {
	c, _ := s.(server.Closer)
	return Handle{Client: server.New(Handle_Methods(nil, s), c)}
}

func Handle_Methods(methods []server.Method, s Handle_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 0)
	}

	return methods
}

type HandleFactory struct{ C.Client }

func (c HandleFactory) NewHandle(ctx context.Context, params func(HandleFactory_newHandle_Params), opts ...C.CallOption) *HandleFactory_newHandle_Results_Promise {
	if c.Client == nil {
		return (*HandleFactory_newHandle_Results_Promise)(C.NewPipeline(C.ErrorAnswer(C.ErrNullClient)))
	}
	return (*HandleFactory_newHandle_Results_Promise)(C.NewPipeline(c.Client.Call(&C.Call{
		Ctx: ctx,
		Method: C.Method{

			InterfaceID:   0x8491a7fe75fe0bce,
			MethodID:      0,
			InterfaceName: "rpc/internal/testcapnp/test.capnp:HandleFactory",
			MethodName:    "newHandle",
		},
		ParamsSize: C.ObjectSize{DataSize: 0, PointerCount: 0},
		ParamsFunc: func(s C.Struct) { params(HandleFactory_newHandle_Params{Struct: s}) },
		Options:    C.NewCallOptions(opts),
	})))
}

type HandleFactory_Server interface {
	NewHandle(HandleFactory_newHandle) error
}

func HandleFactory_ServerToClient(s HandleFactory_Server) HandleFactory {
	c, _ := s.(server.Closer)
	return HandleFactory{Client: server.New(HandleFactory_Methods(nil, s), c)}
}

func HandleFactory_Methods(methods []server.Method, s HandleFactory_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: C.Method{

			InterfaceID:   0x8491a7fe75fe0bce,
			MethodID:      0,
			InterfaceName: "rpc/internal/testcapnp/test.capnp:HandleFactory",
			MethodName:    "newHandle",
		},
		Impl: func(c context.Context, opts C.CallOptions, p, r C.Struct) error {
			call := HandleFactory_newHandle{c, opts, HandleFactory_newHandle_Params{Struct: p}, HandleFactory_newHandle_Results{Struct: r}}
			return s.NewHandle(call)
		},
		ResultsSize: C.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// HandleFactory_newHandle holds the arguments for a server call to HandleFactory.newHandle.
type HandleFactory_newHandle struct {
	Ctx     context.Context
	Options C.CallOptions
	Params  HandleFactory_newHandle_Params
	Results HandleFactory_newHandle_Results
}

type HandleFactory_newHandle_Params struct{ C.Struct }

func NewHandleFactory_newHandle_Params(s *C.Segment) (HandleFactory_newHandle_Params, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return HandleFactory_newHandle_Params{}, err
	}
	return HandleFactory_newHandle_Params{st}, nil
}

func NewRootHandleFactory_newHandle_Params(s *C.Segment) (HandleFactory_newHandle_Params, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return HandleFactory_newHandle_Params{}, err
	}
	return HandleFactory_newHandle_Params{st}, nil
}

func ReadRootHandleFactory_newHandle_Params(msg *C.Message) (HandleFactory_newHandle_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return HandleFactory_newHandle_Params{}, err
	}
	st := C.ToStruct(root)
	return HandleFactory_newHandle_Params{st}, nil
}

// HandleFactory_newHandle_Params_List is a list of HandleFactory_newHandle_Params.
type HandleFactory_newHandle_Params_List struct{ C.List }

// NewHandleFactory_newHandle_Params creates a new list of HandleFactory_newHandle_Params.
func NewHandleFactory_newHandle_Params_List(s *C.Segment, sz int32) (HandleFactory_newHandle_Params_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return HandleFactory_newHandle_Params_List{}, err
	}
	return HandleFactory_newHandle_Params_List{l}, nil
}

func (s HandleFactory_newHandle_Params_List) At(i int) HandleFactory_newHandle_Params {
	return HandleFactory_newHandle_Params{s.List.Struct(i)}
}
func (s HandleFactory_newHandle_Params_List) Set(i int, v HandleFactory_newHandle_Params) {
	s.List.SetStruct(i, v.Struct)
}

type HandleFactory_newHandle_Params_Promise struct{ C.Pipeline }

func (p *HandleFactory_newHandle_Params_Promise) Struct() (HandleFactory_newHandle_Params, error) {
	s, err := p.Pipeline.Struct()
	return HandleFactory_newHandle_Params{s}, err
}

type HandleFactory_newHandle_Results struct{ C.Struct }

func NewHandleFactory_newHandle_Results(s *C.Segment) (HandleFactory_newHandle_Results, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HandleFactory_newHandle_Results{}, err
	}
	return HandleFactory_newHandle_Results{st}, nil
}

func NewRootHandleFactory_newHandle_Results(s *C.Segment) (HandleFactory_newHandle_Results, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HandleFactory_newHandle_Results{}, err
	}
	return HandleFactory_newHandle_Results{st}, nil
}

func ReadRootHandleFactory_newHandle_Results(msg *C.Message) (HandleFactory_newHandle_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return HandleFactory_newHandle_Results{}, err
	}
	st := C.ToStruct(root)
	return HandleFactory_newHandle_Results{st}, nil
}

func (s HandleFactory_newHandle_Results) Handle() Handle {
	p, err := s.Struct.Pointer(0)
	if err != nil {

		return Handle{}
	}
	c := C.ToInterface(p).Client()
	return Handle{Client: c}
}

func (s HandleFactory_newHandle_Results) SetHandle(v Handle) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	ci := seg.Message().AddCap(v.Client)
	return s.Struct.SetPointer(0, C.NewInterface(ci))
}

// HandleFactory_newHandle_Results_List is a list of HandleFactory_newHandle_Results.
type HandleFactory_newHandle_Results_List struct{ C.List }

// NewHandleFactory_newHandle_Results creates a new list of HandleFactory_newHandle_Results.
func NewHandleFactory_newHandle_Results_List(s *C.Segment, sz int32) (HandleFactory_newHandle_Results_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return HandleFactory_newHandle_Results_List{}, err
	}
	return HandleFactory_newHandle_Results_List{l}, nil
}

func (s HandleFactory_newHandle_Results_List) At(i int) HandleFactory_newHandle_Results {
	return HandleFactory_newHandle_Results{s.List.Struct(i)}
}
func (s HandleFactory_newHandle_Results_List) Set(i int, v HandleFactory_newHandle_Results) {
	s.List.SetStruct(i, v.Struct)
}

type HandleFactory_newHandle_Results_Promise struct{ C.Pipeline }

func (p *HandleFactory_newHandle_Results_Promise) Struct() (HandleFactory_newHandle_Results, error) {
	s, err := p.Pipeline.Struct()
	return HandleFactory_newHandle_Results{s}, err
}

func (p *HandleFactory_newHandle_Results_Promise) Handle() Handle {
	return NewHandle((*C.Pipeline)(p).GetPipeline(0).Client())
}

type Hanger struct{ C.Client }

func (c Hanger) Hang(ctx context.Context, params func(Hanger_hang_Params), opts ...C.CallOption) *Hanger_hang_Results_Promise {
	if c.Client == nil {
		return (*Hanger_hang_Results_Promise)(C.NewPipeline(C.ErrorAnswer(C.ErrNullClient)))
	}
	return (*Hanger_hang_Results_Promise)(C.NewPipeline(c.Client.Call(&C.Call{
		Ctx: ctx,
		Method: C.Method{

			InterfaceID:   0x8ae08044aae8a26e,
			MethodID:      0,
			InterfaceName: "rpc/internal/testcapnp/test.capnp:Hanger",
			MethodName:    "hang",
		},
		ParamsSize: C.ObjectSize{DataSize: 0, PointerCount: 0},
		ParamsFunc: func(s C.Struct) { params(Hanger_hang_Params{Struct: s}) },
		Options:    C.NewCallOptions(opts),
	})))
}

type Hanger_Server interface {
	Hang(Hanger_hang) error
}

func Hanger_ServerToClient(s Hanger_Server) Hanger {
	c, _ := s.(server.Closer)
	return Hanger{Client: server.New(Hanger_Methods(nil, s), c)}
}

func Hanger_Methods(methods []server.Method, s Hanger_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: C.Method{

			InterfaceID:   0x8ae08044aae8a26e,
			MethodID:      0,
			InterfaceName: "rpc/internal/testcapnp/test.capnp:Hanger",
			MethodName:    "hang",
		},
		Impl: func(c context.Context, opts C.CallOptions, p, r C.Struct) error {
			call := Hanger_hang{c, opts, Hanger_hang_Params{Struct: p}, Hanger_hang_Results{Struct: r}}
			return s.Hang(call)
		},
		ResultsSize: C.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// Hanger_hang holds the arguments for a server call to Hanger.hang.
type Hanger_hang struct {
	Ctx     context.Context
	Options C.CallOptions
	Params  Hanger_hang_Params
	Results Hanger_hang_Results
}

type Hanger_hang_Params struct{ C.Struct }

func NewHanger_hang_Params(s *C.Segment) (Hanger_hang_Params, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Hanger_hang_Params{}, err
	}
	return Hanger_hang_Params{st}, nil
}

func NewRootHanger_hang_Params(s *C.Segment) (Hanger_hang_Params, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Hanger_hang_Params{}, err
	}
	return Hanger_hang_Params{st}, nil
}

func ReadRootHanger_hang_Params(msg *C.Message) (Hanger_hang_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return Hanger_hang_Params{}, err
	}
	st := C.ToStruct(root)
	return Hanger_hang_Params{st}, nil
}

// Hanger_hang_Params_List is a list of Hanger_hang_Params.
type Hanger_hang_Params_List struct{ C.List }

// NewHanger_hang_Params creates a new list of Hanger_hang_Params.
func NewHanger_hang_Params_List(s *C.Segment, sz int32) (Hanger_hang_Params_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Hanger_hang_Params_List{}, err
	}
	return Hanger_hang_Params_List{l}, nil
}

func (s Hanger_hang_Params_List) At(i int) Hanger_hang_Params {
	return Hanger_hang_Params{s.List.Struct(i)}
}
func (s Hanger_hang_Params_List) Set(i int, v Hanger_hang_Params) { s.List.SetStruct(i, v.Struct) }

type Hanger_hang_Params_Promise struct{ C.Pipeline }

func (p *Hanger_hang_Params_Promise) Struct() (Hanger_hang_Params, error) {
	s, err := p.Pipeline.Struct()
	return Hanger_hang_Params{s}, err
}

type Hanger_hang_Results struct{ C.Struct }

func NewHanger_hang_Results(s *C.Segment) (Hanger_hang_Results, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Hanger_hang_Results{}, err
	}
	return Hanger_hang_Results{st}, nil
}

func NewRootHanger_hang_Results(s *C.Segment) (Hanger_hang_Results, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Hanger_hang_Results{}, err
	}
	return Hanger_hang_Results{st}, nil
}

func ReadRootHanger_hang_Results(msg *C.Message) (Hanger_hang_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return Hanger_hang_Results{}, err
	}
	st := C.ToStruct(root)
	return Hanger_hang_Results{st}, nil
}

// Hanger_hang_Results_List is a list of Hanger_hang_Results.
type Hanger_hang_Results_List struct{ C.List }

// NewHanger_hang_Results creates a new list of Hanger_hang_Results.
func NewHanger_hang_Results_List(s *C.Segment, sz int32) (Hanger_hang_Results_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Hanger_hang_Results_List{}, err
	}
	return Hanger_hang_Results_List{l}, nil
}

func (s Hanger_hang_Results_List) At(i int) Hanger_hang_Results {
	return Hanger_hang_Results{s.List.Struct(i)}
}
func (s Hanger_hang_Results_List) Set(i int, v Hanger_hang_Results) { s.List.SetStruct(i, v.Struct) }

type Hanger_hang_Results_Promise struct{ C.Pipeline }

func (p *Hanger_hang_Results_Promise) Struct() (Hanger_hang_Results, error) {
	s, err := p.Pipeline.Struct()
	return Hanger_hang_Results{s}, err
}

type CallOrder struct{ C.Client }

func (c CallOrder) GetCallSequence(ctx context.Context, params func(CallOrder_getCallSequence_Params), opts ...C.CallOption) *CallOrder_getCallSequence_Results_Promise {
	if c.Client == nil {
		return (*CallOrder_getCallSequence_Results_Promise)(C.NewPipeline(C.ErrorAnswer(C.ErrNullClient)))
	}
	return (*CallOrder_getCallSequence_Results_Promise)(C.NewPipeline(c.Client.Call(&C.Call{
		Ctx: ctx,
		Method: C.Method{

			InterfaceID:   0x92c5ca8314cdd2a5,
			MethodID:      0,
			InterfaceName: "rpc/internal/testcapnp/test.capnp:CallOrder",
			MethodName:    "getCallSequence",
		},
		ParamsSize: C.ObjectSize{DataSize: 8, PointerCount: 0},
		ParamsFunc: func(s C.Struct) { params(CallOrder_getCallSequence_Params{Struct: s}) },
		Options:    C.NewCallOptions(opts),
	})))
}

type CallOrder_Server interface {
	GetCallSequence(CallOrder_getCallSequence) error
}

func CallOrder_ServerToClient(s CallOrder_Server) CallOrder {
	c, _ := s.(server.Closer)
	return CallOrder{Client: server.New(CallOrder_Methods(nil, s), c)}
}

func CallOrder_Methods(methods []server.Method, s CallOrder_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: C.Method{

			InterfaceID:   0x92c5ca8314cdd2a5,
			MethodID:      0,
			InterfaceName: "rpc/internal/testcapnp/test.capnp:CallOrder",
			MethodName:    "getCallSequence",
		},
		Impl: func(c context.Context, opts C.CallOptions, p, r C.Struct) error {
			call := CallOrder_getCallSequence{c, opts, CallOrder_getCallSequence_Params{Struct: p}, CallOrder_getCallSequence_Results{Struct: r}}
			return s.GetCallSequence(call)
		},
		ResultsSize: C.ObjectSize{DataSize: 8, PointerCount: 0},
	})

	return methods
}

// CallOrder_getCallSequence holds the arguments for a server call to CallOrder.getCallSequence.
type CallOrder_getCallSequence struct {
	Ctx     context.Context
	Options C.CallOptions
	Params  CallOrder_getCallSequence_Params
	Results CallOrder_getCallSequence_Results
}

type CallOrder_getCallSequence_Params struct{ C.Struct }

func NewCallOrder_getCallSequence_Params(s *C.Segment) (CallOrder_getCallSequence_Params, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return CallOrder_getCallSequence_Params{}, err
	}
	return CallOrder_getCallSequence_Params{st}, nil
}

func NewRootCallOrder_getCallSequence_Params(s *C.Segment) (CallOrder_getCallSequence_Params, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return CallOrder_getCallSequence_Params{}, err
	}
	return CallOrder_getCallSequence_Params{st}, nil
}

func ReadRootCallOrder_getCallSequence_Params(msg *C.Message) (CallOrder_getCallSequence_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return CallOrder_getCallSequence_Params{}, err
	}
	st := C.ToStruct(root)
	return CallOrder_getCallSequence_Params{st}, nil
}

func (s CallOrder_getCallSequence_Params) Expected() uint32 {
	return s.Struct.Uint32(0)
}

func (s CallOrder_getCallSequence_Params) SetExpected(v uint32) {

	s.Struct.SetUint32(0, v)
}

// CallOrder_getCallSequence_Params_List is a list of CallOrder_getCallSequence_Params.
type CallOrder_getCallSequence_Params_List struct{ C.List }

// NewCallOrder_getCallSequence_Params creates a new list of CallOrder_getCallSequence_Params.
func NewCallOrder_getCallSequence_Params_List(s *C.Segment, sz int32) (CallOrder_getCallSequence_Params_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	if err != nil {
		return CallOrder_getCallSequence_Params_List{}, err
	}
	return CallOrder_getCallSequence_Params_List{l}, nil
}

func (s CallOrder_getCallSequence_Params_List) At(i int) CallOrder_getCallSequence_Params {
	return CallOrder_getCallSequence_Params{s.List.Struct(i)}
}
func (s CallOrder_getCallSequence_Params_List) Set(i int, v CallOrder_getCallSequence_Params) {
	s.List.SetStruct(i, v.Struct)
}

type CallOrder_getCallSequence_Params_Promise struct{ C.Pipeline }

func (p *CallOrder_getCallSequence_Params_Promise) Struct() (CallOrder_getCallSequence_Params, error) {
	s, err := p.Pipeline.Struct()
	return CallOrder_getCallSequence_Params{s}, err
}

type CallOrder_getCallSequence_Results struct{ C.Struct }

func NewCallOrder_getCallSequence_Results(s *C.Segment) (CallOrder_getCallSequence_Results, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return CallOrder_getCallSequence_Results{}, err
	}
	return CallOrder_getCallSequence_Results{st}, nil
}

func NewRootCallOrder_getCallSequence_Results(s *C.Segment) (CallOrder_getCallSequence_Results, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return CallOrder_getCallSequence_Results{}, err
	}
	return CallOrder_getCallSequence_Results{st}, nil
}

func ReadRootCallOrder_getCallSequence_Results(msg *C.Message) (CallOrder_getCallSequence_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return CallOrder_getCallSequence_Results{}, err
	}
	st := C.ToStruct(root)
	return CallOrder_getCallSequence_Results{st}, nil
}

func (s CallOrder_getCallSequence_Results) N() uint32 {
	return s.Struct.Uint32(0)
}

func (s CallOrder_getCallSequence_Results) SetN(v uint32) {

	s.Struct.SetUint32(0, v)
}

// CallOrder_getCallSequence_Results_List is a list of CallOrder_getCallSequence_Results.
type CallOrder_getCallSequence_Results_List struct{ C.List }

// NewCallOrder_getCallSequence_Results creates a new list of CallOrder_getCallSequence_Results.
func NewCallOrder_getCallSequence_Results_List(s *C.Segment, sz int32) (CallOrder_getCallSequence_Results_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	if err != nil {
		return CallOrder_getCallSequence_Results_List{}, err
	}
	return CallOrder_getCallSequence_Results_List{l}, nil
}

func (s CallOrder_getCallSequence_Results_List) At(i int) CallOrder_getCallSequence_Results {
	return CallOrder_getCallSequence_Results{s.List.Struct(i)}
}
func (s CallOrder_getCallSequence_Results_List) Set(i int, v CallOrder_getCallSequence_Results) {
	s.List.SetStruct(i, v.Struct)
}

type CallOrder_getCallSequence_Results_Promise struct{ C.Pipeline }

func (p *CallOrder_getCallSequence_Results_Promise) Struct() (CallOrder_getCallSequence_Results, error) {
	s, err := p.Pipeline.Struct()
	return CallOrder_getCallSequence_Results{s}, err
}

type Echoer struct{ C.Client }

func (c Echoer) Echo(ctx context.Context, params func(Echoer_echo_Params), opts ...C.CallOption) *Echoer_echo_Results_Promise {
	if c.Client == nil {
		return (*Echoer_echo_Results_Promise)(C.NewPipeline(C.ErrorAnswer(C.ErrNullClient)))
	}
	return (*Echoer_echo_Results_Promise)(C.NewPipeline(c.Client.Call(&C.Call{
		Ctx: ctx,
		Method: C.Method{

			InterfaceID:   0x841756c6a41b2a45,
			MethodID:      0,
			InterfaceName: "rpc/internal/testcapnp/test.capnp:Echoer",
			MethodName:    "echo",
		},
		ParamsSize: C.ObjectSize{DataSize: 0, PointerCount: 1},
		ParamsFunc: func(s C.Struct) { params(Echoer_echo_Params{Struct: s}) },
		Options:    C.NewCallOptions(opts),
	})))
}

func (c Echoer) GetCallSequence(ctx context.Context, params func(CallOrder_getCallSequence_Params), opts ...C.CallOption) *CallOrder_getCallSequence_Results_Promise {
	if c.Client == nil {
		return (*CallOrder_getCallSequence_Results_Promise)(C.NewPipeline(C.ErrorAnswer(C.ErrNullClient)))
	}
	return (*CallOrder_getCallSequence_Results_Promise)(C.NewPipeline(c.Client.Call(&C.Call{
		Ctx: ctx,
		Method: C.Method{

			InterfaceID:   0x92c5ca8314cdd2a5,
			MethodID:      0,
			InterfaceName: "rpc/internal/testcapnp/test.capnp:CallOrder",
			MethodName:    "getCallSequence",
		},
		ParamsSize: C.ObjectSize{DataSize: 8, PointerCount: 0},
		ParamsFunc: func(s C.Struct) { params(CallOrder_getCallSequence_Params{Struct: s}) },
		Options:    C.NewCallOptions(opts),
	})))
}

type Echoer_Server interface {
	Echo(Echoer_echo) error

	GetCallSequence(CallOrder_getCallSequence) error
}

func Echoer_ServerToClient(s Echoer_Server) Echoer {
	c, _ := s.(server.Closer)
	return Echoer{Client: server.New(Echoer_Methods(nil, s), c)}
}

func Echoer_Methods(methods []server.Method, s Echoer_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: C.Method{

			InterfaceID:   0x841756c6a41b2a45,
			MethodID:      0,
			InterfaceName: "rpc/internal/testcapnp/test.capnp:Echoer",
			MethodName:    "echo",
		},
		Impl: func(c context.Context, opts C.CallOptions, p, r C.Struct) error {
			call := Echoer_echo{c, opts, Echoer_echo_Params{Struct: p}, Echoer_echo_Results{Struct: r}}
			return s.Echo(call)
		},
		ResultsSize: C.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: C.Method{

			InterfaceID:   0x92c5ca8314cdd2a5,
			MethodID:      0,
			InterfaceName: "rpc/internal/testcapnp/test.capnp:CallOrder",
			MethodName:    "getCallSequence",
		},
		Impl: func(c context.Context, opts C.CallOptions, p, r C.Struct) error {
			call := CallOrder_getCallSequence{c, opts, CallOrder_getCallSequence_Params{Struct: p}, CallOrder_getCallSequence_Results{Struct: r}}
			return s.GetCallSequence(call)
		},
		ResultsSize: C.ObjectSize{DataSize: 8, PointerCount: 0},
	})

	return methods
}

// Echoer_echo holds the arguments for a server call to Echoer.echo.
type Echoer_echo struct {
	Ctx     context.Context
	Options C.CallOptions
	Params  Echoer_echo_Params
	Results Echoer_echo_Results
}

type Echoer_echo_Params struct{ C.Struct }

func NewEchoer_echo_Params(s *C.Segment) (Echoer_echo_Params, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Echoer_echo_Params{}, err
	}
	return Echoer_echo_Params{st}, nil
}

func NewRootEchoer_echo_Params(s *C.Segment) (Echoer_echo_Params, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Echoer_echo_Params{}, err
	}
	return Echoer_echo_Params{st}, nil
}

func ReadRootEchoer_echo_Params(msg *C.Message) (Echoer_echo_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return Echoer_echo_Params{}, err
	}
	st := C.ToStruct(root)
	return Echoer_echo_Params{st}, nil
}

func (s Echoer_echo_Params) Cap() CallOrder {
	p, err := s.Struct.Pointer(0)
	if err != nil {

		return CallOrder{}
	}
	c := C.ToInterface(p).Client()
	return CallOrder{Client: c}
}

func (s Echoer_echo_Params) SetCap(v CallOrder) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	ci := seg.Message().AddCap(v.Client)
	return s.Struct.SetPointer(0, C.NewInterface(ci))
}

// Echoer_echo_Params_List is a list of Echoer_echo_Params.
type Echoer_echo_Params_List struct{ C.List }

// NewEchoer_echo_Params creates a new list of Echoer_echo_Params.
func NewEchoer_echo_Params_List(s *C.Segment, sz int32) (Echoer_echo_Params_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Echoer_echo_Params_List{}, err
	}
	return Echoer_echo_Params_List{l}, nil
}

func (s Echoer_echo_Params_List) At(i int) Echoer_echo_Params {
	return Echoer_echo_Params{s.List.Struct(i)}
}
func (s Echoer_echo_Params_List) Set(i int, v Echoer_echo_Params) { s.List.SetStruct(i, v.Struct) }

type Echoer_echo_Params_Promise struct{ C.Pipeline }

func (p *Echoer_echo_Params_Promise) Struct() (Echoer_echo_Params, error) {
	s, err := p.Pipeline.Struct()
	return Echoer_echo_Params{s}, err
}

func (p *Echoer_echo_Params_Promise) Cap() CallOrder {
	return NewCallOrder((*C.Pipeline)(p).GetPipeline(0).Client())
}

type Echoer_echo_Results struct{ C.Struct }

func NewEchoer_echo_Results(s *C.Segment) (Echoer_echo_Results, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Echoer_echo_Results{}, err
	}
	return Echoer_echo_Results{st}, nil
}

func NewRootEchoer_echo_Results(s *C.Segment) (Echoer_echo_Results, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Echoer_echo_Results{}, err
	}
	return Echoer_echo_Results{st}, nil
}

func ReadRootEchoer_echo_Results(msg *C.Message) (Echoer_echo_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return Echoer_echo_Results{}, err
	}
	st := C.ToStruct(root)
	return Echoer_echo_Results{st}, nil
}

func (s Echoer_echo_Results) Cap() CallOrder {
	p, err := s.Struct.Pointer(0)
	if err != nil {

		return CallOrder{}
	}
	c := C.ToInterface(p).Client()
	return CallOrder{Client: c}
}

func (s Echoer_echo_Results) SetCap(v CallOrder) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	ci := seg.Message().AddCap(v.Client)
	return s.Struct.SetPointer(0, C.NewInterface(ci))
}

// Echoer_echo_Results_List is a list of Echoer_echo_Results.
type Echoer_echo_Results_List struct{ C.List }

// NewEchoer_echo_Results creates a new list of Echoer_echo_Results.
func NewEchoer_echo_Results_List(s *C.Segment, sz int32) (Echoer_echo_Results_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Echoer_echo_Results_List{}, err
	}
	return Echoer_echo_Results_List{l}, nil
}

func (s Echoer_echo_Results_List) At(i int) Echoer_echo_Results {
	return Echoer_echo_Results{s.List.Struct(i)}
}
func (s Echoer_echo_Results_List) Set(i int, v Echoer_echo_Results) { s.List.SetStruct(i, v.Struct) }

type Echoer_echo_Results_Promise struct{ C.Pipeline }

func (p *Echoer_echo_Results_Promise) Struct() (Echoer_echo_Results, error) {
	s, err := p.Pipeline.Struct()
	return Echoer_echo_Results{s}, err
}

func (p *Echoer_echo_Results_Promise) Cap() CallOrder {
	return NewCallOrder((*C.Pipeline)(p).GetPipeline(0).Client())
}

type Adder struct{ C.Client }

func (c Adder) Add(ctx context.Context, params func(Adder_add_Params), opts ...C.CallOption) *Adder_add_Results_Promise {
	if c.Client == nil {
		return (*Adder_add_Results_Promise)(C.NewPipeline(C.ErrorAnswer(C.ErrNullClient)))
	}
	return (*Adder_add_Results_Promise)(C.NewPipeline(c.Client.Call(&C.Call{
		Ctx: ctx,
		Method: C.Method{

			InterfaceID:   0x8f9cac550b1bf41f,
			MethodID:      0,
			InterfaceName: "rpc/internal/testcapnp/test.capnp:Adder",
			MethodName:    "add",
		},
		ParamsSize: C.ObjectSize{DataSize: 8, PointerCount: 0},
		ParamsFunc: func(s C.Struct) { params(Adder_add_Params{Struct: s}) },
		Options:    C.NewCallOptions(opts),
	})))
}

type Adder_Server interface {
	Add(Adder_add) error
}

func Adder_ServerToClient(s Adder_Server) Adder {
	c, _ := s.(server.Closer)
	return Adder{Client: server.New(Adder_Methods(nil, s), c)}
}

func Adder_Methods(methods []server.Method, s Adder_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: C.Method{

			InterfaceID:   0x8f9cac550b1bf41f,
			MethodID:      0,
			InterfaceName: "rpc/internal/testcapnp/test.capnp:Adder",
			MethodName:    "add",
		},
		Impl: func(c context.Context, opts C.CallOptions, p, r C.Struct) error {
			call := Adder_add{c, opts, Adder_add_Params{Struct: p}, Adder_add_Results{Struct: r}}
			return s.Add(call)
		},
		ResultsSize: C.ObjectSize{DataSize: 8, PointerCount: 0},
	})

	return methods
}

// Adder_add holds the arguments for a server call to Adder.add.
type Adder_add struct {
	Ctx     context.Context
	Options C.CallOptions
	Params  Adder_add_Params
	Results Adder_add_Results
}

type Adder_add_Params struct{ C.Struct }

func NewAdder_add_Params(s *C.Segment) (Adder_add_Params, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return Adder_add_Params{}, err
	}
	return Adder_add_Params{st}, nil
}

func NewRootAdder_add_Params(s *C.Segment) (Adder_add_Params, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return Adder_add_Params{}, err
	}
	return Adder_add_Params{st}, nil
}

func ReadRootAdder_add_Params(msg *C.Message) (Adder_add_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return Adder_add_Params{}, err
	}
	st := C.ToStruct(root)
	return Adder_add_Params{st}, nil
}

func (s Adder_add_Params) A() int32 {
	return int32(s.Struct.Uint32(0))
}

func (s Adder_add_Params) SetA(v int32) {

	s.Struct.SetUint32(0, uint32(v))
}

func (s Adder_add_Params) B() int32 {
	return int32(s.Struct.Uint32(4))
}

func (s Adder_add_Params) SetB(v int32) {

	s.Struct.SetUint32(4, uint32(v))
}

// Adder_add_Params_List is a list of Adder_add_Params.
type Adder_add_Params_List struct{ C.List }

// NewAdder_add_Params creates a new list of Adder_add_Params.
func NewAdder_add_Params_List(s *C.Segment, sz int32) (Adder_add_Params_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	if err != nil {
		return Adder_add_Params_List{}, err
	}
	return Adder_add_Params_List{l}, nil
}

func (s Adder_add_Params_List) At(i int) Adder_add_Params     { return Adder_add_Params{s.List.Struct(i)} }
func (s Adder_add_Params_List) Set(i int, v Adder_add_Params) { s.List.SetStruct(i, v.Struct) }

type Adder_add_Params_Promise struct{ C.Pipeline }

func (p *Adder_add_Params_Promise) Struct() (Adder_add_Params, error) {
	s, err := p.Pipeline.Struct()
	return Adder_add_Params{s}, err
}

type Adder_add_Results struct{ C.Struct }

func NewAdder_add_Results(s *C.Segment) (Adder_add_Results, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return Adder_add_Results{}, err
	}
	return Adder_add_Results{st}, nil
}

func NewRootAdder_add_Results(s *C.Segment) (Adder_add_Results, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return Adder_add_Results{}, err
	}
	return Adder_add_Results{st}, nil
}

func ReadRootAdder_add_Results(msg *C.Message) (Adder_add_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return Adder_add_Results{}, err
	}
	st := C.ToStruct(root)
	return Adder_add_Results{st}, nil
}

func (s Adder_add_Results) Result() int32 {
	return int32(s.Struct.Uint32(0))
}

func (s Adder_add_Results) SetResult(v int32) {

	s.Struct.SetUint32(0, uint32(v))
}

// Adder_add_Results_List is a list of Adder_add_Results.
type Adder_add_Results_List struct{ C.List }

// NewAdder_add_Results creates a new list of Adder_add_Results.
func NewAdder_add_Results_List(s *C.Segment, sz int32) (Adder_add_Results_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	if err != nil {
		return Adder_add_Results_List{}, err
	}
	return Adder_add_Results_List{l}, nil
}

func (s Adder_add_Results_List) At(i int) Adder_add_Results {
	return Adder_add_Results{s.List.Struct(i)}
}
func (s Adder_add_Results_List) Set(i int, v Adder_add_Results) { s.List.SetStruct(i, v.Struct) }

type Adder_add_Results_Promise struct{ C.Pipeline }

func (p *Adder_add_Results_Promise) Struct() (Adder_add_Results, error) {
	s, err := p.Pipeline.Struct()
	return Adder_add_Results{s}, err
}
