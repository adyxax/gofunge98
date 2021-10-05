package pointer

import (
	"log"
	"os"
	"strings"
	"time"
	"unsafe"

	"git.adyxax.org/adyxax/gofunge98/pkg/field"
)

func (p *Pointer) Exec(f *field.Field) (done bool, returnValue *int) {
	c := p.Get(*f)
	if p.stringMode {
		if p.lastCharWasSpace {
			for c == ' ' {
				c = p.StepAndGet(*f)
			}
			p.lastCharWasSpace = false
		}
		if c == '"' {
			p.stringMode = false
		} else {
			if c == ' ' {
				p.lastCharWasSpace = true
			}
			p.ss.head.Push(c)
		}
	} else {
		for jumpingMode := false; jumpingMode || c == ' ' || c == ';'; c = p.StepAndGet(*f) {
			if c == ';' {
				jumpingMode = !jumpingMode
			}
		}
		done, returnValue = p.eval(c, f)
	}
	p.Step(*f)
	return
}

func (p *Pointer) eval(c int, f *field.Field) (done bool, returnValue *int) {
	handled := true
	switch c {
	case '@':
		return true, nil
	case 'z':
	case '#':
		p.Step(*f)
	case 'j':
		n := p.ss.head.Pop()
		if n > 0 {
			for j := 0; j < n; j++ {
				p.Step(*f)
			}
		} else {
			p.Reverse()
			for j := 0; j < -n; j++ {
				p.Step(*f)
			}
			p.Reverse()
		}
	case 'q':
		v := p.ss.head.Pop()
		return true, &v
	case 'k':
		x, y := p.x, p.y
		n := p.ss.head.Pop()
		c = p.StepAndGet(*f)
		for jumpingMode := false; jumpingMode || c == ' ' || c == ';'; c = p.StepAndGet(*f) {
			if c == ';' {
				jumpingMode = !jumpingMode
			}
		}
		if n > 0 {
			p.x, p.y = x, y
			if c != ' ' && c != ';' {
				for i := 0; i < n; i++ {
					p.eval(c, f)
				}
			}
		}
	case '!':
		v := p.ss.head.Pop()
		if v == 0 {
			v = 1
		} else {
			v = 0
		}
		p.ss.head.Push(v)
	case '`':
		b, a := p.ss.head.Pop(), p.ss.head.Pop()
		if a > b {
			a = 1
		} else {
			a = 0
		}
		p.ss.head.Push(a)
	case '_':
		v := p.ss.head.Pop()
		if v == 0 {
			p.Redirect('>')
		} else {
			p.Redirect('<')
		}
	case '|':
		v := p.ss.head.Pop()
		if v == 0 {
			p.Redirect('v')
		} else {
			p.Redirect('^')
		}
	case 'w':
		b, a := p.ss.head.Pop(), p.ss.head.Pop()
		if a < b {
			p.Redirect('[')
		} else if a > b {
			p.Redirect(']')
		}
	case '+':
		b, a := p.ss.head.Pop(), p.ss.head.Pop()
		p.ss.head.Push(a + b)
	case '*':
		b, a := p.ss.head.Pop(), p.ss.head.Pop()
		p.ss.head.Push(a * b)
	case '-':
		b, a := p.ss.head.Pop(), p.ss.head.Pop()
		p.ss.head.Push(a - b)
	case '/':
		b, a := p.ss.head.Pop(), p.ss.head.Pop()
		if b == 0 {
			p.ss.head.Push(0)
			return
		}
		p.ss.head.Push(a / b)
	case '%':
		b, a := p.ss.head.Pop(), p.ss.head.Pop()
		if b == 0 {
			p.ss.head.Push(0)
			return
		}
		p.ss.head.Push(a % b)
	case '"':
		p.stringMode = true
	case '\'':
		p.ss.head.Push(p.StepAndGet(*f))
	case 's':
		p.Step(*f)
		f.Set(p.x, p.y, p.ss.head.Pop())
	case '$':
		p.ss.head.Pop()
	case ':':
		p.ss.head.Duplicate()
	case '\\':
		p.ss.head.Swap()
	case 'n':
		p.ss.head.Clear()
	case '{':
		p.ss.Begin(p)
	case '}':
		if p.ss.End(p) {
			p.Reverse()
		}
	case 'u':
		if p.ss.Under() {
			p.Reverse()
		}
	case 'g':
		y, x := p.ss.head.Pop(), p.ss.head.Pop()
		p.ss.head.Push(f.Get(x+p.sox, y+p.soy))
	case 'p':
		y, x, v := p.ss.head.Pop(), p.ss.head.Pop(), p.ss.head.Pop()
		f.Set(x+p.sox, y+p.soy, v)
	case '.':
		p.DecimalOutput(p.ss.head.Pop())
	case ',':
		p.CharacterOutput(p.ss.head.Pop())
	case '&':
		p.ss.head.Push(p.DecimalInput())
	case '~':
		p.ss.head.Push(p.CharacterInput())
	case 'y':
		n := p.ss.head.Pop()
		now := time.Now()
		x, y, lx, ly := f.Dump()
		const uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
		heights := make([]int, p.ss.height)
		s := p.ss.head
		for i := p.ss.height - 1; i >= 0; i-- {
			heights[i] = s.height
			s = s.next
		}
		// 20
		for _, e := range os.Environ() {
			vars := strings.SplitN(e, "=", 2)
			if vars[0] != "LC_ALL" && vars[0] != "PWD" && vars[0] != "PATH" && vars[0] != "DISPLAY" && vars[0] != "USER" && vars[0] != "TERM" && vars[0] != "LANG" && vars[0] != "HOME" && vars[0] != "EDITOR" && vars[0] != "SHELL" {
				continue
			}
			p.ss.head.Push(0)
			for i := len(e) - 1; i >= 0; i-- {
				p.ss.head.Push(int(e[i]))
			}
		}
		// 19
		p.ss.head.Push(0)
		p.ss.head.Push(0)
		for i := len(p.Argv) - 1; i >= 0; i-- {
			p.ss.head.Push(0)
			for j := len(p.Argv[i]) - 1; j >= 0; j-- {
				p.ss.head.Push(int(p.Argv[i][j]))
			}
		}
		// 18
		for i := 0; i < len(heights); i++ {
			p.ss.head.Push(heights[i])
		}
		// 17
		p.ss.head.Push(p.ss.height)
		// 16
		p.ss.head.Push((now.Hour() * 256 * 256) + (now.Minute() * 256) + now.Second())
		// 15
		p.ss.head.Push(((now.Year() - 1900) * 256 * 256) + (int(now.Month()) * 256) + now.Day())
		// 14
		p.ss.head.Push(lx - 1)
		p.ss.head.Push(ly - 1)
		// 13
		p.ss.head.Push(x)
		p.ss.head.Push(y)
		// 12
		p.ss.head.Push(p.sox)
		p.ss.head.Push(p.soy)
		// 11
		p.ss.head.Push(p.dx)
		p.ss.head.Push(p.dy)
		// 10
		p.ss.head.Push(p.x)
		p.ss.head.Push(p.y)
		// 9
		p.ss.head.Push(0)
		// 8
		p.ss.head.Push(*((*int)(unsafe.Pointer(p))))
		// 7
		p.ss.head.Push(2)
		// 6
		p.ss.head.Push('/')
		// 5
		p.ss.head.Push(0) // TODO update when implementing =
		// 4
		p.ss.head.Push(1)
		// 3
		p.ss.head.Push(1048576)
		// 2
		p.ss.head.Push(uintSize / 8)
		// 1
		p.ss.head.Push(0b00000) // TODO update when implementing t, i, o and =
		if n > 0 {
			if n > p.ss.head.height {
				p.ss.head.height = 1
				p.ss.head.data[0] = 0
			} else {
				v := p.ss.head.data[p.ss.head.height-n]
				p.ss.head.height = heights[0]
				p.ss.head.Push(v)
			}
		}
	case '(':
		n := p.ss.head.Pop()
		v := 0
		for i := 0; i < n; i++ {
			v = v*256 + p.ss.head.Pop()
		}
		// No fingerprints supported
		p.Reverse()
	case ')':
		n := p.ss.head.Pop()
		v := 0
		for i := 0; i < n; i++ {
			v = v*256 + p.ss.head.Pop()
		}
		// No fingerprints supported
		p.Reverse()
	case 'i':
		log.Fatalf("Non implemented instruction code %d : %c", c, c)
	case 'o':
		log.Fatalf("Non implemented instruction code %d : %c", c, c)
	case '=':
		log.Fatalf("Non implemented instruction code %d : %c", c, c)
	case 't':
		log.Fatalf("Non implemented instruction code %d : %c", c, c)
	default:
		handled = false
	}
	if !handled {
		switch {
		case p.Redirect(c):
		case c >= '0' && c <= '9':
			p.ss.head.Push(c - '0')
		case c >= 'a' && c <= 'f':
			p.ss.head.Push(c - 'a' + 10)
		default:
			p.Redirect('r')
		}
	}
	return
}
