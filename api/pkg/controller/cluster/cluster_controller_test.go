package cluster

import (
	"log"
	"time"

	"github.com/kubermatic/kubermatic/api/pkg/cluster/client"
	kubermaticfakeclientset "github.com/kubermatic/kubermatic/api/pkg/crd/client/clientset/versioned/fake"
	kubermaticinformers "github.com/kubermatic/kubermatic/api/pkg/crd/client/informers/externalversions"
	"github.com/kubermatic/kubermatic/api/pkg/provider"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	kubefake "k8s.io/client-go/kubernetes/fake"

	ctrlruntimefakeinformer "sigs.k8s.io/controller-runtime/pkg/cache/informertest"
	ctrlruntimefakeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

const TestClusterName = "fqpcvnc6v"
const TestDC = "europe-west3-c"
const TestExternalURL = "dev.kubermatic.io"
const TestExternalPort = 30000

func newTestController(kubeObjects []runtime.Object, kubermaticObjects []runtime.Object) *Controller {
	dcs := buildDatacenterMeta()

	kubeClient := kubefake.NewSimpleClientset(kubeObjects...)
	kubermaticClient := kubermaticfakeclientset.NewSimpleClientset(kubermaticObjects...)

	kubeInformerFactory := informers.NewSharedInformerFactory(kubeClient, time.Minute*5)
	kubermaticInformerFactory := kubermaticinformers.NewSharedInformerFactory(kubermaticClient, time.Minute*5)

	dynamicClient := ctrlruntimefakeclient.NewFakeClient()
	controller, err := NewController(
		kubeClient,
		dynamicClient,
		kubermaticClient,
		TestExternalURL,
		TestDC,
		dcs,
		client.New(kubeInformerFactory.Core().V1().Secrets().Lister()),
		"",
		"",
		"192.0.2.0/24",
		"5Gi",
		"",
		"",
		false,
		false,
		"",
		[]byte{},

		&ctrlruntimefakeinformer.FakeInformers{},
		kubermaticInformerFactory.Kubermatic().V1().Clusters(),
		kubeInformerFactory.Core().V1().Namespaces(),
		kubeInformerFactory.Core().V1().Secrets(),
		kubeInformerFactory.Core().V1().Services(),
		kubeInformerFactory.Core().V1().PersistentVolumeClaims(),
		kubeInformerFactory.Core().V1().ConfigMaps(),
		kubeInformerFactory.Core().V1().ServiceAccounts(),
		kubeInformerFactory.Apps().V1().Deployments(),
		kubeInformerFactory.Apps().V1().StatefulSets(),
		kubeInformerFactory.Batch().V1beta1().CronJobs(),
		kubeInformerFactory.Extensions().V1beta1().Ingresses(),
		kubeInformerFactory.Rbac().V1().Roles(),
		kubeInformerFactory.Rbac().V1().RoleBindings(),
		kubeInformerFactory.Rbac().V1().ClusterRoleBindings(),
		kubeInformerFactory.Policy().V1beta1().PodDisruptionBudgets(),
		"",
		"",
		"",
		true,
		false,
	)
	if err != nil {
		log.Fatal(err)
	}

	kubeInformerFactory.Start(wait.NeverStop)
	kubermaticInformerFactory.Start(wait.NeverStop)

	kubeInformerFactory.WaitForCacheSync(wait.NeverStop)
	kubermaticInformerFactory.WaitForCacheSync(wait.NeverStop)

	return controller
}

func buildDatacenterMeta() map[string]provider.DatacenterMeta {
	seedAlias := "alias-europe-west3-c"
	return map[string]provider.DatacenterMeta{
		"us-central1": {
			Location: "us-central",
			Country:  "US",
			Private:  false,
			IsSeed:   true,
			Spec: provider.DatacenterSpec{
				Digitalocean: &provider.DigitaloceanSpec{
					Region: "ams2",
				},
			},
		},
		"us-central1-byo": {
			Location: "us-central",
			Country:  "US",
			Private:  false,
			Seed:     "us-central1",
			Spec: provider.DatacenterSpec{
				BringYourOwn: &provider.BringYourOwnSpec{},
			},
		},
		"private-do1": {
			Location: "US ",
			Seed:     "us-central1",
			Country:  "NL",
			Private:  true,
			Spec: provider.DatacenterSpec{
				Digitalocean: &provider.DigitaloceanSpec{
					Region: "ams2",
				},
			},
		},
		"regular-do1": {
			Location: "Amsterdam",
			Seed:     "us-central1",
			Country:  "NL",
			Spec: provider.DatacenterSpec{
				Digitalocean: &provider.DigitaloceanSpec{
					Region: "ams2",
				},
			},
		},
		"dns-override-do2": {
			Location:         "Amsterdam",
			Seed:             "us-central1",
			Country:          "NL",
			SeedDNSOverwrite: &seedAlias,
			Spec: provider.DatacenterSpec{
				Digitalocean: &provider.DigitaloceanSpec{
					Region: "ams3",
				},
			},
		},
	}
}
