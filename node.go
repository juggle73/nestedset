// Copyright 2018 Ara Israelyan. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.


package nestedset

// NodeInterface is the interface implemented by types that can be used by nodes in nested set
type NodeInterface interface {
	Type() string // Returns type of node
	Name() string // Returns name of node

	Id() int    // Returns id of node
	Level() int // Returns level of node
	Left() int  // Returns left of node
	Right() int // Returns right of node

	SetId(int)      // Sets node id
	SetName(string) // Sets node name
	SetLevel(int)   // Sets node level
	SetLeft(int)    // Sets node left
	SetRight(int)   // Sets node right
}

// Node represents generic node type with NodeInterface implementation
type Node struct {
	NodeId    int    `json:"id"`
	NodeName  string `json:"node_name"`
	NodeLevel int    `json:"level"`
	NodeLeft  int    `json:"left"`
	NodeRight int    `json:"right"`
}

// NewNode returns a new Node instance
func NewNode() *Node {
	return &Node{}
}

// Type implements NodeInterface.Type() and returns "generic" type
func (n Node) Type() string {
	return "generic"
}

func (n Node) Name() string {
	return n.NodeName
}

func (n Node) Id() int {
	return n.NodeId
}

func (n Node) Level() int {

	return n.NodeLevel
}

func (n Node) Left() int {
	return n.NodeLeft
}

func (n Node) Right() int {
	return n.NodeRight
}

func (n *Node) SetId(id int) {
	n.NodeId = id
}

func (n *Node) SetName(name string) {
	n.NodeName = name
}

func (n *Node) SetLevel(level int) {
	n.NodeLevel = level
}

func (n *Node) SetLeft(left int) {
	n.NodeLeft = left
}

func (n *Node) SetRight(right int) {
	n.NodeRight = right
}
