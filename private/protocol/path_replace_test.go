package protocol

import "testing"

func TestPathReplace(t *testing.T) {
	cases := []struct {
		Orig, ExpPath, ExpRawPath, Key, Val string
	}{
		{
			Orig:       "/{bucket}/{key+}",
			ExpPath:    "/123/{key+}",
			ExpRawPath: "/123/{key+}",
			Key:        "bucket", Val: "123",
		},
		{
			Orig:       "/{bucket}/{key+}",
			ExpPath:    "/{bucket}/abc",
			ExpRawPath: "/{bucket}/abc",
			Key:        "key", Val: "abc",
		},
		{
			Orig:       "/{bucket}/{key+}",
			ExpPath:    "/{bucket}/a/b/c",
			ExpRawPath: "/{bucket}/a/b/c",
			Key:        "key", Val: "a/b/c",
		},
		{
			Orig:       "/{bucket}/{key+}",
			ExpPath:    "/1/2/3/{key+}",
			ExpRawPath: "/1%2F2%2F3/{key+}",
			Key:        "bucket", Val: "1/2/3",
		},
		{
			Orig:       "/{bucket}/{key+}",
			ExpPath:    "/reallylongvaluegoesheregrowingarray/{key+}",
			ExpRawPath: "/reallylongvaluegoesheregrowingarray/{key+}",
			Key:        "bucket", Val: "reallylongvaluegoesheregrowingarray",
		},
	}

	for i, c := range cases {
		r := NewPathReplace(c.Orig)

		r.ReplaceElement(c.Key, c.Val)
		path, rawPath := r.Encode()
		if e, a := c.ExpPath, path; a != e {
			t.Errorf("%d, expect uri path to be %q got %q", i, e, a)
		}
		if e, a := c.ExpRawPath, rawPath; a != e {
			t.Errorf("%d, expect uri raw path to be %q got %q", i, e, a)
		}
	}
}

func TestPathReplace_Multiple(t *testing.T) {
	const (
		origURI    = "/{bucket}/{key+}"
		expPath    = "/something/123/other/thing"
		expRawPath = "/something%2F123/other/thing"
	)

	r := NewPathReplace(origURI)

	r.ReplaceElement("bucket", "something/123")
	r.ReplaceElement("key", "other/thing")

	path, rawPath := r.Encode()
	if e, a := expPath, path; a != e {
		t.Errorf("expect uri path to be %q got %q", e, a)
	}
	if e, a := expRawPath, rawPath; a != e {
		t.Errorf("expect uri path to be %q got %q", e, a)
	}
}
