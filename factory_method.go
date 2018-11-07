package dp

import "fmt"

type ICreator interface {
    CreateDocument(name string) *Document
    NewDocument(name string)
}

type Creator struct {
    ICreator
    docs map[string]*Document
}

func (c *Creator) Init() {
    fmt.Printf("Creator Init\n")
    c.ICreator = c
}

func (c *Creator) CreateDocument(name string) *Document {
    if _, ok := c.ICreator.(*Creator); ok {
        panic("Creator CreateDocument oops\n")
    }
    return c.ICreator.CreateDocument(name)
}

func (c *Creator) NewDocument(name string) {
    d := c.CreateDocument(name)
    c.docs[name] = d
}

func (c *Creator) Show() {
    for _, d := range c.docs {
        d.Show()
    }
}

func NewCreator() *Creator {
    c := &Creator{
        docs: map[string]*Document{},
    }
    c.Init()
    return c
}

type XxxCreator struct {
    Creator
}

func (c *XxxCreator) Init() {
    fmt.Printf("XxxCreator Init\n")
    c.Creator.Init()
    c.ICreator = c
}

func (c *XxxCreator) CreateDocument(name string) *Document {
    return &NewXxxDocument(name).Document
}

func NewXxxCreator() *XxxCreator {
    c := &XxxCreator{
        Creator: Creator{
            docs: map[string]*Document{},
        },
    }
    c.Init()
    return c
}

type ZzzCreator struct {
    Creator
}

func (c *ZzzCreator) Init() {
    fmt.Printf("ZzzCreator Init\n")
    c.Creator.Init()
    c.ICreator = c
}

func (c *ZzzCreator) CreateDocument(name string) *Document {
    return &NewZzzDocument(name).Document
}

func NewZzzCreator() *ZzzCreator {
    c := &ZzzCreator{
        Creator: Creator{
            docs: map[string]*Document{},
        },
    }
    c.Init()
    return c
}

type IDocument interface {
    Show()
}

type Document struct {
    IDocument
    Name string
}

func (d *Document) Init() {
    fmt.Printf("Document Init\n")
    d.IDocument = d
}

func (d *Document) Show() {
    if _, ok := d.IDocument.(*Document); ok {
        panic("Document Show oops\n")
    }
    d.IDocument.Show()
}

func NewDocument() *Document {
    d := &Document{}
    d.Init()
    return d
}

type XxxDocument struct {
    Document
}

func (d *XxxDocument) Init() {
    fmt.Printf("XxxDocument Init\n")
    d.Document.Init()
    d.IDocument = d
}

func (d *XxxDocument) Show() {
    fmt.Printf("XxxDocument Show: %s\n", d.Name)
}

func NewXxxDocument(name string) *XxxDocument {
    d := &XxxDocument{
        Document: Document{
            Name: name,
        },
    }
    d.Init()
    return d
}

type ZzzDocument struct {
    Document
}

func (d *ZzzDocument) Init() {
    fmt.Printf("ZzzDocument Init\n")
    d.Document.Init()
    d.IDocument = d
}

func (d *ZzzDocument) Show() {
    fmt.Printf("ZzzDocument Show: %s\n", d.Name)
}

func NewZzzDocument(name string) *ZzzDocument {
    d := &ZzzDocument{
        Document: Document{
            Name: name,
        },
    }
    d.Init()
    return d
}
