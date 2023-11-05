package proto

func (g Gender) Display() string {
	switch g {
	case Gender_GENDER_MALE:
		return "男"
	case Gender_GENDER_FEMALE:
		return "女"
	default:
		return "未知"
	}
}

func NewGender(s string) Gender {
	switch s {
	case "":
		return Gender_GENDER_MALE
	case "1":
		return Gender_GENDER_FEMALE
	default:
		return Gender_GENDER_UNDEFINED
	}
}
