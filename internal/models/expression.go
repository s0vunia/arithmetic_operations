package models

import "time"

type Expression struct {
	Id          int        `json:"id"`
	Expression  string     `json:"expression"`
	Answer      string     `json:"answer"`
	Status      Status     `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	CompletedAt *time.Time `json:"completedAt,omitempty"`
}

type SubExpression struct {
	Id          string
	Index       int
	Leftnum     float64
	Rightnum    float64
	Operator    string
	ExtraAction string
	Value       float64
	IsDone      bool
}

type InputExpression struct {
	Expression string `json:"expression" validate:"required,expression"`
}

func NewExpressionInProcess(expression string) *Expression {
	exp := NewExpression(expression)
	exp.Status = InProcess
	return exp
}

func NewExpressionInvalid(expression string) *Expression {
	exp := NewExpression(expression)
	exp.Status = Invalid
	return exp
}

func NewExpression(expression string) *Expression {
	return &Expression{
		Expression: expression,
		CreatedAt:  time.Now(),
	}
}
