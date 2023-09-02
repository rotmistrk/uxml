package uxml

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

func (doc *doc) WriteTo(out io.Writer) (int64, error) {
	return doc.RootElem().WriteTo(out)
}

func (text *text) WriteTo(out io.Writer) (int64, error) {
	enc := text.contents
	enc = XmlEncode(enc)
	n, err := out.Write([]byte(enc))
	return int64(n), err
}

func XmlEncode(enc string) string {
	enc = strings.ReplaceAll(enc, "&", "&amp;")
	enc = strings.ReplaceAll(enc, "<", "&lt;")
	enc = strings.ReplaceAll(enc, ">", "&gt;")
	enc = strings.ReplaceAll(enc, "'", "&apos;")
	enc = strings.ReplaceAll(enc, "\"", "&quot;")
	out := ""
	off := 0
	for pos, ch := range enc {
		if ch < ' ' {
			if off < pos {
				out += enc[off:pos]
			}
			out += fmt.Sprintf("&#x%02x;", ch)
			off = pos + 1
		}
	}
	if len(enc) > off {
		out += enc[off:]
	}
	return out
}

func (elem *elem) WriteTo(out io.Writer) (count int64, err error) {
	n, e := fmt.Fprintf(out, "<%s", elem.tag)
	if count += int64(n); e != nil {
		return count, e
	}
	keys := make([]string, 0)
	for k := range elem.attrib {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := fmt.Sprintf("%v", elem.attrib[k])
		n, e = fmt.Fprintf(out, " %s=\"%s\"", k, XmlEncode(v))
		if count += int64(n); e != nil {
			return count, e
		}
	}
	if elem.children == nil || len(elem.children) == 0 {
		n, e = fmt.Fprint(out, "/>")
	} else {
		n, e = fmt.Fprint(out, ">")
		if count += int64(n); e != nil {
			return count, e
		}
		for _, node := range elem.children {
			n64, e := node.WriteTo(out)
			if count += n64; e != nil {
				return count, err
			}
		}
		n, e = fmt.Fprintf(out, "</%s>", elem.tag)
	}
	count += int64(n)
	err = e
	return
}
