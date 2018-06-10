package rdsutils_test

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
)

// Example contains usage of assuming a role and using
// that to build the auth token.
// Usage:
//	./main -user "iamuser" -dbname "foo" -region "us-west-2" -rolearn "arn" -endpoint "dbendpoint" -port 3306
func ExampleRDSUtils_ConnectViaAssumeRole() {
	userPtr := flag.String("user", "", "user of the credentials")
	regionPtr := flag.String("region", "us-east-1", "region to be used when grabbing sts creds")
	roleArnPtr := flag.String("rolearn", "", "role arn to be used when grabbing sts creds")
	endpointPtr := flag.String("endpoint", "", "DB endpoint to be connected to")
	portPtr := flag.Int("port", 3306, "DB port to be connected to")
	tablePtr := flag.String("table", "test_table", "DB table to query against")
	dbNamePtr := flag.String("dbname", "", "DB name to query against")
	flag.Parse()

	// Check required flags. Will exit with status code 1 if
	// required field isn't set.
	if err := requiredFlags(
		userPtr,
		regionPtr,
		roleArnPtr,
		endpointPtr,
		portPtr,
		dbNamePtr,
	); err != nil {
		fmt.Printf("Error: %v\n\n", err)
		flag.PrintDefaults()
		os.Exit(1)
	}

	sess := session.Must(session.NewSession())
	creds := stscreds.NewCredentials(sess, *roleArnPtr)

	endpoint := fmt.Sprintf("%s:%d", *endpointPtr, *portPtr)
	token, err := rdsutils.BuildAuthToken(endpoint, *regionPtr, *userPtr, creds)
	if err != nil {
		panic(fmt.Errorf("failed to build authentication token: %v", err))
	}

	// builds the connection endpoint for the SQL driver to use.
	dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true",
		*userPtr, token, *endpointPtr, *dbNamePtr,
	)

	const dbType = "mysql"

	db, err := sql.Open(dbType, dnsStr)
	// if an error is encountered here, then most likely security groups are incorrect
	// in the database.
	if err != nil {
		panic(fmt.Errorf("failed to open connection to the database"))
	}

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s  LIMIT 1", *tablePtr))
	if err != nil {
		panic(fmt.Errorf("failed to select from table, %q, with %v", *tablePtr, err))
	}

	for rows.Next() {
		columns, err := rows.Columns()
		if err != nil {
			panic(fmt.Errorf("failed to read columns from row: %v", err))
		}

		fmt.Printf("rows colums:\n%d\n", len(columns))
	}
}

func requiredFlags(flags ...interface{}) error {
	for _, f := range flags {
		switch f.(type) {
		case nil:
			return fmt.Errorf("one or more required flags were not set")
		}
	}
	return nil
}
