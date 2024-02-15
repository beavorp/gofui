package node

import (
	"syscall/js"

	"github.com/beavorp/gofui/core"
)

type (
	Body    struct{ *Node[*Body] }
	Div     struct{ *Node[*Div] }
	Aside   struct{ *Node[*Aside] }
	Header  struct{ *Node[*Header] }
	Footer  struct{ *Node[*Footer] }
	H1      struct{ *Node[*H1] }
	H2      struct{ *Node[*H2] }
	H3      struct{ *Node[*H3] }
	H4      struct{ *Node[*H4] }
	H5      struct{ *Node[*H5] }
	H6      struct{ *Node[*H6] }
	P       struct{ *Node[*P] }
	Nav     struct{ *Node[*Nav] }
	I       struct{ *Node[*I] }
	Button  struct{ *Node[*Button] }
	DD      struct{ *Node[*DD] }
	DT      struct{ *Node[*DT] }
	DL      struct{ *Node[*DL] }
	UL      struct{ *Node[*UL] }
	LI      struct{ *Node[*LI] }
	OL      struct{ *Node[*OL] }
	A       struct{ *Node[*A] }
	Span    struct{ *Node[*Span] }
	Table   struct{ *Node[*Table] }
	Caption struct{ *Node[*Caption] }
	THead   struct{ *Node[*THead] }
	TBody   struct{ *Node[*TBody] }
	TR      struct{ *Node[*TR] }
	TH      struct{ *Node[*TH] }
	TD      struct{ *Node[*TD] }
	TFoot   struct{ *Node[*TFoot] }
	Input   struct{ *Node[*Input] }
)

func String(content string) *js.Value {
	v := js.ValueOf(content)
	return &v
}

func NewBody() (p *Body) {
	el := js.Global().Get("document").Get("body")
	p = &Body{}
	p.Node = NewNode(p, &el)
	return
}

func NewDiv() (p *Div) {
	el := core.CreateElement("div")
	p = &Div{}
	p.Node = NewNode(p, &el)
	return
}

func NewAside() (p *Aside) {
	el := core.CreateElement("aside")
	p = &Aside{}
	p.Node = NewNode(p, &el)
	return
}

func NewHeader() (p *Header) {
	el := core.CreateElement("header")
	p = &Header{}
	p.Node = NewNode(p, &el)
	return
}

func NewFooter() (p *Footer) {
	el := core.CreateElement("footer")
	p = &Footer{}
	p.Node = NewNode(p, &el)
	return
}

func NewH1() (p *H1) {
	el := core.CreateElement("h1")
	p = &H1{}
	p.Node = NewNode(p, &el)
	return
}

func NewH2() (p *H2) {
	el := core.CreateElement("h2")
	p = &H2{}
	p.Node = NewNode(p, &el)
	return
}

func NewH3() (p *H3) {
	el := core.CreateElement("h3")
	p = &H3{}
	p.Node = NewNode(p, &el)
	return
}

func NewH4() (p *H4) {
	el := core.CreateElement("h4")
	p = &H4{}
	p.Node = NewNode(p, &el)
	return
}

func NewH5() (p *H5) {
	el := core.CreateElement("h5")
	p = &H5{}
	p.Node = NewNode(p, &el)
	return
}

func NewH6() (p *H6) {
	el := core.CreateElement("h6")
	p = &H6{}
	p.Node = NewNode(p, &el)
	return
}

func NewP() (p *P) {
	el := core.CreateElement("p")
	p = &P{}
	p.Node = NewNode(p, &el)
	return
}

func NewNav() (p *Nav) {
	el := core.CreateElement("nav")
	p = &Nav{}
	p.Node = NewNode(p, &el)
	return
}

func NewI() (p *I) {
	el := core.CreateElement("i")
	p = &I{}
	p.Node = NewNode(p, &el)
	return
}

func NewButton() (p *Button) {
	el := core.CreateElement("button")
	p = &Button{}
	p.Node = NewNode(p, &el)
	return
}

func NewDD() (p *DD) {
	el := core.CreateElement("dd")
	p = &DD{}
	p.Node = NewNode(p, &el)
	return
}

func NewDT() (p *DT) {
	el := core.CreateElement("dt")
	p = &DT{}
	p.Node = NewNode(p, &el)
	return
}

func NewDL() (p *DL) {
	el := core.CreateElement("dl")
	p = &DL{}
	p.Node = NewNode(p, &el)
	return
}

func NewUL() (p *UL) {
	el := core.CreateElement("ul")
	p = &UL{}
	p.Node = NewNode(p, &el)
	return
}

func NewLI() (p *LI) {
	el := core.CreateElement("li")
	p = &LI{}
	p.Node = NewNode(p, &el)
	return
}

func NewOL() (p *OL) {
	el := core.CreateElement("ol")
	p = &OL{}
	p.Node = NewNode(p, &el)
	return
}

func NewA() (p *A) {
	el := core.CreateElement("a")
	p = &A{}
	p.Node = NewNode(p, &el)
	return
}

func NewSpan() (p *Span) {
	el := core.CreateElement("span")
	p = &Span{}
	p.Node = NewNode(p, &el)
	return
}

func NewTable() (p *Table) {
	el := core.CreateElement("table")
	p = &Table{}
	p.Node = NewNode(p, &el)
	return
}

func NewCaption() (p *Caption) {
	el := core.CreateElement("caption")
	p = &Caption{}
	p.Node = NewNode(p, &el)
	return
}

func NewTHead() (p *THead) {
	el := core.CreateElement("thead")
	p = &THead{}
	p.Node = NewNode(p, &el)
	return
}

func NewTBody() (p *TBody) {
	el := core.CreateElement("tbody")
	p = &TBody{}
	p.Node = NewNode(p, &el)
	return
}

func NewTR() (p *TR) {
	el := core.CreateElement("tr")
	p = &TR{}
	p.Node = NewNode(p, &el)
	return
}

func NewTH() (p *TH) {
	el := core.CreateElement("th")
	p = &TH{}
	p.Node = NewNode(p, &el)
	return
}

func NewTD() (p *TD) {
	el := core.CreateElement("td")
	p = &TD{}
	p.Node = NewNode(p, &el)
	return
}

func NewTFoot() (p *TFoot) {
	el := core.CreateElement("tfoot")
	p = &TFoot{}
	p.Node = NewNode(p, &el)
	return
}

func NewInput() (p *Input) {
	el := core.CreateElement("input")
	p = &Input{}
	p.Node = NewNode(p, &el)
	return
}

const (
	TypeText     = "text"
	TypePassword = "password"
	TypeEmail    = "email"
	TypeNumber   = "number"
	TypeDate     = "date"
)

func (i *Input) Type(t string) *Input {
	i.Attr("type", t)
	return i
}
