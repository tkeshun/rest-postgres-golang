https://github.com/jackc/pgx/wiki/Getting-started-with-pgx

db.ExecとQueryの違い

Query
Query 関数は、一つ以上の結果行を返す可能性がある SQL ステートメント（主に SELECT クエリ）を実行するために使用されます。
返されるのは Rows オブジェクトで、これを使用して結果セットを反復処理できます。各行は Next メソッドを使用して取得し、Scan メソッドを使用して個々の列を変数に読み込みます。
例えば、データベースから複数のレコードを取得する場合に使用します。
QueryRow
QueryRow 関数は、単一の結果行のみを期待するクエリに使用されます。この関数は便宜的に用意されており、主に SELECT クエリで1行だけの結果を期待する場合に使います。
返されるのは Row オブジェクトで、このオブジェクトから直接 Scan メソッドを呼び出してデータを変数に読み込みます。
クエリの結果がゼロ行の場合、Scan は ErrNoRows エラーを返します。結果が複数行ある場合、最初の行のみが使用され、残りは無視されます。
Exec
Exec 関数は、結果行を返さない SQL ステートメント（INSERT、UPDATE、DELETE、または DDL 文など）を実行するために使用されます。
返されるのは CommandTag オブジェクトで、これには実行されたクエリによって影響を受けた行の数などの情報が含まれます。
例えば、データをデータベースに挿入したり、更新、削除する場合や、テーブルを作成するなどのDDL操作を行う場合に使用します。