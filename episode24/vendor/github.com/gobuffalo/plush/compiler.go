package plush

import (
	"bytes"
	"fmt"
	"html/template"
	"reflect"
	"regexp"
	"time"

	"github.com/gobuffalo/plush/ast"

	"github.com/pkg/errors"
)

var ErrUnknownIdentifier = errors.New("unknown identifier")

type compiler struct {
	ctx     *Context
	program *ast.Program
	curStmt ast.Statement
}

func (c *compiler) compile() (string, error) {
	bb := &bytes.Buffer{}
	for _, stmt := range c.program.Statements {
		var res interface{}
		var err error
		switch node := stmt.(type) {
		case *ast.ReturnStatement:
			res, err = c.evalReturnStatement(node)
		case *ast.ExpressionStatement:
			if h, ok := node.Expression.(*ast.HTMLLiteral); ok {
				res = template.HTML(h.Value)
			} else {
				_, err = c.evalExpression(node.Expression)
			}
		case *ast.LetStatement:
			res, err = c.evalLetStatement(node)
		}
		if err != nil {
			s := stmt
			if c.curStmt != nil {
				s = c.curStmt
			}
			return "", errors.WithStack(errors.Wrapf(err, "line %d", s.T().LineNumber))
		}

		c.write(bb, res)
	}
	return bb.String(), nil
}

func (c *compiler) write(bb *bytes.Buffer, i interface{}) {
	switch t := i.(type) {
	case time.Time:
		if dtf, ok := c.ctx.Value("TIME_FORMAT").(string); ok {
			bb.WriteString(t.Format(dtf))
			return
		}
		bb.WriteString(t.Format(DefaultTimeFormat))
	case *time.Time:
		c.write(bb, *t)
	case interfaceable:
		c.write(bb, t.Interface())
	case string, ast.Printable, bool:
		bb.WriteString(template.HTMLEscaper(t))
	case template.HTML:
		bb.WriteString(string(t))
	case HTMLer:
		bb.WriteString(string(t.HTML()))
	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64:
		bb.WriteString(fmt.Sprint(t))
	case fmt.Stringer:
		bb.WriteString(t.String())
	case []string:
		for _, ii := range t {
			c.write(bb, ii)
		}
	case []interface{}:
		for _, ii := range t {
			c.write(bb, ii)
		}
	}
}

func (c *compiler) evalExpression(node ast.Expression) (interface{}, error) {
	switch s := node.(type) {
	case *ast.HTMLLiteral:
		return template.HTML(s.Value), nil
	case *ast.StringLiteral:
		return s.Value, nil
	case *ast.IntegerLiteral:
		return s.Value, nil
	case *ast.FloatLiteral:
		return s.Value, nil
	case *ast.InfixExpression:
		return c.evalInfixExpression(s)
	case *ast.HashLiteral:
		return c.evalHashLiteral(s)
	case *ast.IndexExpression:
		return c.evalIndexExpression(s)
	case *ast.CallExpression:
		return c.evalCallExpression(s)
	case *ast.Identifier:
		return c.evalIdentifier(s)
	case *ast.Boolean:
		return s.Value, nil
	case *ast.ArrayLiteral:
		return c.evalArrayLiteral(s)
	case *ast.ForExpression:
		return c.evalForExpression(s)
	case *ast.IfExpression:
		return c.evalIfExpression(s)
	case *ast.PrefixExpression:
		return c.evalPrefixExpression(s)
	case *ast.FunctionLiteral:
		return c.evalFunctionLiteral(s)
	case *ast.AssignExpression:
		return c.evalAssignExpression(s)
	case nil:
		return nil, nil
	}
	return nil, errors.WithStack(errors.Errorf("could not evaluate node %T", node))
}

func (c *compiler) evalAssignExpression(node *ast.AssignExpression) (interface{}, error) {
	v, err := c.evalExpression(node.Value)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	n := node.Name.Value
	if !c.ctx.Has(n) {
		return nil, errors.Errorf("could not find identifier named %s", n)
	}
	c.ctx.Set(n, v)
	return nil, nil
}

func (c *compiler) evalUserFunction(node *userFunction, args []ast.Expression) (interface{}, error) {
	octx := c.ctx
	defer func() { c.ctx = octx }()
	c.ctx = c.ctx.New()
	for i, p := range node.Parameters {
		a := args[i]
		v, err := c.evalExpression(a)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		c.ctx.Set(p.Value, v)

	}
	return c.evalBlockStatement(node.Block)
}

func (c *compiler) evalFunctionLiteral(node *ast.FunctionLiteral) (interface{}, error) {
	params := node.Parameters
	block := node.Block
	return &userFunction{Parameters: params, Block: block}, nil
}

func (c *compiler) evalPrefixExpression(node *ast.PrefixExpression) (interface{}, error) {
	res, err := c.evalExpression(node.Right)
	if err != nil {
		if errors.Cause(err) != ErrUnknownIdentifier {
			return nil, errors.WithStack(err)
		}
	}
	switch node.Operator {
	case "!":
		return !c.isTruthy(res), nil
	}
	return nil, errors.WithStack(errors.Errorf("unknown operator %s", node.Operator))
}

func (c *compiler) evalIfExpression(node *ast.IfExpression) (interface{}, error) {
	// fmt.Println("evalIfExpression")
	con, err := c.evalExpression(node.Condition)
	if err != nil {
		if errors.Cause(err) != ErrUnknownIdentifier {
			return nil, errors.WithStack(err)
		}
	}

	var r interface{}
	if c.isTruthy(con) {
		return c.evalBlockStatement(node.Block)
	} else {
		if node.ElseBlock != nil {
			return c.evalBlockStatement(node.ElseBlock)
		}
	}

	return r, nil
}

func (c *compiler) isTruthy(i interface{}) bool {
	if i == nil {
		return false
	}
	switch t := i.(type) {
	case bool:
		return t
	case string:
		return t != ""
	case template.HTML:
		return t != ""
	}

	return true
}

func (c *compiler) evalIndexExpression(node *ast.IndexExpression) (interface{}, error) {
	index, err := c.evalExpression(node.Index)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	left, err := c.evalExpression(node.Left)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	rv := reflect.ValueOf(left)
	switch rv.Kind() {
	case reflect.Map:
		val := rv.MapIndex(reflect.ValueOf(index))
		if !val.IsValid() {
			return nil, nil
		}
		return val.Interface(), nil
	case reflect.Array, reflect.Slice:
		if i, ok := index.(int); ok {
			return rv.Index(i).Interface(), nil
		}
	}
	return nil, errors.WithStack(errors.Errorf("could not index %T with %T", left, index))
}

func (c *compiler) evalHashLiteral(node *ast.HashLiteral) (interface{}, error) {
	m := map[string]interface{}{}
	for ke, ve := range node.Pairs {
		v, err := c.evalExpression(ve)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		m[ke.String()] = v
	}
	return m, nil
}

func (c *compiler) evalLetStatement(node *ast.LetStatement) (interface{}, error) {
	// fmt.Println("evalLetStatement")
	v, err := c.evalExpression(node.Value)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	c.ctx.Set(node.Name.Value, v)
	return nil, nil
}

func (c *compiler) evalIdentifier(node *ast.Identifier) (interface{}, error) {
	if node.Callee != nil {
		c, err := c.evalExpression(node.Callee)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		rv := reflect.ValueOf(c)
		if !rv.IsValid() {
			return nil, nil
		}
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}
		if rv.Kind() != reflect.Struct {
			return nil, errors.WithStack(errors.Errorf("'%s' does not have a field or method named '%s' (%s)", node.Callee.String(), node.Value, node))
		}
		f := rv.FieldByName(node.Value)
		if !f.IsValid() {
			m := rv.MethodByName(node.Value)
			if !m.IsValid() {
				return nil, errors.WithStack(errors.Errorf("'%s' does not have a field or method named '%s' (%s)", node.Callee.String(), node.Value, node))
			}
			return m.Interface(), nil
		}
		return f.Interface(), nil
	}
	if c.ctx.Has(node.Value) {
		return c.ctx.Value(node.Value), nil
	}
	if node.Value == "nil" {
		return nil, nil
	}
	return nil, errors.Wrap(ErrUnknownIdentifier, node.Value)
}

func (c *compiler) evalInfixExpression(node *ast.InfixExpression) (interface{}, error) {
	// fmt.Println("evalInfixExpression")
	lres, err := c.evalExpression(node.Left)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if node.Operator == "&&" {
		if !c.isTruthy(lres) {
			return false, nil
		}
	}
	rres, err := c.evalExpression(node.Right)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	switch node.Operator {
	case "&&", "||":
		return c.boolsOperator(lres, rres, node.Operator)
	}

	switch t := lres.(type) {
	case string:
		return c.stringsOperator(t, rres, node.Operator)
	case int64:
		if r, ok := rres.(int64); ok {
			return c.intsOperator(int(t), int(r), node.Operator)
		}
	case int:
		if r, ok := rres.(int); ok {
			return c.intsOperator(t, r, node.Operator)
		}
	case float64:
		if r, ok := rres.(float64); ok {
			return c.floatsOperator(t, r, node.Operator)
		}
	case bool:
		return c.boolsOperator(lres, rres, node.Operator)
	case nil:
		return nil, nil
	}
	return nil, errors.WithStack(errors.Errorf("unable to operate (%s) on %T and %T ", node.Operator, lres, rres))
}

func (c *compiler) boolsOperator(l interface{}, r interface{}, op string) (interface{}, error) {
	lt := c.isTruthy(l)
	rt := c.isTruthy(r)
	if op == "||" {
		return lt || rt, nil
	}
	return lt && rt, nil
}

func (c *compiler) intsOperator(l int, r int, op string) (interface{}, error) {
	switch op {
	case "+":
		return l + r, nil
	case "-":
		return l - r, nil
	case "/":
		return l / r, nil
	case "*":
		return l * r, nil
	case "<":
		return l < r, nil
	case ">":
		return l > r, nil
	case "!=":
		return l != r, nil
	case ">=":
		return l >= r, nil
	case "<=":
		return l <= r, nil
	case "==":
		return l == r, nil
	}
	return nil, errors.WithStack(errors.Errorf("unknown operator for integer %s", op))
}

func (c *compiler) floatsOperator(l float64, r float64, op string) (interface{}, error) {
	switch op {
	case "+":
		return l + r, nil
	case "-":
		return l - r, nil
	case "/":
		return l / r, nil
	case "*":
		return l * r, nil
	case "<":
		return l < r, nil
	case ">":
		return l > r, nil
	case "!=":
		return l != r, nil
	case ">=":
		return l >= r, nil
	case "<=":
		return l <= r, nil
	case "==":
		return l == r, nil
	}
	return nil, errors.WithStack(errors.Errorf("unknown operator for float %s", op))
}

func (c *compiler) stringsOperator(l string, r interface{}, op string) (interface{}, error) {
	rr := fmt.Sprint(r)
	switch op {
	case "+":
		return l + rr, nil
	// case "-":
	// 	return l - rr, nil
	// case "/":
	// 	return l / rr, nil
	// case "*":
	// 	return l * rr, nil
	case "<":
		return l < rr, nil
	case ">":
		return l > rr, nil
	case "!=":
		return l != rr, nil
	case ">=":
		return l >= rr, nil
	case "<=":
		return l <= rr, nil
	case "==":
		return l == rr, nil
	case "~=":
		x, err := regexp.Compile(rr)
		if err != nil {
			return nil, errors.WithStack(errors.Errorf("couldn't compile regex %s", rr))
		}
		return x.MatchString(l), nil
	}
	return nil, errors.WithStack(errors.Errorf("unknown operator for string %s", op))
}

func (c *compiler) evalCallExpression(node *ast.CallExpression) (interface{}, error) {
	// fmt.Println("evalCallExpression")
	var rv reflect.Value
	if node.Callee != nil {
		c, err := c.evalExpression(node.Callee)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		rc := reflect.ValueOf(c)
		mname := node.Function.String()
		if i, ok := node.Function.(*ast.Identifier); ok {
			mname = i.Value
		}
		rv = rc.MethodByName(mname)
		if !rv.IsValid() {
			if rv.Kind() == reflect.Slice {
				rv = rc.FieldByName(mname)
				if rv.IsValid() {
					return rv.Interface(), nil
				}
			}
			return rc.Interface(), nil
		}
	} else {
		f, err := c.evalExpression(node.Function)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if ff, ok := f.(*userFunction); ok {
			return c.evalUserFunction(ff, node.Arguments)
		}
		rv = reflect.ValueOf(f)
	}
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if !rv.IsValid() {
		return nil, errors.WithStack(errors.Errorf("%+v (%T) is an invalid function", node.String(), rv))
	}

	rt := rv.Type()
	rtNumIn := rt.NumIn()
	isVariadic := rt.IsVariadic()
	args := []reflect.Value{}

	if !isVariadic {
		if len(node.Arguments) > rtNumIn {
			return nil, errors.WithStack(errors.Errorf("%s too many arguments (%d for %d)", node.String(), len(node.Arguments), rtNumIn))
		}

		for pos, a := range node.Arguments {
			v, err := c.evalExpression(a)
			if err != nil {
				return nil, errors.WithStack(err)
			}

			var ar reflect.Value
			expectedT := rt.In(pos)
			if v != nil {
				ar = reflect.ValueOf(v)
			} else {
				ar = reflect.New(expectedT).Elem()
			}

			actualT := ar.Type()
			if !actualT.AssignableTo(expectedT) {
				return nil, errors.WithStack(errors.Errorf("%+v (%T) is an invalid argument for %s at pos %d: expected (%s)", v, v, node.Function.String(), pos, expectedT))
			}

			args = append(args, ar)
		}

		hc := func(arg reflect.Type) {
			if arg.Name() == helperContextKind {
				hargs := HelperContext{
					Context:  c.ctx,
					compiler: c,
					block:    node.Block,
				}
				args = append(args, reflect.ValueOf(hargs))
				return
			}
			if arg.Name() == "Data" {
				args = append(args, reflect.ValueOf(c.ctx.export()))
				return
			}
			if arg.Kind() == reflect.Map {
				args = append(args, reflect.ValueOf(map[string]interface{}{}))
			}
		}

		if len(args) < rtNumIn {
			// missing some args, let's see if we can figure out what they are.
			diff := rtNumIn - len(args)
			switch diff {
			case 2:
				// check last is help
				// check if last -1 is map
				arg := rt.In(rtNumIn - 2)
				hc(arg)
				last := rt.In(rtNumIn - 1)
				hc(last)
			case 1:
				// check if help or map
				last := rt.In(rtNumIn - 1)
				hc(last)
			}
		}

		if len(args) > rtNumIn {
			return nil, errors.WithStack(errors.Errorf("%s too many arguments (%d for %d) - %+v", node.String(), len(args), rtNumIn, args))
		}
		if len(args) < rtNumIn {
			return nil, errors.WithStack(errors.Errorf("%s too few arguments (%d for %d) - %+v", node.String(), len(args), rtNumIn, args))
		}
	} else {
		// Variadic func
		nodeArgs := node.Arguments
		nodeArgsLen := len(nodeArgs)
		if nodeArgsLen < rtNumIn-1 {
			return nil, errors.WithStack(errors.Errorf("%s too few arguments (%d for %d) - %+v", node.String(), len(args), rtNumIn, args))
		}
		var pos int

		// Handle normal args
		for pos = 0; pos < rtNumIn-1; pos++ {
			v, err := c.evalExpression(nodeArgs[pos])
			if err != nil {
				return nil, errors.WithStack(err)
			}

			var ar reflect.Value
			expectedT := rt.In(pos)
			if v != nil {
				ar = reflect.ValueOf(v)
			} else {
				ar = reflect.New(expectedT).Elem()
			}

			actualT := ar.Type()
			if !actualT.AssignableTo(expectedT) {
				return nil, errors.WithStack(errors.Errorf("%+v (%T) is an invalid argument for %s at pos %d: expected (%s)", v, v, node.Function.String(), pos, expectedT))
			}

			args = append(args, ar)
		}

		// Unroll variadic arg
		expectedT := rt.In(pos).Elem()
		for ; pos < nodeArgsLen; pos++ {
			v, err := c.evalExpression(nodeArgs[pos])
			if err != nil {
				return nil, errors.WithStack(err)
			}

			var ar reflect.Value
			if v != nil {
				ar = reflect.ValueOf(v)
			} else {
				ar = reflect.New(expectedT)
			}

			actualT := ar.Type()
			if !actualT.AssignableTo(expectedT) {
				return nil, errors.WithStack(errors.Errorf("%+v (%T) is an invalid argument for %s at pos %d: expected (%s)", v, v, node.Function.String(), pos, expectedT))
			}

			args = append(args, ar)
		}
	}

	res := rv.Call(args)
	if len(res) > 0 {
		if e, ok := res[len(res)-1].Interface().(error); ok {
			return nil, errors.WithStack(e)
		}
		return res[0].Interface(), nil
	}
	return nil, nil
}

func (c *compiler) evalForExpression(node *ast.ForExpression) (interface{}, error) {
	octx := c.ctx
	defer func() {
		c.ctx = octx
	}()
	c.ctx = octx.New()
	iter, err := c.evalExpression(node.Iterable)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	riter := reflect.ValueOf(iter)
	if riter.Kind() == reflect.Ptr {
		riter = riter.Elem()
	}
	ret := []interface{}{}
	switch riter.Kind() {
	case reflect.Map:
		keys := riter.MapKeys()
		for i := 0; i < len(keys); i++ {
			k := keys[i]
			v := riter.MapIndex(k)
			c.ctx.Set(node.KeyName, k.Interface())
			c.ctx.Set(node.ValueName, v.Interface())
			res, err := c.evalBlockStatement(node.Block)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			ret = append(ret, res)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < riter.Len(); i++ {
			v := riter.Index(i)
			c.ctx.Set(node.KeyName, i)
			c.ctx.Set(node.ValueName, v.Interface())
			res, err := c.evalBlockStatement(node.Block)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			if res != nil {
				ret = append(ret, res)
			}
		}
	default:
		if iter == nil {
			return nil, nil
		}
		if it, ok := iter.(Iterator); ok {
			i := 0
			ii := it.Next()
			for ii != nil {
				c.ctx.Set(node.KeyName, i)
				c.ctx.Set(node.ValueName, ii)
				res, err := c.evalBlockStatement(node.Block)
				if err != nil {
					return nil, errors.WithStack(err)
				}
				if res != nil {
					ret = append(ret, res)
				}
				ii = it.Next()
				i++
			}
			return ret, nil
		}
		return ret, errors.WithStack(errors.Errorf("could not iterate over %T", iter))
	}
	return ret, nil
}

func (c *compiler) evalBlockStatement(node *ast.BlockStatement) (interface{}, error) {
	// fmt.Println("evalBlockStatement")
	res := []interface{}{}
	for _, s := range node.Statements {
		i, err := c.evalStatement(s)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if i != nil {
			res = append(res, i)
		}
	}
	return res, nil
}

func (c *compiler) evalStatement(node ast.Statement) (interface{}, error) {
	c.curStmt = node
	// fmt.Println("evalStatement")
	switch t := node.(type) {
	case *ast.ExpressionStatement:
		s, err := c.evalExpression(t.Expression)
		switch s.(type) {
		case ast.Printable, template.HTML:
			return s, errors.WithStack(err)
		}
		return nil, errors.WithStack(err)
	case *ast.ReturnStatement:
		return c.evalReturnStatement(t)
	case *ast.LetStatement:
		return c.evalLetStatement(t)
	}
	return nil, errors.WithStack(errors.Errorf("could not eval statement %T", node))
}

func (c *compiler) evalReturnStatement(node *ast.ReturnStatement) (interface{}, error) {
	// fmt.Println("evalReturnStatement")
	res, err := c.evalExpression(node.ReturnValue)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (c *compiler) evalArrayLiteral(node *ast.ArrayLiteral) (interface{}, error) {
	res := []interface{}{}
	for _, e := range node.Elements {
		i, err := c.evalExpression(e)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		res = append(res, i)
	}
	return res, nil
}
