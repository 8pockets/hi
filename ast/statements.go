package ast

import (
	"bytes"
	"strconv"

	"github.com/8pockets/hi/token"
)

// Statements
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type VarStatement struct {
	Token token.Token // the token.VAR token
	Name  *Identifier
	Value Expression
}

func (vs *VarStatement) statementNode()       {}
func (vs *VarStatement) TokenLiteral() string { return vs.Token.Literal }
func (vs *VarStatement) String() string {
	var out bytes.Buffer

	out.WriteString(vs.TokenLiteral() + " ")
	out.WriteString(vs.Name.String())
	out.WriteString(" = ")

	if vs.Value != nil {
		out.WriteString(vs.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// ContinueStatement represents "continue" keyword
type ContinueStatement struct {
	Token        token.Token // the 'continue' token
	HierarchyNum int
}

func (cs *ContinueStatement) statementNode() {}

// TokenLiteral returns token's literal
func (cs *ContinueStatement) TokenLiteral() string {
	return cs.Token.Literal
}
func (cs *ContinueStatement) String() string {
	var out bytes.Buffer

	out.WriteString(cs.TokenLiteral())

	if cs.HierarchyNum != 0 {
		out.WriteString(" " + strconv.Itoa(cs.HierarchyNum))
	}

	return out.String()
}

// BreakStatement represents "break" keyword
type BreakStatement struct {
	Token        token.Token // the 'break' token
	HierarchyNum int
}

func (bs *BreakStatement) statementNode() {}

// TokenLiteral returns token's literal
func (bs *BreakStatement) TokenLiteral() string {
	return bs.Token.Literal
}
func (bs *BreakStatement) String() string {
	var out bytes.Buffer

	out.WriteString(bs.TokenLiteral())

	if bs.HierarchyNum != 0 {
		out.WriteString(" " + strconv.Itoa(bs.HierarchyNum))
	}

	return out.String()
}

// AssignStatement represents "=" keyword
type AssignStatement struct {
	Token token.Token // The '=' token
	Left  Expression
	Right Expression
}

func (ae *AssignStatement) statementNode()       {}
func (ae *AssignStatement) TokenLiteral() string { return ae.Token.Literal }
func (ae *AssignStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ae.Left.String())
	out.WriteString(" = ")
	out.WriteString(ae.Right.String())

	return out.String()
}
