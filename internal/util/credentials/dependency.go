/*
Copyright 2025 The ORC Authors.

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
package credentials

import (
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	"github.com/k-orc/openstack-resource-controller/internal/util/dependency"
)

func AddCredentialsWatch[
	objectTP dependency.ObjectType[objectT],
	objectListTP dependency.ObjectListType[objectListT, objectT],
	depTP dependency.DependencyType[depT],

	objectT any, objectListT any, depT any,
](
	log logr.Logger,
	k8sClient client.Client,
	b *builder.Builder,
	credentialsDep dependency.DeletionGuardDependency[objectTP, objectListTP, depTP, objectT, objectListT, depT],
) error {
	credentialsWatchEventHandler, err := credentialsDep.WatchEventHandler(log, k8sClient)
	if err != nil {
		return err
	}

	b.Watches(&corev1.Secret{}, credentialsWatchEventHandler,
		// Only trigger a reconcile when the credentials are created. We
		// don't need to reconcile for updates.
		builder.WithPredicates(predicate.Funcs{
			CreateFunc: func(_ event.TypedCreateEvent[client.Object]) bool {
				return true
			},
		}))

	return nil
}
