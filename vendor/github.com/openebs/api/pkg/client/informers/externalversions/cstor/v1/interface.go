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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/openebs/api/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// CStorPoolClusters returns a CStorPoolClusterInformer.
	CStorPoolClusters() CStorPoolClusterInformer
	// CStorPoolInstances returns a CStorPoolInstanceInformer.
	CStorPoolInstances() CStorPoolInstanceInformer
	// CStorVolumes returns a CStorVolumeInformer.
	CStorVolumes() CStorVolumeInformer
	// CStorVolumeClaims returns a CStorVolumeClaimInformer.
	CStorVolumeClaims() CStorVolumeClaimInformer
	// CStorVolumePolicies returns a CStorVolumePolicyInformer.
	CStorVolumePolicies() CStorVolumePolicyInformer
	// CStorVolumeReplicas returns a CStorVolumeReplicaInformer.
	CStorVolumeReplicas() CStorVolumeReplicaInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// CStorPoolClusters returns a CStorPoolClusterInformer.
func (v *version) CStorPoolClusters() CStorPoolClusterInformer {
	return &cStorPoolClusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CStorPoolInstances returns a CStorPoolInstanceInformer.
func (v *version) CStorPoolInstances() CStorPoolInstanceInformer {
	return &cStorPoolInstanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CStorVolumes returns a CStorVolumeInformer.
func (v *version) CStorVolumes() CStorVolumeInformer {
	return &cStorVolumeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CStorVolumeClaims returns a CStorVolumeClaimInformer.
func (v *version) CStorVolumeClaims() CStorVolumeClaimInformer {
	return &cStorVolumeClaimInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CStorVolumePolicies returns a CStorVolumePolicyInformer.
func (v *version) CStorVolumePolicies() CStorVolumePolicyInformer {
	return &cStorVolumePolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CStorVolumeReplicas returns a CStorVolumeReplicaInformer.
func (v *version) CStorVolumeReplicas() CStorVolumeReplicaInformer {
	return &cStorVolumeReplicaInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
