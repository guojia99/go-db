package tokenizer

//
//// The list of tokens.
//const (
//	// Special tokens
//	ILLEGAL Token = iota
//	EOF
//	COMMENT
//	SPACE
//
//	literal_beg
//	IDENT   // IDENT
//	QIDENT  // "IDENT"
//	STRING  // 'string'
//	BLOB    // ???
//	FLOAT   // 123.45
//	INTEGER // 123
//	NULL    // NULL
//	TRUE    // true
//	FALSE   // false
//	BIND    // ? or ?NNN or :VVV or @VVV or $VVV
//	literal_end
//
//	operator_beg
//	SEMI   // ;
//	LP     // (
//	RP     // )
//	COMMA  // ,
//	NE     // !=
//	EQ     // =
//	LE     // <=
//	LT     // <
//	GT     // >
//	GE     // >=
//	BITAND // &
//	BITOR  // |
//	BITNOT // !
//	LSHIFT // <<
//	RSHIFT // >>
//	PLUS   // +
//	MINUS  // -
//	STAR   // *
//	SLASH  // /
//	REM    // %
//	CONCAT // ||
//	DOT    // .
//	operator_end
//
//	keyword_beg
//	ABORT
//	ACTION
//	ADD
//	AFTER
//	AGG_COLUMN
//	AGG_FUNCTION
//	ALL
//	ALTER
//	ANALYZE
//	AND
//	AS
//	ASC
//	ASTERISK
//	ATTACH
//	AUTOINCREMENT
//	BEFORE
//	BEGIN
//	BETWEEN
//	BY
//	CASCADE
//	CASE
//	CAST
//	CHECK
//	COLUMN
//	COLUMNKW
//	COMMIT
//	CONFLICT
//	CONSTRAINT
//	CREATE
//	CROSS
//	CTIME_KW
//	CURRENT
//	CURRENT_TIME
//	CURRENT_DATE
//	CURRENT_TIMESTAMP
//	DATABASE
//	DEFAULT
//	DEFERRABLE
//	DEFERRED
//	DELETE
//	DESC
//	DETACH
//	DISTINCT
//	DO
//	DROP
//	EACH
//	ELSE
//	END
//	ESCAPE
//	EXCEPT
//	EXCLUDE
//	EXCLUSIVE
//	EXISTS
//	EXPLAIN
//	FAIL
//	FILTER
//	FIRST
//	FOLLOWING
//	FOR
//	FOREIGN
//	FROM
//	FUNCTION
//	GLOB
//	GROUP
//	GROUPS
//	HAVING
//	IF
//	IF_NULL_ROW
//	IGNORE
//	IMMEDIATE
//	IN
//	INDEX
//	INDEXED
//	INITIALLY
//	INNER
//	INSERT
//	INSTEAD
//	INTERSECT
//	INTO
//	IS
//	ISNOT
//	ISNULL // TODO: REMOVE?
//	JOIN
//	KEY
//	LAST
//	LEFT
//	LIKE
//	LIMIT
//	MATCH
//	NATURAL
//	NO
//	NOT
//	NOTBETWEEN
//	NOTEXISTS
//	NOTGLOB
//	NOTHING
//	NOTIN
//	NOTLIKE
//	NOTMATCH
//	NOTNULL
//	NOTREGEXP
//	NULLS
//	OF
//	OFFSET
//	ON
//	OR
//	ORDER
//	OTHERS
//	OUTER
//	OVER
//	PARTITION
//	PLAN
//	PRAGMA
//	PRECEDING
//	PRIMARY
//	QUERY
//	RAISE
//	RANGE
//	RECURSIVE
//	REFERENCES
//	REGEXP
//	REGISTER
//	REINDEX
//	RELEASE
//	RENAME
//	REPLACE
//	RESTRICT
//	ROLLBACK
//	ROW
//	ROWS
//	SAVEPOINT
//	SELECT
//	SELECT_COLUMN
//	SET
//	SPAN
//	TABLE
//	TEMP
//	THEN
//	TIES
//	TO
//	TRANSACTION
//	TRIGGER
//	TRUTH
//	UNBOUNDED
//	UNION
//	UNIQUE
//	UPDATE
//	USING
//	VACUUM
//	VALUES
//	VARIABLE
//	VECTOR
//	VIEW
//	VIRTUAL
//	WHEN
//	WHERE
//	WINDOW
//	WITH
//	WITHOUT
//	keyword_end
//
//	ANY // ???
//	token_end
//)
//
//var tokens = [...]string{
//	ILLEGAL: "ILLEGAL",
//	EOF:     "EOF",
//	COMMENT: "COMMENT",
//	SPACE:   "SPACE",
//
//	IDENT:   "IDENT",
//	QIDENT:  "QIDENT",
//	STRING:  "STRING",
//	BLOB:    "BLOB",
//	FLOAT:   "FLOAT",
//	INTEGER: "INTEGER",
//	NULL:    "NULL",
//	TRUE:    "TRUE",
//	FALSE:   "FALSE",
//	BIND:    "BIND",
//
//	SEMI:   ";",
//	LP:     "(",
//	RP:     ")",
//	COMMA:  ",",
//	NE:     "!=",
//	EQ:     "=",
//	LE:     "<=",
//	LT:     "<",
//	GT:     ">",
//	GE:     ">=",
//	BITAND: "&",
//	BITOR:  "|",
//	BITNOT: "!",
//	LSHIFT: "<<",
//	RSHIFT: ">>",
//	PLUS:   "+",
//	MINUS:  "-",
//	STAR:   "*",
//	SLASH:  "/",
//	REM:    "%",
//	CONCAT: "||",
//	DOT:    ".",
//
//	ABORT:             "ABORT",
//	ACTION:            "ACTION",
//	ADD:               "ADD",
//	AFTER:             "AFTER",
//	AGG_COLUMN:        "AGG_COLUMN",
//	AGG_FUNCTION:      "AGG_FUNCTION",
//	ALL:               "ALL",
//	ALTER:             "ALTER",
//	ANALYZE:           "ANALYZE",
//	AND:               "AND",
//	AS:                "AS",
//	ASC:               "ASC",
//	ASTERISK:          "ASTERISK",
//	ATTACH:            "ATTACH",
//	AUTOINCREMENT:     "AUTOINCREMENT",
//	BEFORE:            "BEFORE",
//	BEGIN:             "BEGIN",
//	BETWEEN:           "BETWEEN",
//	BY:                "BY",
//	CASCADE:           "CASCADE",
//	CASE:              "CASE",
//	CAST:              "CAST",
//	CHECK:             "CHECK",
//	COLUMN:            "COLUMN",
//	COLUMNKW:          "COLUMNKW",
//	COMMIT:            "COMMIT",
//	CONFLICT:          "CONFLICT",
//	CONSTRAINT:        "CONSTRAINT",
//	CREATE:            "CREATE",
//	CROSS:             "CROSS",
//	CTIME_KW:          "CTIME_KW",
//	CURRENT:           "CURRENT",
//	CURRENT_TIME:      "CURRENT_TIME",
//	CURRENT_DATE:      "CURRENT_DATE",
//	CURRENT_TIMESTAMP: "CURRENT_TIMESTAMP",
//	DATABASE:          "DATABASE",
//	DEFAULT:           "DEFAULT",
//	DEFERRABLE:        "DEFERRABLE",
//	DEFERRED:          "DEFERRED",
//	DELETE:            "DELETE",
//	DESC:              "DESC",
//	DETACH:            "DETACH",
//	DISTINCT:          "DISTINCT",
//	DO:                "DO",
//	DROP:              "DROP",
//	EACH:              "EACH",
//	ELSE:              "ELSE",
//	END:               "END",
//	ESCAPE:            "ESCAPE",
//	EXCEPT:            "EXCEPT",
//	EXCLUDE:           "EXCLUDE",
//	EXCLUSIVE:         "EXCLUSIVE",
//	EXISTS:            "EXISTS",
//	EXPLAIN:           "EXPLAIN",
//	FAIL:              "FAIL",
//	FILTER:            "FILTER",
//	FIRST:             "FIRST",
//	FOLLOWING:         "FOLLOWING",
//	FOR:               "FOR",
//	FOREIGN:           "FOREIGN",
//	FROM:              "FROM",
//	FUNCTION:          "FUNCTION",
//	GLOB:              "GLOB",
//	GROUP:             "GROUP",
//	GROUPS:            "GROUPS",
//	HAVING:            "HAVING",
//	IF:                "IF",
//	IF_NULL_ROW:       "IF_NULL_ROW",
//	IGNORE:            "IGNORE",
//	IMMEDIATE:         "IMMEDIATE",
//	IN:                "IN",
//	INDEX:             "INDEX",
//	INDEXED:           "INDEXED",
//	INITIALLY:         "INITIALLY",
//	INNER:             "INNER",
//	INSERT:            "INSERT",
//	INSTEAD:           "INSTEAD",
//	INTERSECT:         "INTERSECT",
//	INTO:              "INTO",
//	IS:                "IS",
//	ISNOT:             "ISNOT",
//	ISNULL:            "ISNULL",
//	JOIN:              "JOIN",
//	KEY:               "KEY",
//	LAST:              "LAST",
//	LEFT:              "LEFT",
//	LIKE:              "LIKE",
//	LIMIT:             "LIMIT",
//	MATCH:             "MATCH",
//	NO:                "NO",
//	NATURAL:           "NATURAL",
//	NOT:               "NOT",
//	NOTBETWEEN:        "NOTBETWEEN",
//	NOTEXISTS:         "NOTEXISTS",
//	NOTGLOB:           "NOTGLOB",
//	NOTHING:           "NOTHING",
//	NOTIN:             "NOTIN",
//	NOTLIKE:           "NOTLIKE",
//	NOTMATCH:          "NOTMATCH",
//	NOTNULL:           "NOTNULL",
//	NOTREGEXP:         "NOTREGEXP",
//	NULLS:             "NULLS",
//	OF:                "OF",
//	OFFSET:            "OFFSET",
//	ON:                "ON",
//	OR:                "OR",
//	ORDER:             "ORDER",
//	OTHERS:            "OTHERS",
//	OUTER:             "OUTER",
//	OVER:              "OVER",
//	PARTITION:         "PARTITION",
//	PLAN:              "PLAN",
//	PRAGMA:            "PRAGMA",
//	PRECEDING:         "PRECEDING",
//	PRIMARY:           "PRIMARY",
//	QUERY:             "QUERY",
//	RAISE:             "RAISE",
//	RANGE:             "RANGE",
//	RECURSIVE:         "RECURSIVE",
//	REFERENCES:        "REFERENCES",
//	REGEXP:            "REGEXP",
//	REGISTER:          "REGISTER",
//	REINDEX:           "REINDEX",
//	RELEASE:           "RELEASE",
//	RENAME:            "RENAME",
//	REPLACE:           "REPLACE",
//	RESTRICT:          "RESTRICT",
//	ROLLBACK:          "ROLLBACK",
//	ROW:               "ROW",
//	ROWS:              "ROWS",
//	SAVEPOINT:         "SAVEPOINT",
//	SELECT:            "SELECT",
//	SELECT_COLUMN:     "SELECT_COLUMN",
//	SET:               "SET",
//	SPAN:              "SPAN",
//	TABLE:             "TABLE",
//	TEMP:              "TEMP",
//	THEN:              "THEN",
//	TIES:              "TIES",
//	TO:                "TO",
//	TRANSACTION:       "TRANSACTION",
//	TRIGGER:           "TRIGGER",
//	TRUTH:             "TRUTH",
//	UNBOUNDED:         "UNBOUNDED",
//	UNION:             "UNION",
//	UNIQUE:            "UNIQUE",
//	UPDATE:            "UPDATE",
//	USING:             "USING",
//	VACUUM:            "VACUUM",
//	VALUES:            "VALUES",
//	VARIABLE:          "VARIABLE",
//	VECTOR:            "VECTOR",
//	VIEW:              "VIEW",
//	VIRTUAL:           "VIRTUAL",
//	WHEN:              "WHEN",
//	WHERE:             "WHERE",
//	WINDOW:            "WINDOW",
//	WITH:              "WITH",
//	WITHOUT:           "WITHOUT",
//}
