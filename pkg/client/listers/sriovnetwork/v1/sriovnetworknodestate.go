/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/pliurh/sriov-network-operator/pkg/apis/sriovnetwork/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SriovNetworkNodeStateLister helps list SriovNetworkNodeStates.
type SriovNetworkNodeStateLister interface {
	// List lists all SriovNetworkNodeStates in the indexer.
	List(selector labels.Selector) (ret []*v1.SriovNetworkNodeState, err error)
	// SriovNetworkNodeStates returns an object that can list and get SriovNetworkNodeStates.
	SriovNetworkNodeStates(namespace string) SriovNetworkNodeStateNamespaceLister
	SriovNetworkNodeStateListerExpansion
}

// sriovNetworkNodeStateLister implements the SriovNetworkNodeStateLister interface.
type sriovNetworkNodeStateLister struct {
	indexer cache.Indexer
}

// NewSriovNetworkNodeStateLister returns a new SriovNetworkNodeStateLister.
func NewSriovNetworkNodeStateLister(indexer cache.Indexer) SriovNetworkNodeStateLister {
	return &sriovNetworkNodeStateLister{indexer: indexer}
}

// List lists all SriovNetworkNodeStates in the indexer.
func (s *sriovNetworkNodeStateLister) List(selector labels.Selector) (ret []*v1.SriovNetworkNodeState, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.SriovNetworkNodeState))
	})
	return ret, err
}

// SriovNetworkNodeStates returns an object that can list and get SriovNetworkNodeStates.
func (s *sriovNetworkNodeStateLister) SriovNetworkNodeStates(namespace string) SriovNetworkNodeStateNamespaceLister {
	return sriovNetworkNodeStateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SriovNetworkNodeStateNamespaceLister helps list and get SriovNetworkNodeStates.
type SriovNetworkNodeStateNamespaceLister interface {
	// List lists all SriovNetworkNodeStates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.SriovNetworkNodeState, err error)
	// Get retrieves the SriovNetworkNodeState from the indexer for a given namespace and name.
	Get(name string) (*v1.SriovNetworkNodeState, error)
	SriovNetworkNodeStateNamespaceListerExpansion
}

// sriovNetworkNodeStateNamespaceLister implements the SriovNetworkNodeStateNamespaceLister
// interface.
type sriovNetworkNodeStateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SriovNetworkNodeStates in the indexer for a given namespace.
func (s sriovNetworkNodeStateNamespaceLister) List(selector labels.Selector) (ret []*v1.SriovNetworkNodeState, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.SriovNetworkNodeState))
	})
	return ret, err
}

// Get retrieves the SriovNetworkNodeState from the indexer for a given namespace and name.
func (s sriovNetworkNodeStateNamespaceLister) Get(name string) (*v1.SriovNetworkNodeState, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("sriovnetworknodestate"), name)
	}
	return obj.(*v1.SriovNetworkNodeState), nil
}