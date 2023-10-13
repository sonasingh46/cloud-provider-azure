// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

// Code generated by MockGen. DO NOT EDIT.
// Source: sigs.k8s.io/cloud-provider-azure/pkg/azclient (interfaces: ClientFactory)

// Package mock_azclient is a generated GoMock package.
package mock_azclient

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	availabilitysetclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/availabilitysetclient"
	deploymentclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/deploymentclient"
	diskclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/diskclient"
	interfaceclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/interfaceclient"
	ipgroupclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/ipgroupclient"
	loadbalancerclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/loadbalancerclient"
	managedclusterclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/managedclusterclient"
	privateendpointclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/privateendpointclient"
	privatelinkserviceclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/privatelinkserviceclient"
	privatezoneclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/privatezoneclient"
	publicipaddressclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/publicipaddressclient"
	publicipprefixclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/publicipprefixclient"
	registryclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/registryclient"
	resourcegroupclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/resourcegroupclient"
	routetableclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/routetableclient"
	securitygroupclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/securitygroupclient"
	snapshotclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/snapshotclient"
	sshpublickeyresourceclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/sshpublickeyresourceclient"
	subnetclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/subnetclient"
	virtualmachineclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachineclient"
	virtualmachinescalesetclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachinescalesetclient"
	virtualmachinescalesetvmclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachinescalesetvmclient"
	virtualnetworkclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualnetworkclient"
)

// MockClientFactory is a mock of ClientFactory interface.
type MockClientFactory struct {
	ctrl     *gomock.Controller
	recorder *MockClientFactoryMockRecorder
}

// MockClientFactoryMockRecorder is the mock recorder for MockClientFactory.
type MockClientFactoryMockRecorder struct {
	mock *MockClientFactory
}

// NewMockClientFactory creates a new mock instance.
func NewMockClientFactory(ctrl *gomock.Controller) *MockClientFactory {
	mock := &MockClientFactory{ctrl: ctrl}
	mock.recorder = &MockClientFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientFactory) EXPECT() *MockClientFactoryMockRecorder {
	return m.recorder
}

// GetAvailabilitySetClient mocks base method.
func (m *MockClientFactory) GetAvailabilitySetClient() availabilitysetclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailabilitySetClient")
	ret0, _ := ret[0].(availabilitysetclient.Interface)
	return ret0
}

// GetAvailabilitySetClient indicates an expected call of GetAvailabilitySetClient.
func (mr *MockClientFactoryMockRecorder) GetAvailabilitySetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailabilitySetClient", reflect.TypeOf((*MockClientFactory)(nil).GetAvailabilitySetClient))
}

// GetDeploymentClient mocks base method.
func (m *MockClientFactory) GetDeploymentClient() deploymentclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeploymentClient")
	ret0, _ := ret[0].(deploymentclient.Interface)
	return ret0
}

// GetDeploymentClient indicates an expected call of GetDeploymentClient.
func (mr *MockClientFactoryMockRecorder) GetDeploymentClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentClient", reflect.TypeOf((*MockClientFactory)(nil).GetDeploymentClient))
}

// GetDiskClient mocks base method.
func (m *MockClientFactory) GetDiskClient() diskclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiskClient")
	ret0, _ := ret[0].(diskclient.Interface)
	return ret0
}

// GetDiskClient indicates an expected call of GetDiskClient.
func (mr *MockClientFactoryMockRecorder) GetDiskClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiskClient", reflect.TypeOf((*MockClientFactory)(nil).GetDiskClient))
}

// GetIPGroupClient mocks base method.
func (m *MockClientFactory) GetIPGroupClient() ipgroupclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIPGroupClient")
	ret0, _ := ret[0].(ipgroupclient.Interface)
	return ret0
}

// GetIPGroupClient indicates an expected call of GetIPGroupClient.
func (mr *MockClientFactoryMockRecorder) GetIPGroupClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPGroupClient", reflect.TypeOf((*MockClientFactory)(nil).GetIPGroupClient))
}

// GetInterfaceClient mocks base method.
func (m *MockClientFactory) GetInterfaceClient() interfaceclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInterfaceClient")
	ret0, _ := ret[0].(interfaceclient.Interface)
	return ret0
}

// GetInterfaceClient indicates an expected call of GetInterfaceClient.
func (mr *MockClientFactoryMockRecorder) GetInterfaceClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInterfaceClient", reflect.TypeOf((*MockClientFactory)(nil).GetInterfaceClient))
}

// GetLoadBalancerClient mocks base method.
func (m *MockClientFactory) GetLoadBalancerClient() loadbalancerclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoadBalancerClient")
	ret0, _ := ret[0].(loadbalancerclient.Interface)
	return ret0
}

// GetLoadBalancerClient indicates an expected call of GetLoadBalancerClient.
func (mr *MockClientFactoryMockRecorder) GetLoadBalancerClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoadBalancerClient", reflect.TypeOf((*MockClientFactory)(nil).GetLoadBalancerClient))
}

// GetManagedClusterClient mocks base method.
func (m *MockClientFactory) GetManagedClusterClient() managedclusterclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagedClusterClient")
	ret0, _ := ret[0].(managedclusterclient.Interface)
	return ret0
}

// GetManagedClusterClient indicates an expected call of GetManagedClusterClient.
func (mr *MockClientFactoryMockRecorder) GetManagedClusterClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedClusterClient", reflect.TypeOf((*MockClientFactory)(nil).GetManagedClusterClient))
}

// GetPrivateEndpointClient mocks base method.
func (m *MockClientFactory) GetPrivateEndpointClient() privateendpointclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateEndpointClient")
	ret0, _ := ret[0].(privateendpointclient.Interface)
	return ret0
}

// GetPrivateEndpointClient indicates an expected call of GetPrivateEndpointClient.
func (mr *MockClientFactoryMockRecorder) GetPrivateEndpointClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateEndpointClient", reflect.TypeOf((*MockClientFactory)(nil).GetPrivateEndpointClient))
}

// GetPrivateLinkServiceClient mocks base method.
func (m *MockClientFactory) GetPrivateLinkServiceClient() privatelinkserviceclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateLinkServiceClient")
	ret0, _ := ret[0].(privatelinkserviceclient.Interface)
	return ret0
}

// GetPrivateLinkServiceClient indicates an expected call of GetPrivateLinkServiceClient.
func (mr *MockClientFactoryMockRecorder) GetPrivateLinkServiceClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateLinkServiceClient", reflect.TypeOf((*MockClientFactory)(nil).GetPrivateLinkServiceClient))
}

// GetPrivateZoneClient mocks base method.
func (m *MockClientFactory) GetPrivateZoneClient() privatezoneclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateZoneClient")
	ret0, _ := ret[0].(privatezoneclient.Interface)
	return ret0
}

// GetPrivateZoneClient indicates an expected call of GetPrivateZoneClient.
func (mr *MockClientFactoryMockRecorder) GetPrivateZoneClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateZoneClient", reflect.TypeOf((*MockClientFactory)(nil).GetPrivateZoneClient))
}

// GetPublicIPAddressClient mocks base method.
func (m *MockClientFactory) GetPublicIPAddressClient() publicipaddressclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicIPAddressClient")
	ret0, _ := ret[0].(publicipaddressclient.Interface)
	return ret0
}

// GetPublicIPAddressClient indicates an expected call of GetPublicIPAddressClient.
func (mr *MockClientFactoryMockRecorder) GetPublicIPAddressClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicIPAddressClient", reflect.TypeOf((*MockClientFactory)(nil).GetPublicIPAddressClient))
}

// GetPublicIPPrefixClient mocks base method.
func (m *MockClientFactory) GetPublicIPPrefixClient() publicipprefixclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicIPPrefixClient")
	ret0, _ := ret[0].(publicipprefixclient.Interface)
	return ret0
}

// GetPublicIPPrefixClient indicates an expected call of GetPublicIPPrefixClient.
func (mr *MockClientFactoryMockRecorder) GetPublicIPPrefixClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicIPPrefixClient", reflect.TypeOf((*MockClientFactory)(nil).GetPublicIPPrefixClient))
}

// GetRegistryClient mocks base method.
func (m *MockClientFactory) GetRegistryClient() registryclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegistryClient")
	ret0, _ := ret[0].(registryclient.Interface)
	return ret0
}

// GetRegistryClient indicates an expected call of GetRegistryClient.
func (mr *MockClientFactoryMockRecorder) GetRegistryClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegistryClient", reflect.TypeOf((*MockClientFactory)(nil).GetRegistryClient))
}

// GetResourceGroupClient mocks base method.
func (m *MockClientFactory) GetResourceGroupClient() resourcegroupclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceGroupClient")
	ret0, _ := ret[0].(resourcegroupclient.Interface)
	return ret0
}

// GetResourceGroupClient indicates an expected call of GetResourceGroupClient.
func (mr *MockClientFactoryMockRecorder) GetResourceGroupClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceGroupClient", reflect.TypeOf((*MockClientFactory)(nil).GetResourceGroupClient))
}

// GetRouteTableClient mocks base method.
func (m *MockClientFactory) GetRouteTableClient() routetableclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRouteTableClient")
	ret0, _ := ret[0].(routetableclient.Interface)
	return ret0
}

// GetRouteTableClient indicates an expected call of GetRouteTableClient.
func (mr *MockClientFactoryMockRecorder) GetRouteTableClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRouteTableClient", reflect.TypeOf((*MockClientFactory)(nil).GetRouteTableClient))
}

// GetSSHPublicKeyResourceClient mocks base method.
func (m *MockClientFactory) GetSSHPublicKeyResourceClient() sshpublickeyresourceclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSSHPublicKeyResourceClient")
	ret0, _ := ret[0].(sshpublickeyresourceclient.Interface)
	return ret0
}

// GetSSHPublicKeyResourceClient indicates an expected call of GetSSHPublicKeyResourceClient.
func (mr *MockClientFactoryMockRecorder) GetSSHPublicKeyResourceClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSSHPublicKeyResourceClient", reflect.TypeOf((*MockClientFactory)(nil).GetSSHPublicKeyResourceClient))
}

// GetSecurityGroupClient mocks base method.
func (m *MockClientFactory) GetSecurityGroupClient() securitygroupclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecurityGroupClient")
	ret0, _ := ret[0].(securitygroupclient.Interface)
	return ret0
}

// GetSecurityGroupClient indicates an expected call of GetSecurityGroupClient.
func (mr *MockClientFactoryMockRecorder) GetSecurityGroupClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecurityGroupClient", reflect.TypeOf((*MockClientFactory)(nil).GetSecurityGroupClient))
}

// GetSnapshotClient mocks base method.
func (m *MockClientFactory) GetSnapshotClient() snapshotclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapshotClient")
	ret0, _ := ret[0].(snapshotclient.Interface)
	return ret0
}

// GetSnapshotClient indicates an expected call of GetSnapshotClient.
func (mr *MockClientFactoryMockRecorder) GetSnapshotClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshotClient", reflect.TypeOf((*MockClientFactory)(nil).GetSnapshotClient))
}

// GetSubnetClient mocks base method.
func (m *MockClientFactory) GetSubnetClient() subnetclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnetClient")
	ret0, _ := ret[0].(subnetclient.Interface)
	return ret0
}

// GetSubnetClient indicates an expected call of GetSubnetClient.
func (mr *MockClientFactoryMockRecorder) GetSubnetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnetClient", reflect.TypeOf((*MockClientFactory)(nil).GetSubnetClient))
}

// GetVirtualMachineClient mocks base method.
func (m *MockClientFactory) GetVirtualMachineClient() virtualmachineclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachineClient")
	ret0, _ := ret[0].(virtualmachineclient.Interface)
	return ret0
}

// GetVirtualMachineClient indicates an expected call of GetVirtualMachineClient.
func (mr *MockClientFactoryMockRecorder) GetVirtualMachineClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachineClient", reflect.TypeOf((*MockClientFactory)(nil).GetVirtualMachineClient))
}

// GetVirtualMachineScaleSetClient mocks base method.
func (m *MockClientFactory) GetVirtualMachineScaleSetClient() virtualmachinescalesetclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachineScaleSetClient")
	ret0, _ := ret[0].(virtualmachinescalesetclient.Interface)
	return ret0
}

// GetVirtualMachineScaleSetClient indicates an expected call of GetVirtualMachineScaleSetClient.
func (mr *MockClientFactoryMockRecorder) GetVirtualMachineScaleSetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachineScaleSetClient", reflect.TypeOf((*MockClientFactory)(nil).GetVirtualMachineScaleSetClient))
}

// GetVirtualMachineScaleSetVMClient mocks base method.
func (m *MockClientFactory) GetVirtualMachineScaleSetVMClient() virtualmachinescalesetvmclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachineScaleSetVMClient")
	ret0, _ := ret[0].(virtualmachinescalesetvmclient.Interface)
	return ret0
}

// GetVirtualMachineScaleSetVMClient indicates an expected call of GetVirtualMachineScaleSetVMClient.
func (mr *MockClientFactoryMockRecorder) GetVirtualMachineScaleSetVMClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachineScaleSetVMClient", reflect.TypeOf((*MockClientFactory)(nil).GetVirtualMachineScaleSetVMClient))
}

// GetVirtualNetworkClient mocks base method.
func (m *MockClientFactory) GetVirtualNetworkClient() virtualnetworkclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualNetworkClient")
	ret0, _ := ret[0].(virtualnetworkclient.Interface)
	return ret0
}

// GetVirtualNetworkClient indicates an expected call of GetVirtualNetworkClient.
func (mr *MockClientFactoryMockRecorder) GetVirtualNetworkClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualNetworkClient", reflect.TypeOf((*MockClientFactory)(nil).GetVirtualNetworkClient))
}
