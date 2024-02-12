package node

import (
	"syscall/js"

	"github.com/beavorp/gofui/core"
)

type (
	Body   struct{ *Node }
	Div    struct{ *Node }
	Aside  struct{ *Node }
	Header struct{ *Node }
	Footer struct{ *Node }
	H1     struct{ *Node }
	H2     struct{ *Node }
	H3     struct{ *Node }
	H4     struct{ *Node }
	H5     struct{ *Node }
	H6     struct{ *Node }
	P      struct{ *Node }
	Nav    struct{ *Node }
	I      struct{ *Node }
	Button struct{ *Node }
	DD     struct{ *Node }
	DT     struct{ *Node }
	DL     struct{ *Node }
	UL     struct{ *Node }
	LI     struct{ *Node }
	OL     struct{ *Node }
	A      struct{ *Node }
	Span   struct{ *Node }
)

func String(content string) *js.Value {
	v := js.ValueOf(content)
	return &v
}

func NewBody() *Body {
	el := js.Global().Get("document").Get("body")
	return &Body{NewNode(&el)}
}

func NewDiv() *Div {
	el := core.CreateElement("div")
	return &Div{NewNode(&el)}
}

func NewAside() *Aside {
	el := core.CreateElement("aside")
	return &Aside{NewNode(&el)}
}

func NewHeader() *Header {
	el := core.CreateElement("header")
	return &Header{NewNode(&el)}
}

func NewFooter() *Footer {
	el := core.CreateElement("footer")
	return &Footer{NewNode(&el)}
}

func NewH1() *H1 {
	el := core.CreateElement("h1")
	return &H1{NewNode(&el)}
}

func NewH2() *H2 {
	el := core.CreateElement("h2")
	return &H2{NewNode(&el)}
}

func NewH3() *H3 {
	el := core.CreateElement("h3")
	return &H3{NewNode(&el)}
}

func NewH4() *H4 {
	el := core.CreateElement("h4")
	return &H4{NewNode(&el)}
}

func NewH5() *H5 {
	el := core.CreateElement("h5")
	return &H5{NewNode(&el)}
}

func NewH6() *H6 {
	el := core.CreateElement("h6")
	return &H6{NewNode(&el)}
}

func NewP() *P {
	el := core.CreateElement("p")
	return &P{NewNode(&el)}
}

func NewNav() *Nav {
	el := core.CreateElement("nav")
	return &Nav{NewNode(&el)}
}

func NewI() *I {
	el := core.CreateElement("i")
	return &I{NewNode(&el)}
}

func NewButton() *Button {
	el := core.CreateElement("button")
	return &Button{NewNode(&el)}
}

func NewDD() *DD {
	el := core.CreateElement("dd")
	return &DD{NewNode(&el)}
}

func NewDT() *DT {
	el := core.CreateElement("dt")
	return &DT{NewNode(&el)}
}

func NewDL() *DL {
	el := core.CreateElement("dl")
	return &DL{NewNode(&el)}
}

func NewUL() *UL {
	el := core.CreateElement("ul")
	return &UL{NewNode(&el)}
}

func NewLI() *LI {
	el := core.CreateElement("li")
	return &LI{NewNode(&el)}
}

func NewOL() *OL {
	el := core.CreateElement("ol")
	return &OL{NewNode(&el)}
}

func NewA() *A {
	el := core.CreateElement("a")
	return &A{NewNode(&el)}
}

func NewSpan() *Span {
	el := core.CreateElement("span")
	return &Span{NewNode(&el)}
}
