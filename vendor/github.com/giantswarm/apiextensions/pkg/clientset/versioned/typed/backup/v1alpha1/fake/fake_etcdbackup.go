/*
Copyright 2020 Giant Swarm GmbH.

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/backup/v1alpha1"
)

// FakeETCDBackups implements ETCDBackupInterface
type FakeETCDBackups struct {
	Fake *FakeBackupV1alpha1
}

var etcdbackupsResource = schema.GroupVersionResource{Group: "backup.giantswarm.io", Version: "v1alpha1", Resource: "etcdbackups"}

var etcdbackupsKind = schema.GroupVersionKind{Group: "backup.giantswarm.io", Version: "v1alpha1", Kind: "ETCDBackup"}

// Get takes name of the eTCDBackup, and returns the corresponding eTCDBackup object, and an error if there is any.
func (c *FakeETCDBackups) Get(name string, options v1.GetOptions) (result *v1alpha1.ETCDBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(etcdbackupsResource, name), &v1alpha1.ETCDBackup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ETCDBackup), err
}

// List takes label and field selectors, and returns the list of ETCDBackups that match those selectors.
func (c *FakeETCDBackups) List(opts v1.ListOptions) (result *v1alpha1.ETCDBackupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(etcdbackupsResource, etcdbackupsKind, opts), &v1alpha1.ETCDBackupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ETCDBackupList{ListMeta: obj.(*v1alpha1.ETCDBackupList).ListMeta}
	for _, item := range obj.(*v1alpha1.ETCDBackupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eTCDBackups.
func (c *FakeETCDBackups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(etcdbackupsResource, opts))
}

// Create takes the representation of a eTCDBackup and creates it.  Returns the server's representation of the eTCDBackup, and an error, if there is any.
func (c *FakeETCDBackups) Create(eTCDBackup *v1alpha1.ETCDBackup) (result *v1alpha1.ETCDBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(etcdbackupsResource, eTCDBackup), &v1alpha1.ETCDBackup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ETCDBackup), err
}

// Update takes the representation of a eTCDBackup and updates it. Returns the server's representation of the eTCDBackup, and an error, if there is any.
func (c *FakeETCDBackups) Update(eTCDBackup *v1alpha1.ETCDBackup) (result *v1alpha1.ETCDBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(etcdbackupsResource, eTCDBackup), &v1alpha1.ETCDBackup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ETCDBackup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeETCDBackups) UpdateStatus(eTCDBackup *v1alpha1.ETCDBackup) (*v1alpha1.ETCDBackup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(etcdbackupsResource, "status", eTCDBackup), &v1alpha1.ETCDBackup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ETCDBackup), err
}

// Delete takes name of the eTCDBackup and deletes it. Returns an error if one occurs.
func (c *FakeETCDBackups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(etcdbackupsResource, name), &v1alpha1.ETCDBackup{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeETCDBackups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(etcdbackupsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ETCDBackupList{})
	return err
}

// Patch applies the patch and returns the patched eTCDBackup.
func (c *FakeETCDBackups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ETCDBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(etcdbackupsResource, name, pt, data, subresources...), &v1alpha1.ETCDBackup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ETCDBackup), err
}
