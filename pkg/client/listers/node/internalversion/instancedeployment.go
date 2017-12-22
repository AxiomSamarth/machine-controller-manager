// This file was automatically generated by lister-gen

package internalversion

import (
	node "github.com/gardener/node-controller-manager/pkg/apis/node"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InstanceDeploymentLister helps list InstanceDeployments.
type InstanceDeploymentLister interface {
	// List lists all InstanceDeployments in the indexer.
	List(selector labels.Selector) (ret []*node.InstanceDeployment, err error)
	// Get retrieves the InstanceDeployment from the index for a given name.
	Get(name string) (*node.InstanceDeployment, error)
	InstanceDeploymentListerExpansion
}

// instanceDeploymentLister implements the InstanceDeploymentLister interface.
type instanceDeploymentLister struct {
	indexer cache.Indexer
}

// NewInstanceDeploymentLister returns a new InstanceDeploymentLister.
func NewInstanceDeploymentLister(indexer cache.Indexer) InstanceDeploymentLister {
	return &instanceDeploymentLister{indexer: indexer}
}

// List lists all InstanceDeployments in the indexer.
func (s *instanceDeploymentLister) List(selector labels.Selector) (ret []*node.InstanceDeployment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*node.InstanceDeployment))
	})
	return ret, err
}

// Get retrieves the InstanceDeployment from the index for a given name.
func (s *instanceDeploymentLister) Get(name string) (*node.InstanceDeployment, error) {
	key := &node.InstanceDeployment{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(node.Resource("instancedeployment"), name)
	}
	return obj.(*node.InstanceDeployment), nil
}
