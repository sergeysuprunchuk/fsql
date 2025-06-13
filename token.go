package fsql

const (
	idbeg uint8 = iota

	add     // +
	sub     // -
	mul     // *
	div     // /
	rem     // %
	eq      // =
	neq     // !=
	less    // <
	more    // >
	lesseq  // <=
	moreeq  // >=
	lparen  // (
	rparen  // )
	and     // and
	or      // or
	not     // not
	boolean // true,false
	integer // 8
	float   // 8.8
	text    // "text"

	idend
)

type baseTok struct {
	id    uint8
	start pos
}

func newBase(id uint8, start pos) baseTok {
	return baseTok{
		id:    id,
		start: start,
	}
}

func (t baseTok) is(ids ...uint8) bool {
	for _, id := range ids {
		if id == t.id {
			return true
		}
	}
	return false
}

func (t baseTok) getVal() string { return "" }

func (t baseTok) getPos() pos { return t.start }

func (t baseTok) getId() uint8 { return t.id }

type valTok struct {
	baseTok
	val string
}

func newVal(id uint8, start pos, val string) valTok {
	return valTok{
		baseTok: newBase(id, start),
		val:     val,
	}
}

func (t valTok) getVal() string { return t.val }

type token interface {
	is(ids ...uint8) bool
	getVal() string
	getPos() pos
	getId() uint8
}
