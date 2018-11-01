package html_test

import (
	"github.com/hill-daniel/alexa-rss-flashbriefing/html"
	"testing"
)

func TestRemoveTags(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{name: "should remove simple tag",
			args: "<html>hello</html>",
			want: "hello",
		},
		{name: "should remove nested tags",
			args: "<html><body><h1>hello</h1></body></html>",
			want: "hello",
		},
		{name: "should remove xml",
			args: "<dc:creator><![CDATA[Heinrich]]>Foo</dc:creator>",
			want: "Foo",
		},
		{name: "should leave plain text as is",
			args: "Foo, Bar, Bazz! <> :P",
			want: "Foo, Bar, Bazz! <> :P",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := html.RemoveTags(tt.args); got != tt.want {
				t.Errorf("RemoveTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
