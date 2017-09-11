# expression
--
    import "github.com/aws/aws-sdk-go/service/dynamodb/expression"

Package expression provides the types and functions to create DynamoDB
Expressions. The goal of the package is to provide an abstraction above DynamoDB
Expressions and allow users to functionally specify relationships between item
attribute names and item attribute values. The package writes the formatted
Expression strings with the right syntax under the hood and users are able to
retrieve the Expressions using getter methods.


#### About DynamoDB Expressions

DynamoDB Expressions are strings that can either modify or specify DynamoDB
operations, such as PutItem and Query. There are different types of DynamoDB
Expressions and each DynamoDB operations utilize a subset of Expression types.
Each type of Expressions support different operations and for further
information on the syntax of Expressions, see
http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.html


#### Using the Package

The goal of the package is to create an Expression struct that has getter
methods that return formatted DynamoDB Expression strings,
ExpressionAttributeNames, and ExpressionAttributeValues. The Expression struct
is essentially a collection of builder structs representing the different types
of DynamoDB Expressions, such as ConditionBuilder, ProjectionBuilder, and so on.

The Builder struct is used to create instances of Expression structs. Users can
use the With() methods to add builders representing DynamoDB Expressions to the
Builder struct. Calling Build() returns an Expression struct and an error.

The builder structs representing the DynamoDB Expressions are created using the
package functions that specify relationships between OperandBuilders.
OperandBuilders represent operands, referring to item attribute names and values
and other objects that make up DynamoDB Expressions.

The different types of DynamoDB Expressions are represented by corresponding
structs such as ConditionBuilder, ProjectionBuilder, and so on. To create the
builder structs, users must use the package functions that specify relationships
between OperandBuilders. OperandBuilders represent operands, referring to item
attribute names and values and other objects that make up DynamoDB Expressions.

Each OperandBuilder has a specific function to create the concrete struct, such
as Name(), Value(), and Key(). NameBuilder and ValueBuilder are separate to draw
a distinction between referencing item attributes and string values. Note that
some OperandBuilders are only to be used in certain builders. For example, the
KeyBuilder struct should only be used in functions used to create
KeyConditionBuilders.

The following example outlines a typical usage of the package.

    keyCond := expression.Name("Artist").Equal(expression.Value("No One You Know"))
    proj := expression.NamesList(expression.Name("SongTitle"))
    expr, err := expression.NewBuilder().WithKeyCondition(keyCond).WithProjection(proj).Build()

    if err != nil {
      fmt.Println(err)
    }

    input := &dynamodb.QueryInput{
      ExpressionAttributeValues: expr.Values(),
      KeyConditionExpression:    expr.KeyCondition(),
      ProjectionExpression:      expr.Projection(),
      TableName:                 aws.String("Music"),
    }


#### Idiosyncrasies of the Expression Package

All of the exported structs in the package are opaque to enforce the idea that
the structs should be initialized using the package functions instead of being
struct initialized. If users create empty structs and use them as arguments to
the package functions, an UnsetParameterError is returned. Similarly, if invalid
values are passed into the package functions, an InvalidParameterError is
returned.

Users are able to create an Expression struct containing many DynamoDB
Expressions due to the fact that some operation inputs, such as
dynamodb.QueryInput, can utilize multiple DynamoDB Expressions. The Expression
struct envelops multiple DynamoDB Expressions to create a consistent map for
ExpressionAttributeValues and ExpressionAttributeNames.

Users must always fill in the ExpressionAttributeNames and
ExpressionAttributeValues member of the input struct when using the Expression
struct since all item attribute names and values are aliased. That means that if
the ExpressionAttributeNames and ExpressionAttributeValues member is not
assigned with the corresponding Names() and Values() methods, the DynamoDB
operation will run into a logic error.

The test files are only built for go versions 1.7 and above since it uses
t.Run()
