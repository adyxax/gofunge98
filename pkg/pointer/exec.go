package pointer

import (
	"log"

	"git.adyxax.org/adyxax/gofunge/pkg/field"
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
	case '#':
		p.Step(*f)
	case 'j':
		n := p.ss.head.Pop()
		for j := 0; j < n; j++ {
			p.Step(*f)
		}
	case 'q':
		v := p.ss.head.Pop()
		return true, &v
	case 'k':
		n := p.ss.head.Pop()
		c = p.StepAndGet(*f)
		for jumpingMode := false; jumpingMode || c == ' ' || c == ';'; c = p.StepAndGet(*f) {
			if c == ';' {
				jumpingMode = !jumpingMode
			}
		}
		for j := 0; j < n; j++ {
			p.eval(c, f)
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
		p.ss.End(p)
	case 'u':
		p.ss.Under()
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
			log.Fatalf("Non implemented instruction code %d : %c", c, c)
		}
	}
	return
}
