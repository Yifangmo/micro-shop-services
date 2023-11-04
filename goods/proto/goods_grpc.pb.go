// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/goods.proto

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GoodsClient is the client API for Goods service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsClient interface {
	// 商品
	GoodsListQuery(ctx context.Context, in *GoodsListQueryRequest, opts ...grpc.CallOption) (*GoodsListResponse, error)
	GetGoodsByIDs(ctx context.Context, in *GoodsIDsRequest, opts ...grpc.CallOption) (*GoodsListResponse, error)
	GetGoodsByID(ctx context.Context, in *GoodsIDRequest, opts ...grpc.CallOption) (*GoodsInfoResponse, error)
	CreateGoods(ctx context.Context, in *GoodsInfoRequest, opts ...grpc.CallOption) (*GoodsInfoResponse, error)
	DeleteGoods(ctx context.Context, in *GoodsIDRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateGoods(ctx context.Context, in *GoodsInfoRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// 商品分类
	AllCategoryList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CategoryListResponse, error)
	GetSubCategory(ctx context.Context, in *GetSubCategoryRequest, opts ...grpc.CallOption) (*SubCategoryListResponse, error)
	CreateCategory(ctx context.Context, in *CategoryInfoRequest, opts ...grpc.CallOption) (*CategoryInfoResponse, error)
	DeleteCategory(ctx context.Context, in *CategoryIDRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateCategory(ctx context.Context, in *CategoryInfoRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// 品牌
	BrandList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*BrandListResponse, error)
	CreateBrand(ctx context.Context, in *BrandInfoRequest, opts ...grpc.CallOption) (*BrandInfoResponse, error)
	DeleteBrand(ctx context.Context, in *BrandIDRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateBrand(ctx context.Context, in *BrandInfoRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// 商品分类与品牌
	CategoryBrandList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*CategoryBrandListResponse, error)
	GetCategoryBrandByBrandID(ctx context.Context, in *BrandIDRequest, opts ...grpc.CallOption) (*BrandListResponse, error)
	CreateCategoryBrand(ctx context.Context, in *CategoryBrandRequest, opts ...grpc.CallOption) (*CategoryBrandResponse, error)
	DeleteCategoryBrand(ctx context.Context, in *CategoryBrandRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateCategoryBrand(ctx context.Context, in *CategoryBrandRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// 轮播图
	BannerList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*BannerListResponse, error)
	CreateBanner(ctx context.Context, in *BannerRequest, opts ...grpc.CallOption) (*BannerResponse, error)
	DeleteBanner(ctx context.Context, in *BannerRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateBanner(ctx context.Context, in *BannerRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type goodsClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsClient(cc grpc.ClientConnInterface) GoodsClient {
	return &goodsClient{cc}
}

func (c *goodsClient) GoodsListQuery(ctx context.Context, in *GoodsListQueryRequest, opts ...grpc.CallOption) (*GoodsListResponse, error) {
	out := new(GoodsListResponse)
	err := c.cc.Invoke(ctx, "/Goods/GoodsListQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetGoodsByIDs(ctx context.Context, in *GoodsIDsRequest, opts ...grpc.CallOption) (*GoodsListResponse, error) {
	out := new(GoodsListResponse)
	err := c.cc.Invoke(ctx, "/Goods/GetGoodsByIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetGoodsByID(ctx context.Context, in *GoodsIDRequest, opts ...grpc.CallOption) (*GoodsInfoResponse, error) {
	out := new(GoodsInfoResponse)
	err := c.cc.Invoke(ctx, "/Goods/GetGoodsByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateGoods(ctx context.Context, in *GoodsInfoRequest, opts ...grpc.CallOption) (*GoodsInfoResponse, error) {
	out := new(GoodsInfoResponse)
	err := c.cc.Invoke(ctx, "/Goods/CreateGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) DeleteGoods(ctx context.Context, in *GoodsIDRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Goods/DeleteGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdateGoods(ctx context.Context, in *GoodsInfoRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Goods/UpdateGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) AllCategoryList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CategoryListResponse, error) {
	out := new(CategoryListResponse)
	err := c.cc.Invoke(ctx, "/Goods/AllCategoryList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetSubCategory(ctx context.Context, in *GetSubCategoryRequest, opts ...grpc.CallOption) (*SubCategoryListResponse, error) {
	out := new(SubCategoryListResponse)
	err := c.cc.Invoke(ctx, "/Goods/GetSubCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateCategory(ctx context.Context, in *CategoryInfoRequest, opts ...grpc.CallOption) (*CategoryInfoResponse, error) {
	out := new(CategoryInfoResponse)
	err := c.cc.Invoke(ctx, "/Goods/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) DeleteCategory(ctx context.Context, in *CategoryIDRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Goods/DeleteCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdateCategory(ctx context.Context, in *CategoryInfoRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Goods/UpdateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) BrandList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*BrandListResponse, error) {
	out := new(BrandListResponse)
	err := c.cc.Invoke(ctx, "/Goods/BrandList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateBrand(ctx context.Context, in *BrandInfoRequest, opts ...grpc.CallOption) (*BrandInfoResponse, error) {
	out := new(BrandInfoResponse)
	err := c.cc.Invoke(ctx, "/Goods/CreateBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) DeleteBrand(ctx context.Context, in *BrandIDRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Goods/DeleteBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdateBrand(ctx context.Context, in *BrandInfoRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Goods/UpdateBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CategoryBrandList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*CategoryBrandListResponse, error) {
	out := new(CategoryBrandListResponse)
	err := c.cc.Invoke(ctx, "/Goods/CategoryBrandList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetCategoryBrandByBrandID(ctx context.Context, in *BrandIDRequest, opts ...grpc.CallOption) (*BrandListResponse, error) {
	out := new(BrandListResponse)
	err := c.cc.Invoke(ctx, "/Goods/GetCategoryBrandByBrandID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateCategoryBrand(ctx context.Context, in *CategoryBrandRequest, opts ...grpc.CallOption) (*CategoryBrandResponse, error) {
	out := new(CategoryBrandResponse)
	err := c.cc.Invoke(ctx, "/Goods/CreateCategoryBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) DeleteCategoryBrand(ctx context.Context, in *CategoryBrandRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Goods/DeleteCategoryBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdateCategoryBrand(ctx context.Context, in *CategoryBrandRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Goods/UpdateCategoryBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) BannerList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*BannerListResponse, error) {
	out := new(BannerListResponse)
	err := c.cc.Invoke(ctx, "/Goods/BannerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateBanner(ctx context.Context, in *BannerRequest, opts ...grpc.CallOption) (*BannerResponse, error) {
	out := new(BannerResponse)
	err := c.cc.Invoke(ctx, "/Goods/CreateBanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) DeleteBanner(ctx context.Context, in *BannerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Goods/DeleteBanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdateBanner(ctx context.Context, in *BannerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Goods/UpdateBanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServer is the server API for Goods service.
// All implementations must embed UnimplementedGoodsServer
// for forward compatibility
type GoodsServer interface {
	// 商品
	GoodsListQuery(context.Context, *GoodsListQueryRequest) (*GoodsListResponse, error)
	GetGoodsByIDs(context.Context, *GoodsIDsRequest) (*GoodsListResponse, error)
	GetGoodsByID(context.Context, *GoodsIDRequest) (*GoodsInfoResponse, error)
	CreateGoods(context.Context, *GoodsInfoRequest) (*GoodsInfoResponse, error)
	DeleteGoods(context.Context, *GoodsIDRequest) (*empty.Empty, error)
	UpdateGoods(context.Context, *GoodsInfoRequest) (*empty.Empty, error)
	// 商品分类
	AllCategoryList(context.Context, *empty.Empty) (*CategoryListResponse, error)
	GetSubCategory(context.Context, *GetSubCategoryRequest) (*SubCategoryListResponse, error)
	CreateCategory(context.Context, *CategoryInfoRequest) (*CategoryInfoResponse, error)
	DeleteCategory(context.Context, *CategoryIDRequest) (*empty.Empty, error)
	UpdateCategory(context.Context, *CategoryInfoRequest) (*empty.Empty, error)
	// 品牌
	BrandList(context.Context, *PageInfo) (*BrandListResponse, error)
	CreateBrand(context.Context, *BrandInfoRequest) (*BrandInfoResponse, error)
	DeleteBrand(context.Context, *BrandIDRequest) (*empty.Empty, error)
	UpdateBrand(context.Context, *BrandInfoRequest) (*empty.Empty, error)
	// 商品分类与品牌
	CategoryBrandList(context.Context, *PageInfo) (*CategoryBrandListResponse, error)
	GetCategoryBrandByBrandID(context.Context, *BrandIDRequest) (*BrandListResponse, error)
	CreateCategoryBrand(context.Context, *CategoryBrandRequest) (*CategoryBrandResponse, error)
	DeleteCategoryBrand(context.Context, *CategoryBrandRequest) (*empty.Empty, error)
	UpdateCategoryBrand(context.Context, *CategoryBrandRequest) (*empty.Empty, error)
	// 轮播图
	BannerList(context.Context, *empty.Empty) (*BannerListResponse, error)
	CreateBanner(context.Context, *BannerRequest) (*BannerResponse, error)
	DeleteBanner(context.Context, *BannerRequest) (*empty.Empty, error)
	UpdateBanner(context.Context, *BannerRequest) (*empty.Empty, error)
	mustEmbedUnimplementedGoodsServer()
}

// UnimplementedGoodsServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsServer struct {
}

func (UnimplementedGoodsServer) GoodsListQuery(context.Context, *GoodsListQueryRequest) (*GoodsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoodsListQuery not implemented")
}
func (UnimplementedGoodsServer) GetGoodsByIDs(context.Context, *GoodsIDsRequest) (*GoodsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsByIDs not implemented")
}
func (UnimplementedGoodsServer) GetGoodsByID(context.Context, *GoodsIDRequest) (*GoodsInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsByID not implemented")
}
func (UnimplementedGoodsServer) CreateGoods(context.Context, *GoodsInfoRequest) (*GoodsInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGoods not implemented")
}
func (UnimplementedGoodsServer) DeleteGoods(context.Context, *GoodsIDRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGoods not implemented")
}
func (UnimplementedGoodsServer) UpdateGoods(context.Context, *GoodsInfoRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGoods not implemented")
}
func (UnimplementedGoodsServer) AllCategoryList(context.Context, *empty.Empty) (*CategoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllCategoryList not implemented")
}
func (UnimplementedGoodsServer) GetSubCategory(context.Context, *GetSubCategoryRequest) (*SubCategoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubCategory not implemented")
}
func (UnimplementedGoodsServer) CreateCategory(context.Context, *CategoryInfoRequest) (*CategoryInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedGoodsServer) DeleteCategory(context.Context, *CategoryIDRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedGoodsServer) UpdateCategory(context.Context, *CategoryInfoRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedGoodsServer) BrandList(context.Context, *PageInfo) (*BrandListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BrandList not implemented")
}
func (UnimplementedGoodsServer) CreateBrand(context.Context, *BrandInfoRequest) (*BrandInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBrand not implemented")
}
func (UnimplementedGoodsServer) DeleteBrand(context.Context, *BrandIDRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBrand not implemented")
}
func (UnimplementedGoodsServer) UpdateBrand(context.Context, *BrandInfoRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBrand not implemented")
}
func (UnimplementedGoodsServer) CategoryBrandList(context.Context, *PageInfo) (*CategoryBrandListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryBrandList not implemented")
}
func (UnimplementedGoodsServer) GetCategoryBrandByBrandID(context.Context, *BrandIDRequest) (*BrandListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryBrandByBrandID not implemented")
}
func (UnimplementedGoodsServer) CreateCategoryBrand(context.Context, *CategoryBrandRequest) (*CategoryBrandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategoryBrand not implemented")
}
func (UnimplementedGoodsServer) DeleteCategoryBrand(context.Context, *CategoryBrandRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategoryBrand not implemented")
}
func (UnimplementedGoodsServer) UpdateCategoryBrand(context.Context, *CategoryBrandRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategoryBrand not implemented")
}
func (UnimplementedGoodsServer) BannerList(context.Context, *empty.Empty) (*BannerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BannerList not implemented")
}
func (UnimplementedGoodsServer) CreateBanner(context.Context, *BannerRequest) (*BannerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBanner not implemented")
}
func (UnimplementedGoodsServer) DeleteBanner(context.Context, *BannerRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBanner not implemented")
}
func (UnimplementedGoodsServer) UpdateBanner(context.Context, *BannerRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBanner not implemented")
}
func (UnimplementedGoodsServer) mustEmbedUnimplementedGoodsServer() {}

// UnsafeGoodsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServer will
// result in compilation errors.
type UnsafeGoodsServer interface {
	mustEmbedUnimplementedGoodsServer()
}

func RegisterGoodsServer(s grpc.ServiceRegistrar, srv GoodsServer) {
	s.RegisterService(&Goods_ServiceDesc, srv)
}

func _Goods_GoodsListQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsListQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GoodsListQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/GoodsListQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GoodsListQuery(ctx, req.(*GoodsListQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetGoodsByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetGoodsByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/GetGoodsByIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetGoodsByIDs(ctx, req.(*GoodsIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetGoodsByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetGoodsByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/GetGoodsByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetGoodsByID(ctx, req.(*GoodsIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/CreateGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateGoods(ctx, req.(*GoodsInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_DeleteGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).DeleteGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/DeleteGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).DeleteGoods(ctx, req.(*GoodsIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdateGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdateGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/UpdateGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdateGoods(ctx, req.(*GoodsInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_AllCategoryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).AllCategoryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/AllCategoryList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).AllCategoryList(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetSubCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetSubCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/GetSubCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetSubCategory(ctx, req.(*GetSubCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateCategory(ctx, req.(*CategoryInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/DeleteCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).DeleteCategory(ctx, req.(*CategoryIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/UpdateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdateCategory(ctx, req.(*CategoryInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_BrandList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).BrandList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/BrandList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).BrandList(ctx, req.(*PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/CreateBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateBrand(ctx, req.(*BrandInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_DeleteBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).DeleteBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/DeleteBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).DeleteBrand(ctx, req.(*BrandIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdateBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdateBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/UpdateBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdateBrand(ctx, req.(*BrandInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CategoryBrandList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CategoryBrandList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/CategoryBrandList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CategoryBrandList(ctx, req.(*PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetCategoryBrandByBrandID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetCategoryBrandByBrandID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/GetCategoryBrandByBrandID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetCategoryBrandByBrandID(ctx, req.(*BrandIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateCategoryBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateCategoryBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/CreateCategoryBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateCategoryBrand(ctx, req.(*CategoryBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_DeleteCategoryBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).DeleteCategoryBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/DeleteCategoryBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).DeleteCategoryBrand(ctx, req.(*CategoryBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdateCategoryBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdateCategoryBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/UpdateCategoryBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdateCategoryBrand(ctx, req.(*CategoryBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_BannerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).BannerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/BannerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).BannerList(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/CreateBanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateBanner(ctx, req.(*BannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_DeleteBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).DeleteBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/DeleteBanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).DeleteBanner(ctx, req.(*BannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdateBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdateBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/UpdateBanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdateBanner(ctx, req.(*BannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Goods_ServiceDesc is the grpc.ServiceDesc for Goods service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Goods_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Goods",
	HandlerType: (*GoodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GoodsListQuery",
			Handler:    _Goods_GoodsListQuery_Handler,
		},
		{
			MethodName: "GetGoodsByIDs",
			Handler:    _Goods_GetGoodsByIDs_Handler,
		},
		{
			MethodName: "GetGoodsByID",
			Handler:    _Goods_GetGoodsByID_Handler,
		},
		{
			MethodName: "CreateGoods",
			Handler:    _Goods_CreateGoods_Handler,
		},
		{
			MethodName: "DeleteGoods",
			Handler:    _Goods_DeleteGoods_Handler,
		},
		{
			MethodName: "UpdateGoods",
			Handler:    _Goods_UpdateGoods_Handler,
		},
		{
			MethodName: "AllCategoryList",
			Handler:    _Goods_AllCategoryList_Handler,
		},
		{
			MethodName: "GetSubCategory",
			Handler:    _Goods_GetSubCategory_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _Goods_CreateCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _Goods_DeleteCategory_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _Goods_UpdateCategory_Handler,
		},
		{
			MethodName: "BrandList",
			Handler:    _Goods_BrandList_Handler,
		},
		{
			MethodName: "CreateBrand",
			Handler:    _Goods_CreateBrand_Handler,
		},
		{
			MethodName: "DeleteBrand",
			Handler:    _Goods_DeleteBrand_Handler,
		},
		{
			MethodName: "UpdateBrand",
			Handler:    _Goods_UpdateBrand_Handler,
		},
		{
			MethodName: "CategoryBrandList",
			Handler:    _Goods_CategoryBrandList_Handler,
		},
		{
			MethodName: "GetCategoryBrandByBrandID",
			Handler:    _Goods_GetCategoryBrandByBrandID_Handler,
		},
		{
			MethodName: "CreateCategoryBrand",
			Handler:    _Goods_CreateCategoryBrand_Handler,
		},
		{
			MethodName: "DeleteCategoryBrand",
			Handler:    _Goods_DeleteCategoryBrand_Handler,
		},
		{
			MethodName: "UpdateCategoryBrand",
			Handler:    _Goods_UpdateCategoryBrand_Handler,
		},
		{
			MethodName: "BannerList",
			Handler:    _Goods_BannerList_Handler,
		},
		{
			MethodName: "CreateBanner",
			Handler:    _Goods_CreateBanner_Handler,
		},
		{
			MethodName: "DeleteBanner",
			Handler:    _Goods_DeleteBanner_Handler,
		},
		{
			MethodName: "UpdateBanner",
			Handler:    _Goods_UpdateBanner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/goods.proto",
}
