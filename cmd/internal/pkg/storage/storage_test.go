package storage

import (
	"testing"
)

type testCase struct {
	name  string
	key   string
	value string
	kind  Kind
}

func TestSetGetWithKind(t *testing.T) {
	cases := []testCase{
		{"hello world", "hello", "world", KindString},
		{"YYYXXX 123", "YYYXXX", "123", KindInt},
		{"U1", "U1", "", KindUndefined},
		{"U2 undefined", "U2", "undefined", KindUndefined},
	}

	s, err := NewStorage()
	if err != nil {
		t.Errorf("new storage: %v", err)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s.Set(c.key, c.value)
			sValue := s.Get(c.key)

			if *sValue != c.value {
				t.Errorf("values not equal")
			}

			if getType(*sValue) != getType(c.value) {
				t.Errorf("value kinds not equal")
			}

			if getType(*sValue) != c.kind {
				t.Errorf("expected value kind: %v", c.kind)
			}
		})
	}
}
