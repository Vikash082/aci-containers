/***
Copyright 2019 Cisco Systems Inc. All rights reserved.

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
	v1 "github.com/noironetworks/aci-containers/pkg/gbpcrd/apis/aci.aw/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EpgLister helps list Epgs.
type EpgLister interface {
	// List lists all Epgs in the indexer.
	List(selector labels.Selector) (ret []*v1.Epg, err error)
	// Epgs returns an object that can list and get Epgs.
	Epgs(namespace string) EpgNamespaceLister
	EpgListerExpansion
}

// epgLister implements the EpgLister interface.
type epgLister struct {
	indexer cache.Indexer
}

// NewEpgLister returns a new EpgLister.
func NewEpgLister(indexer cache.Indexer) EpgLister {
	return &epgLister{indexer: indexer}
}

// List lists all Epgs in the indexer.
func (s *epgLister) List(selector labels.Selector) (ret []*v1.Epg, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Epg))
	})
	return ret, err
}

// Epgs returns an object that can list and get Epgs.
func (s *epgLister) Epgs(namespace string) EpgNamespaceLister {
	return epgNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EpgNamespaceLister helps list and get Epgs.
type EpgNamespaceLister interface {
	// List lists all Epgs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Epg, err error)
	// Get retrieves the Epg from the indexer for a given namespace and name.
	Get(name string) (*v1.Epg, error)
	EpgNamespaceListerExpansion
}

// epgNamespaceLister implements the EpgNamespaceLister
// interface.
type epgNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Epgs in the indexer for a given namespace.
func (s epgNamespaceLister) List(selector labels.Selector) (ret []*v1.Epg, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Epg))
	})
	return ret, err
}

// Get retrieves the Epg from the indexer for a given namespace and name.
func (s epgNamespaceLister) Get(name string) (*v1.Epg, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("epg"), name)
	}
	return obj.(*v1.Epg), nil
}
