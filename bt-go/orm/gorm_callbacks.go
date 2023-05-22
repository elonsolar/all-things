package main

// type DB struct {
//	 *Config
//   Error        error
//   RowsAffected int64
//   Statement    *Statement
//   clone        int
//}

//
// type callbacks struct {
//processors map[string]*processor
//}

//
// type processor struct {
//	db        *DB
//	Clauses   []string
//	fns       []func(*DB)
//	callbacks []*callback
// }//

//  type callback struct {
//  	name      string
//  	before    string
//  	after     string
//  	remove    bool
//  	replace   bool
//  	match     func(*DB) bool
//  	handler   func(*DB)
//  	processor *processor
//  }

// gorm  crud 都是通过call back完成

// 首先在各个驱动包（gorm封装的， 调用  callbacks/callback.go 里的 RegisterDefaultCallbacks 初始化

//  func  RegisterDefaultCallbacks(){
//          createCallback := db.Callback().Create()
//          createCallback.Match(enableTransaction).Register("gorm:begin_transaction", BeginTransaction)
//          createCallback.Register("gorm:before_create", BeforeCreate)
//          createCallback.Register("gorm:save_before_associations", SaveBeforeAssociations(true))
//          createCallback.Register("gorm:create", Create(config))
//          createCallback.Register("gorm:save_after_associations", SaveAfterAssociations(true))
//          createCallback.Register("gorm:after_create", AfterCreate)
//          createCallback.Match(enableTransaction).Register("gorm:commit_or_rollback_transaction", CommitOrRollbackTransaction)
//          createCallback.Clauses = config.CreateClauses

//
// 上面的方法 都注册 进 callbacks.processors["create"]  对应的processor 的callbacks 里了, 每个callback ,包含名称和对应的方法
//
// create,update,find  每一个操作 都有一个processor 对应

// fns ,是callback 排序后里面的func

// for _, f := range p.fns {
//		f(db)
// 	}
//
