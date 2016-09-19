package main

import (
	"fmt"
	"regexp"
)

const dns1035LabelFmt string = "[a-z]([-a-z0-9]*[a-z0-9])?"
const DNS1035LabelMaxLength int = 63

var dns1035LabelRegexp = regexp.MustCompile("^" + dns1035LabelFmt + "$")

const dns1123LabelFmt string = "[a-z0-9]([-a-z0-9]*[a-z0-9])?"
const DNS1123LabelMaxLength int = 63

var dns1123LabelRegexp = regexp.MustCompile("^" + dns1123LabelFmt + "$")

const dns1123SubdomainFmt string = dns1123LabelFmt + "(\\." + dns1123LabelFmt + ")*"
const DNS1123SubdomainMaxLength int = 253

var dns1123SubdomainRegexp = regexp.MustCompile("^" + dns1123SubdomainFmt + "$")

func main() {
	fmt.Printf("PodName, DeploymentName, LimitRangeName, ResourceQuotaName, SecretsName, EndpointsName:\n%s\n", dns1123SubdomainRegexp.String())
	fmt.Printf("ServiceName:\n%s\n", dns1035LabelRegexp.String())
	fmt.Printf("NamespaceName, ContainerName, ServicePortName, EndpointsAddressName:\n%s\n", dns1123LabelRegexp.String())
}
