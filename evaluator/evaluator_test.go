package evaluator

import (
	"testing"

	"github.com/menxqk/my-interpreter/lexer"
	"github.com/menxqk/my-interpreter/object"
	"github.com/menxqk/my-interpreter/parser"
)

func TestEval(t *testing.T) {
	tests := []struct {
		Line       string
		Result     string
		ResultType string
	}{
		{"x;", "null", object.NULL_OBJ},
		{"null;", "null", object.NULL_OBJ},

		{"true;", "true", object.BOOL_OBJ},
		{"false;", "false", object.BOOL_OBJ},
		{"!true;", "false", object.BOOL_OBJ},
		{"!(false);", "true", object.BOOL_OBJ},
		{"!(!true);", "true", object.BOOL_OBJ},
		{"!null;", "true", object.BOOL_OBJ},
		{"!(!null);", "false", object.BOOL_OBJ},

		{"-1;", "-1", object.INT_OBJ},
		{"-3.55;", "-3.550000", object.FLOAT_OBJ},
		{"-s;", "ERROR: unknown operator: -NULL", object.ERROR_OBJ},

		{"int x = 5;", "5", object.INT_OBJ},
		{"float y = 55.55;", "55.550000", object.FLOAT_OBJ},
		{"x;", "5", object.INT_OBJ},
		{"y;", "55.550000", object.FLOAT_OBJ},
		{"x + x;", "10", object.INT_OBJ},
		{"x + y;", "60.550000", object.FLOAT_OBJ},
		{"y + y;", "111.100000", object.FLOAT_OBJ},
		{"y + x;", "60.550000", object.FLOAT_OBJ},
		{"y - y;", "0.000000", object.FLOAT_OBJ},
		{"y - x;", "50.550000", object.FLOAT_OBJ},
		{"x - x;", "0", object.INT_OBJ},
		{"x - y;", "-50.550000", object.FLOAT_OBJ},
		{"x * x;", "25", object.INT_OBJ},
		{"x * y;", "277.750000", object.FLOAT_OBJ},
		{"y * y;", "3085.802500", object.FLOAT_OBJ},
		{"y * x;", "277.750000", object.FLOAT_OBJ},
		{"x / x;", "1", object.INT_OBJ},
		{"x / y;", "0.090009", object.FLOAT_OBJ},
		{"y / y;", "1.000000", object.FLOAT_OBJ},
		{"y / x;", "11.110000", object.FLOAT_OBJ},
		{"x == x;", "true", object.BOOL_OBJ},
		{"x == y;", "false", object.BOOL_OBJ},
		{"y == y;", "true", object.BOOL_OBJ},
		{"y == x;", "false", object.BOOL_OBJ},
		{"x != x;", "false", object.BOOL_OBJ},
		{"x != y;", "true", object.BOOL_OBJ},
		{"y != y;", "false", object.BOOL_OBJ},
		{"y != x;", "true", object.BOOL_OBJ},
		{"x > x;", "false", object.BOOL_OBJ},
		{"x > y;", "false", object.BOOL_OBJ},
		{"y > y;", "false", object.BOOL_OBJ},
		{"y > x;", "true", object.BOOL_OBJ},
		{"x >= x;", "true", object.BOOL_OBJ},
		{"x >= y;", "false", object.BOOL_OBJ},
		{"y >= y;", "true", object.BOOL_OBJ},
		{"y >= x;", "true", object.BOOL_OBJ},
		{"x < x;", "false", object.BOOL_OBJ},
		{"x < y;", "true", object.BOOL_OBJ},
		{"y < y;", "false", object.BOOL_OBJ},
		{"y < x;", "false", object.BOOL_OBJ},
		{"x <= x;", "true", object.BOOL_OBJ},
		{"x <= y;", "true", object.BOOL_OBJ},
		{"y <= y;", "true", object.BOOL_OBJ},
		{"y <= x;", "false", object.BOOL_OBJ},

		{"char w = 'c';", "c", object.CHAR_OBJ},
		{"string z = \"A string\";", "A string", object.STR_OBJ},
		{"w;", "c", object.CHAR_OBJ},
		{"z;", "A string", object.STR_OBJ},
		{"w + z;", "cA string", object.STR_OBJ},
		{"z + w;", "A stringc", object.STR_OBJ},
		{"z == z;", "true", object.BOOL_OBJ},
		{"z == w;", "false", object.BOOL_OBJ},
		{"w == w;", "true", object.BOOL_OBJ},
		{"w == z;", "false", object.BOOL_OBJ},
		{"z != z;", "false", object.BOOL_OBJ},
		{"z != w;", "true", object.BOOL_OBJ},
		{"w != w;", "false", object.BOOL_OBJ},
		{"w != z;", "true", object.BOOL_OBJ},
		{"w > w;", "false", object.BOOL_OBJ},
		{"w > z;", "true", object.BOOL_OBJ},
		{"z > z;", "false", object.BOOL_OBJ},
		{"z > w;", "false", object.BOOL_OBJ},
		{"w >= w;", "true", object.BOOL_OBJ},
		{"w >= z;", "true", object.BOOL_OBJ},
		{"z >= z;", "true", object.BOOL_OBJ},
		{"z >= w;", "false", object.BOOL_OBJ},
		{"w < w;", "false", object.BOOL_OBJ},
		{"w < z;", "false", object.BOOL_OBJ},
		{"z < z;", "false", object.BOOL_OBJ},
		{"z < w;", "true", object.BOOL_OBJ},
		{"w <= w;", "true", object.BOOL_OBJ},
		{"w <= z;", "false", object.BOOL_OBJ},
		{"z <= z;", "true", object.BOOL_OBJ},
		{"z <= w;", "true", object.BOOL_OBJ},

		{"a = x + y;", "ERROR: \"a\" not declared", object.ERROR_OBJ},

		{"int a = x + x;", "10", object.INT_OBJ},
		{"int a = x + y;", "ERROR: cannot assign FLOAT to INTEGER", object.ERROR_OBJ},

		{"float a = x + x;", "ERROR: cannot assign INTEGER to FLOAT", object.ERROR_OBJ},
		{"float a = x + y;", "60.550000", object.FLOAT_OBJ},
		{"a = x + y;", "60.550000", object.FLOAT_OBJ},

		{"char b = w + z;", "ERROR: cannot assign STRING to CHAR", object.ERROR_OBJ},
		{"string b = w + w;", "cc", object.STR_OBJ},
		{"string b = w + z;", "cA string", object.STR_OBJ},
		{"string b = z + w;", "A stringc", object.STR_OBJ},
		{"b = w + z;", "cA string", object.STR_OBJ},
		{"b = z + w;", "A stringc", object.STR_OBJ},

		{"int add(int a, int b) { return a + b; }", "int add(int a, int b) { return (a + b); }", object.FN_OBJ},
		{"add(1, 2);", "3", object.INT_OBJ},
		{"add(x, y);", "ERROR: wrong type for argument 2, got=FLOAT; expected:INTEGER", object.ERROR_OBJ},

		{"if (x > y) { return x; } else { return y; }", "55.550000", object.FLOAT_OBJ},
		{"if (x < y) { return x; } else { return y; }", "5", object.INT_OBJ},
	}

	e := New()
	for _, tt := range tests {
		l := lexer.New(tt.Line)
		p := parser.New(l)
		program := p.ParseProgram()
		result := e.Eval(program)

		if result == nil {
			t.Fatalf("got nil result for %q", tt.Line)
		}

		if result.Inspect() != tt.Result {
			t.Fatalf("expected %q as result, got %q for %q", tt.Result, result.Inspect(), tt.Line)
		}

		if result.Type() != tt.ResultType {
			t.Fatalf("expected %s as result type, got %s for %q", tt.ResultType, result.Type(), tt.Line)
		}
	}
}
