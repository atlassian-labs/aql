parser grammar JQLParser;
options { tokenVocab=JQLLexer; }

// Rules are prefixed with jql to prevent potential conflicts with reserved words when generating code for target languages

jqlQuery
    : jqlWhere? jqlOrderBy? EOF
    ;

jqlWhere
    : jqlOrClause
    ;

jqlOrClause
    : jqlAndClause (OR jqlAndClause)*
    ;

jqlAndClause
    : jqlNotClause (AND jqlNotClause)*
    ;

jqlNotClause
    : (NOT | BANG) jqlNotClause
    | jqlSubClause
    | jqlTerminalClause
    ;

jqlSubClause
    : LPAREN jqlOrClause RPAREN
    ;

jqlTerminalClause
    : jqlField jqlTerminalClauseRhs
    ;

jqlTerminalClauseRhs
    : jqlEqualsOperator (jqlEmpty | jqlValue | jqlFunction) # jqlEqualsClause
    | jqlLikeOperator (jqlEmpty | jqlValue | jqlFunction) # jqlLikeClause
    | jqlComparisonOperator (jqlValue | jqlFunction) # jqlComparisonClause
    | jqlInOperator (jqlList | jqlFunction) # jqlInClause
    | jqlIsOperator jqlEmpty # jqlIsClause
    | jqlWasOperator (jqlEmpty | jqlValue | jqlFunction) jqlWasPredicate* # jqlWasClause
    | jqlWasInOperator (jqlList | jqlFunction) jqlWasPredicate* # jqlWasInClause
    | jqlChangedOperator jqlChangedPredicate* # jqlChangedClause
    ;

jqlEqualsOperator
    : EQUALS | NOT_EQUALS
    ;

jqlLikeOperator
    : LIKE | NOT_LIKE
    ;

jqlComparisonOperator
    : LT | GT | LTEQ | GTEQ
    ;

jqlInOperator
    : IN | NOT IN
    ;

jqlIsOperator
    : IS | IS NOT
    ;

jqlWasOperator
    : WAS | WAS NOT
    ;

jqlWasInOperator
    : WAS IN | WAS NOT IN
    ;

jqlChangedOperator
    : CHANGED
    ;

jqlField
    : jqlNumber #jqlNumberField
    | (
        (jqlString | jqlCustomField)
        (jqlFieldProperty)*
    ) #jqlNonNumberField
    ;

jqlFieldProperty
    : (LBRACKET (jqlArgument) RBRACKET) (jqlPropertyArgument)*;

jqlCustomField
    : CUSTOMFIELD LBRACKET POSNUMBER RBRACKET
    ;

jqlString
    : STRING
    | QUOTE_STRING
    | SQUOTE_STRING
    ;

jqlNumber
	: jqlNum = (POSNUMBER | NEGNUMBER)
	;

jqlOperand
    : jqlEmpty
    | jqlValue
    | jqlFunction
    | jqlList
    ;

jqlEmpty
    : EMPTY // Empty is re-defined as a rule to simplify AST generation
    ;

jqlValue
    : jqlString
    | jqlNumber
    ;

jqlFunction
    : jqlFunctionName LPAREN jqlArgumentList? RPAREN
    ;

jqlFunctionName
    : jqlString
    | jqlNumber
    ;

jqlArgumentList
    : jqlArgument (COMMA jqlArgument)*
    ;

jqlList
    : jqlListStart jqlOperand (COMMA jqlOperand)* jqlListEnd
    ;

// Parenthesis are re-defined as rules in the list context to support an autocomplete feature where we show operands
// at caret positions where a list can been opened, and then auto-insert the opening parenthesis for the user

jqlListStart
    : LPAREN
    ;

jqlListEnd
    : RPAREN
    ;

jqlPropertyArgument
    : jqlArgument
    ;

jqlArgument
    : jqlString
    | jqlNumber
    ;

jqlWasPredicate
    : (jqlDatePredicateOperator | jqlDateRangePredicateOperator | jqlUserPredicateOperator) jqlPredicateOperand
    ;

jqlChangedPredicate
    : (jqlDatePredicateOperator | jqlDateRangePredicateOperator | jqlUserPredicateOperator | jqlValuePredicateOperator) jqlPredicateOperand
    ;

jqlDatePredicateOperator
    : AFTER | BEFORE | ON
    ;

jqlDateRangePredicateOperator
    : DURING
    ;

jqlUserPredicateOperator
    : BY
    ;

jqlValuePredicateOperator
    : FROM | TO
    ;

jqlPredicateOperand
    : jqlOperand
    ;

jqlOrderBy
    : ORDER BY jqlSearchSort (COMMA jqlSearchSort)*
    ;

jqlSearchSort
	: jqlField (DESC | ASC)?
	;
