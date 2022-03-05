lexer grammar CommonLexerRule;

ID: [a-zA-Z]+;
INT: [0-9]+;
NEWLINE: '\r'? '\n' ;
WS: [\t]+ -> skip;