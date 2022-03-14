grammar Errors;
prog
    : classDef+
    ;

classDef
    : 'class' ID '{' member+ '}'
      {fmt.Println("class "+ $ID.text);}
    ;

member
    : 'int' ID ';'
       {fmt.Println("var "+ $ID.text);}
    | 'int' f=ID '(' ID ')' '{' stat '}'
      {fmt.Println("method "+ $f.text);}
    ;

stat
    : expr ';'
      {fmt.Println("found expr "+ $ctx.GetText());}
    | ID '=' expr';'
      {fmt.Println("found assgin "+ $ctx.GetText());}
    ;

expr
    : INT
    | ID '(' INT ')'
    ;

ID
    : [a-zA-Z]+
    ;

INT
    : [0-9]+
    ;
WS
    : [ \t\r\n] -> skip
    ;