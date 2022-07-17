package parser

import (
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
	case token.SUBTRACT:
		return p.parseSubtractStatement()
	case token.MULTIPLY:
		return p.parseMultiplyStatement()
	case token.DIVIDE:
		return p.parseDivideStatement()
	case token.MODULUS:
		return p.parseModulusStatement()
	case token.JUMP:
		return p.parseJumpStatement()
	case token.JUMP_ZERO:
		return p.parseJumpZeroStatement()
	case token.JUMP_NOT_ZERO:
		return p.parseJumpNotZeroStatement()
	case token.DECREMENT_JUMP:
		return p.parseDecrementJumpStatement()
	case token.SPLIT:
		return p.parseSplitStatement()
	case token.COMPARE:
		return p.parseCompareStatement()
	case token.SKIP_EQUAL:
		return p.parseSkipEqualStatement()
	case token.SKIP_NOT_EQUAL:
		return p.parseSkipNotEqualStatement()
	case token.SKIP_LOWER:
		return p.parseSkipLowerStatement()
	case token.LOADP:
		return p.parseLoadStatement()
	case token.SAVEP:
		return p.parseSaveStatement()
	case token.NOOP:
		return p.parseNoopStatement()
	default:
		return nil
	}
}

func (p *Parser) parseDataStatement() ast.Statement {
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

func (p *Parser) parseSubtractStatement() ast.Statement {
	stmt := &ast.SubtractStatement{Token: p.curToken}
	p.nextToken()
	stmt.AField = p.parseExpression()

	if !p.expectPeek(token.COMMA) {
		return nil
	}
	p.nextToken()

	stmt.BField = p.parseExpression()

	return stmt
}

func (p *Parser) parseMultiplyStatement() ast.Statement {
	stmt := &ast.MultiplyStatement{Token: p.curToken}
	p.nextToken()
	stmt.AField = p.parseExpression()

	if !p.expectPeek(token.COMMA) {
		return nil
	}
	p.nextToken()

	stmt.BField = p.parseExpression()

	return stmt
}

func (p *Parser) parseDivideStatement() ast.Statement {
	stmt := &ast.DivideStatement{Token: p.curToken}
	p.nextToken()
	stmt.AField = p.parseExpression()

	if !p.expectPeek(token.COMMA) {
		return nil
	}
	p.nextToken()

	stmt.BField = p.parseExpression()

	return stmt
}

func (p *Parser) parseModulusStatement() ast.Statement {
	stmt := &ast.ModStatement{Token: p.curToken}
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

func (p *Parser) parseJumpZeroStatement() ast.Statement {
	stmt := &ast.JumpZeroStatement{Token: p.curToken}
	p.nextToken()
	stmt.AField = p.parseExpression()

	if !p.expectPeek(token.COMMA) {
		return nil
	}
	p.nextToken()

	stmt.BField = p.parseExpression()

	return stmt
}

func (p *Parser) parseJumpNotZeroStatement() ast.Statement {
	stmt := &ast.JumpNonzeroStatement{Token: p.curToken}
	p.nextToken()
	stmt.AField = p.parseExpression()

	if !p.expectPeek(token.COMMA) {
		return nil
	}
	p.nextToken()

	stmt.BField = p.parseExpression()

	return stmt
}

func (p *Parser) parseDecrementJumpStatement() ast.Statement {
	stmt := &ast.DecrementJumpStatement{Token: p.curToken}
	p.nextToken()
	stmt.AField = p.parseExpression()

	if !p.expectPeek(token.COMMA) {
		return nil
	}
	p.nextToken()

	stmt.BField = p.parseExpression()

	return stmt
}

func (p *Parser) parseSplitStatement() ast.Statement {
	stmt := &ast.SplitStatement{Token: p.curToken}
	p.nextToken()

	stmt.AField = p.parseExpression()
	return stmt
}

func (p *Parser) parseCompareStatement() ast.Statement {
	stmt := &ast.CompareStatement{Token: p.curToken}
	p.nextToken()
	stmt.AField = p.parseExpression()

	if !p.expectPeek(token.COMMA) {
		return nil
	}
	p.nextToken()

	stmt.BField = p.parseExpression()

	return stmt
}

func (p *Parser) parseSkipEqualStatement() ast.Statement {
	stmt := &ast.SkipStatement{Token: p.curToken}
	p.nextToken()
	stmt.AField = p.parseExpression()

	if !p.expectPeek(token.COMMA) {
		return nil
	}
	p.nextToken()

	stmt.BField = p.parseExpression()

	return stmt
}

func (p *Parser) parseSkipNotEqualStatement() ast.Statement {
	stmt := &ast.SkipNeqStatement{Token: p.curToken}
	p.nextToken()
	stmt.AField = p.parseExpression()

	if !p.expectPeek(token.COMMA) {
		return nil
	}
	p.nextToken()

	stmt.BField = p.parseExpression()

	return stmt
}

func (p *Parser) parseSkipLowerStatement() ast.Statement {
	stmt := &ast.SkipLowerStatement{Token: p.curToken}
	p.nextToken()
	stmt.AField = p.parseExpression()

	if !p.expectPeek(token.COMMA) {
		return nil
	}
	p.nextToken()

	stmt.BField = p.parseExpression()

	return stmt
}

func (p *Parser) parseLoadStatement() ast.Statement {
	stmt := &ast.LoadStatement{Token: p.curToken}
	p.nextToken()
	stmt.AField = p.parseExpression()

	if !p.expectPeek(token.COMMA) {
		return nil
	}
	p.nextToken()

	stmt.BField = p.parseExpression()

	return stmt
}

func (p *Parser) parseSaveStatement() ast.Statement {
	stmt := &ast.SaveStatement{Token: p.curToken}
	p.nextToken()
	stmt.AField = p.parseExpression()

	if !p.expectPeek(token.COMMA) {
		return nil
	}
	p.nextToken()

	stmt.BField = p.parseExpression()

	return stmt
}

func (p *Parser) parseNoopStatement() ast.Statement {
	stmt := &ast.NopStatement{Token: p.curToken}
	p.nextToken()
	return stmt
}
