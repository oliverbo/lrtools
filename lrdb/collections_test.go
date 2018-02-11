package lrdb

import "testing"

func TestGetCollectionById(t *testing.T) {
	// Create new collection
	c := GetCollectionById(5)
	if c == nil {
		t.Error("No collection pointer returned")
	}
	if c.localId != 5 || c.Name != "" {
		t.Errorf("Incorrect values %d, %s", c.localId, c.Name)
	}


	// retrieve existing collection
	c.Name = "Test"
	c = GetCollectionById(5)
	if c == nil {
		t.Error("No collection pointer returned")
	}
	if c.localId != 5 || c.Name != "Test" {
		t.Errorf("Incorrect values %d, %s", c.localId, c.Name)
	}
}

func TestCollection_AppendChild(t *testing.T) {
	c := Collection {localId:5, Name:"Test"}
	child := Collection {localId:50, Name:"TestChild"}

	// Append first child
	c.AppendChild(&child)

	if c.first != &child || c.last != &child {
		t.Error("First Child: Parent collection not correctly wired")
	}
	if child.parent != &c || child.nextSibling != nil {
		t.Error("First Child: Child collection not correctly wired")
	}

	// Append second child
	child2 := Collection{localId:51, Name:"TestChild2"}
	c.AppendChild(&child2)
	if c.first != &child || c.last != &child2 {
		t.Error("2nd Child: Parent not correctly wired")
	}
	if child.nextSibling != &child2 {
		t.Error("2nd Child: First Child not correctly wired")
	}
	if child2.parent != &c || child2.nextSibling != nil {
		t.Error("2nd Child: 2nd child collection not correctly wired")
	}
}