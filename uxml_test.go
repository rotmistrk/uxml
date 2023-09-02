package uxml

import (
	"bytes"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		generate func() Doc
		wantStr  string
	}{
		{
			name: "single tag",
			generate: func() Doc {
				return NewDoc("single")
			},
			wantStr: "<single/>",
		},
		{
			name: "single tag with attrs",
			generate: func() Doc {
				return NewDoc("single", Attrib{
					"one":   "une",
					"two":   "deux",
					"three": "trois",
				})
			},
			wantStr: "<single one=\"une\" three=\"trois\" two=\"deux\"/>",
		},
		{
			name: "single tag with text",
			generate: func() Doc {
				doc := NewDoc("single", Attrib{
					"attrib": "<whisper>\"don't let chip & dail sing\"",
				})
				doc.RootElem().AddText("some very \"fancy\" text")
				return doc
			},
			wantStr: "<single attrib=\"&lt;whisper&gt;&quot;don&apos;t let chip &amp; dail sing&quot;\">some very &quot;fancy&quot; text</single>",
		},
		{
			name: "with child elements",
			generate: func() Doc {
				doc := NewDoc("root", Attrib{
					"type": "root",
				})
				lst := doc.RootElem().AddElem("list")
				lst.AddElem("li", Attrib{"ord": 1}).AddText("une")
				lst.AddElem("li", Attrib{"ord": "2"}).AddText("deux")
				lst.AddElem("li", Attrib{"ord": 3.5}).AddText("trois")
				par := doc.RootElem().AddElem("par", Attrib{"type": "regular"})
				par.AddText("This is ")
				par.AddElem("b").AddText("formatted")
				par.AddText(" text")
				return doc
			},
			wantStr: "<root type=\"root\"><list><li ord=\"1\">une</li><li ord=\"2\">deux</li><li ord=\"3.5\">trois</li></list><par type=\"regular\">This is <b>formatted</b> text</par></root>",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buf := bytes.Buffer{}
			doc := test.generate()
			_, _ = doc.WriteTo(&buf)
			if got := buf.String(); got != test.wantStr {
				t.Errorf("%v: got = %v, want %v", test.name, got, test.wantStr)
			}
		})
	}
}

func TestIterating(t *testing.T) {
	doc := NewDoc("root")
	root := doc.RootElem()
	root.AddElem("first").AddText("une")
	root.AddText("some text")
	root.AddElem("second", Attrib{"class": "someclass"})

	t.Run("iterating", func(t *testing.T) {
		for _, child := range root.Children() {
			if child.Type() == NodeElem {
				child.(Elem).Attrib()["visible"] = "true"
			}
		}
		wantStr := `<root><first visible="true">une</first>some text<second class="someclass" visible="true"/></root>`
		buf := bytes.Buffer{}
		written, err := doc.WriteTo(&buf)
		if err != nil {
			t.Errorf("iterating: unexpected error %v", err)
		}
		if got := buf.String(); got != wantStr {
			t.Errorf("iterating: got = %v, want %v", got, wantStr)
		}
		if written != int64(len(wantStr)) {
			t.Errorf("iterating: unexpected length: %v <> %v", written, len(wantStr))
		}
	})
}
