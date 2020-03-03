/*
Copyright 2020 The OpenEBS Authors

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	cstorv1 "github.com/openebs/api/pkg/apis/cstor/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCStorPoolInstances implements CStorPoolInstanceInterface
type FakeCStorPoolInstances struct {
	Fake *FakeCstorV1
	ns   string
}

var cstorpoolinstancesResource = schema.GroupVersionResource{Group: "cstor.openebs.io", Version: "v1", Resource: "cstorpoolinstances"}

var cstorpoolinstancesKind = schema.GroupVersionKind{Group: "cstor.openebs.io", Version: "v1", Kind: "CStorPoolInstance"}

// Get takes name of the cStorPoolInstance, and returns the corresponding cStorPoolInstance object, and an error if there is any.
func (c *FakeCStorPoolInstances) Get(name string, options v1.GetOptions) (result *cstorv1.CStorPoolInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cstorpoolinstancesResource, c.ns, name), &cstorv1.CStorPoolInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cstorv1.CStorPoolInstance), err
}

// List takes label and field selectors, and returns the list of CStorPoolInstances that match those selectors.
func (c *FakeCStorPoolInstances) List(opts v1.ListOptions) (result *cstorv1.CStorPoolInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cstorpoolinstancesResource, cstorpoolinstancesKind, c.ns, opts), &cstorv1.CStorPoolInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cstorv1.CStorPoolInstanceList{ListMeta: obj.(*cstorv1.CStorPoolInstanceList).ListMeta}
	for _, item := range obj.(*cstorv1.CStorPoolInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cStorPoolInstances.
func (c *FakeCStorPoolInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cstorpoolinstancesResource, c.ns, opts))

}

// Create takes the representation of a cStorPoolInstance and creates it.  Returns the server's representation of the cStorPoolInstance, and an error, if there is any.
func (c *FakeCStorPoolInstances) Create(cStorPoolInstance *cstorv1.CStorPoolInstance) (result *cstorv1.CStorPoolInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cstorpoolinstancesResource, c.ns, cStorPoolInstance), &cstorv1.CStorPoolInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cstorv1.CStorPoolInstance), err
}

// Update takes the representation of a cStorPoolInstance and updates it. Returns the server's representation of the cStorPoolInstance, and an error, if there is any.
func (c *FakeCStorPoolInstances) Update(cStorPoolInstance *cstorv1.CStorPoolInstance) (result *cstorv1.CStorPoolInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cstorpoolinstancesResource, c.ns, cStorPoolInstance), &cstorv1.CStorPoolInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cstorv1.CStorPoolInstance), err
}

// Delete takes name of the cStorPoolInstance and deletes it. Returns an error if one occurs.
func (c *FakeCStorPoolInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cstorpoolinstancesResource, c.ns, name), &cstorv1.CStorPoolInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCStorPoolInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cstorpoolinstancesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &cstorv1.CStorPoolInstanceList{})
	return err
}

// Patch applies the patch and returns the patched cStorPoolInstance.
func (c *FakeCStorPoolInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cstorv1.CStorPoolInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cstorpoolinstancesResource, c.ns, name, pt, data, subresources...), &cstorv1.CStorPoolInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cstorv1.CStorPoolInstance), err
}
