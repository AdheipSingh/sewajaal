// Code generated by client-gen. DO NOT EDIT.

package scheme

import (
	authenticationv1alpha1 "istio.io/client-go/pkg/apis/authentication/v1alpha1"
	configv1alpha2 "istio.io/client-go/pkg/apis/config/v1alpha2"
	networkingv1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	rbacv1alpha1 "istio.io/client-go/pkg/apis/rbac/v1alpha1"
	securityv1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var ParameterCodec = runtime.NewParameterCodec(Scheme)
var localSchemeBuilder = runtime.SchemeBuilder{
	authenticationv1alpha1.AddToScheme,
	configv1alpha2.AddToScheme,
	networkingv1alpha3.AddToScheme,
	rbacv1alpha1.AddToScheme,
	securityv1beta1.AddToScheme,
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//   import (
//     "k8s.io/client-go/kubernetes"
//     clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//     aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//   )
//
//   kclientset, _ := kubernetes.NewForConfig(c)
//   _ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(Scheme))
}
