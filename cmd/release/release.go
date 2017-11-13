package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Target struct {
	Name    string   `json:"name"`
	Aliases []string `json:"aliases,omitempty"`
}

type Muscle struct {
	Name    string   `json:"name"`
	Aliases []string `json:"aliases,omitempty"`
	Targets []Target `json:"targets,omitempty"`
}

type Group struct {
	Name    string   `json:"name"`
	Aliases []string `json:"aliases,omitempty"`
	Muscles []Muscle `json:"muscles,omitempty"`
}

type Release struct {
	Groups []Group `json:"groups,omitempty"`
}

var (
	Aka        = []byte("aka: ")
	H1         = []byte("# ")
	H2         = []byte("## ")
	Asterisk   = []byte("* ")
	OpenParen  = []byte(" (")
	CloseParen = []byte(")")
	Comma      = ", "
)

func NewRelease(dir string) (Release, error) {
	var r Release
	d, err := ioutil.ReadDir(dir)
	if err != nil {
		return r, err
	}

	for _, f := range d {
		if !strings.HasSuffix(f.Name(), ".md") {
			continue
		}

		b, err := ioutil.ReadFile(filepath.Join(dir, f.Name()))
		if err != nil {
			return r, err
		}

		var g Group
		var m Muscle
		buf := bufio.NewReader(bytes.NewBuffer(b))

		for {
			l, _, err := buf.ReadLine()
			if err != nil {
				if err == io.EOF {
					g.Muscles = append(g.Muscles, m)
					break
				}

				return r, err
			}

			if len(l) == 0 {
				continue
			}

			if bytes.HasPrefix(l, H1) {
				g.Name = string(l[len(H1):])
			}

			if bytes.HasPrefix(l, Aka) {
				aliases := string(l[len(Aka):])
				g.Aliases = strings.Split(aliases, Comma)
			}

			if bytes.HasPrefix(l, H2) {
				if m.Name != "" {
					g.Muscles = append(g.Muscles, m)
				}

				m = Muscle{}
				m.Name = string(l[len(H2):])
			}

			if bytes.HasPrefix(l, Asterisk) {
				var t Target
				target := l[len(Asterisk):]
				if !bytes.Contains(target, OpenParen) {
					t.Name = string(target)
					m.Targets = append(m.Targets, t)
				} else {
					op := bytes.Index(target, OpenParen)
					cp := bytes.Index(target, CloseParen)
					n := op + len(OpenParen) + len(Aka)
					aliases := string(target[n:cp])

					t.Name = string(target[:op])
					t.Aliases = strings.Split(aliases, Comma)
					m.Targets = append(m.Targets, t)
				}
			}
		}

		r.Groups = append(r.Groups, g)
	}

	return r, nil
}
