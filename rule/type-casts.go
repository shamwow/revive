package rule

import (
	"fmt"
	"go/ast"
	"regexp"

	"github.com/shamwow/revive/lint"
	"github.com/pkg/errors"
)

// TypeCastsRule lints given else constructs.
type TypeCastsRule struct{}

// Apply applies the rule to given file.
func (r *TypeCastsRule) Apply(
	file *lint.File, arguments lint.Arguments) []lint.Failure {

	excludedFiles := make([]string, 0)
	if len(arguments) > 0 {
		excludedFilesRaw, ok := arguments[0].([]interface{})
		if !ok {
			panic(errors.Errorf("invalid value passed as argument excluded_files " +
				"to the 'type-casts' rule"))
		}
		for _, excludedFileRaw := range excludedFilesRaw {
			excludedFiles = append(excludedFiles,
				fmt.Sprintf("%v", excludedFileRaw))
		}
	}

	for _, excludedFile := range excludedFiles {
		match, err := regexp.Match(excludedFile, []byte(file.Name))
		if err != nil {
			panic(errors.Wrap(err, "could not regex match exluded files"))
		}
		if match {
			return nil
		}
	}

	var failures []lint.Failure

	walker := lintTypeCasts{
		onFailure: func(failure lint.Failure) {
			failures = append(failures, failure)
		},
	}

	ast.Walk(walker, file.AST)

	return failures
}

// Name returns the rule name.
func (r *TypeCastsRule) Name() string {
	return "type-casts"
}

type lintTypeCasts struct {
	// If 0, we are not in a type switch statement. If N > 0, we are N nodes
	// deep into a AST whose root is a type switch statement node. Only keeps
	// track of the inner most type switch statement.
	//
	// Used to not raise the lint error for unchecked type switch assignments.
	typeSwitchStmtDepth int
	onFailure           func(lint.Failure)
}

func (w lintTypeCasts) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	if w.typeSwitchStmtDepth > 0 {
		w.typeSwitchStmtDepth++
	}

	switch node := n.(type) {
	case *ast.TypeSwitchStmt:
		w.typeSwitchStmtDepth = 1
		return w
	case *ast.AssignStmt:
		if w.typeSwitchStmtDepth != 2 && isAssignTypeAssert(node) &&
			len(node.Lhs) != 2 {

			w.onFailure(lint.Failure{
				Confidence: 1,
				Failure:    "type assert can panic",
				Node:       node,
			})
			return w
		}
	}

	return w
}

func isAssignTypeAssert(node *ast.AssignStmt) bool {
	for _, rhsNode := range node.Rhs {
		_, isTypeAssert := rhsNode.(*ast.TypeAssertExpr)
		if isTypeAssert {
			return true
		}
	}
	return false
}
