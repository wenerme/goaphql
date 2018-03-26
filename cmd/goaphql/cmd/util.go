package cmd

import (
	"github.com/Sirupsen/logrus"
	"strings"
)

func sliceToMap(a []string) map[string]string {
	m := map[string]string{}
	for _, v := range a {
		i := strings.IndexRune(v, '=')
		if i < 0 {
			logrus.WithField("Option", v).Fatal("incorrect map option")
		}
		m[strings.TrimSpace(v[:i])] = strings.TrimSpace(v[i+1:])
	}
	return m
}
