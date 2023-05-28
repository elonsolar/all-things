package expr

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"unicode"
)

// 1+2*3/4*(5+4)

type TreeNode struct {
	val      string
	nodeType string //
	left     *TreeNode
	right    *TreeNode
}

func TestConvertSuffix(t *testing.T) {

	// fmt.Println(ConvertSuffix(" 2*(1+2)"))
	fmt.Println(ConvertSuffix(" 1 +2*3 /4* (5 + 4) "))

}

func TestMidToTree(t *testing.T) {

	tokens, _ := ConvertSuffix(" 1 +2*3 /4* (5 + 4) ")

	// ast 栈，和 op栈
	astStack := make(Stack[*TreeNode], 0, 10)
	opStack := make(Stack[string], 0, 10)

	// 遇到数字 入ast
	// 遇到操作符，如果优先级高于栈顶，入栈，否则依次弹出栈顶，取出ast 栈顶2个元素，组成新的ast,入栈
	for i, token := range tokens {

		switch token.tag {

		case INT:
			astStack = append(astStack, &TreeNode{val: string(token.lexme)})
		case OPERATOR:

			for top, err := opStack.peek(); err == nil && ((!isPriHigher(string(token.lexme), top) && top != "(") || (top != "(" && string(token.lexme) == ")")); top, err = opStack.peek() {

				opStack.pop()

				topFirstAst, err := astStack.pop()
				if err != nil {
					panic(err)
				}
				topSecondAst, err := astStack.pop()
				if err != nil {
					panic(fmt.Sprintf("%d: %s", i, err.Error()))
				}

				newAstNode := TreeNode{
					val:   top,
					left:  topSecondAst,
					right: topFirstAst,
				}

				astStack = append(astStack, &newAstNode)
			}
			if string(token.lexme) != ")" {
				opStack.push(string(token.lexme))
			}
			if top, _ := opStack.peek(); top == "(" && string(token.lexme) == ")" {
				opStack.pop()
			}

		}

		if i == len(tokens)-1 {

			// holly , mis top,err=opStack.peek()
			for top, err := opStack.peek(); err == nil; top, err = opStack.peek() {

				opStack.pop()

				topFirstAst, err := astStack.pop()
				if err != nil {
					panic(err)
				}
				topSecondAst, err := astStack.pop()
				if err != nil {
					panic(err)
				}

				newAstNode := TreeNode{
					val:   top,
					left:  topSecondAst,
					right: topFirstAst,
				}

				astStack = append(astStack, &newAstNode)
			}
			break
		}
	}

	fmt.Println(astStack)
	fmt.Println(opStack)
}

func TestSuffixToTree(t *testing.T) {}

func ConvertSuffix(expr string) ([]Token, []string) {
	scanner := &Scanner{
		text:     []rune(expr),
		position: 0,
	}

	tokens := make([]Token, 0, 10)

	opStack := make(Stack[string], 0, 10)

	sequence := make([]string, 0, 10)

	for {
		token, err := scanner.scan()
		if err != nil && err != EOF {
			panic(err)
		}

		// fmt.Println("token: ", string(token.lexme))

		if err == nil && token != nil {

			tokens = append(tokens, *token)
		}

		if err == EOF {
			for !opStack.isEmpty() {
				op, _ := opStack.pop()
				sequence = append(sequence, op)
			}
			break
		} else if token.tag == INT {
			sequence = append(sequence, string(token.lexme))
		} else if token.tag == OPERATOR {

			// 如果当前操作符，优先级大于栈顶，将其入栈，否则从栈内出栈依次取出操作符，直到小于当前操作符
			// 如果遇到( 直接入栈， 如果遇到） 弹出操作符，直到(

			currentOp := string(token.lexme)
			if opStack.isEmpty() {
				opStack.push(string(token.lexme))
			} else {
				lastOp, err := opStack.peek()
				if err != nil {
					panic(err)
				}

				switch currentOp {

				case "+", "-", "*", "/":
					if isPriHigher(currentOp, lastOp) {
						opStack.push(currentOp)
					} else {
						for !isPriHigher(currentOp, lastOp) && lastOp != "(" && !opStack.isEmpty() {

							lastOp, _ = opStack.pop()
							sequence = append(sequence, lastOp)

							lastOp, _ = opStack.peek()
						}
						opStack.push(currentOp)
					}

				case "(":
					opStack.push(currentOp)
				case ")":
					for lastOp != "(" && !opStack.isEmpty() {
						lastOp, _ = opStack.pop()
						sequence = append(sequence, lastOp)
						lastOp, _ = opStack.peek()
					}

					if lastOp == "(" {
						opStack.pop()
					}

				}

			}

		}
	}

	return tokens, sequence
}

var (
	opLevelMap = map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"(": 9,
		")": 9,
	}
)

func isPriHigher(left, right string) bool {
	return opLevelMap[left] > opLevelMap[right]
}

type Tag int

const (
	INT Tag = iota
	FLOAT
	OPERATOR //+- */

	INDENTIFIER
	PARENTHESIS //()

	UNKNOWN //()
)

type Token struct {
	tag   Tag
	lexme []rune
}

func newToken(tag Tag, lexme []rune) *Token {
	return &Token{tag: tag, lexme: lexme}
}

type Scanner struct {
	text     []rune
	position int

	endPostion int

	err error
}

var (
	EOF = errors.New("EOF")
)

func (s *Scanner) scan() (*Token, error) {

	s.position = s.endPostion

	if s.err != nil {
		return nil, s.err
	}

	s.skipWhiteSpace()

	next, err := s.readNext()
	if err != nil {
		return nil, s.addError(err)
	}

	if unicode.IsDigit(next) {
		for {
			next, err = s.readNext()
			if err == EOF {
				s.addError(err)
				return newToken(INT, s.text[s.position:s.endPostion]), nil // 是否包含endPostion，需要减1 吗？
			}

			if err != nil {
				return nil, s.addError(err)
			}

			if unicode.IsDigit(next) {
				continue
			}

			s.endPostion--

			return newToken(INT, s.text[s.position:s.endPostion]), nil // 是否包含endPostion，需要减1 吗？
		}

	} else if unicode.IsLetter(next) {
		for {
			next, err = s.readNext()
			if err == EOF {
				s.addError(err)
				return newToken(INDENTIFIER, s.text[s.position:s.endPostion]), nil // 是否包含endPostion，需要减1 吗？
			}

			if unicode.IsLetter(next) || unicode.IsDigit(next) {
				continue
			}

			return newToken(INDENTIFIER, s.text[s.position:s.endPostion]), nil
		}

	} else if string(next) == "*" || string(next) == "/" || string(next) == "+" || string(next) == "-" {
		return newToken(OPERATOR, []rune{next}), nil
	} else if string(next) == "(" || string(next) == ")" {
		return newToken(OPERATOR, []rune{next}), nil
	}

	return newToken(UNKNOWN, nil), nil
}

func (s *Scanner) skipWhiteSpace() {
	for s.position < len(s.text)-1 && s.text[s.position] == ' ' {
		s.position++
	}
	s.endPostion = s.position
}

func (s *Scanner) readNext() (rune, error) {
	// s.endPostion++
	if s.endPostion >= len(s.text) {
		return 0, EOF
	}
	defer func() {
		s.endPostion++

		// ensPostion =2
		// fmt.Println(s.endPostion)
	}()

	// endPostion =1
	// 但这里的值还是返回加一前的
	return s.text[s.endPostion], nil
}

func (s *Scanner) addError(err error) error {
	s.err = err
	return err
}

type Stack[T any] []T

func (s *Stack[T]) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) push(t T) {
	*s = append(*s, t)
}

func (s *Stack[T]) peek() (t T, err error) {
	if len(*s) == 0 {
		fmt.Println(reflect.TypeOf(t).Name())

		return reflect.Zero(reflect.TypeOf(t)).Interface().(T), errors.New("eof")
	}

	t = (*s)[len(*s)-1]

	return
}

func (s *Stack[T]) pop() (t T, err error) {
	if len(*s) == 0 {

		fmt.Println(reflect.TypeOf(t).Name())

		return reflect.Zero(reflect.TypeOf(t)).Interface().(T), errors.New("eof")
	}

	t = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return
}

func TestStach(t *testing.T) {

	var ages Stack[int] = []int{}

	ages.push((5))
	fmt.Println(ages)

	v5, err := ages.pop()
	if err != nil {
		panic(err)
	}

	fmt.Println(v5)

	v0, err := ages.pop()
	fmt.Println(v0)
	if err != nil {
		panic(err)
	}

}
