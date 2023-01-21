server:
	CompileDaemon -command="./marcellinuselbert-be"

migration:
	go run migrate/migrate.go