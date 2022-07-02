package parser

import (
	"fmt"
	"strconv"

	"github.com/andey-robins/GoMARS/ast"
	"github.com/andey-robins/GoMARS/token"
)

func (p *Parser) parseExpression() ast.Expression {
	switch p.curToken.Type {
	case token.HASH:
		return p.parseImmediateExpression()
	case token.AT:
		return p.parseIndirectExpression()
	default:
		return p.parseDirectExpression()
	}
}

func (p *Parser) parseImmediateExpression() ast.Expression {
	if !p.curTokenIs(token.HASH) {
		return nil
	}
	stmt := &ast.ImmediateExpression{Token: p.curToken}
	p.nextToken()
	if val := p.parseIntegerLiteral(); val != nil {
		stmt.Value = *val
	} else {
		return nil
	}
	return stmt
}

func (p *Parser) parseIndirectExpression() ast.Expression {
	if !p.curTokenIs(token.AT) {
		return nil
	}
	stmt := &ast.IndirectExpression{Token: p.curToken}
	p.nextToken()
	if val := p.parseIntegerLiteral(); val != nil {
		stmt.Value = *val
	} else {
		return nil
	}
	return stmt
}

func (p *Parser) parseDirectExpression() ast.Expression {
	stmt := &ast.DirectExpression{Token: p.curToken}
	if val := p.parseIntegerLiteral(); val != nil {
		stmt.Value = *val
	} else {
		return nil
	}

	return stmt
}

func (p *Parser) parseIntegerLiteral() *ast.IntegerLiteral {

	negative := false
	if p.curTokenIs(token.MINUS) {
		negative = true
		p.nextToken()
	}

	lit := &ast.IntegerLiteral{Token: p.curToken}

	numString := p.curToken.Literal
	if negative {
		numString = "-" + numString
	}

	value, err := strconv.ParseInt(numString, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}
