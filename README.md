# nestedset
Package for manage nested sets in golang projects.

### Install

```
go get github.com/juggleru/nestedset
```

### Usage

To manage in nested set your data types add to your type `*Node` and init it.
```go
package main

import (
    "github.com/juggleru/nestedset"
    "fmt"
    "strings"
)

type MySomeType struct {
    *nestedset.Node // add it to your any type

    // type vars
    MyId string
    MyName string
}

// Init it in instance creation
func NewMySomeType() *MySomeType {
    return &MySomeType{
        Node: nestedset.NewNode(),
    }
}

// You can redefine NodeInterface functions

// Return your type
func (t *MySomeType) Type() string {
    return "mysometype"
}

// Return your inner id
func (t *MySomeType) Id() string {
    return t.MyId
}

// Return your inner name
func (t *MySomeType) Name() string {
    return t.MyName
}

// Set your inner id or generate it
func (t *MySomeType) SetId(id int) {
    t.MyId = id // or t.MyId = getSomeNewId()
}

// Set your inner name
func (t *MySomeType) SetName(name string) {
    t.MyName = name
}

func main() { 
    
    ns := nestedset.NewNestedSet()

    // create 3 new nodes
    node1 := NewMySomeType()
    node1.MyName = "Node 1"
    node2 := NewMySomeType()
    node2.MyName = "Node 2"
    node3 := NewMySomeType()
    node3.MyName = "Node 3"

    ns.Add(node1, nil)   // add node to root
    ns.Add(node2, nil)   // add node to root
    ns.Add(node3, node1) // add node to node1

    ns.Move(node3, node2) // move node3 from node1 to node2

    branch := ns.Branch(nil) // get full tree

    // print tree with indents	
    for _, n := range branch {
   	    fmt.Print(strings.Repeat("..", n.Level()))
   	    fmt.Printf("%s lvl:%d, left:%d, right:%d\n", n.Name(), n.Level(), n.Left(), n.Right())
    }
}
```
### Documentation

https://godoc.org/github.com/juggleru/nestedset

### TODO

Add implementation moving node up/down in same branch.

### Support

<a href="https://www.buymeacoffee.com/juggle" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>
