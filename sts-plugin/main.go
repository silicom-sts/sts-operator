// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2020-2021 Intel Corporation

package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	stsv1alpha1 "github.com/silicom-sts/sts-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
{"class":"POLL","time":"2010-06-04T10:31:00.289Z","active":1,
    "tpv":[{"class":"TPV","device":"/dev/ttyUSB0",
            "time":"2010-09-08T13:33:06.095Z",
            "ept":0.005,"lat":40.035093060,
            "lon":-75.519748733,"track":99.4319,"speed":0.123,"mode":2}],
*/
type GPSStatusRsp struct {
	Tpvs   []TPV  `json:"tpv"`
	Time   string `json:"time"`
	Active int    `json:"active"`
	Class  string `json:"class"`
}

type TPV struct {
	Time   string  `json:"time"`
	Lat    float32 `json:"lat"`
	Lon    float32 `json:"lon"`
	Device string  `json:"device"`
	Mode   int     `json:"mode"`
}

type GpsVersionRsp struct {
	Class      string `json:"class"`
	Release    string `json:"release"`
	Rev        string `json:"rev"`
	ProtoMajor int    `json:"proto_major"`
	ProtoMinor int    `json:"proto_minor"`
}

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(stsv1alpha1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func query_host(stsNode *stsv1alpha1.StsNode) {
	var idx int

	out, err := exec.Command("lspci", "-n", "-d", "8086:1591").Output()
	if err != nil {
		fmt.Errorf("error with lspci %v", err)
		return
	}

	stsNode.Status.EthInterfaces = []stsv1alpha1.StsNodeInterfaceStatus{}

	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		var nodeInterface stsv1alpha1.StsNodeInterfaceStatus

		id := strings.Split(scanner.Text(), " ")
		path := fmt.Sprintf("/sys/bus/pci/devices/0000:%s/net/*", id[0])
		res, _ := filepath.Glob(path)

		if len(res) < 1 {
			stsNode.Status.DriverAvailable = false
			break
		}

		stsNode.Status.DriverAvailable = true
		iface := string(filepath.Base(res[0]))

		nodeInterface.EthName = iface
		nodeInterface.PciAddr = id[0]
		nodeInterface.EthPort = idx
		idx = idx + 1

		out, err := exec.Command("ip", "link", "show", "dev", iface).Output()
		if err != nil {
			fmt.Printf("Error with ip link %v\n", err)
			return
		}

		if strings.Contains(string(out), "state UP") {
			nodeInterface.Status = "up"
		} else {
			nodeInterface.Status = "down"
		}

		stsNode.Status.EthInterfaces = append(stsNode.Status.EthInterfaces, nodeInterface)
	}
}

func main() {
	var node corev1.Node
	stsNode := &stsv1alpha1.StsNode{}
	nodeName := os.Getenv("NODE_NAME")
	namespace := os.Getenv("NAMESPACE")

	stsNode.Name = nodeName
	stsNode.Namespace = namespace

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	k8sClient, err := client.New(config, client.Options{Scheme: scheme})
	if err != nil {
		panic(err.Error())
	}

	for {

		err = k8sClient.Get(context.Background(),
			client.ObjectKey{
				Namespace: namespace,
				Name:      nodeName,
			}, stsNode)
		if err != nil {
			if err = k8sClient.Create(context.TODO(), stsNode); err != nil {
				panic(err.Error())
			}
		}

		err = k8sClient.Get(context.TODO(),
			client.ObjectKey{
				Name: os.Getenv("NODE_NAME"),
			}, &node)
		if err != nil {
			panic(err.Error())
		}

		query_host(stsNode)

		if stsNode.Status.DriverAvailable {
			node.Labels["sts.silicom.com/ice-driver-available"] = "true"
		} else {
			node.Labels["sts.silicom.com/ice-driver-available"] = "false"
		}

		err = k8sClient.Update(context.TODO(), &node)
		if err != nil {
			panic(err.Error())
		}

		if err := k8sClient.Status().Update(context.TODO(), stsNode); err != nil {
			fmt.Printf("Update failed: %v\n", err)
		}

		time.Sleep(30 * time.Second)
	}
}
