package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

const FileName = "export"

type Exporter func(string, Release) error

func ToJSON(dir string, r Release) error {
	f, err := os.Create(filepath.Join(dir, FileName+".json"))
	if err != nil {
		return err
	}

	defer f.Close()
	if err := json.NewEncoder(f).Encode(r); err != nil {
		return err
	}

	return nil
}

func ToCSV(dir string, r Release) error {
	f, err := os.Create(filepath.Join(dir, FileName+".csv"))
	if err != nil {
		return err
	}

	defer f.Close()
	w := csv.NewWriter(f)
	w.Write([]string{
		"Group",
		"Group Aliases",
		"Muscle",
		"Target",
		"Target Aliases",
	})

	for _, g := range r.Groups {
		for _, m := range g.Muscles {
			if m.Targets == nil {
				err := w.Write([]string{
					g.Name,
					strings.Join(g.Aliases, ","),
					m.Name,
					"",
					"",
				})

				if err != nil {
					return err
				}

				continue
			}

			for _, t := range m.Targets {
				err := w.Write([]string{
					g.Name,
					strings.Join(g.Aliases, ","),
					m.Name,
					t.Name,
					strings.Join(t.Aliases, ","),
				})

				if err != nil {
					return err
				}
			}
		}
	}

	w.Flush()
	return w.Error()
}
