package repository

import (
	"strings"
	"testing"

	"github.com/cloudskiff/driftctl/pkg/remote/cache"
	"github.com/stretchr/testify/mock"

	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/r3labs/diff/v2"
	"github.com/stretchr/testify/assert"
)

func Test_ec2Repository_ListAllImages(t *testing.T) {

	tests := []struct {
		name    string
		mocks   func(client *MockEC2Client)
		want    []*ec2.Image
		wantErr error
	}{
		{
			name: "List all images",
			mocks: func(client *MockEC2Client) {
				client.On("DescribeImages",
					&ec2.DescribeImagesInput{
						Owners: []*string{
							aws.String("self"),
						},
					}).Return(&ec2.DescribeImagesOutput{
					Images: []*ec2.Image{
						{ImageId: aws.String("1")},
						{ImageId: aws.String("2")},
						{ImageId: aws.String("3")},
						{ImageId: aws.String("4")},
					},
				}, nil).Once()
			},
			want: []*ec2.Image{
				{ImageId: aws.String("1")},
				{ImageId: aws.String("2")},
				{ImageId: aws.String("3")},
				{ImageId: aws.String("4")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := cache.New(1)
			client := &MockEC2Client{}
			tt.mocks(client)
			r := &ec2Repository{
				client: client,
				cache:  store,
			}
			got, err := r.ListAllImages()
			assert.Equal(t, tt.wantErr, err)

			if err == nil {
				// Check that results were cached
				cachedData, err := r.ListAllImages()
				assert.NoError(t, err)
				assert.Equal(t, got, cachedData)
				assert.IsType(t, []*ec2.Image{}, store.Get("ec2ListAllImages"))
			}

			changelog, err := diff.Diff(got, tt.want)
			assert.Nil(t, err)
			if len(changelog) > 0 {
				for _, change := range changelog {
					t.Errorf("%s: %s -> %s", strings.Join(change.Path, "."), change.From, change.To)
				}
				t.Fail()
			}
		})
	}
}

func Test_ec2Repository_ListAllSnapshots(t *testing.T) {
	tests := []struct {
		name    string
		mocks   func(client *MockEC2Client)
		want    []*ec2.Snapshot
		wantErr error
	}{
		{name: "List with 2 pages",
			mocks: func(client *MockEC2Client) {
				client.On("DescribeSnapshotsPages",
					&ec2.DescribeSnapshotsInput{
						OwnerIds: []*string{
							aws.String("self"),
						},
					},
					mock.MatchedBy(func(callback func(res *ec2.DescribeSnapshotsOutput, lastPage bool) bool) bool {
						callback(&ec2.DescribeSnapshotsOutput{
							Snapshots: []*ec2.Snapshot{
								{VolumeId: aws.String("1")},
								{VolumeId: aws.String("2")},
								{VolumeId: aws.String("3")},
								{VolumeId: aws.String("4")},
							},
						}, false)
						callback(&ec2.DescribeSnapshotsOutput{
							Snapshots: []*ec2.Snapshot{
								{VolumeId: aws.String("5")},
								{VolumeId: aws.String("6")},
								{VolumeId: aws.String("7")},
								{VolumeId: aws.String("8")},
							},
						}, true)
						return true
					})).Return(nil).Once()
			},
			want: []*ec2.Snapshot{
				{VolumeId: aws.String("1")},
				{VolumeId: aws.String("2")},
				{VolumeId: aws.String("3")},
				{VolumeId: aws.String("4")},
				{VolumeId: aws.String("5")},
				{VolumeId: aws.String("6")},
				{VolumeId: aws.String("7")},
				{VolumeId: aws.String("8")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := cache.New(1)
			client := &MockEC2Client{}
			tt.mocks(client)
			r := &ec2Repository{
				client: client,
				cache:  store,
			}
			got, err := r.ListAllSnapshots()
			assert.Equal(t, tt.wantErr, err)

			if err == nil {
				// Check that results were cached
				cachedData, err := r.ListAllSnapshots()
				assert.NoError(t, err)
				assert.Equal(t, got, cachedData)
				assert.IsType(t, []*ec2.Snapshot{}, store.Get("ec2ListAllSnapshots"))
			}

			changelog, err := diff.Diff(got, tt.want)
			assert.Nil(t, err)
			if len(changelog) > 0 {
				for _, change := range changelog {
					t.Errorf("%s: %s -> %s", strings.Join(change.Path, "."), change.From, change.To)
				}
				t.Fail()
			}
		})
	}
}

func Test_ec2Repository_ListAllVolumes(t *testing.T) {
	tests := []struct {
		name    string
		mocks   func(client *MockEC2Client)
		want    []*ec2.Volume
		wantErr error
	}{
		{name: "List with 2 pages",
			mocks: func(client *MockEC2Client) {
				client.On("DescribeVolumesPages",
					&ec2.DescribeVolumesInput{},
					mock.MatchedBy(func(callback func(res *ec2.DescribeVolumesOutput, lastPage bool) bool) bool {
						callback(&ec2.DescribeVolumesOutput{
							Volumes: []*ec2.Volume{
								{VolumeId: aws.String("1")},
								{VolumeId: aws.String("2")},
								{VolumeId: aws.String("3")},
								{VolumeId: aws.String("4")},
							},
						}, false)
						callback(&ec2.DescribeVolumesOutput{
							Volumes: []*ec2.Volume{
								{VolumeId: aws.String("5")},
								{VolumeId: aws.String("6")},
								{VolumeId: aws.String("7")},
								{VolumeId: aws.String("8")},
							},
						}, true)
						return true
					})).Return(nil).Once()
			},
			want: []*ec2.Volume{
				{VolumeId: aws.String("1")},
				{VolumeId: aws.String("2")},
				{VolumeId: aws.String("3")},
				{VolumeId: aws.String("4")},
				{VolumeId: aws.String("5")},
				{VolumeId: aws.String("6")},
				{VolumeId: aws.String("7")},
				{VolumeId: aws.String("8")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := cache.New(1)
			client := &MockEC2Client{}
			tt.mocks(client)
			r := &ec2Repository{
				client: client,
				cache:  store,
			}
			got, err := r.ListAllVolumes()
			assert.Equal(t, tt.wantErr, err)

			if err == nil {
				// Check that results were cached
				cachedData, err := r.ListAllVolumes()
				assert.NoError(t, err)
				assert.Equal(t, got, cachedData)
				assert.IsType(t, []*ec2.Volume{}, store.Get("ec2ListAllVolumes"))
			}

			changelog, err := diff.Diff(got, tt.want)
			assert.Nil(t, err)
			if len(changelog) > 0 {
				for _, change := range changelog {
					t.Errorf("%s: %s -> %s", strings.Join(change.Path, "."), change.From, change.To)
				}
				t.Fail()
			}
		})
	}
}

func Test_ec2Repository_ListAllAddresses(t *testing.T) {
	tests := []struct {
		name    string
		mocks   func(client *MockEC2Client)
		want    []*ec2.Address
		wantErr error
	}{
		{
			name: "List address",
			mocks: func(client *MockEC2Client) {
				client.On("DescribeAddresses", &ec2.DescribeAddressesInput{}).
					Return(&ec2.DescribeAddressesOutput{
						Addresses: []*ec2.Address{
							{AssociationId: aws.String("1")},
							{AssociationId: aws.String("2")},
							{AssociationId: aws.String("3")},
							{AssociationId: aws.String("4")},
						},
					}, nil).Once()
			},
			want: []*ec2.Address{
				{AssociationId: aws.String("1")},
				{AssociationId: aws.String("2")},
				{AssociationId: aws.String("3")},
				{AssociationId: aws.String("4")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := cache.New(1)
			client := &MockEC2Client{}
			tt.mocks(client)
			r := &ec2Repository{
				client: client,
				cache:  store,
			}
			got, err := r.ListAllAddresses()
			assert.Equal(t, tt.wantErr, err)

			if err == nil {
				// Check that results were cached
				cachedData, err := r.ListAllAddresses()
				assert.NoError(t, err)
				assert.Equal(t, got, cachedData)
				assert.IsType(t, []*ec2.Address{}, store.Get("ec2ListAllAddresses"))
			}

			changelog, err := diff.Diff(got, tt.want)
			assert.Nil(t, err)
			if len(changelog) > 0 {
				for _, change := range changelog {
					t.Errorf("%s: %s -> %s", strings.Join(change.Path, "."), change.From, change.To)
				}
				t.Fail()
			}
		})
	}
}

func Test_ec2Repository_ListAllAddressesAssociation(t *testing.T) {
	tests := []struct {
		name    string
		mocks   func(client *MockEC2Client)
		want    []string
		wantErr error
	}{
		{
			name: "List address",
			mocks: func(client *MockEC2Client) {
				client.On("DescribeAddresses", &ec2.DescribeAddressesInput{}).
					Return(&ec2.DescribeAddressesOutput{
						Addresses: []*ec2.Address{
							{AssociationId: aws.String("1")},
							{AssociationId: aws.String("2")},
							{AssociationId: aws.String("3")},
							{AssociationId: aws.String("4")},
						},
					}, nil).Once()
			},
			want: []string{
				"1",
				"2",
				"3",
				"4",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := cache.New(1)
			client := &MockEC2Client{}
			tt.mocks(client)
			r := &ec2Repository{
				client: client,
				cache:  store,
			}
			got, err := r.ListAllAddressesAssociation()
			assert.Equal(t, tt.wantErr, err)

			if err == nil {
				// Check that results were cached
				cachedData, err := r.ListAllAddressesAssociation()
				assert.NoError(t, err)
				assert.Equal(t, got, cachedData)
				assert.IsType(t, []string{}, store.Get("ec2ListAllAddressesAssociation"))
			}

			changelog, err := diff.Diff(got, tt.want)
			assert.Nil(t, err)
			if len(changelog) > 0 {
				for _, change := range changelog {
					t.Errorf("%s: %s -> %s", strings.Join(change.Path, "."), change.From, change.To)
				}
				t.Fail()
			}
		})
	}
}

func Test_ec2Repository_ListAllInstances(t *testing.T) {
	tests := []struct {
		name    string
		mocks   func(client *MockEC2Client)
		want    []*ec2.Instance
		wantErr error
	}{
		{name: "List with 2 pages",
			mocks: func(client *MockEC2Client) {
				client.On("DescribeInstancesPages",
					&ec2.DescribeInstancesInput{},
					mock.MatchedBy(func(callback func(res *ec2.DescribeInstancesOutput, lastPage bool) bool) bool {
						callback(&ec2.DescribeInstancesOutput{
							Reservations: []*ec2.Reservation{
								{
									Instances: []*ec2.Instance{
										{ImageId: aws.String("1")},
										{ImageId: aws.String("2")},
										{ImageId: aws.String("3")},
									},
								},
								{
									Instances: []*ec2.Instance{
										{ImageId: aws.String("4")},
										{ImageId: aws.String("5")},
										{ImageId: aws.String("6")},
									},
								},
							},
						}, false)
						callback(&ec2.DescribeInstancesOutput{
							Reservations: []*ec2.Reservation{
								{
									Instances: []*ec2.Instance{
										{ImageId: aws.String("7")},
										{ImageId: aws.String("8")},
										{ImageId: aws.String("9")},
									},
								},
								{
									Instances: []*ec2.Instance{
										{ImageId: aws.String("10")},
										{ImageId: aws.String("11")},
										{ImageId: aws.String("12")},
									},
								},
							},
						}, true)
						return true
					})).Return(nil).Once()
			},
			want: []*ec2.Instance{
				{ImageId: aws.String("1")},
				{ImageId: aws.String("2")},
				{ImageId: aws.String("3")},
				{ImageId: aws.String("4")},
				{ImageId: aws.String("5")},
				{ImageId: aws.String("6")},
				{ImageId: aws.String("7")},
				{ImageId: aws.String("8")},
				{ImageId: aws.String("9")},
				{ImageId: aws.String("10")},
				{ImageId: aws.String("11")},
				{ImageId: aws.String("12")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := cache.New(1)
			client := &MockEC2Client{}
			tt.mocks(client)
			r := &ec2Repository{
				client: client,
				cache:  store,
			}
			got, err := r.ListAllInstances()
			assert.Equal(t, tt.wantErr, err)

			if err == nil {
				// Check that results were cached
				cachedData, err := r.ListAllInstances()
				assert.NoError(t, err)
				assert.Equal(t, got, cachedData)
				assert.IsType(t, []*ec2.Instance{}, store.Get("ec2ListAllInstances"))
			}
			changelog, err := diff.Diff(got, tt.want)
			assert.Nil(t, err)
			if len(changelog) > 0 {
				for _, change := range changelog {
					t.Errorf("%s: %s -> %s", strings.Join(change.Path, "."), change.From, change.To)
				}
				t.Fail()
			}
		})
	}
}

func Test_ec2Repository_ListAllKeyPairs(t *testing.T) {
	tests := []struct {
		name    string
		mocks   func(client *MockEC2Client)
		want    []*ec2.KeyPairInfo
		wantErr error
	}{
		{
			name: "List address",
			mocks: func(client *MockEC2Client) {
				client.On("DescribeKeyPairs", &ec2.DescribeKeyPairsInput{}).
					Return(&ec2.DescribeKeyPairsOutput{
						KeyPairs: []*ec2.KeyPairInfo{
							{KeyPairId: aws.String("1")},
							{KeyPairId: aws.String("2")},
							{KeyPairId: aws.String("3")},
							{KeyPairId: aws.String("4")},
						},
					}, nil).Once()
			},
			want: []*ec2.KeyPairInfo{
				{KeyPairId: aws.String("1")},
				{KeyPairId: aws.String("2")},
				{KeyPairId: aws.String("3")},
				{KeyPairId: aws.String("4")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := cache.New(1)
			client := &MockEC2Client{}
			tt.mocks(client)
			r := &ec2Repository{
				client: client,
				cache:  store,
			}
			got, err := r.ListAllKeyPairs()
			assert.Equal(t, tt.wantErr, err)

			if err == nil {
				// Check that results were cached
				cachedData, err := r.ListAllKeyPairs()
				assert.NoError(t, err)
				assert.Equal(t, got, cachedData)
				assert.IsType(t, []*ec2.KeyPairInfo{}, store.Get("ec2ListAllKeyPairs"))
			}

			changelog, err := diff.Diff(got, tt.want)
			assert.Nil(t, err)
			if len(changelog) > 0 {
				for _, change := range changelog {
					t.Errorf("%s: %s -> %s", strings.Join(change.Path, "."), change.From, change.To)
				}
				t.Fail()
			}
		})
	}
}

func Test_ec2Repository_ListAllInternetGateways(t *testing.T) {
	tests := []struct {
		name    string
		mocks   func(client *MockEC2Client)
		want    []*ec2.InternetGateway
		wantErr error
	}{
		{
			name: "List only gateways with multiple pages",
			mocks: func(client *MockEC2Client) {
				client.On("DescribeInternetGatewaysPages",
					&ec2.DescribeInternetGatewaysInput{},
					mock.MatchedBy(func(callback func(res *ec2.DescribeInternetGatewaysOutput, lastPage bool) bool) bool {
						callback(&ec2.DescribeInternetGatewaysOutput{
							InternetGateways: []*ec2.InternetGateway{
								{
									InternetGatewayId: aws.String("Internet-0"),
								},
								{
									InternetGatewayId: aws.String("Internet-1"),
								},
							},
						}, false)
						callback(&ec2.DescribeInternetGatewaysOutput{
							InternetGateways: []*ec2.InternetGateway{
								{
									InternetGatewayId: aws.String("Internet-2"),
								},
								{
									InternetGatewayId: aws.String("Internet-3"),
								},
							},
						}, true)
						return true
					})).Return(nil)
			},
			want: []*ec2.InternetGateway{
				{
					InternetGatewayId: aws.String("Internet-0"),
				},
				{
					InternetGatewayId: aws.String("Internet-1"),
				},
				{
					InternetGatewayId: aws.String("Internet-2"),
				},
				{
					InternetGatewayId: aws.String("Internet-3"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &MockEC2Client{}
			tt.mocks(client)
			r := &ec2Repository{
				client: client,
			}
			got, err := r.ListAllInternetGateways()
			assert.Equal(t, tt.wantErr, err)
			changelog, err := diff.Diff(got, tt.want)
			assert.Nil(t, err)
			if len(changelog) > 0 {
				for _, change := range changelog {
					t.Errorf("%s: %s -> %s", strings.Join(change.Path, "."), change.From, change.To)
				}
				t.Fail()
			}
		})
	}
}

func Test_ec2Repository_ListAllSubnets(t *testing.T) {
	tests := []struct {
		name              string
		mocks             func(client *MockEC2Client)
		wantSubnet        []*ec2.Subnet
		wantDefaultSubnet []*ec2.Subnet
		wantErr           error
	}{
		{
			name: "List with 2 pages",
			mocks: func(client *MockEC2Client) {
				client.On("DescribeSubnetsPages",
					&ec2.DescribeSubnetsInput{},
					mock.MatchedBy(func(callback func(res *ec2.DescribeSubnetsOutput, lastPage bool) bool) bool {
						callback(&ec2.DescribeSubnetsOutput{
							Subnets: []*ec2.Subnet{
								{
									SubnetId:     aws.String("subnet-0b13f1e0eacf67424"), // subnet2
									DefaultForAz: aws.Bool(false),
								},
								{
									SubnetId:     aws.String("subnet-0c9b78001fe186e22"), // subnet3
									DefaultForAz: aws.Bool(false),
								},
								{
									SubnetId:     aws.String("subnet-05810d3f933925f6d"), // subnet1
									DefaultForAz: aws.Bool(false),
								},
							},
						}, false)
						callback(&ec2.DescribeSubnetsOutput{
							Subnets: []*ec2.Subnet{
								{
									SubnetId:     aws.String("subnet-44fe0c65"), // us-east-1a
									DefaultForAz: aws.Bool(true),
								},
								{
									SubnetId:     aws.String("subnet-65e16628"), // us-east-1b
									DefaultForAz: aws.Bool(true),
								},
								{
									SubnetId:     aws.String("subnet-afa656f0"), // us-east-1c
									DefaultForAz: aws.Bool(true),
								},
							},
						}, true)
						return true
					})).Return(nil)
			},
			wantSubnet: []*ec2.Subnet{
				{
					SubnetId:     aws.String("subnet-0b13f1e0eacf67424"), // subnet2
					DefaultForAz: aws.Bool(false),
				},
				{
					SubnetId:     aws.String("subnet-0c9b78001fe186e22"), // subnet3
					DefaultForAz: aws.Bool(false),
				},
				{
					SubnetId:     aws.String("subnet-05810d3f933925f6d"), // subnet1
					DefaultForAz: aws.Bool(false),
				},
			},
			wantDefaultSubnet: []*ec2.Subnet{
				{
					SubnetId:     aws.String("subnet-44fe0c65"), // us-east-1a
					DefaultForAz: aws.Bool(true),
				},
				{
					SubnetId:     aws.String("subnet-65e16628"), // us-east-1b
					DefaultForAz: aws.Bool(true),
				},
				{
					SubnetId:     aws.String("subnet-afa656f0"), // us-east-1c
					DefaultForAz: aws.Bool(true),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &MockEC2Client{}
			tt.mocks(client)
			r := &ec2Repository{
				client: client,
			}
			gotSubnet, gotDefaultSubnet, err := r.ListAllSubnets()
			assert.Equal(t, tt.wantErr, err)
			changelog, err := diff.Diff(gotSubnet, tt.wantSubnet)
			assert.Nil(t, err)
			if len(changelog) > 0 {
				for _, change := range changelog {
					t.Errorf("%s: %s -> %s", strings.Join(change.Path, "."), change.From, change.To)
				}
				t.Fail()
			}
			changelog, err = diff.Diff(gotDefaultSubnet, tt.wantDefaultSubnet)
			assert.Nil(t, err)
			if len(changelog) > 0 {
				for _, change := range changelog {
					t.Errorf("%s: %s -> %s", strings.Join(change.Path, "."), change.From, change.To)
				}
				t.Fail()
			}
		})
	}
}

func Test_ec2Repository_ListAllNatGateways(t *testing.T) {
	tests := []struct {
		name    string
		mocks   func(client *MockEC2Client)
		want    []*ec2.NatGateway
		wantErr error
	}{
		{
			name: "List only gateways with multiple pages",
			mocks: func(client *MockEC2Client) {
				client.On("DescribeNatGatewaysPages",
					&ec2.DescribeNatGatewaysInput{},
					mock.MatchedBy(func(callback func(res *ec2.DescribeNatGatewaysOutput, lastPage bool) bool) bool {
						callback(&ec2.DescribeNatGatewaysOutput{
							NatGateways: []*ec2.NatGateway{
								{
									NatGatewayId: aws.String("nat-0"),
								},
								{
									NatGatewayId: aws.String("nat-1"),
								},
							},
						}, false)
						callback(&ec2.DescribeNatGatewaysOutput{
							NatGateways: []*ec2.NatGateway{
								{
									NatGatewayId: aws.String("nat-2"),
								},
								{
									NatGatewayId: aws.String("nat-3"),
								},
							},
						}, true)
						return true
					})).Return(nil)
			},
			want: []*ec2.NatGateway{
				{
					NatGatewayId: aws.String("nat-0"),
				},
				{
					NatGatewayId: aws.String("nat-1"),
				},
				{
					NatGatewayId: aws.String("nat-2"),
				},
				{
					NatGatewayId: aws.String("nat-3"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &MockEC2Client{}
			tt.mocks(client)
			r := &ec2Repository{
				client: client,
			}
			got, err := r.ListAllNatGateways()
			assert.Equal(t, tt.wantErr, err)
			changelog, err := diff.Diff(got, tt.want)
			assert.Nil(t, err)
			if len(changelog) > 0 {
				for _, change := range changelog {
					t.Errorf("%s: %s -> %s", strings.Join(change.Path, "."), change.From, change.To)
				}
				t.Fail()
			}
		})
	}
}
