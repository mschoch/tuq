%{
package main
import "fmt"
import "log"

func logDebugGrammar(format string, v ...interface{}) {
    if *debugGrammar && len(v) > 0 {
        log.Printf("DEBUG GRAMMAR " + format, v)
    } else if *debugGrammar {
        log.Printf("DEBUG GRAMMAR " + format)
    }
}
%}

%union { 
s string 
n int
f float64}

%token INT REAL STRING IDENTIFIER PROPERTY NEWLINE
%token LPAREN RPAREN COMMA LBRACE RBRACE
%token SELECT DISTINCT ALL AS FROM WHERE GROUP BY
%token HAVING FLATTEN EACH UNION INTERSECT EXCEPT ORDER
%token LIMIT OFFSET ASC DESC TRUE FALSE LBRACKET RBRACKET
%token QUESTION COLON MAX MIN AVG COUNT SUM DOT
%token PLUS MINUS MULT DIV MOD AND OR NOT EQ LT LTE
%token GT GTE NE PRAGMA ASSIGN DEBUG
%left OR
%left AND
%left EQ LT LTE GT GTE NE
%left PLUS MINUS MULT DIV MOD
%right NOT
%right QUESTION
%%
input: select_stmt { logDebugGrammar("INPUT") }
        | pragma_stmt
;

pragma_stmt:    PRAGMA pragma_name ASSIGN expression { logDebugGrammar("PRAGMA: %v", $1)
                                                       right := parsingStack.Pop()
                                                       left := parsingStack.Pop()
                                                       ProcessPragma(left.(map[string]interface{}), right.(map[string]interface{}))
                                                     }
;

pragma_name:    DEBUG { curr := make(map[string]interface{})
                        curr["pragma"] = "debug"
                        parsingStack.Push(curr) }
;

select_stmt:     select_compound select_order select_limit_offset { logDebugGrammar("SELECT_STMT")
                                                                    parsingQuery.parsedSuccessfully = true }
;

select_order:   /* empty */
        |   ORDER BY sorting_list
;

select_limit_offset:    /* empty */
        |   select_limit
        |   select_limit select_offset
;

select_limit:   LIMIT expression { curr := make(map[string]interface{})
                                   curr["op"] = "limit"
                                   curr["expression"] = parsingStack.Pop()
                                   parsingQuery.limit = curr
                                 }
;

select_offset:  OFFSET expression { curr := make(map[string]interface{})
                                   curr["op"] = "offset"
                                   curr["expression"] = parsingStack.Pop()
                                   parsingQuery.offset = curr
                                 }
;

sorting_list:   sorting_single
        | sorting_single COMMA sorting_list
;

sorting_single: expression { curr := make(map[string]interface{})
                             curr["op"] = "sort"
                             curr["expression"] = parsingStack.Pop()
                             curr["ascending"] = true
                             parsingQuery.orderby = append(parsingQuery.orderby, curr)
                           }
        |   expression ASC { curr := make(map[string]interface{})
                             curr["op"] = "sort"
                             curr["expression"] = parsingStack.Pop()
                             curr["ascending"] = true
                             parsingQuery.orderby = append(parsingQuery.orderby, curr)
                           }
        |   expression DESC { curr := make(map[string]interface{})
                              curr["op"] = "sort"
                              curr["expression"] = parsingStack.Pop()
                              curr["ascending"] = false
                              parsingQuery.orderby = append(parsingQuery.orderby, curr)
                           }
;

select_compound:    select_core { logDebugGrammar("SELECT_COMPOUND") }
        |   select_core compound_operator select_compound
;

compound_operator:  UNION
        |   UNION ALL
        |   INTERSECT
        |   EXCEPT
;

select_core:    select_select select_from select_where select_group_having { logDebugGrammar("SELECT_CORE") }
;

select_group: GROUP BY expression_list { logDebugGrammar("SELECT GROUP")
                                         parsingQuery.isAggregateQuery = true 
                                         parsingQuery.groupby = parsingStack.Pop().([]interface{}) }
;

select_group_having:    /*empty*/ { logDebugGrammar("SELECT GROUP HAVING - EMPTY") }
        |   select_group { logDebugGrammar("SELECT GROUP HAVING - SELECT GROUP") }
        |   select_group select_having { logDebugGrammar("SELECT GROUP HAVING - SELECT GROUP SELECT HAVING") }
;



select_having:  HAVING expression
;

select_where:   /* empty */ { logDebugGrammar("SELECT WHERE - EMPTY") }
        |   WHERE expression { logDebugGrammar("SELECT WHERE - EXPR")
                               where_part := parsingStack.Pop()
                               parsingQuery.where = where_part.(map[string]interface{}) }
;

select_from:    /* empty */
        |   FROM data_source_list   { logDebugGrammar("SELECT_FROM") }
;

data_source_list:   data_source_single
        |   data_source_single COMMA data_source_list
;

data_source_single:  data_source { ds := NewDataSource($1.s)
                                   parsingQuery.AddDataSource(ds) 
                                 }
        |   data_source AS IDENTIFIER   { ds := NewDataSourceWithAs($1.s, $3.s) 
                                          parsingQuery.AddDataSource(ds) 
                                        }
;

data_source:    IDENTIFIER

select_select:  select_select_head select_select_tail { logDebugGrammar("SELECT_SELECT") }
;

select_select_head:  SELECT { logDebugGrammar("SELECT_SELECT_HEAD") }
        |       SELECT ALL
        |       SELECT DISTINCT
;

select_select_tail:     /* empty */ { logDebugGrammar("SELECT SELECT TAIL - EMPTY") } 
        |      expression { logDebugGrammar("SELECT SELECT TAIL - EXPR")
                            curr := make(map[string]interface{})
                            curr["op"] = "select"
                            curr["expression"] = parsingStack.Pop()
                            parsingQuery.sel = curr
                          }
        |      expression AS IDENTIFIER { logDebugGrammar("SELECT SELECT TAIL - EXPR AS IDENTIFIER")
                                          curr := make(map[string]interface{})
                                          curr["op"] = "select"
                                          curr["expression"] = parsingStack.Pop()
                                          curr["as"] = $3.s
                                          parsingQuery.sel = curr
                                        }
        |      AS IDENTIFIER { logDebugGrammar("SELECT SELECT TAIL - AS IDENTIFIER")
                               curr := make(map[string]interface{})
                               curr["op"] = "select"
                               curr["as"] = $2.s
                               parsingQuery.sel = curr
                             }
;

expression: expr { logDebugGrammar("EXPRESSION") }
    |   expr QUESTION expression COLON expression
;

expr: expr PLUS expr
        |   expr MINUS expr
        |   expr MULT expr
        |   expr DIV expr
        |   expr AND expr {  logDebugGrammar("EXPR - AND")
                             
                            curr := make(map[string]interface{})
                            curr["right"] = parsingStack.Pop()
                            curr["left"] = parsingStack.Pop()
                            logDebugGrammar("LHS %v and RHS %v", curr["left"], curr["right"])
                            curr["op"] = "&&"
                            parsingStack.Push(curr)
                         }
        |   expr OR expr {  logDebugGrammar("EXPR - OR")
                            curr := make(map[string]interface{})
                            curr["right"] = parsingStack.Pop()
                            curr["left"] = parsingStack.Pop()
                            logDebugGrammar("LHS %v and RHS %v", curr["left"], curr["right"])
                            curr["op"] = "||"
                            parsingStack.Push(curr)
                         }
        |   expr EQ expr {  logDebugGrammar("EXPR - EQ")
                            curr := make(map[string]interface{})
                            curr["right"] = parsingStack.Pop()
                            curr["left"] = parsingStack.Pop()
                            logDebugGrammar("LHS %v and RHS %v", curr["left"], curr["right"])
                            curr["op"] = "=="
                            parsingStack.Push(curr)
                         }
        |   expr LT expr {  logDebugGrammar("EXPR - LT")
                            curr := make(map[string]interface{})
                            curr["right"] = parsingStack.Pop()
                            curr["left"] = parsingStack.Pop()
                            logDebugGrammar("LHS %v and RHS %v", curr["left"], curr["right"])
                            curr["op"] = "<"
                            parsingStack.Push(curr)
                         }
        |   expr LTE expr {  logDebugGrammar("EXPR - LTE")
                            curr := make(map[string]interface{})
                            curr["right"] = parsingStack.Pop()
                            curr["left"] = parsingStack.Pop()
                            logDebugGrammar("LHS %v and RHS %v", curr["left"], curr["right"])
                            curr["op"] = "<="
                            parsingStack.Push(curr)
                         }
        |   expr GT expr {  logDebugGrammar("EXPR - GT")
                            curr := make(map[string]interface{})
                            right := parsingStack.Pop()
                            left := parsingStack.Pop()
                            curr["right"] = right
                            curr["left"] = left
                            logDebugGrammar("LHS %v and RHS %v", left, right)
                            curr["op"] = ">"
                            parsingStack.Push(curr)
                         }
        |   expr GTE expr {  logDebugGrammar("EXPR - GTE")
                            curr := make(map[string]interface{})
                            curr["right"] = parsingStack.Pop()
                            curr["left"] = parsingStack.Pop()
                            logDebugGrammar("LHS %v and RHS %v", curr["left"], curr["right"])
                            curr["op"] = ">="
                            parsingStack.Push(curr)
                         }
        |   expr NE expr {  logDebugGrammar("EXPR - NE")
                            curr := make(map[string]interface{})
                            curr["right"] = parsingStack.Pop()
                            curr["left"] = parsingStack.Pop()
                            logDebugGrammar("LHS %v and RHS %v", curr["left"], curr["right"])
                            curr["op"] = "!="
                            parsingStack.Push(curr)
                         }
        |   expr MOD expr        
        |   prefix_expr
;


prefix_expr: NOT prefix_expr
    | suffix_expr
;

suffix_expr: atom { logDebugGrammar("SUFFIX_EXPR") }
;

atom: property { curr := make(map[string]interface{})
                 curr["property"] = $1.s
                 parsingStack.Push(curr) }
    |   IDENTIFIER { curr := make(map[string]interface{})
                 curr["identifier"] = $1.s
                 parsingStack.Push(curr) }
    |   INT { curr := make(map[string]interface{})
                 curr["int"] = $1.n
                 parsingStack.Push(curr) }
    |   REAL { curr := make(map[string]interface{})
                 curr["real"] = $1.f
                 parsingStack.Push(curr) }
    |   STRING { curr := make(map[string]interface{})
                 curr["string"] = $1.s
                 parsingStack.Push(curr) }
    |   TRUE { curr := make(map[string]interface{})
                 curr["bool"] = true
                 parsingStack.Push(curr) }
    |   FALSE { curr := make(map[string]interface{})
                 curr["bool"] = false
                 parsingStack.Push(curr) }
    |   LBRACE named_expression_list RBRACE { logDebugGrammar("ATOM - {}")
                                              curr := make(map[string]interface{})
                                              curr["object"] = parsingStack.Pop()
                                              parsingStack.Push(curr)
                                            }
    |   LBRACKET expression_list RBRACKET { logDebugGrammar("ATOM - []")
                                            curr := make(map[string]interface{})
                                            curr["array"] = parsingStack.Pop()
                                            parsingStack.Push(curr)
                                          }
    |   function_name LPAREN expression_list RPAREN { logDebugGrammar("FUNCTION - $1.s")
                                                      expression_list := parsingStack.Pop().([]interface{})
                                                      function_map := parsingStack.Pop().(map[string]interface{})
                                                      function_map["expression_list"] = expression_list
                                                      parsingStack.Push(function_map)
                                                    }
    |   LPAREN expression RPAREN
    |   LPAREN select_stmt RPAREN
;

expression_list:  expression { logDebugGrammar("EXPRESSION_LIST - EXPRESSION")
                                curr := make([]interface{},0)
                               curr = append(curr, parsingStack.Pop())
                               parsingStack.Push(curr)
                             }
        |   expression COMMA expression_list { logDebugGrammar("EXPRESSION_LIST - EXPRESSION COMMA EXPRESSION_LIST")
                                               rest := parsingStack.Pop().([]interface{})
                                               last := parsingStack.Pop().(interface{})
                                               curr := make([]interface{},0)
                                               curr = append(curr, last)
                                               for _,v := range rest {
                                                 curr = append(curr, v)
                                               }
                                               parsingStack.Push(curr)
                                             }
;

named_expression_list:    named_expression_single 
        |   named_expression_single COMMA named_expression_list { last := parsingStack.Pop().(map[string]interface{})
                                                                  rest := parsingStack.Pop().(map[string]interface{})
                                                                  for k,v := range last {
                                                                    rest[k] = v
                                                                  }
                                                                  parsingStack.Push(rest)
                                                                }
;

named_expression_single:   STRING COLON expression { curr := make(map[string]interface{})
                                                     curr[$1.s] = parsingStack.Pop()
                                                     parsingStack.Push(curr)
                                                   }
;

property:   PROPERTY
;

function_name: MIN { parsingQuery.isAggregateQuery = true
                     curr := make(map[string]interface{})
                     curr["function"] = "min"
                     parsingStack.Push(curr) }
        |   MAX { parsingQuery.isAggregateQuery = true
                  curr := make(map[string]interface{})
                  curr["function"] = "max"
                  parsingStack.Push(curr) }
        |   AVG { parsingQuery.isAggregateQuery = true
                  curr := make(map[string]interface{})
                  curr["function"] = "avg"
                  parsingStack.Push(curr) }
        |   COUNT { parsingQuery.isAggregateQuery = true
                    curr := make(map[string]interface{})
                    curr["function"] = "count"
                    parsingStack.Push(curr) }
        |   SUM { parsingQuery.isAggregateQuery = true
                    curr := make(map[string]interface{})
                    curr["function"] = "sum"
                    parsingStack.Push(curr) }
;
%%
