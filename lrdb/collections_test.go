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

func TestCollection_VisitChildren(t *testing.T) {
	c := Collection{localId:1, Name:"Parent"}
	c1 := Collection{localId:2, Name:"Child1"}
	c2 := Collection{localId:3, Name:"Child2"}
	c3 := Collection{localId:4, Name:"Child3"}
	c.AppendChild(&c1)
	c.AppendChild(&c2)
	c.AppendChild(&c3)

	var visited string
	f := func(c *Collection) {visited += c.Name}
	c.VisitChildren(f)

	if visited != "Child1Child2Child3" {
		t.Errorf("Vistor failed (%s)", visited)
	}
}