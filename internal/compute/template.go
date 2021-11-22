package compute

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// Template is just a list of Atom, where an Atom is either a Variable or a Constant string.
type Template []Atom

type Atom interface {
	atom()
	String() string
}

type Attribute string

const (
	LengthAttr Attribute = "length"
	RangeAttr  Attribute = "range"
)

// Variable represents a variable in the template that may be substituted for. A
// variable is optionally qualified by an attribute, which is data associated
// with a variable (e.g., length, range). Attributes are currently unused, and
// exist for future expansion.
type Variable struct {
	Name      string
	Attribute Attribute
}

type Constant string

func (Variable) atom() {}
func (Constant) atom() {}

func (v Variable) String() string {
	if v.Attribute != "" {
		return v.Name + "." + string(v.Attribute)
	}
	return v.Name
}
func (c Constant) String() string { return string(c) }

const varAllowed = "abcdefghijklmnopqrstuvwxyzABCEDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"

// scanTemplate scans an input string to produce a Template. Recognized
// metavariable syntax is `$(varAllowed+)`.
func scanTemplate(buf []byte) (*Template, error) {
	// Tracks whether the current token is a variable.
	var isVariable bool

	var start int
	var r rune
	var token []rune
	var result []Atom

	next := func() rune {
		r, start := utf8.DecodeRune(buf)
		buf = buf[start:]
		return r
	}

	appendAtom := func(atom Atom) {
		if a, ok := atom.(Constant); ok && len(a) == 0 {
			return
		}
		if a, ok := atom.(Variable); ok && len(a.Name) == 0 {
			return
		}
		result = append(result, atom)
		// Reset token, but reuse the backing memory.
		token = token[:0]
	}

	for len(buf) > 0 {
		r = next()
		switch r {
		case '$':
			if len(buf[start:]) > 0 {
				if r, _ = utf8.DecodeRune(buf); strings.ContainsRune(varAllowed, r) {
					// Start of a recognized variable.
					if isVariable {
						// We were busy scanning a variable.
						appendAtom(Variable{Name: string(token)}) // Push variable.
					} else {
						// We were busy scanning a constant.
						appendAtom(Constant(token))
					}
					token = append(token, '$')
					isVariable = true
					continue
				}
				// Something else, push the '$' we saw and continue.
				token = append(token, '$')
				isVariable = false
				continue
			}
			// Trailing '$'
			if isVariable {
				appendAtom(Variable{Name: string(token)}) // Push variable.
				isVariable = false
			} else {
				appendAtom(Constant(token))
			}
			token = append(token, '$')
		case '\\':
			if isVariable {
				// We were busy scanning a variable. A '\' always terminates it.
				appendAtom(Variable{Name: string(token)}) // Push variable.
				isVariable = false
			}
			if len(buf[start:]) > 0 {
				r = next()
				switch r {
				case 'n':
					token = append(token, '\n')
				case 'r':
					token = append(token, '\r')
				case 't':
					token = append(token, '\t')
				case '\\', '$', ' ', '.':
					token = append(token, r)
				default:
					token = append(token, '\\', r)
				}
				continue
			}
			// Trailing '\'
			token = append(token, '\\')
		default:
			if isVariable && !strings.ContainsRune(varAllowed, r) {
				appendAtom(Variable{Name: string(token)}) // Push variable.
				isVariable = false
			}
			token = append(token, r)
		}
	}
	if len(token) > 0 {
		if isVariable {
			appendAtom(Variable{Name: string(token)})
		} else {
			appendAtom(Constant(token))
		}
	}
	t := Template(result)
	return &t, nil
}

func toJSON(atom Atom) interface{} {
	switch a := atom.(type) {
	case Constant:
		return struct {
			Value string `json:"constant"`
		}{
			Value: string(a),
		}
	case Variable:
		return struct {
			Name      string `json:"variable"`
			Attribute string `json:"attribute,omitempty"`
		}{
			Name:      a.Name,
			Attribute: string(a.Attribute),
		}
	}
	panic("unreachable")
}

func toJSONString(template *Template) string {
	var jsons []interface{}
	for _, atom := range *template {
		jsons = append(jsons, toJSON(atom))
	}
	json, _ := json.Marshal(jsons)
	return string(json)
}
