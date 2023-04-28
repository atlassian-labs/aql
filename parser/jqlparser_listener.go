// Code generated from JQLParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // JQLParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// JQLParserListener is a complete listener for a parse tree produced by JQLParser.
type JQLParserListener interface {
	antlr.ParseTreeListener

	// EnterJqlQuery is called when entering the jqlQuery production.
	EnterJqlQuery(c *JqlQueryContext)

	// EnterJqlWhere is called when entering the jqlWhere production.
	EnterJqlWhere(c *JqlWhereContext)

	// EnterJqlOrClause is called when entering the jqlOrClause production.
	EnterJqlOrClause(c *JqlOrClauseContext)

	// EnterJqlAndClause is called when entering the jqlAndClause production.
	EnterJqlAndClause(c *JqlAndClauseContext)

	// EnterJqlNotClause is called when entering the jqlNotClause production.
	EnterJqlNotClause(c *JqlNotClauseContext)

	// EnterJqlSubClause is called when entering the jqlSubClause production.
	EnterJqlSubClause(c *JqlSubClauseContext)

	// EnterJqlTerminalClause is called when entering the jqlTerminalClause production.
	EnterJqlTerminalClause(c *JqlTerminalClauseContext)

	// EnterJqlEqualsClause is called when entering the jqlEqualsClause production.
	EnterJqlEqualsClause(c *JqlEqualsClauseContext)

	// EnterJqlLikeClause is called when entering the jqlLikeClause production.
	EnterJqlLikeClause(c *JqlLikeClauseContext)

	// EnterJqlComparisonClause is called when entering the jqlComparisonClause production.
	EnterJqlComparisonClause(c *JqlComparisonClauseContext)

	// EnterJqlInClause is called when entering the jqlInClause production.
	EnterJqlInClause(c *JqlInClauseContext)

	// EnterJqlIsClause is called when entering the jqlIsClause production.
	EnterJqlIsClause(c *JqlIsClauseContext)

	// EnterJqlWasClause is called when entering the jqlWasClause production.
	EnterJqlWasClause(c *JqlWasClauseContext)

	// EnterJqlWasInClause is called when entering the jqlWasInClause production.
	EnterJqlWasInClause(c *JqlWasInClauseContext)

	// EnterJqlChangedClause is called when entering the jqlChangedClause production.
	EnterJqlChangedClause(c *JqlChangedClauseContext)

	// EnterJqlEqualsOperator is called when entering the jqlEqualsOperator production.
	EnterJqlEqualsOperator(c *JqlEqualsOperatorContext)

	// EnterJqlLikeOperator is called when entering the jqlLikeOperator production.
	EnterJqlLikeOperator(c *JqlLikeOperatorContext)

	// EnterJqlComparisonOperator is called when entering the jqlComparisonOperator production.
	EnterJqlComparisonOperator(c *JqlComparisonOperatorContext)

	// EnterJqlInOperator is called when entering the jqlInOperator production.
	EnterJqlInOperator(c *JqlInOperatorContext)

	// EnterJqlIsOperator is called when entering the jqlIsOperator production.
	EnterJqlIsOperator(c *JqlIsOperatorContext)

	// EnterJqlWasOperator is called when entering the jqlWasOperator production.
	EnterJqlWasOperator(c *JqlWasOperatorContext)

	// EnterJqlWasInOperator is called when entering the jqlWasInOperator production.
	EnterJqlWasInOperator(c *JqlWasInOperatorContext)

	// EnterJqlChangedOperator is called when entering the jqlChangedOperator production.
	EnterJqlChangedOperator(c *JqlChangedOperatorContext)

	// EnterJqlNumberField is called when entering the jqlNumberField production.
	EnterJqlNumberField(c *JqlNumberFieldContext)

	// EnterJqlNonNumberField is called when entering the jqlNonNumberField production.
	EnterJqlNonNumberField(c *JqlNonNumberFieldContext)

	// EnterJqlFieldProperty is called when entering the jqlFieldProperty production.
	EnterJqlFieldProperty(c *JqlFieldPropertyContext)

	// EnterJqlCustomField is called when entering the jqlCustomField production.
	EnterJqlCustomField(c *JqlCustomFieldContext)

	// EnterJqlString is called when entering the jqlString production.
	EnterJqlString(c *JqlStringContext)

	// EnterJqlNumber is called when entering the jqlNumber production.
	EnterJqlNumber(c *JqlNumberContext)

	// EnterJqlOperand is called when entering the jqlOperand production.
	EnterJqlOperand(c *JqlOperandContext)

	// EnterJqlEmpty is called when entering the jqlEmpty production.
	EnterJqlEmpty(c *JqlEmptyContext)

	// EnterJqlValue is called when entering the jqlValue production.
	EnterJqlValue(c *JqlValueContext)

	// EnterJqlFunction is called when entering the jqlFunction production.
	EnterJqlFunction(c *JqlFunctionContext)

	// EnterJqlFunctionName is called when entering the jqlFunctionName production.
	EnterJqlFunctionName(c *JqlFunctionNameContext)

	// EnterJqlArgumentList is called when entering the jqlArgumentList production.
	EnterJqlArgumentList(c *JqlArgumentListContext)

	// EnterJqlList is called when entering the jqlList production.
	EnterJqlList(c *JqlListContext)

	// EnterJqlListStart is called when entering the jqlListStart production.
	EnterJqlListStart(c *JqlListStartContext)

	// EnterJqlListEnd is called when entering the jqlListEnd production.
	EnterJqlListEnd(c *JqlListEndContext)

	// EnterJqlPropertyArgument is called when entering the jqlPropertyArgument production.
	EnterJqlPropertyArgument(c *JqlPropertyArgumentContext)

	// EnterJqlArgument is called when entering the jqlArgument production.
	EnterJqlArgument(c *JqlArgumentContext)

	// EnterJqlWasPredicate is called when entering the jqlWasPredicate production.
	EnterJqlWasPredicate(c *JqlWasPredicateContext)

	// EnterJqlChangedPredicate is called when entering the jqlChangedPredicate production.
	EnterJqlChangedPredicate(c *JqlChangedPredicateContext)

	// EnterJqlDatePredicateOperator is called when entering the jqlDatePredicateOperator production.
	EnterJqlDatePredicateOperator(c *JqlDatePredicateOperatorContext)

	// EnterJqlDateRangePredicateOperator is called when entering the jqlDateRangePredicateOperator production.
	EnterJqlDateRangePredicateOperator(c *JqlDateRangePredicateOperatorContext)

	// EnterJqlUserPredicateOperator is called when entering the jqlUserPredicateOperator production.
	EnterJqlUserPredicateOperator(c *JqlUserPredicateOperatorContext)

	// EnterJqlValuePredicateOperator is called when entering the jqlValuePredicateOperator production.
	EnterJqlValuePredicateOperator(c *JqlValuePredicateOperatorContext)

	// EnterJqlPredicateOperand is called when entering the jqlPredicateOperand production.
	EnterJqlPredicateOperand(c *JqlPredicateOperandContext)

	// EnterJqlOrderBy is called when entering the jqlOrderBy production.
	EnterJqlOrderBy(c *JqlOrderByContext)

	// EnterJqlSearchSort is called when entering the jqlSearchSort production.
	EnterJqlSearchSort(c *JqlSearchSortContext)

	// ExitJqlQuery is called when exiting the jqlQuery production.
	ExitJqlQuery(c *JqlQueryContext)

	// ExitJqlWhere is called when exiting the jqlWhere production.
	ExitJqlWhere(c *JqlWhereContext)

	// ExitJqlOrClause is called when exiting the jqlOrClause production.
	ExitJqlOrClause(c *JqlOrClauseContext)

	// ExitJqlAndClause is called when exiting the jqlAndClause production.
	ExitJqlAndClause(c *JqlAndClauseContext)

	// ExitJqlNotClause is called when exiting the jqlNotClause production.
	ExitJqlNotClause(c *JqlNotClauseContext)

	// ExitJqlSubClause is called when exiting the jqlSubClause production.
	ExitJqlSubClause(c *JqlSubClauseContext)

	// ExitJqlTerminalClause is called when exiting the jqlTerminalClause production.
	ExitJqlTerminalClause(c *JqlTerminalClauseContext)

	// ExitJqlEqualsClause is called when exiting the jqlEqualsClause production.
	ExitJqlEqualsClause(c *JqlEqualsClauseContext)

	// ExitJqlLikeClause is called when exiting the jqlLikeClause production.
	ExitJqlLikeClause(c *JqlLikeClauseContext)

	// ExitJqlComparisonClause is called when exiting the jqlComparisonClause production.
	ExitJqlComparisonClause(c *JqlComparisonClauseContext)

	// ExitJqlInClause is called when exiting the jqlInClause production.
	ExitJqlInClause(c *JqlInClauseContext)

	// ExitJqlIsClause is called when exiting the jqlIsClause production.
	ExitJqlIsClause(c *JqlIsClauseContext)

	// ExitJqlWasClause is called when exiting the jqlWasClause production.
	ExitJqlWasClause(c *JqlWasClauseContext)

	// ExitJqlWasInClause is called when exiting the jqlWasInClause production.
	ExitJqlWasInClause(c *JqlWasInClauseContext)

	// ExitJqlChangedClause is called when exiting the jqlChangedClause production.
	ExitJqlChangedClause(c *JqlChangedClauseContext)

	// ExitJqlEqualsOperator is called when exiting the jqlEqualsOperator production.
	ExitJqlEqualsOperator(c *JqlEqualsOperatorContext)

	// ExitJqlLikeOperator is called when exiting the jqlLikeOperator production.
	ExitJqlLikeOperator(c *JqlLikeOperatorContext)

	// ExitJqlComparisonOperator is called when exiting the jqlComparisonOperator production.
	ExitJqlComparisonOperator(c *JqlComparisonOperatorContext)

	// ExitJqlInOperator is called when exiting the jqlInOperator production.
	ExitJqlInOperator(c *JqlInOperatorContext)

	// ExitJqlIsOperator is called when exiting the jqlIsOperator production.
	ExitJqlIsOperator(c *JqlIsOperatorContext)

	// ExitJqlWasOperator is called when exiting the jqlWasOperator production.
	ExitJqlWasOperator(c *JqlWasOperatorContext)

	// ExitJqlWasInOperator is called when exiting the jqlWasInOperator production.
	ExitJqlWasInOperator(c *JqlWasInOperatorContext)

	// ExitJqlChangedOperator is called when exiting the jqlChangedOperator production.
	ExitJqlChangedOperator(c *JqlChangedOperatorContext)

	// ExitJqlNumberField is called when exiting the jqlNumberField production.
	ExitJqlNumberField(c *JqlNumberFieldContext)

	// ExitJqlNonNumberField is called when exiting the jqlNonNumberField production.
	ExitJqlNonNumberField(c *JqlNonNumberFieldContext)

	// ExitJqlFieldProperty is called when exiting the jqlFieldProperty production.
	ExitJqlFieldProperty(c *JqlFieldPropertyContext)

	// ExitJqlCustomField is called when exiting the jqlCustomField production.
	ExitJqlCustomField(c *JqlCustomFieldContext)

	// ExitJqlString is called when exiting the jqlString production.
	ExitJqlString(c *JqlStringContext)

	// ExitJqlNumber is called when exiting the jqlNumber production.
	ExitJqlNumber(c *JqlNumberContext)

	// ExitJqlOperand is called when exiting the jqlOperand production.
	ExitJqlOperand(c *JqlOperandContext)

	// ExitJqlEmpty is called when exiting the jqlEmpty production.
	ExitJqlEmpty(c *JqlEmptyContext)

	// ExitJqlValue is called when exiting the jqlValue production.
	ExitJqlValue(c *JqlValueContext)

	// ExitJqlFunction is called when exiting the jqlFunction production.
	ExitJqlFunction(c *JqlFunctionContext)

	// ExitJqlFunctionName is called when exiting the jqlFunctionName production.
	ExitJqlFunctionName(c *JqlFunctionNameContext)

	// ExitJqlArgumentList is called when exiting the jqlArgumentList production.
	ExitJqlArgumentList(c *JqlArgumentListContext)

	// ExitJqlList is called when exiting the jqlList production.
	ExitJqlList(c *JqlListContext)

	// ExitJqlListStart is called when exiting the jqlListStart production.
	ExitJqlListStart(c *JqlListStartContext)

	// ExitJqlListEnd is called when exiting the jqlListEnd production.
	ExitJqlListEnd(c *JqlListEndContext)

	// ExitJqlPropertyArgument is called when exiting the jqlPropertyArgument production.
	ExitJqlPropertyArgument(c *JqlPropertyArgumentContext)

	// ExitJqlArgument is called when exiting the jqlArgument production.
	ExitJqlArgument(c *JqlArgumentContext)

	// ExitJqlWasPredicate is called when exiting the jqlWasPredicate production.
	ExitJqlWasPredicate(c *JqlWasPredicateContext)

	// ExitJqlChangedPredicate is called when exiting the jqlChangedPredicate production.
	ExitJqlChangedPredicate(c *JqlChangedPredicateContext)

	// ExitJqlDatePredicateOperator is called when exiting the jqlDatePredicateOperator production.
	ExitJqlDatePredicateOperator(c *JqlDatePredicateOperatorContext)

	// ExitJqlDateRangePredicateOperator is called when exiting the jqlDateRangePredicateOperator production.
	ExitJqlDateRangePredicateOperator(c *JqlDateRangePredicateOperatorContext)

	// ExitJqlUserPredicateOperator is called when exiting the jqlUserPredicateOperator production.
	ExitJqlUserPredicateOperator(c *JqlUserPredicateOperatorContext)

	// ExitJqlValuePredicateOperator is called when exiting the jqlValuePredicateOperator production.
	ExitJqlValuePredicateOperator(c *JqlValuePredicateOperatorContext)

	// ExitJqlPredicateOperand is called when exiting the jqlPredicateOperand production.
	ExitJqlPredicateOperand(c *JqlPredicateOperandContext)

	// ExitJqlOrderBy is called when exiting the jqlOrderBy production.
	ExitJqlOrderBy(c *JqlOrderByContext)

	// ExitJqlSearchSort is called when exiting the jqlSearchSort production.
	ExitJqlSearchSort(c *JqlSearchSortContext)
}
