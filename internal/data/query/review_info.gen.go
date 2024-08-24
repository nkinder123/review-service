// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"review-server/internal/data/model"
)

func newReviewInfo(db *gorm.DB, opts ...gen.DOOption) reviewInfo {
	_reviewInfo := reviewInfo{}

	_reviewInfo.reviewInfoDo.UseDB(db, opts...)
	_reviewInfo.reviewInfoDo.UseModel(&model.ReviewInfo{})

	tableName := _reviewInfo.reviewInfoDo.TableName()
	_reviewInfo.ALL = field.NewAsterisk(tableName)
	_reviewInfo.ID = field.NewInt64(tableName, "id")
	_reviewInfo.CreateBy = field.NewString(tableName, "create_by")
	_reviewInfo.UpdateBy = field.NewString(tableName, "update_by")
	_reviewInfo.CreateAt = field.NewTime(tableName, "create_at")
	_reviewInfo.UpdateAt = field.NewTime(tableName, "update_at")
	_reviewInfo.DeleteAt = field.NewTime(tableName, "delete_at")
	_reviewInfo.Version = field.NewInt32(tableName, "version")
	_reviewInfo.ReviewID = field.NewInt64(tableName, "review_id")
	_reviewInfo.Content = field.NewString(tableName, "content")
	_reviewInfo.Score = field.NewInt32(tableName, "score")
	_reviewInfo.ServiceScore = field.NewInt32(tableName, "service_score")
	_reviewInfo.ExpressScore = field.NewInt32(tableName, "express_score")
	_reviewInfo.HasMedia = field.NewInt32(tableName, "has_media")
	_reviewInfo.OrderID = field.NewInt64(tableName, "order_id")
	_reviewInfo.SkuID = field.NewInt64(tableName, "sku_id")
	_reviewInfo.SpuID = field.NewInt64(tableName, "spu_id")
	_reviewInfo.StoreID = field.NewInt64(tableName, "store_id")
	_reviewInfo.UserID = field.NewInt64(tableName, "user_id")
	_reviewInfo.Anonymous = field.NewInt32(tableName, "anonymous")
	_reviewInfo.Tags = field.NewString(tableName, "tags")
	_reviewInfo.PicInfo = field.NewString(tableName, "pic_info")
	_reviewInfo.VideoInfo = field.NewString(tableName, "video_info")
	_reviewInfo.Status = field.NewInt32(tableName, "status")
	_reviewInfo.IsDefault = field.NewInt32(tableName, "is_default")
	_reviewInfo.HasReply = field.NewInt32(tableName, "has_reply")
	_reviewInfo.OpReason = field.NewString(tableName, "op_reason")
	_reviewInfo.OpRemarks = field.NewString(tableName, "op_remarks")
	_reviewInfo.OpUser = field.NewString(tableName, "op_user")
	_reviewInfo.GoodsSnapshoot = field.NewString(tableName, "goods_snapshoot")
	_reviewInfo.ExtJSON = field.NewString(tableName, "ext_json")
	_reviewInfo.CtrlJSON = field.NewString(tableName, "ctrl_json")

	_reviewInfo.fillFieldMap()

	return _reviewInfo
}

type reviewInfo struct {
	reviewInfoDo reviewInfoDo

	ALL            field.Asterisk
	ID             field.Int64  // 主键
	CreateBy       field.String // 创建方标识
	UpdateBy       field.String // 更新方标识
	CreateAt       field.Time   // 创建时间
	UpdateAt       field.Time   // 更新时间
	DeleteAt       field.Time   // 逻辑删除标记
	Version        field.Int32  // 乐观锁标记
	ReviewID       field.Int64  // 评价id
	Content        field.String // 评价内容
	Score          field.Int32  // 评分
	ServiceScore   field.Int32  // 商家服务评分
	ExpressScore   field.Int32  // 物流评分
	HasMedia       field.Int32  // 是否有图或视频
	OrderID        field.Int64  // 订单id
	SkuID          field.Int64  // sku id
	SpuID          field.Int64  // spu id
	StoreID        field.Int64  // 店铺id
	UserID         field.Int64  // 用户id
	Anonymous      field.Int32  // 是否匿名
	Tags           field.String // 标签json
	PicInfo        field.String // 媒体信息：图片
	VideoInfo      field.String // 媒体信息：视频
	Status         field.Int32  // 状态:10待审核；20审核通过；30审核不通过；40隐藏
	IsDefault      field.Int32  // 是否默认评价
	HasReply       field.Int32  // 是否有商家回复:0无;1有
	OpReason       field.String // 运营审核拒绝原因
	OpRemarks      field.String // 运营备注
	OpUser         field.String // 运营者标识
	GoodsSnapshoot field.String // 商品快照信息
	ExtJSON        field.String // 信息扩展
	CtrlJSON       field.String // 控制扩展

	fieldMap map[string]field.Expr
}

func (r reviewInfo) Table(newTableName string) *reviewInfo {
	r.reviewInfoDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r reviewInfo) As(alias string) *reviewInfo {
	r.reviewInfoDo.DO = *(r.reviewInfoDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *reviewInfo) updateTableName(table string) *reviewInfo {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.CreateBy = field.NewString(table, "create_by")
	r.UpdateBy = field.NewString(table, "update_by")
	r.CreateAt = field.NewTime(table, "create_at")
	r.UpdateAt = field.NewTime(table, "update_at")
	r.DeleteAt = field.NewTime(table, "delete_at")
	r.Version = field.NewInt32(table, "version")
	r.ReviewID = field.NewInt64(table, "review_id")
	r.Content = field.NewString(table, "content")
	r.Score = field.NewInt32(table, "score")
	r.ServiceScore = field.NewInt32(table, "service_score")
	r.ExpressScore = field.NewInt32(table, "express_score")
	r.HasMedia = field.NewInt32(table, "has_media")
	r.OrderID = field.NewInt64(table, "order_id")
	r.SkuID = field.NewInt64(table, "sku_id")
	r.SpuID = field.NewInt64(table, "spu_id")
	r.StoreID = field.NewInt64(table, "store_id")
	r.UserID = field.NewInt64(table, "user_id")
	r.Anonymous = field.NewInt32(table, "anonymous")
	r.Tags = field.NewString(table, "tags")
	r.PicInfo = field.NewString(table, "pic_info")
	r.VideoInfo = field.NewString(table, "video_info")
	r.Status = field.NewInt32(table, "status")
	r.IsDefault = field.NewInt32(table, "is_default")
	r.HasReply = field.NewInt32(table, "has_reply")
	r.OpReason = field.NewString(table, "op_reason")
	r.OpRemarks = field.NewString(table, "op_remarks")
	r.OpUser = field.NewString(table, "op_user")
	r.GoodsSnapshoot = field.NewString(table, "goods_snapshoot")
	r.ExtJSON = field.NewString(table, "ext_json")
	r.CtrlJSON = field.NewString(table, "ctrl_json")

	r.fillFieldMap()

	return r
}

func (r *reviewInfo) WithContext(ctx context.Context) IReviewInfoDo {
	return r.reviewInfoDo.WithContext(ctx)
}

func (r reviewInfo) TableName() string { return r.reviewInfoDo.TableName() }

func (r reviewInfo) Alias() string { return r.reviewInfoDo.Alias() }

func (r reviewInfo) Columns(cols ...field.Expr) gen.Columns { return r.reviewInfoDo.Columns(cols...) }

func (r *reviewInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *reviewInfo) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 31)
	r.fieldMap["id"] = r.ID
	r.fieldMap["create_by"] = r.CreateBy
	r.fieldMap["update_by"] = r.UpdateBy
	r.fieldMap["create_at"] = r.CreateAt
	r.fieldMap["update_at"] = r.UpdateAt
	r.fieldMap["delete_at"] = r.DeleteAt
	r.fieldMap["version"] = r.Version
	r.fieldMap["review_id"] = r.ReviewID
	r.fieldMap["content"] = r.Content
	r.fieldMap["score"] = r.Score
	r.fieldMap["service_score"] = r.ServiceScore
	r.fieldMap["express_score"] = r.ExpressScore
	r.fieldMap["has_media"] = r.HasMedia
	r.fieldMap["order_id"] = r.OrderID
	r.fieldMap["sku_id"] = r.SkuID
	r.fieldMap["spu_id"] = r.SpuID
	r.fieldMap["store_id"] = r.StoreID
	r.fieldMap["user_id"] = r.UserID
	r.fieldMap["anonymous"] = r.Anonymous
	r.fieldMap["tags"] = r.Tags
	r.fieldMap["pic_info"] = r.PicInfo
	r.fieldMap["video_info"] = r.VideoInfo
	r.fieldMap["status"] = r.Status
	r.fieldMap["is_default"] = r.IsDefault
	r.fieldMap["has_reply"] = r.HasReply
	r.fieldMap["op_reason"] = r.OpReason
	r.fieldMap["op_remarks"] = r.OpRemarks
	r.fieldMap["op_user"] = r.OpUser
	r.fieldMap["goods_snapshoot"] = r.GoodsSnapshoot
	r.fieldMap["ext_json"] = r.ExtJSON
	r.fieldMap["ctrl_json"] = r.CtrlJSON
}

func (r reviewInfo) clone(db *gorm.DB) reviewInfo {
	r.reviewInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r reviewInfo) replaceDB(db *gorm.DB) reviewInfo {
	r.reviewInfoDo.ReplaceDB(db)
	return r
}

type reviewInfoDo struct{ gen.DO }

type IReviewInfoDo interface {
	gen.SubQuery
	Debug() IReviewInfoDo
	WithContext(ctx context.Context) IReviewInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IReviewInfoDo
	WriteDB() IReviewInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IReviewInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IReviewInfoDo
	Not(conds ...gen.Condition) IReviewInfoDo
	Or(conds ...gen.Condition) IReviewInfoDo
	Select(conds ...field.Expr) IReviewInfoDo
	Where(conds ...gen.Condition) IReviewInfoDo
	Order(conds ...field.Expr) IReviewInfoDo
	Distinct(cols ...field.Expr) IReviewInfoDo
	Omit(cols ...field.Expr) IReviewInfoDo
	Join(table schema.Tabler, on ...field.Expr) IReviewInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IReviewInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IReviewInfoDo
	Group(cols ...field.Expr) IReviewInfoDo
	Having(conds ...gen.Condition) IReviewInfoDo
	Limit(limit int) IReviewInfoDo
	Offset(offset int) IReviewInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IReviewInfoDo
	Unscoped() IReviewInfoDo
	Create(values ...*model.ReviewInfo) error
	CreateInBatches(values []*model.ReviewInfo, batchSize int) error
	Save(values ...*model.ReviewInfo) error
	First() (*model.ReviewInfo, error)
	Take() (*model.ReviewInfo, error)
	Last() (*model.ReviewInfo, error)
	Find() ([]*model.ReviewInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ReviewInfo, err error)
	FindInBatches(result *[]*model.ReviewInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ReviewInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IReviewInfoDo
	Assign(attrs ...field.AssignExpr) IReviewInfoDo
	Joins(fields ...field.RelationField) IReviewInfoDo
	Preload(fields ...field.RelationField) IReviewInfoDo
	FirstOrInit() (*model.ReviewInfo, error)
	FirstOrCreate() (*model.ReviewInfo, error)
	FindByPage(offset int, limit int) (result []*model.ReviewInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IReviewInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r reviewInfoDo) Debug() IReviewInfoDo {
	return r.withDO(r.DO.Debug())
}

func (r reviewInfoDo) WithContext(ctx context.Context) IReviewInfoDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r reviewInfoDo) ReadDB() IReviewInfoDo {
	return r.Clauses(dbresolver.Read)
}

func (r reviewInfoDo) WriteDB() IReviewInfoDo {
	return r.Clauses(dbresolver.Write)
}

func (r reviewInfoDo) Session(config *gorm.Session) IReviewInfoDo {
	return r.withDO(r.DO.Session(config))
}

func (r reviewInfoDo) Clauses(conds ...clause.Expression) IReviewInfoDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r reviewInfoDo) Returning(value interface{}, columns ...string) IReviewInfoDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r reviewInfoDo) Not(conds ...gen.Condition) IReviewInfoDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r reviewInfoDo) Or(conds ...gen.Condition) IReviewInfoDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r reviewInfoDo) Select(conds ...field.Expr) IReviewInfoDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r reviewInfoDo) Where(conds ...gen.Condition) IReviewInfoDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r reviewInfoDo) Order(conds ...field.Expr) IReviewInfoDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r reviewInfoDo) Distinct(cols ...field.Expr) IReviewInfoDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r reviewInfoDo) Omit(cols ...field.Expr) IReviewInfoDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r reviewInfoDo) Join(table schema.Tabler, on ...field.Expr) IReviewInfoDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r reviewInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IReviewInfoDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r reviewInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IReviewInfoDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r reviewInfoDo) Group(cols ...field.Expr) IReviewInfoDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r reviewInfoDo) Having(conds ...gen.Condition) IReviewInfoDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r reviewInfoDo) Limit(limit int) IReviewInfoDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r reviewInfoDo) Offset(offset int) IReviewInfoDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r reviewInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IReviewInfoDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r reviewInfoDo) Unscoped() IReviewInfoDo {
	return r.withDO(r.DO.Unscoped())
}

func (r reviewInfoDo) Create(values ...*model.ReviewInfo) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r reviewInfoDo) CreateInBatches(values []*model.ReviewInfo, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r reviewInfoDo) Save(values ...*model.ReviewInfo) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r reviewInfoDo) First() (*model.ReviewInfo, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReviewInfo), nil
	}
}

func (r reviewInfoDo) Take() (*model.ReviewInfo, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReviewInfo), nil
	}
}

func (r reviewInfoDo) Last() (*model.ReviewInfo, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReviewInfo), nil
	}
}

func (r reviewInfoDo) Find() ([]*model.ReviewInfo, error) {
	result, err := r.DO.Find()
	return result.([]*model.ReviewInfo), err
}

func (r reviewInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ReviewInfo, err error) {
	buf := make([]*model.ReviewInfo, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r reviewInfoDo) FindInBatches(result *[]*model.ReviewInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r reviewInfoDo) Attrs(attrs ...field.AssignExpr) IReviewInfoDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r reviewInfoDo) Assign(attrs ...field.AssignExpr) IReviewInfoDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r reviewInfoDo) Joins(fields ...field.RelationField) IReviewInfoDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r reviewInfoDo) Preload(fields ...field.RelationField) IReviewInfoDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r reviewInfoDo) FirstOrInit() (*model.ReviewInfo, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReviewInfo), nil
	}
}

func (r reviewInfoDo) FirstOrCreate() (*model.ReviewInfo, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReviewInfo), nil
	}
}

func (r reviewInfoDo) FindByPage(offset int, limit int) (result []*model.ReviewInfo, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r reviewInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r reviewInfoDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r reviewInfoDo) Delete(models ...*model.ReviewInfo) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *reviewInfoDo) withDO(do gen.Dao) *reviewInfoDo {
	r.DO = *do.(*gen.DO)
	return r
}
