package keywords

import "strings"

type KeyWord int

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
)

var keyWordByLowerString = map[KeyWord]string{
	KwAbort:            "abort",
	KwAction:           "action",
	KwAdd:              "add",
	KwAfter:            "after",
	KwAll:              "all",
	KwAlter:            "alter",
	KwAnalyze:          "analyze",
	KwAnd:              "and",
	KwAs:               "as",
	KwAsc:              "asc",
	KwAttach:           "attach",
	KwAutoincrement:    "autoincrement",
	KwBefore:           "before",
	KwBegin:            "begin",
	KwBetween:          "between",
	KwBy:               "by",
	KwCascade:          "cascade",
	KwCase:             "case",
	KwCast:             "cast",
	KwCheck:            "check",
	KwCollate:          "collate",
	KwColumn:           "column",
	KwCommit:           "commit",
	KwConflict:         "conflict",
	KwConstraint:       "constraint",
	KwCreate:           "create",
	KwCross:            "cross",
	KwCurrentDate:      "current_date",
	KwCurrentTime:      "current_time",
	KwCurrentTimestamp: "current_timestamp",
	KwDatabase:         "database",
	KwDefault:          "default",
	KwDeferrable:       "deferrable",
	KwDeferred:         "deferred",
	KwDelete:           "delete",
	KwDesc:             "desc",
	KwDetach:           "detach",
	KwDistinct:         "distinct",
	KwDrop:             "drop",
	KwEach:             "each",
	KwElse:             "else",
	KwEnd:              "end",
	KwEscape:           "escape",
	KwExcept:           "except",
	KwExclusive:        "exclusive",
	KwExists:           "exists",
	KwExplain:          "explain",
	KwFail:             "fail",
	KwFor:              "for",
	KwForeign:          "foreign",
	KwFrom:             "from",
	KwFull:             "full",
	KwFalse:            "false",
	KwGlob:             "glob",
	KwGroup:            "group",
	KwHaving:           "having",
	KwIf:               "if",
	KwIgnore:           "ignore",
	KwImmediate:        "immediate",
	KwIn:               "in",
	KwIndex:            "index",
	KwIndexed:          "indexed",
	KwInitially:        "initially",
	KwInner:            "inner",
	KwInsert:           "insert",
	KwInstead:          "instead",
	KwIntersect:        "intersect",
	KwInto:             "into",
	KwIs:               "is",
	KwIsnull:           "isnull",
	KwJoin:             "join",
	KwKey:              "key",
	KwLeft:             "left",
	KwLike:             "like",
	KwLimit:            "limit",
	KwMatch:            "match",
	KwNatural:          "natural",
	KwNo:               "no",
	KwNot:              "not",
	KwNotnull:          "notnull",
	KwNull:             "null",
	KwOf:               "of",
	KwOffset:           "offset",
	KwOn:               "on",
	KwOr:               "or",
	KwOrder:            "order",
	KwOuter:            "outer",
	KwPlan:             "plan",
	KwPragma:           "pragma",
	KwPrimary:          "primary",
	KwQuery:            "query",
	KwRaise:            "raise",
	KwRecursive:        "recursive",
	KwReferences:       "references",
	KwRegexp:           "regexp",
	KwReindex:          "reindex",
	KwRelease:          "release",
	KwRename:           "rename",
	KwReplace:          "replace",
	KwRestrict:         "restrict",
	KwRight:            "right",
	KwRollback:         "rollback",
	KwRow:              "row",
	KwSavepoint:        "savepoint",
	KwSelect:           "select",
	KwSet:              "set",
	KwTable:            "table",
	KwTemp:             "temp",
	KwTemporary:        "temporary",
	KwThen:             "then",
	KwTrue:             "true",
	KwTo:               "to",
	KwTransaction:      "transaction",
	KwTrigger:          "trigger",
	KwUnion:            "union",
	KwUnique:           "unique",
	KwUpdate:           "update",
	KwUsing:            "using",
	KwVacuum:           "vacuum",
	KwValues:           "values",
	KwView:             "view",
	KwVirtual:          "virtual",
	KwWhen:             "when",
	KwWhere:            "where",
	KwWith:             "with",
	KwWithout:          "without",
}

var lowerStringByKeyWord = map[string]KeyWord{
	"abort":             KwAbort,
	"action":            KwAction,
	"add":               KwAdd,
	"after":             KwAfter,
	"all":               KwAll,
	"alter":             KwAlter,
	"analyze":           KwAnalyze,
	"and":               KwAnd,
	"as":                KwAs,
	"asc":               KwAsc,
	"attach":            KwAttach,
	"autoincrement":     KwAutoincrement,
	"before":            KwBefore,
	"begin":             KwBegin,
	"between":           KwBetween,
	"by":                KwBy,
	"cascade":           KwCascade,
	"case":              KwCase,
	"cast":              KwCast,
	"check":             KwCheck,
	"collate":           KwCollate,
	"column":            KwColumn,
	"commit":            KwCommit,
	"conflict":          KwConflict,
	"constraint":        KwConstraint,
	"create":            KwCreate,
	"cross":             KwCross,
	"current_date":      KwCurrentDate,
	"current_time":      KwCurrentTime,
	"current_timestamp": KwCurrentTimestamp,
	"database":          KwDatabase,
	"default":           KwDefault,
	"deferrable":        KwDeferrable,
	"deferred":          KwDeferred,
	"delete":            KwDelete,
	"desc":              KwDesc,
	"detach":            KwDetach,
	"distinct":          KwDistinct,
	"drop":              KwDrop,
	"each":              KwEach,
	"else":              KwElse,
	"end":               KwEnd,
	"escape":            KwEscape,
	"except":            KwExcept,
	"exclusive":         KwExclusive,
	"exists":            KwExists,
	"explain":           KwExplain,
	"fail":              KwFail,
	"for":               KwFor,
	"foreign":           KwForeign,
	"from":              KwFrom,
	"full":              KwFull,
	"false":             KwFalse,
	"glob":              KwGlob,
	"group":             KwGroup,
	"having":            KwHaving,
	"if":                KwIf,
	"ignore":            KwIgnore,
	"immediate":         KwImmediate,
	"in":                KwIn,
	"index":             KwIndex,
	"indexed":           KwIndexed,
	"initially":         KwInitially,
	"inner":             KwInner,
	"insert":            KwInsert,
	"instead":           KwInstead,
	"intersect":         KwIntersect,
	"into":              KwInto,
	"is":                KwIs,
	"isnull":            KwIsnull,
	"join":              KwJoin,
	"key":               KwKey,
	"left":              KwLeft,
	"like":              KwLike,
	"limit":             KwLimit,
	"match":             KwMatch,
	"natural":           KwNatural,
	"no":                KwNo,
	"not":               KwNot,
	"notnull":           KwNotnull,
	"null":              KwNull,
	"of":                KwOf,
	"offset":            KwOffset,
	"on":                KwOn,
	"or":                KwOr,
	"order":             KwOrder,
	"outer":             KwOuter,
	"plan":              KwPlan,
	"pragma":            KwPragma,
	"primary":           KwPrimary,
	"query":             KwQuery,
	"raise":             KwRaise,
	"recursive":         KwRecursive,
	"references":        KwReferences,
	"regexp":            KwRegexp,
	"reindex":           KwReindex,
	"release":           KwRelease,
	"rename":            KwRename,
	"replace":           KwReplace,
	"restrict":          KwRestrict,
	"right":             KwRight,
	"rollback":          KwRollback,
	"row":               KwRow,
	"savepoint":         KwSavepoint,
	"select":            KwSelect,
	"set":               KwSet,
	"table":             KwTable,
	"temp":              KwTemp,
	"temporary":         KwTemporary,
	"then":              KwThen,
	"true":              KwTrue,
	"to":                KwTo,
	"transaction":       KwTransaction,
	"trigger":           KwTrigger,
	"union":             KwUnion,
	"unique":            KwUnique,
	"update":            KwUpdate,
	"using":             KwUsing,
	"vacuum":            KwVacuum,
	"values":            KwValues,
	"view":              KwView,
	"virtual":           KwVirtual,
	"when":              KwWhen,
	"where":             KwWhere,
	"with":              KwWith,
	"without":           KwWithout,
}

func (kw KeyWord) String() string { return keyWordByLowerString[kw] }
func IsKeyWord(in string) bool    { return lowerStringByKeyWord[strings.ToLower(in)] != KwEOF }
func GetKeyWord(in string) KeyWord {
	val, ok := lowerStringByKeyWord[strings.ToLower(in)]
	if ok {
		return val
	}
	return KwEOF
}
