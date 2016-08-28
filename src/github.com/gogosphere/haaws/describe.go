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
	//DescribeInstances()
}

func DescribeLoadBalancers() map[string][]string {

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
	return elbwithinsts
}

func DescribeSubnets() map[int][]string {

	uri := "describe-subnets.json"
	results := getdata(uri)
	var s SubnetStruct
	json.Unmarshal(results, &s)

	vpctoaz := make(map[string][]string)
	vpctosub := make(map[string][]string)
	vpctoinst := make(map[string][]string)
	i := 0
	q := make(map[int][]string)
	for _, v := range s.Subnets {

		vpctoaz[v.VpcId] = append(vpctoaz[v.VpcId], v.AvailabilityZone)
		vpctosub[v.VpcId] = append(vpctosub[v.VpcId], v.SubnetId)
		vpctoinst[v.VpcId] = append(vpctoinst[v.VpcId], v.CidrBlock)

		q[i] = []string{v.AvailabilityZone, v.VpcId, v.SubnetId}
		i++
	}
	zz := DescribeInstances(q)
	return zz
}

func DescribeInstances(zz map[int][]string) map[int][]string {
	uri := "describe-instances.json"
	results := getdata(uri)
	var s InstanceStruct
	//instancetotag := make(map[string][]string)
	//instancetotagkey := make(map[string]map[string][]string)
	//subnetaskey := make(map[string]map[string][]string)
	appendinstancemap := make(map[string][]string)
	json.Unmarshal(results, &s)
	for _, v := range s.Reservations {
		//fmt.Println(k, v.Instances[0].InstanceId, v.ReservationId, v.OwnerId, v.Groups)
		appendinstancemap[v.Instances[0].SubnetId] = append(appendinstancemap[v.Instances[0].SubnetId],v.Instances[0].InstanceId)
	}

	for q, r := range zz {
		_ = q
		_ = r
		zz[q] = append(zz[q], appendinstancemap[r[2]]...)
	}

	/*
	for k, v := range subnetaskey {
		_ = v
		_ = k
		fmt.Println(k)
	}
	*/
	return zz
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
