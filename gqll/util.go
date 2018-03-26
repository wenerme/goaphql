package gqll

func NameOf(v interface{}) string {
	if feat, ok := v.(HasName); ok {
		return feat.GetName()
	}
	return ""
}
