// Code generated by sdkgen. DO NOT EDIT.

// nolint
package organizationmanager

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	organizationmanager "github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
)

//revive:disable

// GroupServiceClient is a organizationmanager.GroupServiceClient with
// lazy GRPC connection initialization.
type GroupServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements organizationmanager.GroupServiceClient
func (c *GroupServiceClient) Create(ctx context.Context, in *organizationmanager.CreateGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return organizationmanager.NewGroupServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements organizationmanager.GroupServiceClient
func (c *GroupServiceClient) Delete(ctx context.Context, in *organizationmanager.DeleteGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return organizationmanager.NewGroupServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements organizationmanager.GroupServiceClient
func (c *GroupServiceClient) Get(ctx context.Context, in *organizationmanager.GetGroupRequest, opts ...grpc.CallOption) (*organizationmanager.Group, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return organizationmanager.NewGroupServiceClient(conn).Get(ctx, in, opts...)
}

// List implements organizationmanager.GroupServiceClient
func (c *GroupServiceClient) List(ctx context.Context, in *organizationmanager.ListGroupsRequest, opts ...grpc.CallOption) (*organizationmanager.ListGroupsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return organizationmanager.NewGroupServiceClient(conn).List(ctx, in, opts...)
}

type GroupIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *GroupServiceClient
	request *organizationmanager.ListGroupsRequest

	items []*organizationmanager.Group
}

func (c *GroupServiceClient) GroupIterator(ctx context.Context, req *organizationmanager.ListGroupsRequest, opts ...grpc.CallOption) *GroupIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &GroupIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *GroupIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Groups
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *GroupIterator) Take(size int64) ([]*organizationmanager.Group, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*organizationmanager.Group

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *GroupIterator) TakeAll() ([]*organizationmanager.Group, error) {
	return it.Take(0)
}

func (it *GroupIterator) Value() *organizationmanager.Group {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *GroupIterator) Error() error {
	return it.err
}

// ListAccessBindings implements organizationmanager.GroupServiceClient
func (c *GroupServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return organizationmanager.NewGroupServiceClient(conn).ListAccessBindings(ctx, in, opts...)
}

type GroupAccessBindingsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *GroupServiceClient
	request *access.ListAccessBindingsRequest

	items []*access.AccessBinding
}

func (c *GroupServiceClient) GroupAccessBindingsIterator(ctx context.Context, req *access.ListAccessBindingsRequest, opts ...grpc.CallOption) *GroupAccessBindingsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &GroupAccessBindingsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *GroupAccessBindingsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListAccessBindings(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.AccessBindings
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *GroupAccessBindingsIterator) Take(size int64) ([]*access.AccessBinding, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*access.AccessBinding

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *GroupAccessBindingsIterator) TakeAll() ([]*access.AccessBinding, error) {
	return it.Take(0)
}

func (it *GroupAccessBindingsIterator) Value() *access.AccessBinding {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *GroupAccessBindingsIterator) Error() error {
	return it.err
}

// ListMembers implements organizationmanager.GroupServiceClient
func (c *GroupServiceClient) ListMembers(ctx context.Context, in *organizationmanager.ListGroupMembersRequest, opts ...grpc.CallOption) (*organizationmanager.ListGroupMembersResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return organizationmanager.NewGroupServiceClient(conn).ListMembers(ctx, in, opts...)
}

type GroupMembersIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *GroupServiceClient
	request *organizationmanager.ListGroupMembersRequest

	items []*organizationmanager.GroupMember
}

func (c *GroupServiceClient) GroupMembersIterator(ctx context.Context, req *organizationmanager.ListGroupMembersRequest, opts ...grpc.CallOption) *GroupMembersIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &GroupMembersIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *GroupMembersIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListMembers(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Members
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *GroupMembersIterator) Take(size int64) ([]*organizationmanager.GroupMember, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*organizationmanager.GroupMember

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *GroupMembersIterator) TakeAll() ([]*organizationmanager.GroupMember, error) {
	return it.Take(0)
}

func (it *GroupMembersIterator) Value() *organizationmanager.GroupMember {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *GroupMembersIterator) Error() error {
	return it.err
}

// ListOperations implements organizationmanager.GroupServiceClient
func (c *GroupServiceClient) ListOperations(ctx context.Context, in *organizationmanager.ListGroupOperationsRequest, opts ...grpc.CallOption) (*organizationmanager.ListGroupOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return organizationmanager.NewGroupServiceClient(conn).ListOperations(ctx, in, opts...)
}

type GroupOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *GroupServiceClient
	request *organizationmanager.ListGroupOperationsRequest

	items []*operation.Operation
}

func (c *GroupServiceClient) GroupOperationsIterator(ctx context.Context, req *organizationmanager.ListGroupOperationsRequest, opts ...grpc.CallOption) *GroupOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &GroupOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *GroupOperationsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *GroupOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*operation.Operation

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *GroupOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *GroupOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *GroupOperationsIterator) Error() error {
	return it.err
}

// SetAccessBindings implements organizationmanager.GroupServiceClient
func (c *GroupServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return organizationmanager.NewGroupServiceClient(conn).SetAccessBindings(ctx, in, opts...)
}

// Update implements organizationmanager.GroupServiceClient
func (c *GroupServiceClient) Update(ctx context.Context, in *organizationmanager.UpdateGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return organizationmanager.NewGroupServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateAccessBindings implements organizationmanager.GroupServiceClient
func (c *GroupServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return organizationmanager.NewGroupServiceClient(conn).UpdateAccessBindings(ctx, in, opts...)
}

// UpdateMembers implements organizationmanager.GroupServiceClient
func (c *GroupServiceClient) UpdateMembers(ctx context.Context, in *organizationmanager.UpdateGroupMembersRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return organizationmanager.NewGroupServiceClient(conn).UpdateMembers(ctx, in, opts...)
}
