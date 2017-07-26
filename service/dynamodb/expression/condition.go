package expression

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// ConditionMode will specify the types of the struct Condition
type ConditionMode int

const (
	// EqualCond will represent the Equal Clause Condition
	EqualCond ConditionMode = iota + 1
)

// Condition will represent the ConditionExpressions
type Condition struct {
	OperandList   []OperandBuilder
	ConditionList []Condition
	Mode          ConditionMode
}

// Equal

// Equal will create a CompareBuilder. This will be the method PathBuilder.
func (p PathBuilder) Equal(right OperandBuilder) Condition {
	return Condition{
		OperandList: []OperandBuilder{p, right},
		Mode:        EqualCond,
	}
}

// Equal will create a CompareBuilder. This will be the method ValueBuilder.
func (v ValueBuilder) Equal(right OperandBuilder) Condition {
	return Condition{
		OperandList: []OperandBuilder{v, right},
		Mode:        EqualCond,
	}
}

// Equal will create a CompareBuilder. This will be the method SizeBuilder.
func (s SizeBuilder) Equal(right OperandBuilder) Condition {
	return Condition{
		OperandList: []OperandBuilder{s, right},
		Mode:        EqualCond,
	}
}

// Size

// SizeBuilder will implement OperandBuilder thus being an Operand. This
// reflects the fact that the function Size() returns type that is used in place
// of an Operand
type SizeBuilder struct {
	path PathBuilder
}

// Size will
func (p PathBuilder) Size() SizeBuilder {
	return SizeBuilder{
		path: p,
	}
}

// BuildOperand will allow SizeBuilder to implement the interface OperandBuilder
func (s SizeBuilder) BuildOperand(al AliasList) (Expression, error) {
	expr, err := s.path.BuildOperand(al)
	if err != nil {
		return Expression{}, err
	}
	expr.Expression = "size (" + expr.Expression + ")"
	return expr, nil
}

// ListOperand will allow SizeBuilder to implement the interface OperandBuilder
func (s SizeBuilder) ListOperand() (AliasList, error) {
	return s.path.ListOperand()
}

func (cond Condition) buildCondition(al AliasList) (Expression, error) {
	switch cond.Mode {
	case EqualCond:
		return equalBuildCondition(cond, al)
	}
	return Expression{}, fmt.Errorf("No matching Mode to %v", cond.Mode)
}

func (cond Condition) buildList() (AliasList, error) {
	al := AliasList{
		NamesList: make([]string, 0),
	}

	for _, opeList := range cond.OperandList {
		tempAl, err := opeList.ListOperand()
		if err != nil {
			return AliasList{}, err
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

	for _, condList := range cond.ConditionList {
		tempAl, err := condList.buildList()
		if err != nil {
			return AliasList{}, err
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
func (cond Condition) BuildCondition() (Expression, error) {

	al, err := cond.buildList()
	if err != nil {
		return Expression{}, err
	}
	al.ValuesCounter = aws.Int(0)

	ret, err := cond.buildCondition(al)
	if err != nil {
		return Expression{}, err
	}

	return ret, nil
}

func equalBuildCondition(c Condition, al AliasList) (Expression, error) {
	if len(c.ConditionList) != 0 {
		return Expression{}, fmt.Errorf("Invalid Condition. Expected 0 Conditions")
	}

	if len(c.OperandList) != 2 {
		return Expression{}, fmt.Errorf("Invalid Condition. Expected 2 Operands")
	}

	left, err := c.OperandList[0].BuildOperand(al)
	if err != nil {
		return Expression{}, err
	}

	right, err := c.OperandList[1].BuildOperand(al)
	if err != nil {
		return Expression{}, err
	}

	if left.Names == nil && right.Names != nil {
		left.Names = make(map[string]*string)
	}
	for alias, name := range right.Names {
		left.Names[alias] = name
	}

	if left.Values == nil && right.Values != nil {
		left.Values = make(map[string]*dynamodb.AttributeValue)
	}
	for alias, value := range right.Values {
		left.Values[alias] = value
	}

	left.Expression = left.Expression + " = " + right.Expression

	return left, nil
}
