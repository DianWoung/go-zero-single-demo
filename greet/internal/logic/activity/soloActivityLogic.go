package activity

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"
)

type SoloActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSoloActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SoloActivityLogic {
	return &SoloActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SoloActivityLogic) GetSoloActivityList(req *types.PageReq) (resp *types.Response, err error) {
	var total int64 = 1
	var list []svc.SoloActivityInfo

	/**
	 * 1、普通查询数据
	 * 1）svc.SoloActivity是model的表结构
	 * 2）将查询的数据映射到 list 中
	 */
	////`SELECT `solo_activity`.`id`,`solo_activity`.`title` FROM `solo_activity` LIMIT 10` 说明他会根据 list结构 查询指定字段，并不是*
	//query := l.svcCtx.Db.Model(&svc.SoloActivity{}).Select("school_id,title,id").Limit(10).Find(&list)

	//`SELECT `id`,`title`,`school_id` FROM `solo_activity` LIMIT 10` 也可以根据Select方法来指定字段
	query := l.svcCtx.Db.Model(&svc.SoloActivity{}).Select("school_id,title,id").Limit(10).Find(&list)

	/**
	 * 2、for update查询 （SELECT `solo_activity`.`id`,`solo_activity`.`title` FROM `solo_activity` LIMIT 10 FOR UPDATE）
	 */
	//query := l.svcCtx.Db.Model(&svc.SoloActivity{}).Clauses(clause.Locking{Strength: "UPDATE"}).Limit(10).Find(&list)

	/**
	 * 3、 where条件
	 * 1）使用map[string]interface{}{} 形式的命名参数
	 * 2）使用svc.SoloActivity 形式的命名参数
	 */
	//var id int = 20
	//var title string = "活动活动"
	////SELECT `solo_activity`.`id`,`solo_activity`.`title` FROM `solo_activity` WHERE title = '活动活动' OR id > 20 LIMIT 10
	//query := l.svcCtx.Db.Model(&svc.SoloActivity{}).
	//	Where("title = @title OR id > @id", map[string]interface{}{"title": title, "id": id}).
	//	Limit(10).Find(&list)
	//
	////query := l.svcCtx.Db.Model(&svc.SoloActivity{}).Where(svc.SoloActivity{Title: title}).Limit(10).Find(&list)

	/**
	 * 4、pluck用于从数据库查询单个列 + In查询
	 */
	//var idList []int
	//l.svcCtx.Db.Model(&svc.SoloActivity{}).Limit(10).Pluck("id", &idList)
	////fmt.Println(idList) //输出[9 10 11 12 22 27 33 37 71 153]
	//
	//query := l.svcCtx.Db.Model(&svc.SoloActivity{}).
	//	Where("title in @idList ", map[string]interface{}{"idList": idList}).
	//	Limit(10).Find(&list)

	/**
	 * 5、Group By
	 */
	//var id int = 20
	//var schoolCount []svc.ActivityGroupBySchoolCount
	//
	////SELECT school_id,COUNT(id) as count FROM `solo_activity` GROUP BY `school_id` HAVING AVG(id) > (20) LIMIT 10
	//query := l.svcCtx.Db.Model(&svc.SoloActivity{}).Select("school_id,COUNT(id) as count").
	//	Group("school_id").
	//	Having("AVG(id) > (?)", id).Limit(10).Find(&schoolCount)
	//fmt.Println(schoolCount) //[{1 103} {2 15} {3 26} {15 23} {22 2} {30 2} {31 1} {33 1}]

	if err := query.Error; err != nil {
		return &types.Response{
			Code: -1,
			Msg:  "获取失败",
		}, nil
	}

	return &types.Response{
		Code: 0,
		Msg:  "Hello go-zero",
		Data: types.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
			List:     list,
		},
	}, nil
}
