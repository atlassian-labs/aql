// Code generated from JQLParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // JQLParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type JQLParser struct {
	*antlr.BaseParser
}

var jqlparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jqlparserParserInit() {
	staticData := &jqlparserParserStaticData
	staticData.literalNames = []string{
		"", "'('", "')'", "','", "'['", "']'", "'!'", "'<'", "'>'", "'>='",
		"'<='", "'='", "'!='", "'~'", "'!~'",
	}
	staticData.symbolicNames = []string{
		"", "LPAREN", "RPAREN", "COMMA", "LBRACKET", "RBRACKET", "BANG", "LT",
		"GT", "GTEQ", "LTEQ", "EQUALS", "NOT_EQUALS", "LIKE", "NOT_LIKE", "IN",
		"IS", "AND", "OR", "NOT", "EMPTY", "WAS", "CHANGED", "BEFORE", "AFTER",
		"FROM", "TO", "ON", "DURING", "ORDER", "BY", "ASC", "DESC", "POSNUMBER",
		"NEGNUMBER", "CUSTOMFIELD", "RESERVED_WORD", "STRING", "MATCHWS", "ERROR_RESERVED",
		"ERRORCHAR", "QUOTE_STRING", "UNCLOSED_QUOTE_STRING", "INVALID_QUOTE_STRING",
		"SQUOTE_STRING", "UNCLOSED_SQUOTE_STRING", "INVALID_SQUOTE_STRING",
	}
	staticData.ruleNames = []string{
		"jqlQuery", "jqlWhere", "jqlOrClause", "jqlAndClause", "jqlNotClause",
		"jqlSubClause", "jqlTerminalClause", "jqlTerminalClauseRhs", "jqlEqualsOperator",
		"jqlLikeOperator", "jqlComparisonOperator", "jqlInOperator", "jqlIsOperator",
		"jqlWasOperator", "jqlWasInOperator", "jqlChangedOperator", "jqlField",
		"jqlFieldProperty", "jqlCustomField", "jqlString", "jqlNumber", "jqlOperand",
		"jqlEmpty", "jqlValue", "jqlFunction", "jqlFunctionName", "jqlArgumentList",
		"jqlList", "jqlListStart", "jqlListEnd", "jqlPropertyArgument", "jqlArgument",
		"jqlWasPredicate", "jqlChangedPredicate", "jqlDatePredicateOperator",
		"jqlDateRangePredicateOperator", "jqlUserPredicateOperator", "jqlValuePredicateOperator",
		"jqlPredicateOperand", "jqlOrderBy", "jqlSearchSort",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 46, 332, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 1, 0, 3, 0, 84,
		8, 0, 1, 0, 3, 0, 87, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 5,
		2, 96, 8, 2, 10, 2, 12, 2, 99, 9, 2, 1, 3, 1, 3, 1, 3, 5, 3, 104, 8, 3,
		10, 3, 12, 3, 107, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 113, 8, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 126,
		8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 132, 8, 7, 1, 7, 1, 7, 1, 7, 3, 7,
		137, 8, 7, 1, 7, 1, 7, 1, 7, 3, 7, 142, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 3, 7, 151, 8, 7, 1, 7, 5, 7, 154, 8, 7, 10, 7, 12, 7, 157,
		9, 7, 1, 7, 1, 7, 1, 7, 3, 7, 162, 8, 7, 1, 7, 5, 7, 165, 8, 7, 10, 7,
		12, 7, 168, 9, 7, 1, 7, 1, 7, 5, 7, 172, 8, 7, 10, 7, 12, 7, 175, 9, 7,
		3, 7, 177, 8, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		11, 3, 11, 188, 8, 11, 1, 12, 1, 12, 1, 12, 3, 12, 193, 8, 12, 1, 13, 1,
		13, 1, 13, 3, 13, 198, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14,
		205, 8, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 3, 16, 212, 8, 16, 1, 16,
		5, 16, 215, 8, 16, 10, 16, 12, 16, 218, 9, 16, 3, 16, 220, 8, 16, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 227, 8, 17, 10, 17, 12, 17, 230, 9,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 21, 1, 21, 3, 21, 245, 8, 21, 1, 22, 1, 22, 1, 23, 1, 23, 3,
		23, 251, 8, 23, 1, 24, 1, 24, 1, 24, 3, 24, 256, 8, 24, 1, 24, 1, 24, 1,
		25, 1, 25, 3, 25, 262, 8, 25, 1, 26, 1, 26, 1, 26, 5, 26, 267, 8, 26, 10,
		26, 12, 26, 270, 9, 26, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 276, 8, 27,
		10, 27, 12, 27, 279, 9, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 31, 1, 31, 3, 31, 291, 8, 31, 1, 32, 1, 32, 1, 32, 3, 32,
		296, 8, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 304, 8, 33,
		1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1,
		38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 5, 39, 323, 8, 39, 10, 39,
		12, 39, 326, 9, 39, 1, 40, 1, 40, 3, 40, 330, 8, 40, 1, 40, 0, 0, 41, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74,
		76, 78, 80, 0, 9, 2, 0, 6, 6, 19, 19, 1, 0, 11, 12, 1, 0, 13, 14, 1, 0,
		7, 10, 3, 0, 37, 37, 41, 41, 44, 44, 1, 0, 33, 34, 2, 0, 23, 24, 27, 27,
		1, 0, 25, 26, 1, 0, 31, 32, 339, 0, 83, 1, 0, 0, 0, 2, 90, 1, 0, 0, 0,
		4, 92, 1, 0, 0, 0, 6, 100, 1, 0, 0, 0, 8, 112, 1, 0, 0, 0, 10, 114, 1,
		0, 0, 0, 12, 118, 1, 0, 0, 0, 14, 176, 1, 0, 0, 0, 16, 178, 1, 0, 0, 0,
		18, 180, 1, 0, 0, 0, 20, 182, 1, 0, 0, 0, 22, 187, 1, 0, 0, 0, 24, 192,
		1, 0, 0, 0, 26, 197, 1, 0, 0, 0, 28, 204, 1, 0, 0, 0, 30, 206, 1, 0, 0,
		0, 32, 219, 1, 0, 0, 0, 34, 221, 1, 0, 0, 0, 36, 231, 1, 0, 0, 0, 38, 236,
		1, 0, 0, 0, 40, 238, 1, 0, 0, 0, 42, 244, 1, 0, 0, 0, 44, 246, 1, 0, 0,
		0, 46, 250, 1, 0, 0, 0, 48, 252, 1, 0, 0, 0, 50, 261, 1, 0, 0, 0, 52, 263,
		1, 0, 0, 0, 54, 271, 1, 0, 0, 0, 56, 282, 1, 0, 0, 0, 58, 284, 1, 0, 0,
		0, 60, 286, 1, 0, 0, 0, 62, 290, 1, 0, 0, 0, 64, 295, 1, 0, 0, 0, 66, 303,
		1, 0, 0, 0, 68, 307, 1, 0, 0, 0, 70, 309, 1, 0, 0, 0, 72, 311, 1, 0, 0,
		0, 74, 313, 1, 0, 0, 0, 76, 315, 1, 0, 0, 0, 78, 317, 1, 0, 0, 0, 80, 327,
		1, 0, 0, 0, 82, 84, 3, 2, 1, 0, 83, 82, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0,
		84, 86, 1, 0, 0, 0, 85, 87, 3, 78, 39, 0, 86, 85, 1, 0, 0, 0, 86, 87, 1,
		0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 5, 0, 0, 1, 89, 1, 1, 0, 0, 0, 90,
		91, 3, 4, 2, 0, 91, 3, 1, 0, 0, 0, 92, 97, 3, 6, 3, 0, 93, 94, 5, 18, 0,
		0, 94, 96, 3, 6, 3, 0, 95, 93, 1, 0, 0, 0, 96, 99, 1, 0, 0, 0, 97, 95,
		1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 5, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0,
		100, 105, 3, 8, 4, 0, 101, 102, 5, 17, 0, 0, 102, 104, 3, 8, 4, 0, 103,
		101, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105, 106,
		1, 0, 0, 0, 106, 7, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108, 109, 7, 0, 0,
		0, 109, 113, 3, 8, 4, 0, 110, 113, 3, 10, 5, 0, 111, 113, 3, 12, 6, 0,
		112, 108, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 111, 1, 0, 0, 0, 113,
		9, 1, 0, 0, 0, 114, 115, 5, 1, 0, 0, 115, 116, 3, 4, 2, 0, 116, 117, 5,
		2, 0, 0, 117, 11, 1, 0, 0, 0, 118, 119, 3, 32, 16, 0, 119, 120, 3, 14,
		7, 0, 120, 13, 1, 0, 0, 0, 121, 125, 3, 16, 8, 0, 122, 126, 3, 44, 22,
		0, 123, 126, 3, 46, 23, 0, 124, 126, 3, 48, 24, 0, 125, 122, 1, 0, 0, 0,
		125, 123, 1, 0, 0, 0, 125, 124, 1, 0, 0, 0, 126, 177, 1, 0, 0, 0, 127,
		131, 3, 18, 9, 0, 128, 132, 3, 44, 22, 0, 129, 132, 3, 46, 23, 0, 130,
		132, 3, 48, 24, 0, 131, 128, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131, 130,
		1, 0, 0, 0, 132, 177, 1, 0, 0, 0, 133, 136, 3, 20, 10, 0, 134, 137, 3,
		46, 23, 0, 135, 137, 3, 48, 24, 0, 136, 134, 1, 0, 0, 0, 136, 135, 1, 0,
		0, 0, 137, 177, 1, 0, 0, 0, 138, 141, 3, 22, 11, 0, 139, 142, 3, 54, 27,
		0, 140, 142, 3, 48, 24, 0, 141, 139, 1, 0, 0, 0, 141, 140, 1, 0, 0, 0,
		142, 177, 1, 0, 0, 0, 143, 144, 3, 24, 12, 0, 144, 145, 3, 44, 22, 0, 145,
		177, 1, 0, 0, 0, 146, 150, 3, 26, 13, 0, 147, 151, 3, 44, 22, 0, 148, 151,
		3, 46, 23, 0, 149, 151, 3, 48, 24, 0, 150, 147, 1, 0, 0, 0, 150, 148, 1,
		0, 0, 0, 150, 149, 1, 0, 0, 0, 151, 155, 1, 0, 0, 0, 152, 154, 3, 64, 32,
		0, 153, 152, 1, 0, 0, 0, 154, 157, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 155,
		156, 1, 0, 0, 0, 156, 177, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 158, 161,
		3, 28, 14, 0, 159, 162, 3, 54, 27, 0, 160, 162, 3, 48, 24, 0, 161, 159,
		1, 0, 0, 0, 161, 160, 1, 0, 0, 0, 162, 166, 1, 0, 0, 0, 163, 165, 3, 64,
		32, 0, 164, 163, 1, 0, 0, 0, 165, 168, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0,
		166, 167, 1, 0, 0, 0, 167, 177, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 169,
		173, 3, 30, 15, 0, 170, 172, 3, 66, 33, 0, 171, 170, 1, 0, 0, 0, 172, 175,
		1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 177, 1, 0,
		0, 0, 175, 173, 1, 0, 0, 0, 176, 121, 1, 0, 0, 0, 176, 127, 1, 0, 0, 0,
		176, 133, 1, 0, 0, 0, 176, 138, 1, 0, 0, 0, 176, 143, 1, 0, 0, 0, 176,
		146, 1, 0, 0, 0, 176, 158, 1, 0, 0, 0, 176, 169, 1, 0, 0, 0, 177, 15, 1,
		0, 0, 0, 178, 179, 7, 1, 0, 0, 179, 17, 1, 0, 0, 0, 180, 181, 7, 2, 0,
		0, 181, 19, 1, 0, 0, 0, 182, 183, 7, 3, 0, 0, 183, 21, 1, 0, 0, 0, 184,
		188, 5, 15, 0, 0, 185, 186, 5, 19, 0, 0, 186, 188, 5, 15, 0, 0, 187, 184,
		1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 188, 23, 1, 0, 0, 0, 189, 193, 5, 16,
		0, 0, 190, 191, 5, 16, 0, 0, 191, 193, 5, 19, 0, 0, 192, 189, 1, 0, 0,
		0, 192, 190, 1, 0, 0, 0, 193, 25, 1, 0, 0, 0, 194, 198, 5, 21, 0, 0, 195,
		196, 5, 21, 0, 0, 196, 198, 5, 19, 0, 0, 197, 194, 1, 0, 0, 0, 197, 195,
		1, 0, 0, 0, 198, 27, 1, 0, 0, 0, 199, 200, 5, 21, 0, 0, 200, 205, 5, 15,
		0, 0, 201, 202, 5, 21, 0, 0, 202, 203, 5, 19, 0, 0, 203, 205, 5, 15, 0,
		0, 204, 199, 1, 0, 0, 0, 204, 201, 1, 0, 0, 0, 205, 29, 1, 0, 0, 0, 206,
		207, 5, 22, 0, 0, 207, 31, 1, 0, 0, 0, 208, 220, 3, 40, 20, 0, 209, 212,
		3, 38, 19, 0, 210, 212, 3, 36, 18, 0, 211, 209, 1, 0, 0, 0, 211, 210, 1,
		0, 0, 0, 212, 216, 1, 0, 0, 0, 213, 215, 3, 34, 17, 0, 214, 213, 1, 0,
		0, 0, 215, 218, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0,
		217, 220, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 219, 208, 1, 0, 0, 0, 219,
		211, 1, 0, 0, 0, 220, 33, 1, 0, 0, 0, 221, 222, 5, 4, 0, 0, 222, 223, 3,
		62, 31, 0, 223, 224, 5, 5, 0, 0, 224, 228, 1, 0, 0, 0, 225, 227, 3, 60,
		30, 0, 226, 225, 1, 0, 0, 0, 227, 230, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0,
		228, 229, 1, 0, 0, 0, 229, 35, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 231, 232,
		5, 35, 0, 0, 232, 233, 5, 4, 0, 0, 233, 234, 5, 33, 0, 0, 234, 235, 5,
		5, 0, 0, 235, 37, 1, 0, 0, 0, 236, 237, 7, 4, 0, 0, 237, 39, 1, 0, 0, 0,
		238, 239, 7, 5, 0, 0, 239, 41, 1, 0, 0, 0, 240, 245, 3, 44, 22, 0, 241,
		245, 3, 46, 23, 0, 242, 245, 3, 48, 24, 0, 243, 245, 3, 54, 27, 0, 244,
		240, 1, 0, 0, 0, 244, 241, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 244, 243,
		1, 0, 0, 0, 245, 43, 1, 0, 0, 0, 246, 247, 5, 20, 0, 0, 247, 45, 1, 0,
		0, 0, 248, 251, 3, 38, 19, 0, 249, 251, 3, 40, 20, 0, 250, 248, 1, 0, 0,
		0, 250, 249, 1, 0, 0, 0, 251, 47, 1, 0, 0, 0, 252, 253, 3, 50, 25, 0, 253,
		255, 5, 1, 0, 0, 254, 256, 3, 52, 26, 0, 255, 254, 1, 0, 0, 0, 255, 256,
		1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 258, 5, 2, 0, 0, 258, 49, 1, 0,
		0, 0, 259, 262, 3, 38, 19, 0, 260, 262, 3, 40, 20, 0, 261, 259, 1, 0, 0,
		0, 261, 260, 1, 0, 0, 0, 262, 51, 1, 0, 0, 0, 263, 268, 3, 62, 31, 0, 264,
		265, 5, 3, 0, 0, 265, 267, 3, 62, 31, 0, 266, 264, 1, 0, 0, 0, 267, 270,
		1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 53, 1, 0,
		0, 0, 270, 268, 1, 0, 0, 0, 271, 272, 3, 56, 28, 0, 272, 277, 3, 42, 21,
		0, 273, 274, 5, 3, 0, 0, 274, 276, 3, 42, 21, 0, 275, 273, 1, 0, 0, 0,
		276, 279, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 277, 278, 1, 0, 0, 0, 278,
		280, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 280, 281, 3, 58, 29, 0, 281, 55,
		1, 0, 0, 0, 282, 283, 5, 1, 0, 0, 283, 57, 1, 0, 0, 0, 284, 285, 5, 2,
		0, 0, 285, 59, 1, 0, 0, 0, 286, 287, 3, 62, 31, 0, 287, 61, 1, 0, 0, 0,
		288, 291, 3, 38, 19, 0, 289, 291, 3, 40, 20, 0, 290, 288, 1, 0, 0, 0, 290,
		289, 1, 0, 0, 0, 291, 63, 1, 0, 0, 0, 292, 296, 3, 68, 34, 0, 293, 296,
		3, 70, 35, 0, 294, 296, 3, 72, 36, 0, 295, 292, 1, 0, 0, 0, 295, 293, 1,
		0, 0, 0, 295, 294, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297, 298, 3, 76, 38,
		0, 298, 65, 1, 0, 0, 0, 299, 304, 3, 68, 34, 0, 300, 304, 3, 70, 35, 0,
		301, 304, 3, 72, 36, 0, 302, 304, 3, 74, 37, 0, 303, 299, 1, 0, 0, 0, 303,
		300, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 303, 302, 1, 0, 0, 0, 304, 305,
		1, 0, 0, 0, 305, 306, 3, 76, 38, 0, 306, 67, 1, 0, 0, 0, 307, 308, 7, 6,
		0, 0, 308, 69, 1, 0, 0, 0, 309, 310, 5, 28, 0, 0, 310, 71, 1, 0, 0, 0,
		311, 312, 5, 30, 0, 0, 312, 73, 1, 0, 0, 0, 313, 314, 7, 7, 0, 0, 314,
		75, 1, 0, 0, 0, 315, 316, 3, 42, 21, 0, 316, 77, 1, 0, 0, 0, 317, 318,
		5, 29, 0, 0, 318, 319, 5, 30, 0, 0, 319, 324, 3, 80, 40, 0, 320, 321, 5,
		3, 0, 0, 321, 323, 3, 80, 40, 0, 322, 320, 1, 0, 0, 0, 323, 326, 1, 0,
		0, 0, 324, 322, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 79, 1, 0, 0, 0,
		326, 324, 1, 0, 0, 0, 327, 329, 3, 32, 16, 0, 328, 330, 7, 8, 0, 0, 329,
		328, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330, 81, 1, 0, 0, 0, 34, 83, 86,
		97, 105, 112, 125, 131, 136, 141, 150, 155, 161, 166, 173, 176, 187, 192,
		197, 204, 211, 216, 219, 228, 244, 250, 255, 261, 268, 277, 290, 295, 303,
		324, 329,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// JQLParserInit initializes any static state used to implement JQLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJQLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JQLParserInit() {
	staticData := &jqlparserParserStaticData
	staticData.once.Do(jqlparserParserInit)
}

// NewJQLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJQLParser(input antlr.TokenStream) *JQLParser {
	JQLParserInit()
	this := new(JQLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &jqlparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "JQLParser.g4"

	return this
}

// JQLParser tokens.
const (
	JQLParserEOF                    = antlr.TokenEOF
	JQLParserLPAREN                 = 1
	JQLParserRPAREN                 = 2
	JQLParserCOMMA                  = 3
	JQLParserLBRACKET               = 4
	JQLParserRBRACKET               = 5
	JQLParserBANG                   = 6
	JQLParserLT                     = 7
	JQLParserGT                     = 8
	JQLParserGTEQ                   = 9
	JQLParserLTEQ                   = 10
	JQLParserEQUALS                 = 11
	JQLParserNOT_EQUALS             = 12
	JQLParserLIKE                   = 13
	JQLParserNOT_LIKE               = 14
	JQLParserIN                     = 15
	JQLParserIS                     = 16
	JQLParserAND                    = 17
	JQLParserOR                     = 18
	JQLParserNOT                    = 19
	JQLParserEMPTY                  = 20
	JQLParserWAS                    = 21
	JQLParserCHANGED                = 22
	JQLParserBEFORE                 = 23
	JQLParserAFTER                  = 24
	JQLParserFROM                   = 25
	JQLParserTO                     = 26
	JQLParserON                     = 27
	JQLParserDURING                 = 28
	JQLParserORDER                  = 29
	JQLParserBY                     = 30
	JQLParserASC                    = 31
	JQLParserDESC                   = 32
	JQLParserPOSNUMBER              = 33
	JQLParserNEGNUMBER              = 34
	JQLParserCUSTOMFIELD            = 35
	JQLParserRESERVED_WORD          = 36
	JQLParserSTRING                 = 37
	JQLParserMATCHWS                = 38
	JQLParserERROR_RESERVED         = 39
	JQLParserERRORCHAR              = 40
	JQLParserQUOTE_STRING           = 41
	JQLParserUNCLOSED_QUOTE_STRING  = 42
	JQLParserINVALID_QUOTE_STRING   = 43
	JQLParserSQUOTE_STRING          = 44
	JQLParserUNCLOSED_SQUOTE_STRING = 45
	JQLParserINVALID_SQUOTE_STRING  = 46
)

// JQLParser rules.
const (
	JQLParserRULE_jqlQuery                      = 0
	JQLParserRULE_jqlWhere                      = 1
	JQLParserRULE_jqlOrClause                   = 2
	JQLParserRULE_jqlAndClause                  = 3
	JQLParserRULE_jqlNotClause                  = 4
	JQLParserRULE_jqlSubClause                  = 5
	JQLParserRULE_jqlTerminalClause             = 6
	JQLParserRULE_jqlTerminalClauseRhs          = 7
	JQLParserRULE_jqlEqualsOperator             = 8
	JQLParserRULE_jqlLikeOperator               = 9
	JQLParserRULE_jqlComparisonOperator         = 10
	JQLParserRULE_jqlInOperator                 = 11
	JQLParserRULE_jqlIsOperator                 = 12
	JQLParserRULE_jqlWasOperator                = 13
	JQLParserRULE_jqlWasInOperator              = 14
	JQLParserRULE_jqlChangedOperator            = 15
	JQLParserRULE_jqlField                      = 16
	JQLParserRULE_jqlFieldProperty              = 17
	JQLParserRULE_jqlCustomField                = 18
	JQLParserRULE_jqlString                     = 19
	JQLParserRULE_jqlNumber                     = 20
	JQLParserRULE_jqlOperand                    = 21
	JQLParserRULE_jqlEmpty                      = 22
	JQLParserRULE_jqlValue                      = 23
	JQLParserRULE_jqlFunction                   = 24
	JQLParserRULE_jqlFunctionName               = 25
	JQLParserRULE_jqlArgumentList               = 26
	JQLParserRULE_jqlList                       = 27
	JQLParserRULE_jqlListStart                  = 28
	JQLParserRULE_jqlListEnd                    = 29
	JQLParserRULE_jqlPropertyArgument           = 30
	JQLParserRULE_jqlArgument                   = 31
	JQLParserRULE_jqlWasPredicate               = 32
	JQLParserRULE_jqlChangedPredicate           = 33
	JQLParserRULE_jqlDatePredicateOperator      = 34
	JQLParserRULE_jqlDateRangePredicateOperator = 35
	JQLParserRULE_jqlUserPredicateOperator      = 36
	JQLParserRULE_jqlValuePredicateOperator     = 37
	JQLParserRULE_jqlPredicateOperand           = 38
	JQLParserRULE_jqlOrderBy                    = 39
	JQLParserRULE_jqlSearchSort                 = 40
)

// IJqlQueryContext is an interface to support dynamic dispatch.
type IJqlQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	JqlWhere() IJqlWhereContext
	JqlOrderBy() IJqlOrderByContext

	// IsJqlQueryContext differentiates from other interfaces.
	IsJqlQueryContext()
}

type JqlQueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlQueryContext() *JqlQueryContext {
	var p = new(JqlQueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlQuery
	return p
}

func (*JqlQueryContext) IsJqlQueryContext() {}

func NewJqlQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlQueryContext {
	var p = new(JqlQueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlQuery

	return p
}

func (s *JqlQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlQueryContext) EOF() antlr.TerminalNode {
	return s.GetToken(JQLParserEOF, 0)
}

func (s *JqlQueryContext) JqlWhere() IJqlWhereContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlWhereContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlWhereContext)
}

func (s *JqlQueryContext) JqlOrderBy() IJqlOrderByContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlOrderByContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlOrderByContext)
}

func (s *JqlQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlQuery(s)
	}
}

func (s *JqlQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlQuery(s)
	}
}

func (p *JQLParser) JqlQuery() (localctx IJqlQueryContext) {
	this := p
	_ = this

	localctx = NewJqlQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JQLParserRULE_jqlQuery)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&19988778319938) != 0 {
		{
			p.SetState(82)
			p.JqlWhere()
		}

	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JQLParserORDER {
		{
			p.SetState(85)
			p.JqlOrderBy()
		}

	}
	{
		p.SetState(88)
		p.Match(JQLParserEOF)
	}

	return localctx
}

// IJqlWhereContext is an interface to support dynamic dispatch.
type IJqlWhereContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JqlOrClause() IJqlOrClauseContext

	// IsJqlWhereContext differentiates from other interfaces.
	IsJqlWhereContext()
}

type JqlWhereContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlWhereContext() *JqlWhereContext {
	var p = new(JqlWhereContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlWhere
	return p
}

func (*JqlWhereContext) IsJqlWhereContext() {}

func NewJqlWhereContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlWhereContext {
	var p = new(JqlWhereContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlWhere

	return p
}

func (s *JqlWhereContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlWhereContext) JqlOrClause() IJqlOrClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlOrClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlOrClauseContext)
}

func (s *JqlWhereContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlWhereContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlWhereContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlWhere(s)
	}
}

func (s *JqlWhereContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlWhere(s)
	}
}

func (p *JQLParser) JqlWhere() (localctx IJqlWhereContext) {
	this := p
	_ = this

	localctx = NewJqlWhereContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JQLParserRULE_jqlWhere)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.JqlOrClause()
	}

	return localctx
}

// IJqlOrClauseContext is an interface to support dynamic dispatch.
type IJqlOrClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllJqlAndClause() []IJqlAndClauseContext
	JqlAndClause(i int) IJqlAndClauseContext
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsJqlOrClauseContext differentiates from other interfaces.
	IsJqlOrClauseContext()
}

type JqlOrClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlOrClauseContext() *JqlOrClauseContext {
	var p = new(JqlOrClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlOrClause
	return p
}

func (*JqlOrClauseContext) IsJqlOrClauseContext() {}

func NewJqlOrClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlOrClauseContext {
	var p = new(JqlOrClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlOrClause

	return p
}

func (s *JqlOrClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlOrClauseContext) AllJqlAndClause() []IJqlAndClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJqlAndClauseContext); ok {
			len++
		}
	}

	tst := make([]IJqlAndClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJqlAndClauseContext); ok {
			tst[i] = t.(IJqlAndClauseContext)
			i++
		}
	}

	return tst
}

func (s *JqlOrClauseContext) JqlAndClause(i int) IJqlAndClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlAndClauseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlAndClauseContext)
}

func (s *JqlOrClauseContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(JQLParserOR)
}

func (s *JqlOrClauseContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(JQLParserOR, i)
}

func (s *JqlOrClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlOrClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlOrClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlOrClause(s)
	}
}

func (s *JqlOrClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlOrClause(s)
	}
}

func (p *JQLParser) JqlOrClause() (localctx IJqlOrClauseContext) {
	this := p
	_ = this

	localctx = NewJqlOrClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JQLParserRULE_jqlOrClause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.JqlAndClause()
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JQLParserOR {
		{
			p.SetState(93)
			p.Match(JQLParserOR)
		}
		{
			p.SetState(94)
			p.JqlAndClause()
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IJqlAndClauseContext is an interface to support dynamic dispatch.
type IJqlAndClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllJqlNotClause() []IJqlNotClauseContext
	JqlNotClause(i int) IJqlNotClauseContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode

	// IsJqlAndClauseContext differentiates from other interfaces.
	IsJqlAndClauseContext()
}

type JqlAndClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlAndClauseContext() *JqlAndClauseContext {
	var p = new(JqlAndClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlAndClause
	return p
}

func (*JqlAndClauseContext) IsJqlAndClauseContext() {}

func NewJqlAndClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlAndClauseContext {
	var p = new(JqlAndClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlAndClause

	return p
}

func (s *JqlAndClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlAndClauseContext) AllJqlNotClause() []IJqlNotClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJqlNotClauseContext); ok {
			len++
		}
	}

	tst := make([]IJqlNotClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJqlNotClauseContext); ok {
			tst[i] = t.(IJqlNotClauseContext)
			i++
		}
	}

	return tst
}

func (s *JqlAndClauseContext) JqlNotClause(i int) IJqlNotClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlNotClauseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlNotClauseContext)
}

func (s *JqlAndClauseContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(JQLParserAND)
}

func (s *JqlAndClauseContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(JQLParserAND, i)
}

func (s *JqlAndClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlAndClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlAndClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlAndClause(s)
	}
}

func (s *JqlAndClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlAndClause(s)
	}
}

func (p *JQLParser) JqlAndClause() (localctx IJqlAndClauseContext) {
	this := p
	_ = this

	localctx = NewJqlAndClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JQLParserRULE_jqlAndClause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.JqlNotClause()
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JQLParserAND {
		{
			p.SetState(101)
			p.Match(JQLParserAND)
		}
		{
			p.SetState(102)
			p.JqlNotClause()
		}

		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IJqlNotClauseContext is an interface to support dynamic dispatch.
type IJqlNotClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JqlNotClause() IJqlNotClauseContext
	NOT() antlr.TerminalNode
	BANG() antlr.TerminalNode
	JqlSubClause() IJqlSubClauseContext
	JqlTerminalClause() IJqlTerminalClauseContext

	// IsJqlNotClauseContext differentiates from other interfaces.
	IsJqlNotClauseContext()
}

type JqlNotClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlNotClauseContext() *JqlNotClauseContext {
	var p = new(JqlNotClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlNotClause
	return p
}

func (*JqlNotClauseContext) IsJqlNotClauseContext() {}

func NewJqlNotClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlNotClauseContext {
	var p = new(JqlNotClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlNotClause

	return p
}

func (s *JqlNotClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlNotClauseContext) JqlNotClause() IJqlNotClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlNotClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlNotClauseContext)
}

func (s *JqlNotClauseContext) NOT() antlr.TerminalNode {
	return s.GetToken(JQLParserNOT, 0)
}

func (s *JqlNotClauseContext) BANG() antlr.TerminalNode {
	return s.GetToken(JQLParserBANG, 0)
}

func (s *JqlNotClauseContext) JqlSubClause() IJqlSubClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlSubClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlSubClauseContext)
}

func (s *JqlNotClauseContext) JqlTerminalClause() IJqlTerminalClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlTerminalClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlTerminalClauseContext)
}

func (s *JqlNotClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlNotClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlNotClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlNotClause(s)
	}
}

func (s *JqlNotClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlNotClause(s)
	}
}

func (p *JQLParser) JqlNotClause() (localctx IJqlNotClauseContext) {
	this := p
	_ = this

	localctx = NewJqlNotClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JQLParserRULE_jqlNotClause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(112)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JQLParserBANG, JQLParserNOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			_la = p.GetTokenStream().LA(1)

			if !(_la == JQLParserBANG || _la == JQLParserNOT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(109)
			p.JqlNotClause()
		}

	case JQLParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.JqlSubClause()
		}

	case JQLParserPOSNUMBER, JQLParserNEGNUMBER, JQLParserCUSTOMFIELD, JQLParserSTRING, JQLParserQUOTE_STRING, JQLParserSQUOTE_STRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(111)
			p.JqlTerminalClause()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IJqlSubClauseContext is an interface to support dynamic dispatch.
type IJqlSubClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	JqlOrClause() IJqlOrClauseContext
	RPAREN() antlr.TerminalNode

	// IsJqlSubClauseContext differentiates from other interfaces.
	IsJqlSubClauseContext()
}

type JqlSubClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlSubClauseContext() *JqlSubClauseContext {
	var p = new(JqlSubClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlSubClause
	return p
}

func (*JqlSubClauseContext) IsJqlSubClauseContext() {}

func NewJqlSubClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlSubClauseContext {
	var p = new(JqlSubClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlSubClause

	return p
}

func (s *JqlSubClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlSubClauseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JQLParserLPAREN, 0)
}

func (s *JqlSubClauseContext) JqlOrClause() IJqlOrClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlOrClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlOrClauseContext)
}

func (s *JqlSubClauseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JQLParserRPAREN, 0)
}

func (s *JqlSubClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlSubClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlSubClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlSubClause(s)
	}
}

func (s *JqlSubClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlSubClause(s)
	}
}

func (p *JQLParser) JqlSubClause() (localctx IJqlSubClauseContext) {
	this := p
	_ = this

	localctx = NewJqlSubClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JQLParserRULE_jqlSubClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(JQLParserLPAREN)
	}
	{
		p.SetState(115)
		p.JqlOrClause()
	}
	{
		p.SetState(116)
		p.Match(JQLParserRPAREN)
	}

	return localctx
}

// IJqlTerminalClauseContext is an interface to support dynamic dispatch.
type IJqlTerminalClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JqlField() IJqlFieldContext
	JqlTerminalClauseRhs() IJqlTerminalClauseRhsContext

	// IsJqlTerminalClauseContext differentiates from other interfaces.
	IsJqlTerminalClauseContext()
}

type JqlTerminalClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlTerminalClauseContext() *JqlTerminalClauseContext {
	var p = new(JqlTerminalClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlTerminalClause
	return p
}

func (*JqlTerminalClauseContext) IsJqlTerminalClauseContext() {}

func NewJqlTerminalClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlTerminalClauseContext {
	var p = new(JqlTerminalClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlTerminalClause

	return p
}

func (s *JqlTerminalClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlTerminalClauseContext) JqlField() IJqlFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlFieldContext)
}

func (s *JqlTerminalClauseContext) JqlTerminalClauseRhs() IJqlTerminalClauseRhsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlTerminalClauseRhsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlTerminalClauseRhsContext)
}

func (s *JqlTerminalClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlTerminalClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlTerminalClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlTerminalClause(s)
	}
}

func (s *JqlTerminalClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlTerminalClause(s)
	}
}

func (p *JQLParser) JqlTerminalClause() (localctx IJqlTerminalClauseContext) {
	this := p
	_ = this

	localctx = NewJqlTerminalClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JQLParserRULE_jqlTerminalClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.JqlField()
	}
	{
		p.SetState(119)
		p.JqlTerminalClauseRhs()
	}

	return localctx
}

// IJqlTerminalClauseRhsContext is an interface to support dynamic dispatch.
type IJqlTerminalClauseRhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsJqlTerminalClauseRhsContext differentiates from other interfaces.
	IsJqlTerminalClauseRhsContext()
}

type JqlTerminalClauseRhsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlTerminalClauseRhsContext() *JqlTerminalClauseRhsContext {
	var p = new(JqlTerminalClauseRhsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlTerminalClauseRhs
	return p
}

func (*JqlTerminalClauseRhsContext) IsJqlTerminalClauseRhsContext() {}

func NewJqlTerminalClauseRhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlTerminalClauseRhsContext {
	var p = new(JqlTerminalClauseRhsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlTerminalClauseRhs

	return p
}

func (s *JqlTerminalClauseRhsContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlTerminalClauseRhsContext) CopyFrom(ctx *JqlTerminalClauseRhsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *JqlTerminalClauseRhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlTerminalClauseRhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type JqlWasClauseContext struct {
	*JqlTerminalClauseRhsContext
}

func NewJqlWasClauseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JqlWasClauseContext {
	var p = new(JqlWasClauseContext)

	p.JqlTerminalClauseRhsContext = NewEmptyJqlTerminalClauseRhsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JqlTerminalClauseRhsContext))

	return p
}

func (s *JqlWasClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlWasClauseContext) JqlWasOperator() IJqlWasOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlWasOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlWasOperatorContext)
}

func (s *JqlWasClauseContext) JqlEmpty() IJqlEmptyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlEmptyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlEmptyContext)
}

func (s *JqlWasClauseContext) JqlValue() IJqlValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlValueContext)
}

func (s *JqlWasClauseContext) JqlFunction() IJqlFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlFunctionContext)
}

func (s *JqlWasClauseContext) AllJqlWasPredicate() []IJqlWasPredicateContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJqlWasPredicateContext); ok {
			len++
		}
	}

	tst := make([]IJqlWasPredicateContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJqlWasPredicateContext); ok {
			tst[i] = t.(IJqlWasPredicateContext)
			i++
		}
	}

	return tst
}

func (s *JqlWasClauseContext) JqlWasPredicate(i int) IJqlWasPredicateContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlWasPredicateContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlWasPredicateContext)
}

func (s *JqlWasClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlWasClause(s)
	}
}

func (s *JqlWasClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlWasClause(s)
	}
}

type JqlLikeClauseContext struct {
	*JqlTerminalClauseRhsContext
}

func NewJqlLikeClauseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JqlLikeClauseContext {
	var p = new(JqlLikeClauseContext)

	p.JqlTerminalClauseRhsContext = NewEmptyJqlTerminalClauseRhsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JqlTerminalClauseRhsContext))

	return p
}

func (s *JqlLikeClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlLikeClauseContext) JqlLikeOperator() IJqlLikeOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlLikeOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlLikeOperatorContext)
}

func (s *JqlLikeClauseContext) JqlEmpty() IJqlEmptyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlEmptyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlEmptyContext)
}

func (s *JqlLikeClauseContext) JqlValue() IJqlValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlValueContext)
}

func (s *JqlLikeClauseContext) JqlFunction() IJqlFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlFunctionContext)
}

func (s *JqlLikeClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlLikeClause(s)
	}
}

func (s *JqlLikeClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlLikeClause(s)
	}
}

type JqlWasInClauseContext struct {
	*JqlTerminalClauseRhsContext
}

func NewJqlWasInClauseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JqlWasInClauseContext {
	var p = new(JqlWasInClauseContext)

	p.JqlTerminalClauseRhsContext = NewEmptyJqlTerminalClauseRhsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JqlTerminalClauseRhsContext))

	return p
}

func (s *JqlWasInClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlWasInClauseContext) JqlWasInOperator() IJqlWasInOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlWasInOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlWasInOperatorContext)
}

func (s *JqlWasInClauseContext) JqlList() IJqlListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlListContext)
}

func (s *JqlWasInClauseContext) JqlFunction() IJqlFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlFunctionContext)
}

func (s *JqlWasInClauseContext) AllJqlWasPredicate() []IJqlWasPredicateContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJqlWasPredicateContext); ok {
			len++
		}
	}

	tst := make([]IJqlWasPredicateContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJqlWasPredicateContext); ok {
			tst[i] = t.(IJqlWasPredicateContext)
			i++
		}
	}

	return tst
}

func (s *JqlWasInClauseContext) JqlWasPredicate(i int) IJqlWasPredicateContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlWasPredicateContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlWasPredicateContext)
}

func (s *JqlWasInClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlWasInClause(s)
	}
}

func (s *JqlWasInClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlWasInClause(s)
	}
}

type JqlComparisonClauseContext struct {
	*JqlTerminalClauseRhsContext
}

func NewJqlComparisonClauseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JqlComparisonClauseContext {
	var p = new(JqlComparisonClauseContext)

	p.JqlTerminalClauseRhsContext = NewEmptyJqlTerminalClauseRhsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JqlTerminalClauseRhsContext))

	return p
}

func (s *JqlComparisonClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlComparisonClauseContext) JqlComparisonOperator() IJqlComparisonOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlComparisonOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlComparisonOperatorContext)
}

func (s *JqlComparisonClauseContext) JqlValue() IJqlValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlValueContext)
}

func (s *JqlComparisonClauseContext) JqlFunction() IJqlFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlFunctionContext)
}

func (s *JqlComparisonClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlComparisonClause(s)
	}
}

func (s *JqlComparisonClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlComparisonClause(s)
	}
}

type JqlEqualsClauseContext struct {
	*JqlTerminalClauseRhsContext
}

func NewJqlEqualsClauseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JqlEqualsClauseContext {
	var p = new(JqlEqualsClauseContext)

	p.JqlTerminalClauseRhsContext = NewEmptyJqlTerminalClauseRhsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JqlTerminalClauseRhsContext))

	return p
}

func (s *JqlEqualsClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlEqualsClauseContext) JqlEqualsOperator() IJqlEqualsOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlEqualsOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlEqualsOperatorContext)
}

func (s *JqlEqualsClauseContext) JqlEmpty() IJqlEmptyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlEmptyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlEmptyContext)
}

func (s *JqlEqualsClauseContext) JqlValue() IJqlValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlValueContext)
}

func (s *JqlEqualsClauseContext) JqlFunction() IJqlFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlFunctionContext)
}

func (s *JqlEqualsClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlEqualsClause(s)
	}
}

func (s *JqlEqualsClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlEqualsClause(s)
	}
}

type JqlInClauseContext struct {
	*JqlTerminalClauseRhsContext
}

func NewJqlInClauseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JqlInClauseContext {
	var p = new(JqlInClauseContext)

	p.JqlTerminalClauseRhsContext = NewEmptyJqlTerminalClauseRhsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JqlTerminalClauseRhsContext))

	return p
}

func (s *JqlInClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlInClauseContext) JqlInOperator() IJqlInOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlInOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlInOperatorContext)
}

func (s *JqlInClauseContext) JqlList() IJqlListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlListContext)
}

func (s *JqlInClauseContext) JqlFunction() IJqlFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlFunctionContext)
}

func (s *JqlInClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlInClause(s)
	}
}

func (s *JqlInClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlInClause(s)
	}
}

type JqlIsClauseContext struct {
	*JqlTerminalClauseRhsContext
}

func NewJqlIsClauseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JqlIsClauseContext {
	var p = new(JqlIsClauseContext)

	p.JqlTerminalClauseRhsContext = NewEmptyJqlTerminalClauseRhsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JqlTerminalClauseRhsContext))

	return p
}

func (s *JqlIsClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlIsClauseContext) JqlIsOperator() IJqlIsOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlIsOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlIsOperatorContext)
}

func (s *JqlIsClauseContext) JqlEmpty() IJqlEmptyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlEmptyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlEmptyContext)
}

func (s *JqlIsClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlIsClause(s)
	}
}

func (s *JqlIsClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlIsClause(s)
	}
}

type JqlChangedClauseContext struct {
	*JqlTerminalClauseRhsContext
}

func NewJqlChangedClauseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JqlChangedClauseContext {
	var p = new(JqlChangedClauseContext)

	p.JqlTerminalClauseRhsContext = NewEmptyJqlTerminalClauseRhsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JqlTerminalClauseRhsContext))

	return p
}

func (s *JqlChangedClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlChangedClauseContext) JqlChangedOperator() IJqlChangedOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlChangedOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlChangedOperatorContext)
}

func (s *JqlChangedClauseContext) AllJqlChangedPredicate() []IJqlChangedPredicateContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJqlChangedPredicateContext); ok {
			len++
		}
	}

	tst := make([]IJqlChangedPredicateContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJqlChangedPredicateContext); ok {
			tst[i] = t.(IJqlChangedPredicateContext)
			i++
		}
	}

	return tst
}

func (s *JqlChangedClauseContext) JqlChangedPredicate(i int) IJqlChangedPredicateContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlChangedPredicateContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlChangedPredicateContext)
}

func (s *JqlChangedClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlChangedClause(s)
	}
}

func (s *JqlChangedClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlChangedClause(s)
	}
}

func (p *JQLParser) JqlTerminalClauseRhs() (localctx IJqlTerminalClauseRhsContext) {
	this := p
	_ = this

	localctx = NewJqlTerminalClauseRhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JQLParserRULE_jqlTerminalClauseRhs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewJqlEqualsClauseContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.JqlEqualsOperator()
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(122)
				p.JqlEmpty()
			}

		case 2:
			{
				p.SetState(123)
				p.JqlValue()
			}

		case 3:
			{
				p.SetState(124)
				p.JqlFunction()
			}

		}

	case 2:
		localctx = NewJqlLikeClauseContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(127)
			p.JqlLikeOperator()
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(128)
				p.JqlEmpty()
			}

		case 2:
			{
				p.SetState(129)
				p.JqlValue()
			}

		case 3:
			{
				p.SetState(130)
				p.JqlFunction()
			}

		}

	case 3:
		localctx = NewJqlComparisonClauseContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(133)
			p.JqlComparisonOperator()
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(134)
				p.JqlValue()
			}

		case 2:
			{
				p.SetState(135)
				p.JqlFunction()
			}

		}

	case 4:
		localctx = NewJqlInClauseContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(138)
			p.JqlInOperator()
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case JQLParserLPAREN:
			{
				p.SetState(139)
				p.JqlList()
			}

		case JQLParserPOSNUMBER, JQLParserNEGNUMBER, JQLParserSTRING, JQLParserQUOTE_STRING, JQLParserSQUOTE_STRING:
			{
				p.SetState(140)
				p.JqlFunction()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	case 5:
		localctx = NewJqlIsClauseContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(143)
			p.JqlIsOperator()
		}
		{
			p.SetState(144)
			p.JqlEmpty()
		}

	case 6:
		localctx = NewJqlWasClauseContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(146)
			p.JqlWasOperator()
		}
		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(147)
				p.JqlEmpty()
			}

		case 2:
			{
				p.SetState(148)
				p.JqlValue()
			}

		case 3:
			{
				p.SetState(149)
				p.JqlFunction()
			}

		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1501560832) != 0 {
			{
				p.SetState(152)
				p.JqlWasPredicate()
			}

			p.SetState(157)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 7:
		localctx = NewJqlWasInClauseContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(158)
			p.JqlWasInOperator()
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case JQLParserLPAREN:
			{
				p.SetState(159)
				p.JqlList()
			}

		case JQLParserPOSNUMBER, JQLParserNEGNUMBER, JQLParserSTRING, JQLParserQUOTE_STRING, JQLParserSQUOTE_STRING:
			{
				p.SetState(160)
				p.JqlFunction()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1501560832) != 0 {
			{
				p.SetState(163)
				p.JqlWasPredicate()
			}

			p.SetState(168)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 8:
		localctx = NewJqlChangedClauseContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(169)
			p.JqlChangedOperator()
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1602224128) != 0 {
			{
				p.SetState(170)
				p.JqlChangedPredicate()
			}

			p.SetState(175)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IJqlEqualsOperatorContext is an interface to support dynamic dispatch.
type IJqlEqualsOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQUALS() antlr.TerminalNode
	NOT_EQUALS() antlr.TerminalNode

	// IsJqlEqualsOperatorContext differentiates from other interfaces.
	IsJqlEqualsOperatorContext()
}

type JqlEqualsOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlEqualsOperatorContext() *JqlEqualsOperatorContext {
	var p = new(JqlEqualsOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlEqualsOperator
	return p
}

func (*JqlEqualsOperatorContext) IsJqlEqualsOperatorContext() {}

func NewJqlEqualsOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlEqualsOperatorContext {
	var p = new(JqlEqualsOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlEqualsOperator

	return p
}

func (s *JqlEqualsOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlEqualsOperatorContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(JQLParserEQUALS, 0)
}

func (s *JqlEqualsOperatorContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(JQLParserNOT_EQUALS, 0)
}

func (s *JqlEqualsOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlEqualsOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlEqualsOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlEqualsOperator(s)
	}
}

func (s *JqlEqualsOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlEqualsOperator(s)
	}
}

func (p *JQLParser) JqlEqualsOperator() (localctx IJqlEqualsOperatorContext) {
	this := p
	_ = this

	localctx = NewJqlEqualsOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JQLParserRULE_jqlEqualsOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		_la = p.GetTokenStream().LA(1)

		if !(_la == JQLParserEQUALS || _la == JQLParserNOT_EQUALS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IJqlLikeOperatorContext is an interface to support dynamic dispatch.
type IJqlLikeOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LIKE() antlr.TerminalNode
	NOT_LIKE() antlr.TerminalNode

	// IsJqlLikeOperatorContext differentiates from other interfaces.
	IsJqlLikeOperatorContext()
}

type JqlLikeOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlLikeOperatorContext() *JqlLikeOperatorContext {
	var p = new(JqlLikeOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlLikeOperator
	return p
}

func (*JqlLikeOperatorContext) IsJqlLikeOperatorContext() {}

func NewJqlLikeOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlLikeOperatorContext {
	var p = new(JqlLikeOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlLikeOperator

	return p
}

func (s *JqlLikeOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlLikeOperatorContext) LIKE() antlr.TerminalNode {
	return s.GetToken(JQLParserLIKE, 0)
}

func (s *JqlLikeOperatorContext) NOT_LIKE() antlr.TerminalNode {
	return s.GetToken(JQLParserNOT_LIKE, 0)
}

func (s *JqlLikeOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlLikeOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlLikeOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlLikeOperator(s)
	}
}

func (s *JqlLikeOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlLikeOperator(s)
	}
}

func (p *JQLParser) JqlLikeOperator() (localctx IJqlLikeOperatorContext) {
	this := p
	_ = this

	localctx = NewJqlLikeOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JQLParserRULE_jqlLikeOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		_la = p.GetTokenStream().LA(1)

		if !(_la == JQLParserLIKE || _la == JQLParserNOT_LIKE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IJqlComparisonOperatorContext is an interface to support dynamic dispatch.
type IJqlComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	LTEQ() antlr.TerminalNode
	GTEQ() antlr.TerminalNode

	// IsJqlComparisonOperatorContext differentiates from other interfaces.
	IsJqlComparisonOperatorContext()
}

type JqlComparisonOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlComparisonOperatorContext() *JqlComparisonOperatorContext {
	var p = new(JqlComparisonOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlComparisonOperator
	return p
}

func (*JqlComparisonOperatorContext) IsJqlComparisonOperatorContext() {}

func NewJqlComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlComparisonOperatorContext {
	var p = new(JqlComparisonOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlComparisonOperator

	return p
}

func (s *JqlComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlComparisonOperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(JQLParserLT, 0)
}

func (s *JqlComparisonOperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(JQLParserGT, 0)
}

func (s *JqlComparisonOperatorContext) LTEQ() antlr.TerminalNode {
	return s.GetToken(JQLParserLTEQ, 0)
}

func (s *JqlComparisonOperatorContext) GTEQ() antlr.TerminalNode {
	return s.GetToken(JQLParserGTEQ, 0)
}

func (s *JqlComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlComparisonOperator(s)
	}
}

func (s *JqlComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlComparisonOperator(s)
	}
}

func (p *JQLParser) JqlComparisonOperator() (localctx IJqlComparisonOperatorContext) {
	this := p
	_ = this

	localctx = NewJqlComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JQLParserRULE_jqlComparisonOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1920) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IJqlInOperatorContext is an interface to support dynamic dispatch.
type IJqlInOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IN() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsJqlInOperatorContext differentiates from other interfaces.
	IsJqlInOperatorContext()
}

type JqlInOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlInOperatorContext() *JqlInOperatorContext {
	var p = new(JqlInOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlInOperator
	return p
}

func (*JqlInOperatorContext) IsJqlInOperatorContext() {}

func NewJqlInOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlInOperatorContext {
	var p = new(JqlInOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlInOperator

	return p
}

func (s *JqlInOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlInOperatorContext) IN() antlr.TerminalNode {
	return s.GetToken(JQLParserIN, 0)
}

func (s *JqlInOperatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(JQLParserNOT, 0)
}

func (s *JqlInOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlInOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlInOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlInOperator(s)
	}
}

func (s *JqlInOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlInOperator(s)
	}
}

func (p *JQLParser) JqlInOperator() (localctx IJqlInOperatorContext) {
	this := p
	_ = this

	localctx = NewJqlInOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JQLParserRULE_jqlInOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(187)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JQLParserIN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(184)
			p.Match(JQLParserIN)
		}

	case JQLParserNOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(185)
			p.Match(JQLParserNOT)
		}
		{
			p.SetState(186)
			p.Match(JQLParserIN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IJqlIsOperatorContext is an interface to support dynamic dispatch.
type IJqlIsOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IS() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsJqlIsOperatorContext differentiates from other interfaces.
	IsJqlIsOperatorContext()
}

type JqlIsOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlIsOperatorContext() *JqlIsOperatorContext {
	var p = new(JqlIsOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlIsOperator
	return p
}

func (*JqlIsOperatorContext) IsJqlIsOperatorContext() {}

func NewJqlIsOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlIsOperatorContext {
	var p = new(JqlIsOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlIsOperator

	return p
}

func (s *JqlIsOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlIsOperatorContext) IS() antlr.TerminalNode {
	return s.GetToken(JQLParserIS, 0)
}

func (s *JqlIsOperatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(JQLParserNOT, 0)
}

func (s *JqlIsOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlIsOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlIsOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlIsOperator(s)
	}
}

func (s *JqlIsOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlIsOperator(s)
	}
}

func (p *JQLParser) JqlIsOperator() (localctx IJqlIsOperatorContext) {
	this := p
	_ = this

	localctx = NewJqlIsOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JQLParserRULE_jqlIsOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.Match(JQLParserIS)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Match(JQLParserIS)
		}
		{
			p.SetState(191)
			p.Match(JQLParserNOT)
		}

	}

	return localctx
}

// IJqlWasOperatorContext is an interface to support dynamic dispatch.
type IJqlWasOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WAS() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsJqlWasOperatorContext differentiates from other interfaces.
	IsJqlWasOperatorContext()
}

type JqlWasOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlWasOperatorContext() *JqlWasOperatorContext {
	var p = new(JqlWasOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlWasOperator
	return p
}

func (*JqlWasOperatorContext) IsJqlWasOperatorContext() {}

func NewJqlWasOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlWasOperatorContext {
	var p = new(JqlWasOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlWasOperator

	return p
}

func (s *JqlWasOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlWasOperatorContext) WAS() antlr.TerminalNode {
	return s.GetToken(JQLParserWAS, 0)
}

func (s *JqlWasOperatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(JQLParserNOT, 0)
}

func (s *JqlWasOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlWasOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlWasOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlWasOperator(s)
	}
}

func (s *JqlWasOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlWasOperator(s)
	}
}

func (p *JQLParser) JqlWasOperator() (localctx IJqlWasOperatorContext) {
	this := p
	_ = this

	localctx = NewJqlWasOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, JQLParserRULE_jqlWasOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(194)
			p.Match(JQLParserWAS)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(195)
			p.Match(JQLParserWAS)
		}
		{
			p.SetState(196)
			p.Match(JQLParserNOT)
		}

	}

	return localctx
}

// IJqlWasInOperatorContext is an interface to support dynamic dispatch.
type IJqlWasInOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WAS() antlr.TerminalNode
	IN() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsJqlWasInOperatorContext differentiates from other interfaces.
	IsJqlWasInOperatorContext()
}

type JqlWasInOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlWasInOperatorContext() *JqlWasInOperatorContext {
	var p = new(JqlWasInOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlWasInOperator
	return p
}

func (*JqlWasInOperatorContext) IsJqlWasInOperatorContext() {}

func NewJqlWasInOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlWasInOperatorContext {
	var p = new(JqlWasInOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlWasInOperator

	return p
}

func (s *JqlWasInOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlWasInOperatorContext) WAS() antlr.TerminalNode {
	return s.GetToken(JQLParserWAS, 0)
}

func (s *JqlWasInOperatorContext) IN() antlr.TerminalNode {
	return s.GetToken(JQLParserIN, 0)
}

func (s *JqlWasInOperatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(JQLParserNOT, 0)
}

func (s *JqlWasInOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlWasInOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlWasInOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlWasInOperator(s)
	}
}

func (s *JqlWasInOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlWasInOperator(s)
	}
}

func (p *JQLParser) JqlWasInOperator() (localctx IJqlWasInOperatorContext) {
	this := p
	_ = this

	localctx = NewJqlWasInOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, JQLParserRULE_jqlWasInOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.Match(JQLParserWAS)
		}
		{
			p.SetState(200)
			p.Match(JQLParserIN)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(201)
			p.Match(JQLParserWAS)
		}
		{
			p.SetState(202)
			p.Match(JQLParserNOT)
		}
		{
			p.SetState(203)
			p.Match(JQLParserIN)
		}

	}

	return localctx
}

// IJqlChangedOperatorContext is an interface to support dynamic dispatch.
type IJqlChangedOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CHANGED() antlr.TerminalNode

	// IsJqlChangedOperatorContext differentiates from other interfaces.
	IsJqlChangedOperatorContext()
}

type JqlChangedOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlChangedOperatorContext() *JqlChangedOperatorContext {
	var p = new(JqlChangedOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlChangedOperator
	return p
}

func (*JqlChangedOperatorContext) IsJqlChangedOperatorContext() {}

func NewJqlChangedOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlChangedOperatorContext {
	var p = new(JqlChangedOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlChangedOperator

	return p
}

func (s *JqlChangedOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlChangedOperatorContext) CHANGED() antlr.TerminalNode {
	return s.GetToken(JQLParserCHANGED, 0)
}

func (s *JqlChangedOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlChangedOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlChangedOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlChangedOperator(s)
	}
}

func (s *JqlChangedOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlChangedOperator(s)
	}
}

func (p *JQLParser) JqlChangedOperator() (localctx IJqlChangedOperatorContext) {
	this := p
	_ = this

	localctx = NewJqlChangedOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, JQLParserRULE_jqlChangedOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(JQLParserCHANGED)
	}

	return localctx
}

// IJqlFieldContext is an interface to support dynamic dispatch.
type IJqlFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsJqlFieldContext differentiates from other interfaces.
	IsJqlFieldContext()
}

type JqlFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlFieldContext() *JqlFieldContext {
	var p = new(JqlFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlField
	return p
}

func (*JqlFieldContext) IsJqlFieldContext() {}

func NewJqlFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlFieldContext {
	var p = new(JqlFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlField

	return p
}

func (s *JqlFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlFieldContext) CopyFrom(ctx *JqlFieldContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *JqlFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type JqlNumberFieldContext struct {
	*JqlFieldContext
}

func NewJqlNumberFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JqlNumberFieldContext {
	var p = new(JqlNumberFieldContext)

	p.JqlFieldContext = NewEmptyJqlFieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JqlFieldContext))

	return p
}

func (s *JqlNumberFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlNumberFieldContext) JqlNumber() IJqlNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlNumberContext)
}

func (s *JqlNumberFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlNumberField(s)
	}
}

func (s *JqlNumberFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlNumberField(s)
	}
}

type JqlNonNumberFieldContext struct {
	*JqlFieldContext
}

func NewJqlNonNumberFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JqlNonNumberFieldContext {
	var p = new(JqlNonNumberFieldContext)

	p.JqlFieldContext = NewEmptyJqlFieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JqlFieldContext))

	return p
}

func (s *JqlNonNumberFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlNonNumberFieldContext) JqlString() IJqlStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlStringContext)
}

func (s *JqlNonNumberFieldContext) JqlCustomField() IJqlCustomFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlCustomFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlCustomFieldContext)
}

func (s *JqlNonNumberFieldContext) AllJqlFieldProperty() []IJqlFieldPropertyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJqlFieldPropertyContext); ok {
			len++
		}
	}

	tst := make([]IJqlFieldPropertyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJqlFieldPropertyContext); ok {
			tst[i] = t.(IJqlFieldPropertyContext)
			i++
		}
	}

	return tst
}

func (s *JqlNonNumberFieldContext) JqlFieldProperty(i int) IJqlFieldPropertyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlFieldPropertyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlFieldPropertyContext)
}

func (s *JqlNonNumberFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlNonNumberField(s)
	}
}

func (s *JqlNonNumberFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlNonNumberField(s)
	}
}

func (p *JQLParser) JqlField() (localctx IJqlFieldContext) {
	this := p
	_ = this

	localctx = NewJqlFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, JQLParserRULE_jqlField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(219)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JQLParserPOSNUMBER, JQLParserNEGNUMBER:
		localctx = NewJqlNumberFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(208)
			p.JqlNumber()
		}

	case JQLParserCUSTOMFIELD, JQLParserSTRING, JQLParserQUOTE_STRING, JQLParserSQUOTE_STRING:
		localctx = NewJqlNonNumberFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(211)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case JQLParserSTRING, JQLParserQUOTE_STRING, JQLParserSQUOTE_STRING:
			{
				p.SetState(209)
				p.JqlString()
			}

		case JQLParserCUSTOMFIELD:
			{
				p.SetState(210)
				p.JqlCustomField()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JQLParserLBRACKET {
			{
				p.SetState(213)
				p.JqlFieldProperty()
			}

			p.SetState(218)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IJqlFieldPropertyContext is an interface to support dynamic dispatch.
type IJqlFieldPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	AllJqlPropertyArgument() []IJqlPropertyArgumentContext
	JqlPropertyArgument(i int) IJqlPropertyArgumentContext
	JqlArgument() IJqlArgumentContext

	// IsJqlFieldPropertyContext differentiates from other interfaces.
	IsJqlFieldPropertyContext()
}

type JqlFieldPropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlFieldPropertyContext() *JqlFieldPropertyContext {
	var p = new(JqlFieldPropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlFieldProperty
	return p
}

func (*JqlFieldPropertyContext) IsJqlFieldPropertyContext() {}

func NewJqlFieldPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlFieldPropertyContext {
	var p = new(JqlFieldPropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlFieldProperty

	return p
}

func (s *JqlFieldPropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlFieldPropertyContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(JQLParserLBRACKET, 0)
}

func (s *JqlFieldPropertyContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(JQLParserRBRACKET, 0)
}

func (s *JqlFieldPropertyContext) AllJqlPropertyArgument() []IJqlPropertyArgumentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJqlPropertyArgumentContext); ok {
			len++
		}
	}

	tst := make([]IJqlPropertyArgumentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJqlPropertyArgumentContext); ok {
			tst[i] = t.(IJqlPropertyArgumentContext)
			i++
		}
	}

	return tst
}

func (s *JqlFieldPropertyContext) JqlPropertyArgument(i int) IJqlPropertyArgumentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlPropertyArgumentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlPropertyArgumentContext)
}

func (s *JqlFieldPropertyContext) JqlArgument() IJqlArgumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlArgumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlArgumentContext)
}

func (s *JqlFieldPropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlFieldPropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlFieldPropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlFieldProperty(s)
	}
}

func (s *JqlFieldPropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlFieldProperty(s)
	}
}

func (p *JQLParser) JqlFieldProperty() (localctx IJqlFieldPropertyContext) {
	this := p
	_ = this

	localctx = NewJqlFieldPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, JQLParserRULE_jqlFieldProperty)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Match(JQLParserLBRACKET)
	}

	{
		p.SetState(222)
		p.JqlArgument()
	}

	{
		p.SetState(223)
		p.Match(JQLParserRBRACKET)
	}

	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&19954418057216) != 0 {
		{
			p.SetState(225)
			p.JqlPropertyArgument()
		}

		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IJqlCustomFieldContext is an interface to support dynamic dispatch.
type IJqlCustomFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CUSTOMFIELD() antlr.TerminalNode
	LBRACKET() antlr.TerminalNode
	POSNUMBER() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode

	// IsJqlCustomFieldContext differentiates from other interfaces.
	IsJqlCustomFieldContext()
}

type JqlCustomFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlCustomFieldContext() *JqlCustomFieldContext {
	var p = new(JqlCustomFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlCustomField
	return p
}

func (*JqlCustomFieldContext) IsJqlCustomFieldContext() {}

func NewJqlCustomFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlCustomFieldContext {
	var p = new(JqlCustomFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlCustomField

	return p
}

func (s *JqlCustomFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlCustomFieldContext) CUSTOMFIELD() antlr.TerminalNode {
	return s.GetToken(JQLParserCUSTOMFIELD, 0)
}

func (s *JqlCustomFieldContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(JQLParserLBRACKET, 0)
}

func (s *JqlCustomFieldContext) POSNUMBER() antlr.TerminalNode {
	return s.GetToken(JQLParserPOSNUMBER, 0)
}

func (s *JqlCustomFieldContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(JQLParserRBRACKET, 0)
}

func (s *JqlCustomFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlCustomFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlCustomFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlCustomField(s)
	}
}

func (s *JqlCustomFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlCustomField(s)
	}
}

func (p *JQLParser) JqlCustomField() (localctx IJqlCustomFieldContext) {
	this := p
	_ = this

	localctx = NewJqlCustomFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, JQLParserRULE_jqlCustomField)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.Match(JQLParserCUSTOMFIELD)
	}
	{
		p.SetState(232)
		p.Match(JQLParserLBRACKET)
	}
	{
		p.SetState(233)
		p.Match(JQLParserPOSNUMBER)
	}
	{
		p.SetState(234)
		p.Match(JQLParserRBRACKET)
	}

	return localctx
}

// IJqlStringContext is an interface to support dynamic dispatch.
type IJqlStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	QUOTE_STRING() antlr.TerminalNode
	SQUOTE_STRING() antlr.TerminalNode

	// IsJqlStringContext differentiates from other interfaces.
	IsJqlStringContext()
}

type JqlStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlStringContext() *JqlStringContext {
	var p = new(JqlStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlString
	return p
}

func (*JqlStringContext) IsJqlStringContext() {}

func NewJqlStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlStringContext {
	var p = new(JqlStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlString

	return p
}

func (s *JqlStringContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlStringContext) STRING() antlr.TerminalNode {
	return s.GetToken(JQLParserSTRING, 0)
}

func (s *JqlStringContext) QUOTE_STRING() antlr.TerminalNode {
	return s.GetToken(JQLParserQUOTE_STRING, 0)
}

func (s *JqlStringContext) SQUOTE_STRING() antlr.TerminalNode {
	return s.GetToken(JQLParserSQUOTE_STRING, 0)
}

func (s *JqlStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlString(s)
	}
}

func (s *JqlStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlString(s)
	}
}

func (p *JQLParser) JqlString() (localctx IJqlStringContext) {
	this := p
	_ = this

	localctx = NewJqlStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, JQLParserRULE_jqlString)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&19928648253440) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IJqlNumberContext is an interface to support dynamic dispatch.
type IJqlNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetJqlNum returns the jqlNum token.
	GetJqlNum() antlr.Token

	// SetJqlNum sets the jqlNum token.
	SetJqlNum(antlr.Token)

	// Getter signatures
	POSNUMBER() antlr.TerminalNode
	NEGNUMBER() antlr.TerminalNode

	// IsJqlNumberContext differentiates from other interfaces.
	IsJqlNumberContext()
}

type JqlNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	jqlNum antlr.Token
}

func NewEmptyJqlNumberContext() *JqlNumberContext {
	var p = new(JqlNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlNumber
	return p
}

func (*JqlNumberContext) IsJqlNumberContext() {}

func NewJqlNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlNumberContext {
	var p = new(JqlNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlNumber

	return p
}

func (s *JqlNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlNumberContext) GetJqlNum() antlr.Token { return s.jqlNum }

func (s *JqlNumberContext) SetJqlNum(v antlr.Token) { s.jqlNum = v }

func (s *JqlNumberContext) POSNUMBER() antlr.TerminalNode {
	return s.GetToken(JQLParserPOSNUMBER, 0)
}

func (s *JqlNumberContext) NEGNUMBER() antlr.TerminalNode {
	return s.GetToken(JQLParserNEGNUMBER, 0)
}

func (s *JqlNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlNumber(s)
	}
}

func (s *JqlNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlNumber(s)
	}
}

func (p *JQLParser) JqlNumber() (localctx IJqlNumberContext) {
	this := p
	_ = this

	localctx = NewJqlNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, JQLParserRULE_jqlNumber)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*JqlNumberContext).jqlNum = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == JQLParserPOSNUMBER || _la == JQLParserNEGNUMBER) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*JqlNumberContext).jqlNum = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IJqlOperandContext is an interface to support dynamic dispatch.
type IJqlOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JqlEmpty() IJqlEmptyContext
	JqlValue() IJqlValueContext
	JqlFunction() IJqlFunctionContext
	JqlList() IJqlListContext

	// IsJqlOperandContext differentiates from other interfaces.
	IsJqlOperandContext()
}

type JqlOperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlOperandContext() *JqlOperandContext {
	var p = new(JqlOperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlOperand
	return p
}

func (*JqlOperandContext) IsJqlOperandContext() {}

func NewJqlOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlOperandContext {
	var p = new(JqlOperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlOperand

	return p
}

func (s *JqlOperandContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlOperandContext) JqlEmpty() IJqlEmptyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlEmptyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlEmptyContext)
}

func (s *JqlOperandContext) JqlValue() IJqlValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlValueContext)
}

func (s *JqlOperandContext) JqlFunction() IJqlFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlFunctionContext)
}

func (s *JqlOperandContext) JqlList() IJqlListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlListContext)
}

func (s *JqlOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlOperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlOperand(s)
	}
}

func (s *JqlOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlOperand(s)
	}
}

func (p *JQLParser) JqlOperand() (localctx IJqlOperandContext) {
	this := p
	_ = this

	localctx = NewJqlOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, JQLParserRULE_jqlOperand)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(240)
			p.JqlEmpty()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(241)
			p.JqlValue()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(242)
			p.JqlFunction()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(243)
			p.JqlList()
		}

	}

	return localctx
}

// IJqlEmptyContext is an interface to support dynamic dispatch.
type IJqlEmptyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EMPTY() antlr.TerminalNode

	// IsJqlEmptyContext differentiates from other interfaces.
	IsJqlEmptyContext()
}

type JqlEmptyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlEmptyContext() *JqlEmptyContext {
	var p = new(JqlEmptyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlEmpty
	return p
}

func (*JqlEmptyContext) IsJqlEmptyContext() {}

func NewJqlEmptyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlEmptyContext {
	var p = new(JqlEmptyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlEmpty

	return p
}

func (s *JqlEmptyContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlEmptyContext) EMPTY() antlr.TerminalNode {
	return s.GetToken(JQLParserEMPTY, 0)
}

func (s *JqlEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlEmptyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlEmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlEmpty(s)
	}
}

func (s *JqlEmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlEmpty(s)
	}
}

func (p *JQLParser) JqlEmpty() (localctx IJqlEmptyContext) {
	this := p
	_ = this

	localctx = NewJqlEmptyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, JQLParserRULE_jqlEmpty)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Match(JQLParserEMPTY)
	}

	return localctx
}

// IJqlValueContext is an interface to support dynamic dispatch.
type IJqlValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JqlString() IJqlStringContext
	JqlNumber() IJqlNumberContext

	// IsJqlValueContext differentiates from other interfaces.
	IsJqlValueContext()
}

type JqlValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlValueContext() *JqlValueContext {
	var p = new(JqlValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlValue
	return p
}

func (*JqlValueContext) IsJqlValueContext() {}

func NewJqlValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlValueContext {
	var p = new(JqlValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlValue

	return p
}

func (s *JqlValueContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlValueContext) JqlString() IJqlStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlStringContext)
}

func (s *JqlValueContext) JqlNumber() IJqlNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlNumberContext)
}

func (s *JqlValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlValue(s)
	}
}

func (s *JqlValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlValue(s)
	}
}

func (p *JQLParser) JqlValue() (localctx IJqlValueContext) {
	this := p
	_ = this

	localctx = NewJqlValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, JQLParserRULE_jqlValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(250)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JQLParserSTRING, JQLParserQUOTE_STRING, JQLParserSQUOTE_STRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(248)
			p.JqlString()
		}

	case JQLParserPOSNUMBER, JQLParserNEGNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(249)
			p.JqlNumber()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IJqlFunctionContext is an interface to support dynamic dispatch.
type IJqlFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JqlFunctionName() IJqlFunctionNameContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	JqlArgumentList() IJqlArgumentListContext

	// IsJqlFunctionContext differentiates from other interfaces.
	IsJqlFunctionContext()
}

type JqlFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlFunctionContext() *JqlFunctionContext {
	var p = new(JqlFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlFunction
	return p
}

func (*JqlFunctionContext) IsJqlFunctionContext() {}

func NewJqlFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlFunctionContext {
	var p = new(JqlFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlFunction

	return p
}

func (s *JqlFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlFunctionContext) JqlFunctionName() IJqlFunctionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlFunctionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlFunctionNameContext)
}

func (s *JqlFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JQLParserLPAREN, 0)
}

func (s *JqlFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JQLParserRPAREN, 0)
}

func (s *JqlFunctionContext) JqlArgumentList() IJqlArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlArgumentListContext)
}

func (s *JqlFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlFunction(s)
	}
}

func (s *JqlFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlFunction(s)
	}
}

func (p *JQLParser) JqlFunction() (localctx IJqlFunctionContext) {
	this := p
	_ = this

	localctx = NewJqlFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, JQLParserRULE_jqlFunction)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.JqlFunctionName()
	}
	{
		p.SetState(253)
		p.Match(JQLParserLPAREN)
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&19954418057216) != 0 {
		{
			p.SetState(254)
			p.JqlArgumentList()
		}

	}
	{
		p.SetState(257)
		p.Match(JQLParserRPAREN)
	}

	return localctx
}

// IJqlFunctionNameContext is an interface to support dynamic dispatch.
type IJqlFunctionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JqlString() IJqlStringContext
	JqlNumber() IJqlNumberContext

	// IsJqlFunctionNameContext differentiates from other interfaces.
	IsJqlFunctionNameContext()
}

type JqlFunctionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlFunctionNameContext() *JqlFunctionNameContext {
	var p = new(JqlFunctionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlFunctionName
	return p
}

func (*JqlFunctionNameContext) IsJqlFunctionNameContext() {}

func NewJqlFunctionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlFunctionNameContext {
	var p = new(JqlFunctionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlFunctionName

	return p
}

func (s *JqlFunctionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlFunctionNameContext) JqlString() IJqlStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlStringContext)
}

func (s *JqlFunctionNameContext) JqlNumber() IJqlNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlNumberContext)
}

func (s *JqlFunctionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlFunctionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlFunctionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlFunctionName(s)
	}
}

func (s *JqlFunctionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlFunctionName(s)
	}
}

func (p *JQLParser) JqlFunctionName() (localctx IJqlFunctionNameContext) {
	this := p
	_ = this

	localctx = NewJqlFunctionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, JQLParserRULE_jqlFunctionName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(261)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JQLParserSTRING, JQLParserQUOTE_STRING, JQLParserSQUOTE_STRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(259)
			p.JqlString()
		}

	case JQLParserPOSNUMBER, JQLParserNEGNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(260)
			p.JqlNumber()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IJqlArgumentListContext is an interface to support dynamic dispatch.
type IJqlArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllJqlArgument() []IJqlArgumentContext
	JqlArgument(i int) IJqlArgumentContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsJqlArgumentListContext differentiates from other interfaces.
	IsJqlArgumentListContext()
}

type JqlArgumentListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlArgumentListContext() *JqlArgumentListContext {
	var p = new(JqlArgumentListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlArgumentList
	return p
}

func (*JqlArgumentListContext) IsJqlArgumentListContext() {}

func NewJqlArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlArgumentListContext {
	var p = new(JqlArgumentListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlArgumentList

	return p
}

func (s *JqlArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlArgumentListContext) AllJqlArgument() []IJqlArgumentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJqlArgumentContext); ok {
			len++
		}
	}

	tst := make([]IJqlArgumentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJqlArgumentContext); ok {
			tst[i] = t.(IJqlArgumentContext)
			i++
		}
	}

	return tst
}

func (s *JqlArgumentListContext) JqlArgument(i int) IJqlArgumentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlArgumentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlArgumentContext)
}

func (s *JqlArgumentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JQLParserCOMMA)
}

func (s *JqlArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JQLParserCOMMA, i)
}

func (s *JqlArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlArgumentList(s)
	}
}

func (s *JqlArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlArgumentList(s)
	}
}

func (p *JQLParser) JqlArgumentList() (localctx IJqlArgumentListContext) {
	this := p
	_ = this

	localctx = NewJqlArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, JQLParserRULE_jqlArgumentList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.JqlArgument()
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JQLParserCOMMA {
		{
			p.SetState(264)
			p.Match(JQLParserCOMMA)
		}
		{
			p.SetState(265)
			p.JqlArgument()
		}

		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IJqlListContext is an interface to support dynamic dispatch.
type IJqlListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JqlListStart() IJqlListStartContext
	AllJqlOperand() []IJqlOperandContext
	JqlOperand(i int) IJqlOperandContext
	JqlListEnd() IJqlListEndContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsJqlListContext differentiates from other interfaces.
	IsJqlListContext()
}

type JqlListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlListContext() *JqlListContext {
	var p = new(JqlListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlList
	return p
}

func (*JqlListContext) IsJqlListContext() {}

func NewJqlListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlListContext {
	var p = new(JqlListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlList

	return p
}

func (s *JqlListContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlListContext) JqlListStart() IJqlListStartContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlListStartContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlListStartContext)
}

func (s *JqlListContext) AllJqlOperand() []IJqlOperandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJqlOperandContext); ok {
			len++
		}
	}

	tst := make([]IJqlOperandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJqlOperandContext); ok {
			tst[i] = t.(IJqlOperandContext)
			i++
		}
	}

	return tst
}

func (s *JqlListContext) JqlOperand(i int) IJqlOperandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlOperandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlOperandContext)
}

func (s *JqlListContext) JqlListEnd() IJqlListEndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlListEndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlListEndContext)
}

func (s *JqlListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JQLParserCOMMA)
}

func (s *JqlListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JQLParserCOMMA, i)
}

func (s *JqlListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlList(s)
	}
}

func (s *JqlListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlList(s)
	}
}

func (p *JQLParser) JqlList() (localctx IJqlListContext) {
	this := p
	_ = this

	localctx = NewJqlListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, JQLParserRULE_jqlList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		p.JqlListStart()
	}
	{
		p.SetState(272)
		p.JqlOperand()
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JQLParserCOMMA {
		{
			p.SetState(273)
			p.Match(JQLParserCOMMA)
		}
		{
			p.SetState(274)
			p.JqlOperand()
		}

		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(280)
		p.JqlListEnd()
	}

	return localctx
}

// IJqlListStartContext is an interface to support dynamic dispatch.
type IJqlListStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode

	// IsJqlListStartContext differentiates from other interfaces.
	IsJqlListStartContext()
}

type JqlListStartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlListStartContext() *JqlListStartContext {
	var p = new(JqlListStartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlListStart
	return p
}

func (*JqlListStartContext) IsJqlListStartContext() {}

func NewJqlListStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlListStartContext {
	var p = new(JqlListStartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlListStart

	return p
}

func (s *JqlListStartContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlListStartContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JQLParserLPAREN, 0)
}

func (s *JqlListStartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlListStartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlListStartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlListStart(s)
	}
}

func (s *JqlListStartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlListStart(s)
	}
}

func (p *JQLParser) JqlListStart() (localctx IJqlListStartContext) {
	this := p
	_ = this

	localctx = NewJqlListStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, JQLParserRULE_jqlListStart)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Match(JQLParserLPAREN)
	}

	return localctx
}

// IJqlListEndContext is an interface to support dynamic dispatch.
type IJqlListEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RPAREN() antlr.TerminalNode

	// IsJqlListEndContext differentiates from other interfaces.
	IsJqlListEndContext()
}

type JqlListEndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlListEndContext() *JqlListEndContext {
	var p = new(JqlListEndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlListEnd
	return p
}

func (*JqlListEndContext) IsJqlListEndContext() {}

func NewJqlListEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlListEndContext {
	var p = new(JqlListEndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlListEnd

	return p
}

func (s *JqlListEndContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlListEndContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JQLParserRPAREN, 0)
}

func (s *JqlListEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlListEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlListEndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlListEnd(s)
	}
}

func (s *JqlListEndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlListEnd(s)
	}
}

func (p *JQLParser) JqlListEnd() (localctx IJqlListEndContext) {
	this := p
	_ = this

	localctx = NewJqlListEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, JQLParserRULE_jqlListEnd)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.Match(JQLParserRPAREN)
	}

	return localctx
}

// IJqlPropertyArgumentContext is an interface to support dynamic dispatch.
type IJqlPropertyArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JqlArgument() IJqlArgumentContext

	// IsJqlPropertyArgumentContext differentiates from other interfaces.
	IsJqlPropertyArgumentContext()
}

type JqlPropertyArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlPropertyArgumentContext() *JqlPropertyArgumentContext {
	var p = new(JqlPropertyArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlPropertyArgument
	return p
}

func (*JqlPropertyArgumentContext) IsJqlPropertyArgumentContext() {}

func NewJqlPropertyArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlPropertyArgumentContext {
	var p = new(JqlPropertyArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlPropertyArgument

	return p
}

func (s *JqlPropertyArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlPropertyArgumentContext) JqlArgument() IJqlArgumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlArgumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlArgumentContext)
}

func (s *JqlPropertyArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlPropertyArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlPropertyArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlPropertyArgument(s)
	}
}

func (s *JqlPropertyArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlPropertyArgument(s)
	}
}

func (p *JQLParser) JqlPropertyArgument() (localctx IJqlPropertyArgumentContext) {
	this := p
	_ = this

	localctx = NewJqlPropertyArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, JQLParserRULE_jqlPropertyArgument)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.JqlArgument()
	}

	return localctx
}

// IJqlArgumentContext is an interface to support dynamic dispatch.
type IJqlArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JqlString() IJqlStringContext
	JqlNumber() IJqlNumberContext

	// IsJqlArgumentContext differentiates from other interfaces.
	IsJqlArgumentContext()
}

type JqlArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlArgumentContext() *JqlArgumentContext {
	var p = new(JqlArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlArgument
	return p
}

func (*JqlArgumentContext) IsJqlArgumentContext() {}

func NewJqlArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlArgumentContext {
	var p = new(JqlArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlArgument

	return p
}

func (s *JqlArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlArgumentContext) JqlString() IJqlStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlStringContext)
}

func (s *JqlArgumentContext) JqlNumber() IJqlNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlNumberContext)
}

func (s *JqlArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlArgument(s)
	}
}

func (s *JqlArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlArgument(s)
	}
}

func (p *JQLParser) JqlArgument() (localctx IJqlArgumentContext) {
	this := p
	_ = this

	localctx = NewJqlArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, JQLParserRULE_jqlArgument)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(290)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JQLParserSTRING, JQLParserQUOTE_STRING, JQLParserSQUOTE_STRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(288)
			p.JqlString()
		}

	case JQLParserPOSNUMBER, JQLParserNEGNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(289)
			p.JqlNumber()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IJqlWasPredicateContext is an interface to support dynamic dispatch.
type IJqlWasPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JqlPredicateOperand() IJqlPredicateOperandContext
	JqlDatePredicateOperator() IJqlDatePredicateOperatorContext
	JqlDateRangePredicateOperator() IJqlDateRangePredicateOperatorContext
	JqlUserPredicateOperator() IJqlUserPredicateOperatorContext

	// IsJqlWasPredicateContext differentiates from other interfaces.
	IsJqlWasPredicateContext()
}

type JqlWasPredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlWasPredicateContext() *JqlWasPredicateContext {
	var p = new(JqlWasPredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlWasPredicate
	return p
}

func (*JqlWasPredicateContext) IsJqlWasPredicateContext() {}

func NewJqlWasPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlWasPredicateContext {
	var p = new(JqlWasPredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlWasPredicate

	return p
}

func (s *JqlWasPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlWasPredicateContext) JqlPredicateOperand() IJqlPredicateOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlPredicateOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlPredicateOperandContext)
}

func (s *JqlWasPredicateContext) JqlDatePredicateOperator() IJqlDatePredicateOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlDatePredicateOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlDatePredicateOperatorContext)
}

func (s *JqlWasPredicateContext) JqlDateRangePredicateOperator() IJqlDateRangePredicateOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlDateRangePredicateOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlDateRangePredicateOperatorContext)
}

func (s *JqlWasPredicateContext) JqlUserPredicateOperator() IJqlUserPredicateOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlUserPredicateOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlUserPredicateOperatorContext)
}

func (s *JqlWasPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlWasPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlWasPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlWasPredicate(s)
	}
}

func (s *JqlWasPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlWasPredicate(s)
	}
}

func (p *JQLParser) JqlWasPredicate() (localctx IJqlWasPredicateContext) {
	this := p
	_ = this

	localctx = NewJqlWasPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, JQLParserRULE_jqlWasPredicate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(295)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JQLParserBEFORE, JQLParserAFTER, JQLParserON:
		{
			p.SetState(292)
			p.JqlDatePredicateOperator()
		}

	case JQLParserDURING:
		{
			p.SetState(293)
			p.JqlDateRangePredicateOperator()
		}

	case JQLParserBY:
		{
			p.SetState(294)
			p.JqlUserPredicateOperator()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(297)
		p.JqlPredicateOperand()
	}

	return localctx
}

// IJqlChangedPredicateContext is an interface to support dynamic dispatch.
type IJqlChangedPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JqlPredicateOperand() IJqlPredicateOperandContext
	JqlDatePredicateOperator() IJqlDatePredicateOperatorContext
	JqlDateRangePredicateOperator() IJqlDateRangePredicateOperatorContext
	JqlUserPredicateOperator() IJqlUserPredicateOperatorContext
	JqlValuePredicateOperator() IJqlValuePredicateOperatorContext

	// IsJqlChangedPredicateContext differentiates from other interfaces.
	IsJqlChangedPredicateContext()
}

type JqlChangedPredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlChangedPredicateContext() *JqlChangedPredicateContext {
	var p = new(JqlChangedPredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlChangedPredicate
	return p
}

func (*JqlChangedPredicateContext) IsJqlChangedPredicateContext() {}

func NewJqlChangedPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlChangedPredicateContext {
	var p = new(JqlChangedPredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlChangedPredicate

	return p
}

func (s *JqlChangedPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlChangedPredicateContext) JqlPredicateOperand() IJqlPredicateOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlPredicateOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlPredicateOperandContext)
}

func (s *JqlChangedPredicateContext) JqlDatePredicateOperator() IJqlDatePredicateOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlDatePredicateOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlDatePredicateOperatorContext)
}

func (s *JqlChangedPredicateContext) JqlDateRangePredicateOperator() IJqlDateRangePredicateOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlDateRangePredicateOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlDateRangePredicateOperatorContext)
}

func (s *JqlChangedPredicateContext) JqlUserPredicateOperator() IJqlUserPredicateOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlUserPredicateOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlUserPredicateOperatorContext)
}

func (s *JqlChangedPredicateContext) JqlValuePredicateOperator() IJqlValuePredicateOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlValuePredicateOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlValuePredicateOperatorContext)
}

func (s *JqlChangedPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlChangedPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlChangedPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlChangedPredicate(s)
	}
}

func (s *JqlChangedPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlChangedPredicate(s)
	}
}

func (p *JQLParser) JqlChangedPredicate() (localctx IJqlChangedPredicateContext) {
	this := p
	_ = this

	localctx = NewJqlChangedPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, JQLParserRULE_jqlChangedPredicate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(303)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JQLParserBEFORE, JQLParserAFTER, JQLParserON:
		{
			p.SetState(299)
			p.JqlDatePredicateOperator()
		}

	case JQLParserDURING:
		{
			p.SetState(300)
			p.JqlDateRangePredicateOperator()
		}

	case JQLParserBY:
		{
			p.SetState(301)
			p.JqlUserPredicateOperator()
		}

	case JQLParserFROM, JQLParserTO:
		{
			p.SetState(302)
			p.JqlValuePredicateOperator()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(305)
		p.JqlPredicateOperand()
	}

	return localctx
}

// IJqlDatePredicateOperatorContext is an interface to support dynamic dispatch.
type IJqlDatePredicateOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AFTER() antlr.TerminalNode
	BEFORE() antlr.TerminalNode
	ON() antlr.TerminalNode

	// IsJqlDatePredicateOperatorContext differentiates from other interfaces.
	IsJqlDatePredicateOperatorContext()
}

type JqlDatePredicateOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlDatePredicateOperatorContext() *JqlDatePredicateOperatorContext {
	var p = new(JqlDatePredicateOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlDatePredicateOperator
	return p
}

func (*JqlDatePredicateOperatorContext) IsJqlDatePredicateOperatorContext() {}

func NewJqlDatePredicateOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlDatePredicateOperatorContext {
	var p = new(JqlDatePredicateOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlDatePredicateOperator

	return p
}

func (s *JqlDatePredicateOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlDatePredicateOperatorContext) AFTER() antlr.TerminalNode {
	return s.GetToken(JQLParserAFTER, 0)
}

func (s *JqlDatePredicateOperatorContext) BEFORE() antlr.TerminalNode {
	return s.GetToken(JQLParserBEFORE, 0)
}

func (s *JqlDatePredicateOperatorContext) ON() antlr.TerminalNode {
	return s.GetToken(JQLParserON, 0)
}

func (s *JqlDatePredicateOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlDatePredicateOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlDatePredicateOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlDatePredicateOperator(s)
	}
}

func (s *JqlDatePredicateOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlDatePredicateOperator(s)
	}
}

func (p *JQLParser) JqlDatePredicateOperator() (localctx IJqlDatePredicateOperatorContext) {
	this := p
	_ = this

	localctx = NewJqlDatePredicateOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, JQLParserRULE_jqlDatePredicateOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&159383552) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IJqlDateRangePredicateOperatorContext is an interface to support dynamic dispatch.
type IJqlDateRangePredicateOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DURING() antlr.TerminalNode

	// IsJqlDateRangePredicateOperatorContext differentiates from other interfaces.
	IsJqlDateRangePredicateOperatorContext()
}

type JqlDateRangePredicateOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlDateRangePredicateOperatorContext() *JqlDateRangePredicateOperatorContext {
	var p = new(JqlDateRangePredicateOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlDateRangePredicateOperator
	return p
}

func (*JqlDateRangePredicateOperatorContext) IsJqlDateRangePredicateOperatorContext() {}

func NewJqlDateRangePredicateOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlDateRangePredicateOperatorContext {
	var p = new(JqlDateRangePredicateOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlDateRangePredicateOperator

	return p
}

func (s *JqlDateRangePredicateOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlDateRangePredicateOperatorContext) DURING() antlr.TerminalNode {
	return s.GetToken(JQLParserDURING, 0)
}

func (s *JqlDateRangePredicateOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlDateRangePredicateOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlDateRangePredicateOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlDateRangePredicateOperator(s)
	}
}

func (s *JqlDateRangePredicateOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlDateRangePredicateOperator(s)
	}
}

func (p *JQLParser) JqlDateRangePredicateOperator() (localctx IJqlDateRangePredicateOperatorContext) {
	this := p
	_ = this

	localctx = NewJqlDateRangePredicateOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, JQLParserRULE_jqlDateRangePredicateOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.Match(JQLParserDURING)
	}

	return localctx
}

// IJqlUserPredicateOperatorContext is an interface to support dynamic dispatch.
type IJqlUserPredicateOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BY() antlr.TerminalNode

	// IsJqlUserPredicateOperatorContext differentiates from other interfaces.
	IsJqlUserPredicateOperatorContext()
}

type JqlUserPredicateOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlUserPredicateOperatorContext() *JqlUserPredicateOperatorContext {
	var p = new(JqlUserPredicateOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlUserPredicateOperator
	return p
}

func (*JqlUserPredicateOperatorContext) IsJqlUserPredicateOperatorContext() {}

func NewJqlUserPredicateOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlUserPredicateOperatorContext {
	var p = new(JqlUserPredicateOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlUserPredicateOperator

	return p
}

func (s *JqlUserPredicateOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlUserPredicateOperatorContext) BY() antlr.TerminalNode {
	return s.GetToken(JQLParserBY, 0)
}

func (s *JqlUserPredicateOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlUserPredicateOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlUserPredicateOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlUserPredicateOperator(s)
	}
}

func (s *JqlUserPredicateOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlUserPredicateOperator(s)
	}
}

func (p *JQLParser) JqlUserPredicateOperator() (localctx IJqlUserPredicateOperatorContext) {
	this := p
	_ = this

	localctx = NewJqlUserPredicateOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, JQLParserRULE_jqlUserPredicateOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		p.Match(JQLParserBY)
	}

	return localctx
}

// IJqlValuePredicateOperatorContext is an interface to support dynamic dispatch.
type IJqlValuePredicateOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FROM() antlr.TerminalNode
	TO() antlr.TerminalNode

	// IsJqlValuePredicateOperatorContext differentiates from other interfaces.
	IsJqlValuePredicateOperatorContext()
}

type JqlValuePredicateOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlValuePredicateOperatorContext() *JqlValuePredicateOperatorContext {
	var p = new(JqlValuePredicateOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlValuePredicateOperator
	return p
}

func (*JqlValuePredicateOperatorContext) IsJqlValuePredicateOperatorContext() {}

func NewJqlValuePredicateOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlValuePredicateOperatorContext {
	var p = new(JqlValuePredicateOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlValuePredicateOperator

	return p
}

func (s *JqlValuePredicateOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlValuePredicateOperatorContext) FROM() antlr.TerminalNode {
	return s.GetToken(JQLParserFROM, 0)
}

func (s *JqlValuePredicateOperatorContext) TO() antlr.TerminalNode {
	return s.GetToken(JQLParserTO, 0)
}

func (s *JqlValuePredicateOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlValuePredicateOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlValuePredicateOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlValuePredicateOperator(s)
	}
}

func (s *JqlValuePredicateOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlValuePredicateOperator(s)
	}
}

func (p *JQLParser) JqlValuePredicateOperator() (localctx IJqlValuePredicateOperatorContext) {
	this := p
	_ = this

	localctx = NewJqlValuePredicateOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, JQLParserRULE_jqlValuePredicateOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		_la = p.GetTokenStream().LA(1)

		if !(_la == JQLParserFROM || _la == JQLParserTO) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IJqlPredicateOperandContext is an interface to support dynamic dispatch.
type IJqlPredicateOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JqlOperand() IJqlOperandContext

	// IsJqlPredicateOperandContext differentiates from other interfaces.
	IsJqlPredicateOperandContext()
}

type JqlPredicateOperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlPredicateOperandContext() *JqlPredicateOperandContext {
	var p = new(JqlPredicateOperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlPredicateOperand
	return p
}

func (*JqlPredicateOperandContext) IsJqlPredicateOperandContext() {}

func NewJqlPredicateOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlPredicateOperandContext {
	var p = new(JqlPredicateOperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlPredicateOperand

	return p
}

func (s *JqlPredicateOperandContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlPredicateOperandContext) JqlOperand() IJqlOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlOperandContext)
}

func (s *JqlPredicateOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlPredicateOperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlPredicateOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlPredicateOperand(s)
	}
}

func (s *JqlPredicateOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlPredicateOperand(s)
	}
}

func (p *JQLParser) JqlPredicateOperand() (localctx IJqlPredicateOperandContext) {
	this := p
	_ = this

	localctx = NewJqlPredicateOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, JQLParserRULE_jqlPredicateOperand)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(315)
		p.JqlOperand()
	}

	return localctx
}

// IJqlOrderByContext is an interface to support dynamic dispatch.
type IJqlOrderByContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ORDER() antlr.TerminalNode
	BY() antlr.TerminalNode
	AllJqlSearchSort() []IJqlSearchSortContext
	JqlSearchSort(i int) IJqlSearchSortContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsJqlOrderByContext differentiates from other interfaces.
	IsJqlOrderByContext()
}

type JqlOrderByContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlOrderByContext() *JqlOrderByContext {
	var p = new(JqlOrderByContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlOrderBy
	return p
}

func (*JqlOrderByContext) IsJqlOrderByContext() {}

func NewJqlOrderByContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlOrderByContext {
	var p = new(JqlOrderByContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlOrderBy

	return p
}

func (s *JqlOrderByContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlOrderByContext) ORDER() antlr.TerminalNode {
	return s.GetToken(JQLParserORDER, 0)
}

func (s *JqlOrderByContext) BY() antlr.TerminalNode {
	return s.GetToken(JQLParserBY, 0)
}

func (s *JqlOrderByContext) AllJqlSearchSort() []IJqlSearchSortContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJqlSearchSortContext); ok {
			len++
		}
	}

	tst := make([]IJqlSearchSortContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJqlSearchSortContext); ok {
			tst[i] = t.(IJqlSearchSortContext)
			i++
		}
	}

	return tst
}

func (s *JqlOrderByContext) JqlSearchSort(i int) IJqlSearchSortContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlSearchSortContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlSearchSortContext)
}

func (s *JqlOrderByContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JQLParserCOMMA)
}

func (s *JqlOrderByContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JQLParserCOMMA, i)
}

func (s *JqlOrderByContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlOrderByContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlOrderByContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlOrderBy(s)
	}
}

func (s *JqlOrderByContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlOrderBy(s)
	}
}

func (p *JQLParser) JqlOrderBy() (localctx IJqlOrderByContext) {
	this := p
	_ = this

	localctx = NewJqlOrderByContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, JQLParserRULE_jqlOrderBy)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(317)
		p.Match(JQLParserORDER)
	}
	{
		p.SetState(318)
		p.Match(JQLParserBY)
	}
	{
		p.SetState(319)
		p.JqlSearchSort()
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JQLParserCOMMA {
		{
			p.SetState(320)
			p.Match(JQLParserCOMMA)
		}
		{
			p.SetState(321)
			p.JqlSearchSort()
		}

		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IJqlSearchSortContext is an interface to support dynamic dispatch.
type IJqlSearchSortContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JqlField() IJqlFieldContext
	DESC() antlr.TerminalNode
	ASC() antlr.TerminalNode

	// IsJqlSearchSortContext differentiates from other interfaces.
	IsJqlSearchSortContext()
}

type JqlSearchSortContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJqlSearchSortContext() *JqlSearchSortContext {
	var p = new(JqlSearchSortContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JQLParserRULE_jqlSearchSort
	return p
}

func (*JqlSearchSortContext) IsJqlSearchSortContext() {}

func NewJqlSearchSortContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JqlSearchSortContext {
	var p = new(JqlSearchSortContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JQLParserRULE_jqlSearchSort

	return p
}

func (s *JqlSearchSortContext) GetParser() antlr.Parser { return s.parser }

func (s *JqlSearchSortContext) JqlField() IJqlFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJqlFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJqlFieldContext)
}

func (s *JqlSearchSortContext) DESC() antlr.TerminalNode {
	return s.GetToken(JQLParserDESC, 0)
}

func (s *JqlSearchSortContext) ASC() antlr.TerminalNode {
	return s.GetToken(JQLParserASC, 0)
}

func (s *JqlSearchSortContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JqlSearchSortContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JqlSearchSortContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.EnterJqlSearchSort(s)
	}
}

func (s *JqlSearchSortContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JQLParserListener); ok {
		listenerT.ExitJqlSearchSort(s)
	}
}

func (p *JQLParser) JqlSearchSort() (localctx IJqlSearchSortContext) {
	this := p
	_ = this

	localctx = NewJqlSearchSortContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, JQLParserRULE_jqlSearchSort)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)
		p.JqlField()
	}
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JQLParserASC || _la == JQLParserDESC {
		{
			p.SetState(328)
			_la = p.GetTokenStream().LA(1)

			if !(_la == JQLParserASC || _la == JQLParserDESC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}
