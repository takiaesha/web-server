package e2e_test

import (
	"context"
	"flag"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	tappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/utils/pointer"
	"kmodules.xyz/client-go/tools/portforward"
	"path/filepath"
	"testing"
	"time"
)

func TestWebServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WebServer Suite")

}

var dep tappsv1.DeploymentInterface
var tunnel *portforward.Tunnel

var _ = BeforeSuite(func() {

	var kubeconfig *string

	ctx := context.Background()

	home := homedir.HomeDir()

	kubeconfig = flag.String("Kubeconfig", filepath.Join(home, ".kube", "config"), "directory ")

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		config, err = rest.InClusterConfig()
		if err != nil {
			return
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	dep = clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	deployment := &appsv1.Deployment{
		ObjectMeta: v1.ObjectMeta{
			Name: "apiserver-deploy",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32Ptr(3),
			Selector: &v1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "api",
							Image: "takia111/new-image",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 8080,
								},
							},
							Args: []string{
								"server",
								"-p",
								"8080",
							},
						},
					},
				},
			},
		},
	}

	By("Deployment creation: ")
	_, err = dep.Create(ctx, deployment, v1.CreateOptions{})
	Expect(err).NotTo(HaveOccurred())
	time.Sleep(20 * time.Second)

	tunnel = portforward.NewTunnel(portforward.TunnelOptions{
		Config:    config,
		Client:    clientset.CoreV1().RESTClient(),
		Resource:  "deployments",
		Name:      "apiserver-deploy",
		Namespace: "default",
		Remote:    8080,
	})
	err = tunnel.ForwardPort()
	Expect(err).NotTo(HaveOccurred())
})

// /delete deployment
var _ = AfterSuite(func() {
	///delete tunnel
	tunnel.Close()

	By("Delete deployment: ")

	dlt := v1.DeletePropagationForeground
	//Expect(dlt).NotTo(HaveOccurred())  //its a constant
	if err := dep.Delete(context.Background(), "apiserver-deploy", v1.DeleteOptions{
		PropagationPolicy: &dlt,
	}); err != nil {
		return
	}
})
