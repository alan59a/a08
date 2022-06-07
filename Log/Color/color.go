package Colour

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type Colorize struct {
	Colors  map[string]string
	Disable bool
	Reset   bool
}

var (
	DefaultColors map[string]string
	def           Colorize
	parseReRaw    = `\[[a-z0-9_-]+\]`
	parseRe       = regexp.MustCompile(`(?i)` + parseReRaw)
	prefixRe      = regexp.MustCompile(`^(?i)(` + parseReRaw + `)+`)
)

func Color(v string) string {
	return def.Color(v)
}

func ColorPrefix(v string) string {
	return def.ColorPrefix(v)
}

func (c *Colorize) Color(v string) string {
	matches := parseRe.FindAllStringIndex(v, -1)
	if len(matches) == 0 {
		return v
	}
	result := new(bytes.Buffer)
	colored := false
	m := []int{0, 0}
	for _, nm := range matches {
		result.WriteString(v[m[1]:nm[0]])
		m = nm
		var replace string
		if code, ok := c.Colors[v[m[0]+1:m[1]-1]]; ok {
			colored = true
			if !c.Disable {
				replace = fmt.Sprintf("\033[%sm", code)
			}
		} else {
			replace = v[m[0]:m[1]]
		}
		result.WriteString(replace)
	}
	result.WriteString(v[m[1]:])
	if colored && c.Reset && !c.Disable {
		result.WriteString("\033[0m")
	}
	return result.String()
}

func (c *Colorize) ColorPrefix(v string) string {
	return prefixRe.FindString(strings.TrimSpace(v))
}

func init() {
	DefaultColors = map[string]string{
		"default":   "39",
		"_default_": "49",

		"black":         "30",
		"red":           "31",
		"green":         "32",
		"yellow":        "33",
		"blue":          "34",
		"magenta":       "35",
		"cyan":          "36",
		"light_gray":    "37",
		"dark_gray":     "90",
		"light_red":     "91",
		"light_green":   "92",
		"light_yellow":  "93",
		"light_blue":    "94",
		"light_magenta": "95",
		"light_cyan":    "96",
		"white":         "97",

		"_black_":         "40",
		"_red_":           "41",
		"_green_":         "42",
		"_yellow_":        "43",
		"_blue_":          "44",
		"_magenta_":       "45",
		"_cyan_":          "46",
		"_light_gray_":    "47",
		"_dark_gray_":     "100",
		"_light_red_":     "101",
		"_light_green_":   "102",
		"_light_yellow_":  "103",
		"_light_blue_":    "104",
		"_light_magenta_": "105",
		"_light_cyan_":    "106",
		"_white_":         "107",

		"bold":       "1",
		"dim":        "2",
		"underline":  "4",
		"blink_slow": "5",
		"blink_fast": "6",
		"invert":     "7",
		"hidden":     "8",

		"reset":      "0",
		"reset_bold": "21",
	}

	def = Colorize{
		Colors: DefaultColors,
		Reset:  true,
	}
}

func Print(a string) (n int, err error) {
	return fmt.Print(Color(a))
}

func Println(a string) (n int, err error) {
	return fmt.Println(Color(a))
}

func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(Color(format), a...)
}

func Fprint(w io.Writer, a string) (n int, err error) {
	return fmt.Fprint(w, Color(a))
}

func Fprintln(w io.Writer, a string) (n int, err error) {
	return fmt.Fprintln(w, Color(a))
}

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, Color(format), a...)
}
