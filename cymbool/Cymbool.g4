grammar Cymbool;

file
    : (functionDecl | varDecl)+
    ;

functionDecl
    : type ID '(' formalParameters? ')' block
    ;


formalParameters
    : formalParameter (',' formalParameter)*
    ;

formalParameter
    :  type ID
    ;

varDecl
    : type ID ('=' expr)? ';'
    ;

type
    : 'float'
    | 'int'
    | 'void'
    ;

block
    : '{' stat* '}'
    ;

stat
    : block
    | varDecl
    | 'if' expr 'then' stat+ ('else' stat+)?
    | 'return' expr? ';'
    | expr '=' expr ';'
    | expr ';'
    ;

expr
    : ID '(' exprList? ')'
    | expr '[' expr ']'
    | '-' expr
    | '!' expr
    | expr '*' expr
    | expr ('+' | '-' ) expr
    | expr '==' expr
    | ID
    | INT
    | '(' expr ')'
    ;

exprList
    : expr (',' expr)*
    ;

ID
    : [a-zA-Z] [0-9a-zA-Z]*
    ;

INT
    : '0'
    | [1-9] [0-9]*
    ;

WS
    : [ \t\r\n] -> skip
    ;