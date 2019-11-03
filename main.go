package main

import (
	"log"
	"os"

	authenticationv1alpha1 "istio.io/api/authentication/v1alpha1"
	"istio.io/client-go/pkg/apis/authentication/v1alpha1"
	"istio.io/client-go/pkg/clientset/versioned"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := os.Getenv("KUBECONFIG")
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Failed to create k8s rest client: %s", err)
	}

	ic, err := versioned.NewForConfig(restConfig)
	if err != nil {
		log.Fatalf("Failed to create istio client: %s", err)
	}
	vsList, err := ic.NetworkingV1alpha3().VirtualServices("default").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to get VirtualService in %s namespace: %s", "default", err)
	}
	for i := range vsList.Items {
		vs := vsList.Items[i]
		log.Printf("Index: %d VirtualService Hosts: %+v\n", i, vs.Spec.GetHosts())
	}
	// apiVersion: authentication.istio.io/v1alpha1
	// kind: Policy
	// metadata:
	//   name: default
	//   namespace: frod
	// spec:
	//   peers:
	//   - mtls:
	policy := &v1alpha1.Policy{
		ObjectMeta: metav1.ObjectMeta{
			Name: "default",
		},
		Spec: authenticationv1alpha1.Policy{
			Peers: []*authenticationv1alpha1.PeerAuthenticationMethod{
				{
					Params: &authenticationv1alpha1.PeerAuthenticationMethod_Mtls{
						Mtls: &authenticationv1alpha1.MutualTls{
							Mode: 1,
						},
					},
				},
			},
		},
	}

	a, err := ic.AuthenticationV1alpha1().Policies("default").Create(policy)
	if err != nil {
		log.Printf("%v", err)
	}

	log.Printf("%v", a)

}
