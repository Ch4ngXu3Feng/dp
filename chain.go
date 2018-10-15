package dp

import "fmt"

type IHandler interface {
    Handle()
}

type Handler struct {
    IHandler
    next IHandler
}

func (h *Handler) Init() {
    h.IHandler = h
    fmt.Printf("Handler Init\n")
}

func (h *Handler) Handle() {
    if h.next != nil {
        h.next.Handle()
    }
}

func NewHandler(handler IHandler) *Handler {
    h := &Handler{
        next: handler,
    }
    h.Init()
    return h
}

type XxxHandler struct {
    Handler
}

func (h *XxxHandler) Init() {
    h.Handler.Init()
    h.IHandler = h
    fmt.Printf("XxxHandler Init\n")
}

func (h *XxxHandler) Handle() {
    fmt.Printf("XxxHandler Handle\n")
    h.Handler.Handle()
}

func NewXxxHandler(handler IHandler) *XxxHandler {
    h := &XxxHandler{
        Handler: Handler{
            next: handler,
        },
    }
    h.Init()
    return h
}

type ZzzHandler struct {
    Handler
}

func (h *ZzzHandler) Init() {
    h.Handler.Init()
    h.IHandler = h
    fmt.Printf("ZzzHandler Init\n")
}

func (h *ZzzHandler) Handle() {
    fmt.Printf("ZzzHandler Handle\n")
    h.Handler.Handle()
}

func NewZzzHandler(handler IHandler) *ZzzHandler {
    h := &ZzzHandler{
        Handler: Handler{
            next: handler,
        },
    }
    h.Init()
    return h
}
