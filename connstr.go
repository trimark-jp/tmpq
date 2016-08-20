package tmpq

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type (
	// SSLMode implements sslmode in connection string.
	SSLMode string

	// ConnectionString implements connection string for PostgreSQL.
	// See folloing links:
	//  https://www.postgresql.org/docs/current/static/libpq-connect.html
	//  https://godoc.org/github.com/lib/pq
	ConnectionString struct {
		DBName                  string  `json:"dbname" connect:"dbname"`
		User                    string  `json:"user" connect:"user"`
		Password                string  `json:"password" connect:"password"`
		Host                    string  `json:"host" connect:"host"`
		Port                    int     `json:"port" connect:"port"`
		SSlMode                 SSLMode `json:"sslmode" connect:"sslmode"`
		FallbackApplicationName string  `json:"fallback_application_name" connect:"fallback_application_name"`
		ConnectTimeoutSeconds   int     `json:"connect_timeout" connect:"connect_timeout"`
		SSLCertFilePath         string  `json:"sslcert" connect:"sslcert"`
		SSLKeyFilePath          string  `json:"sslkey" connect:"sslkey"`
		SSLRootCertFilePath     string  `json:"sslrootcert" connect:"sslrootcert"`
	}
)

// Constants for SSLMode.
const (
	SSLModeDisable    SSLMode = "disable"
	SSLModeRequire    SSLMode = "require"
	SSLModeVerifyCA   SSLMode = "verify-ca"
	SSLModeVerifyFull SSLMode = "verify-full"
)

const (
	connectionStringStringParamFormat = "%s=%s"
	connectionStringIntParamFormat    = "%s=%d"
)

var (
	whiteSpaceRegex  = regexp.MustCompile(`\s+`)
	singleQuoteRegex = regexp.MustCompile(`'`)
)

func (c *ConnectionString) String() string {
	return strings.Join(c.collectParams(), " ")
}

// StringParam returns string parameter in connection string format.
func (c *ConnectionString) collectParams() []string {

	val := reflect.ValueOf(c).Elem()
	t := val.Type()

	result := []string{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if tag, ok := f.Tag.Lookup("connect"); ok {
			p := c.param(val.Field(i), tag)
			if p != "" {
				result = append(result, p)
			}
		}
	}

	return result
}

func (c *ConnectionString) param(val reflect.Value, tag string) string {
	switch val.Kind() {
	case reflect.String:
		return stringParam(val.String(), tag)
	case reflect.Int:
		return intParam(int(val.Int()), tag)
	}
	return ""
}

func stringParam(str string, tag string) string {
	if str == "" {
		return ""
	}
	v := str
	if whiteSpaceRegex.MatchString(str) {
		if singleQuoteRegex.MatchString(str) {
			v = singleQuoteRegex.ReplaceAllString(str, `\'`)
		}
		v = `'` + v + `'`
	}

	return fmt.Sprintf(connectionStringStringParamFormat,
		tag, v)

}

func intParam(n int, tag string) string {
	if n == 0 {
		return ""
	}
	return fmt.Sprintf(connectionStringIntParamFormat,
		tag, n)
}
