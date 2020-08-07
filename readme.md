# pala

原创的第四种分词器

阅读下列示例代码，请对比参考之前的 [pili](https://github.com/fumeboy/pili)

./pala_test.go

```go
func Test1(t *testing.T)  {
    // 两种匹配文本方式
    // c.DELIMIT  略过空格，匹配文本
    // c.TILL 读取文本，直到遇到 Matcher 匹配的文本
	var function_name = func(c *C, end ...*Matcher) (string, string) {
        // end : 接下来会遇到的符号对应的 Matcher
		c.DELIMIT(ANY)
		token, delimiter := c.TILL(append(end, BLANK)...)
		switch delimiter {
		case BLANKi:
			delimiter = c.DELIMIT(end...)
		}
		return token,delimiter
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

```