/*
Copyright 2024 The Kubermatic Kubernetes Platform contributors.

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

package resources

import (
	"k8c.io/reconciler/pkg/reconciling"

	rbacv1 "k8s.io/api/rbac/v1"
)

const (
	operatorRoleName = "flatcar-linux-update-operator"
)

func OperatorRoleReconciler() reconciling.NamedRoleReconcilerFactory {
	return func() (string, reconciling.RoleReconciler) {
		return operatorRoleName, func(cr *rbacv1.Role) (*rbacv1.Role, error) {
			cr.Rules = []rbacv1.PolicyRule{
				{
					APIGroups: []string{""},
					Resources: []string{"configmaps"},
					Verbs: []string{
						"create",
					},
				},
				{
					APIGroups:     []string{""},
					Resources:     []string{"configmaps"},
					ResourceNames: []string{"flatcar-linux-update-operator-lock"},
					Verbs: []string{
						"get",
						"update",
					},
				},
				{
					APIGroups: []string{""},
					Resources: []string{"events"},
					Verbs: []string{
						"create",
						"watch",
					},
				},
				{
					APIGroups: []string{"coordination.k8s.io"},
					Resources: []string{"leases"},
					Verbs: []string{
						"create",
					},
				},
				{
					APIGroups:     []string{"coordination.k8s.io"},
					Resources:     []string{"leases"},
					ResourceNames: []string{"flatcar-linux-update-operator-lock"},
					Verbs: []string{
						"get",
						"update",
					},
				},
			}
			return cr, nil
		}
	}
}