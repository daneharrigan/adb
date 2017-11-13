package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestExporters(t *testing.T) {
	tests := []struct {
		Name     string
		Exporter Exporter
		Fixture  string
		Want     string
		Got      string
	}{
		{
			Name:     "export ToCSV",
			Exporter: ToCSV,
			Fixture:  "testing/fixture.json",
			Want:     "testing/export.csv",
			Got:      "export.csv",
		},
		{
			Name:     "export ToJSON",
			Exporter: ToJSON,
			Fixture:  "testing/fixture.json",
			Want:     "testing/export.json",
			Got:      "export.json",
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			f, err := os.Open(tc.Fixture)
			if err != nil {
				t.Fatal(err)
			}

			defer f.Close()

			var r Release
			if err := json.NewDecoder(f).Decode(&r); err != nil {
				t.Fatal(err)
			}

			dir, err := ioutil.TempDir("", "")
			if err != nil {
				t.Fatal(err)
			}

			defer os.RemoveAll(dir)

			if err := tc.Exporter(dir, r); err != nil {
				t.Fatal(err)
			}

			got, err := ioutil.ReadFile(filepath.Join(dir, tc.Got))
			if err != nil {
				t.Fatal(err)
			}

			want, err := ioutil.ReadFile(tc.Want)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(got, want) {
				t.Fatalf("\ngot: %s\nwant: %s", got, want)
			}
		})
	}
}
