Update(ctx context.Context, {{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error
List(ctx context.Context, page, limit int64) ([]*{{.upperStartCamelObject}}, error)
TransUpdate(ctx context.Context, session sqlx.Session, data *{{.upperStartCamelObject}}) error
Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error