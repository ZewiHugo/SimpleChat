package secrets

var JwtSecret string

func init() {
	JwtSecret = "my secret"
}