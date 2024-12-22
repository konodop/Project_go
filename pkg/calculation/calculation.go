package calculation

import (
	"errors"
	"strconv"
)

func Calc(expression string) (float64, error) {
	opn := 0
	cls := 0
	if len(expression) == 0 {
		return 0, errors.New("empty expression")
	}
	for i, j := range expression {
		if j == '(' {
			if i+1 < len(expression) {
				if expression[i+1] == ')' {
					return 0, errors.New("syntax Error")
				}
			}
			opn++
		} else if j == ')' {
			cls++
		}
		if opn < cls {
			return 0, errors.New("syntax Error")
		}

		if j == '+' || j == '*' || j == '/' {
			if i == 0 || i == len(expression)-1 {
				return 0, errors.New("syntax Error")
			} else if expression[i-1] == '(' || expression[i+1] == '+' || expression[i+1] == '*' || expression[i+1] == '/' || expression[i+1] == ')' {
				return 0, errors.New("syntax Error")
			}

		} else if j == '-' {
			if i == len(expression)-1 {
				return 0, errors.New("syntax Error")
			} else if expression[i+1] == '-' || expression[i+1] == '+' || expression[i+1] == '*' || expression[i+1] == '/' || expression[i+1] == ')' {
				return 0, errors.New("syntax Error")
			}
		}
	}

	if opn != cls {
		return 0, errors.New("syntax Error")
	}
	for {
		var s int = 0
		for i := range len(expression) {
			if expression[i] == '(' {
				s = i + 1
			}
		}
		var q rune = ' '
		i := s
		index := 0
		for i < len(expression) && expression[i] != ')' {
			if expression[i] == '*' {
				index = i
				q = '*'
				break
			} else if expression[i] == '/' {
				index = i
				q = '/'
				break
			} else if expression[i] == '+' && q == ' ' {
				q = '+'
				index = i
			} else if expression[i] == '-' && q == ' ' {
				q = '-'
				index = i
			}
			i++
		}
		if q == ' ' {
			if s != 0 {
				if expression[s-1] == '(' && i < len(expression) {
					expression = expression[:s-1] + expression[s:i] + expression[i+1:]
					continue
				}
			}
			break
		}
		right := ""
		left := ""
		i = index - 1
		for i >= 0 {
			g := expression[i]
			if g == '(' {
				break
			} else if g == '+' || g == '*' || g == '/' {
				break
			} else if g == '-' {
				if i == s {
					left = "-" + left
					i--
				} else if expression[i-1] == '-' || expression[i-1] == '+' || expression[i-1] == '*' || expression[i-1] == '/' {
					left = "-" + left
					i--
				}
				break
			}
			left = expression[i:i+1] + left
			i--
		}
		lt := i + 1
		i = index + 1
		for i < len(expression) {
			g := expression[i]
			if g == ')' {
				break
			} else if g == '+' || g == '*' || g == '/' {
				break
			} else if g == '-' {
				if expression[i-1] == '-' || expression[i-1] == '+' || expression[i-1] == '*' || expression[i-1] == '/' {
				} else {
					break
				}
			}
			right = right + expression[i:i+1]
			i++
		}
		rt := i
		ans := ""
		lefted, _ := strconv.ParseFloat(left, 64)
		righted, _ := strconv.ParseFloat(right, 64)
		if q == '*' {
			ans = strconv.FormatFloat(lefted*righted, 'f', 6, 64)
		} else if q == '/' {
			if righted == 0 {
				return 0, errors.New("cannot be divided by zero")
			}
			ans = strconv.FormatFloat(lefted/righted, 'f', 6, 64)
		} else if q == '-' {
			ans = strconv.FormatFloat(lefted-righted, 'f', 6, 64)
			if len(left) == 0 {
				break
			}
		} else {
			ans = strconv.FormatFloat(lefted+righted, 'f', 6, 64)
		}
		expression = expression[:lt] + ans + expression[rt:]
	}
	ex, _ := strconv.ParseFloat(expression, 64)
	return ex, nil
}
