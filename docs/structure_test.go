package docs

import "testing"

func TestCategoriesFilter(t *testing.T) {
	pages := Categories{
		NewCategory(
			"Forms",
			NewPage("Button"),
			NewPage("Checkbox"),
		),
		NewCategory(
			"Navigation",
			NewPage("Tabs"),
		),
	}

	tests := []struct {
		name       string
		query      string
		wantCounts map[string]int
	}{
		{
			name:  "matches page names case insensitively",
			query: "bUtToN",
			wantCounts: map[string]int{
				"Forms": 1,
			},
		},
		{
			name:  "matches category names",
			query: "forms",
			wantCounts: map[string]int{
				"Forms": 2,
			},
		},
		{
			name:  "matches multiple terms across categories",
			query: "forms tabs",
			wantCounts: map[string]int{
				"Forms":      2,
				"Navigation": 1,
			},
		},
		{
			name:  "empty query returns all pages",
			query: "",
			wantCounts: map[string]int{
				"Forms":      2,
				"Navigation": 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pages.Filter(tt.query)

			if len(got) != len(tt.wantCounts) {
				t.Fatalf("len(result) = %d, want %d", len(got), len(tt.wantCounts))
			}

			for _, category := range got {
				wantPages, ok := tt.wantCounts[category.Name]
				if !ok {
					t.Fatalf("unexpected category %q in result", category.Name)
				}

				if len(category.Pages) != wantPages {
					t.Fatalf("len(%s.Pages) = %d, want %d", category.Name, len(category.Pages), wantPages)
				}
			}
		})
	}
}
