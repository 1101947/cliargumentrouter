package treewrapper

import (
	"fmt"
)

type WrapperNode struct {
	WNode *WrappedNode
	Children map[string]*WrapperNode
}

func (W WrapperNode) Purify() (wrapperNode, error) {
	children := &(W.WNode).GetChildren()
	if len(W.Children) != len(children) {
		return wrapperNode{}, fmt.Errorf("Number of wrapper node's children is not equal to number of wrapped node children, but it must be equal")
	}
	for k,wrappedChildren := range(children) {
		wrapperChildren, exists := W.Children[k]
		if !exists {
			return wrapperNode{}, fmt.Errorf("Didn't find wrapper for wrapped node children, but must to") 
		}
	}
	pChildren := map[string]*wrapperNode{}
	for k,v := range(W.Children) {
		pNode, err := &v.Purify()
		if err != nil {
			return wrapperNode{}, fmt.Errorf("Purifing node children, got :: %w", err) 
		}
		pChildren[k] = pNode 
	}
	return wrapperNode{
		wnode: W.WNode,
		children: pChildren,
	}, nil
}

type wrapperNode struct {
	wnode *WrappedNode
	children map[string]*wrpperNode
}

func (w wrapperNode) GetNode() (WrappedNode, error) {
	if w.wnode == nil {
		return wrappedNode{}, fmt.Errorf("No wrapped node linked")
	}
	return &w.wnode, nil
}

type WrappedDirtyNode interface {
	GetChildren() map[string]*WrappedNode
}

type WrappedPureNode interface {
	GetChildren() map[string]*wrappedNode
}
