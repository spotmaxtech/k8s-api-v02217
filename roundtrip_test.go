/*
Copyright 2018 The Kubernetes Authors.

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

package testing

import (
	"math/rand"
	"testing"

	admissionv1 "github.com/spotmaxtech/k8s-api-v02217/admission/v1"
	admissionv1beta1 "github.com/spotmaxtech/k8s-api-v02217/admission/v1beta1"
	admissionregv1 "github.com/spotmaxtech/k8s-api-v02217/admissionregistration/v1"
	admissionregv1beta1 "github.com/spotmaxtech/k8s-api-v02217/admissionregistration/v1beta1"
	apiserverinternalv1alpha1 "github.com/spotmaxtech/k8s-api-v02217/apiserverinternal/v1alpha1"
	appsv1 "github.com/spotmaxtech/k8s-api-v02217/apps/v1"
	appsv1beta1 "github.com/spotmaxtech/k8s-api-v02217/apps/v1beta1"
	appsv1beta2 "github.com/spotmaxtech/k8s-api-v02217/apps/v1beta2"
	authenticationv1 "github.com/spotmaxtech/k8s-api-v02217/authentication/v1"
	authenticationv1beta1 "github.com/spotmaxtech/k8s-api-v02217/authentication/v1beta1"
	authorizationv1 "github.com/spotmaxtech/k8s-api-v02217/authorization/v1"
	authorizationv1beta1 "github.com/spotmaxtech/k8s-api-v02217/authorization/v1beta1"
	autoscalingv1 "github.com/spotmaxtech/k8s-api-v02217/autoscaling/v1"
	autoscalingv2beta1 "github.com/spotmaxtech/k8s-api-v02217/autoscaling/v2beta1"
	autoscalingv2beta2 "github.com/spotmaxtech/k8s-api-v02217/autoscaling/v2beta2"
	batchv1 "github.com/spotmaxtech/k8s-api-v02217/batch/v1"
	batchv1beta1 "github.com/spotmaxtech/k8s-api-v02217/batch/v1beta1"
	certificatesv1 "github.com/spotmaxtech/k8s-api-v02217/certificates/v1"
	certificatesv1beta1 "github.com/spotmaxtech/k8s-api-v02217/certificates/v1beta1"
	coordinationv1 "github.com/spotmaxtech/k8s-api-v02217/coordination/v1"
	coordinationv1beta1 "github.com/spotmaxtech/k8s-api-v02217/coordination/v1beta1"
	corev1 "github.com/spotmaxtech/k8s-api-v02217/core/v1"
	discoveryv1 "github.com/spotmaxtech/k8s-api-v02217/discovery/v1"
	discoveryv1beta1 "github.com/spotmaxtech/k8s-api-v02217/discovery/v1beta1"
	eventsv1 "github.com/spotmaxtech/k8s-api-v02217/events/v1"
	eventsv1beta1 "github.com/spotmaxtech/k8s-api-v02217/events/v1beta1"
	extensionsv1beta1 "github.com/spotmaxtech/k8s-api-v02217/extensions/v1beta1"
	flowcontrolv1alpha1 "github.com/spotmaxtech/k8s-api-v02217/flowcontrol/v1alpha1"
	flowcontrolv1beta1 "github.com/spotmaxtech/k8s-api-v02217/flowcontrol/v1beta1"
	imagepolicyv1alpha1 "github.com/spotmaxtech/k8s-api-v02217/imagepolicy/v1alpha1"
	networkingv1 "github.com/spotmaxtech/k8s-api-v02217/networking/v1"
	networkingv1beta1 "github.com/spotmaxtech/k8s-api-v02217/networking/v1beta1"
	nodev1 "github.com/spotmaxtech/k8s-api-v02217/node/v1"
	nodev1alpha1 "github.com/spotmaxtech/k8s-api-v02217/node/v1alpha1"
	nodev1beta1 "github.com/spotmaxtech/k8s-api-v02217/node/v1beta1"
	policyv1 "github.com/spotmaxtech/k8s-api-v02217/policy/v1"
	policyv1beta1 "github.com/spotmaxtech/k8s-api-v02217/policy/v1beta1"
	rbacv1 "github.com/spotmaxtech/k8s-api-v02217/rbac/v1"
	rbacv1alpha1 "github.com/spotmaxtech/k8s-api-v02217/rbac/v1alpha1"
	rbacv1beta1 "github.com/spotmaxtech/k8s-api-v02217/rbac/v1beta1"
	schedulingv1 "github.com/spotmaxtech/k8s-api-v02217/scheduling/v1"
	schedulingv1alpha1 "github.com/spotmaxtech/k8s-api-v02217/scheduling/v1alpha1"
	schedulingv1beta1 "github.com/spotmaxtech/k8s-api-v02217/scheduling/v1beta1"
	storagev1 "github.com/spotmaxtech/k8s-api-v02217/storage/v1"
	storagev1alpha1 "github.com/spotmaxtech/k8s-api-v02217/storage/v1alpha1"
	storagev1beta1 "github.com/spotmaxtech/k8s-api-v02217/storage/v1beta1"

	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/apitesting/fuzzer"
	"k8s.io/apimachinery/pkg/api/apitesting/roundtrip"
	genericfuzzer "k8s.io/apimachinery/pkg/apis/meta/fuzzer"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

var groups = []runtime.SchemeBuilder{
	admissionv1beta1.SchemeBuilder,
	admissionv1.SchemeBuilder,
	admissionregv1beta1.SchemeBuilder,
	admissionregv1.SchemeBuilder,
	apiserverinternalv1alpha1.SchemeBuilder,
	appsv1beta1.SchemeBuilder,
	appsv1beta2.SchemeBuilder,
	appsv1.SchemeBuilder,
	authenticationv1beta1.SchemeBuilder,
	authenticationv1.SchemeBuilder,
	authorizationv1beta1.SchemeBuilder,
	authorizationv1.SchemeBuilder,
	autoscalingv1.SchemeBuilder,
	autoscalingv2beta1.SchemeBuilder,
	autoscalingv2beta2.SchemeBuilder,
	batchv1beta1.SchemeBuilder,
	batchv1.SchemeBuilder,
	certificatesv1.SchemeBuilder,
	certificatesv1beta1.SchemeBuilder,
	coordinationv1.SchemeBuilder,
	coordinationv1beta1.SchemeBuilder,
	corev1.SchemeBuilder,
	discoveryv1.SchemeBuilder,
	discoveryv1beta1.SchemeBuilder,
	eventsv1.SchemeBuilder,
	eventsv1beta1.SchemeBuilder,
	extensionsv1beta1.SchemeBuilder,
	flowcontrolv1alpha1.SchemeBuilder,
	flowcontrolv1beta1.SchemeBuilder,
	imagepolicyv1alpha1.SchemeBuilder,
	networkingv1.SchemeBuilder,
	networkingv1beta1.SchemeBuilder,
	nodev1.SchemeBuilder,
	nodev1alpha1.SchemeBuilder,
	nodev1beta1.SchemeBuilder,
	policyv1.SchemeBuilder,
	policyv1beta1.SchemeBuilder,
	rbacv1alpha1.SchemeBuilder,
	rbacv1beta1.SchemeBuilder,
	rbacv1.SchemeBuilder,
	schedulingv1alpha1.SchemeBuilder,
	schedulingv1beta1.SchemeBuilder,
	schedulingv1.SchemeBuilder,
	storagev1alpha1.SchemeBuilder,
	storagev1beta1.SchemeBuilder,
	storagev1.SchemeBuilder,
}

func TestRoundTripExternalTypes(t *testing.T) {
	scheme := runtime.NewScheme()
	codecs := serializer.NewCodecFactory(scheme)
	for _, builder := range groups {
		require.NoError(t, builder.AddToScheme(scheme))
	}
	seed := rand.Int63()
	// I'm only using the generic fuzzer funcs, but at some point in time we might need to
	// switch to specialized. For now we're happy with the current serialization test.
	fuzzer := fuzzer.FuzzerFor(genericfuzzer.Funcs, rand.NewSource(seed), codecs)

	roundtrip.RoundTripExternalTypes(t, scheme, codecs, fuzzer, nil)
}

func TestCompatibility(t *testing.T) {
	scheme := runtime.NewScheme()
	for _, builder := range groups {
		require.NoError(t, builder.AddToScheme(scheme))
	}
	roundtrip.NewCompatibilityTestOptions(scheme).Complete(t).Run(t)
}
