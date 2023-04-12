package common

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"github.com/open-feature/open-feature-operator/pkg/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func CreateConfigMap(
	ctx context.Context, log logr.Logger, c client.Client, namespace string, name string, ownerReferences []metav1.OwnerReference,
) error {
	log.V(1).Info(fmt.Sprintf("Creating configmap %s", name))
	references := []metav1.OwnerReference{
		ownerReferences[0],
	}
	references[0].Controller = utils.FalseVal()
	ff := FeatureFlag(ctx, c, namespace, name)
	if ff.Name == "" {
		return fmt.Errorf("feature flag configuration %s/%s not found", namespace, name)
	}
	references = append(references, ff.GetReference())

	cm := ff.GenerateConfigMap(name, namespace, references)

	return c.Create(ctx, &cm)
}