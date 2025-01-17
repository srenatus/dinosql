package dinosql

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kyleconroy/dinosql/internal/pg"
)

func TestColumnsToStruct(t *testing.T) {
	cols := []pg.Column{
		{Name: "other", DataType: "text", NotNull: true},
		{Name: "count", DataType: "integer", NotNull: true},
		{Name: "count", DataType: "integer", NotNull: true},
	}

	r := Result{}
	actual := r.columnsToStruct("Foo", cols)
	expected := &GoStruct{
		Name: "Foo",
		Fields: []GoField{
			{Name: "Other", Type: "string", Tags: map[string]string{"json": "other"}},
			{Name: "Count", Type: "int", Tags: map[string]string{"json": "count"}},
			{Name: "Count_2", Type: "int", Tags: map[string]string{"json": "count_2"}},
		},
	}
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("struct mismatch: \n%s", diff)
	}
}
