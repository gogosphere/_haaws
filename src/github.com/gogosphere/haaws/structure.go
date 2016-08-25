package haaws

type ELBStruct struct {
	LoadBalancerDescriptions []struct {
		AvailabilityZones         []string
		BackendServerDescriptions []interface{}
		CanonicalHostedZoneName   string
		CanonicalHostedZoneNameID string
		CreatedTime               string
		DNSName                   string
		HealthCheck               struct {
			HealthyThreshold   float64
			Interval           float64
			Target             string
			Timeout            float64
			UnhealthyThreshold float64
		}
		Instances []struct {
			InstanceId string
		}
		ListenerDescriptions []struct {
			Listener struct {
				InstancePort     float64
				InstanceProtocol string
				LoadBalancerPort float64
				Protocol         string
			}
			PolicyNames []interface{}
		}
		LoadBalancerName string
		Policies         struct {
			AppCookieStickinessPolicies []interface{}
			LBCookieStickinessPolicies  []interface{}
			OtherPolicies               []interface{}
		}
		Scheme              string
		SecurityGroups      []string
		SourceSecurityGroup struct {
			GroupName  string
			OwnerAlias string
		}
		Subnets []string
		VPCId   string
	}
}

type SubnetStruct struct {
	Subnets []struct {
		AvailabilityZone        string
		AvailableIpAddressCount float64
		CidrBlock               string
		DefaultForAz            bool
		MapPublicIpOnLaunch     bool
		State                   string
		SubnetId                string
		VpcId                   string
	}
}

type TagsStruct struct {
	Tags []struct {
		Key          string
		ResourceId   string
		ResourceType string
		Value        string
	}
}

type InstanceStruct struct {
	Reservations []struct {
		Groups    []interface{}
		Instances []struct {
			AmiLaunchIndex      float64
			Architecture        string
			BlockDeviceMappings []struct {
				DeviceName string
				Ebs        struct {
					AttachTime          string
					DeleteOnTermination bool
					Status              string
					VolumeId            string
				}
			}
			ClientToken        string
			EbsOptimized       bool
			Hypervisor         string
			IamInstanceProfile struct {
				Arn string
				ID  string `json:"Id,omitempty"`
			}
			ImageId      string
			InstanceId   string
			InstanceType string
			KeyName      string
			LaunchTime   string
			Monitoring   struct {
				State string
			}
			NetworkInterfaces []struct {
				Attachment struct {
					AttachTime          string
					AttachmentId        string
					DeleteOnTermination bool
					DeviceIndex         float64
					Status              string
				}
				Description string
				Groups      []struct {
					GroupId   string
					GroupName string
				}
				MacAddress         string
				NetworkInterfaceId string
				OwnerId            string
				PrivateDnsName     string
				PrivateIpAddress   string
				PrivateIpAddresses []struct {
					Primary          bool
					PrivateDnsName   string
					PrivateIpAddress string
				}
				SourceDestCheck bool
				Status          string
				SubnetId        string
				VpcId           string
			}
			Placement struct {
				AvailabilityZone string
				GroupName        string
				Tenancy          string
			}
			PrivateDnsName   string
			PrivateIpAddress string
			ProductCodes     []interface{}
			PublicDnsName    string
			RootDeviceName   string
			RootDeviceType   string
			SecurityGroups   []struct {
				GroupId   string
				GroupName string
			}
			SourceDestCheck bool
			State           struct {
				Code float64
				Name string
			}
			StateReason struct {
				Code    string
				Message string
			}
			StateTransitionReason string
			SubnetId              string
			Tags                  []struct {
				Key   string
				Value string
			}
			VirtualizationType string
			VpcId              string
		}
		OwnerId       string
		ReservationId string
	}
}
