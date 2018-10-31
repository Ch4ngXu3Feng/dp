package dp

import (
    "fmt"
)

type Product struct {
    header string
    body string
}

func (p *Product) Init() {
    fmt.Printf("Product Init\n")
}

func (p *Product) SetHeader(text string) {
    p.header = text
}

func (p *Product) SetBody(text string) {
    p.body = text
}

func (p *Product) Show() {
    fmt.Printf("Product Show header: %s, body: %s\n", p.header, p.body)
}

func NewProduct() *Product {
    p := &Product{}
    p.Init()
    return p
}

type Director struct {
}

func (d *Director) Init() {
    fmt.Printf("Director Init\n")
}

func (d *Director) Create(b IBuilder) *Product {
    p := NewProduct()
    b.CreateHeader(p)
    b.CreateBody(p)
    return p
}

func NewDirector() *Director {
    d := &Director{}
    d.Init()
    return d
}

type IBuilder interface {
    CreateHeader(*Product)
    CreateBody(*Product)
}

type Builder struct {
    IBuilder
}

func (b *Builder) Init() {
    b.IBuilder = b
    fmt.Printf("Builder Init\n")
}

func (b *Builder) CreateHeader(p *Product) {
    if _, ok := b.IBuilder.(*Builder); ok {
        panic("Builder CreateHeader oops\n")
    }
    b.IBuilder.CreateHeader(p)
}

func (b *Builder) CreateBody(p *Product) {
    if _, ok := b.IBuilder.(*Builder); ok {
        panic("Builder CreateBody oops\n")
    }
    b.IBuilder.CreateBody(p)
}

func NewBuilder() *Builder {
    b := &Builder{}
    b.Init()
    return b
}

type XxxBuilder struct {
    Builder
}

func (b *XxxBuilder) Init() {
    b.Builder.Init()
    b.IBuilder = b
    fmt.Printf("XxxBuilder Init\n")
}

func (b *XxxBuilder) CreateHeader(p *Product) {
    fmt.Printf("XxxBuilder Header\n")
    p.SetHeader("XxxHeader")
}

func (b *XxxBuilder) CreateBody(p *Product) {
    fmt.Printf("XxxBuilder Body\n")
    p.SetBody("XxxBody")
}

func NewXxxBuilder() *XxxBuilder {
    b := &XxxBuilder{}
    b.Init()
    return b
}

type ZzzBuilder struct {
    Builder
}

func (b *ZzzBuilder) Init() {
    b.Builder.Init()
    b.IBuilder = b
    fmt.Printf("ZzzBuilder Init\n")
}

func (b *ZzzBuilder) CreateHeader(p *Product) {
    fmt.Printf("ZzzBuilder Header\n")
    p.SetHeader("ZzzHeader")
}

func (b *ZzzBuilder) CreateBody(p *Product) {
    fmt.Printf("ZzzBuilder Body\n")
    p.SetBody("ZzzBody")
}

func NewZzzBuilder() *ZzzBuilder {
    b := &ZzzBuilder{}
    b.Init()
    return b
}
