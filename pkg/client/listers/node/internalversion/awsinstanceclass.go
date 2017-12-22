// This file was automatically generated by lister-gen

package internalversion

import (
	node "github.com/gardener/node-controller-manager/pkg/apis/node"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AWSInstanceClassLister helps list AWSInstanceClasses.
type AWSInstanceClassLister interface {
	// List lists all AWSInstanceClasses in the indexer.
	List(selector labels.Selector) (ret []*node.AWSInstanceClass, err error)
	// Get retrieves the AWSInstanceClass from the index for a given name.
	Get(name string) (*node.AWSInstanceClass, error)
	AWSInstanceClassListerExpansion
}

// aWSInstanceClassLister implements the AWSInstanceClassLister interface.
type aWSInstanceClassLister struct {
	indexer cache.Indexer
}

// NewAWSInstanceClassLister returns a new AWSInstanceClassLister.
func NewAWSInstanceClassLister(indexer cache.Indexer) AWSInstanceClassLister {
	return &aWSInstanceClassLister{indexer: indexer}
}

// List lists all AWSInstanceClasses in the indexer.
func (s *aWSInstanceClassLister) List(selector labels.Selector) (ret []*node.AWSInstanceClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*node.AWSInstanceClass))
	})
	return ret, err
}

// Get retrieves the AWSInstanceClass from the index for a given name.
func (s *aWSInstanceClassLister) Get(name string) (*node.AWSInstanceClass, error) {
	key := &node.AWSInstanceClass{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(node.Resource("awsinstanceclass"), name)
	}
	return obj.(*node.AWSInstanceClass), nil
}
