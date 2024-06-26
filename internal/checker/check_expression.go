package checker

import (
	"arithmetic_operations/internal/stack"
	"arithmetic_operations/internal/topostfix"
	"errors"
	"fmt"
	"log/slog"
	"regexp"
	"strings"
)

// RemoveAllSpaces убирает пробелы в выражении
func RemoveAllSpaces(a string) string {
	r := regexp.MustCompile(`\s+`)
	noSpaces := r.ReplaceAllString(a, "")

	return noSpaces
}

// CheckExpression проверяет на все возможные ошибки
func CheckExpression(log *slog.Logger, expression string) error {
	log.Info("start check expression", slog.String("expr", expression))
	if len(expression) == 0 {
		log.Error("length of expression is 0", slog.String("expr", expression))
		return errors.New("length of expression is 0")
	}
	RemoveAllSpaces(expression)
	var errSlice []error
	errSlice = append(errSlice, HasDoubleSymbol(expression))
	errSlice = append(errSlice, ExpressionStartsWithNumber(expression))
	errSlice = append(errSlice, IsValidParentheses(expression))
	errSlice = append(errSlice, HasDivizionByZero(expression))
	errSlice = append(errSlice, HasValidCharacters(expression))
	errSlice = append(errSlice, HasAtLeastOneExpression(expression))
	errSlice = append(errSlice, ContainsCorrectFloatPoint(expression))

	if len(errSlice) == 0 {
		return nil
	} else {
		return errSlice[0]
	}

}

// HasDoubleSymbol проверяет на двойной символ
func HasDoubleSymbol(s string) error {
	for i := 0; i < len(s); i++ {
		if i > 0 && (s[i] == '*' && s[i-1] == '*' || s[i] == '/' && s[i-1] == '/' || s[i] == '+' && s[i-1] == '+' || s[i] == '-' && s[i-1] == '-') {
			return errors.New("expression has doubled symbol")
		}
	}
	return nil
}

// IsValidParentheses проверяет скобочную последовательность
func IsValidParentheses(s string) error {
	stack := &stack.Stack{}
	countOpen := 0
	countClose := 0

	for _, char := range s {
		if char == '(' {
			countOpen++
		} else if char == ')' {
			countClose++
		}
	}

	if countOpen != countClose {
		return errors.New("expression has invalid parentheses")
	}

	for _, r := range s {
		switch r {
		case '(':
			stack.Push('(')
		case ')':
			if stack.Empty() {
				return errors.New("expression has invalid parentheses")
			}
			stack.Pop()
		default:
			continue // Skip non-parentheses characters
		}
	}

	return nil
}

// HasDivizionByZero проверяет есть ли деление на ноль
func HasDivizionByZero(a string) error {
	if strings.Contains(a, "/0") {
		return errors.New("expression has division by zero")
	}
	return nil
}

// HasValidCharacters проверяет на допустимые символы
func HasValidCharacters(a string) error {
	validChars := "1234567890+-*/(). " // Список допустимых символов
	for _, char := range a {
		//(string(s[j]) == "." && number != "")
		if !strings.ContainsRune(validChars, char) {
			return errors.New(fmt.Sprintf("expression has invalid character %d", char))
		}
	}
	return nil
}

// ContainsCorrectFloatPoint проверяет на точку в правильном месте(должно быть в числе float)
func ContainsCorrectFloatPoint(expr string) error {
	length := len(expr)
	for i := 0; i < length; i++ {
		if string(expr[0]) == "." || string(expr[length-1]) == "." {
			return errors.New("expression has a dot in a wrong place")
		}
		if i > 0 && length-1 > i {
			if string(expr[i]) == "." {
				if !(topostfix.IsOperand(expr[i-1]) && topostfix.IsOperand(expr[i+1])) {
					return errors.New("expression has a dot in a wrong place")
				}
			}
		}
	}

	return nil
}

// HasAtLeastOneExpression проверяет на хотя бы одно выражение число оператор число
func HasAtLeastOneExpression(expr string) error {
	// Regular expression to match the pattern "number operator number"
	pattern := `\d+\s*[\+\-\*/]\s*\d+`
	r := regexp.MustCompile(pattern)
	// Check if the expression contains at least one match
	matched := r.MatchString(expr)
	if matched {
		return nil
	} else {
		return errors.New("this string doesn't has at least one expression")
	}

}

// ExpressionStartsWithNumber проверяет - первым ли идёт число или скобка в выражении
func ExpressionStartsWithNumber(expr string) error {
	// Регулярное выражение для проверки начала строки с числом или скобкой
	regexPattern := `^(\()?\d+`
	matched, _ := regexp.MatchString(regexPattern, expr)
	if matched {
		return nil
	} else {
		return errors.New("this string doesn't start with number or bracket")
	}

}
