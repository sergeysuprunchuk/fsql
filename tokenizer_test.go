package fsql

import (
	"testing"
)

func Test_next(t *testing.T) {
	const data = `func (t *tokenizer) skip() {
	for unicode.IsSpace(t.char()) {
		t.next()
	}
}`

	tests := []struct {
		n    int  //сколько раз вызывать метод next
		pos       //значение t.pos после n вызовов next
		char rune //значение t.char() после n вызовов next
	}{
		{0, pos{0, 0}, 'f'},
		{2, pos{0, 2}, 'n'},
		{4, pos{0, 4}, ' '},
		{8, pos{0, 8}, '*'},
		{16, pos{0, 16}, 'e'},
		{32, pos{1, 3}, 'r'},
		{64, pos{2, 2}, 't'},
	}

	for ti, test := range tests {
		tok := newTokenizer(data)
		for range test.n {
			tok.next()
		}

		if tok.char() != test.char {
			t.Errorf("%d: char: ожидал %s, получил %s",
				ti+1, string(test.char), string(tok.char()))
		}

		if tok.pos != test.pos {
			t.Errorf("%d: pos: ожидал %q, получил %q",
				ti+1, test.pos.string(), tok.pos.string())
		}

		if tok.cursor != test.n {
			t.Errorf("%d: cursor: ожидал %d, получил %d",
				ti+1, test.n, tok.cursor)
		}
	}
}

func Test_skip(t *testing.T) {
	tests := []struct {
		data   string
		pos         //финальный t.pos
		char   rune //финальный t.char()
		cursor int  //финальный cursor
	}{
		{`func`, pos{0, 0}, 'f', 0},
		{`	 func`, pos{0, 2}, 'f', 2},
		{`			
func`, pos{1, 0}, 'f', 4},
		{`			
 	func`, pos{1, 2}, 'f', 6},
	}

	for ti, test := range tests {
		tok := newTokenizer(test.data)
		tok.skip()

		if tok.char() != test.char {
			t.Errorf("%d: char: ожидал %s, получил %s",
				ti+1, string(test.char), string(tok.char()))
		}

		if tok.pos != test.pos {
			t.Errorf("%d: pos: ожидал %q, получил %q",
				ti+1, test.pos.string(), tok.pos.string())
		}

		if tok.cursor != test.cursor {
			t.Errorf("%d: cursor: ожидал %d, получил %d",
				ti+1, test.cursor, tok.cursor)
		}
	}
}
