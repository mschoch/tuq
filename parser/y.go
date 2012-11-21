
//line unql.y:2
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

//line unql.y:15
type yySymType struct {
	yys int 
s string 
n int
f float64}

const INT = 57346
const REAL = 57347
const STRING = 57348
const IDENTIFIER = 57349
const PROPERTY = 57350
const NEWLINE = 57351
const LPAREN = 57352
const RPAREN = 57353
const COMMA = 57354
const LBRACE = 57355
const RBRACE = 57356
const SELECT = 57357
const DISTINCT = 57358
const ALL = 57359
const AS = 57360
const FROM = 57361
const WHERE = 57362
const GROUP = 57363
const BY = 57364
const HAVING = 57365
const FLATTEN = 57366
const EACH = 57367
const UNION = 57368
const INTERSECT = 57369
const EXCEPT = 57370
const ORDER = 57371
const LIMIT = 57372
const OFFSET = 57373
const ASC = 57374
const DESC = 57375
const TRUE = 57376
const FALSE = 57377
const LBRACKET = 57378
const RBRACKET = 57379
const QUESTION = 57380
const COLON = 57381
const MAX = 57382
const MIN = 57383
const AVG = 57384
const COUNT = 57385
const SUM = 57386
const DOT = 57387
const PLUS = 57388
const MINUS = 57389
const MULT = 57390
const DIV = 57391
const MOD = 57392
const AND = 57393
const OR = 57394
const NOT = 57395
const EQ = 57396
const LT = 57397
const LTE = 57398
const GT = 57399
const GTE = 57400
const NE = 57401
const PRAGMA = 57402
const ASSIGN = 57403
const EXPLAIN = 57404

var yyToknames = []string{
	"INT",
	"REAL",
	"STRING",
	"IDENTIFIER",
	"PROPERTY",
	"NEWLINE",
	"LPAREN",
	"RPAREN",
	"COMMA",
	"LBRACE",
	"RBRACE",
	"SELECT",
	"DISTINCT",
	"ALL",
	"AS",
	"FROM",
	"WHERE",
	"GROUP",
	"BY",
	"HAVING",
	"FLATTEN",
	"EACH",
	"UNION",
	"INTERSECT",
	"EXCEPT",
	"ORDER",
	"LIMIT",
	"OFFSET",
	"ASC",
	"DESC",
	"TRUE",
	"FALSE",
	"LBRACKET",
	"RBRACKET",
	"QUESTION",
	"COLON",
	"MAX",
	"MIN",
	"AVG",
	"COUNT",
	"SUM",
	"DOT",
	"PLUS",
	"MINUS",
	"MULT",
	"DIV",
	"MOD",
	"AND",
	"OR",
	"NOT",
	"EQ",
	"LT",
	"LTE",
	"GT",
	"GTE",
	"NE",
	"PRAGMA",
	"ASSIGN",
	"EXPLAIN",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line unql.y:343


//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 90
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 229

var yyAct = []int{

	68, 67, 112, 64, 80, 5, 12, 6, 48, 122,
	103, 44, 104, 129, 130, 20, 21, 22, 19, 29,
	110, 28, 74, 49, 25, 7, 38, 39, 40, 70,
	13, 50, 51, 52, 53, 62, 54, 55, 36, 56,
	57, 58, 59, 60, 61, 23, 24, 26, 2, 86,
	87, 31, 30, 32, 33, 34, 50, 51, 52, 53,
	62, 132, 133, 76, 15, 75, 117, 79, 42, 120,
	84, 106, 77, 6, 11, 111, 114, 71, 47, 46,
	118, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 14, 101, 128, 119, 105, 102,
	126, 135, 108, 107, 124, 69, 123, 125, 83, 121,
	63, 127, 85, 66, 65, 27, 18, 17, 16, 43,
	10, 82, 81, 136, 134, 131, 116, 115, 78, 114,
	41, 137, 9, 138, 37, 139, 50, 51, 52, 53,
	62, 54, 8, 113, 56, 57, 58, 59, 60, 61,
	50, 51, 52, 53, 62, 109, 73, 72, 56, 57,
	58, 59, 60, 61, 20, 21, 22, 19, 29, 35,
	28, 4, 3, 25, 1, 0, 0, 0, 45, 20,
	21, 22, 19, 29, 0, 28, 0, 0, 25, 0,
	0, 0, 0, 0, 23, 24, 26, 0, 0, 0,
	31, 30, 32, 33, 34, 0, 0, 0, 0, 23,
	24, 26, 0, 15, 0, 31, 30, 32, 33, 34,
	0, 0, 0, 0, 0, 0, 0, 0, 15,
}
var yyPact = []int{

	-55, -1000, -1000, -1000, 59, 175, -1000, 9, 0, 49,
	160, 62, -53, -15, -1000, 175, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 107, 175, 95, 11, -1000,
	-1000, -1000, -1000, -1000, -1000, -8, 43, 59, 55, -1000,
	-1000, 47, 101, -1000, 52, 105, -1000, -1000, 175, 175,
	175, 175, 175, 175, 175, 175, 175, 175, 175, 175,
	175, 175, 175, -1000, 81, 87, -29, -25, 86, 175,
	92, 91, -1000, -11, 175, 175, -1000, -1000, 45, 175,
	-1000, 85, 51, -1000, 102, -1000, -1000, -30, -1000, -1000,
	-1000, -1000, 104, 90, 10, 10, 10, 10, 10, 10,
	-1000, -1000, 107, 175, -1000, 175, 89, -1000, -1000, -1000,
	175, -1000, -1000, 84, -19, -1000, 38, 40, -1000, 101,
	94, -1000, 175, -1000, -1000, -1000, -1000, -1000, 175, -1000,
	-1000, -1000, 175, 175, -1000, -1000, -1000, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 174, 48, 172, 0, 171, 25, 169, 157, 2,
	156, 155, 143, 142, 134, 132, 130, 128, 127, 126,
	1, 125, 4, 122, 121, 120, 119, 30, 94, 118,
	117, 116, 3, 115, 114,
}
var yyR1 = []int{

	0, 1, 1, 3, 2, 5, 5, 7, 7, 8,
	8, 8, 10, 11, 9, 9, 12, 12, 12, 6,
	6, 14, 14, 14, 14, 13, 19, 18, 18, 18,
	21, 17, 17, 16, 16, 22, 22, 23, 23, 24,
	15, 25, 25, 25, 26, 26, 26, 26, 4, 4,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 28, 28, 29, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 20,
	20, 32, 32, 34, 31, 33, 33, 33, 33, 33,
}
var yyR2 = []int{

	0, 1, 1, 4, 4, 0, 1, 0, 3, 0,
	1, 2, 2, 2, 1, 3, 1, 2, 2, 1,
	3, 1, 2, 1, 1, 4, 3, 0, 1, 2,
	2, 0, 2, 0, 2, 1, 3, 1, 3, 1,
	2, 1, 2, 2, 0, 1, 3, 2, 1, 5,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 3, 3, 4, 3, 3, 1,
	3, 1, 3, 3, 1, 1, 1, 1, 1, 1,
}
var yyChk = []int{

	-1000, -1, -2, -3, -5, 60, 62, -6, -13, -15,
	-25, 15, -4, -27, -28, 53, -29, -30, -31, 7,
	4, 5, 6, 34, 35, 13, 36, -33, 10, 8,
	41, 40, 42, 43, 44, -7, 29, -14, 26, 27,
	28, -16, 19, -26, -4, 18, 17, 16, 61, 38,
	46, 47, 48, 49, 51, 52, 54, 55, 56, 57,
	58, 59, 50, -28, -32, -34, 6, -20, -4, 10,
	-4, -2, -8, -10, 30, 22, -6, 17, -17, 20,
	-22, -23, -24, 7, 18, 7, -4, -4, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, 14, 12, 39, 37, 12, -20, 11, 11, -11,
	31, -4, -9, -12, -4, -18, -19, 21, -4, 12,
	18, 7, 39, -32, -4, -20, 11, -4, 12, 32,
	33, -21, 23, 22, -22, 7, -4, -9, -4, -20,
}
var yyDef = []int{

	5, -2, 1, 2, 0, 0, 6, 7, 19, 33,
	44, 41, 0, 48, 63, 0, 65, 66, 67, 68,
	69, 70, 71, 72, 73, 0, 0, 0, 5, 84,
	85, 86, 87, 88, 89, 9, 0, 0, 21, 23,
	24, 31, 0, 40, 45, 0, 42, 43, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 0, 81, 0, 0, 79, 0,
	0, 0, 4, 10, 0, 0, 20, 22, 27, 0,
	34, 35, 37, 39, 0, 47, 3, 0, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 74, 0, 0, 75, 0, 0, 77, 78, 11,
	0, 12, 8, 14, 16, 25, 28, 0, 32, 0,
	0, 46, 0, 82, 83, 80, 76, 13, 0, 17,
	18, 29, 0, 0, 36, 38, 49, 15, 30, 26,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c > 0 && c <= len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return fmt.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return fmt.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		fmt.Printf("lex %U %s\n", uint(char), yyTokname(c))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		fmt.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				fmt.Printf("%s", yyStatname(yystate))
				fmt.Printf("saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					fmt.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				fmt.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		fmt.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line unql.y:35
		{ logDebugGrammar("INPUT") }
	case 3:
		//line unql.y:39
		{ logDebugGrammar("PRAGMA: %v", yyS[yypt-3])
	                                                       right := parsingStack.Pop()
	                                                       left := parsingStack.Pop()
	                                                       ProcessPragma(left.(Expression), right.(Expression))
	                                                     }
	case 4:
		//line unql.y:46
		{ logDebugGrammar("SELECT_STMT")
	                                                                    parsingQuery.parsedSuccessfully = true }
	case 6:
		//line unql.y:51
		{  parsingQuery.isExplainOnly = true  }
	case 12:
		//line unql.y:63
		{ thisExpression := parsingStack.Pop()
	                                   parsingQuery.Limit = thisExpression.(Expression)
	                                 }
	case 13:
		//line unql.y:68
		{ thisExpression := parsingStack.Pop()
	                                   parsingQuery.Offset = thisExpression.(Expression)
	                                 }
	case 16:
		//line unql.y:77
		{ thisExpression := NewSortItem(parsingStack.Pop().(Expression), true)
	                            parsingQuery.Orderby = append(parsingQuery.Orderby, *thisExpression)
	                           }
	case 17:
		//line unql.y:80
		{ thisExpression := NewSortItem(parsingStack.Pop().(Expression), true)
	                            parsingQuery.Orderby = append(parsingQuery.Orderby, *thisExpression)
	                           }
	case 18:
		//line unql.y:83
		{ thisExpression := NewSortItem(parsingStack.Pop().(Expression), false)
	                            parsingQuery.Orderby = append(parsingQuery.Orderby, *thisExpression)
	                           }
	case 19:
		//line unql.y:88
		{ logDebugGrammar("SELECT_COMPOUND") }
	case 25:
		//line unql.y:98
		{ logDebugGrammar("SELECT_CORE") }
	case 26:
		//line unql.y:101
		{ logDebugGrammar("SELECT GROUP")
	                                         parsingQuery.isAggregateQuery = true 
	                                         parsingQuery.Groupby = parsingStack.Pop().(ExpressionList) }
	case 27:
		//line unql.y:106
		{ logDebugGrammar("SELECT GROUP HAVING - EMPTY") }
	case 28:
		//line unql.y:107
		{ logDebugGrammar("SELECT GROUP HAVING - SELECT GROUP") }
	case 29:
		//line unql.y:108
		{ logDebugGrammar("SELECT GROUP HAVING - SELECT GROUP SELECT HAVING") }
	case 30:
		//line unql.y:113
		{ parsingQuery.Having = parsingStack.Pop().(Expression) }
	case 31:
		//line unql.y:116
		{ logDebugGrammar("SELECT WHERE - EMPTY") }
	case 32:
		//line unql.y:117
		{ logDebugGrammar("SELECT WHERE - EXPR")
	                               where_part := parsingStack.Pop()
	                               parsingQuery.Where = where_part.(Expression) }
	case 34:
		//line unql.y:123
		{ logDebugGrammar("SELECT_FROM") }
	case 37:
		//line unql.y:130
		{ ds := NewDataSource(yyS[yypt-0].s)
	                                   parsingQuery.AddDataSource(ds) 
	                                 }
	case 38:
		//line unql.y:133
		{ ds := NewDataSourceWithAs(yyS[yypt-2].s, yyS[yypt-0].s) 
	                                          parsingQuery.AddDataSource(ds) 
	                                        }
	case 40:
		//line unql.y:140
		{ logDebugGrammar("SELECT_SELECT") }
	case 41:
		//line unql.y:143
		{ logDebugGrammar("SELECT_SELECT_HEAD") }
	case 44:
		//line unql.y:148
		{ logDebugGrammar("SELECT SELECT TAIL - EMPTY") }
	case 45:
		//line unql.y:149
		{ logDebugGrammar("SELECT SELECT TAIL - EXPR")
	                            thisExpression := parsingStack.Pop()
	                            parsingQuery.Sel = thisExpression.(Expression)
	                          }
	case 46:
		//line unql.y:153
		{ logDebugGrammar("SELECT SELECT TAIL - EXPR AS IDENTIFIER")
	                                          thisExpression := parsingStack.Pop()
	                                          parsingQuery.Sel = thisExpression.(Expression)
	                                          parsingQuery.SelAs = yyS[yypt-0].s
	                                        }
	case 47:
		//line unql.y:158
		{ logDebugGrammar("SELECT SELECT TAIL - AS IDENTIFIER")
	                               parsingQuery.SelAs = yyS[yypt-0].s
	                             }
	case 48:
		//line unql.y:163
		{ logDebugGrammar("EXPRESSION") }
	case 49:
		//line unql.y:164
		{ logDebugGrammar("EXPRESSION - TERNARY") }
	case 50:
		//line unql.y:167
		{  logDebugGrammar("EXPR - PLUS")
	                        right := parsingStack.Pop()
	                        left := parsingStack.Pop()
	                        thisExpression := NewPlusExpression(left.(Expression), right.(Expression)) 
	                        parsingStack.Push(thisExpression)
	                     }
	case 51:
		//line unql.y:173
		{  logDebugGrammar("EXPR - MINUS")
	                               right := parsingStack.Pop()
	                               left := parsingStack.Pop()
	                               thisExpression := NewMinusExpression(left.(Expression), right.(Expression)) 
	                               parsingStack.Push(thisExpression)
	                            }
	case 52:
		//line unql.y:179
		{  logDebugGrammar("EXPR - MULT")
	                              right := parsingStack.Pop()
	                              left := parsingStack.Pop()
	                              thisExpression := NewMultiplyExpression(left.(Expression), right.(Expression)) 
	                              parsingStack.Push(thisExpression)
	                           }
	case 53:
		//line unql.y:185
		{  logDebugGrammar("EXPR - DIV")
	                             right := parsingStack.Pop()
	                             left := parsingStack.Pop()
	                             thisExpression := NewDivideExpression(left.(Expression), right.(Expression)) 
	                             parsingStack.Push(thisExpression)
	                          }
	case 54:
		//line unql.y:191
		{  logDebugGrammar("EXPR - AND")
	                             right := parsingStack.Pop()
	                             left := parsingStack.Pop()
	                             thisExpression := NewAndExpression(left.(Expression), right.(Expression)) 
	                             parsingStack.Push(thisExpression)
	                         }
	case 55:
		//line unql.y:197
		{  logDebugGrammar("EXPR - OR")
	                            right := parsingStack.Pop()
	                            left := parsingStack.Pop()
	                            thisExpression := NewOrExpression(left.(Expression), right.(Expression)) 
	                            parsingStack.Push(thisExpression)
	                         }
	case 56:
		//line unql.y:203
		{  logDebugGrammar("EXPR - EQ")
	                            right := parsingStack.Pop()
	                            left := parsingStack.Pop()
	                            thisExpression := NewEqualsExpression(left.(Expression), right.(Expression)) 
	                            parsingStack.Push(thisExpression)
	                         }
	case 57:
		//line unql.y:209
		{  logDebugGrammar("EXPR - LT")
	                            right := parsingStack.Pop()
	                            left := parsingStack.Pop()
	                            thisExpression := NewLessThanExpression(left.(Expression), right.(Expression)) 
	                            parsingStack.Push(thisExpression)
	                         }
	case 58:
		//line unql.y:215
		{  logDebugGrammar("EXPR - LTE")
	                             right := parsingStack.Pop()
	                             left := parsingStack.Pop()
	                             thisExpression := NewLessThanOrEqualExpression(left.(Expression), right.(Expression)) 
	                             parsingStack.Push(thisExpression)
	                         }
	case 59:
		//line unql.y:221
		{  logDebugGrammar("EXPR - GT")
	                            right := parsingStack.Pop()
	                            left := parsingStack.Pop()
	                            thisExpression := NewGreaterThanExpression(left.(Expression), right.(Expression)) 
	                            parsingStack.Push(thisExpression)
	                         }
	case 60:
		//line unql.y:227
		{  logDebugGrammar("EXPR - GTE")
	                             right := parsingStack.Pop()
	                             left := parsingStack.Pop()
	                             thisExpression := NewGreaterThanOrEqualExpression(left.(Expression), right.(Expression)) 
	                             parsingStack.Push(thisExpression)
	                         }
	case 61:
		//line unql.y:233
		{  logDebugGrammar("EXPR - NE")
	                            right := parsingStack.Pop()
	                            left := parsingStack.Pop()
	                            thisExpression := NewNotEqualsExpression(left.(Expression), right.(Expression)) 
	                            parsingStack.Push(thisExpression)
	                         }
	case 66:
		//line unql.y:248
		{ logDebugGrammar("SUFFIX_EXPR") }
	case 67:
		//line unql.y:251
		{ thisExpression := NewProperty(yyS[yypt-0].s) 
	                 parsingStack.Push(thisExpression) }
	case 68:
		//line unql.y:253
		{ thisExpression := NewProperty(yyS[yypt-0].s) 
	                 parsingStack.Push(thisExpression) }
	case 69:
		//line unql.y:255
		{ thisExpression := NewIntegerLiteral(yyS[yypt-0].n) 
	                 parsingStack.Push(thisExpression) }
	case 70:
		//line unql.y:257
		{ thisExpression := NewFloatLiteral(yyS[yypt-0].f) 
	                 parsingStack.Push(thisExpression) }
	case 71:
		//line unql.y:259
		{ thisExpression := NewStringLiteral(yyS[yypt-0].s) 
	                 parsingStack.Push(thisExpression) }
	case 72:
		//line unql.y:261
		{ thisExpression := NewBoolLiteral(true) 
	                 parsingStack.Push(thisExpression) }
	case 73:
		//line unql.y:263
		{ thisExpression := NewBoolLiteral(false) 
	                 parsingStack.Push(thisExpression)}
	case 74:
		//line unql.y:265
		{ logDebugGrammar("ATOM - {}")
	                                            }
	case 75:
		//line unql.y:267
		{ logDebugGrammar("ATOM - []")
	                                            exp_list := parsingStack.Pop().(ExpressionList)
	                                            thisExpression := NewArrayLiteral(exp_list)
	                                            parsingStack.Push(thisExpression)
	                                          }
	case 76:
		//line unql.y:272
		{ logDebugGrammar("FUNCTION - $1.s")
	                                                      exp_list := parsingStack.Pop().(ExpressionList)
	                                                      function := parsingStack.Pop().(*Function)
	                                                      function.AddArguments(exp_list)
	                                                      parsingStack.Push(function)
	                                                    }
	case 79:
		//line unql.y:282
		{ logDebugGrammar("EXPRESSION_LIST - EXPRESSION")
	                               exp_list := make(ExpressionList, 0)
	                               exp_list = append(exp_list, parsingStack.Pop().(Expression))
	                               parsingStack.Push(exp_list)
	                             }
	case 80:
		//line unql.y:287
		{ logDebugGrammar("EXPRESSION_LIST - EXPRESSION COMMA EXPRESSION_LIST")
	                                               rest := parsingStack.Pop().(ExpressionList)
	                                               last := parsingStack.Pop()
	                                               new_list := make(ExpressionList, 0)
	                                               new_list = append(new_list, last.(Expression))
	                                               for _, v := range rest {
	                                                new_list = append(new_list, v)
	                                               }
	                                               parsingStack.Push(new_list)
	                                             }
	case 82:
		//line unql.y:300
		{ last := parsingStack.Pop().(*ObjectLiteral)
	                                                                  rest := parsingStack.Pop().(*ObjectLiteral)
	                                                                  rest.AddAll(last)
	                                                                  parsingStack.Push(rest)
	                                                                }
	case 83:
		//line unql.y:307
		{ thisKey := yyS[yypt-2].s
	                                                     thisValue := parsingStack.Pop().(Expression)
	                                                     thisExpression := NewObjectLiteral(Object{thisKey: thisValue})
	                                                     parsingStack.Push(thisExpression) 
	                                                   }
	case 85:
		//line unql.y:317
		{ 
	                     parsingQuery.isAggregateQuery = true
	                     thisExpression := NewFunction("min")
	                     parsingStack.Push(thisExpression)
	                   }
	case 86:
		//line unql.y:322
		{ 
	                  parsingQuery.isAggregateQuery = true
	                  thisExpression := NewFunction("max")
	                  parsingStack.Push(thisExpression)
	                }
	case 87:
		//line unql.y:327
		{ 
	                  parsingQuery.isAggregateQuery = true
	                  thisExpression := NewFunction("avg")
	                  parsingStack.Push(thisExpression)
	                }
	case 88:
		//line unql.y:332
		{ 
	                   parsingQuery.isAggregateQuery = true
	                   thisExpression := NewFunction("count")
	                   parsingStack.Push(thisExpression)
	                  }
	case 89:
		//line unql.y:337
		{ 
	                  parsingQuery.isAggregateQuery = true
	                  thisExpression := NewFunction("sum")
	                  parsingStack.Push(thisExpression)
	                }
	}
	goto yystack /* stack new state and value */
}
