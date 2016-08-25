package haaws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func DescribeAll() {
	DescribeLoadBalancers()
	DescribeSubnets()
	DescribeTags()
	DescribeInstances()
}

func DescribeLoadBalancers() {

	//avals := make(map[string][]string)
	uri := "describe-load-balancers.json"
	results := getdata(uri)
	var r ELBStruct
	json.Unmarshal(results, &r)

	elbwithinsts := make(map[string][]string)
	vpcwithelb := make(map[string][]string)

	for _, v := range r.LoadBalancerDescriptions {

		vpcwithelb[v.VPCId] = append(vpcwithelb[v.VPCId], v.DNSName)

		for _, vv := range v.Instances {
			elbwithinsts[v.DNSName] = append(elbwithinsts[v.DNSName], vv.InstanceId)
		}
	}
}

func DescribeSubnets() {

	uri := "describe-subnets.json"
	results := getdata(uri)
	var s SubnetStruct
	json.Unmarshal(results, &s)
	vpctoaz := make(map[string][]string)
	vpctosub := make(map[string][]string)
	vpctoinst := make(map[string][]string)
	for _, v := range s.Subnets {

		vpctoaz[v.VpcId] = append(vpctoaz[v.VpcId], v.AvailabilityZone)
		vpctosub[v.VpcId] = append(vpctosub[v.VpcId], v.SubnetId)
		vpctoinst[v.VpcId] = append(vpctoinst[v.VpcId], v.CidrBlock)
	}
}

func DescribeTags() {

	uri := "describe-tags.json"
	results := getdata(uri)
	var s TagsStruct
	unique := make(map[string][]string)
	json.Unmarshal(results, &s)
	for k, v := range s.Tags {
		_ = k
		unique[v.ResourceType] = append(unique[v.ResourceType], v.ResourceId)
	}
	for _, v := range unique {
		fmt.Println(v)
	}

}

func DescribeInstances() {
	uri := "describe-instances.json"
	results := getdata(uri)
	var s InstanceStruct
	//unique := make(map[string][]string)
	json.Unmarshal(results, &s)
	for k, v := range s.Reservations {
		fmt.Println(k, v.Instances[0].InstanceId, v.ReservationId, v.OwnerId, v.Groups)
	}

}

func getdata(uri string) []byte {
	var jsonStr []byte

	url := "http://localhost:8080/" + uri
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal("Die: ", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return body
}
