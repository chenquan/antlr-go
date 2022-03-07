grammar CSV;

file
    : hdr row+;

hdr
    : row;

row
    : field (',' field)* '\r'? '\n';

field
    : TEXT      # text
    | STRING    # string
    |           # empty
    ;

TEXT
    : ~[,\t\n]+;

STRING
    : '"' ('""' | ~'"')* '"';
