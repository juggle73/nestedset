// Copyright 2018 Ara Israelyan. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.


package nestedset

// NodeInterface is the interface implemented by types that can be used by nodes in nested set
type NodeInterface interface {
	Type() string // Returns type of node
	Name() string // Returns name of node

	Id() int64    // Returns id of node
	Level() int64 // Returns level of node
	Left() int64  // Returns left of node
	Right() int64 // Returns right of node

	SetId(int64)      // Sets node id
	SetName(string) // Sets node name
	SetLevel(int64)   // Sets node level
	SetLeft(int64)    // Sets node left
	SetRight(int64)   // Sets node right
}

// Node represents generic node type with NodeInterface implementation
type Node struct {
	NodeId    int64    `json:"id"`
	NodeName  string   `json:"node_name"`
	NodeLevel int64    `json:"level"`
	NodeLeft  int64    `json:"left"`
	NodeRight int64    `json:"right"`
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

func (n Node) Id() int64 {
	return n.NodeId
}

func (n Node) Level() int64 {

	return n.NodeLevel
}

func (n Node) Left() int64 {
	return n.NodeLeft
}

func (n Node) Right() int64 {
	return n.NodeRight
}

func (n *Node) SetId(id int64) {
	n.NodeId = id
}

func (n *Node) SetName(name string) {
	n.NodeName = name
}

func (n *Node) SetLevel(level int64) {
	n.NodeLevel = level
}

func (n *Node) SetLeft(left int64) {
	n.NodeLeft = left
}

func (n *Node) SetRight(right int64) {
	n.NodeRight = right
}
