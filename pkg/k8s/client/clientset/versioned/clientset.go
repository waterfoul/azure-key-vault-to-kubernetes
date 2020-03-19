/*
Copyright Sparebanken Vest

Based on the Kubernetes controller example at
https://github.com/kubernetes/sample-controller

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

package versioned

import (
	azurekeyvaultv1 "github.com/SparebankenVest/azure-key-vault-to-kubernetes/pkg/k8s/client/clientset/versioned/typed/azurekeyvault/v1"
	azurekeyvaultv1alpha1 "github.com/SparebankenVest/azure-key-vault-to-kubernetes/pkg/k8s/client/clientset/versioned/typed/azurekeyvault/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AzurekeyvaultV1alpha1() azurekeyvaultv1alpha1.AzurekeyvaultV1alpha1Interface
	AzurekeyvaultV1() azurekeyvaultv1.AzurekeyvaultV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Azurekeyvault() azurekeyvaultv1.AzurekeyvaultV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	azurekeyvaultV1alpha1 *azurekeyvaultv1alpha1.AzurekeyvaultV1alpha1Client
	azurekeyvaultV1       *azurekeyvaultv1.AzurekeyvaultV1Client
}

// AzurekeyvaultV1alpha1 retrieves the AzurekeyvaultV1alpha1Client
func (c *Clientset) AzurekeyvaultV1alpha1() azurekeyvaultv1alpha1.AzurekeyvaultV1alpha1Interface {
	return c.azurekeyvaultV1alpha1
}

// AzurekeyvaultV1 retrieves the AzurekeyvaultV1Client
func (c *Clientset) AzurekeyvaultV1() azurekeyvaultv1.AzurekeyvaultV1Interface {
	return c.azurekeyvaultV1
}

// Deprecated: Azurekeyvault retrieves the default version of AzurekeyvaultClient.
// Please explicitly pick a version.
func (c *Clientset) Azurekeyvault() azurekeyvaultv1.AzurekeyvaultV1Interface {
	return c.azurekeyvaultV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.azurekeyvaultV1alpha1, err = azurekeyvaultv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.azurekeyvaultV1, err = azurekeyvaultv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.azurekeyvaultV1alpha1 = azurekeyvaultv1alpha1.NewForConfigOrDie(c)
	cs.azurekeyvaultV1 = azurekeyvaultv1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.azurekeyvaultV1alpha1 = azurekeyvaultv1alpha1.New(c)
	cs.azurekeyvaultV1 = azurekeyvaultv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
