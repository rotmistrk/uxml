package uxml

import "fmt"

type text struct {
	parent   Elem
	contents string
}

type elem struct {
	parent   *elem
	tag      string
	attrib   Attrib
	children []Node
}

type doc struct {
	rootElem *elem
}

func (t *text) Parent() Elem {
	return t.parent
}

func (t *text) Type() NodeType {
	return NodeText
}

func (t *text) Text() string {
	return t.contents
}

func (e *elem) Parent() Elem {
	return e.parent
}

func (e *elem) Type() NodeType {
	return NodeElem
}

func (e *elem) Tag() string {
	return e.tag
}

func (e *elem) Attrib() Attrib {
	return e.attrib
}

func (e *elem) Children() []Node {
	return e.children
}

func (e *elem) AddElem(tag string, attr ...Attrib) Elem {
	child := &elem{
		parent:   e,
		tag:      tag,
		attrib:   make(Attrib),
		children: make([]Node, 0, 0),
	}
	copyAttr(attr, child)
	e.children = append(e.children, child)
	return child
}

func copyAttr(attr []Attrib, e *elem) {
	for _, a := range attr {
		for k, v := range a {
			e.attrib[k] = v
		}
	}
}

func (e *elem) AddText(contents ...any) Text {
	child := &text{
		parent:   e,
		contents: fmt.Sprint(contents...),
	}
	e.children = append(e.children, child)
	return child
}

func (e *elem) AddTextf(format string, contents ...any) Text {
	child := &text{
		parent:   e,
		contents: fmt.Sprintf(format, contents...),
	}
	e.children = append(e.children, child)
	return child
}

func (doc *doc) RootElem() Elem {
	return doc.rootElem
}

func NewDoc(rootTag string, attr ...Attrib) Doc {
	out := &doc{
		rootElem: &elem{
			parent:   nil,
			tag:      rootTag,
			attrib:   make(Attrib),
			children: make([]Node, 0, 0),
		},
	}
	copyAttr(attr, out.rootElem)
	return out
}
