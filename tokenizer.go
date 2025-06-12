package fsql

import (
	"fmt"
	"unicode"
)

type pos struct {
	line   int
	offset int
}

func (p *pos) string() string {
	return fmt.Sprintf("строка %d символ %d", p.line+1, p.offset+1)
}

type tokenizer struct {
	data   []rune //запрос
	cursor int    //индекс текущего символа
	pos           //позиция текущего символа (строка, символ)
}

func newTokenizer(query string) *tokenizer {
	return &tokenizer{
		data: []rune(query),
	}
}

// возвращает текущий символ или 0, если это конец строки
func (t *tokenizer) char() rune {
	if t.cursor < len(t.data) {
		return t.data[t.cursor]
	}
	return 0
}

// перемещает курсор на следующий символ
func (t *tokenizer) next() {
	if t.char() == 0 {
		return
	}

	if t.char() == '\n' {
		t.line++
		t.offset = 0
		t.cursor++
		return
	}

	t.cursor++
	t.offset++
}

// пропускает пробельные символы
func (t *tokenizer) skip() {
	for unicode.IsSpace(t.char()) {
		t.next()
	}
}
