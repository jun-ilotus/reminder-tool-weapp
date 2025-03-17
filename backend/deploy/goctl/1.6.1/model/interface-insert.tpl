Insert(ctx context.Context, data *{{.upperStartCamelObject}}) (sql.Result,error)
TransInsert(ctx context.Context, session sqlx.Session, data *{{.upperStartCamelObject}}) (sql.Result, error)