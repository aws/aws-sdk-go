package expression

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// ConditionMode will specify the types of the struct ConditionBuilder
type ConditionMode int

const (
	// EqualCond will represent the Equal Clause ConditionBuilder
	EqualCond ConditionMode = iota + 1
)

// ConditionBuilder will represent the ConditionExpressions
type ConditionBuilder struct {
	operandList   []OperandExpression
	conditionList []ConditionBuilder
	err           error
	Mode          ConditionMode
}

// Equal

// Equal will create a ConditionBuilder. This will be the function call
func Equal(left, right OperandBuilder) ConditionBuilder {
	leftOpe, opeErr := left.BuildOperand()
	if opeErr != nil {
		return ConditionBuilder{
			err: opeErr,
		}
	}

	rightOpe, opeErr := right.BuildOperand()
	if opeErr != nil {
		return ConditionBuilder{
			err: opeErr,
		}
	}

	return ConditionBuilder{
		operandList: []OperandExpression{leftOpe, rightOpe},
		Mode:        EqualCond,
	}
}

// Equal will create a ConditionBuilder. This will be the method for PathBuilder
func (p PathBuilder) Equal(right OperandBuilder) ConditionBuilder {
	return Equal(p, right)
}

// Equal will create a ConditionBuilder. This will be the method for
// ValueBuilder
func (v ValueBuilder) Equal(right OperandBuilder) ConditionBuilder {
	return Equal(v, right)
}

// Equal will create a ConditionBuilder. This will be the method for SizeBuilder
func (s SizeBuilder) Equal(right OperandBuilder) ConditionBuilder {
	return Equal(s, right)
}

// buildExpression will iterate over the tree of ConditionBuilders and
// OperandExpressions and builds the Expression
func (cond ConditionBuilder) buildExpression(al aliasList) (Expression, error) {
	switch cond.Mode {
	case EqualCond:
		return equalBuildExpression(cond, al)
	}
	return Expression{}, fmt.Errorf("No matching Mode to %v", cond.Mode)
}

// buildList will iterate over the tree of ConditionBuilders and
// OperandExpressions and create an aliasList
func (cond ConditionBuilder) buildList() (aliasList, error) {
	al := aliasList{
		NamesList:  make([]string, 0),
		ValuesList: make([]dynamodb.AttributeValue, 0),
	}

	for _, opeList := range cond.operandList {
		tempAl, err := opeList.buildList()
		if err != nil {
			return aliasList{}, err
		}

		unique := true
		for _, newName := range tempAl.NamesList {
			for _, oldNames := range al.NamesList {
				if newName == oldNames {
					unique = false
					break
				}
			}
			if unique == false {
				unique = true
				continue
			}
			al.NamesList = append(al.NamesList, newName)
		}
	}

	for _, condList := range cond.conditionList {
		tempAl, err := condList.buildList()
		if err != nil {
			return aliasList{}, err
		}

		unique := true
		for _, newName := range tempAl.NamesList {
			for _, oldNames := range al.NamesList {
				if newName == oldNames {
					unique = false
					break
				}
			}
			if unique == false {
				unique = true
				continue
			}
			al.NamesList = append(al.NamesList, newName)
		}

		// for key, _ := range tempAl.NamesMap {
		// 	if al.NamesMap[key] == nil {
		// 		al.NamesMap[key] = aws.Int(len(al.NamesMap))
		// 	}
		// }
	}
	return al, nil
}

// BuildCondition will create the Expression represented by CompareBuilder
func (cond ConditionBuilder) BuildCondition() (Expression, error) {

	al, err := cond.buildList()
	if err != nil {
		return Expression{}, err
	}

	ret, err := cond.buildExpression(al)
	if err != nil {
		return Expression{}, err
	}

	return ret, nil
}

func equalBuildExpression(c ConditionBuilder, al aliasList) (Expression, error) {
	if len(c.conditionList) != 0 {
		return Expression{}, fmt.Errorf("Invalid ConditionBuilder. Expected 0 ConditionBuilders")
	}

	if len(c.operandList) != 2 {
		return Expression{}, fmt.Errorf("Invalid ConditionBuilder. Expected 2 Operands")
	}

	operandList := make([]Expression, 0)
	for _, ope := range c.operandList {
		expr, err := ope.buildExpression(al)
		if err != nil {
			return Expression{}, err
		}
		operandList = append(operandList, expr)
	}

	ret := mergeExpressionMaps(operandList)

	ret.Expression = "(" + operandList[0].Expression + ") = (" + operandList[1].Expression + ")"

	return ret, nil
}

func mergeExpressionMaps(lists ...[]Expression) Expression {
	ret := Expression{}
	for _, list := range lists {
		for _, expr := range list {
			for alias, name := range expr.Names {
				if ret.Names == nil {
					ret.Names = make(map[string]*string)
				}
				ret.Names[alias] = name
			}

			for alias, value := range expr.Values {
				if ret.Values == nil {
					ret.Values = make(map[string]*dynamodb.AttributeValue)
				}
				ret.Values[alias] = value
			}
		}
	}
	return ret
}
