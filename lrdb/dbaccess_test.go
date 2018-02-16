package lrdb

import (
	"testing"
	"database/sql"
	"strconv"
)

type testRow struct {
	localId int64
	name string
	parent sql.NullInt64
}

type testRows struct {
	index int
	data []testRow
}

var testData = []testRow {
	{1, "Collection1", sql.NullInt64{0, false}},
	{2, "Collection2", sql.NullInt64{0, false}},
	{3, "Collection3", sql.NullInt64{1, true}},
	{4, "Collection4", sql.NullInt64{1, true}},
	{5, "Collection5", sql.NullInt64{2, true}},
	{6, "Collection6", sql.NullInt64{2, true}},
	{7, "Collection7", sql.NullInt64{5, true}}}

func TestBuildCollectionTree(t *testing.T) {
	var rows RowLoader = &testRows{-1, testData}

	buildCollectionTree(rows)

	root := GetCollectionRoot()

	var idstring string
	f := func(c *Collection) {
		idstring += strconv.FormatInt(c.localId, 10)
	}
	root.VisitChildren(f)
	if idstring != "12" {
		t.Errorf("First level incorrect (%s)", idstring)
	}

	var c *Collection

	c = GetCollectionById(1)
	idstring = ""
	c.VisitChildren(f)
	if idstring != "34" {
		t.Errorf("Second level '1' incorrect (%s)", idstring)
	}

	c = GetCollectionById(2)
	idstring = ""
	c.VisitChildren(f)
	if idstring != "56" {
		t.Errorf("Second level '2' incorrect (%s)", idstring)
	}

	c = GetCollectionById(5)
	idstring = ""
	c.VisitChildren(f)
	if idstring != "7" {
		t.Errorf("Third level '5' incorrect (%s)", idstring)
	}
}

func (tr *testRows) Next() bool {
	tr.index++
	return tr.index < len(tr.data)
}

func (tr *testRows) Scan(dest ...interface{}) error {
	*(dest[0].(*int64)) = tr.data[tr.index].localId
	*(dest[1].(*string)) = tr.data[tr.index].name
	*(dest[2].(*sql.NullInt64)) = tr.data[tr.index].parent
	return nil
}