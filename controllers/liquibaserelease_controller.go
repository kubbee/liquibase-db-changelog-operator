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

	corev1 "k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
)

// LiquiBaseReleaseReconciler reconciles a LiquiBaseRelease object
type LiquiBaseReleaseReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=liquibase.kubbee.tech,resources=liquibasereleases,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=liquibase.kubbee.tech,resources=liquibasereleases/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=liquibase.kubbee.tech,resources=liquibasereleases/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the LiquiBaseRelease object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *LiquiBaseReleaseReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := ctrl.LoggerFrom(ctx)

	//
	logger.Info("start::LiquiBaseRelease")

	liquiBase := &liquibasev1beta1.LiquiBaseRelease{}

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

	return r.applyChangelog(ctx, req, liquiBase)
}

// applyChangelog
func (r *LiquiBaseReleaseReconciler) applyChangelog(ctx context.Context, req ctrl.Request, liquiBase *liquibasev1beta1.LiquiBaseRelease) (ctrl.Result, error) {

	r.readConfigMap()

	// TODO: se tiver habilitado o secretmanager deve buscas user e senha lá

	// TODO: implamentar qual banco de dado será

	// TODO: criar um aquivo de propriedade

	// TODO: copiar o valor do consifg map e criar no os para executar

	// TODO: chamar o update passando os dados

	// TODO: chamar o tag passando os dados

	// TODO: chamar o processe que salva no s3

	return ctrl.Result{}, nil
}

// readConfigMap
func (r *LiquiBaseReleaseReconciler) readConfigMap() (corev1.ConfigMap, error) {

	return corev1.ConfigMap{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *LiquiBaseReleaseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&liquibasev1beta1.LiquiBaseRelease{}).
		Complete(r)
}
