package resourcecreator

import (
	"fmt"

	nais "github.com/nais/naiserator/pkg/apis/nais.io/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ServiceAccount(app *nais.Application, options ResourceOptions) *corev1.ServiceAccount {
	objectMeta := app.CreateObjectMeta()
	if len(options.GoogleProjectId) > 0 {
		objectMeta.Annotations["iam.gke.io/gcp-service-account"] = GcpServiceAccountName(app, options.GoogleProjectId)
	}

	return &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceAccount",
			APIVersion: "v1",
		},
		ObjectMeta: objectMeta,
	}
}

func GcpServiceAccountName(app *nais.Application, projectId string) string {
	return fmt.Sprintf("%s@%s.iam.gserviceaccount.com", app.CreateAppNamespaceHash(), projectId)
}
