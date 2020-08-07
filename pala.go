package pala

type C struct {
	input []byte
	offset int
	len_in int
}
func NEWcontext(input []byte) *C {
	c := &C{input: input}
	c.len_in = len(c.input)
	return c
}

type Matcher struct {
	fn func(input []byte) (offset int,ok bool)
	name string
}

func ER(sep string) *Matcher {
	f := func(input []byte) (offset int,ok bool) {
		if len(input) >= offset+len(sep) && sep == string(input[offset:offset+len(sep)]) {
			return len(sep), true
		}
		return 0, false
	}
	return &Matcher{fn: f, name: sep}
}

const EOFi = " EOF"
var EOF = &Matcher{
	fn: func(input []byte) (offset int, ok bool) {
		if len(input) == 0 {
			return 0, true
		} else {
			return 0, false
		}
	},
	name: EOFi,
}
const BLANKi = " BLANK"
var BLANK = &Matcher{name:BLANKi,fn: func(input []byte) (offset int, ok bool) {
	j := 0
	l := len(input)
	for ;j<l;j++{
		if input[j] != ' '{break}
	}
	if j > 0{
		return j, true
	}else{
		return 0, false
	}
}}
const ANYi = " ANY"
var ANY = &Matcher{
	fn: func(input []byte) (offset int, ok bool) {
		if len(input)>0 &&input[0] != ' '{
			return 0, true
		}else{
			return 0, false
		}
	},
	name: ANYi,
}

func (c *C) DELIMIT(matchers ...*Matcher) string {
	var ifBLANK bool
	var ifNOTBLANK bool
	matchers_ := []*Matcher{}
	for i, l := 0, len(matchers); i < l; i++ {
		m := matchers[i]
		if m == ANY {
			ifNOTBLANK = true
			continue
		}
		if m == BLANK {
			ifBLANK = true
			continue
		}
		matchers_ = append(matchers_, m)
	}
	j := c.offset
	if ifBLANK && c.input[j] == ' '{
		o, _ := BLANK.fn(c.input[j:])
		c.offset = j + o
		return BLANK.name
	}
	for ; j < c.len_in; j++ {
		if c.input[j] != ' ' {
			break
		}
	}
	in := c.input[j:]
	for i, l := 0, len(matchers_); i < l; i++ {
		o, ok := matchers_[i].fn(in)
		if ok {
			c.offset = o + j
			return matchers_[i].name
		}
	}
	if ifNOTBLANK {
		o, ok := ANY.fn(in)
		if ok {
			c.offset = o + j
			return ANY.name
		}
	}
	panic("find fail:")
}

func (c *C) TILL (matchers ...*Matcher) (string, string) {
	j := c.offset
	for ; j <= c.len_in; j++ {
		in := c.input[j:]
		for i, l := 0, len(matchers); i < l; i++ {
			o, ok := (matchers[i]).fn(in)
			if ok {
				token := string(c.input[c.offset:j])
				c.offset = o + j
				return token, matchers[i].name
			}
		}
	}
	panic("find fail2")
}

