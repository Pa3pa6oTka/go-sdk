// Code generated by sdkgen. DO NOT EDIT.

//nolint
package compute

import (
	"context"

	"google.golang.org/grpc"

	compute "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// SnapshotScheduleServiceClient is a compute.SnapshotScheduleServiceClient with
// lazy GRPC connection initialization.
type SnapshotScheduleServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements compute.SnapshotScheduleServiceClient
func (c *SnapshotScheduleServiceClient) Create(ctx context.Context, in *compute.CreateSnapshotScheduleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewSnapshotScheduleServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements compute.SnapshotScheduleServiceClient
func (c *SnapshotScheduleServiceClient) Delete(ctx context.Context, in *compute.DeleteSnapshotScheduleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewSnapshotScheduleServiceClient(conn).Delete(ctx, in, opts...)
}

// Disable implements compute.SnapshotScheduleServiceClient
func (c *SnapshotScheduleServiceClient) Disable(ctx context.Context, in *compute.DisableSnapshotScheduleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewSnapshotScheduleServiceClient(conn).Disable(ctx, in, opts...)
}

// Enable implements compute.SnapshotScheduleServiceClient
func (c *SnapshotScheduleServiceClient) Enable(ctx context.Context, in *compute.EnableSnapshotScheduleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewSnapshotScheduleServiceClient(conn).Enable(ctx, in, opts...)
}

// Get implements compute.SnapshotScheduleServiceClient
func (c *SnapshotScheduleServiceClient) Get(ctx context.Context, in *compute.GetSnapshotScheduleRequest, opts ...grpc.CallOption) (*compute.SnapshotSchedule, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewSnapshotScheduleServiceClient(conn).Get(ctx, in, opts...)
}

// List implements compute.SnapshotScheduleServiceClient
func (c *SnapshotScheduleServiceClient) List(ctx context.Context, in *compute.ListSnapshotSchedulesRequest, opts ...grpc.CallOption) (*compute.ListSnapshotSchedulesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewSnapshotScheduleServiceClient(conn).List(ctx, in, opts...)
}

type SnapshotScheduleIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *SnapshotScheduleServiceClient
	request *compute.ListSnapshotSchedulesRequest

	items []*compute.SnapshotSchedule
}

func (c *SnapshotScheduleServiceClient) SnapshotScheduleIterator(ctx context.Context, req *compute.ListSnapshotSchedulesRequest, opts ...grpc.CallOption) *SnapshotScheduleIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &SnapshotScheduleIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *SnapshotScheduleIterator) Next() bool {
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

	it.items = response.SnapshotSchedules
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *SnapshotScheduleIterator) Take(size int64) ([]*compute.SnapshotSchedule, error) {
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

	var result []*compute.SnapshotSchedule

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *SnapshotScheduleIterator) TakeAll() ([]*compute.SnapshotSchedule, error) {
	return it.Take(0)
}

func (it *SnapshotScheduleIterator) Value() *compute.SnapshotSchedule {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *SnapshotScheduleIterator) Error() error {
	return it.err
}

// ListDisks implements compute.SnapshotScheduleServiceClient
func (c *SnapshotScheduleServiceClient) ListDisks(ctx context.Context, in *compute.ListSnapshotScheduleDisksRequest, opts ...grpc.CallOption) (*compute.ListSnapshotScheduleDisksResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewSnapshotScheduleServiceClient(conn).ListDisks(ctx, in, opts...)
}

type SnapshotScheduleDisksIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *SnapshotScheduleServiceClient
	request *compute.ListSnapshotScheduleDisksRequest

	items []*compute.Disk
}

func (c *SnapshotScheduleServiceClient) SnapshotScheduleDisksIterator(ctx context.Context, req *compute.ListSnapshotScheduleDisksRequest, opts ...grpc.CallOption) *SnapshotScheduleDisksIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &SnapshotScheduleDisksIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *SnapshotScheduleDisksIterator) Next() bool {
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

	response, err := it.client.ListDisks(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Disks
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *SnapshotScheduleDisksIterator) Take(size int64) ([]*compute.Disk, error) {
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

	var result []*compute.Disk

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *SnapshotScheduleDisksIterator) TakeAll() ([]*compute.Disk, error) {
	return it.Take(0)
}

func (it *SnapshotScheduleDisksIterator) Value() *compute.Disk {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *SnapshotScheduleDisksIterator) Error() error {
	return it.err
}

// ListOperations implements compute.SnapshotScheduleServiceClient
func (c *SnapshotScheduleServiceClient) ListOperations(ctx context.Context, in *compute.ListSnapshotScheduleOperationsRequest, opts ...grpc.CallOption) (*compute.ListSnapshotScheduleOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewSnapshotScheduleServiceClient(conn).ListOperations(ctx, in, opts...)
}

type SnapshotScheduleOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *SnapshotScheduleServiceClient
	request *compute.ListSnapshotScheduleOperationsRequest

	items []*operation.Operation
}

func (c *SnapshotScheduleServiceClient) SnapshotScheduleOperationsIterator(ctx context.Context, req *compute.ListSnapshotScheduleOperationsRequest, opts ...grpc.CallOption) *SnapshotScheduleOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &SnapshotScheduleOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *SnapshotScheduleOperationsIterator) Next() bool {
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

func (it *SnapshotScheduleOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
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

func (it *SnapshotScheduleOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *SnapshotScheduleOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *SnapshotScheduleOperationsIterator) Error() error {
	return it.err
}

// ListSnapshots implements compute.SnapshotScheduleServiceClient
func (c *SnapshotScheduleServiceClient) ListSnapshots(ctx context.Context, in *compute.ListSnapshotScheduleSnapshotsRequest, opts ...grpc.CallOption) (*compute.ListSnapshotScheduleSnapshotsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewSnapshotScheduleServiceClient(conn).ListSnapshots(ctx, in, opts...)
}

type SnapshotScheduleSnapshotsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *SnapshotScheduleServiceClient
	request *compute.ListSnapshotScheduleSnapshotsRequest

	items []*compute.Snapshot
}

func (c *SnapshotScheduleServiceClient) SnapshotScheduleSnapshotsIterator(ctx context.Context, req *compute.ListSnapshotScheduleSnapshotsRequest, opts ...grpc.CallOption) *SnapshotScheduleSnapshotsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &SnapshotScheduleSnapshotsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *SnapshotScheduleSnapshotsIterator) Next() bool {
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

	response, err := it.client.ListSnapshots(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Snapshots
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *SnapshotScheduleSnapshotsIterator) Take(size int64) ([]*compute.Snapshot, error) {
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

	var result []*compute.Snapshot

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *SnapshotScheduleSnapshotsIterator) TakeAll() ([]*compute.Snapshot, error) {
	return it.Take(0)
}

func (it *SnapshotScheduleSnapshotsIterator) Value() *compute.Snapshot {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *SnapshotScheduleSnapshotsIterator) Error() error {
	return it.err
}

// Update implements compute.SnapshotScheduleServiceClient
func (c *SnapshotScheduleServiceClient) Update(ctx context.Context, in *compute.UpdateSnapshotScheduleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewSnapshotScheduleServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateDisks implements compute.SnapshotScheduleServiceClient
func (c *SnapshotScheduleServiceClient) UpdateDisks(ctx context.Context, in *compute.UpdateSnapshotScheduleDisksRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewSnapshotScheduleServiceClient(conn).UpdateDisks(ctx, in, opts...)
}
