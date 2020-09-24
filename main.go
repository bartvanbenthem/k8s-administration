package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/tabwriter"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Host struct {
	HostName  string
	Namespace string
}

type K8s struct{}

func main() {
	// Check if there are empty ENV Variables that need to be set
	CheckEmptyEnVar()
	// print hostnames per namespace
	PrintHostnames()
}

func CheckEmptyEnVar() {
	vars := []string{"K8S_KUBECONFIG"}

	for _, v := range vars {
		if os.Getenv(v) == "" {
			log.Fatalf("Fatal Error: env variable [ %v ] is empty\n", v)
		}
	}
}

func (k *K8s) GetCurrentContext() string {
	cmd := exec.Command("kubectl", "config", "current-context")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	return strings.TrimSuffix(string(stdoutStderr), "\n")
}

func (k *K8s) CreateClientSet() *kubernetes.Clientset {
	// When running the binary inside of a pod in a cluster,
	// the kubelet will automatically mount a service account into the container at:
	// /var/run/secrets/kubernetes.io/serviceaccount.
	// It replaces the kubeconfig file and is turned into a rest.Config via the rest.InClusterConfig() method
	config, err := rest.InClusterConfig()
	if err != nil {
		// fallback to kubeconfig
		kubeconfig := filepath.Join("~", ".kube", "config")
		if envvar := os.Getenv("K8S_KUBECONFIG"); len(envvar) > 0 {
			kubeconfig = envvar
		}
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			fmt.Printf("The kubeconfig cannot be loaded: %v\n", err)
			os.Exit(1)
		}
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	return clientset
}

func (k *K8s) GetHostname(clientset *kubernetes.Clientset) ([]Host, error) {
	var hosts []Host

	ns, err := clientset.CoreV1().Namespaces().List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, n := range ns.Items {
		ing, err := clientset.NetworkingV1beta1().Ingresses(n.GetName()).List(v1.ListOptions{})
		if err != nil {
			return nil, err
		}

		var host Host
		for _, i := range ing.Items {
			rules := i.Spec.Rules
			for _, r := range rules {
				host.HostName = r.Host
				host.Namespace = n.GetName()
				hosts = append(hosts, host)
			}
		}
	}

	return hosts, err
}

func PrintHostnames() {
	var kube K8s
	// Create the hostname output
	hosts, err := kube.GetHostname(kube.CreateClientSet())
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintln(writer, "HostName\tNamespace\tContext\t")
	fmt.Fprintln(writer, "--------\t---------\t-------\t")

	for _, h := range hosts {
		cluster := kube.GetCurrentContext()
		fmt.Fprintln(writer, fmt.Sprintf("%v\t%v\t%v\t", h.HostName, h.Namespace, cluster))
	}
	writer.Flush()
	fmt.Println()
}
