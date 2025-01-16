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
8. request validation.
9. middleware for idempotency, and redis implementation. 
idempotency key will not allow duplicate requests before a logical request is processed (fail / success).
every time the form is refreshed there is going to be a new idempotency key. added tests for middleware.
 dockerisation [in-progress]
versioning, context, pagination, Dependabot, pre-commit webhooks. [To-Do]


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
