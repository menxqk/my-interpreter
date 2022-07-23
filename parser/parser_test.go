package parser

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/menxqk/my-interpreter/ast"
	"github.com/menxqk/my-interpreter/lexer"
)

func TestParseProgram(t *testing.T) {
	tests := []struct {
		Line string
		Stmt ast.Statement
	}{
		{"int c = 50;", &ast.VariableDeclarationStatement{}},
		{"float add(float a, float b) { a + b; }", &ast.FunctionDeclarationStatement{}},
		{"x = 20;", &ast.AssignmentStatement{}},
		{"1 + 1;", &ast.ExpressionStatement{}},
	}

	lines := []string{}
	for _, input := range tests {
		lines = append(lines, input.Line)
	}

	l := lexer.New(strings.Join(lines, "\n"))
	p := New(l)
	program := p.ParseProgram()

	if p.HasErrors() {
		var out bytes.Buffer
		for i, e := range p.errors {
			out.WriteString(fmt.Sprintf("%d - %s\n", i, e))
		}
		t.Fatalf("expected zero errors, got=%d\n%s", len(p.errors), out.String())
	}

	if len(program.Statements) != len(tests) {
		t.Fatalf("expected %d statements, got=%d", len(tests), len(program.Statements))
	}

	for i, input := range tests {
		stmt := program.Statements[i]
		if reflect.TypeOf(stmt) != reflect.TypeOf(input.Stmt) {
			t.Fatalf("expected type %s, got=%s", reflect.TypeOf(input.Stmt), reflect.TypeOf(stmt))
		}
	}

	tests = []struct {
		Line string
		Stmt ast.Statement
	}{
		{"x = 20", nil},
		{"string s =", nil},
	}

	lines = []string{}
	for _, input := range tests {
		lines = append(lines, input.Line)
	}

	l = lexer.New(strings.Join(lines, "\n"))
	p = New(l)
	program = p.ParseProgram()

	if len(p.errors) != len(tests) {
		t.Fatalf("expected %d error, got=%d", len(tests), len(p.errors))
	}

	if len(program.Statements) != 0 {
		t.Fatalf("expected 0 statements, got=%d", len(program.Statements))
	}
}

func TestParseDeclarationStatements(t *testing.T) {
	tests := []struct {
		Line string
		Stmt ast.Statement
	}{
		{"string s = \"A string\";", &ast.VariableDeclarationStatement{}},
		{"string help() { return \"help\";", &ast.FunctionDeclarationStatement{}},
	}

	lines := []string{}
	for _, input := range tests {
		lines = append(lines, input.Line)
	}

	l := lexer.New(strings.Join(lines, "\n"))
	p := New(l)
	program := p.ParseProgram()

	if p.HasErrors() {
		var out bytes.Buffer
		for i, e := range p.errors {
			out.WriteString(fmt.Sprintf("%d - %s\n", i, e))
		}
		t.Fatalf("expected zero errors, got=%d\n%s", len(p.errors), out.String())
	}

	if len(program.Statements) != len(tests) {
		t.Fatalf("expected %d statements, got=%d", len(tests), len(program.Statements))
	}

	for i, input := range tests {
		stmt := program.Statements[i]
		if reflect.TypeOf(stmt) != reflect.TypeOf(input.Stmt) {
			t.Fatalf("expected type %s, got=%s", reflect.TypeOf(input.Stmt), reflect.TypeOf(stmt))
		}
	}

	tests = []struct {
		Line string
		Stmt ast.Statement
	}{
		{"string s = \"A string\"", nil},
		{"string help() { return \"help\"", nil},
	}

	lines = []string{}
	for _, input := range tests {
		lines = append(lines, input.Line)
	}

	l = lexer.New(strings.Join(lines, "\n"))
	p = New(l)
	program = p.ParseProgram()

	if len(p.errors) != len(tests) {
		t.Fatalf("expected %d error, got=%d", len(tests), len(p.errors))
	}

	if len(program.Statements) != 0 {
		t.Fatalf("expected 0 statements, got=%d", len(program.Statements))
	}
}

func TestAssignmentStatement(t *testing.T) {
	tests := []struct {
		Line string
		Stmt ast.Statement
	}{
		{"x = 20;", &ast.AssignmentStatement{}},
		{"y = 35.60;", &ast.AssignmentStatement{}},
		{"s = \"string\";", &ast.AssignmentStatement{}},
		{"c = 'c';", &ast.AssignmentStatement{}},
	}

	lines := []string{}
	for _, input := range tests {
		lines = append(lines, input.Line)
	}

	l := lexer.New(strings.Join(lines, "\n"))
	p := New(l)
	program := p.ParseProgram()

	if p.HasErrors() {
		var out bytes.Buffer
		for i, e := range p.errors {
			out.WriteString(fmt.Sprintf("%d - %s\n", i, e))
		}
		t.Fatalf("expected zero errors, got=%d\n%s", len(p.errors), out.String())
	}

	if len(program.Statements) != len(tests) {
		t.Fatalf("expected %d statements, got=%d", len(tests), len(program.Statements))
	}

	for i, input := range tests {
		stmt := program.Statements[i]
		if reflect.TypeOf(stmt) != reflect.TypeOf(input.Stmt) {
			t.Fatalf("expected type %s, got=%s", reflect.TypeOf(input.Stmt), reflect.TypeOf(stmt))
		}
	}

	tests = []struct {
		Line string
		Stmt ast.Statement
	}{
		{"x = 20", nil},
		{"y = 35.60", nil},
		{"s = \"string\"", nil},
		{"c = 'c'", nil},
	}

	lines = []string{}
	for _, input := range tests {
		lines = append(lines, input.Line)
	}

	l = lexer.New(strings.Join(lines, "\n"))
	p = New(l)
	program = p.ParseProgram()

	if len(p.errors) != len(tests) {
		t.Fatalf("expected %d error, got=%d", len(tests), len(p.errors))
	}

	if len(program.Statements) != 0 {
		t.Fatalf("expected 0 statements, got=%d", len(program.Statements))
	}
}

func TestParseExpressionStatement(t *testing.T) {
	tests := []struct {
		Line string
		Stmt ast.Statement
		Exp  ast.Expression
	}{
		{"abc;", &ast.ExpressionStatement{}, &ast.Identifier{}},
		{"!1;", &ast.ExpressionStatement{}, &ast.PrefixExpression{}},
		{"(1);", &ast.ExpressionStatement{}, &ast.GroupedExpression{}},
		{"1+1;", &ast.ExpressionStatement{}, &ast.InfixExpression{}},

		{"if (x < 1) { x = 2; } else { y = 3; }", &ast.ExpressionStatement{}, &ast.IfExpression{}},
		{"add(a, 1);", &ast.ExpressionStatement{}, &ast.CallExpression{}},
		{"1;", &ast.ExpressionStatement{}, &ast.IntegerLiteral{}},
		{"3.560;", &ast.ExpressionStatement{}, &ast.FloatLiteral{}},
		{"'c';", &ast.ExpressionStatement{}, &ast.CharLiteral{}},
		{"\"A string\";", &ast.ExpressionStatement{}, &ast.StringLiteral{}},
	}

	lines := []string{}
	for _, input := range tests {
		lines = append(lines, input.Line)
	}

	l := lexer.New(strings.Join(lines, "\n"))
	p := New(l)
	program := p.ParseProgram()

	if p.HasErrors() {
		var out bytes.Buffer
		for i, e := range p.errors {
			out.WriteString(fmt.Sprintf("%d - %s\n", i, e))
		}
		t.Fatalf("expected zero errors, got=%d\n%s", len(p.errors), out.String())
	}

	if len(program.Statements) != len(tests) {
		t.Fatalf("expected %d statements, got=%d", len(tests), len(program.Statements))
	}

	for i, input := range tests {
		stmt := program.Statements[i]
		if reflect.TypeOf(stmt) != reflect.TypeOf(input.Stmt) {
			t.Fatalf("expected type %s, got=%s", reflect.TypeOf(input.Stmt), reflect.TypeOf(stmt))
		}

		switch stmt := stmt.(type) {
		case *ast.ExpressionStatement:
			exp := stmt.Expression
			if reflect.TypeOf(exp) != reflect.TypeOf(input.Exp) {
				t.Fatalf("expected %T for expression, got= %T", input.Exp, exp)
			}
		default:
			t.Fatalf("expected *ast.ExpressionStatement, got= %s", stmt)
		}
	}

	tests = []struct {
		Line string
		Stmt ast.Statement
		Exp  ast.Expression
	}{
		{"abc", &ast.ExpressionStatement{}, nil},
		{"!1", &ast.ExpressionStatement{}, nil},
		{"(1)", &ast.ExpressionStatement{}, nil},
		{"1+1", &ast.ExpressionStatement{}, nil},

		{"if;", &ast.ExpressionStatement{}, nil},
		{"add(a, 1)", &ast.ExpressionStatement{}, nil},
		{"1", &ast.ExpressionStatement{}, nil},
		{"3.560", &ast.ExpressionStatement{}, nil},
		{"'c'", &ast.ExpressionStatement{}, nil},
		{"\"A string\"", &ast.ExpressionStatement{}, nil},
	}

	lines = []string{}
	for _, input := range tests {
		lines = append(lines, input.Line)
	}

	l = lexer.New(strings.Join(lines, "\n"))
	p = New(l)
	program = p.ParseProgram()

	if len(p.errors) != len(tests) {
		t.Fatalf("expected %d error, got=%d", len(tests), len(p.errors))
	}

	if len(program.Statements) != 0 {
		t.Fatalf("expected 0 statements, got=%d", len(program.Statements))
	}
}

func TestParseReturnStatement(t *testing.T) {
	tests := []struct {
		Line string
		Stmt ast.Statement
	}{
		{"return a;", &ast.ReturnStatement{}},
		{"return a + 1;", &ast.ReturnStatement{}},
	}

	lines := []string{}
	for _, input := range tests {
		lines = append(lines, input.Line)
	}

	l := lexer.New(strings.Join(lines, "\n"))
	p := New(l)
	program := p.ParseProgram()

	if p.HasErrors() {
		var out bytes.Buffer
		for i, e := range p.errors {
			out.WriteString(fmt.Sprintf("%d - %s\n", i, e))
		}
		t.Fatalf("expected zero errors, got=%d\n%s", len(p.errors), out.String())
	}

	if len(program.Statements) != len(tests) {
		t.Fatalf("expected %d statements, got=%d", len(tests), len(program.Statements))
	}

	for i, input := range tests {
		stmt := program.Statements[i]
		if reflect.TypeOf(stmt) != reflect.TypeOf(input.Stmt) {
			t.Fatalf("expected type %s, got=%s", reflect.TypeOf(input.Stmt), reflect.TypeOf(stmt))
		}
	}

	tests = []struct {
		Line string
		Stmt ast.Statement
	}{
		{"return a", nil},
		{"return a + 1", nil},
	}

	lines = []string{}
	for _, input := range tests {
		lines = append(lines, input.Line)
	}

	l = lexer.New(strings.Join(lines, "\n"))
	p = New(l)
	program = p.ParseProgram()

	if len(p.errors) != len(tests) {
		t.Fatalf("expected %d error, got=%d", len(tests), len(p.errors))
	}

	if len(program.Statements) != 0 {
		t.Fatalf("expected 0 statements, got=%d", len(program.Statements))
	}
}

func TestParseBlockStatement(t *testing.T) {
	tests := []struct {
		Line string
		Stmt ast.Statement
	}{
		{"int x = 1;", &ast.VariableDeclarationStatement{}},
		{"float y = 3.2;", &ast.VariableDeclarationStatement{}},
		{"c = 4;", &ast.AssignmentStatement{}},
		{"string s = \"a string\";", &ast.VariableDeclarationStatement{}},
	}

	lines := []string{}
	for _, input := range tests {
		lines = append(lines, input.Line)
	}

	l := lexer.New("{\n" + strings.Join(lines, "\n") + "}")
	p := New(l)

	block := p.parseBlockStatement()
	if block == nil {
		t.Error("could not parse block statement")
	}

	if p.HasErrors() {
		var out bytes.Buffer
		for i, e := range p.errors {
			out.WriteString(fmt.Sprintf("%d - %s\n", i, e))
		}
		t.Fatalf("expected zero errors, got=%d\n%s", len(p.errors), out.String())
	}

	if len(block.Statements) != len(tests) {
		t.Fatalf("expected %d statements in the block, got=%d", len(tests), len(block.Statements))
	}

	for i, input := range tests {
		stmt := block.Statements[i]
		if reflect.TypeOf(stmt) != reflect.TypeOf(input.Stmt) {
			t.Fatalf("expected type %s, got=%s", reflect.TypeOf(input.Stmt), reflect.TypeOf(stmt))
		}
	}

	tests = []struct {
		Line string
		Stmt ast.Statement
	}{
		{"int x = ;", nil},
	}

	lines = []string{}
	for _, input := range tests {
		lines = append(lines, input.Line)
	}

	l = lexer.New("{\n" + strings.Join(lines, "\n") + "}")
	p = New(l)

	block = p.parseBlockStatement()
	if block != nil {
		t.Error(fmt.Sprintf("block should be nil but is [%T]", block))
	}

	if len(p.errors) != len(tests) {
		t.Fatalf("expected %d error, got=%d", len(tests), len(p.errors))
	}
}
