/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package keywords

const (
	KwEOF KeyWord = iota
	KwAbort
	KwAction
	KwAdd
	KwAfter
	KwAll
	KwAlter
	KwAnalyze
	KwAnd
	KwAs
	KwAsc
	KwAttach
	KwAutoincrement
	KwBefore
	KwBegin
	KwBetween
	KwBy
	KwCascade
	KwCase
	KwCast
	KwCheck
	KwCollate
	KwColumn
	KwCommit
	KwConflict
	KwConstraint
	KwCreate
	KwCross
	KwCurrentDate
	KwCurrentTime
	KwCurrentTimestamp
	KwDatabase
	KwDefault
	KwDeferrable
	KwDeferred
	KwDelete
	KwDesc
	KwDetach
	KwDistinct
	KwDrop
	KwEach
	KwElse
	KwEnd
	KwEscape
	KwExcept
	KwExclusive
	KwExists
	KwExplain
	KwFail
	KwFor
	KwForeign
	KwFrom
	KwFull
	KwFalse
	KwGlob
	KwGroup
	KwHaving
	KwIf
	KwIgnore
	KwImmediate
	KwIn
	KwIndex
	KwIndexed
	KwInitially
	KwInner
	KwInsert
	KwInstead
	KwIntersect
	KwInto
	KwIs
	KwIsnull
	KwJoin
	KwKey
	KwLeft
	KwLike
	KwLimit
	KwMatch
	KwNatural
	KwNo
	KwNot
	KwNotnull
	KwNull
	KwOf
	KwOffset
	KwOn
	KwOr
	KwOrder
	KwOuter
	KwPlan
	KwPragma
	KwPrimary
	KwQuery
	KwRaise
	KwRecursive
	KwReferences
	KwRegexp
	KwReindex
	KwRelease
	KwRename
	KwReplace
	KwRestrict
	KwRight
	KwRollback
	KwRow
	KwSavepoint
	KwSelect
	KwSet
	KwTable
	KwTemp
	KwTemporary
	KwThen
	KwTrue
	KwTo
	KwTransaction
	KwTrigger
	KwUnion
	KwUnique
	KwUpdate
	KwUsing
	KwVacuum
	KwValues
	KwView
	KwVirtual
	KwWhen
	KwWhere
	KwWith
	KwWithout
	KwNotBetween
	KwIsNot
	KwNotIn
	KwNotLike
	KwNotGlob
	KwNotMatch
	KwNotRegexp
)

var keyWordByLowerString = map[KeyWord]string{
	KwAbort:            "ABORT",
	KwAction:           "ACTION",
	KwAdd:              "ADD",
	KwAfter:            "AFTER",
	KwAll:              "ALL",
	KwAlter:            "ALTER",
	KwAnalyze:          "ANALYZE",
	KwAnd:              "AND",
	KwAs:               "AS",
	KwAsc:              "ASC",
	KwAttach:           "ATTACH",
	KwAutoincrement:    "AUTOINCREMENT",
	KwBefore:           "BEFORE",
	KwBegin:            "BEGIN",
	KwBetween:          "BETWEEN",
	KwBy:               "BY",
	KwCascade:          "CASCADE",
	KwCase:             "CASE",
	KwCast:             "CAST",
	KwCheck:            "CHECK",
	KwCollate:          "COLLATE",
	KwColumn:           "COLUMN",
	KwCommit:           "COMMIT",
	KwConflict:         "CONFLICT",
	KwConstraint:       "CONSTRAINT",
	KwCreate:           "CREATE",
	KwCross:            "CROSS",
	KwCurrentDate:      "CURRENT_DATE",
	KwCurrentTime:      "CURRENT_TIME",
	KwCurrentTimestamp: "CURRENT_TIMESTAMP",
	KwDatabase:         "DATABASE",
	KwDefault:          "DEFAULT",
	KwDeferrable:       "DEFERRABLE",
	KwDeferred:         "DEFERRED",
	KwDelete:           "DELETE",
	KwDesc:             "DESC",
	KwDetach:           "DETACH",
	KwDistinct:         "DISTINCT",
	KwDrop:             "DROP",
	KwEach:             "EACH",
	KwElse:             "ELSE",
	KwEnd:              "END",
	KwEscape:           "ESCAPE",
	KwExcept:           "EXCEPT",
	KwExclusive:        "EXCLUSIVE",
	KwExists:           "EXISTS",
	KwExplain:          "EXPLAIN",
	KwFail:             "FAIL",
	KwFor:              "FOR",
	KwForeign:          "FOREIGN",
	KwFrom:             "FROM",
	KwFull:             "FULL",
	KwFalse:            "FALSE",
	KwGlob:             "GLOB",
	KwGroup:            "GROUP",
	KwHaving:           "HAVING",
	KwIf:               "IF",
	KwIgnore:           "IGNORE",
	KwImmediate:        "IMMEDIATE",
	KwIn:               "IN",
	KwIndex:            "INDEX",
	KwIndexed:          "INDEXED",
	KwInitially:        "INITIALLY",
	KwInner:            "INNER",
	KwInsert:           "INSERT",
	KwInstead:          "INSTEAD",
	KwIntersect:        "INTERSECT",
	KwInto:             "INTO",
	KwIs:               "IS",
	KwIsnull:           "ISNULL",
	KwJoin:             "JOIN",
	KwKey:              "KEY",
	KwLeft:             "LEFT",
	KwLike:             "LIKE",
	KwLimit:            "LIMIT",
	KwMatch:            "MATCH",
	KwNatural:          "NATURAL",
	KwNo:               "NO",
	KwNot:              "NOT",
	KwNotnull:          "NOTNULL",
	KwNull:             "NULL",
	KwOf:               "OF",
	KwOffset:           "OFFSET",
	KwOn:               "ON",
	KwOr:               "OR",
	KwOrder:            "ORDER",
	KwOuter:            "OUTER",
	KwPlan:             "PLAN",
	KwPragma:           "PRAGMA",
	KwPrimary:          "PRIMARY",
	KwQuery:            "QUERY",
	KwRaise:            "RAISE",
	KwRecursive:        "RECURSIVE",
	KwReferences:       "REFERENCES",
	KwRegexp:           "REGEXP",
	KwReindex:          "REINDEX",
	KwRelease:          "RELEASE",
	KwRename:           "RENAME",
	KwReplace:          "REPLACE",
	KwRestrict:         "RESTRICT",
	KwRight:            "RIGHT",
	KwRollback:         "ROLLBACK",
	KwRow:              "ROW",
	KwSavepoint:        "SAVEPOINT",
	KwSelect:           "SELECT",
	KwSet:              "SET",
	KwTable:            "TABLE",
	KwTemp:             "TEMP",
	KwTemporary:        "TEMPORARY",
	KwThen:             "THEN",
	KwTrue:             "TRUE",
	KwTo:               "TO",
	KwTransaction:      "TRANSACTION",
	KwTrigger:          "TRIGGER",
	KwUnion:            "UNION",
	KwUnique:           "UNIQUE",
	KwUpdate:           "UPDATE",
	KwUsing:            "USING",
	KwVacuum:           "VACUUM",
	KwValues:           "VALUES",
	KwView:             "VIEW",
	KwVirtual:          "VIRTUAL",
	KwWhen:             "WHEN",
	KwWhere:            "WHERE",
	KwWith:             "WITH",
	KwWithout:          "WITHOUT",
	KwNotBetween:       "NOT BETWEEN",
	KwIsNot:            "IS NOT",
	KwNotIn:            "NOT IN",
	KwNotLike:          "NOT LIKE",
	KwNotGlob:          "NOT GLOB",
	KwNotMatch:         "NOT MATCH",
	KwNotRegexp:        "NOT REGEXP",
}

var lowerStringByKeyWord = map[string]KeyWord{
	"ABORT":             KwAbort,
	"ACTION":            KwAction,
	"ADD":               KwAdd,
	"AFTER":             KwAfter,
	"ALL":               KwAll,
	"ALTER":             KwAlter,
	"ANALYZE":           KwAnalyze,
	"AND":               KwAnd,
	"AS":                KwAs,
	"ASC":               KwAsc,
	"ATTACH":            KwAttach,
	"AUTOINCREMENT":     KwAutoincrement,
	"BEFORE":            KwBefore,
	"BEGIN":             KwBegin,
	"BETWEEN":           KwBetween,
	"BY":                KwBy,
	"CASCADE":           KwCascade,
	"CASE":              KwCase,
	"CAST":              KwCast,
	"CHECK":             KwCheck,
	"COLLATE":           KwCollate,
	"COLUMN":            KwColumn,
	"COMMIT":            KwCommit,
	"CONFLICT":          KwConflict,
	"CONSTRAINT":        KwConstraint,
	"CREATE":            KwCreate,
	"CROSS":             KwCross,
	"CURRENT_DATE":      KwCurrentDate,
	"CURRENT_TIME":      KwCurrentTime,
	"CURRENT_TIMESTAMP": KwCurrentTimestamp,
	"DATABASE":          KwDatabase,
	"DEFAULT":           KwDefault,
	"DEFERRABLE":        KwDeferrable,
	"DEFERRED":          KwDeferred,
	"DELETE":            KwDelete,
	"DESC":              KwDesc,
	"DETACH":            KwDetach,
	"DISTINCT":          KwDistinct,
	"DROP":              KwDrop,
	"EACH":              KwEach,
	"ELSE":              KwElse,
	"END":               KwEnd,
	"ESCAPE":            KwEscape,
	"EXCEPT":            KwExcept,
	"EXCLUSIVE":         KwExclusive,
	"EXISTS":            KwExists,
	"EXPLAIN":           KwExplain,
	"FAIL":              KwFail,
	"FOR":               KwFor,
	"FOREIGN":           KwForeign,
	"FROM":              KwFrom,
	"FULL":              KwFull,
	"FALSE":             KwFalse,
	"GLOB":              KwGlob,
	"GROUP":             KwGroup,
	"HAVING":            KwHaving,
	"IF":                KwIf,
	"IGNORE":            KwIgnore,
	"IMMEDIATE":         KwImmediate,
	"IN":                KwIn,
	"INDEX":             KwIndex,
	"INDEXED":           KwIndexed,
	"INITIALLY":         KwInitially,
	"INNER":             KwInner,
	"INSERT":            KwInsert,
	"INSTEAD":           KwInstead,
	"INTERSECT":         KwIntersect,
	"INTO":              KwInto,
	"IS":                KwIs,
	"ISNULL":            KwIsnull,
	"JOIN":              KwJoin,
	"KEY":               KwKey,
	"LEFT":              KwLeft,
	"LIKE":              KwLike,
	"LIMIT":             KwLimit,
	"MATCH":             KwMatch,
	"NATURAL":           KwNatural,
	"NO":                KwNo,
	"NOT":               KwNot,
	"NOTNULL":           KwNotnull,
	"NULL":              KwNull,
	"OF":                KwOf,
	"OFFSET":            KwOffset,
	"ON":                KwOn,
	"OR":                KwOr,
	"ORDER":             KwOrder,
	"OUTER":             KwOuter,
	"PLAN":              KwPlan,
	"PRAGMA":            KwPragma,
	"PRIMARY":           KwPrimary,
	"QUERY":             KwQuery,
	"RAISE":             KwRaise,
	"RECURSIVE":         KwRecursive,
	"REFERENCES":        KwReferences,
	"REGEXP":            KwRegexp,
	"REINDEX":           KwReindex,
	"RELEASE":           KwRelease,
	"RENAME":            KwRename,
	"REPLACE":           KwReplace,
	"RESTRICT":          KwRestrict,
	"RIGHT":             KwRight,
	"ROLLBACK":          KwRollback,
	"ROW":               KwRow,
	"SAVEPOINT":         KwSavepoint,
	"SELECT":            KwSelect,
	"SET":               KwSet,
	"TABLE":             KwTable,
	"TEMP":              KwTemp,
	"TEMPORARY":         KwTemporary,
	"THEN":              KwThen,
	"TRUE":              KwTrue,
	"TO":                KwTo,
	"TRANSACTION":       KwTransaction,
	"TRIGGER":           KwTrigger,
	"UNION":             KwUnion,
	"UNIQUE":            KwUnique,
	"UPDATE":            KwUpdate,
	"USING":             KwUsing,
	"VACUUM":            KwVacuum,
	"VALUES":            KwValues,
	"VIEW":              KwView,
	"VIRTUAL":           KwVirtual,
	"WHEN":              KwWhen,
	"WHERE":             KwWhere,
	"WITH":              KwWith,
	"WITHOUT":           KwWithout,
	"NOT BETWEEN":       KwNotBetween,
	"IS NOT":            KwIsNot,
	"NOT IN":            KwNotIn,
	"NOT LIKE":          KwNotLike,
	"NOT GLOB":          KwNotGlob,
	"NOT MATCH":         KwNotMatch,
	"NOT REGEXP":        KwNotRegexp,
}
