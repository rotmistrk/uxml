# uxml (micro-XML)


### micro-XML

```go
doc := uxml.NewDoc("root", uxml.Attrib{"ver": "0.1.0"})
root := doc.RootElem()
first := root.AddElem("first", uxml.Attrin{"ort":"one",, "class": "theClass"}).AddText("someText")
second := root.AddElem("second", uxml.Attrin{"ort":"two",, "class": "theClass"}).AddText("otherText")
second.AddElem("an-elem").AddText("some value")
for _, child := range root.Children() {
	if child.Type() == uxml.NodeElem {
        child.(Elem).Attrib()["visible"] = "true"
    }
}
doc.WriteTo(myWriter)
```

