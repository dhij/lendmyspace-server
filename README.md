# LendMySpace API

## Purpose

- Implement Domain Driven Architecture
- Learn CI/CD with Github Action
- Push the image to ECR and deploy on EKS

## Tech stack

- Golang
- PostgreSQL
- AWS ECR/EKS

## Core dependency

- [sqlx](https://github.com/jmoiron/sqlx)
  - Database entity mapping
- [viper](https://github.com/spf13/viper)
  - Manage configuration
- model
  - DTO & entity mapping

## Structure

- Inspired from https://github.com/bxcodec/go-clean-arch

## Make command

- createdb
  - create lendmyspace db in lendmyspace_postgres container
- drop db
  - drop lendmyspace db
- migrateup
  - migrate up (golang-migrate)
- migrateup1
  - migrate up 1 step
- migratedown
  - migrate down
- migratedown1
  - migrate down 1 step
- postgres
  - run postgres container
- test
  - run all tests
- server
  - run the server locally

## Design and Planning Docs

- [Figma](https://www.figma.com/file/wj9zJzOud1QSKTBwcqCJ4A/lendmyspace?node-id=0%3A1)

### Semantic Commit messages

- feat: :zap: new feature
- fix: :bug: fix bug
- refactor: :hammer: refactor code
