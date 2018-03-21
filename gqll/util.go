package gqll

func NameOf(v interface{}) string {
	if feat, ok := v.(featName); ok {
		return feat.GetName()
	}
	return ""
}
