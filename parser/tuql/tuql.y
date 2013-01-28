%{
package tuql
import "fmt"
import "log"
import "github.com/mschoch/tuq/parser"

func logDebugGrammar(format string, v ...interface{}) {
    if parser.DebugGrammar && len(v) > 0 {
        log.Printf("DEBUG GRAMMAR " + format, v)
    } else if parser.DebugGrammar {
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
                                                       parser.ProcessPragma(left.(parser.Expression), right.(parser.Expression))
                                                     }
;

select_stmt:     select_explain select_compound select_order select_limit_offset { logDebugGrammar("SELECT_STMT")
                                                                    parsingQuery.ParsedSuccessfully = true }
;

select_explain:  /* empty */
        |   EXPLAIN {  parsingQuery.IsExplainOnly = true  }
;

select_order:   /* empty */
        |   ORDER BY sorting_list
;

select_limit_offset:    /* empty */
        |   select_limit
        |   select_limit select_offset
;

select_limit:   LIMIT expression { thisExpression := parsingStack.Pop()
                                   parsingQuery.Limit = thisExpression.(parser.Expression)
                                 }
;

select_offset:  OFFSET expression { thisExpression := parsingStack.Pop()
                                   parsingQuery.Offset = thisExpression.(parser.Expression)
                                 }
;

sorting_list:   sorting_single
        | sorting_single COMMA sorting_list
;

sorting_single: expression { thisExpression := parser.NewSortItem(parsingStack.Pop().(parser.Expression), true)
                            parsingQuery.Orderby = append(parsingQuery.Orderby, *thisExpression)
                           }
        |   expression ASC { thisExpression := parser.NewSortItem(parsingStack.Pop().(parser.Expression), true)
                            parsingQuery.Orderby = append(parsingQuery.Orderby, *thisExpression)
                           }
        |   expression DESC { thisExpression := parser.NewSortItem(parsingStack.Pop().(parser.Expression), false)
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
                                         parsingQuery.IsAggregateQuery = true 
                                         parsingQuery.Groupby = parsingStack.Pop().(parser.ExpressionList) }
;

select_group_having:    /*empty*/ { logDebugGrammar("SELECT GROUP HAVING - EMPTY") }
        |   select_group { logDebugGrammar("SELECT GROUP HAVING - SELECT GROUP") }
        |   select_group select_having { logDebugGrammar("SELECT GROUP HAVING - SELECT GROUP SELECT HAVING") }
;



select_having:  HAVING expression { parsingQuery.Having = parsingStack.Pop().(parser.Expression) }
;

select_where:   /* empty */ { logDebugGrammar("SELECT WHERE - EMPTY") }
        |   WHERE expression { logDebugGrammar("SELECT WHERE - EXPR")
                               where_part := parsingStack.Pop()
                               parsingQuery.Where = where_part.(parser.Expression) }
;

select_from:    /* empty */
        |   FROM data_source_list   { logDebugGrammar("SELECT_FROM") }
;

data_source_list:   data_source_single
        |   data_source_single COMMA data_source_list
;

data_source_single:  data_source { ds := parser.NewDataSource($1.s)
                                   parsingQuery.AddDataSource(ds) 
                                 }
        |   data_source data_source_over_list { ds := parser.NewDataSource($1.s)
                                          nextOver := parsingStack.Pop()
                                          for nextOver != nil {
                                            ds.AddOver(nextOver.(*parser.Over))
                                            nextOver = parsingStack.Pop()
                                          }
                                          parsingQuery.AddDataSource(ds)
                                        }
        |   data_source AS IDENTIFIER   { ds := parser.NewDataSourceWithAs($1.s, $3.s) 
                                          parsingQuery.AddDataSource(ds) 
                                        }
        |   data_source AS IDENTIFIER data_source_over_list { ds := parser.NewDataSourceWithAs($1.s, $3.s)
                                          nextOver := parsingStack.Pop()
                                          for nextOver != nil {
                                            ds.AddOver(nextOver.(*parser.Over))
                                            nextOver = parsingStack.Pop()
                                          }
                                          parsingQuery.AddDataSource(ds)
                                        }
;

data_source_over_list:  data_source_over
        |   data_source_over data_source_over_list
;

data_source_over:   OVER property AS IDENTIFIER {   prop := parsingStack.Pop().(*parser.Property)
                                                    over := parser.NewOver(prop, $4.s)
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
                            parsingQuery.Sel = thisExpression.(parser.Expression)
                          }
;

expression: expr { logDebugGrammar("EXPRESSION") }
    |   expr QUESTION expression COLON expression { logDebugGrammar("EXPRESSION - TERNARY")
                                                    elsee := parsingStack.Pop().(parser.Expression)
                                                    thenn := parsingStack.Pop().(parser.Expression)
                                                    iff := parsingStack.Pop().(parser.Expression)
                                                    thisExpr := parser.NewTernaryExpression(iff, thenn, elsee)
                                                    parsingStack.Push(thisExpr)
                                                  }
;

expr: expr PLUS expr {  logDebugGrammar("EXPR - PLUS")
                        right := parsingStack.Pop()
                        left := parsingStack.Pop()
                        thisExpression := parser.NewPlusExpression(left.(parser.Expression), right.(parser.Expression)) 
                        parsingStack.Push(thisExpression)
                     }
        |   expr MINUS expr {  logDebugGrammar("EXPR - MINUS")
                               right := parsingStack.Pop()
                               left := parsingStack.Pop()
                               thisExpression := parser.NewMinusExpression(left.(parser.Expression), right.(parser.Expression)) 
                               parsingStack.Push(thisExpression)
                            }
        |   expr MULT expr {  logDebugGrammar("EXPR - MULT")
                              right := parsingStack.Pop()
                              left := parsingStack.Pop()
                              thisExpression := parser.NewMultiplyExpression(left.(parser.Expression), right.(parser.Expression)) 
                              parsingStack.Push(thisExpression)
                           }
        |   expr DIV expr {  logDebugGrammar("EXPR - DIV")
                             right := parsingStack.Pop()
                             left := parsingStack.Pop()
                             thisExpression := parser.NewDivideExpression(left.(parser.Expression), right.(parser.Expression)) 
                             parsingStack.Push(thisExpression)
                          }
        |   expr AND expr {  logDebugGrammar("EXPR - AND")
                             right := parsingStack.Pop()
                             left := parsingStack.Pop()
                             thisExpression := parser.NewAndExpression(left.(parser.Expression), right.(parser.Expression)) 
                             parsingStack.Push(thisExpression)
                         }
        |   expr OR expr {  logDebugGrammar("EXPR - OR")
                            right := parsingStack.Pop()
                            left := parsingStack.Pop()
                            thisExpression := parser.NewOrExpression(left.(parser.Expression), right.(parser.Expression)) 
                            parsingStack.Push(thisExpression)
                         }
        |   expr EQ expr {  logDebugGrammar("EXPR - EQ")
                            right := parsingStack.Pop()
                            left := parsingStack.Pop()
                            thisExpression := parser.NewEqualsExpression(left.(parser.Expression), right.(parser.Expression)) 
                            parsingStack.Push(thisExpression)
                         }
        |   expr LT expr {  logDebugGrammar("EXPR - LT")
                            right := parsingStack.Pop()
                            left := parsingStack.Pop()
                            thisExpression := parser.NewLessThanExpression(left.(parser.Expression), right.(parser.Expression)) 
                            parsingStack.Push(thisExpression)
                         }
        |   expr LTE expr {  logDebugGrammar("EXPR - LTE")
                             right := parsingStack.Pop()
                             left := parsingStack.Pop()
                             thisExpression := parser.NewLessThanOrEqualExpression(left.(parser.Expression), right.(parser.Expression)) 
                             parsingStack.Push(thisExpression)
                         }
        |   expr GT expr {  logDebugGrammar("EXPR - GT")
                            right := parsingStack.Pop()
                            left := parsingStack.Pop()
                            thisExpression := parser.NewGreaterThanExpression(left.(parser.Expression), right.(parser.Expression)) 
                            parsingStack.Push(thisExpression)
                         }
        |   expr GTE expr {  logDebugGrammar("EXPR - GTE")
                             right := parsingStack.Pop()
                             left := parsingStack.Pop()
                             thisExpression := parser.NewGreaterThanOrEqualExpression(left.(parser.Expression), right.(parser.Expression)) 
                             parsingStack.Push(thisExpression)
                         }
        |   expr NE expr {  logDebugGrammar("EXPR - NE")
                            right := parsingStack.Pop()
                            left := parsingStack.Pop()
                            thisExpression := parser.NewNotEqualsExpression(left.(parser.Expression), right.(parser.Expression)) 
                            parsingStack.Push(thisExpression)
                         }
        |   expr MOD expr        
        |   prefix_expr
;


prefix_expr: NOT prefix_expr { logDebugGrammar("EXPR - NOT")
                               curr := parsingStack.Pop().(parser.Expression)
                               thisExpression := parser.NewNotExpression(curr)
                               parsingStack.Push(thisExpression)
                             }
    | suffix_expr
;

suffix_expr: atom { logDebugGrammar("SUFFIX_EXPR") }
;

atom: NULL { logDebugGrammar("NULL")
             thisExpression := parser.NewNull()
             parsingStack.Push(thisExpression)
           }
    | property {  }
    | property LBRACKET expression RBRACKET {     logDebugGrammar("ATOM - prop[]")
                                                  rightExpr := parsingStack.Pop().(parser.Expression)
                                                  leftProp := parsingStack.Pop().(*parser.Property)
                                                  thisExpression := parser.NewBracketMemberExpression(leftProp, rightExpr)
                                                  parsingStack.Push(thisExpression)
                                            }
    |   INT { thisExpression := parser.NewIntegerLiteral($1.n) 
                 parsingStack.Push(thisExpression) }
    |   MINUS INT { thisExpression := parser.NewIntegerLiteral(-$1.n)
                 parsingStack.Push(thisExpression) }
    |   REAL { thisExpression := parser.NewFloatLiteral($1.f) 
                 parsingStack.Push(thisExpression) }
    |   MINUS REAL { thisExpression := parser.NewFloatLiteral(-$1.f)
                 parsingStack.Push(thisExpression) }
    |   STRING { thisExpression := parser.NewStringLiteral($1.s) 
                 parsingStack.Push(thisExpression) }
    |   TRUE { thisExpression := parser.NewBoolLiteral(true) 
                 parsingStack.Push(thisExpression) }
    |   FALSE { thisExpression := parser.NewBoolLiteral(false) 
                 parsingStack.Push(thisExpression)}
    |   LBRACE named_expression_list RBRACE { logDebugGrammar("ATOM - {}")
                                            }
    |   LBRACKET expression_list RBRACKET { logDebugGrammar("ATOM - []")
                                            exp_list := parsingStack.Pop().(parser.ExpressionList)
                                            thisExpression := parser.NewArrayLiteral(exp_list)
                                            parsingStack.Push(thisExpression)
                                          }
    |   function_name LPAREN expression_list RPAREN { logDebugGrammar("FUNCTION - $1.s")
                                                      exp_list := parsingStack.Pop().(parser.ExpressionList)
                                                      function := parsingStack.Pop().(*parser.Function)
                                                      function.AddArguments(exp_list)
                                                      parsingStack.Push(function)
                                                    }
    |   LPAREN expression RPAREN
    |   LPAREN select_stmt RPAREN
;

expression_list:  expression { logDebugGrammar("EXPRESSION_LIST - EXPRESSION")
                               exp_list := make(parser.ExpressionList, 0)
                               exp_list = append(exp_list, parsingStack.Pop().(parser.Expression))
                               parsingStack.Push(exp_list)
                             }
        |   expression COMMA expression_list { logDebugGrammar("EXPRESSION_LIST - EXPRESSION COMMA EXPRESSION_LIST")
                                               rest := parsingStack.Pop().(parser.ExpressionList)
                                               last := parsingStack.Pop()
                                               new_list := make(parser.ExpressionList, 0)
                                               new_list = append(new_list, last.(parser.Expression))
                                               for _, v := range rest {
                                                new_list = append(new_list, v)
                                               }
                                               parsingStack.Push(new_list)
                                             }
;

named_expression_list:    named_expression_single 
        |   named_expression_single COMMA named_expression_list { last := parsingStack.Pop().(*parser.ObjectLiteral)
                                                                  rest := parsingStack.Pop().(*parser.ObjectLiteral)
                                                                  rest.AddAll(last)
                                                                  parsingStack.Push(rest)
                                                                }
;

named_expression_single:   STRING COLON expression { thisKey := $1.s
                                                     thisValue := parsingStack.Pop().(parser.Expression)
                                                     thisExpression := parser.NewObjectLiteral(parser.Object{thisKey: thisValue})
                                                     parsingStack.Push(thisExpression) 
                                                   }
;

property:   IDENTIFIER {
                         thisExpression := parser.NewProperty($1.s) 
                         parsingStack.Push(thisExpression) 
                       }
        | IDENTIFIER DOT property {
                                    thisValue := parsingStack.Pop().(*parser.Property)
                                    thisExpression := parser.NewProperty($1.s + "." + thisValue.Symbol)
                                    parsingStack.Push(thisExpression)
                                  }
;

function_name: MIN { 
                     parsingQuery.IsAggregateQuery = true
                     thisExpression := parser.NewFunction("min")
                     parsingStack.Push(thisExpression)
                   }
        |   MAX { 
                  parsingQuery.IsAggregateQuery = true
                  thisExpression := parser.NewFunction("max")
                  parsingStack.Push(thisExpression)
                }
        |   AVG { 
                  parsingQuery.IsAggregateQuery = true
                  thisExpression := parser.NewFunction("avg")
                  parsingStack.Push(thisExpression)
                }
        |   COUNT { 
                   parsingQuery.IsAggregateQuery = true
                   thisExpression := parser.NewFunction("count")
                   parsingStack.Push(thisExpression)
                  }
        |   SUM { 
                  parsingQuery.IsAggregateQuery = true
                  thisExpression := parser.NewFunction("sum")
                  parsingStack.Push(thisExpression)
                }
;
%%
