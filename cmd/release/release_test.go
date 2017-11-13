package main

import (
	"reflect"
	"testing"
)

func TestRelease(t *testing.T) {
	want := Release{
		Groups: []Group{
			{
				Name:    "Example Group",
				Aliases: []string{"Example Group Alias"},
				Muscles: []Muscle{
					{
						Name: "Example Muscle",
						Targets: []Target{
							{
								Name: "Example Target",
								Aliases: []string{
									"Example Target Alias",
									"Example Target Alias 2",
								},
							},
						},
					},
					{
						Name: "Example 2 Muscle",
					},
				},
			},
		},
	}

	got, err := NewRelease("testing")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\nwant: %#+v\ngot: %#+v", want, got)
	}
}
