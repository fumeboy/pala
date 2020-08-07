package pala_test

import (
	"fmt"
	. "pala"
	"testing"
)

func Test1(t *testing.T)  {
	var function_name = func(c *C, end ...*Matcher) (string, string) {
		c.DELIMIT(ANY)
		t, s := c.TILL(append(end, BLANK)...)
		switch s {
		case BLANKi:
			s = c.DELIMIT(end...)
		}
		return t,s
	}

	var param = func(c *C, end ...*Matcher) (string, string) {
		switch c.DELIMIT(ER("'"), ANY) {
		case "'":
			t,s:= c.TILL(ER("'"))
			s = c.DELIMIT(end...)
			return t,s
		case ANYi:
			t, s := c.TILL(append(end, BLANK)...)
			switch s {
			case BLANKi:
				s = c.DELIMIT(end...)
			}
			return t,s
		default:
			panic("")
		}
	}

	var sentence = func(c *C) {
		c.DELIMIT(ER("let"))
		c.DELIMIT(BLANK)
		fn, _ := function_name(c,ER("("))
		fmt.Println(fn)
		repeat:
		for {
			p, sep := param(c, ER(","), ER(")"))
			fmt.Println(p)
			switch sep {
			case ",":
				continue
			case ")":
				break repeat
			default:
			}
		}
	}

	ctx := NEWcontext([]byte("let apple(a,'abc')"))
	sentence(ctx)
}
