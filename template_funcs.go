package webls

import (
	"html/template"
	"time"

	"unicode/utf8"

	"github.com/ammario/micropkg/strutil"
	humanize "github.com/dustin/go-humanize"
	"github.com/willf/pad"
)

//TemplateFuncs sets up basic functions for t
var TemplateFuncs = template.FuncMap{
	"pad": pad.Right,
	"clean_time": func(t time.Time) string {
		return t.Format("01/02/2006 15:04:05")
	},
	"pad_rest": func(str string, ln int, padder string) string {
		return pad.Right("", ln-utf8.RuneCountInString(str), padder)
	},
	"epad": func(str string, ln int, padder string) string {
		return pad.Right(strutil.Ellipsis(str, ln), ln, padder)
	},
	"ellipsis": strutil.Ellipsis,
	"bytes": func(size int64) string {
		return pad.Right(humanize.Bytes(uint64(size)), 10, " ")
	},
}
