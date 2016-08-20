package tmpq

import "testing"

func TestConnectionString(t *testing.T) {
	expected := "dbname=mydb user=dbuser password=pass host=localhost port=5432 sslmode=disable fallback_application_name=tmpq-test connect_timeout=60 sslcert=path/to/cert sslkey=path/to/key sslrootcert=path/to/rootcert"
	cs := &ConnectionString{
		DBName:                  "mydb",
		User:                    "dbuser",
		Password:                "pass",
		Host:                    "localhost",
		Port:                    5432,
		SSlMode:                 SSLModeDisable,
		FallbackApplicationName: "tmpq-test",
		ConnectTimeoutSeconds:   60,
		SSLCertFilePath:         "path/to/cert",
		SSLKeyFilePath:          "path/to/key",
		SSLRootCertFilePath:     "path/to/rootcert",
	}
	actual := cs.String()

	if expected != actual {
		t.Fatal("Invalid ConnectionString", actual)
	}

}
