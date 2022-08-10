package env

import (
	"reflect"
	"os"
	"strings"
)

func CleanByEnvTags(prefix string, i interface{}) (err error) {
	if prefix != "" {
		prefix += "_"
	}

	t := reflect.TypeOf(i).Elem()
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Tag.Get("envconfig")
		if key == "" {
			continue
		}

		err = os.Unsetenv(strings.ToUpper(prefix + key))
		if err != nil {
			return err
		}
	}

	return nil
}
