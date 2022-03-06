lexer grammar CommonLexerRule;

ID:         [a-zA-Z]+;
NUMBER:     [0-9]+('.'[0-9]+)?;
NEWLINE:    '\r'? '\n' ;
WS:         [\t]+ -> skip;

MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
