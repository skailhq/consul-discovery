package consuldiscovery

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func getCurrentNodeName() string {
	hostName, _ := os.Hostname()
	return hostName
}

func TestHealth(t *testing.T) {
	Convey("HealthByNode", t, func() {
		client := getClient(t)
		checks, err := client.HealthByNode(getCurrentNodeName())
		So(err, ShouldEqual, nil)
		So(len(checks), ShouldEqual, 3)
	})

	Convey("HealthByState", t, func() {
		client := getClient(t)
		checks, err := client.HealthByState("critical")
		So(err, ShouldEqual, nil)
		So(len(checks), ShouldNotEqual, 0)
	})

	Convey("HealthByState", t, func() {
		client := getClient(t)
		checks, err := client.HealthByState("passing")
		So(err, ShouldEqual, nil)
		So(len(checks), ShouldNotEqual, 0)
	})

	Convey("HealthByService", t, func() {
		client := getClient(t)
		nodes, err := client.HealthByService("simple_service")
		So(err, ShouldEqual, nil)
		So(len(nodes), ShouldEqual, 1)
		node := nodes[0]
		So(node.Service.ServiceID, ShouldEqual, "simple_service")
		So(node.Service.ServiceName, ShouldEqual, "simple_service")
		So(len(node.Checks), ShouldEqual, 3)
	})
}
