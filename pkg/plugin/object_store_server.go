package plugin

import (
	"io"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	proto "github.com/heptio/ark/pkg/plugin/generated"
	"github.com/heptio/ark/pkg/plugin/interface/objectinterface"
)

// GRPCServer registers an ObjectStore gRPC server.
func (p *ObjectStorePlugin) GRPCServer(s *grpc.Server) error {
	proto.RegisterObjectStoreServer(s, &ObjectStoreGRPCServer{mux: p.serverMux})
	return nil
}

// ObjectStoreGRPCServer implements the proto-generated ObjectStoreServer interface, and accepts
// gRPC calls and forwards them to an implementation of the pluggable interface.
type ObjectStoreGRPCServer struct {
	mux *serverMux
}

func (s *ObjectStoreGRPCServer) getImpl(name string) (objectinterface.ObjectStore, error) {
	impl, err := s.mux.getHandler(name)
	if err != nil {
		return nil, err
	}

	os, ok := impl.(objectinterface.ObjectStore)
	if !ok {
		return nil, errors.Errorf("%T is not an object store", impl)
	}

	return os, nil
}

// Init prepares the ObjectStore for usage using the provided map of
// configuration key-value pairs. It returns an error if the ObjectStore
// cannot be initialized from the provided config.
func (s *ObjectStoreGRPCServer) Init(ctx context.Context, req *proto.InitRequest) (*proto.Empty, error) {
	impl, err := s.getImpl(req.Plugin)
	if err != nil {
		return nil, err
	}

	if err := impl.Init(req.Config); err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}

// PutObject creates a new object using the data in body within the specified
// object storage bucket with the given key.
func (s *ObjectStoreGRPCServer) PutObject(stream proto.ObjectStore_PutObjectServer) error {
	// we need to read the first chunk ahead of time to get the bucket and key;
	// in our receive method, we'll use `first` on the first call
	firstChunk, err := stream.Recv()
	if err != nil {
		return err
	}

	impl, err := s.getImpl(firstChunk.Plugin)
	if err != nil {
		return err
	}

	bucket := firstChunk.Bucket
	key := firstChunk.Key

	receive := func() ([]byte, error) {
		if firstChunk != nil {
			res := firstChunk.Body
			firstChunk = nil
			return res, nil
		}

		data, err := stream.Recv()
		if err != nil {
			return nil, err
		}
		return data.Body, nil
	}

	close := func() error {
		return nil
	}

	if err := impl.PutObject(bucket, key, &StreamReadCloser{receive: receive, close: close}); err != nil {
		return err
	}

	return stream.SendAndClose(&proto.Empty{})
}

// GetObject retrieves the object with the given key from the specified
// bucket in object storage.
func (s *ObjectStoreGRPCServer) GetObject(req *proto.GetObjectRequest, stream proto.ObjectStore_GetObjectServer) error {
	impl, err := s.getImpl(req.Plugin)
	if err != nil {
		return err
	}

	rdr, err := impl.GetObject(req.Bucket, req.Key)
	if err != nil {
		return err
	}

	chunk := make([]byte, byteChunkSize)
	for {
		n, err := rdr.Read(chunk)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			return nil
		}

		if err := stream.Send(&proto.Bytes{Data: chunk[0:n]}); err != nil {
			return err
		}
	}
}

// ListCommonPrefixes gets a list of all object key prefixes that start with
// the specified prefix and stop at the next instance of the provided delimiter
// (this is often used to simulate a directory hierarchy in object storage).
func (s *ObjectStoreGRPCServer) ListCommonPrefixes(ctx context.Context, req *proto.ListCommonPrefixesRequest) (*proto.ListCommonPrefixesResponse, error) {
	impl, err := s.getImpl(req.Plugin)
	if err != nil {
		return nil, err
	}

	prefixes, err := impl.ListCommonPrefixes(req.Bucket, req.Prefix, req.Delimiter)
	if err != nil {
		return nil, err
	}

	return &proto.ListCommonPrefixesResponse{Prefixes: prefixes}, nil
}

// ListObjects gets a list of all objects in bucket that have the same prefix.
func (s *ObjectStoreGRPCServer) ListObjects(ctx context.Context, req *proto.ListObjectsRequest) (*proto.ListObjectsResponse, error) {
	impl, err := s.getImpl(req.Plugin)
	if err != nil {
		return nil, err
	}

	keys, err := impl.ListObjects(req.Bucket, req.Prefix)
	if err != nil {
		return nil, err
	}

	return &proto.ListObjectsResponse{Keys: keys}, nil
}

// DeleteObject removes object with the specified key from the given
// bucket.
func (s *ObjectStoreGRPCServer) DeleteObject(ctx context.Context, req *proto.DeleteObjectRequest) (*proto.Empty, error) {
	impl, err := s.getImpl(req.Plugin)
	if err != nil {
		return nil, err
	}

	if err := impl.DeleteObject(req.Bucket, req.Key); err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}

// CreateSignedURL creates a pre-signed URL for the given bucket and key that expires after ttl.
func (s *ObjectStoreGRPCServer) CreateSignedURL(ctx context.Context, req *proto.CreateSignedURLRequest) (*proto.CreateSignedURLResponse, error) {
	impl, err := s.getImpl(req.Plugin)
	if err != nil {
		return nil, err
	}

	url, err := impl.CreateSignedURL(req.Bucket, req.Key, time.Duration(req.Ttl))
	if err != nil {
		return nil, err
	}

	return &proto.CreateSignedURLResponse{Url: url}, nil
}
