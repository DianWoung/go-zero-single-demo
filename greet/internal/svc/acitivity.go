package svc

type SoloActivityInfo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type ActivityGroupBySchoolCount struct {
	SchoolId int `json:"school_id"`
	Count    int `json:"count"`
}

// CourseSnapshot表 + Course表信息（使用left join）
type CourseSnapshotInfo struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	CourseId   int64  `json:"course_id"`   // 课程id；相关的外键都要展示出来，不然查不出结构
	ClassHour  int64  `json:"class_hour"`  // 课时
	CoverImage string `json:"cover_image"` // 封面图片

	//Course表的字段（使用left join）
	Type      int64  `json:"type"`       // 账号所属的机构类型,1:触点管理员，2：第三方机构，3：学校，4：教育局，5：运营公司
	IsDefault int64  `json:"is_default"` // 是否默认课程：1-是，0-否
	IsCloud   int64  `json:"is_cloud"`   // 是否云课程1不是2是
	CourseNo  string `json:"course_no"`  // 课程编号
}
