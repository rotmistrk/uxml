package uxml

import "testing"

func TestXmlEncode(t *testing.T) {
	type args struct {
		enc string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no escaping",
			args: args{"Hello, world!"},
			want: "Hello, world!",
		},
		{
			name: "amp escaping",
			args: args{"chip & dail"},
			want: "chip &amp; dail",
		},
		{
			name: "less escaping",
			args: args{"15 < 25"},
			want: "15 &lt; 25",
		},
		{
			name: "greater escaping",
			args: args{"35 > 25"},
			want: "35 &gt; 25",
		},
		{
			name: "quote escaping",
			args: args{"the \"special\" case of \"\"string\"\""},
			want: "the &quot;special&quot; case of &quot;&quot;string&quot;&quot;",
		},
		{
			name: "apos escaping",
			args: args{"let's go"},
			want: "let&apos;s go",
		},
		{
			name: "control chars in the middle",
			args: args{"here\twe\nare"},
			want: "here&#x09;we&#x0a;are",
		},
		{
			name: "control chars in the ends",
			args: args{"\nthat is funny & more funny\n"},
			want: "&#x0a;that is funny &amp; more funny&#x0a;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XmlEncode(tt.args.enc); got != tt.want {
				t.Errorf("XmlEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}
