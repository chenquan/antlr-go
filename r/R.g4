grammar R;

prog
    : ( expr_or_assgin (';' | NL) | NL)*
    ;

expr_or_assgin
    : expr ('<-' | '=' | '<<-') expr_or_assgin
    | expr
    ;

expr
    : expr '[[' sub_list ']]'
    | expr '[' sub_list ']'
    | expr  ('::' | ':::') expr
    | <assoc=right> expr '^' expr
    | ('-' | '+') expr
    | expr ':' expr
    | expr USER_OP expr
    | expr ('*' | '/') expr
    | expr ('+' | '-') expr
    | expr ('>' | '>=' | '<' | '<=' | '==' | '!=') expr
    | '!' expr
    | expr ('&' | '&&') expr
    | expr ('|' | '||') expr
    | '~' expr
    | expr '~' expr
    | expr ('->' | '->>' | ':=' )expr
    | 'function' '(' form_list? ')' expr
    | expr '(' sub_list ')'
    | '{' expr_list '}'
    | 'if' '(' expr ')' expr
    | 'if' '(' expr ')' expr 'else' expr
    | 'for' '(' ID 'in' expr ')' expr
    | 'while' '(' expr ')' expr
    | 'repeat' expr
    | '?' expr
    | 'next'
    | 'break'
    | '(' expr ')'
    | ID
    | STRING
    | HEX
    | INT
    | FLOAT
    | COMPLEX
    | 'NULL'
    | 'NA'
    | 'Inf'
    | 'NaN'
    | 'TRUE'
    | 'FALSE'

    ;

expr_list
    : expr_or_assgin ((';' | NL) expr_or_assgin?)*
    ;

form_list
    : form (',' form)*
    ;

form
    : ID
    | ID '=' expr
    | '...'
    ;

sub_list
    : sub (',' sub)*
    ;

sub
    : expr
    | ID '='
    | ID '=' expr
    | STRING '='
    | STRING '=' expr
    | 'NULL' '='
    | 'NULL' '=' expr
    | '...'
    |
    ;

NL
    : '\r'? '\n'
    ;


HEX
    : '0' ('x'|'X') HEXDIGIT+ [Ll]?
    ;

INT
    : DIGIT+ [Ll]?
    ;

fragment HEXDIGIT
    : ('0'..'9'|'a'..'f'|'A'..'F')
    ;

FLOAT
    : DIGIT+ '.' DIGIT* EXP? [Ll]?
    | DIGIT+ EXP? [Ll]?
    | '.' DIGIT+ EXP? [Ll]?
    ;

fragment DIGIT
    : '0'..'9'
    ;

fragment EXP :   [Ee] ('+' | '-')? INT ;

COMPLEX
    : INT 'i'
    | FLOAT 'i'
    ;

STRING
    : '"' ( ESC | ~[\\"] )*? '"'
    | '\'' ( ESC | ~[\\'] )*? '\''
    ;

fragment ESC
    : '\\' ([abtnfrv]|'"'|'\'')
    | UNICODE_ESCAPE
    | HEX_ESCAPE
    | OCTAL_ESCAPE
    ;

fragment UNICODE_ESCAPE
    : '\\' 'u' HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT
    | '\\' 'u' '{' HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT '}'
    ;

fragment OCTAL_ESCAPE
    : '\\' [0-3] [0-7] [0-7]
    | '\\' [0-7] [0-7]
    | '\\' [0-7]
    ;

fragment HEX_ESCAPE
    : '\\' HEXDIGIT HEXDIGIT?
    ;

ID
    : '.' (LETTER|'_'|'.') (LETTER|DIGIT|'_'|'.')*
    | LETTER (LETTER|DIGIT|'_'|'.')*
    ;
fragment LETTER
    : [a-zA-Z]
    ;

USER_OP
    :'%' .*? '%'
    ;

COMMENT :   '#' .*? '\r'? '\n' -> type(NL) ;

WS      :   [ \t]+ -> skip ;
