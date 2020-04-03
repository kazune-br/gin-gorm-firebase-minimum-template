${ENV}:
  dialect: mysql
  dir: db/migrations
  datasource: ${DB_USER}:${DB_PASS}@tcp(${DB_HOST})/${DB_NAME}?parseTime=true&charset=utf8mb4
  table: migrations