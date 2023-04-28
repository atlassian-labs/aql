// Code generated from JQLParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // JQLParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseJQLParserListener is a complete listener for a parse tree produced by JQLParser.
type BaseJQLParserListener struct{}

var _ JQLParserListener = &BaseJQLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJQLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJQLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJQLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJQLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterJqlQuery is called when production jqlQuery is entered.
func (s *BaseJQLParserListener) EnterJqlQuery(ctx *JqlQueryContext) {}

// ExitJqlQuery is called when production jqlQuery is exited.
func (s *BaseJQLParserListener) ExitJqlQuery(ctx *JqlQueryContext) {}

// EnterJqlWhere is called when production jqlWhere is entered.
func (s *BaseJQLParserListener) EnterJqlWhere(ctx *JqlWhereContext) {}

// ExitJqlWhere is called when production jqlWhere is exited.
func (s *BaseJQLParserListener) ExitJqlWhere(ctx *JqlWhereContext) {}

// EnterJqlOrClause is called when production jqlOrClause is entered.
func (s *BaseJQLParserListener) EnterJqlOrClause(ctx *JqlOrClauseContext) {}

// ExitJqlOrClause is called when production jqlOrClause is exited.
func (s *BaseJQLParserListener) ExitJqlOrClause(ctx *JqlOrClauseContext) {}

// EnterJqlAndClause is called when production jqlAndClause is entered.
func (s *BaseJQLParserListener) EnterJqlAndClause(ctx *JqlAndClauseContext) {}

// ExitJqlAndClause is called when production jqlAndClause is exited.
func (s *BaseJQLParserListener) ExitJqlAndClause(ctx *JqlAndClauseContext) {}

// EnterJqlNotClause is called when production jqlNotClause is entered.
func (s *BaseJQLParserListener) EnterJqlNotClause(ctx *JqlNotClauseContext) {}

// ExitJqlNotClause is called when production jqlNotClause is exited.
func (s *BaseJQLParserListener) ExitJqlNotClause(ctx *JqlNotClauseContext) {}

// EnterJqlSubClause is called when production jqlSubClause is entered.
func (s *BaseJQLParserListener) EnterJqlSubClause(ctx *JqlSubClauseContext) {}

// ExitJqlSubClause is called when production jqlSubClause is exited.
func (s *BaseJQLParserListener) ExitJqlSubClause(ctx *JqlSubClauseContext) {}

// EnterJqlTerminalClause is called when production jqlTerminalClause is entered.
func (s *BaseJQLParserListener) EnterJqlTerminalClause(ctx *JqlTerminalClauseContext) {}

// ExitJqlTerminalClause is called when production jqlTerminalClause is exited.
func (s *BaseJQLParserListener) ExitJqlTerminalClause(ctx *JqlTerminalClauseContext) {}

// EnterJqlEqualsClause is called when production jqlEqualsClause is entered.
func (s *BaseJQLParserListener) EnterJqlEqualsClause(ctx *JqlEqualsClauseContext) {}

// ExitJqlEqualsClause is called when production jqlEqualsClause is exited.
func (s *BaseJQLParserListener) ExitJqlEqualsClause(ctx *JqlEqualsClauseContext) {}

// EnterJqlLikeClause is called when production jqlLikeClause is entered.
func (s *BaseJQLParserListener) EnterJqlLikeClause(ctx *JqlLikeClauseContext) {}

// ExitJqlLikeClause is called when production jqlLikeClause is exited.
func (s *BaseJQLParserListener) ExitJqlLikeClause(ctx *JqlLikeClauseContext) {}

// EnterJqlComparisonClause is called when production jqlComparisonClause is entered.
func (s *BaseJQLParserListener) EnterJqlComparisonClause(ctx *JqlComparisonClauseContext) {}

// ExitJqlComparisonClause is called when production jqlComparisonClause is exited.
func (s *BaseJQLParserListener) ExitJqlComparisonClause(ctx *JqlComparisonClauseContext) {}

// EnterJqlInClause is called when production jqlInClause is entered.
func (s *BaseJQLParserListener) EnterJqlInClause(ctx *JqlInClauseContext) {}

// ExitJqlInClause is called when production jqlInClause is exited.
func (s *BaseJQLParserListener) ExitJqlInClause(ctx *JqlInClauseContext) {}

// EnterJqlIsClause is called when production jqlIsClause is entered.
func (s *BaseJQLParserListener) EnterJqlIsClause(ctx *JqlIsClauseContext) {}

// ExitJqlIsClause is called when production jqlIsClause is exited.
func (s *BaseJQLParserListener) ExitJqlIsClause(ctx *JqlIsClauseContext) {}

// EnterJqlWasClause is called when production jqlWasClause is entered.
func (s *BaseJQLParserListener) EnterJqlWasClause(ctx *JqlWasClauseContext) {}

// ExitJqlWasClause is called when production jqlWasClause is exited.
func (s *BaseJQLParserListener) ExitJqlWasClause(ctx *JqlWasClauseContext) {}

// EnterJqlWasInClause is called when production jqlWasInClause is entered.
func (s *BaseJQLParserListener) EnterJqlWasInClause(ctx *JqlWasInClauseContext) {}

// ExitJqlWasInClause is called when production jqlWasInClause is exited.
func (s *BaseJQLParserListener) ExitJqlWasInClause(ctx *JqlWasInClauseContext) {}

// EnterJqlChangedClause is called when production jqlChangedClause is entered.
func (s *BaseJQLParserListener) EnterJqlChangedClause(ctx *JqlChangedClauseContext) {}

// ExitJqlChangedClause is called when production jqlChangedClause is exited.
func (s *BaseJQLParserListener) ExitJqlChangedClause(ctx *JqlChangedClauseContext) {}

// EnterJqlEqualsOperator is called when production jqlEqualsOperator is entered.
func (s *BaseJQLParserListener) EnterJqlEqualsOperator(ctx *JqlEqualsOperatorContext) {}

// ExitJqlEqualsOperator is called when production jqlEqualsOperator is exited.
func (s *BaseJQLParserListener) ExitJqlEqualsOperator(ctx *JqlEqualsOperatorContext) {}

// EnterJqlLikeOperator is called when production jqlLikeOperator is entered.
func (s *BaseJQLParserListener) EnterJqlLikeOperator(ctx *JqlLikeOperatorContext) {}

// ExitJqlLikeOperator is called when production jqlLikeOperator is exited.
func (s *BaseJQLParserListener) ExitJqlLikeOperator(ctx *JqlLikeOperatorContext) {}

// EnterJqlComparisonOperator is called when production jqlComparisonOperator is entered.
func (s *BaseJQLParserListener) EnterJqlComparisonOperator(ctx *JqlComparisonOperatorContext) {}

// ExitJqlComparisonOperator is called when production jqlComparisonOperator is exited.
func (s *BaseJQLParserListener) ExitJqlComparisonOperator(ctx *JqlComparisonOperatorContext) {}

// EnterJqlInOperator is called when production jqlInOperator is entered.
func (s *BaseJQLParserListener) EnterJqlInOperator(ctx *JqlInOperatorContext) {}

// ExitJqlInOperator is called when production jqlInOperator is exited.
func (s *BaseJQLParserListener) ExitJqlInOperator(ctx *JqlInOperatorContext) {}

// EnterJqlIsOperator is called when production jqlIsOperator is entered.
func (s *BaseJQLParserListener) EnterJqlIsOperator(ctx *JqlIsOperatorContext) {}

// ExitJqlIsOperator is called when production jqlIsOperator is exited.
func (s *BaseJQLParserListener) ExitJqlIsOperator(ctx *JqlIsOperatorContext) {}

// EnterJqlWasOperator is called when production jqlWasOperator is entered.
func (s *BaseJQLParserListener) EnterJqlWasOperator(ctx *JqlWasOperatorContext) {}

// ExitJqlWasOperator is called when production jqlWasOperator is exited.
func (s *BaseJQLParserListener) ExitJqlWasOperator(ctx *JqlWasOperatorContext) {}

// EnterJqlWasInOperator is called when production jqlWasInOperator is entered.
func (s *BaseJQLParserListener) EnterJqlWasInOperator(ctx *JqlWasInOperatorContext) {}

// ExitJqlWasInOperator is called when production jqlWasInOperator is exited.
func (s *BaseJQLParserListener) ExitJqlWasInOperator(ctx *JqlWasInOperatorContext) {}

// EnterJqlChangedOperator is called when production jqlChangedOperator is entered.
func (s *BaseJQLParserListener) EnterJqlChangedOperator(ctx *JqlChangedOperatorContext) {}

// ExitJqlChangedOperator is called when production jqlChangedOperator is exited.
func (s *BaseJQLParserListener) ExitJqlChangedOperator(ctx *JqlChangedOperatorContext) {}

// EnterJqlNumberField is called when production jqlNumberField is entered.
func (s *BaseJQLParserListener) EnterJqlNumberField(ctx *JqlNumberFieldContext) {}

// ExitJqlNumberField is called when production jqlNumberField is exited.
func (s *BaseJQLParserListener) ExitJqlNumberField(ctx *JqlNumberFieldContext) {}

// EnterJqlNonNumberField is called when production jqlNonNumberField is entered.
func (s *BaseJQLParserListener) EnterJqlNonNumberField(ctx *JqlNonNumberFieldContext) {}

// ExitJqlNonNumberField is called when production jqlNonNumberField is exited.
func (s *BaseJQLParserListener) ExitJqlNonNumberField(ctx *JqlNonNumberFieldContext) {}

// EnterJqlFieldProperty is called when production jqlFieldProperty is entered.
func (s *BaseJQLParserListener) EnterJqlFieldProperty(ctx *JqlFieldPropertyContext) {}

// ExitJqlFieldProperty is called when production jqlFieldProperty is exited.
func (s *BaseJQLParserListener) ExitJqlFieldProperty(ctx *JqlFieldPropertyContext) {}

// EnterJqlCustomField is called when production jqlCustomField is entered.
func (s *BaseJQLParserListener) EnterJqlCustomField(ctx *JqlCustomFieldContext) {}

// ExitJqlCustomField is called when production jqlCustomField is exited.
func (s *BaseJQLParserListener) ExitJqlCustomField(ctx *JqlCustomFieldContext) {}

// EnterJqlString is called when production jqlString is entered.
func (s *BaseJQLParserListener) EnterJqlString(ctx *JqlStringContext) {}

// ExitJqlString is called when production jqlString is exited.
func (s *BaseJQLParserListener) ExitJqlString(ctx *JqlStringContext) {}

// EnterJqlNumber is called when production jqlNumber is entered.
func (s *BaseJQLParserListener) EnterJqlNumber(ctx *JqlNumberContext) {}

// ExitJqlNumber is called when production jqlNumber is exited.
func (s *BaseJQLParserListener) ExitJqlNumber(ctx *JqlNumberContext) {}

// EnterJqlOperand is called when production jqlOperand is entered.
func (s *BaseJQLParserListener) EnterJqlOperand(ctx *JqlOperandContext) {}

// ExitJqlOperand is called when production jqlOperand is exited.
func (s *BaseJQLParserListener) ExitJqlOperand(ctx *JqlOperandContext) {}

// EnterJqlEmpty is called when production jqlEmpty is entered.
func (s *BaseJQLParserListener) EnterJqlEmpty(ctx *JqlEmptyContext) {}

// ExitJqlEmpty is called when production jqlEmpty is exited.
func (s *BaseJQLParserListener) ExitJqlEmpty(ctx *JqlEmptyContext) {}

// EnterJqlValue is called when production jqlValue is entered.
func (s *BaseJQLParserListener) EnterJqlValue(ctx *JqlValueContext) {}

// ExitJqlValue is called when production jqlValue is exited.
func (s *BaseJQLParserListener) ExitJqlValue(ctx *JqlValueContext) {}

// EnterJqlFunction is called when production jqlFunction is entered.
func (s *BaseJQLParserListener) EnterJqlFunction(ctx *JqlFunctionContext) {}

// ExitJqlFunction is called when production jqlFunction is exited.
func (s *BaseJQLParserListener) ExitJqlFunction(ctx *JqlFunctionContext) {}

// EnterJqlFunctionName is called when production jqlFunctionName is entered.
func (s *BaseJQLParserListener) EnterJqlFunctionName(ctx *JqlFunctionNameContext) {}

// ExitJqlFunctionName is called when production jqlFunctionName is exited.
func (s *BaseJQLParserListener) ExitJqlFunctionName(ctx *JqlFunctionNameContext) {}

// EnterJqlArgumentList is called when production jqlArgumentList is entered.
func (s *BaseJQLParserListener) EnterJqlArgumentList(ctx *JqlArgumentListContext) {}

// ExitJqlArgumentList is called when production jqlArgumentList is exited.
func (s *BaseJQLParserListener) ExitJqlArgumentList(ctx *JqlArgumentListContext) {}

// EnterJqlList is called when production jqlList is entered.
func (s *BaseJQLParserListener) EnterJqlList(ctx *JqlListContext) {}

// ExitJqlList is called when production jqlList is exited.
func (s *BaseJQLParserListener) ExitJqlList(ctx *JqlListContext) {}

// EnterJqlListStart is called when production jqlListStart is entered.
func (s *BaseJQLParserListener) EnterJqlListStart(ctx *JqlListStartContext) {}

// ExitJqlListStart is called when production jqlListStart is exited.
func (s *BaseJQLParserListener) ExitJqlListStart(ctx *JqlListStartContext) {}

// EnterJqlListEnd is called when production jqlListEnd is entered.
func (s *BaseJQLParserListener) EnterJqlListEnd(ctx *JqlListEndContext) {}

// ExitJqlListEnd is called when production jqlListEnd is exited.
func (s *BaseJQLParserListener) ExitJqlListEnd(ctx *JqlListEndContext) {}

// EnterJqlPropertyArgument is called when production jqlPropertyArgument is entered.
func (s *BaseJQLParserListener) EnterJqlPropertyArgument(ctx *JqlPropertyArgumentContext) {}

// ExitJqlPropertyArgument is called when production jqlPropertyArgument is exited.
func (s *BaseJQLParserListener) ExitJqlPropertyArgument(ctx *JqlPropertyArgumentContext) {}

// EnterJqlArgument is called when production jqlArgument is entered.
func (s *BaseJQLParserListener) EnterJqlArgument(ctx *JqlArgumentContext) {}

// ExitJqlArgument is called when production jqlArgument is exited.
func (s *BaseJQLParserListener) ExitJqlArgument(ctx *JqlArgumentContext) {}

// EnterJqlWasPredicate is called when production jqlWasPredicate is entered.
func (s *BaseJQLParserListener) EnterJqlWasPredicate(ctx *JqlWasPredicateContext) {}

// ExitJqlWasPredicate is called when production jqlWasPredicate is exited.
func (s *BaseJQLParserListener) ExitJqlWasPredicate(ctx *JqlWasPredicateContext) {}

// EnterJqlChangedPredicate is called when production jqlChangedPredicate is entered.
func (s *BaseJQLParserListener) EnterJqlChangedPredicate(ctx *JqlChangedPredicateContext) {}

// ExitJqlChangedPredicate is called when production jqlChangedPredicate is exited.
func (s *BaseJQLParserListener) ExitJqlChangedPredicate(ctx *JqlChangedPredicateContext) {}

// EnterJqlDatePredicateOperator is called when production jqlDatePredicateOperator is entered.
func (s *BaseJQLParserListener) EnterJqlDatePredicateOperator(ctx *JqlDatePredicateOperatorContext) {}

// ExitJqlDatePredicateOperator is called when production jqlDatePredicateOperator is exited.
func (s *BaseJQLParserListener) ExitJqlDatePredicateOperator(ctx *JqlDatePredicateOperatorContext) {}

// EnterJqlDateRangePredicateOperator is called when production jqlDateRangePredicateOperator is entered.
func (s *BaseJQLParserListener) EnterJqlDateRangePredicateOperator(ctx *JqlDateRangePredicateOperatorContext) {
}

// ExitJqlDateRangePredicateOperator is called when production jqlDateRangePredicateOperator is exited.
func (s *BaseJQLParserListener) ExitJqlDateRangePredicateOperator(ctx *JqlDateRangePredicateOperatorContext) {
}

// EnterJqlUserPredicateOperator is called when production jqlUserPredicateOperator is entered.
func (s *BaseJQLParserListener) EnterJqlUserPredicateOperator(ctx *JqlUserPredicateOperatorContext) {}

// ExitJqlUserPredicateOperator is called when production jqlUserPredicateOperator is exited.
func (s *BaseJQLParserListener) ExitJqlUserPredicateOperator(ctx *JqlUserPredicateOperatorContext) {}

// EnterJqlValuePredicateOperator is called when production jqlValuePredicateOperator is entered.
func (s *BaseJQLParserListener) EnterJqlValuePredicateOperator(ctx *JqlValuePredicateOperatorContext) {
}

// ExitJqlValuePredicateOperator is called when production jqlValuePredicateOperator is exited.
func (s *BaseJQLParserListener) ExitJqlValuePredicateOperator(ctx *JqlValuePredicateOperatorContext) {
}

// EnterJqlPredicateOperand is called when production jqlPredicateOperand is entered.
func (s *BaseJQLParserListener) EnterJqlPredicateOperand(ctx *JqlPredicateOperandContext) {}

// ExitJqlPredicateOperand is called when production jqlPredicateOperand is exited.
func (s *BaseJQLParserListener) ExitJqlPredicateOperand(ctx *JqlPredicateOperandContext) {}

// EnterJqlOrderBy is called when production jqlOrderBy is entered.
func (s *BaseJQLParserListener) EnterJqlOrderBy(ctx *JqlOrderByContext) {}

// ExitJqlOrderBy is called when production jqlOrderBy is exited.
func (s *BaseJQLParserListener) ExitJqlOrderBy(ctx *JqlOrderByContext) {}

// EnterJqlSearchSort is called when production jqlSearchSort is entered.
func (s *BaseJQLParserListener) EnterJqlSearchSort(ctx *JqlSearchSortContext) {}

// ExitJqlSearchSort is called when production jqlSearchSort is exited.
func (s *BaseJQLParserListener) ExitJqlSearchSort(ctx *JqlSearchSortContext) {}
