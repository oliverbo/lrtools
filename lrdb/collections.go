package lrdb

// Collection describes a Lightroom collection by its name
type Collection struct {
	Name string
	localId int
	parent *Collection
	nextSibling *Collection
	first *Collection
	last *Collection
}

const CollectionRootId  = -1

var root *Collection
var index map[int] *Collection

// init initializes the collection tree
func init() {
	index = make(map[int] *Collection)
	root = new(Collection)
	root.localId = CollectionRootId
	root.Name = "ROOT"
}

// GetCollectionById returns the collection with the ID 'localId'. If this collection does not yet exist,
// a new record will be created
func GetCollectionById(localId int) *Collection {
	c, ok := index[localId]
	if !ok {
		c = new(Collection)
		c.localId = localId
		index[localId] = c
	}
	return c
}

// GetCollectionRoot returns the root of the collection tree
func GetCollectionRoot() *Collection {
	return root
}

// AppendChild appends a new child 'child' to the collection
func (c *Collection) AppendChild(child *Collection) {
	if c.first != nil {
		c.last.nextSibling = child
		c.last = child
	} else {
		c.first = child
		c.last = child
	}
	child.parent = c
}

func (c *Collection) VisitChildren(v func(c *Collection)) {
	for f := c.first; f != nil; f = f.nextSibling {
		v(f)
	}
}