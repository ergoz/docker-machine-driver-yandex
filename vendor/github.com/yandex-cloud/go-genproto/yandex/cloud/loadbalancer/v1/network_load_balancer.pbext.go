// Code generated by protoc-gen-goext. DO NOT EDIT.

package loadbalancer

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *NetworkLoadBalancer) SetId(v string) {
	m.Id = v
}

func (m *NetworkLoadBalancer) SetFolderId(v string) {
	m.FolderId = v
}

func (m *NetworkLoadBalancer) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *NetworkLoadBalancer) SetName(v string) {
	m.Name = v
}

func (m *NetworkLoadBalancer) SetDescription(v string) {
	m.Description = v
}

func (m *NetworkLoadBalancer) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *NetworkLoadBalancer) SetRegionId(v string) {
	m.RegionId = v
}

func (m *NetworkLoadBalancer) SetStatus(v NetworkLoadBalancer_Status) {
	m.Status = v
}

func (m *NetworkLoadBalancer) SetType(v NetworkLoadBalancer_Type) {
	m.Type = v
}

func (m *NetworkLoadBalancer) SetSessionAffinity(v NetworkLoadBalancer_SessionAffinity) {
	m.SessionAffinity = v
}

func (m *NetworkLoadBalancer) SetListeners(v []*Listener) {
	m.Listeners = v
}

func (m *NetworkLoadBalancer) SetAttachedTargetGroups(v []*AttachedTargetGroup) {
	m.AttachedTargetGroups = v
}

func (m *NetworkLoadBalancer) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *AttachedTargetGroup) SetTargetGroupId(v string) {
	m.TargetGroupId = v
}

func (m *AttachedTargetGroup) SetHealthChecks(v []*HealthCheck) {
	m.HealthChecks = v
}

func (m *Listener) SetName(v string) {
	m.Name = v
}

func (m *Listener) SetAddress(v string) {
	m.Address = v
}

func (m *Listener) SetPort(v int64) {
	m.Port = v
}

func (m *Listener) SetProtocol(v Listener_Protocol) {
	m.Protocol = v
}

func (m *Listener) SetTargetPort(v int64) {
	m.TargetPort = v
}

func (m *Listener) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *Listener) SetIpVersion(v IpVersion) {
	m.IpVersion = v
}

func (m *TargetState) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *TargetState) SetAddress(v string) {
	m.Address = v
}

func (m *TargetState) SetStatus(v TargetState_Status) {
	m.Status = v
}
