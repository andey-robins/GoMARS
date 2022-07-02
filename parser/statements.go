package parser

import (
	"fmt"

	"github.com/andey-robins/GoMARS/ast"
	"github.com/andey-robins/GoMARS/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.DATA:
		return p.parseDataStatement()
	case token.MOVE:
		return p.parseMoveStatement()
	case token.ADD:
		return p.parseAddStatement()
	case token.JUMP:
		return p.parseJumpStatement()
	default:
		return nil
	}
}

func (p *Parser) parseDataStatement() ast.Statement {
	fmt.Println("Parsing data statement")
	stmt := &ast.DataStatement{Token: p.curToken}
	p.nextToken()
	stmt.AField = p.parseExpression()

	if !p.expectPeek(token.COMMA) {
		return nil
	}
	p.nextToken()

	stmt.BField = p.parseExpression()

	return stmt
}

func (p *Parser) parseMoveStatement() ast.Statement {
	stmt := &ast.MoveStatement{Token: p.curToken}
	p.nextToken()
	stmt.AField = p.parseExpression()

	if !p.expectPeek(token.COMMA) {
		return nil
	}
	p.nextToken()

	stmt.BField = p.parseExpression()

	return stmt
}

func (p *Parser) parseAddStatement() ast.Statement {
	stmt := &ast.AddStatement{Token: p.curToken}
	p.nextToken()
	stmt.AField = p.parseExpression()

	if !p.expectPeek(token.COMMA) {
		return nil
	}
	p.nextToken()

	stmt.BField = p.parseExpression()

	return stmt
}

func (p *Parser) parseJumpStatement() ast.Statement {
	stmt := &ast.JumpStatement{Token: p.curToken}
	p.nextToken()

	stmt.AField = p.parseExpression()
	return stmt
}
