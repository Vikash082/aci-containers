/*
Copyright 2019 The Kubernetes Authors.

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

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	authorization "k8s.io/kubernetes/pkg/apis/authorization"
)

// SubjectAccessReviewLister helps list SubjectAccessReviews.
type SubjectAccessReviewLister interface {
	// List lists all SubjectAccessReviews in the indexer.
	List(selector labels.Selector) (ret []*authorization.SubjectAccessReview, err error)
	// Get retrieves the SubjectAccessReview from the index for a given name.
	Get(name string) (*authorization.SubjectAccessReview, error)
	SubjectAccessReviewListerExpansion
}

// subjectAccessReviewLister implements the SubjectAccessReviewLister interface.
type subjectAccessReviewLister struct {
	indexer cache.Indexer
}

// NewSubjectAccessReviewLister returns a new SubjectAccessReviewLister.
func NewSubjectAccessReviewLister(indexer cache.Indexer) SubjectAccessReviewLister {
	return &subjectAccessReviewLister{indexer: indexer}
}

// List lists all SubjectAccessReviews in the indexer.
func (s *subjectAccessReviewLister) List(selector labels.Selector) (ret []*authorization.SubjectAccessReview, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*authorization.SubjectAccessReview))
	})
	return ret, err
}

// Get retrieves the SubjectAccessReview from the index for a given name.
func (s *subjectAccessReviewLister) Get(name string) (*authorization.SubjectAccessReview, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(authorization.Resource("subjectaccessreview"), name)
	}
	return obj.(*authorization.SubjectAccessReview), nil
}
