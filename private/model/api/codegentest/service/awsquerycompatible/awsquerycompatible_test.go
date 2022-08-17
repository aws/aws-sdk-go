//go:build go1.8 && codegen
// +build go1.8,codegen

package awsquerycompatible

import ( "testing" )

func Test_MappedErrorCode(t *testing.T) {
	assertErrorCode("AWS.SimpleQueueService.QueueDeletedRecently", ErrCodeQueueDeletedRecently, t)
}

func Test_UnmappedErrorCode(t *testing.T) {
	assertErrorCode("QueueNameExists", ErrCodeQueueNameExists, t)
}

func assertErrorCode(expected string, actual string, t *testing.T) {
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
