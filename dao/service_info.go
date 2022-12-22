package dao

import (
	"go_gateway/dto"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServiceInfo struct {
	ID          int64     `json:"id" gorm:"primary_key" description:"自增主键"`
	LoadType    int       `json:"load_type" gorm:"column:load_type" description:"负载类型 0-http 1-tcp 2-grpc"`
	ServiceName string    `json:"service_name" gorm:"column:service_name" description:"服务名称"`
	ServiceDesc string    `json:"service_desc" gorm:"column:service_desc" description:"服务描述"`
	UpdatedAt   time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt   time.Time `json:"create_at" gorm:"column:create_at" description:"添加时间"`
	IsDelete    int8      `json:"is_delete" gorm:"column:is_delete" description:"是否删除；0:否，1:是"`
}

func (t *ServiceInfo) TableName() string {
	return "gateway_service_info"
}

func (t *ServiceInfo) Find(ctx *gin.Context, tx *gorm.DB, search *ServiceInfo) (*ServiceInfo, error) {
	out := &ServiceInfo{}
	// err := tx.SetCtx(public.GetGinTraceContext(ctx)).Where(search).Find(out).Error
	err := tx.WithContext(ctx).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (t *ServiceInfo) Save(ctx *gin.Context, tx *gorm.DB) error {
	// err := tx.SetCtx(public.GetGinTraceContext(ctx)).Where(search).Find(out).Error
	return tx.WithContext(ctx).Save(t).Error
}

func (t ServiceInfo) ServiceDetail(ctx *gin.Context, tx *gorm.DB, search *ServiceInfo) (*ServiceDetail, error) {
	httpRule := &HttpRule{ServiceID: search.ID}
	httpRule, err := httpRule.Find(ctx, tx, httpRule)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	tcpRule := &TcpRule{ServiceID: search.ID}
	tcpRule, err = tcpRule.Find(ctx, tx, tcpRule)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	grpcRule := &GrpcRule{ServiceID: search.ID}
	grpcRule, err = grpcRule.Find(ctx, tx, grpcRule)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	accessControl := &AccessControl{ServiceID: search.ID}
	accessControl, err = accessControl.Find(ctx, tx, accessControl)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	loadBalance := &LoadBalance{ServiceID: search.ID}
	loadBalance, err = loadBalance.Find(ctx, tx, loadBalance)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	detail := &ServiceDetail{
		Info:          search,
		HTTPRule:      httpRule,
		TCPRule:       tcpRule,
		GRPCRule:      grpcRule,
		LoadBalance:   loadBalance,
		AccessControl: accessControl,
	}
	return detail, nil
}

// PageList 分页
func (t *ServiceInfo) PageList(ctx *gin.Context, tx *gorm.DB, param *dto.ServiceListInput) ([]ServiceInfo, int64, error) {
	total := int64(0)
	// list := []ServiceInfo{}
	var list []ServiceInfo
	offset := (param.PageNo - 1) * param.PageSize
	query := tx.WithContext(ctx)
	query = query.Table(t.TableName()).Where("is_delete=0")
	if param.Info != "" {
		query = query.Where("(service_name like ? or service_desc like ?)", "%"+param.Info+"%", "%"+param.Info+"%")
	}
	err := query.Limit(param.PageSize).Offset(offset).Order("id desc").Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	query.Limit(param.PageSize).Offset(offset).Count(&total)
	return list, total, nil
}
