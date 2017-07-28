package expression

// func TestListOperand(t *testing.T) {
// 	cases := []struct {
// 		input               OperandBuilder
// 		expected            aliasList
// 		incompletePathError bool
// 		emptyPathError      bool
// 	}{
// 		{
// 			input: NewPath("foo"),
// 			expected: aliasList{
// 				NamesList: []string{
// 					"foo",
// 				},
// 				ValuesCounter: nil,
// 			},
// 		},
// 		{
// 			input: NewValue(5),
// 			expected: aliasList{
// 				NamesList:     nil,
// 				ValuesCounter: nil,
// 			},
// 		},
// 		{
// 			input: NewPath("foo.bar[7].baz"),
// 			expected: aliasList{
// 				NamesList: []string{
// 					"foo",
// 					"bar",
// 					"baz",
// 				},
// 				ValuesCounter: nil,
// 			},
// 		},
// 		{
// 			input:          NewPath(""),
// 			expected:       aliasList{},
// 			emptyPathError: true,
// 		},
// 		{
// 			input:               NewPath("foo..bar"),
// 			expected:            aliasList{},
// 			incompletePathError: true,
// 		},
// 	}
//
// 	for testNumber, c := range cases {
// 		al, err := c.input.ListOperand()
//
// 		if c.emptyPathError {
// 			if err == nil {
// 				t.Errorf("TestListOperand Test Number %#v: Expected empty path error but got no error", testNumber)
// 			} else {
// 				continue
// 			}
// 		}
// 		if c.incompletePathError {
// 			if err == nil {
// 				t.Errorf("TestListOperand Test Number %#v: Expected incomplete path error but got no error", testNumber)
// 			} else {
// 				continue
// 			}
// 		}
//
// 		if err != nil {
// 			t.Errorf("TestListOperand Test Number %#v: Unexpected Error %#v", testNumber, err)
// 		}
//
// 		if reflect.DeepEqual(al, c.expected) != true {
// 			t.Errorf("TestListOperand Test Number %#v: Expected %#v, got %#v", testNumber, c.expected, al)
// 		}
// 	}
// }

// func TestBuildOperand(t *testing.T) {
// 	cases := []struct {
// 		input        OperandBuilder
// 		expected     Expression
// 		counterError bool
// 		alError      bool
// 	}{
// 		{
// 			input: NewPath("foo"),
// 			expected: Expression{
// 				Names: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				Expression: "#0",
// 			},
// 		},
// 		{
// 			input: NewValue(5),
// 			expected: Expression{
// 				Values: map[string]*dynamodb.AttributeValue{
// 					":0": &dynamodb.AttributeValue{
// 						N: aws.String("5"),
// 					},
// 				},
// 				Expression: ":0",
// 			},
// 		},
// 		{
// 			input: NewPath("foo.bar"),
// 			expected: Expression{
// 				Names: map[string]*string{
// 					"#0": aws.String("foo"),
// 					"#1": aws.String("bar"),
// 				},
// 				Expression: "#0.#1",
// 			},
// 		},
// 		{
// 			input: NewPath("foo.bar[0].baz"),
// 			expected: Expression{
// 				Names: map[string]*string{
// 					"#0": aws.String("foo"),
// 					"#1": aws.String("bar"),
// 					"#2": aws.String("baz"),
// 				},
// 				Expression: "#0.#1[0].#2",
// 			},
// 		},
// 		{
// 			input:        NewValue(5),
// 			expected:     Expression{},
// 			counterError: true,
// 		},
// 		{
// 			input:    NewPath("foo"),
// 			expected: Expression{},
// 			alError:  true,
// 		},
// 		{
// 			input:    NewPath("foo").Size(),
// 			expected: Expression{},
// 			alError:  true,
// 		},
// 	}
//
// 	for testNumber, c := range cases {
// 		al, err := c.input.ListOperand()
// 		if err != nil {
// 			t.Error(err)
// 		}
//
// 		if !c.counterError {
// 			al.ValuesCounter = aws.Int(0)
// 		}
// 		if c.alError {
// 			al.NamesList = al.NamesList[1:]
// 		}
//
// 		operand, err := c.input.BuildOperand(al)
// 		if c.counterError {
// 			if err == nil {
// 				t.Errorf("TestBuildOperand Test Number %#v: Expected counter error but got no error", testNumber)
// 			} else {
// 				continue
// 			}
// 		}
//
// 		if c.alError {
// 			if err == nil {
// 				t.Errorf("TestBuildOperand Test Number %#v: Expected List error but got no error", testNumber)
// 			} else {
// 				continue
// 			}
// 		}
//
// 		if err != nil {
// 			t.Errorf("TestBuildOperand Test Number %#v: Unexpected Error %#v", testNumber, err)
// 		}
//
// 		if operand.Expression != c.expected.Expression {
// 			t.Errorf("TestBuildOperand Test Number %#v: BuildOperand returned an unexpected Expression string %#v, expected %#v\n", testNumber, operand.Expression, c.expected.Expression)
// 		}
//
// 		if reflect.DeepEqual(c.expected.Names, operand.Names) != true {
// 			t.Errorf("TestBuildOperand Test Number %#v: BuildOperand returned an unexpected Name Map %#v, expected %#v\n", testNumber, operand.Names, c.expected.Names)
// 		}
//
// 		if reflect.DeepEqual(c.expected.Values, operand.Values) != true {
// 			t.Errorf("TestBuildOperand Test Number %#v: BuildOperand returned an unexpected Name Map %#v, expected %#v\n", testNumber, operand.Values, c.expected.Values)
// 		}
// 	}
// }
