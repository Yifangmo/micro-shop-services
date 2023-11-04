package service

// import (
// 	"context"
// 	"fmt"

// 	goredislib "github.com/go-redis/redis/v8"
// 	"github.com/go-redsync/redsync/v4"
// 	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// 	"google.golang.org/protobuf/types/known/emptypb"

// 	"github.com/Yifangmo/micro-shop-services/inventory/global"
// 	"github.com/Yifangmo/micro-shop-services/inventory/models"
// 	"github.com/Yifangmo/micro-shop-services/inventory/proto"
// )

// func (*InventoryServer) TrySell(ctx context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
// 	//扣减库存， 本地事务 [1:10,  2:5, 3: 20]
// 	//数据库基本的一个应用场景：数据库事务
// 	//并发情况之下 可能会出现超卖 1
// 	client := goredislib.NewClient(&goredislib.Options{
// 		Addr: "192.168.1.10:6379",
// 	})
// 	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)
// 	rs := redsync.New(pool)

// 	tx := global.DB.Begin()
// 	//m.Lock() //获取锁 这把锁有问题吗？  假设有10w的并发， 这里并不是请求的同一件商品  这个锁就没有问题了吗？
// 	for _, goodInfo := range req.GoodsInfo {
// 		var inv models.InventoryNew
// 		//if result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where(&model.Inventory{Goods:goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
// 		//	tx.Rollback() //回滚之前的操作
// 		//	return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
// 		//}

// 		//for {
// 		mutex := rs.NewMutex(fmt.Sprintf("goods_%d", goodInfo.GoodsId))
// 		if err := mutex.Lock(); err != nil {
// 			return nil, status.Errorf(codes.Internal, "获取redis分布式锁异常")
// 		}

// 		if result := global.DB.Where(&models.Inventory{Goods: goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
// 			tx.Rollback() //回滚之前的操作
// 			return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
// 		}
// 		//判断库存是否充足
// 		if inv.Stocks < goodInfo.Num {
// 			tx.Rollback() //回滚之前的操作
// 			return nil, status.Errorf(codes.ResourceExhausted, "库存不足")
// 		}
// 		//扣减， 会出现数据不一致的问题 - 锁，分布式锁
// 		//inv.Stocks -= goodInfo.Num
// 		inv.Freeze += goodInfo.Num
// 		tx.Save(&inv)

// 		if ok, err := mutex.Unlock(); !ok || err != nil {
// 			return nil, status.Errorf(codes.Internal, "释放redis分布式锁异常")
// 		}
// 		//update inventory set stocks = stocks-1, version=version+1 where goods=goods and version=version
// 		//这种写法有瑕疵，为什么？
// 		//零值 对于int类型来说 默认值是0 这种会被gorm给忽略掉
// 		//if result := tx.Model(&model.Inventory{}).Select("Stocks", "Version").Where("goods = ? and version= ?", goodInfo.GoodsId, inv.Version).Updates(model.Inventory{Stocks: inv.Stocks, Version: inv.Version+1}); result.RowsAffected == 0 {
// 		//	zap.S().Info("库存扣减失败")
// 		//}else{
// 		//	break
// 		//}
// 		//}
// 		//tx.Save(&inv)
// 	}
// 	tx.Commit() // 需要自己手动提交操作
// 	//m.Unlock() //释放锁
// 	return &emptypb.Empty{}, nil
// }

// func (*InventoryServer) ConfirmSell(ctx context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
// 	//扣减库存， 本地事务 [1:10,  2:5, 3: 20]
// 	//数据库基本的一个应用场景：数据库事务
// 	//并发情况之下 可能会出现超卖 1
// 	client := goredislib.NewClient(&goredislib.Options{
// 		Addr: "192.168.1.10:6379",
// 	})
// 	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)
// 	rs := redsync.New(pool)

// 	tx := global.DB.Begin()
// 	//m.Lock() //获取锁 这把锁有问题吗？  假设有10w的并发， 这里并不是请求的同一件商品  这个锁就没有问题了吗？
// 	for _, goodInfo := range req.GoodsInfo {
// 		var inv models.InventoryNew
// 		//if result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where(&model.Inventory{Goods:goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
// 		//	tx.Rollback() //回滚之前的操作
// 		//	return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
// 		//}

// 		//for {
// 		mutex := rs.NewMutex(fmt.Sprintf("goods_%d", goodInfo.GoodsId))
// 		if err := mutex.Lock(); err != nil {
// 			return nil, status.Errorf(codes.Internal, "获取redis分布式锁异常")
// 		}

// 		if result := global.DB.Where(&models.Inventory{Goods: goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
// 			tx.Rollback() //回滚之前的操作
// 			return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
// 		}
// 		//判断库存是否充足
// 		if inv.Stocks < goodInfo.Num {
// 			tx.Rollback() //回滚之前的操作
// 			return nil, status.Errorf(codes.ResourceExhausted, "库存不足")
// 		}
// 		//扣减， 会出现数据不一致的问题 - 锁，分布式锁
// 		inv.Stocks -= goodInfo.Num
// 		inv.Freeze -= goodInfo.Num
// 		tx.Save(&inv)

// 		if ok, err := mutex.Unlock(); !ok || err != nil {
// 			return nil, status.Errorf(codes.Internal, "释放redis分布式锁异常")
// 		}
// 		//update inventory set stocks = stocks-1, version=version+1 where goods=goods and version=version
// 		//这种写法有瑕疵，为什么？
// 		//零值 对于int类型来说 默认值是0 这种会被gorm给忽略掉
// 		//if result := tx.Model(&model.Inventory{}).Select("Stocks", "Version").Where("goods = ? and version= ?", goodInfo.GoodsId, inv.Version).Updates(model.Inventory{Stocks: inv.Stocks, Version: inv.Version+1}); result.RowsAffected == 0 {
// 		//	zap.S().Info("库存扣减失败")
// 		//}else{
// 		//	break
// 		//}
// 		//}
// 		//tx.Save(&inv)
// 	}
// 	tx.Commit() // 需要自己手动提交操作
// 	//m.Unlock() //释放锁
// 	return &emptypb.Empty{}, nil
// }

// func (*InventoryServer) CancelSell(ctx context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
// 	//扣减库存， 本地事务 [1:10,  2:5, 3: 20]
// 	//数据库基本的一个应用场景：数据库事务
// 	//并发情况之下 可能会出现超卖 1
// 	client := goredislib.NewClient(&goredislib.Options{
// 		Addr: "192.168.1.10:6379",
// 	})
// 	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)
// 	rs := redsync.New(pool)

// 	tx := global.DB.Begin()
// 	//m.Lock() //获取锁 这把锁有问题吗？  假设有10w的并发， 这里并不是请求的同一件商品  这个锁就没有问题了吗？
// 	for _, goodInfo := range req.GoodsInfo {
// 		var inv models.InventoryNew
// 		//if result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where(&model.Inventory{Goods:goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
// 		//	tx.Rollback() //回滚之前的操作
// 		//	return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
// 		//}

// 		//for {
// 		mutex := rs.NewMutex(fmt.Sprintf("goods_%d", goodInfo.GoodsId))
// 		if err := mutex.Lock(); err != nil {
// 			return nil, status.Errorf(codes.Internal, "获取redis分布式锁异常")
// 		}

// 		if result := global.DB.Where(&models.Inventory{Goods: goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
// 			tx.Rollback() //回滚之前的操作
// 			return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
// 		}
// 		//判断库存是否充足
// 		if inv.Stocks < goodInfo.Num {
// 			tx.Rollback() //回滚之前的操作
// 			return nil, status.Errorf(codes.ResourceExhausted, "库存不足")
// 		}
// 		//扣减， 会出现数据不一致的问题 - 锁，分布式锁
// 		inv.Freeze -= goodInfo.Num
// 		tx.Save(&inv)

// 		if ok, err := mutex.Unlock(); !ok || err != nil {
// 			return nil, status.Errorf(codes.Internal, "释放redis分布式锁异常")
// 		}
// 		//update inventory set stocks = stocks-1, version=version+1 where goods=goods and version=version
// 		//这种写法有瑕疵，为什么？
// 		//零值 对于int类型来说 默认值是0 这种会被gorm给忽略掉
// 		//if result := tx.Model(&model.Inventory{}).Select("Stocks", "Version").Where("goods = ? and version= ?", goodInfo.GoodsId, inv.Version).Updates(model.Inventory{Stocks: inv.Stocks, Version: inv.Version+1}); result.RowsAffected == 0 {
// 		//	zap.S().Info("库存扣减失败")
// 		//}else{
// 		//	break
// 		//}
// 		//}
// 		//tx.Save(&inv)
// 	}
// 	tx.Commit() // 需要自己手动提交操作
// 	//m.Unlock() //释放锁
// 	return &emptypb.Empty{}, nil
// }
