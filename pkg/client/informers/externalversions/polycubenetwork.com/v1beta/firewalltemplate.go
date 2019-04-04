/*
Copyright The Kubernetes Authors.

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

package v1beta

import (
	time "time"

	polycubenetworkcomv1beta "github.com/SunSince90/polycube-firewall-template/pkg/apis/polycubenetwork.com/v1beta"
	versioned "github.com/SunSince90/polycube-firewall-template/pkg/client/clientset/versioned"
	internalinterfaces "github.com/SunSince90/polycube-firewall-template/pkg/client/informers/externalversions/internalinterfaces"
	v1beta "github.com/SunSince90/polycube-firewall-template/pkg/client/listers/polycubenetwork.com/v1beta"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FirewallTemplateInformer provides access to a shared informer and lister for
// FirewallTemplates.
type FirewallTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta.FirewallTemplateLister
}

type firewallTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFirewallTemplateInformer constructs a new informer for FirewallTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFirewallTemplateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFirewallTemplateInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFirewallTemplateInformer constructs a new informer for FirewallTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFirewallTemplateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolycubenetworkV1beta().FirewallTemplates(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolycubenetworkV1beta().FirewallTemplates(namespace).Watch(options)
			},
		},
		&polycubenetworkcomv1beta.FirewallTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *firewallTemplateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFirewallTemplateInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *firewallTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&polycubenetworkcomv1beta.FirewallTemplate{}, f.defaultInformer)
}

func (f *firewallTemplateInformer) Lister() v1beta.FirewallTemplateLister {
	return v1beta.NewFirewallTemplateLister(f.Informer().GetIndexer())
}