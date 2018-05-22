// +build integration

package s3

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

func TestSelectObjectContent(t *testing.T) {
	keyName := "selectObject.csv"
	putTestFile(t, filepath.Join("testdata", "positive_select.csv"), keyName)

	resp, err := svc.SelectObjectContent(&s3.SelectObjectContentInput{
		Bucket:         bucketName,
		Key:            &keyName,
		Expression:     aws.String("Select * from S3Object"),
		ExpressionType: aws.String(s3.ExpressionTypeSql),
		InputSerialization: &s3.InputSerialization{
			CSV: &s3.CSVInput{
				FieldDelimiter: aws.String(","),
				FileHeaderInfo: aws.String(s3.FileHeaderInfoIgnore),
			},
		},
		OutputSerialization: &s3.OutputSerialization{
			CSV: &s3.CSVOutput{
				FieldDelimiter: aws.String(","),
			},
		},
	})
	if err != nil {
		t.Fatalf("expect no error, %v", err)
	}
	defer resp.EventStream.Close()

	var sum int64
	var processed int64
	for event := range resp.EventStream.Events {
		switch tv := event.(type) {
		case *s3.RecordsEvent:
			sum += int64(len(tv.Payload))
		case *s3.StatsEvent:
			processed = *tv.Details.BytesProcessed
		}
	}

	if sum == 0 {
		t.Errorf("expect selected content, got none")
	}

	if processed == 0 {
		t.Errorf("expect selected status bytes processed, got none")
	}

	if err := resp.EventStream.Err(); err != nil {
		t.Fatalf("exect no error, %v", err)
	}
}

func TestSelectObjectContent_Error(t *testing.T) {
	keyName := "negativeSelect.csv"

	buf := make([]byte, 0, 1024*1024*6)
	buf = append(buf, []byte("name,number\n")...)
	line := []byte("jj,0\n")
	for i := 0; i < (cap(buf)/len(line))-2; i++ {
		buf = append(buf, line...)
	}
	buf = append(buf, []byte("gg,NaN\n")...)

	putTestContent(t, bytes.NewReader(buf), keyName)

	resp, err := svc.SelectObjectContent(&s3.SelectObjectContentInput{
		Bucket:         bucketName,
		Key:            &keyName,
		Expression:     aws.String("SELECT name FROM S3Object WHERE cast(number as int) < 1"),
		ExpressionType: aws.String(s3.ExpressionTypeSql),
		InputSerialization: &s3.InputSerialization{
			CSV: &s3.CSVInput{
				FileHeaderInfo: aws.String(s3.FileHeaderInfoUse),
			},
		},
		OutputSerialization: &s3.OutputSerialization{
			CSV: &s3.CSVOutput{
				FieldDelimiter: aws.String(","),
			},
		},
	})
	if err != nil {
		t.Fatalf("expect no error, %v", err)
	}
	defer resp.EventStream.Close()

	var sum int64
	for event := range resp.EventStream.Events {
		switch tv := event.(type) {
		case *s3.RecordsEvent:
			sum += int64(len(tv.Payload))
		}
	}

	if sum == 0 {
		t.Errorf("expect selected content")
	}

	err = resp.EventStream.Err()
	if err == nil {
		t.Fatalf("exepct error")
	}

	aerr := err.(awserr.Error)
	if a := aerr.Code(); len(a) == 0 {
		t.Errorf("expect, error code")
	}
	if a := aerr.Message(); len(a) == 0 {
		t.Errorf("expect, error message")
	}
}
