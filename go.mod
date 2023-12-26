module example.com/mod

go 1.21.4

require github.com/go-chi/chi/v5 v5.0.10

require github.com/golang-jwt/jwt v3.2.2+incompatible

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/cobra v1.8.0
	// Cobra must be installed where the main project/file (go.mod) files stays, otherwise shows error, same for all packages)

	github.com/spf13/pflag v1.0.5 // indirect
)
