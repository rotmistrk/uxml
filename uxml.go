package uxml

import "io"

type NodeType int

const (
	NodeDoc  = NodeType(0)
	NodeElem = NodeType(1)
	NodeText = NodeType(2)
)

type Node interface {
	Parent() Elem
	Type() NodeType
	WriteTo(writer io.Writer) (int64, error)
}

type Attrib map[string]string

type Elem interface {
	Node
	Tag() string
	Attrib() Attrib
	Children() []Node
	AddElem(string, ...Attrib) Elem
	AddText(string) Text
}

type Text interface {
	Node
	Text() string
}

type Doc interface {
	RootElem() Elem
	WriteTo(writer io.Writer) (int64, error)
}
