package pointer

import (
	"log"
	"os"
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
			p.lastCharWasSpace = false
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
		n := p.ss.head.Pop()
		c = p.StepAndGet(*f)
		steps := 1
		for jumpingMode := false; jumpingMode || c == ' ' || c == ';'; c = p.StepAndGet(*f) {
			steps += 1
			if c == ';' {
				jumpingMode = !jumpingMode
			}
		}
		if n > 0 {
			// we need to reverse that step
			p.Reverse()
			for i := 0; i < steps; i++ {
				p.Step(*f)
			}
			p.Reverse()
			if c != ' ' && c != ';' {
				if n > 0 {
					for i := 0; i < n; i++ {
						p.eval(c, f)
					}
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
		if n > 22+p.ss.height {
			n = n - 20
			if p.ss.head.height <= n {
				p.ss.head.Push(0)
			} else {
				p.ss.head.Push(p.ss.head.data[p.ss.head.height-n])
			}
			return
		}
		now := time.Now()
		x, y, lx, ly := f.Dump()
		const uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
		heights := make([]int, p.ss.height)
		s := p.ss.head
		for i := p.ss.height - 1; i >= 0; i-- {
			heights[i] = s.height
			s = s.next
		}
		if n <= 0 {
			for _, e := range os.Environ() {
				p.ss.head.Push(0)
				for i := len(e) - 1; i >= 0; i-- {
					p.ss.head.Push(int(e[i]))
				}
			}
		}
		if n <= 0 {
			p.ss.head.Push(0)
			p.ss.head.Push(0)
			for i := len(p.Argv) - 1; i >= 0; i-- {
				p.ss.head.Push(0)
				for j := len(p.Argv[i]) - 1; j >= 0; j-- {
					p.ss.head.Push(int(p.Argv[i][j]))
				}
			}
		}
		if (n > 22 && n <= 22+p.ss.height) || n <= 0 {
			for i := 0; i < len(heights); i++ {
				p.ss.head.Push(heights[i])
			}
		}
		if n == 22 || n <= 0 {
			p.ss.head.Push(p.ss.height)
		}
		if n == 21 || n <= 0 {
			p.ss.head.Push((now.Hour() * 256 * 256) + (now.Minute() * 256) + now.Second())
		}
		if n == 20 || n <= 0 {
			p.ss.head.Push(((now.Year() - 1900) * 256 * 256) + (int(now.Month()) * 256) + now.Day())
		}
		if n == 19 || n <= 0 {
			p.ss.head.Push(lx + x)
		}
		if n == 18 || n <= 0 {
			p.ss.head.Push(ly + y)
		}
		if n == 17 || n <= 0 {
			p.ss.head.Push(x)
		}
		if n == 16 || n <= 0 {
			p.ss.head.Push(y)
		}
		if n == 15 || n <= 0 {
			p.ss.head.Push(p.sox)
		}
		if n == 14 || n <= 0 {
			p.ss.head.Push(p.soy)
		}
		if n == 13 || n <= 0 {
			p.ss.head.Push(p.dx)
		}
		if n == 12 || n <= 0 {
			p.ss.head.Push(p.dy)
		}
		if n == 11 || n <= 0 {
			p.ss.head.Push(p.x)
		}
		if n == 10 || n <= 0 {
			p.ss.head.Push(p.y)
		}
		if n == 9 || n <= 0 {
			p.ss.head.Push(0)
		}
		if n == 8 || n <= 0 {
			p.ss.head.Push(*((*int)(unsafe.Pointer(p))))
		}
		if n == 7 || n <= 0 {
			p.ss.head.Push(2)
		}
		if n == 6 || n <= 0 {
			p.ss.head.Push('/')
		}
		if n == 5 || n <= 0 {
			p.ss.head.Push(0) // TODO update when implementing =
		}
		if n == 4 || n <= 0 {
			p.ss.head.Push(1)
		}
		if n == 3 || n <= 0 {
			p.ss.head.Push(1048576)
		}
		if n == 2 || n <= 0 {
			p.ss.head.Push(uintSize / 8)
		}
		if n == 1 || n <= 0 {
			p.ss.head.Push(0b00000) // TODO update when implementing t, i, o and =
		}
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
