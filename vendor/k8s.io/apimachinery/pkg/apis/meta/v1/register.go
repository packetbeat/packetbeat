/*
Copyright 2014 The Kubernetes Authors.

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

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

// GroupName is the group name for this API.
const GroupName = "meta.k8s.io"

var (
	// localSchemeBuilder is used to make compiler happy for autogenerated
	// conversions. However, it's not used.
	schemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &schemeBuilder
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1"}

// Unversioned is group version for unversioned API objects
// TODO: this should be v1 probably
var Unversioned = schema.GroupVersion{Group: "", Version: "v1"}

// WatchEventKind is name reserved for serializing watch events.
const WatchEventKind = "WatchEvent"

// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// scheme is the registry for the common types that adhere to the meta v1 API spec.
var scheme = runtime.NewScheme()

// ParameterCodec knows about query parameters used with the meta v1 API spec.
var ParameterCodec = runtime.NewParameterCodec(scheme)

var optionsTypes = []runtime.Object{
	&ListOptions{},
	&ExportOptions{},
	&GetOptions{},
	&DeleteOptions{},
	&CreateOptions{},
	&UpdateOptions{},
	&PatchOptions{},
}

// AddToGroupVersion registers common meta types into schemas.
func AddToGroupVersion(scheme *runtime.Scheme, groupVersion schema.GroupVersion) {
	scheme.AddKnownTypeWithName(groupVersion.WithKind(WatchEventKind), &WatchEvent{})
	scheme.AddKnownTypeWithName(
		schema.GroupVersion{Group: groupVersion.Group, Version: runtime.APIVersionInternal}.WithKind(WatchEventKind),
		&InternalEvent{},
	)
	// Supports legacy code paths, most callers should use metav1.ParameterCodec for now
	scheme.AddKnownTypes(groupVersion, optionsTypes...)
	// Register Unversioned types under their own special group
	scheme.AddUnversionedTypes(Unversioned,
		&Status{},
		&APIVersions{},
		&APIGroupList{},
		&APIGroup{},
		&APIResourceList{},
	)

	// register manually. This usually goes through the SchemeBuilder, which we cannot use here.
	utilruntime.Must(RegisterConversions(scheme))
	utilruntime.Must(RegisterDefaults(scheme))
}

// AddMetaToScheme registers base meta types into schemas.
func AddMetaToScheme(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Table{},
		&TableOptions{},
		&PartialObjectMetadata{},
		&PartialObjectMetadataList{},
	)

	return nil
}

func init() {
	scheme.AddUnversionedTypes(SchemeGroupVersion, optionsTypes...)

	utilruntime.Must(AddMetaToScheme(scheme))

	// register manually. This usually goes through the SchemeBuilder, which we cannot use here.
	utilruntime.Must(RegisterDefaults(scheme))
}
