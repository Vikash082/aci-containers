// Copyright 2016 Cisco Systems, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"errors"
	"flag"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/client/cache"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	"k8s.io/kubernetes/pkg/client/restclient"
	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/controller/informers"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/wait"
	"k8s.io/kubernetes/pkg/watch"
)

var (
	log         = logrus.New()
	kubeconfig  = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	nodeip      = flag.String("node", "", "IP of current node")
	metadataDir = flag.String("cnimetadatadir", "/var/lib/cni/opflex-networks", "Directory containing OpFlex CNI metadata")
	network     = flag.String("cninetwork", "opflex-k8s-network", "Name of CNI network")
	endpointDir = flag.String("endpointdir", "/var/lib/opflex-agent-ovs/endpoints/", "Directory for writing OpFlex endpoint metadata")
	serviceDir  = flag.String("servicedir", "/var/lib/opflex-agent-ovs/services/", "Directory for writing OpFlex anycast service metadata")
	vrfTenant   = flag.String("vrftenant", "common", "ACI tenant containing the VRF for kubernetes")
	vrf         = flag.String("vrf", "kubernetes-vrf", "ACI VRF name for for kubernetes")
	defaultEg   = flag.String("default-endpoint-group", "", "Default endpoint group annotation value")
	defaultSg   = flag.String("default-security-group", "", "Default security group annotation value")

	indexMutex     = &sync.Mutex{}
	depPods        = make(map[string]string)
	opflexEps      = make(map[string]*opflexEndpoint)
	opflexServices = make(map[string]*opflexService)

	namespaceInformer  cache.SharedIndexInformer
	podInformer        cache.SharedIndexInformer
	endpointsInformer  cache.SharedIndexInformer
	serviceInformer    cache.SharedIndexInformer
	deploymentInformer cache.SharedIndexInformer

	syncEnabled = false
)

func initNamespaceInformer(kubeClient *clientset.Clientset) {
	namespaceInformer = cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options api.ListOptions) (runtime.Object, error) {
				return kubeClient.Core().Namespaces().List(options)
			},
			WatchFunc: func(options api.ListOptions) (watch.Interface, error) {
				return kubeClient.Core().Namespaces().Watch(options)
			},
		},
		&api.Namespace{},
		controller.NoResyncPeriodFunc(),
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)
	namespaceInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    namespaceChanged,
		UpdateFunc: namespaceUpdated,
		DeleteFunc: namespaceChanged,
	})

	go namespaceInformer.GetController().Run(wait.NeverStop)
	go namespaceInformer.Run(wait.NeverStop)
}

func initPodInformer(kubeClient *clientset.Clientset) {
	podInformer = informers.NewPodInformer(kubeClient,
		controller.NoResyncPeriodFunc())
	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    podAdded,
		UpdateFunc: podUpdated,
		DeleteFunc: podDeleted,
	})

	go podInformer.GetController().Run(wait.NeverStop)
	go podInformer.Run(wait.NeverStop)
}

func initEndpointsInformer(kubeClient *clientset.Clientset) {
	endpointsInformer = cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options api.ListOptions) (runtime.Object, error) {
				return kubeClient.Core().Endpoints(api.NamespaceAll).List(options)
			},
			WatchFunc: func(options api.ListOptions) (watch.Interface, error) {
				return kubeClient.Core().Endpoints(api.NamespaceAll).Watch(options)
			},
		},
		&api.Endpoints{},
		controller.NoResyncPeriodFunc(),
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)
	endpointsInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    endpointsChanged,
		UpdateFunc: endpointsUpdated,
		DeleteFunc: endpointsChanged,
	})

	go endpointsInformer.GetController().Run(wait.NeverStop)
	go endpointsInformer.Run(wait.NeverStop)
}

func initServiceInformer(kubeClient *clientset.Clientset) {
	serviceInformer = cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options api.ListOptions) (runtime.Object, error) {
				return kubeClient.Core().Services(api.NamespaceAll).List(options)
			},
			WatchFunc: func(options api.ListOptions) (watch.Interface, error) {
				return kubeClient.Core().Services(api.NamespaceAll).Watch(options)
			},
		},
		&api.Service{},
		controller.NoResyncPeriodFunc(),
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)
	serviceInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    serviceAdded,
		UpdateFunc: serviceUpdated,
		DeleteFunc: serviceDeleted,
	})

	go serviceInformer.GetController().Run(wait.NeverStop)
	go serviceInformer.Run(wait.NeverStop)
}

func initDeploymentInformer(kubeClient *clientset.Clientset) {
	deploymentInformer = cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options api.ListOptions) (runtime.Object, error) {
				return kubeClient.Extensions().Deployments(api.NamespaceAll).List(options)
			},
			WatchFunc: func(options api.ListOptions) (watch.Interface, error) {
				return kubeClient.Extensions().Deployments(api.NamespaceAll).Watch(options)
			},
		},
		&extensions.Deployment{},
		controller.NoResyncPeriodFunc(),
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)
	deploymentInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    deploymentAdded,
		UpdateFunc: deploymentUpdated,
		DeleteFunc: deploymentDeleted,
	})

	go deploymentInformer.GetController().Run(wait.NeverStop)
	go deploymentInformer.Run(wait.NeverStop)
}

func main() {
	flag.Parse()

	if nodeip == nil || *nodeip == "" {
		err := errors.New("Node IP not specified")
		log.Error(err.Error())
		panic(err.Error())
	}

	log.WithFields(logrus.Fields{
		"kubeconfig": *kubeconfig,
		"nodeip":     *nodeip,
	}).Info("Starting")

	var config *restclient.Config
	var err error
	if kubeconfig != nil {
		// use kubeconfig file from command line
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}
	} else {
		// creates the in-cluster config
		config, err = restclient.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
	}

	// creates the client
	kubeClient, err := clientset.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	var wg sync.WaitGroup
	wg.Add(1)

	initNamespaceInformer(kubeClient)
	initDeploymentInformer(kubeClient)
	initPodInformer(kubeClient)

	initEndpointsInformer(kubeClient)
	initServiceInformer(kubeClient)

	go func() {
		time.Sleep(time.Second * 5)
		syncEnabled = true
		indexMutex.Lock()
		defer indexMutex.Unlock()
		syncServices()
		syncEps()
	}()

	wg.Wait()
}