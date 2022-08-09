package parser

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/menxqk/my-interpreter/ast"
	"github.com/menxqk/my-interpreter/lexer"
	"github.com/menxqk/my-interpreter/token"
)

func TestParseProgram(t *testing.T) {
	tests := []struct {
		Line string
		Stmt ast.Statement
	}{
		{"int c = 50;", &ast.VariableDeclarationStatement{
			Identifier: ast.Identifier{Name: "c", Type: token.INT_TYPE, TypeLiteral: "int"},
			Expression: &ast.IntegerLiteral{Value: 50}},
		},
		{"string concat(string s1, string s2) { return s1 + s2; }", &ast.FunctionDeclarationStatement{
			Function: &ast.FunctionExpression{
				Identifier: ast.Identifier{Name: "concat", Type: token.STRING_TYPE, TypeLiteral: "string"},
				Parameters: []*ast.Identifier{
					{Name: "s1", Type: token.STRING_TYPE, TypeLiteral: "string"},
					{Name: "s2", Type: token.STRING_TYPE, TypeLiteral: "string"},
				},
				Body: &ast.BlockStatement{
					Statements: []ast.Statement{
						&ast.ReturnStatement{
							ReturnValue: &ast.InfixExpression{
								Left:     &ast.Identifier{Name: "s1"},
								Operator: "+",
								Right:    &ast.Identifier{Name: "s2"},
							},
						},
					},
				},
			}},
		},
		{"x = 20;", &ast.AssignmentStatement{
			Identifier: ast.Identifier{Name: "x"},
			Expression: &ast.IntegerLiteral{Value: 20}},
		},
		{"1 + 1;", &ast.ExpressionStatement{
			Expression: &ast.InfixExpression{
				Left:     &ast.IntegerLiteral{Value: 1},
				Operator: "+",
				Right:    &ast.IntegerLiteral{Value: 1},
			}},
		},
		{"concat(\"string1\",\"string2\");", &ast.ExpressionStatement{
			Expression: &ast.CallExpression{
				Identifier: ast.Identifier{Name: "concat"},
				Arguments: []ast.Expression{
					&ast.StringLiteral{Value: "string1"},
					&ast.StringLiteral{Value: "string2"},
				},
			}},
		},
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

	for i, tt := range tests {
		stmt := program.Statements[i]
		checkStatements(t, stmt, tt.Stmt)
	}

	tests = []struct {
		Line string
		Stmt ast.Statement
	}{
		{"int c = +", nil},
		{"string concat(-", nil},
		{"x = (+", nil},
		{"string s = /", nil},
		{"concat(*", nil},
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

func TestParseStatement(t *testing.T) {
	tests := []struct {
		Line string
		Stmt ast.Statement
	}{
		{"string s = \"A string\";", &ast.VariableDeclarationStatement{
			Identifier: ast.Identifier{Name: "s", Type: token.STRING_TYPE, TypeLiteral: "string"},
			Expression: &ast.StringLiteral{Value: "A string"}},
		},
		{"string help() { return \"help\";", &ast.FunctionDeclarationStatement{
			Function: &ast.FunctionExpression{
				Identifier: ast.Identifier{Name: "help", Type: token.STRING_TYPE, TypeLiteral: "string"},
				Parameters: []*ast.Identifier{},
				Body: &ast.BlockStatement{
					Statements: []ast.Statement{
						&ast.ReturnStatement{
							ReturnValue: &ast.StringLiteral{
								Value: "help",
							},
						},
					},
				},
			}},
		},
		{"int add(int a, int b) { return a + b; }", &ast.FunctionDeclarationStatement{
			Function: &ast.FunctionExpression{
				Identifier: ast.Identifier{Name: "add", Type: token.INT_TYPE, TypeLiteral: "int"},
				Parameters: []*ast.Identifier{
					{Name: "a", Type: token.INT_TYPE, TypeLiteral: "int"},
					{Name: "b", Type: token.INT_TYPE, TypeLiteral: "int"},
				},
				Body: &ast.BlockStatement{
					Statements: []ast.Statement{
						&ast.ReturnStatement{
							ReturnValue: &ast.InfixExpression{
								Left:     &ast.Identifier{Name: "a"},
								Operator: "+",
								Right:    &ast.Identifier{Name: "b"},
							},
						},
					},
				},
			}},
		},
		{"x = 20;", &ast.AssignmentStatement{
			Identifier: ast.Identifier{Name: "x"},
			Expression: &ast.IntegerLiteral{Value: 20}},
		},
		{"y = 35.60;", &ast.AssignmentStatement{
			Identifier: ast.Identifier{Name: "y"},
			Expression: &ast.FloatLiteral{Value: 35.60}},
		},
		{"s = \"string\";", &ast.AssignmentStatement{
			Identifier: ast.Identifier{Name: "s"},
			Expression: &ast.StringLiteral{Value: "string"}},
		},
		{"c = 'c';", &ast.AssignmentStatement{
			Identifier: ast.Identifier{Name: "c"},
			Expression: &ast.CharLiteral{Value: 'c'}},
		},
		{"d = 1 + 3;", &ast.AssignmentStatement{
			Identifier: ast.Identifier{Name: "d"},
			Expression: &ast.InfixExpression{
				Left:     &ast.IntegerLiteral{Value: 1},
				Operator: "+",
				Right:    &ast.IntegerLiteral{Value: 3},
			}},
		},
		{"abc;", &ast.ExpressionStatement{
			Expression: &ast.Identifier{Name: "abc"}},
		},
		{"!1;", &ast.ExpressionStatement{
			Expression: &ast.PrefixExpression{
				Operator:   "!",
				Expression: &ast.IntegerLiteral{Value: 1},
			}},
		},
		{"(2);", &ast.ExpressionStatement{
			Expression: &ast.GroupedExpression{
				Expression: &ast.IntegerLiteral{Value: 2}}},
		},
		{"3 + 3;", &ast.ExpressionStatement{
			Expression: &ast.InfixExpression{
				Left:     &ast.IntegerLiteral{Value: 3},
				Operator: "+",
				Right:    &ast.IntegerLiteral{Value: 3},
			}},
		},
		{"if (x < 1) { x = 2; } else { y = 3.60; }", &ast.ExpressionStatement{
			Expression: &ast.IfExpression{
				Condition: &ast.InfixExpression{
					Left:     &ast.Identifier{Name: "x"},
					Operator: "<",
					Right:    &ast.IntegerLiteral{Value: 1}},
				Consequence: &ast.BlockStatement{
					Statements: []ast.Statement{
						&ast.AssignmentStatement{
							Identifier: ast.Identifier{Name: "x"},
							Expression: &ast.IntegerLiteral{Value: 2},
						},
					},
				},
				Alternative: &ast.BlockStatement{
					Statements: []ast.Statement{
						&ast.AssignmentStatement{
							Identifier: ast.Identifier{Name: "y"},
							Expression: &ast.FloatLiteral{Value: 3.60},
						},
					},
				},
			}},
		},
		{"add(a, 1 + 1);", &ast.ExpressionStatement{
			Expression: &ast.CallExpression{
				Identifier: ast.Identifier{Name: "add"},
				Arguments: []ast.Expression{
					&ast.Identifier{Name: "a"},
					&ast.InfixExpression{
						Left:     &ast.IntegerLiteral{Value: 1},
						Operator: "+",
						Right:    &ast.IntegerLiteral{Value: 1},
					},
				},
			}},
		},
		{"1;", &ast.ExpressionStatement{Expression: &ast.IntegerLiteral{Value: 1}}},
		{"3.560;", &ast.ExpressionStatement{Expression: &ast.FloatLiteral{Value: 3.560}}},
		{"'c';", &ast.ExpressionStatement{Expression: &ast.CharLiteral{Value: 'c'}}},
		{"\"A string\";", &ast.ExpressionStatement{Expression: &ast.StringLiteral{Value: "A string"}}},
		{"return a;", &ast.ReturnStatement{
			ReturnValue: &ast.Identifier{Name: "a"}},
		},
		{"return a + 1;", &ast.ReturnStatement{
			ReturnValue: &ast.InfixExpression{
				Left:     &ast.Identifier{Name: "a"},
				Operator: "+",
				Right:    &ast.IntegerLiteral{Value: 1},
			}},
		},
	}
	for _, tt := range tests {
		l := lexer.New(tt.Line)
		p := New(l)
		stmt := p.parseStatement()

		if len(p.errors) != 0 {
			t.Errorf("expected zero errors, got %d", len(p.errors))
			fmt.Println(p.errors)
			fmt.Println(stmt)
		}

		checkStatements(t, stmt, tt.Stmt)
	}

	tests = []struct {
		Line string
		Stmt ast.Statement
	}{
		{"string s = \"A string\"", nil},
		{"string help() { return \"help\"", nil},
		{"int add(a)", nil},
		{"x = 20", nil},
		{"y = 35.60", nil},
		{"s = \"string\"", nil},
		{"c = 'c'", nil},
		{"d = 1 + 3", nil},
		{"abc", nil},
		{"!1", nil},
		{"(2)", nil},
		{"3 + 3", nil},
		{"if (x < )", nil},
		{"add(a, 1 + 1)", nil},
		{"1", nil},
		{"3.560", nil},
		{"'c'", nil},
		{"\"A string\"", nil},
		{"return a", nil},
		{"return a + 1", nil},
		{"{ int x = 1\nfloat y = 3.2\nc = 'c'\nstring s = \"a string\"\nreturn x / y }", nil},
	}
	for _, tt := range tests {
		l := lexer.New(tt.Line)
		p := New(l)
		// program := p.ParseProgram()
		stmt := p.parseStatement()

		if len(p.errors) != 1 {
			t.Errorf("expected 1 error, got %d", len(p.errors))
		}

		if stmt != nil && !reflect.ValueOf(stmt).IsNil() {
			t.Errorf("expected nil statement, got %T", stmt)
		}
	}

}

func TestParseExpression(t *testing.T) {
	tests := []struct {
		Line string
		Exp  ast.Expression
	}{
		{"abc", &ast.Identifier{Name: "abc"}},
		{"1", &ast.IntegerLiteral{Value: 1}},
		{"28.90", &ast.FloatLiteral{Value: 28.90}},
		{"'c'", &ast.CharLiteral{Value: 'c'}},
		{"\"a string\"", &ast.StringLiteral{Value: "a string"}},
		{"!2", &ast.PrefixExpression{Operator: "!", Expression: &ast.IntegerLiteral{Value: 2}}},
		{"-3.80", &ast.PrefixExpression{Operator: "-", Expression: &ast.FloatLiteral{Value: 3.80}}},
		{"(4)", &ast.GroupedExpression{Expression: &ast.IntegerLiteral{Value: 4}}},
		{"5 + 6", &ast.InfixExpression{
			Left:     &ast.IntegerLiteral{Value: 5},
			Operator: "+",
			Right:    &ast.IntegerLiteral{Value: 6}},
		},
		{"if (x < 1) { x = x + 1; } else { x = x + 2; }", &ast.IfExpression{
			Condition: &ast.InfixExpression{
				Left:     &ast.Identifier{Name: "x"},
				Operator: "<",
				Right:    &ast.IntegerLiteral{Value: 1},
			},
			Consequence: &ast.BlockStatement{
				Statements: []ast.Statement{
					&ast.AssignmentStatement{
						Identifier: ast.Identifier{Name: "x"},
						Expression: &ast.InfixExpression{
							Left:     &ast.Identifier{Name: "x"},
							Operator: "+",
							Right:    &ast.IntegerLiteral{Value: 1},
						},
					},
				},
			},
			Alternative: &ast.BlockStatement{
				Statements: []ast.Statement{
					&ast.AssignmentStatement{
						Identifier: ast.Identifier{Name: "x"},
						Expression: &ast.InfixExpression{
							Left:     &ast.Identifier{Name: "x"},
							Operator: "+",
							Right:    &ast.IntegerLiteral{Value: 2},
						},
					},
				},
			}},
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.Line)
		p := New(l)
		exp := p.parseExpression(LOWEST)

		if exp == nil {
			t.Fatalf("parseExpression returned nil for %s", tt.Exp.String())
		}

		checkExpressions(t, exp, tt.Exp)
	}
}

func checkStatements(t *testing.T, stmt ast.Statement, ttStmt ast.Statement) {
	if reflect.TypeOf(stmt) != reflect.TypeOf(ttStmt) {
		t.Errorf("expected %s, got %s", reflect.TypeOf(ttStmt), reflect.TypeOf(stmt))
		return
	}

	switch stmt := stmt.(type) {
	case *ast.ExpressionStatement:
		ttStmt := ttStmt.(*ast.ExpressionStatement)
		checkExpressions(t, stmt.Expression, ttStmt.Expression)
	case *ast.BlockStatement:
		ttStmt := ttStmt.(*ast.BlockStatement)
		if len(stmt.Statements) != len(ttStmt.Statements) {
			t.Errorf("expected '%d' statements, got '%d'", len(ttStmt.Statements), len(stmt.Statements))
		}
		for i, stmt := range stmt.Statements {
			ttStmt := ttStmt.Statements[i]
			checkStatements(t, stmt, ttStmt)
		}
	case *ast.VariableDeclarationStatement:
		ttStmt := ttStmt.(*ast.VariableDeclarationStatement)
		checkExpressions(t, &stmt.Identifier, &ttStmt.Identifier)
		checkExpressions(t, stmt.Expression, ttStmt.Expression)
	case *ast.FunctionDeclarationStatement:
		ttStmt := ttStmt.(*ast.FunctionDeclarationStatement)
		checkExpressions(t, stmt.Function, ttStmt.Function)
	case *ast.AssignmentStatement:
		ttStmt := ttStmt.(*ast.AssignmentStatement)
		checkExpressions(t, &stmt.Identifier, &ttStmt.Identifier)
		checkExpressions(t, stmt.Expression, ttStmt.Expression)
	case *ast.ReturnStatement:
		ttStmt := ttStmt.(*ast.ReturnStatement)
		checkExpressions(t, stmt.ReturnValue, ttStmt.ReturnValue)
	}
}

func checkExpressions(t *testing.T, exp ast.Expression, ttExp ast.Expression) {
	if reflect.TypeOf(exp) != reflect.TypeOf(ttExp) {
		t.Errorf("expected %s, got %s", reflect.TypeOf(ttExp), reflect.TypeOf(exp))
		return
	}

	switch exp := exp.(type) {
	case *ast.Identifier:
		ttExp := ttExp.(*ast.Identifier)
		if exp.Name != ttExp.Name {
			t.Errorf("expected Name '%s', got '%s'", ttExp.Name, exp.Name)
		}
		if exp.Type != ttExp.Type {
			t.Errorf("expected Type '%s', got '%s'", ttExp.Type, exp.Type)
		}
		if exp.TypeLiteral != ttExp.TypeLiteral {
			t.Errorf("expected TypeLiteral '%s', got '%s'", ttExp.TypeLiteral, exp.TypeLiteral)
		}
	case *ast.IntegerLiteral:
		ttExp := ttExp.(*ast.IntegerLiteral)
		if exp.Value != ttExp.Value {
			t.Errorf("expected Value '%d', got '%d'", ttExp.Value, exp.Value)
		}
	case *ast.FloatLiteral:
		ttExp := ttExp.(*ast.FloatLiteral)
		if exp.Value != ttExp.Value {
			t.Errorf("expected Value '%f', got '%f'", ttExp.Value, exp.Value)
		}
	case *ast.CharLiteral:
		ttExp := ttExp.(*ast.CharLiteral)
		if exp.Value != ttExp.Value {
			t.Errorf("expected Value '%c', got '%c'", ttExp.Value, exp.Value)
		}
	case *ast.StringLiteral:
		ttExp := ttExp.(*ast.StringLiteral)
		if exp.Value != ttExp.Value {
			t.Errorf("expected Value '%s', got '%s'", ttExp.Value, exp.Value)
		}
	case *ast.PrefixExpression:
		ttExp := ttExp.(*ast.PrefixExpression)
		if exp.Operator != ttExp.Operator {
			t.Errorf("expected Operator '%s', got '%s'", ttExp.Operator, exp.Operator)
		}
		checkExpressions(t, exp.Expression, ttExp.Expression)
	case *ast.GroupedExpression:
		ttExp := ttExp.(*ast.GroupedExpression)
		checkExpressions(t, exp.Expression, ttExp.Expression)
	case *ast.InfixExpression:
		ttExp := ttExp.(*ast.InfixExpression)
		if exp.Operator != ttExp.Operator {
			t.Errorf("expected Operator '%s', got '%s'", ttExp.Operator, exp.Operator)
		}
		checkExpressions(t, exp.Left, ttExp.Left)
		checkExpressions(t, exp.Right, ttExp.Right)
	case *ast.IfExpression:
		ttExp := ttExp.(*ast.IfExpression)
		checkExpressions(t, exp.Condition, ttExp.Condition)
		if exp.Consequence != nil {
			if ttExp.Consequence == nil {
				t.Errorf("expected not nil consequece block")
			} else {
				checkStatements(t, exp.Consequence, ttExp.Consequence)
			}
		}
		if exp.Alternative != nil {
			if ttExp.Alternative == nil {
				t.Errorf("expected not nil alternative block")
			} else {
				checkStatements(t, exp.Alternative, ttExp.Alternative)
			}
		}
	case *ast.FunctionExpression:
		ttExp := ttExp.(*ast.FunctionExpression)
		checkExpressions(t, &exp.Identifier, &ttExp.Identifier)
		if len(exp.Parameters) != len(ttExp.Parameters) {
			t.Errorf("expected %d parameter, got %d", len(ttExp.Parameters), len(exp.Parameters))
		} else {
			for i, param := range exp.Parameters {
				ttParam := ttExp.Parameters[i]
				checkExpressions(t, param, ttParam)
			}
		}
		checkStatements(t, exp.Body, ttExp.Body)
	case *ast.CallExpression:
		ttExp := ttExp.(*ast.CallExpression)
		checkExpressions(t, &exp.Identifier, &ttExp.Identifier)
		if len(exp.Arguments) != len(ttExp.Arguments) {
			t.Errorf("expected %d arguments, got %d", len(ttExp.Arguments), len(exp.Arguments))
		} else {
			for i, arg := range exp.Arguments {
				ttArg := ttExp.Arguments[i]
				checkExpressions(t, arg, ttArg)
			}
		}
	}
}
