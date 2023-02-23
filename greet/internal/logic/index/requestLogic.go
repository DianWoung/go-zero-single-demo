package index

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/svc/constant"
	"go-zero-demo/greet/internal/types"
	"go-zero-demo/greet/utils"
	"gorm.io/gorm"
	"time"
)

type RequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RequestLogic {
	return &RequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RequestLogic) GetCourseList(req *types.PageReq) (resp *types.Response, err error) {
	var total int64 = 1
	var list1 []svc.CourseSnapshot
	var list2 []svc.CourseSnapshotInfo
	var list3 []svc.CourseSnapshotInfo

	jwtToken, err := utils.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, map[string]interface{}{
		"name": "huangbin",
		"age":  "23",
	})
	if err != nil {
		return types.NewDefaultError(err.Error()), nil
	}

	println("ttoken")
	println(jwtToken)

	mycaim, err := utils.ParseToken(l.svcCtx.Config.Auth.AccessSecret, jwtToken)

	if err != nil {
		return types.NewDefaultError(err.Error()), nil
	}

	// fmt.Printf("%#v", mycaim)
	fmt.Println("rrrrrrr")
	fmt.Printf("%#v", mycaim)

	res := mycaim["data"].(map[string]interface{})

	fmt.Println(res["age"])
	fmt.Println(res["name"])
	//fmt.Println(mycaim.data)
	//fmt.Println(data.)

	//-1、获取常量信息
	println(constant.ACTIVITY_FAMILY)

	//0、错误结果抛出
	//return types.NewDefaultError("默认错误信息"), nil
	//return types.NewCodeError(-2,"默认错误信息"), nil

	/**
	 * 1、N:N查询
	 * 1）N:N关系定义在struct中
	 * 2）Find查询结果是列表，First查询的是单条数据
	 */
	query1 := l.svcCtx.Db.Model(&svc.CourseSnapshot{}).
		Preload("CourseStudentList", func(db *gorm.DB) *gorm.DB {
			return db.Where("id > @id", map[string]interface{}{"id": 3774}).
				Select("id,student_id,course_id,course_snapshot_id").Order("id asc")
		}).
		Preload("CourseInfo").
		Preload("CourseCitingInfo").
		Select("id,course_id,title,cover_image,class_hour").
		Order("course_id desc").Limit(10).Find(&list1)

	if err := query1.Error; err != nil {
		return types.NewDefaultError("默认错误信息"), nil
	}

	/**
	 * 2、left join查询 + 链式操作
	 */
	query2 := l.svcCtx.Db.Model(&svc.CourseSnapshot{}).
		Joins("left join course on course.id = course_snapshot.course_id where course_snapshot.course_id > ?", 0).
		Select("" +
			"course_snapshot.id,course_snapshot.course_id,course_snapshot.title,course_snapshot.cover_image,course_snapshot.class_hour" +
			",course.type,course.is_default,course.is_cloud,course.course_no").
		Order("course_id desc")

	if 1 == 1 {
		query2.Limit(10).Find(&list2)
	}

	if err := query2.Error; err != nil {
		return types.NewDefaultError("默认错误信息"), nil
	}

	/**
	 * 3、分页
	 */
	query3 := l.svcCtx.Db.Model(&svc.CourseSnapshot{}).
		Joins("left join course on course.id = course_snapshot.course_id where course_snapshot.course_id > ?", 0)

	var count int64
	var page int = 2
	var pageSize int = 2
	query3.Count(&count) //总行数

	query3.Select("" +
		"course_snapshot.id,course_snapshot.course_id,course_snapshot.title,course_snapshot.cover_image,course_snapshot.class_hour" +
		",course.type,course.is_default,course.is_cloud,course.course_no").
		Order("course_id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list3)

	if err := query3.Error; err != nil {
		return types.NewDefaultError("默认错误信息"), nil
	}

	return &types.Response{
		Code: 0,
		Msg:  "Hello go-zero",
		Data: types.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
			List: map[string]interface{}{
				"list1": list1,
				"list2": list2,
				"list3": list3,
			},
		},
	}, nil
}

func (l *RequestLogic) TestJwtToken(req *types.PageReq) (resp *types.Response, err error) {
	err = l.svcCtx.Rdb.Set("liu", "huiling", time.Duration(600)*time.Second).Err()
	if err != nil {
		fmt.Println("设置key失败")
		return
	}

	jwtToken, err := utils.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, map[string]interface{}{
		"name": "huangbin123",
		"age":  "18",
	})
	if err != nil {
		return types.NewDefaultError(err.Error()), nil
	}
	println("token")
	println(jwtToken)

	mycaim, err := utils.ParseToken(l.svcCtx.Config.Auth.AccessSecret, jwtToken)

	if err != nil {
		return types.NewDefaultError(err.Error()), nil
	}

	fmt.Printf("%#v", mycaim)
	result := mycaim["data"].(map[string]interface{})

	return &types.Response{
		Code: 0,
		Msg:  "Hello go-zero",
		Data: map[string]string{
			"token": jwtToken,
			"age":   result["age"].(string),
			"name":  result["name"].(string),
		},
	}, nil
}
