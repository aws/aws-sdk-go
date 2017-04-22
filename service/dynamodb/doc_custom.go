// AttributeValue Marshaling and Unmarshaling Helpers
//
// The package dynamodbattribute nested under this package provides marshaling
// utilities for marshaling to AttributeValue types and unmarshaling to Go
// value types. These utilities allow you to marshal slices, maps, structs,
// and scalar values to and from dynamodb.AttributeValue. See the
// dynamodbattribute package for more information.
//
// AttributeValue Marshaling
//
// To marshal a Go type to a dynamodbAttributeValue you can use the Marshal
// functions in the dynamodbattribute package. There are specialized versions
// of these functions for collections of Attributevalue, such as maps and lists.
//
// The following example uses MarshalMap to convert the Record Go type to a
// dynamodb.AttributeValue type and use the value to make a PutItem API request.
//
//     type Record struct {
//         ID     string
//         URLs   []string
//     }
//
//     //...
//
//     r := Record{
//         ID:   "ABC123",
//         URLs: []string{
//             "https://example.com/first/link",
//             "https://example.com/second/url",
//         },
//     }
//     av, err := dynamodbattribute.MarshalMap(r)
//     if err != nil {
//         panic(fmt.Sprintf("failed to DynamoDB marshal Record, %v", err))
//     }
//
//     _, err = svc.PutItem(&dynamodb.PutItemInput{
//         TableName: aws.String(myTableName),
//         Item:      av,
//     })
//     if err != nil {
//         panic(fmt.Sprintf("failed to put Record to DynamoDB, %v", err))
//     }
//
// AttributeValue Unmarshaling
//
// To unmarshal a dynamodb.AttributeValue to a Go type you can use the Unmarshal
// functions in the dynamodbattribute package. There are specialized versions
// of these functions for collections of Attributevalue, such as maps and lists.
//
// The following example will unmarshal the DynamoDB's Scan API operation. The
// Items returned by the operation will be unmarshaled into the slice of Records
// Go type.
//
//     type Record struct {
//         ID     string
//         URLs   []string
//     }
//
//     //...
//
//     var records []Record
//
//     // Use the ScanPages method to perform the scan with pagination. Use
//     // just Scan method to make the API call without pagination.
//     err := svc.ScanPages(&dynamodb.ScanInput{
//         TableName: aws.String(myTableName),
//     }, func(page *dynamodb.ScanOutput, last bool) bool {
//         recs := []Record{}
//
//         err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &recs)
//         if err != nil {
//              panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
//         }
//
//         records = append(records, recs...)
//
//         return true // keep paging
//     })
package dynamodb
