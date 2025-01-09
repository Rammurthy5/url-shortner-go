# url-shortner-go
url shortner app in go - full stack app. 

This project is to learn how to create a Full stack Go app.

production-ready standards followed for:
1. Folder structure.
2. Configuration via Viper.
3. Loading templates.
4. Logging via zap.
5. golang-migrate for db migrations.
6. sqlc and pgx for database interaction.
7. cicd pipeline.
more tests, Dependabot, pre-commit webhooks. [To-Do]


# Migrate commands
make migrate_up

### rollback migration
make migrate_down

### fix a migration
make migrate_fix

## to run tests
make test

## to audit
make audit