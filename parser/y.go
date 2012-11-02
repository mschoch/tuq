
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
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line unql.y:319


//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 88
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 204

var yyAct = []int{

	70, 69, 87, 78, 13, 66, 12, 51, 50, 42,
	9, 124, 107, 108, 85, 52, 53, 54, 55, 64,
	56, 57, 48, 58, 59, 60, 61, 62, 63, 72,
	52, 53, 54, 55, 64, 56, 4, 11, 58, 59,
	60, 61, 62, 63, 52, 53, 54, 55, 64, 86,
	89, 90, 91, 122, 123, 5, 2, 92, 93, 94,
	95, 96, 97, 98, 99, 100, 101, 102, 103, 104,
	130, 131, 74, 110, 36, 37, 38, 49, 116, 52,
	53, 54, 55, 64, 115, 73, 120, 58, 59, 60,
	61, 62, 63, 77, 40, 118, 82, 75, 20, 21,
	22, 19, 29, 14, 28, 45, 44, 25, 126, 9,
	9, 127, 125, 105, 121, 117, 109, 106, 128, 65,
	112, 132, 89, 111, 134, 135, 71, 133, 23, 24,
	26, 136, 81, 137, 31, 30, 32, 33, 34, 20,
	21, 22, 19, 29, 119, 28, 83, 15, 25, 68,
	67, 27, 18, 43, 20, 21, 22, 19, 29, 17,
	28, 16, 41, 25, 8, 80, 79, 129, 114, 23,
	24, 26, 113, 76, 39, 31, 30, 32, 33, 34,
	7, 35, 6, 88, 23, 24, 26, 84, 15, 47,
	31, 30, 32, 33, 34, 46, 10, 3, 1, 0,
	0, 0, 0, 15,
}
var yyPact = []int{

	-5, -1000, -1000, -1000, 8, 150, 48, 75, 135, 89,
	-8, 55, -53, -31, -1000, 150, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 143, 150, 116, 94, -1000,
	-1000, -1000, -1000, -1000, -1000, 95, 80, -1000, -1000, 73,
	125, -1000, 78, 139, -1000, -1000, -1000, -17, 150, 150,
	150, 150, 150, 150, 150, 150, 150, 150, 150, 150,
	150, 150, 150, 150, 150, -1000, 99, 105, -27, -24,
	104, 150, 112, 109, -1000, -1000, 63, 150, -1000, 103,
	77, -1000, 137, -1000, -1000, 150, -1000, -1000, 102, 21,
	-1000, -28, -1000, -1000, -1000, -1000, 33, -16, -2, -2,
	-2, -2, -2, -2, -1000, -1000, 143, 150, -1000, 150,
	107, -1000, -1000, -1000, 47, 49, -1000, 125, 120, -1000,
	-1000, 150, -1000, -1000, 150, -1000, -1000, -1000, -1000, -1000,
	150, 150, -1000, -1000, -1000, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 198, 56, 197, 0, 36, 196, 195, 2, 189,
	187, 183, 182, 181, 180, 174, 173, 172, 168, 1,
	167, 3, 166, 165, 164, 162, 4, 103, 161, 159,
	152, 5, 151, 150,
}
var yyR1 = []int{

	0, 1, 1, 3, 2, 6, 6, 7, 7, 7,
	9, 10, 8, 8, 11, 11, 11, 5, 5, 13,
	13, 13, 13, 12, 18, 17, 17, 17, 20, 16,
	16, 15, 15, 21, 21, 22, 22, 23, 14, 24,
	24, 24, 25, 25, 25, 25, 4, 4, 26, 26,
	26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
	26, 26, 27, 27, 28, 29, 29, 29, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 19, 19, 31,
	31, 33, 30, 32, 32, 32, 32, 32,
}
var yyR2 = []int{

	0, 1, 1, 4, 3, 0, 3, 0, 1, 2,
	2, 2, 1, 3, 1, 2, 2, 1, 3, 1,
	2, 1, 1, 4, 3, 0, 1, 2, 2, 0,
	2, 0, 2, 1, 3, 1, 3, 1, 2, 1,
	2, 2, 0, 1, 3, 2, 1, 5, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 3, 4, 3, 3, 1, 3, 1,
	3, 3, 1, 1, 1, 1, 1, 1,
}
var yyChk = []int{

	-1000, -1, -2, -3, -5, 60, -12, -14, -24, 15,
	-6, 29, -4, -26, -27, 53, -28, -29, -30, 7,
	4, 5, 6, 34, 35, 13, 36, -32, 10, 8,
	41, 40, 42, 43, 44, -13, 26, 27, 28, -15,
	19, -25, -4, 18, 17, 16, -7, -9, 30, 22,
	61, 38, 46, 47, 48, 49, 51, 52, 54, 55,
	56, 57, 58, 59, 50, -27, -31, -33, 6, -19,
	-4, 10, -4, -2, -5, 17, -16, 20, -21, -22,
	-23, 7, 18, 7, -10, 31, -4, -8, -11, -4,
	-4, -4, -26, -26, -26, -26, -26, -26, -26, -26,
	-26, -26, -26, -26, -26, 14, 12, 39, 37, 12,
	-19, 11, 11, -17, -18, 21, -4, 12, 18, 7,
	-4, 12, 32, 33, 39, -31, -4, -19, 11, -20,
	23, 22, -21, 7, -8, -4, -4, -19,
}
var yyDef = []int{

	0, -2, 1, 2, 5, 0, 17, 31, 42, 39,
	7, 0, 0, 46, 61, 0, 63, 64, 65, 66,
	67, 68, 69, 70, 71, 0, 0, 0, 0, 82,
	83, 84, 85, 86, 87, 0, 19, 21, 22, 29,
	0, 38, 43, 0, 40, 41, 4, 8, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 62, 0, 79, 0, 0,
	77, 0, 0, 0, 18, 20, 25, 0, 32, 33,
	35, 37, 0, 45, 9, 0, 10, 6, 12, 14,
	3, 0, 48, 49, 50, 51, 52, 53, 54, 55,
	56, 57, 58, 59, 60, 72, 0, 0, 73, 0,
	0, 75, 76, 23, 26, 0, 30, 0, 0, 44,
	11, 0, 15, 16, 0, 80, 81, 78, 74, 27,
	0, 0, 34, 36, 13, 47, 28, 24,
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
	                                                       ProcessPragma(left.(map[string]interface{}), right.(map[string]interface{}))
	                                                     }
	case 4:
		//line unql.y:46
		{ logDebugGrammar("SELECT_STMT")
	                                                                    parsingQuery.parsedSuccessfully = true }
	case 10:
		//line unql.y:59
		{ thisExpression := parsingStack.Pop()
	                                   parsingQuery.Limit = thisExpression
	                                 }
	case 11:
		//line unql.y:64
		{ thisExpression := parsingStack.Pop()
	                                   parsingQuery.Offset = thisExpression
	                                 }
	case 14:
		//line unql.y:73
		{ thisExpression := NewSortItem(parsingStack.Pop(), true)
	                            parsingQuery.Orderby = append(parsingQuery.Orderby, *thisExpression)
	                           }
	case 15:
		//line unql.y:76
		{ thisExpression := NewSortItem(parsingStack.Pop(), true)
	                            parsingQuery.Orderby = append(parsingQuery.Orderby, *thisExpression)
	                           }
	case 16:
		//line unql.y:79
		{ thisExpression := NewSortItem(parsingStack.Pop(), false)
	                            parsingQuery.Orderby = append(parsingQuery.Orderby, *thisExpression)
	                           }
	case 17:
		//line unql.y:84
		{ logDebugGrammar("SELECT_COMPOUND") }
	case 23:
		//line unql.y:94
		{ logDebugGrammar("SELECT_CORE") }
	case 24:
		//line unql.y:97
		{ logDebugGrammar("SELECT GROUP")
	                                         parsingQuery.isAggregateQuery = true 
	                                         parsingQuery.Groupby = parsingStack.Pop().(ExpressionList) }
	case 25:
		//line unql.y:102
		{ logDebugGrammar("SELECT GROUP HAVING - EMPTY") }
	case 26:
		//line unql.y:103
		{ logDebugGrammar("SELECT GROUP HAVING - SELECT GROUP") }
	case 27:
		//line unql.y:104
		{ logDebugGrammar("SELECT GROUP HAVING - SELECT GROUP SELECT HAVING") }
	case 29:
		//line unql.y:112
		{ logDebugGrammar("SELECT WHERE - EMPTY") }
	case 30:
		//line unql.y:113
		{ logDebugGrammar("SELECT WHERE - EXPR")
	                               where_part := parsingStack.Pop()
	                               parsingQuery.Where = where_part }
	case 32:
		//line unql.y:119
		{ logDebugGrammar("SELECT_FROM") }
	case 35:
		//line unql.y:126
		{ ds := NewDataSource(yyS[yypt-0].s)
	                                   parsingQuery.AddDataSource(ds) 
	                                 }
	case 36:
		//line unql.y:129
		{ ds := NewDataSourceWithAs(yyS[yypt-2].s, yyS[yypt-0].s) 
	                                          parsingQuery.AddDataSource(ds) 
	                                        }
	case 38:
		//line unql.y:136
		{ logDebugGrammar("SELECT_SELECT") }
	case 39:
		//line unql.y:139
		{ logDebugGrammar("SELECT_SELECT_HEAD") }
	case 42:
		//line unql.y:144
		{ logDebugGrammar("SELECT SELECT TAIL - EMPTY") }
	case 43:
		//line unql.y:145
		{ logDebugGrammar("SELECT SELECT TAIL - EXPR")
	                            thisExpression := parsingStack.Pop()
	                            parsingQuery.Sel = thisExpression
	                          }
	case 44:
		//line unql.y:149
		{ logDebugGrammar("SELECT SELECT TAIL - EXPR AS IDENTIFIER")
	                                          thisExpression := parsingStack.Pop()
	                                          parsingQuery.Sel = thisExpression
	                                          parsingQuery.SelAs = yyS[yypt-0].s
	                                        }
	case 45:
		//line unql.y:154
		{ logDebugGrammar("SELECT SELECT TAIL - AS IDENTIFIER")
	                               parsingQuery.SelAs = yyS[yypt-0].s
	                             }
	case 46:
		//line unql.y:159
		{ logDebugGrammar("EXPRESSION") }
	case 47:
		//line unql.y:160
		{ logDebugGrammar("EXPRESSION - TERNARY") }
	case 52:
		//line unql.y:167
		{  logDebugGrammar("EXPR - AND")
	                             right := parsingStack.Pop()
	                             left := parsingStack.Pop()
	                             thisExpression := NewAndExpression(left, right) 
	                             parsingStack.Push(*thisExpression)
	                         }
	case 53:
		//line unql.y:173
		{  logDebugGrammar("EXPR - OR")
	                            right := parsingStack.Pop()
	                            left := parsingStack.Pop()
	                            thisExpression := NewOrExpression(left, right) 
	                            parsingStack.Push(*thisExpression)
	                         }
	case 54:
		//line unql.y:179
		{  logDebugGrammar("EXPR - EQ")
	                            right := parsingStack.Pop()
	                            left := parsingStack.Pop()
	                            thisExpression := NewEqualsExpression(left, right) 
	                            parsingStack.Push(*thisExpression)
	                         }
	case 55:
		//line unql.y:185
		{  logDebugGrammar("EXPR - LT")
	                            right := parsingStack.Pop()
	                            left := parsingStack.Pop()
	                            thisExpression := NewLessThanExpression(left, right) 
	                            parsingStack.Push(*thisExpression)
	                         }
	case 56:
		//line unql.y:191
		{  logDebugGrammar("EXPR - LTE")
	                             right := parsingStack.Pop()
	                             left := parsingStack.Pop()
	                             thisExpression := NewLessThanOrEqualExpression(left, right) 
	                             parsingStack.Push(*thisExpression)
	                         }
	case 57:
		//line unql.y:197
		{  logDebugGrammar("EXPR - GT")
	                            right := parsingStack.Pop()
	                            left := parsingStack.Pop()
	                            thisExpression := NewGreaterThanExpression(left, right) 
	                            parsingStack.Push(*thisExpression)
	                         }
	case 58:
		//line unql.y:203
		{  logDebugGrammar("EXPR - GTE")
	                             right := parsingStack.Pop()
	                             left := parsingStack.Pop()
	                             thisExpression := NewGreaterThanOrEqualExpression(left, right) 
	                             parsingStack.Push(*thisExpression)
	                         }
	case 59:
		//line unql.y:209
		{  logDebugGrammar("EXPR - NE")
	                            right := parsingStack.Pop()
	                            left := parsingStack.Pop()
	                            thisExpression := NewNotEqualsExpression(left, right) 
	                            parsingStack.Push(*thisExpression)
	                         }
	case 64:
		//line unql.y:224
		{ logDebugGrammar("SUFFIX_EXPR") }
	case 65:
		//line unql.y:227
		{ thisExpression := NewProperty(yyS[yypt-0].s) 
	                 parsingStack.Push(*thisExpression) }
	case 66:
		//line unql.y:229
		{ thisExpression := NewProperty(yyS[yypt-0].s) 
	                 parsingStack.Push(*thisExpression) }
	case 67:
		//line unql.y:231
		{ thisExpression := NewIntegerLiteral(yyS[yypt-0].n) 
	                 parsingStack.Push(*thisExpression) }
	case 68:
		//line unql.y:233
		{ thisExpression := NewFloatLiteral(yyS[yypt-0].f) 
	                 parsingStack.Push(*thisExpression) }
	case 69:
		//line unql.y:235
		{ thisExpression := NewStringLiteral(yyS[yypt-0].s) 
	                 parsingStack.Push(*thisExpression) }
	case 70:
		//line unql.y:237
		{ thisExpression := NewBoolLiteral(true) 
	                 parsingStack.Push(*thisExpression) }
	case 71:
		//line unql.y:239
		{ thisExpression := NewBoolLiteral(false) 
	                 parsingStack.Push(*thisExpression)}
	case 72:
		//line unql.y:241
		{ logDebugGrammar("ATOM - {}")
	                                            }
	case 73:
		//line unql.y:243
		{ logDebugGrammar("ATOM - []")
	                                            exp_list := parsingStack.Pop().(ExpressionList)
	                                            thisExpression := NewArrayLiteral(exp_list)
	                                            parsingStack.Push(*thisExpression)
	                                          }
	case 74:
		//line unql.y:248
		{ logDebugGrammar("FUNCTION - $1.s")
	                                                      exp_list := parsingStack.Pop().(ExpressionList)
	                                                      function := parsingStack.Pop().(Function)
	                                                      function.AddArguments(exp_list)
	                                                      parsingStack.Push(function)
	                                                    }
	case 77:
		//line unql.y:258
		{ logDebugGrammar("EXPRESSION_LIST - EXPRESSION")
	                               exp_list := make(ExpressionList, 0)
	                               exp_list = append(exp_list, parsingStack.Pop())
	                               parsingStack.Push(exp_list)
	                             }
	case 78:
		//line unql.y:263
		{ logDebugGrammar("EXPRESSION_LIST - EXPRESSION COMMA EXPRESSION_LIST")
	                                               rest := parsingStack.Pop().(ExpressionList)
	                                               last := parsingStack.Pop()
	                                               new_list := make(ExpressionList, 0)
	                                               new_list = append(new_list, last)
	                                               for _, v := range rest {
	                                                new_list = append(new_list, v)
	                                               }
	                                               parsingStack.Push(new_list)
	                                             }
	case 80:
		//line unql.y:276
		{ last := parsingStack.Pop().(ObjectLiteral)
	                                                                  rest := parsingStack.Pop().(ObjectLiteral)
	                                                                  rest.AddAll(last)
	                                                                  parsingStack.Push(rest)
	                                                                }
	case 81:
		//line unql.y:283
		{ thisKey := yyS[yypt-2].s
	                                                     thisValue := parsingStack.Pop() 
	                                                     thisExpression := NewObjectLiteral(Object{thisKey: thisValue})
	                                                     parsingStack.Push(*thisExpression) 
	                                                   }
	case 83:
		//line unql.y:293
		{ 
	                     parsingQuery.isAggregateQuery = true
	                     thisExpression := NewFunction("min")
	                     parsingStack.Push(*thisExpression)
	                   }
	case 84:
		//line unql.y:298
		{ 
	                  parsingQuery.isAggregateQuery = true
	                  thisExpression := NewFunction("max")
	                  parsingStack.Push(*thisExpression)
	                }
	case 85:
		//line unql.y:303
		{ 
	                  parsingQuery.isAggregateQuery = true
	                  thisExpression := NewFunction("avg")
	                  parsingStack.Push(*thisExpression)
	                }
	case 86:
		//line unql.y:308
		{ 
	                   parsingQuery.isAggregateQuery = true
	                   thisExpression := NewFunction("count")
	                   parsingStack.Push(*thisExpression)
	                  }
	case 87:
		//line unql.y:313
		{ 
	                  parsingQuery.isAggregateQuery = true
	                  thisExpression := NewFunction("sum")
	                  parsingStack.Push(*thisExpression)
	                }
	}
	goto yystack /* stack new state and value */
}
