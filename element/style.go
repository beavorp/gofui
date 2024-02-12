package element

import "syscall/js"

type Style struct {
	el      js.Value
	elStyle js.Value
}

func NewStyle(el js.Value) *Style {
	return &Style{
		el:      el,
		elStyle: el.Get("style"),
	}
}

func (s Style) SetBackgroundColor(color string) {
	s.elStyle.Set("backgroundColor", color)
}

func (s Style) SetColor(color string) {
	s.elStyle.Set("color", color)
}

func (s Style) SetWidth(width string) {
	s.elStyle.Set("width", width)
}

func (s Style) SetHeight(height string) {
	s.elStyle.Set("height", height)
}

func (s Style) H_PaMa0() {
	s.elStyle.Set("padding", 0)
	s.elStyle.Set("margin", 0)
}

func (s Style) SetDisplay(display string) {
	s.elStyle.Set("display", display)
}

func (s Style) SetOverflow(overflow string) {
	s.elStyle.Set("overflow", overflow)
}

func (s Style) SetClassName(className string) {
	s.el.Set("className", className)
}

func (s Style) AddClassName(className string) {
	s.el.Get("classList").Call("add", className)
}

func (s Style) RemoveClassName(className string) {
	s.el.Get("classList").Call("remove", className)
}
