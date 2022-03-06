grammar JSON;

json
    : object
    | array
    ;

object
    : LEFT_BRACE pair (',' pair)* RIGHT_BRACE
    | LEFT_BRACE RIGHT_BRACE
    ;

pair
    : STRING COLON value
    ;

array
    : OPEN_BRACKET value (COMMA value)* CLOSE_BRACKET
    | OPEN_BRACKET CLOSE_BRACKET
    ;

value
    : STRING
    | NUMBER
    | object
    | array
    | TRUE
    | FALSE
    | NULL
    ;

STRING
    : '"' (ESC | ~["\\])* '"'
    ;

fragment ESC
    : '\\' (["\\ /bfnrt] | UNICODE)
    ;

fragment UNICODE
    : 'u' HEX HEX HEX HEX
    ;

fragment HEX
    : [0-9a-fA-F]
    ;

// float
NUMBER
    : '-'? INT '.' INT EXP?
    | '-'? INT EXP
    | '-'? INT
    ;

fragment INT
    : '0'
    | [1-9] [0-9]*
    ;

fragment EXP
    : [Ee] [+/-]? INT
    ;

TRUE
    : 'true'
    ;

FALSE
    : 'false'
    ;

NULL
    : 'null'
    ;

LEFT_BRACE
    : '{'
    ;

RIGHT_BRACE
    : '}'
    ;

OPEN_BRACKET
    : '['
    ;

CLOSE_BRACKET
    : ']'
    ;

COMMA
    : ','
    ;

COLON
    : ':'
    ;

WS
    : [ \t\n\r]+ -> skip
    ;