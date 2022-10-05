/*
Copyright 2022.

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

package controllers

import (
	"context"
	"errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	liquibasev1beta1 "github.com/kubbee/liquibase-db-changelog-operator/api/v1beta1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
)

// LiquiBaseRollbackReconciler reconciles a LiquiBaseRollback object
type LiquiBaseRollbackReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=liquibase.kubbee.tech,resources=liquibaserollbacks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=liquibase.kubbee.tech,resources=liquibaserollbacks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=liquibase.kubbee.tech,resources=liquibaserollbacks/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the LiquiBaseRollback object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *LiquiBaseRollbackReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := ctrl.LoggerFrom(ctx)

	//
	logger.Info("start::LiquiBaseRelease")

	liquiBase := &liquibasev1beta1.LiquiBaseRollback{}

	//
	if err := r.Get(ctx, req.NamespacedName, liquiBase); err != nil {
		if k8sErrors.IsNotFound(err) {
			logger.Info("Kafka Consumenr Not Found.")

			if !liquiBase.ObjectMeta.DeletionTimestamp.IsZero() {
				logger.Info("Was marked for deletion.")
				return reconcile.Result{}, nil // implementing the nil in the future
			}
		}
		return reconcile.Result{}, nil
	}

	if req.NamespacedName.Namespace != liquiBase.Namespace {
		logger.Info("Was marked for deletion.")
		return reconcile.Result{}, errors.New("the namespace resquested is different of the informed")
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *LiquiBaseRollbackReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&liquibasev1beta1.LiquiBaseRollback{}).
		Complete(r)
}
