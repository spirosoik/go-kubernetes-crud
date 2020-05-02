package kubecrud

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // it is needed for k8 auth
)

// Create creates a namespace in k8
func (k *KubeService) Create(ctx context.Context, name string) error {
	ns := &v1.Namespace{ObjectMeta: metav1.ObjectMeta{
		Name: name,
	}}
	createOptions := metav1.CreateOptions{}
	_, err := k.ClientSet.CoreV1().Namespaces().Create(ctx, ns, createOptions)
	if err != nil {
		return err
	}
	return nil
}

// Get return details of the provided namespace
func (k *KubeService) Get(ctx context.Context, name string) (*v1.Namespace, error) {
	getOptions := metav1.GetOptions{}
	ns, err := k.ClientSet.CoreV1().Namespaces().Get(ctx, name, getOptions)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

// Update updates the provided namespace
func (k *KubeService) Update(ctx context.Context, namespace string, labelsToUpdate map[string]string) error {
	if len(labelsToUpdate) == 0 {
		return nil
	}
	ns, err := k.Get(ctx, namespace)
	if err != nil {
		return err
	}
	labels := overrideAttributes(ns.Labels, labelsToUpdate)
	nsMeta := &v1.Namespace{ObjectMeta: metav1.ObjectMeta{
		Name:   namespace,
		Labels: labels,
	}}
	updateOptions := metav1.UpdateOptions{}
	_, err = k.ClientSet.CoreV1().Namespaces().Update(ctx, nsMeta, updateOptions)
	if err != nil {
		return err
	}

	if len(labelsToUpdate) != 0 {
		var data string
		for k, v := range labels {
			data += fmt.Sprintf("%s=%s, ", k, v)
		}
	}
	return nil
}

// Delete deletes a namespace
func (k *KubeService) Delete(ctx context.Context, name string) error {
	deleteOptions := metav1.DeleteOptions{}
	err := k.ClientSet.CoreV1().Namespaces().Delete(ctx, name, deleteOptions)
	if err != nil {
		return err
	}
	return nil
}

// Exist returns true if the namespace exists in k8
func (k *KubeService) Exist(ctx context.Context, name string) (bool, error) {
	namespace, err := k.Get(ctx, name)
	if err != nil {
		return false, err
	}
	status := namespace.Status.Phase
	if namespace.Status.Phase != v1.NamespaceActive {
		return false, fmt.Errorf("Namespace status is not yet ready: %s", status)
	}
	return true, nil
}

// GetPods returns a pods list
func (k *KubeService) GetPods(ctx context.Context, namespace string) (*v1.PodList, error) {
	listOptions := metav1.ListOptions{}
	pods, err := k.ClientSet.CoreV1().Pods(namespace).List(ctx, listOptions)
	if err != nil {
		return nil, err
	}

	return pods, nil
}

// GetEndpoints returns an endpoints list
func (k *KubeService) GetEndpoints(ctx context.Context, namespace string) (*v1.EndpointsList, error) {
	listOptions := metav1.ListOptions{}
	endpoints, err := k.ClientSet.CoreV1().Endpoints(namespace).List(ctx, listOptions)
	if err != nil {
		return nil, err
	}

	return endpoints, nil
}

func overrideAttributes(currentAttributes, attributesToUpdate map[string]string) map[string]string {
	attributes := currentAttributes
	if len(attributes) == 0 {
		attributes = attributesToUpdate
	} else {
		for k, v := range attributesToUpdate {
			attributes[k] = v
		}
	}
	return attributes
}
