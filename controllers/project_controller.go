/*
Copyright 2021.

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

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	p6sv1alpha1 "github.com/logandavies181/p6s/api/v1alpha1"
)

// ProjectReconciler reconciles a Project object
type ProjectReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=p6s.logan.kiwi,resources=projects,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=p6s.logan.kiwi,resources=projects/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=p6s.logan.kiwi,resources=projects/finalizers,verbs=update

// Reconcile creates a namespace Project.Metadata.Name and resources as per the ProjectTemplate "Default" within the
// p6s-system namespace
func (r *ProjectReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues("project", req.NamespacedName)

	name := req.Name

	ns := &corev1.Namespace{}
	err := r.Get(ctx, req.NamespacedName, ns)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Creating namespace", "Namespace.Name", name)
		err = r.Create(ctx, ns)
		if err != nil {
			logger.Error(err, "Could not create namespace", "Namespace.Name", name)
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ProjectReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&p6sv1alpha1.Project{}).
		Owns(&p6sv1alpha1.Project{}).
		Complete(r)
}
