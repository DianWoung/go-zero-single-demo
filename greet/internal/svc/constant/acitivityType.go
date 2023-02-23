package constant

const DEFAULT_ERROR_CODE int = -1 //默认错误码

// 活动状态,1:学校活动，2：班级活动，3：校外活动，4：家庭活动，5：课后活动 6：研学活动
const (
	ACTIVITY_SCHOOL     = iota + 1 //1
	ACTIVITY_CLASSIFY              //2
	ACTIVITY_SCHOOL_OUT            //3
	ACTIVITY_FAMILY                //4
	ACTIVITY_AFTER                 //5
	ACTIVITY_STUDIES               //6
)
