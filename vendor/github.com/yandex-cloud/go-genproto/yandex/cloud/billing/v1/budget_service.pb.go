// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: yandex/cloud/billing/v1/budget_service.proto

package billing

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetBudgetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the budget to return.
	// To get the budget ID, use [BudgetService.List] request.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBudgetRequest) Reset() {
	*x = GetBudgetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_budget_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBudgetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBudgetRequest) ProtoMessage() {}

func (x *GetBudgetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_budget_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBudgetRequest.ProtoReflect.Descriptor instead.
func (*GetBudgetRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_budget_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetBudgetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListBudgetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the billing account to list budgets corresponding to.
	// To get the billing account ID, use [BillingAccountService.List] request.
	BillingAccountId string `protobuf:"bytes,1,opt,name=billing_account_id,json=billingAccountId,proto3" json:"billing_account_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size],
	// the service returns a [ListBudgetsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results,
	// set [page_token] to the [ListBudgetsResponse.next_page_token]
	// returned by a previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListBudgetsRequest) Reset() {
	*x = ListBudgetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_budget_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBudgetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBudgetsRequest) ProtoMessage() {}

func (x *ListBudgetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_budget_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBudgetsRequest.ProtoReflect.Descriptor instead.
func (*ListBudgetsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_budget_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListBudgetsRequest) GetBillingAccountId() string {
	if x != nil {
		return x.BillingAccountId
	}
	return ""
}

func (x *ListBudgetsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListBudgetsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListBudgetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of budgets.
	Budgets []*Budget `protobuf:"bytes,1,rep,name=budgets,proto3" json:"budgets,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListBudgetsRequest.page_size], use
	// [next_page_token] as the value
	// for the [ListBudgetsRequest.page_token] query parameter
	// in the next list request. Each subsequent list request will have its own
	// [next_page_token] to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListBudgetsResponse) Reset() {
	*x = ListBudgetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_budget_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBudgetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBudgetsResponse) ProtoMessage() {}

func (x *ListBudgetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_budget_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBudgetsResponse.ProtoReflect.Descriptor instead.
func (*ListBudgetsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_budget_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListBudgetsResponse) GetBudgets() []*Budget {
	if x != nil {
		return x.Budgets
	}
	return nil
}

func (x *ListBudgetsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateBudgetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the billing account to list budgets corresponding to.
	// To get the billing account ID, use [yandex.cloud.billing.v1.BillingAccountService.List] request.
	BillingAccountId string `protobuf:"bytes,1,opt,name=billing_account_id,json=billingAccountId,proto3" json:"billing_account_id,omitempty"`
	// Name of the budget.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Specification of the budget.
	//
	// Types that are assignable to BudgetSpec:
	//
	//	*CreateBudgetRequest_CostBudgetSpec
	//	*CreateBudgetRequest_ExpenseBudgetSpec
	//	*CreateBudgetRequest_BalanceBudgetSpec
	BudgetSpec isCreateBudgetRequest_BudgetSpec `protobuf_oneof:"budget_spec"`
}

func (x *CreateBudgetRequest) Reset() {
	*x = CreateBudgetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_budget_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBudgetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBudgetRequest) ProtoMessage() {}

func (x *CreateBudgetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_budget_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBudgetRequest.ProtoReflect.Descriptor instead.
func (*CreateBudgetRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_budget_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBudgetRequest) GetBillingAccountId() string {
	if x != nil {
		return x.BillingAccountId
	}
	return ""
}

func (x *CreateBudgetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *CreateBudgetRequest) GetBudgetSpec() isCreateBudgetRequest_BudgetSpec {
	if m != nil {
		return m.BudgetSpec
	}
	return nil
}

func (x *CreateBudgetRequest) GetCostBudgetSpec() *CostBudgetSpec {
	if x, ok := x.GetBudgetSpec().(*CreateBudgetRequest_CostBudgetSpec); ok {
		return x.CostBudgetSpec
	}
	return nil
}

func (x *CreateBudgetRequest) GetExpenseBudgetSpec() *ExpenseBudgetSpec {
	if x, ok := x.GetBudgetSpec().(*CreateBudgetRequest_ExpenseBudgetSpec); ok {
		return x.ExpenseBudgetSpec
	}
	return nil
}

func (x *CreateBudgetRequest) GetBalanceBudgetSpec() *BalanceBudgetSpec {
	if x, ok := x.GetBudgetSpec().(*CreateBudgetRequest_BalanceBudgetSpec); ok {
		return x.BalanceBudgetSpec
	}
	return nil
}

type isCreateBudgetRequest_BudgetSpec interface {
	isCreateBudgetRequest_BudgetSpec()
}

type CreateBudgetRequest_CostBudgetSpec struct {
	// Cost budget specification.
	CostBudgetSpec *CostBudgetSpec `protobuf:"bytes,3,opt,name=cost_budget_spec,json=costBudgetSpec,proto3,oneof"`
}

type CreateBudgetRequest_ExpenseBudgetSpec struct {
	// Expense budget specification.
	ExpenseBudgetSpec *ExpenseBudgetSpec `protobuf:"bytes,4,opt,name=expense_budget_spec,json=expenseBudgetSpec,proto3,oneof"`
}

type CreateBudgetRequest_BalanceBudgetSpec struct {
	// Balance budget specification.
	BalanceBudgetSpec *BalanceBudgetSpec `protobuf:"bytes,5,opt,name=balance_budget_spec,json=balanceBudgetSpec,proto3,oneof"`
}

func (*CreateBudgetRequest_CostBudgetSpec) isCreateBudgetRequest_BudgetSpec() {}

func (*CreateBudgetRequest_ExpenseBudgetSpec) isCreateBudgetRequest_BudgetSpec() {}

func (*CreateBudgetRequest_BalanceBudgetSpec) isCreateBudgetRequest_BudgetSpec() {}

type CreateBudgetMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the budget.
	BudgetId string `protobuf:"bytes,1,opt,name=budget_id,json=budgetId,proto3" json:"budget_id,omitempty"`
}

func (x *CreateBudgetMetadata) Reset() {
	*x = CreateBudgetMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_budget_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBudgetMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBudgetMetadata) ProtoMessage() {}

func (x *CreateBudgetMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_budget_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBudgetMetadata.ProtoReflect.Descriptor instead.
func (*CreateBudgetMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_budget_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateBudgetMetadata) GetBudgetId() string {
	if x != nil {
		return x.BudgetId
	}
	return ""
}

var File_yandex_cloud_billing_v1_budget_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_billing_v1_budget_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d,
	0x35, 0x30, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa3, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a,
	0x12, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a,
	0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x10, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0xfa, 0xc7,
	0x31, 0x06, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x30, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x28, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30,
	0x30, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x78, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x07, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x12, 0x26,
	0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x91, 0x03, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a,
	0x0a, 0x12, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01,
	0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x10, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x53, 0x0a, 0x10, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x62, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x6f, 0x73, 0x74, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x5c, 0x0a, 0x13, 0x65, 0x78, 0x70,
	0x65, 0x6e, 0x73, 0x65, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x53, 0x70,
	0x65, 0x63, 0x48, 0x00, 0x52, 0x11, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x5c, 0x0a, 0x13, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63,
	0x48, 0x00, 0x52, 0x11, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x42, 0x13, 0x0a, 0x0b, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x12, 0x04, 0xc0, 0xc1, 0x31, 0x01, 0x22, 0x33, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x49, 0x64, 0x32,
	0xa2, 0x03, 0x0a, 0x0d, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x73, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x29, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7e, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x15, 0x12, 0x13, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x12, 0x9b, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x2c, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x40, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x3a,
	0x01, 0x2a, 0xb2, 0xd2, 0x2a, 0x1e, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x06, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x42, 0x62, 0x0a, 0x1b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x3b, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_billing_v1_budget_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_billing_v1_budget_service_proto_rawDescData = file_yandex_cloud_billing_v1_budget_service_proto_rawDesc
)

func file_yandex_cloud_billing_v1_budget_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_billing_v1_budget_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_billing_v1_budget_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_billing_v1_budget_service_proto_rawDescData)
	})
	return file_yandex_cloud_billing_v1_budget_service_proto_rawDescData
}

var file_yandex_cloud_billing_v1_budget_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_billing_v1_budget_service_proto_goTypes = []interface{}{
	(*GetBudgetRequest)(nil),     // 0: yandex.cloud.billing.v1.GetBudgetRequest
	(*ListBudgetsRequest)(nil),   // 1: yandex.cloud.billing.v1.ListBudgetsRequest
	(*ListBudgetsResponse)(nil),  // 2: yandex.cloud.billing.v1.ListBudgetsResponse
	(*CreateBudgetRequest)(nil),  // 3: yandex.cloud.billing.v1.CreateBudgetRequest
	(*CreateBudgetMetadata)(nil), // 4: yandex.cloud.billing.v1.CreateBudgetMetadata
	(*Budget)(nil),               // 5: yandex.cloud.billing.v1.Budget
	(*CostBudgetSpec)(nil),       // 6: yandex.cloud.billing.v1.CostBudgetSpec
	(*ExpenseBudgetSpec)(nil),    // 7: yandex.cloud.billing.v1.ExpenseBudgetSpec
	(*BalanceBudgetSpec)(nil),    // 8: yandex.cloud.billing.v1.BalanceBudgetSpec
	(*operation.Operation)(nil),  // 9: yandex.cloud.operation.Operation
}
var file_yandex_cloud_billing_v1_budget_service_proto_depIdxs = []int32{
	5, // 0: yandex.cloud.billing.v1.ListBudgetsResponse.budgets:type_name -> yandex.cloud.billing.v1.Budget
	6, // 1: yandex.cloud.billing.v1.CreateBudgetRequest.cost_budget_spec:type_name -> yandex.cloud.billing.v1.CostBudgetSpec
	7, // 2: yandex.cloud.billing.v1.CreateBudgetRequest.expense_budget_spec:type_name -> yandex.cloud.billing.v1.ExpenseBudgetSpec
	8, // 3: yandex.cloud.billing.v1.CreateBudgetRequest.balance_budget_spec:type_name -> yandex.cloud.billing.v1.BalanceBudgetSpec
	0, // 4: yandex.cloud.billing.v1.BudgetService.Get:input_type -> yandex.cloud.billing.v1.GetBudgetRequest
	1, // 5: yandex.cloud.billing.v1.BudgetService.List:input_type -> yandex.cloud.billing.v1.ListBudgetsRequest
	3, // 6: yandex.cloud.billing.v1.BudgetService.Create:input_type -> yandex.cloud.billing.v1.CreateBudgetRequest
	5, // 7: yandex.cloud.billing.v1.BudgetService.Get:output_type -> yandex.cloud.billing.v1.Budget
	2, // 8: yandex.cloud.billing.v1.BudgetService.List:output_type -> yandex.cloud.billing.v1.ListBudgetsResponse
	9, // 9: yandex.cloud.billing.v1.BudgetService.Create:output_type -> yandex.cloud.operation.Operation
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_billing_v1_budget_service_proto_init() }
func file_yandex_cloud_billing_v1_budget_service_proto_init() {
	if File_yandex_cloud_billing_v1_budget_service_proto != nil {
		return
	}
	file_yandex_cloud_billing_v1_budget_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_billing_v1_budget_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBudgetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_billing_v1_budget_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBudgetsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_billing_v1_budget_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBudgetsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_billing_v1_budget_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBudgetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_billing_v1_budget_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBudgetMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_yandex_cloud_billing_v1_budget_service_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*CreateBudgetRequest_CostBudgetSpec)(nil),
		(*CreateBudgetRequest_ExpenseBudgetSpec)(nil),
		(*CreateBudgetRequest_BalanceBudgetSpec)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_billing_v1_budget_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_billing_v1_budget_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_billing_v1_budget_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_billing_v1_budget_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_billing_v1_budget_service_proto = out.File
	file_yandex_cloud_billing_v1_budget_service_proto_rawDesc = nil
	file_yandex_cloud_billing_v1_budget_service_proto_goTypes = nil
	file_yandex_cloud_billing_v1_budget_service_proto_depIdxs = nil
}
