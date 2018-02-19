package lrdb

import "testing"

func TestFindCollectionById(t *testing.T) {
	// Create new collection
	c := FindCollectionById(5)
	if c == nil {
		t.Error("No collection pointer returned")
	}
	if c.localId != 5 || c.Name != "" {
		t.Errorf("Incorrect values %d, %s", c.localId, c.Name)
	}


	// retrieve existing collection
	c.Name = "Test"
	c = FindCollectionById(5)
	if c == nil {
		t.Error("No collection pointer returned")
	}
	if c.localId != 5 || c.Name != "Test" {
		t.Errorf("Incorrect values %d, %s", c.localId, c.Name)
	}
}

func TestFindCollectionByPath(t *testing.T) {
	root = &Collection{localId:CollectionRootId, Name: "ROOT"}
	c1 := Collection{localId:1, Name: "Level 1"}
	c2 := Collection{localId:2, Name: "Level 2"}
	c3 := Collection{localId:3, Name: "Level 3"}
	c4 := Collection{localId:4, Name: "Level 4"}
	root.AppendChild(&c1)
	c1.AppendChild(&c2)
	c2.AppendChild(&c3)
	c3.AppendChild(&c4)

	// Find collection
	c := FindCollectionByPath("/Level 1/Level 2")
	if c == nil {
		t.Error("No collection returned")
	} else if c != &c2 {
		t.Errorf("1: Incorrect collection found %d", c.localId)
	}

	c = FindCollectionByPath("Level 1/Level 2/Level 3")
	if c == nil {
		t.Error("No collection returned")
	} else if c != &c3 {
		t.Errorf("2: Incorrect collection found %d", c.localId)
	}

	// Not find collection
	c = FindCollectionByPath("Invalid")
	if c != nil {
		t.Error("3: Should have returned nil")
	}
}

func TestCollection_FindChildByName(t *testing.T) {
	root = &Collection{localId:1, Name:"Parent"}
	c1 := Collection{localId:2, Name:"Child1"}
	c2 := Collection{localId:3, Name:"Child2"}
	c3 := Collection{localId:4, Name:"Child3"}
	root.AppendChild(&c1)
	root.AppendChild(&c2)
	root.AppendChild(&c3)

	cc := root.FindChildByName("Child2")
	if cc == nil {
		t.Error("1: No collection found")
	} else if cc != &c2 {
		t.Errorf("1: Incorrect collection found %d", cc.localId)
	}

	cc = root.FindChildByName("Invalid")
	if cc != nil {
		t.Error("2: Should have returned nil")
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
	root = &Collection{localId:1, Name:"Root"}
	c1 := Collection{localId:2, Name:"Child1"}
	c2 := Collection{localId:3, Name:"Child2"}
	c3 := Collection{localId:4, Name:"Child3"}
	c11 := Collection{localId:11, Name:"Child1-1"}
	c12 := Collection{localId:11, Name:"Child1-2"}
	c121 := Collection{localId:11, Name:"Child1-2-1"}
	c13 := Collection{localId:11, Name:"Child1-3"}
	root.AppendChild(&c1)
	c1.AppendChild(&c11)
	c1.AppendChild(&c12)
	c12.AppendChild(&c121)
	c1.AppendChild(&c13)
	root.AppendChild(&c2)
	root.AppendChild(&c3)

	var visited string
	f := func(l int, c *Collection) {visited += c.Name}
	// One level
	root.VisitChildren(1, true, f)
	if visited != "Child1Child2Child3" {
		t.Errorf("1: Vistor failed (%s)", visited)
	}
	// Two Levels
	visited = ""
	root.VisitChildren(2, true, f)
	if visited != "Child1Child1-1Child1-2Child1-3Child2Child3" {
		t.Errorf("2: Vistor failed (%s)", visited)
	}
	// Two Levels, one down
	visited = ""
	c1.VisitChildren(2, true, f)
	if visited != "Child1-1Child1-2Child1-2-1Child1-3" {
		t.Errorf("3: Vistor failed (%s)", visited)
	}
	// Two Levels, don't include parent
	visited = ""
	root.VisitChildren(2, false, f)
	if visited != "Child1-1Child1-2Child1-3" {
		t.Errorf("4: Vistor failed (%s)", visited)
	}
}

func TestCollection_Path(t *testing.T) {
	root = &Collection{localId:CollectionRootId, Name: "ROOT"}
	c1 := Collection{localId:1, Name: "Level 1"}
	c2 := Collection{localId:2, Name: "Level 2"}
	c3 := Collection{localId:3, Name: "Level 3"}
	c4 := Collection{localId:4, Name: "Level 4"}
	root.AppendChild(&c1)
	c1.AppendChild(&c2)
	c2.AppendChild(&c3)
	c3.AppendChild(&c4)

	path := c4.Path()

	if path != "/Level 1/Level 2/Level 3/Level 4" {
		t.Errorf("1: Incorrect collection path: '%s'", path)
	}

	path = c1.Path()
	if path != "/Level 1" {
		t.Errorf("2: Incorrect collection path: '%s'", path)
	}
}