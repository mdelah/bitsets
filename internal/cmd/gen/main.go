package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	t := template.Must(template.New("").ParseGlob("internal/*.go.tmpl"))
	for _, v := range []Variant{8, 16, 32, 64, 1024, 8192, -1} {
		must(os.MkdirAll(v.Pkg(), 0o755))
		for _, fn := range []string{"set.go", "set_test.go"} {
			fp, err := os.Create(filepath.Join(v.Pkg(), fn))
			must(err)
			must(t.ExecuteTemplate(fp, fn+".tmpl", v))
			must(fp.Close())
		}
	}
}

type (
	Variant  int
	Operator struct {
		Name  string
		Expr  string
		Expr2 string
		Desc  string
	}
)

func (d Variant) Cap() int       { return int(d) }
func (d Variant) Large() bool    { return d > 64 || d == -1 }
func (d Variant) Dynamic() bool  { return d == -1 }
func (d Variant) PageCount() int { return int(d / 64) }
func (d Variant) LastPage() int  { return d.PageCount() - 1 }

func (d Variant) Max() string {
	if d.Dynamic() {
		return "math.Max"
	}
	return fmt.Sprint(d - 1)
}

func (d Variant) None() string {
	if d.Large() {
		return d.Pkg() + ".None()"
	}
	return d.Pkg() + ".None"
}

func (d Variant) All() string {
	if d.Large() {
		return d.Pkg() + ".All()"
	}
	return d.Pkg() + ".All"
}

func (d Variant) Pkg() string {
	switch {
	case d.Dynamic():
		return "vbit"
	case d == 1024:
		return "kbit"
	case d.Large():
		return fmt.Sprintf("kbit%d", d/1024)
	default:
		return fmt.Sprintf("bit%d", d)
	}
}

func (d Variant) Operators() []Operator {
	return []Operator{
		{"Sub", "s&^other", "*s&=^other", "values present on the left but not right"},
		{"And", "s&other", "*s&=other", "values common to both sides"},
		{"Or", "s|other", "*s|=other", "values on either side (or both)"},
		{"Xor", "s^other", "*s^=other", "values on exactly one side"},
		{"Nor", "^s&^other", "*s=^*s&^other", "values absent from both sides"},
		{"Iff", "s^^other", "*s^=^other", "values on both sides, or neither"},
		{"Imply", "^s|other", "*s=^*s|other", "values on the right side, or not the left"},
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
