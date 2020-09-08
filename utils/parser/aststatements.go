package parser

//
//import "fmt"
//
//type Statement interface {
//	Execute() error
//	ToString() string
//	getNodesList() []Statement
//}
//
//type BasicStatement struct {
//	node Node
//}
//
//func (ps *BasicStatement) Execute() error {
//	_, err := ps.node.Execute()
//	return err
//}
//
//func (ps *BasicStatement) ToString() string {
//	return ps.node.ToString()
//}
//
//func (ps *BasicStatement) getNodesList() []Statement {
//	return []Statement{&BasicStatement{ps.node}}
//}
//
////////////////////
//
///* Assignment Statement */
//
//type AssignmentStatement struct {
//	Variable   string
//	Expression Node
//}
//
//func (as *AssignmentStatement) New(variable string, node Node) *AssignmentStatement {
//	return &AssignmentStatement{Variable: variable, Expression: node}
//}
//
//func (as *AssignmentStatement) Execute() error {
//	result, ok := as.Expression.Execute()
//	if ok == nil {
//		VarTable[as.Variable] = result
//	}
//	return nil
//}
//
//func (as *AssignmentStatement) ToString() string {
//	return "ASSIGNMENT STATEMENT (Statement) Identifier: '" + as.Variable + "'"
//}
//
//func (as *AssignmentStatement) getNodesList() []Statement {
//	return []Statement{&BasicStatement{as.Expression}}
//}
//
////////////////////
//
///* Print Statement */
//
//type PrintStatement struct {
//	node Node
//}
//
//func (ps *PrintStatement) Execute() error {
//	result, _ := ps.node.Execute()
//	fmt.Println(result.AsString())
//	return nil
//}
//
//func (ps *PrintStatement) ToString() string {
//	return "PRINT OPERATOR (Keyword)"
//}
//
//func (ps *PrintStatement) getNodesList() []Statement {
//	return []Statement{&BasicStatement{ps.node}}
//}
//
//
////////////////////
//
///* If Statement */
//
//type IfStatement struct {
//	node Node
//	ifStmt Statement
//	elseStmt Statement
//}
//
//func (is *IfStatement) Execute() error {
//	result, _ := is.node.Execute()
//
//	var err error
//	if result.AsNumber() != 0 {
//		err = is.ifStmt.Execute()
//	} else if is.elseStmt != nil {
//		err = is.elseStmt.Execute()
//	}
//
//	return err
//}
//
//func (is *IfStatement) ToString() string {
//	result := "IF STATEMENT (Statement) "
//	result += is.node.ToString()
//	result += "\n    !--> "
//	result += is.ifStmt.ToString()
//	result += string('\n')
//	if is.elseStmt != nil {
//		result += "!--> ELSE STATEMENT (Statement) "
//		result += "\n    !--> "
//		result += is.elseStmt.ToString()
//	}
//	return result
//}
//
//func (is *IfStatement) getNodesList() []Statement {
//	return []Statement{}
//}
