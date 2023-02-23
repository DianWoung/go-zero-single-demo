package svc

import (
	"database/sql"
	"time"
)

//模型定义文档 https://gorm.io/zh_CN/docs/models.html

type SoloActivity struct {
	ID               int    `gorm:"not null;unique;primarykey"`
	SchoolTermId     int    `json:"schoolTermId"  gorm:"comment:学期Id"`
	SchoolId         int    `json:"schoolId" gorm:"comment:学校Id"`
	StudentId        int    `json:"studentId" gorm:"comment:学生Id"`
	Title            string `json:"title" gorm:"comment:活动标题"`
	Content          string `json:"content" gorm:"comment:活动内容"`
	ClassHour        int    `json:"classHour" gorm:"comment:活动课时"`
	Archived         int    `json:"archived"`
	Grade            int    `json:"grade"  gorm:"comment:活动课时"`
	Class            int    `json:"class"  gorm:"comment:活动课时"`
	SchoolClassId    int    `json:"schoolClassId" gorm:"comment:班级id"`
	Type             int    `json:"type" gorm:"default:0;comment:活动类型 4-家庭活动 5-课后活动"`
	ActivityType     int    `json:"activityType" gorm:"default:0;comment:活动课程类型：1-实践课程，2-劳动课程"`
	ActivityWay      int    `json:"activityWay" gorm:"default:0;comment:活动方式：1-考察探究，2-社会服务，3-设计制作，4-职业体验，5-日常生活劳动，6-服务劳动，7-生产劳动"`
	ActivityTask     int    `json:"activityTask" gorm:"default:0;comment:劳动任务群：1、清洁与卫生 2、整理与收纳 3、烹饪与营养 4、家用器具使用于维护 5、农业生产劳动 6、传统工艺制作 7、工业生产劳动 8、新技术体验与应用 "`
	IsReject         int    `json:"isReject" gorm:"default:0;comment:是否驳回0否 1是"`
	RejectMessage    int    `json:"rejectMessage" gorm:"default:0;comment:驳回内容"`
	ReadStatus       int    `json:"readStatus" gorm:"default:0;comment:活动读状态 0未读 1已读"`
	RecordReadStatus int    `json:"recordReadStatus" gorm:"default:0;comment:记录都状态 0未读 1已读"`
	StartAt          time.Time
	EndAt            time.Time
}

type Course struct {
	Id               int64        `json:"id"`
	CourseSnapshotId int64        `json:"course_snapshot_id"` // 课程快照id
	AdminId          int64        `json:"admin_id"`           // 管理员账号id
	TeacherId        int64        `json:"teacher_id"`         // 教师/教官id
	InstitutionId    int64        `json:"institution_id"`     // 机构id
	Type             int64        `json:"type"`               // 账号所属的机构类型,1:触点管理员，2：第三方机构，3：学校，4：教育局，5：运营公司
	IsDefault        int64        `json:"is_default"`         // 是否默认课程：1-是，0-否
	IsCloud          int64        `json:"is_cloud"`           // 是否云课程1不是2是
	CourseNo         string       `json:"course_no"`          // 课程编号
	CreatedAt        sql.NullTime `json:"created_at"`         // 创建时间
	UpdatedAt        sql.NullTime `json:"updated_at"`         // 更新时间
	DeletedAt        sql.NullTime `json:"deleted_at"`         // 删除时间
}

type CourseSnapshot struct {
	/**
	 * FOREIGN KEY 和 REFERENCES区别
	 *	假设两张表，表1(学号,姓名,性别),学号为主键；表2(学号,课程,成绩)。
	 *	可以为表2的学号定义外键(FOREIGN KEY)，该外键的取值范围参照(REFERENCES)表1的学号
	 */

	//HasMany关系：使用`[]CourseStudent`表示；外键为`course_student`表的 course_snapshot_id，引用为`表的course_snapshot`表的 id
	CourseStudentList []CourseStudent `gorm:"foreignKey:CourseSnapshotId;references:Id;"`

	//Belong to关系：外键为`表的course_snapshot`表的 course_id，引用为`表的course`表的 id
	CourseInfo Course `gorm:"foreignKey:CourseId;references:Id"` // 外键指的是`引用表的外键`，即`CourseStudent表`的外键

	//Has one关系：外键为`表的course_citing`表的 course_snapshot_id，引用为`表的course_snapshot`表的 id
	CourseCitingInfo CourseCiting `gorm:"foreignKey:CourseSnapshotId;references:Id"`

	Id                 int64        `json:"id"`
	CourseId           int64        `json:"course_id"`           // 课程id
	Title              string       `json:"title"`               // 课程名称
	ClassHour          int64        `json:"class_hour"`          // 课时
	Type               int64        `json:"type"`                // 课程类型：1-实践课程，2-劳动课程
	ActivityWay        int64        `json:"activity_way"`        // 活动方式：1-考察探究，2-社会服务，3-设计制作，4-职业体验，5-日常生活劳动，6-服务劳动，7-生成劳动
	ActivityTask       int64        `json:"activity_task"`       // 劳动任务群：1、清洁与卫生 2、整理与收纳 3、烹饪与营养 4、家用器具使用于维护 5、农业生产劳动 6、传统工艺制作 7、工业生产劳动 8、新技术体验与应用 9、现代服务业劳动 10、公益劳动与志愿服务
	CourseContent      string       `json:"course_content"`      // 课程内容
	Address            string       `json:"address"`             // 实施地点
	SignatureIntroduce string       `json:"signature_introduce"` // 署名介绍
	CoverImage         string       `json:"cover_image"`         // 封面图片
	CreatedAt          sql.NullTime `json:"created_at"`
	UpdatedAt          sql.NullTime `json:"updated_at"`
	DeletedAt          sql.NullTime `json:"deleted_at"`
}

type CourseStudent struct {
	Id               int64        `json:"id"`
	StudentId        int64        `json:"student_id"`         // 学生id
	CourseId         int64        `json:"course_id"`          // 课程id
	CourseSnapshotId int64        `json:"course_snapshot_id"` // 课程快照id
	ActivityId       int64        `json:"activity_id"`        // 活动id
	Type             int64        `json:"type"`               // 活动类型：1-学校活动，2-班级活动，3-基地活动 ,4-家庭活动 5-课后活动 6-研学活动
	IsCourseTask     int64        `json:"is_course_task"`     // 是否完成课程任务书 1-完成 0-未完成
	CreatedAt        sql.NullTime `json:"created_at"`
	UpdatedAt        sql.NullTime `json:"updated_at"`
	DeletedAt        sql.NullTime `json:"deleted_at"`
}

type CourseCiting struct {
	Id               int64        `json:"id"`
	CourseSnapshotId int64        `json:"course_snapshot_id"` // 课程快照id
	Type             int64        `json:"type"`               // 类型 课程快照-1 外部资料-2
	Name             string       `json:"name"`               // 快照活动名称 或 外部名称
	Ext              string       `json:"ext"`                // 快照id 或者 外部url
	CreatedAt        sql.NullTime `json:"created_at"`
}
