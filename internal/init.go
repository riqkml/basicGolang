package internal

var connection string

func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
