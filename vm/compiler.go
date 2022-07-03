package vm

import (
	"errors"

	"github.com/andey-robins/GoMARS/ast"
	"github.com/andey-robins/GoMARS/lexer"
	"github.com/andey-robins/GoMARS/parser"
)

func compile(prog string) ([]Cell, error) {
	mem := make([]Cell, 0)

	l := lexer.New(prog)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		return nil, errors.New("error compiling program")
	}

	// convert AST into memory
	for _, stmt := range program.Statements {
		switch stmt.(type) {
		case *ast.DataStatement:
			mem = append(mem, programDataStatement(stmt.(*ast.DataStatement)))
		case *ast.MoveStatement:
			mem = append(mem, programMoveStatement(stmt.(*ast.MoveStatement)))
		case *ast.AddStatement:
			mem = append(mem, programAddStatement(stmt.(*ast.AddStatement)))
		case *ast.SubtractStatement:
			mem = append(mem, programSubtractStatement(stmt.(*ast.SubtractStatement)))
		case *ast.MultiplyStatement:
			mem = append(mem, programMultiplyStatement(stmt.(*ast.MultiplyStatement)))
		case *ast.DivideStatement:
			mem = append(mem, programDivideStatement(stmt.(*ast.DivideStatement)))
		case *ast.ModStatement:
			mem = append(mem, programModStatement(stmt.(*ast.ModStatement)))
		case *ast.JumpStatement:
			mem = append(mem, programJumpStatement(stmt.(*ast.JumpStatement)))
		case *ast.JumpZeroStatement:
			mem = append(mem, programJumpZeroStatement(stmt.(*ast.JumpZeroStatement)))
		case *ast.JumpNonzeroStatement:
			mem = append(mem, programJumpNonzeroStatement(stmt.(*ast.JumpNonzeroStatement)))
		case *ast.DecrementJumpStatement:
			mem = append(mem, programDecrementJumpStatement(stmt.(*ast.DecrementJumpStatement)))
		case *ast.SplitStatement:
			mem = append(mem, programSplitStatement(stmt.(*ast.SplitStatement)))
		case *ast.CompareStatement:
			mem = append(mem, programCompareStatement(stmt.(*ast.CompareStatement)))
		case *ast.SkipStatement:
			mem = append(mem, programSkipStatement(stmt.(*ast.SkipStatement)))
		case *ast.SkipNeqStatement:
			mem = append(mem, programSkipNeqStatement(stmt.(*ast.SkipNeqStatement)))
		case *ast.SkipLowerStatement:
			mem = append(mem, programSkipLowerStatement(stmt.(*ast.SkipLowerStatement)))
		case *ast.LoadStatement:
			mem = append(mem, programLoadStatement(stmt.(*ast.LoadStatement)))
		case *ast.SaveStatement:
			mem = append(mem, programSaveStatement(stmt.(*ast.SaveStatement)))
		case *ast.NopStatement:
			mem = append(mem, programNopStatement(stmt.(*ast.NopStatement)))
		default:
			return nil, errors.New("encountered an error decoding ast")
		}
	}

	return mem, nil
}

func programDataStatement(stmt *ast.DataStatement) Cell {
	a, err := programExpression(stmt.AField)
	if err != nil {
		panic(err)
	}
	b, err := programExpression(stmt.BField)
	if err != nil {
		panic(err)
	}

	return Cell{
		Operation: DAT,
		AField:    a,
		BField:    b,
	}
}

func programMoveStatement(stmt *ast.MoveStatement) Cell {
	a, err := programExpression(stmt.AField)
	if err != nil {
		panic(err)
	}
	b, err := programExpression(stmt.BField)
	if err != nil {
		panic(err)
	}

	return Cell{
		Operation: MOV,
		AField:    a,
		BField:    b,
	}
}

func programAddStatement(stmt *ast.AddStatement) Cell {
	a, err := programExpression(stmt.AField)
	if err != nil {
		panic(err)
	}
	b, err := programExpression(stmt.BField)
	if err != nil {
		panic(err)
	}

	return Cell{
		Operation: ADD,
		AField:    a,
		BField:    b,
	}
}

func programSubtractStatement(stmt *ast.SubtractStatement) Cell {
	a, err := programExpression(stmt.AField)
	if err != nil {
		panic(err)
	}
	b, err := programExpression(stmt.BField)
	if err != nil {
		panic(err)
	}

	return Cell{
		Operation: SUB,
		AField:    a,
		BField:    b,
	}
}

func programMultiplyStatement(stmt *ast.MultiplyStatement) Cell {
	a, err := programExpression(stmt.AField)
	if err != nil {
		panic(err)
	}
	b, err := programExpression(stmt.BField)
	if err != nil {
		panic(err)
	}

	return Cell{
		Operation: MUL,
		AField:    a,
		BField:    b,
	}
}

func programDivideStatement(stmt *ast.DivideStatement) Cell {
	a, err := programExpression(stmt.AField)
	if err != nil {
		panic(err)
	}
	b, err := programExpression(stmt.BField)
	if err != nil {
		panic(err)
	}

	return Cell{
		Operation: DIV,
		AField:    a,
		BField:    b,
	}
}

func programModStatement(stmt *ast.ModStatement) Cell {
	a, err := programExpression(stmt.AField)
	if err != nil {
		panic(err)
	}
	b, err := programExpression(stmt.BField)
	if err != nil {
		panic(err)
	}

	return Cell{
		Operation: MOD,
		AField:    a,
		BField:    b,
	}
}

func programJumpStatement(stmt *ast.JumpStatement) Cell {
	a, err := programExpression(stmt.AField)
	if err != nil {
		panic(err)
	}

	return Cell{
		Operation: JMP,
		AField:    a,
		BField:    Field{},
	}
}

func programJumpZeroStatement(stmt *ast.JumpZeroStatement) Cell {
	a, err := programExpression(stmt.AField)
	if err != nil {
		panic(err)
	}
	b, err := programExpression(stmt.BField)
	if err != nil {
		panic(err)
	}

	return Cell{
		Operation: JMZ,
		AField:    a,
		BField:    b,
	}
}

func programJumpNonzeroStatement(stmt *ast.JumpNonzeroStatement) Cell {
	a, err := programExpression(stmt.AField)
	if err != nil {
		panic(err)
	}
	b, err := programExpression(stmt.BField)
	if err != nil {
		panic(err)
	}

	return Cell{
		Operation: JMN,
		AField:    a,
		BField:    b,
	}
}

func programDecrementJumpStatement(stmt *ast.DecrementJumpStatement) Cell {
	a, err := programExpression(stmt.AField)
	if err != nil {
		panic(err)
	}
	b, err := programExpression(stmt.BField)
	if err != nil {
		panic(err)
	}

	return Cell{
		Operation: DJN,
		AField:    a,
		BField:    b,
	}
}

func programSplitStatement(stmt *ast.SplitStatement) Cell {
	a, err := programExpression(stmt.AField)
	if err != nil {
		panic(err)
	}

	return Cell{
		Operation: SPL,
		AField:    a,
		BField:    Field{},
	}
}

func programCompareStatement(stmt *ast.CompareStatement) Cell {
	a, err := programExpression(stmt.AField)
	if err != nil {
		panic(err)
	}
	b, err := programExpression(stmt.BField)
	if err != nil {
		panic(err)
	}

	return Cell{
		Operation: CMP,
		AField:    a,
		BField:    b,
	}
}

func programSkipStatement(stmt *ast.SkipStatement) Cell {
	a, err := programExpression(stmt.AField)
	if err != nil {
		panic(err)
	}
	b, err := programExpression(stmt.BField)
	if err != nil {
		panic(err)
	}

	return Cell{
		Operation: SEQ,
		AField:    a,
		BField:    b,
	}
}

func programSkipNeqStatement(stmt *ast.SkipNeqStatement) Cell {
	a, err := programExpression(stmt.AField)
	if err != nil {
		panic(err)
	}
	b, err := programExpression(stmt.BField)
	if err != nil {
		panic(err)
	}

	return Cell{
		Operation: SNE,
		AField:    a,
		BField:    b,
	}
}

func programSkipLowerStatement(stmt *ast.SkipLowerStatement) Cell {
	a, err := programExpression(stmt.AField)
	if err != nil {
		panic(err)
	}
	b, err := programExpression(stmt.BField)
	if err != nil {
		panic(err)
	}

	return Cell{
		Operation: SLT,
		AField:    a,
		BField:    b,
	}
}

func programLoadStatement(stmt *ast.LoadStatement) Cell {
	a, err := programExpression(stmt.AField)
	if err != nil {
		panic(err)
	}
	b, err := programExpression(stmt.BField)
	if err != nil {
		panic(err)
	}

	return Cell{
		Operation: LDP,
		AField:    a,
		BField:    b,
	}
}

func programSaveStatement(stmt *ast.SaveStatement) Cell {
	a, err := programExpression(stmt.AField)
	if err != nil {
		panic(err)
	}
	b, err := programExpression(stmt.BField)
	if err != nil {
		panic(err)
	}

	return Cell{
		Operation: STP,
		AField:    a,
		BField:    b,
	}
}

func programNopStatement(stmt *ast.NopStatement) Cell {
	a, err := programExpression(stmt.AField)
	if err != nil {
		panic(err)
	}
	b, err := programExpression(stmt.BField)
	if err != nil {
		panic(err)
	}

	return Cell{
		Operation: NOP,
		AField:    a,
		BField:    b,
	}
}

func programExpression(exp interface{}) (Field, error) {
	switch exp.(type) {
	case *ast.ImmediateExpression:
		return programImmediateExpression(exp.(*ast.ImmediateExpression)), nil
	case *ast.DirectExpression:
		return programDirectExpression(exp.(*ast.DirectExpression)), nil
	case *ast.IndirectExpression:
		return programIndirectExpression(exp.(*ast.IndirectExpression)), nil
	default:
		return Field{}, errors.New("invalid expression type encountered")
	}
}

func programImmediateExpression(exp *ast.ImmediateExpression) Field {
	return Field{
		Mode:  IMMEDIATE,
		Value: int(exp.Value.Value),
	}
}

func programDirectExpression(exp *ast.DirectExpression) Field {
	return Field{
		Mode:  DIRECT,
		Value: int(exp.Value.Value),
	}
}

func programIndirectExpression(exp *ast.IndirectExpression) Field {
	return Field{
		Mode:  INDIRECT,
		Value: int(exp.Value.Value),
	}
}
