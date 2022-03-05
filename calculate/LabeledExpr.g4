grammar LabeledExpr;
import CommonLexerRule;

prog: stat+;
stat: expr NEWLINE              # printExpr
    |ID '=' expr NEWLINE        # assgin
    | NEWLINE                   # blank
    ;
expr: expr op=('*'|'/') expr    # MulDiv
    | expr op=('-'|'+') expr    # AddSub
    |INT                        # int
    |ID                         # id
    |'(' expr ')'               # parens
    ;

MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';