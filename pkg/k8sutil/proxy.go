package k8sutil

import (
	"time"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func WaitForKotsadm(clientset *kubernetes.Clientset, namespace string, timeoutWaitingForWeb time.Duration) (string, error) {
	start := time.Now()

	for {
		// todo, find service, not pod
		pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{LabelSelector: "app=kotsadm"})
		if err != nil {
			return "", errors.Wrap(err, "failed to list pods")
		}

		for _, pod := range pods.Items {
			if pod.Status.Phase == corev1.PodRunning {
				if pod.Status.ContainerStatuses[0].Ready == true {
					return pod.Name, nil
				}
			}
		}

		time.Sleep(time.Second)

		if time.Now().Sub(start) > timeoutWaitingForWeb {
			return "", errors.New("timeout waiting for kotsadm pod")
		}
	}
}
