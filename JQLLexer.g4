lexer grammar JQLLexer;

// Some significant characters that need to be matched
LPAREN : '(';
RPAREN : ')';
COMMA : ',';
LBRACKET : '[';
RBRACKET : ']';

// Opening quotes
LQUOTE : QUOTE -> more, mode(QUOTED_STR);
LSQUOTE : SQUOTE -> more, mode(SQUOTED_STR);

// JQL Operators
BANG : '!';
LT : '<';
GT : '>';
GTEQ : '>=';
LTEQ : '<=';
EQUALS : '=';
NOT_EQUALS : '!=';
LIKE : '~';
NOT_LIKE : '!~';
IN : I N;
IS : I S;
AND : (A N D) | AMPER | AMPER_AMPER;
OR : (O R) | PIPE | PIPE_PIPE;
NOT : N O T;
EMPTY : (E M P T Y) | (N U L L);
WAS : W A S;
CHANGED : C H A N G E D;
BEFORE : B E F O R E;
AFTER : A F T E R;
FROM : F R O M;
TO : T O;
ON : O N;
DURING : D U R I N G;

// Order by
ORDER : O R D E R;
BY : B Y;
ASC : A S C;
DESC : D E S C;

// Numbers
POSNUMBER : DIGIT+;
NEGNUMBER : MINUS DIGIT+;

// Custom field prefix
CUSTOMFIELD : C F;

// Tokens defined above don't need to be included in the list of reserved words (they will match before reaching this point)
RESERVED_WORD : (
 A B O R T | A C C E S S | A D D | A L I A S | A L T E R | A N Y | A S | A U D I T | A V G | B E G I N | B E T W E E N
 | B O O L E A N | B R E A K | B Y T E | C A T C H | C H A R | C H A R A C T E R | C H E C K | C H E C K P O I N T
 | C O L L A T E | C O L L A T I O N | C O L U M N | C O M M I T | C O N N E C T | C O N T I N U E | C O U N T
 | C R E A T E | C U R R E N T | D A T E | D E C I M A L | D E C L A R E | D E C R E M E N T | D E F A U L T
 | D E F A U L T S | D E F I N E | D E L E T E | D E L I M I T E R | D I F F E R E N C E | D I S T I N C T
 | D I V I D E | D O | D O U B L E | D R O P | E L S E | E N C O D I N G | E N D | E Q U A L S | E S C A P E
 | E X C L U S I V E | E X E C | E X E C U T E | E X I S T S | E X P L A I N | F A L S E | F E T C H | F I L E
 | F I E L D | F I R S T | F L O A T | F O R | F U N C T I O N | G O | G O T O | G R A N T | G R E A T E R | G R O U P
 | H A V I N G | I D E N T I F I E D | I F | I M M E D I A T E | I N C R E M E N T | I N D E X | I N I T I A L
 | I N N E R | I N O U T | I N P U T | I N S E R T | I N T | I N T E G E R | I N T E R S E C T | I N T E R S E C T I O N
 | I N T O | I S E M P T Y | I S N U L L | J O I N | L A S T | L E F T | L E S S | L I K E | L I M I T | L O C K
 | L O N G | M A X | M I N | M I N U S | M O D E | M O D I F Y | M O D U L O | M O R E | M U L T I P L Y | N E X T
 | N O A U D I T | N O T I N | N O W A I T | N U M B E R | O B J E C T | O F | O P T I O N | O U T E R | O U T P U T
 | P O W E R | P R E V I O U S | P R I O R | P R I V I L E G E S | P U B L I C | R A I S E | R A W | R E M A I N D E R
 | R E N A M E | R E S O U R C E | R E T U R N | R E T U R N S | R E V O K E | R I G H T | R O W | R O W I D
 | R O W N U M | R O W S | S E L E C T | S E S S I O N | S E T | S H A R E | S I Z E | S Q R T | S T A R T
 | S T R I C T | S T R I N G | S U B T R A C T | S U M | S Y N O N Y M | T A B L E | T H E N | T R A N S
 | T R A N S A C T I O N | T R I G G E R | T R U E | U I D | U N I O N | U N I Q U E | U P D A T E | U S E R
 | V A L I D A T E | V A L U E S | V I E W | W H E N | W H E N E V E R | W H E R E | W H I L E | W I T H
);

// String handling
STRING : (ESCAPE | VALID_UNQUOTED_CHARS)+;

// Match any whitespace and then ignore it
MATCHWS : WS+ -> channel(HIDDEN);

/**
 * This is yet another large hack. We want this to be a different error message for reserved characters so we just match
 * them and try to recover.
 */
ERROR_RESERVED : RESERVED_CHARS;

/**
 * This is a large HACK. Match any character that we did not match using one of the previous rules. It must be an error
 * so lets try and do something.
 */
ERRORCHAR : .;


// ####### STRING MODES #######

// See https://github.com/antlr/antlr4/blob/master/doc/lexer-rules.md for more info on lexical modes
// TL;DR: Lexer modes allow us to have a finer grain control on what tokens we return to the main lexer mode
// (DEFAULT_MODE) in certain scenarios, e.g. after opening a quote. Two tokens can be returned here:
//  - QUOTE_STRING / SQUOTE_STRING : valid string
//  - INVALID_QUOTE_STRING / INVALID_SQUOTE_STRING : invalid / unclosed string

mode QUOTED_STR;
QUOTE_STRING : QUOTE -> mode(DEFAULT_MODE);
QUOTED_CONTENT : (ESCAPE | VALID_QUOTED_CHARS) -> more;
UNCLOSED_QUOTE_STRING : EOF -> mode(DEFAULT_MODE);
INVALID_QUOTE_STRING : . -> mode(DEFAULT_MODE);

mode SQUOTED_STR;
SQUOTE_STRING : SQUOTE -> mode(DEFAULT_MODE);
SQUOTED_CONTENT : (ESCAPE | VALID_SQUOTED_CHARS) -> more;
UNCLOSED_SQUOTE_STRING : EOF -> mode(DEFAULT_MODE);
INVALID_SQUOTE_STRING : . -> mode(DEFAULT_MODE);


// ####### FRAGMENTS #######

// Rule references are not supported in a set while negating in ANTLR4
fragment VALID_UNQUOTED_CHARS : ~(
    '\\' // BSLASH
    | ' ' | '\t' | '\r' | '\n' // WS
    // NOTE: This list needs to be synchronised with JqlStringSupportImpl.isJqlControlCharacter
    | '\u0000'..'\u0009' | '\u000b'..'\u000c' | '\u000e'..'\u001f' | '\u007f'..'\u009f'| '\ufdd0'..'\ufdef'| '\ufffe'..'\uffff' // CONTROLCHARS
    | '"' // QUOTE
    | '\'' // SQUOTE
    | '=' // EQUALS
    | '!' // BANG
    | '<' // LT
    | '>' // GT
    | '(' // LPAREN
    | ')' // RPAREN
    | '~' // LIKE
    | ',' // COMMA
    | '[' // LBRACKET
    | ']' // RBRACKET
    | '|' // PIPE
    | '&' // AMPER
    | '{' | '}' | '*' | '/' | '%' | '+' | '^' | '$' | '#' | '@' | '?' | ';' // RESERVED_CHARS
);

// Rule references are not supported in a set while negating in ANTLR4
fragment VALID_QUOTED_CHARS : ~(
    '\\' // BSLASH
    | '"' // QUOTE
    // NOTE: This list needs to be synchronised with JqlStringSupportImpl.isJqlControlCharacter
    | '\u0000'..'\u0009' | '\u000b'..'\u000c' | '\u000e'..'\u001f' | '\u007f'..'\u009f'| '\ufdd0'..'\ufdef'| '\ufffe'..'\uffff' // CONTROLCHARS
);

// Rule references are not supported in a set while negating in ANTLR4
fragment VALID_SQUOTED_CHARS : ~(
    '\\' // BSLASH
    | '\'' // SQUOTE
    // NOTE: This list needs to be synchronised with JqlStringSupportImpl.isJqlControlCharacter
    | '\u0000'..'\u0009' | '\u000b'..'\u000c' | '\u000e'..'\u001f' | '\u007f'..'\u009f'| '\ufdd0'..'\ufdef'| '\ufffe'..'\uffff' // CONTROLCHARS
);

/**
 * These are some characters that we do not use now but we want to reserve. We have not reserved MINUS because we
 * really really really don't want to force people into quoting issues keys and dates.
 */
fragment RESERVED_CHARS
    : '{' | '}'
    | '*' | '/' | '%' | '+' | '^'
    | '$' | '#' | '@'
    | '?' | ';'
;

fragment QUOTE : '"';
fragment SQUOTE : '\'';
fragment BSLASH : '\\';
fragment NL : '\r';
fragment CR : '\n';
fragment SPACE : ' ';
fragment AMPER : '&';
fragment AMPER_AMPER : '&&';
fragment PIPE : '|';
fragment PIPE_PIPE : '||';

fragment ESCAPE
    :   BSLASH
    (
        't'
        | 'n'
        | 'r'
        | QUOTE
        | SQUOTE
        | BSLASH
        | SPACE
        | 'u' HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT
    )
;

fragment NEWLINE : NL | CR;

fragment HEXDIGIT : DIGIT | A | B | C | D | E | F;

fragment WS : (SPACE|'\t'|NEWLINE);

fragment MINUS : '-';

fragment DIGIT : [0-9];

fragment A : [aA];
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];
