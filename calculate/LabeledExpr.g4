grammar LabeledExpr;
import CommonLexerRule;

prog
    : stat+
    ;

stat
    : expr NEWLINE            # printExpr
    | ID '=' expr NEWLINE     # assgin
    | NEWLINE                 # blank
    ;

expr
    : expr op=('*'|'/') expr  # MulDiv
    | expr op=('-'|'+') expr  # AddSub
    | NUMBER                  # number
    | ID                      # id
    | '(' expr ')'            # parens
    ;