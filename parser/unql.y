%{
package parser
import "fmt"
import "log"

func logDebugGrammar(format string, v ...interface{}) {
    if debugGrammar && len(v) > 0 {
        log.Printf("DEBUG GRAMMAR " + format, v)
    } else if debugGrammar {
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
%token HAVING UNION INTERSECT EXCEPT ORDER
%token LIMIT OFFSET ASC DESC TRUE FALSE LBRACKET RBRACKET
%token QUESTION COLON MAX MIN AVG COUNT SUM DOT
%token PLUS MINUS MULT DIV MOD AND OR NOT EQ LT LTE
%token GT GTE NE PRAGMA ASSIGN EXPLAIN NULL OVER
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

pragma_stmt:    PRAGMA expression ASSIGN expression { logDebugGrammar("PRAGMA: %v", $1)
                                                       right := parsingStack.Pop()
                                                       left := parsingStack.Pop()
                                                       ProcessPragma(left.(Expression), right.(Expression))
                                                     }
;

select_stmt:     select_explain select_compound select_order select_limit_offset { logDebugGrammar("SELECT_STMT")
                                                                    parsingQuery.parsedSuccessfully = true }
;

select_explain:  /* empty */
        |   EXPLAIN {  parsingQuery.isExplainOnly = true  }
;

select_order:   /* empty */
        |   ORDER BY sorting_list
;

select_limit_offset:    /* empty */
        |   select_limit
        |   select_limit select_offset
;

select_limit:   LIMIT expression { thisExpression := parsingStack.Pop()
                                   parsingQuery.Limit = thisExpression.(Expression)
                                 }
;

select_offset:  OFFSET expression { thisExpression := parsingStack.Pop()
                                   parsingQuery.Offset = thisExpression.(Expression)
                                 }
;

sorting_list:   sorting_single
        | sorting_single COMMA sorting_list
;

sorting_single: expression { thisExpression := NewSortItem(parsingStack.Pop().(Expression), true)
                            parsingQuery.Orderby = append(parsingQuery.Orderby, *thisExpression)
                           }
        |   expression ASC { thisExpression := NewSortItem(parsingStack.Pop().(Expression), true)
                            parsingQuery.Orderby = append(parsingQuery.Orderby, *thisExpression)
                           }
        |   expression DESC { thisExpression := NewSortItem(parsingStack.Pop().(Expression), false)
                            parsingQuery.Orderby = append(parsingQuery.Orderby, *thisExpression)
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
                                         parsingQuery.Groupby = parsingStack.Pop().(ExpressionList) }
;

select_group_having:    /*empty*/ { logDebugGrammar("SELECT GROUP HAVING - EMPTY") }
        |   select_group { logDebugGrammar("SELECT GROUP HAVING - SELECT GROUP") }
        |   select_group select_having { logDebugGrammar("SELECT GROUP HAVING - SELECT GROUP SELECT HAVING") }
;



select_having:  HAVING expression { parsingQuery.Having = parsingStack.Pop().(Expression) }
;

select_where:   /* empty */ { logDebugGrammar("SELECT WHERE - EMPTY") }
        |   WHERE expression { logDebugGrammar("SELECT WHERE - EXPR")
                               where_part := parsingStack.Pop()
                               parsingQuery.Where = where_part.(Expression) }
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
        |   data_source data_source_over_list { ds := NewDataSource($1.s)
                                          nextOver := parsingStack.Pop()
                                          for nextOver != nil {
                                            ds.AddOver(nextOver.(*Over))
                                            nextOver = parsingStack.Pop()
                                          }
                                          parsingQuery.AddDataSource(ds)
                                        }
        |   data_source AS IDENTIFIER   { ds := NewDataSourceWithAs($1.s, $3.s) 
                                          parsingQuery.AddDataSource(ds) 
                                        }
        |   data_source AS IDENTIFIER data_source_over_list { ds := NewDataSourceWithAs($1.s, $3.s)
                                          nextOver := parsingStack.Pop()
                                          for nextOver != nil {
                                            ds.AddOver(nextOver.(*Over))
                                            nextOver = parsingStack.Pop()
                                          }
                                          parsingQuery.AddDataSource(ds)
                                        }
;

data_source_over_list:  data_source_over
        |   data_source_over data_source_over_list
;

data_source_over:   OVER property AS IDENTIFIER {   prop := parsingStack.Pop().(*Property)
                                                    over := NewOver(prop, $4.s)
                                                    parsingStack.Push(over)
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
                            thisExpression := parsingStack.Pop()
                            parsingQuery.Sel = thisExpression.(Expression)
                          }
;

expression: expr { logDebugGrammar("EXPRESSION") }
    |   expr QUESTION expression COLON expression { logDebugGrammar("EXPRESSION - TERNARY")
                                                    elsee := parsingStack.Pop().(Expression)
                                                    thenn := parsingStack.Pop().(Expression)
                                                    iff := parsingStack.Pop().(Expression)
                                                    thisExpr := NewTernaryExpression(iff, thenn, elsee)
                                                    parsingStack.Push(thisExpr)
                                                  }
;

expr: expr PLUS expr {  logDebugGrammar("EXPR - PLUS")
                        right := parsingStack.Pop()
                        left := parsingStack.Pop()
                        thisExpression := NewPlusExpression(left.(Expression), right.(Expression)) 
                        parsingStack.Push(thisExpression)
                     }
        |   expr MINUS expr {  logDebugGrammar("EXPR - MINUS")
                               right := parsingStack.Pop()
                               left := parsingStack.Pop()
                               thisExpression := NewMinusExpression(left.(Expression), right.(Expression)) 
                               parsingStack.Push(thisExpression)
                            }
        |   expr MULT expr {  logDebugGrammar("EXPR - MULT")
                              right := parsingStack.Pop()
                              left := parsingStack.Pop()
                              thisExpression := NewMultiplyExpression(left.(Expression), right.(Expression)) 
                              parsingStack.Push(thisExpression)
                           }
        |   expr DIV expr {  logDebugGrammar("EXPR - DIV")
                             right := parsingStack.Pop()
                             left := parsingStack.Pop()
                             thisExpression := NewDivideExpression(left.(Expression), right.(Expression)) 
                             parsingStack.Push(thisExpression)
                          }
        |   expr AND expr {  logDebugGrammar("EXPR - AND")
                             right := parsingStack.Pop()
                             left := parsingStack.Pop()
                             thisExpression := NewAndExpression(left.(Expression), right.(Expression)) 
                             parsingStack.Push(thisExpression)
                         }
        |   expr OR expr {  logDebugGrammar("EXPR - OR")
                            right := parsingStack.Pop()
                            left := parsingStack.Pop()
                            thisExpression := NewOrExpression(left.(Expression), right.(Expression)) 
                            parsingStack.Push(thisExpression)
                         }
        |   expr EQ expr {  logDebugGrammar("EXPR - EQ")
                            right := parsingStack.Pop()
                            left := parsingStack.Pop()
                            thisExpression := NewEqualsExpression(left.(Expression), right.(Expression)) 
                            parsingStack.Push(thisExpression)
                         }
        |   expr LT expr {  logDebugGrammar("EXPR - LT")
                            right := parsingStack.Pop()
                            left := parsingStack.Pop()
                            thisExpression := NewLessThanExpression(left.(Expression), right.(Expression)) 
                            parsingStack.Push(thisExpression)
                         }
        |   expr LTE expr {  logDebugGrammar("EXPR - LTE")
                             right := parsingStack.Pop()
                             left := parsingStack.Pop()
                             thisExpression := NewLessThanOrEqualExpression(left.(Expression), right.(Expression)) 
                             parsingStack.Push(thisExpression)
                         }
        |   expr GT expr {  logDebugGrammar("EXPR - GT")
                            right := parsingStack.Pop()
                            left := parsingStack.Pop()
                            thisExpression := NewGreaterThanExpression(left.(Expression), right.(Expression)) 
                            parsingStack.Push(thisExpression)
                         }
        |   expr GTE expr {  logDebugGrammar("EXPR - GTE")
                             right := parsingStack.Pop()
                             left := parsingStack.Pop()
                             thisExpression := NewGreaterThanOrEqualExpression(left.(Expression), right.(Expression)) 
                             parsingStack.Push(thisExpression)
                         }
        |   expr NE expr {  logDebugGrammar("EXPR - NE")
                            right := parsingStack.Pop()
                            left := parsingStack.Pop()
                            thisExpression := NewNotEqualsExpression(left.(Expression), right.(Expression)) 
                            parsingStack.Push(thisExpression)
                         }
        |   expr MOD expr        
        |   prefix_expr
;


prefix_expr: NOT prefix_expr { logDebugGrammar("EXPR - NOT")
                               curr := parsingStack.Pop().(Expression)
                               thisExpression := NewNotExpression(curr)
                               parsingStack.Push(thisExpression)
                             }
    | suffix_expr
;

suffix_expr: atom { logDebugGrammar("SUFFIX_EXPR") }
;

atom: NULL { logDebugGrammar("NULL")
             thisExpression := NewNull()
             parsingStack.Push(thisExpression)
           }
    | property {  }
    | property LBRACKET expression RBRACKET {     logDebugGrammar("ATOM - prop[]")
                                                  rightExpr := parsingStack.Pop().(Expression)
                                                  leftProp := parsingStack.Pop().(*Property)
                                                  thisExpression := NewBracketMemberExpression(leftProp, rightExpr)
                                                  parsingStack.Push(thisExpression)
                                            }
    |   INT { thisExpression := NewIntegerLiteral($1.n) 
                 parsingStack.Push(thisExpression) }
    |   MINUS INT { thisExpression := NewIntegerLiteral(-$1.n)
                 parsingStack.Push(thisExpression) }
    |   REAL { thisExpression := NewFloatLiteral($1.f) 
                 parsingStack.Push(thisExpression) }
    |   MINUS REAL { thisExpression := NewFloatLiteral(-$1.f)
                 parsingStack.Push(thisExpression) }
    |   STRING { thisExpression := NewStringLiteral($1.s) 
                 parsingStack.Push(thisExpression) }
    |   TRUE { thisExpression := NewBoolLiteral(true) 
                 parsingStack.Push(thisExpression) }
    |   FALSE { thisExpression := NewBoolLiteral(false) 
                 parsingStack.Push(thisExpression)}
    |   LBRACE named_expression_list RBRACE { logDebugGrammar("ATOM - {}")
                                            }
    |   LBRACKET expression_list RBRACKET { logDebugGrammar("ATOM - []")
                                            exp_list := parsingStack.Pop().(ExpressionList)
                                            thisExpression := NewArrayLiteral(exp_list)
                                            parsingStack.Push(thisExpression)
                                          }
    |   function_name LPAREN expression_list RPAREN { logDebugGrammar("FUNCTION - $1.s")
                                                      exp_list := parsingStack.Pop().(ExpressionList)
                                                      function := parsingStack.Pop().(*Function)
                                                      function.AddArguments(exp_list)
                                                      parsingStack.Push(function)
                                                    }
    |   LPAREN expression RPAREN
    |   LPAREN select_stmt RPAREN
;

expression_list:  expression { logDebugGrammar("EXPRESSION_LIST - EXPRESSION")
                               exp_list := make(ExpressionList, 0)
                               exp_list = append(exp_list, parsingStack.Pop().(Expression))
                               parsingStack.Push(exp_list)
                             }
        |   expression COMMA expression_list { logDebugGrammar("EXPRESSION_LIST - EXPRESSION COMMA EXPRESSION_LIST")
                                               rest := parsingStack.Pop().(ExpressionList)
                                               last := parsingStack.Pop()
                                               new_list := make(ExpressionList, 0)
                                               new_list = append(new_list, last.(Expression))
                                               for _, v := range rest {
                                                new_list = append(new_list, v)
                                               }
                                               parsingStack.Push(new_list)
                                             }
;

named_expression_list:    named_expression_single 
        |   named_expression_single COMMA named_expression_list { last := parsingStack.Pop().(*ObjectLiteral)
                                                                  rest := parsingStack.Pop().(*ObjectLiteral)
                                                                  rest.AddAll(last)
                                                                  parsingStack.Push(rest)
                                                                }
;

named_expression_single:   STRING COLON expression { thisKey := $1.s
                                                     thisValue := parsingStack.Pop().(Expression)
                                                     thisExpression := NewObjectLiteral(Object{thisKey: thisValue})
                                                     parsingStack.Push(thisExpression) 
                                                   }
;

property:   IDENTIFIER {
                         thisExpression := NewProperty($1.s) 
                         parsingStack.Push(thisExpression) 
                       }
        | IDENTIFIER DOT property {
                                    thisValue := parsingStack.Pop().(*Property)
                                    thisExpression := NewProperty($1.s + "." + thisValue.Symbol)
                                    parsingStack.Push(thisExpression)
                                  }
;

function_name: MIN { 
                     parsingQuery.isAggregateQuery = true
                     thisExpression := NewFunction("min")
                     parsingStack.Push(thisExpression)
                   }
        |   MAX { 
                  parsingQuery.isAggregateQuery = true
                  thisExpression := NewFunction("max")
                  parsingStack.Push(thisExpression)
                }
        |   AVG { 
                  parsingQuery.isAggregateQuery = true
                  thisExpression := NewFunction("avg")
                  parsingStack.Push(thisExpression)
                }
        |   COUNT { 
                   parsingQuery.isAggregateQuery = true
                   thisExpression := NewFunction("count")
                   parsingStack.Push(thisExpression)
                  }
        |   SUM { 
                  parsingQuery.isAggregateQuery = true
                  thisExpression := NewFunction("sum")
                  parsingStack.Push(thisExpression)
                }
;
%%
