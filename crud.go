package kubecrud

import (
	"os"

	"path/filepath"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

//KubeService includes the kube context
type KubeService struct {
	ClientSet *kubernetes.Clientset
}

// NewService factory method to create a new
// Kubernete sertvice to interact with crud methods
func NewService(kubecontext *string) (*KubeService, error) {
	var (
		config *rest.Config
		err error
	)
	if kubecontext == nil {
		config, err = rest.InClusterConfig()
	} else {
		config, err = clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			&clientcmd.ClientConfigLoadingRules{ExplicitPath: getKubeconfig()},
			&clientcmd.ConfigOverrides{
				CurrentContext: *kubecontext,
			}).ClientConfig()
	}

	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &KubeService{
		ClientSet: clientset,
	}, nil
}


func getKubeconfig() string {
	var kubeconfig string
	if kubeConfigPath := os.Getenv("KUBECONFIG"); kubeConfigPath != "" {
		kubeconfig = kubeConfigPath
	} else {
		kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "config")
	}
	return kubeconfig
}