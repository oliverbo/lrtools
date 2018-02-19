package lrdb

import (
	"container/list"
	"strings"
)

// Collection describes a Lightroom collection by its name
type Collection struct {
	Name string
	localId int64
	parent *Collection
	nextSibling *Collection
	first *Collection
	last *Collection
}

const CollectionRootId  = -1
const PathSeparator = "/"

var root *Collection
var index map[int64] *Collection

// init initializes the collection tree
func init() {
	index = make(map[int64] *Collection)
	root = new(Collection)
	root.localId = CollectionRootId
	root.Name = "ROOT"
	index[CollectionRootId] = root
}

// FindCollectionById returns the collection with the ID 'localId'. If this collection does not yet exist,
// a new record will be created
func FindCollectionById(localId int64) *Collection {
	c, ok := index[localId]
	if !ok {
		c = new(Collection)
		c.localId = localId
		index[localId] = c
	}
	return c
}

// FindCollectionByPath returns the collection with the supplied absolute path in the /name/name/.../name format
func FindCollectionByPath(path string) *Collection {
	names := strings.Split(path, PathSeparator)

	c := GetCollectionRoot()
	for _, n := range names {
		if n != "" {
			c = c.FindChildByName(n)
			if c == nil {
				break
			}
		}
	}

	return c
}

// GetCollectionRoot returns the root of the collection tree
func GetCollectionRoot() *Collection {
	return root
}

func (c Collection) FindChildByName(name string) *Collection {
	for cc := c.first; cc != nil; cc = cc.nextSibling {
		if cc.Name == name {
			return cc
		}
	}
	return nil
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

func (c *Collection) VisitChildren(level int, includeParent bool, v func(level int, c *Collection)) {
	for f := c.first; f != nil; f = f.nextSibling {
		if includeParent || level == 1 {
			v(level, f)
		}
		if level > 1 {
			f.VisitChildren(level-1, includeParent, v)
		}
	}
}

func (c Collection) Path() string {
	var names list.List
	names.PushBack(c.Name)
	for p := c.parent; p != nil; p = p.parent {
		if p.localId != CollectionRootId {
			names.PushBack(p.Name)
		}
	}
	var path string
	for e := names.Back(); e != nil; e = e.Prev() {
		path = path + PathSeparator + e.Value.(string)
	}

	return path
}