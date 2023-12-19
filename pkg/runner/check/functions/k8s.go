package functions

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/client"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func jpKubernetesList(arguments []any) (any, error) {
	var client client.Client
	var apiVersion, kind, namespace string
	if err := getArg(arguments, 0, &client); err != nil {
		return nil, err
	}
	if err := getArg(arguments, 1, &apiVersion); err != nil {
		return nil, err
	}
	if err := getArg(arguments, 2, &kind); err != nil {
		return nil, err
	}
	if len(arguments) == 4 {
		if err := getArg(arguments, 3, &namespace); err != nil {
			return nil, err
		}
	}
	var list unstructured.UnstructuredList
	list.SetAPIVersion(apiVersion)
	list.SetKind(kind)
	var listOptions []ctrlclient.ListOption
	if namespace != "" {
		listOptions = append(listOptions, ctrlclient.InNamespace(namespace))
	}
	if err := client.List(context.TODO(), &list, listOptions...); err != nil {
		return nil, err
	}
	return list.Items, nil
}