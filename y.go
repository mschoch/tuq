
//line unql.y:2
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
const DOT = 57386
const PLUS = 57387
const MINUS = 57388
const MULT = 57389
const DIV = 57390
const MOD = 57391
const AND = 57392
const OR = 57393
const NOT = 57394
const EQ = 57395
const LT = 57396
const LTE = 57397
const GT = 57398
const GTE = 57399
const NE = 57400
const PRAGMA = 57401
const ASSIGN = 57402
const DEBUG = 57403

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
	"DEBUG",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line unql.y:358


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

const yyLast = 205

var yyAct = []int{

	80, 79, 87, 76, 55, 23, 61, 13, 50, 21,
	9, 129, 114, 62, 63, 64, 65, 74, 66, 67,
	2, 68, 69, 70, 71, 72, 73, 62, 63, 64,
	65, 74, 62, 63, 64, 65, 74, 115, 85, 82,
	68, 69, 70, 71, 72, 73, 122, 123, 48, 86,
	89, 90, 11, 126, 5, 94, 15, 16, 17, 83,
	125, 24, 98, 49, 93, 54, 19, 96, 99, 100,
	101, 102, 103, 104, 105, 106, 107, 108, 109, 110,
	111, 59, 52, 117, 45, 44, 120, 75, 9, 112,
	121, 4, 116, 113, 95, 133, 119, 118, 81, 128,
	127, 30, 31, 32, 29, 39, 51, 38, 58, 97,
	35, 60, 9, 78, 77, 131, 37, 130, 132, 28,
	27, 26, 89, 20, 134, 8, 135, 57, 136, 56,
	137, 33, 34, 36, 124, 92, 91, 41, 40, 42,
	43, 30, 31, 32, 29, 39, 53, 38, 18, 25,
	35, 7, 14, 6, 88, 22, 30, 31, 32, 29,
	39, 84, 38, 47, 46, 35, 10, 12, 3, 1,
	0, 33, 34, 36, 0, 0, 0, 41, 40, 42,
	43, 0, 0, 0, 0, 0, 33, 34, 36, 25,
	0, 0, 41, 40, 42, 43, 0, 0, 0, 0,
	0, 0, 0, 0, 25,
}
var yyPact = []int{

	-5, -1000, -1000, -1000, 23, -54, 30, 47, 137, 68,
	18, 41, -52, -1000, 73, 65, -1000, -1000, 45, 101,
	-1000, 63, 104, -32, -1000, 152, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 107, 152, 88, 97, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 7, 152, 152,
	152, -1000, -1000, 43, 152, -1000, 82, 49, -1000, 102,
	-1000, 152, 152, 152, 152, 152, 152, 152, 152, 152,
	152, 152, 152, 152, 152, -1000, 75, 81, -27, 0,
	80, 152, 86, 85, -1000, 152, -1000, -1000, 78, 14,
	-1000, -1000, 37, 31, -1000, 101, 92, -1000, -28, -1000,
	-1000, -1000, -1000, -13, -13, -18, -18, -18, -18, -18,
	-18, -1000, -1000, 107, 152, -1000, 152, 84, -1000, -1000,
	-1000, 152, -1000, -1000, -1000, 152, 152, -1000, -1000, 152,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 169, 20, 168, 167, 0, 91, 166, 164, 2,
	163, 161, 154, 153, 152, 151, 148, 146, 136, 135,
	1, 134, 4, 129, 127, 125, 123, 5, 61, 121,
	120, 119, 3, 116, 114,
}
var yyR1 = []int{

	0, 1, 1, 3, 4, 2, 7, 7, 8, 8,
	8, 10, 11, 9, 9, 12, 12, 12, 6, 6,
	14, 14, 14, 14, 13, 19, 18, 18, 18, 21,
	17, 17, 16, 16, 22, 22, 23, 23, 24, 15,
	25, 25, 25, 26, 26, 26, 26, 5, 5, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 28, 28, 29, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 20, 20,
	32, 32, 34, 31, 33, 33, 33, 33,
}
var yyR2 = []int{

	0, 1, 1, 4, 1, 3, 0, 3, 0, 1,
	2, 2, 2, 1, 3, 1, 2, 2, 1, 3,
	1, 2, 1, 1, 4, 3, 0, 1, 2, 2,
	0, 2, 0, 2, 1, 3, 1, 3, 1, 2,
	1, 2, 2, 0, 1, 3, 2, 1, 5, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 1, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 3, 4, 3, 3, 1, 3,
	1, 3, 3, 1, 1, 1, 1, 1,
}
var yyChk = []int{

	-1000, -1, -2, -3, -6, 59, -13, -15, -25, 15,
	-7, 29, -4, 61, -14, 26, 27, 28, -16, 19,
	-26, -5, 18, -27, -28, 52, -29, -30, -31, 7,
	4, 5, 6, 34, 35, 13, 36, -33, 10, 8,
	41, 40, 42, 43, 17, 16, -8, -10, 30, 22,
	60, -6, 17, -17, 20, -22, -23, -24, 7, 18,
	7, 38, 45, 46, 47, 48, 50, 51, 53, 54,
	55, 56, 57, 58, 49, -28, -32, -34, 6, -20,
	-5, 10, -5, -2, -11, 31, -5, -9, -12, -5,
	-5, -18, -19, 21, -5, 12, 18, 7, -5, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, 14, 12, 39, 37, 12, -20, 11, 11,
	-5, 12, 32, 33, -21, 23, 22, -22, 7, 39,
	-32, -5, -20, 11, -9, -5, -20, -5,
}
var yyDef = []int{

	0, -2, 1, 2, 6, 0, 18, 32, 43, 40,
	8, 0, 0, 4, 0, 20, 22, 23, 30, 0,
	39, 44, 0, 47, 62, 0, 64, 65, 66, 67,
	68, 69, 70, 71, 72, 0, 0, 0, 0, 83,
	84, 85, 86, 87, 41, 42, 5, 9, 0, 0,
	0, 19, 21, 26, 0, 33, 34, 36, 38, 0,
	46, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 63, 0, 80, 0, 0,
	78, 0, 0, 0, 10, 0, 11, 7, 13, 15,
	3, 24, 27, 0, 31, 0, 0, 45, 0, 49,
	50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
	60, 61, 73, 0, 0, 74, 0, 0, 76, 77,
	12, 0, 16, 17, 28, 0, 0, 35, 37, 0,
	81, 82, 79, 75, 14, 29, 25, 48,
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
		//line unql.y:34
		{ logDebugGrammar("INPUT") }
	case 3:
		//line unql.y:38
		{ logDebugGrammar("PRAGMA: %v", yyS[yypt-3])
	                                                       right := parsingStack.Pop()
	                                                       left := parsingStack.Pop()
	                                                       ProcessPragma(left.(map[string]interface{}), right.(map[string]interface{}))
	                                                     }
	case 4:
		//line unql.y:45
		{ curr := make(map[string]interface{})
	                        curr["pragma"] = "debug"
	                        parsingStack.Push(curr) }
	case 5:
		//line unql.y:50
		{ logDebugGrammar("SELECT_STMT")
	                                                                    parsingQuery.parsedSuccessfully = true }
	case 11:
		//line unql.y:63
		{ curr := make(map[string]interface{})
	                                   curr["op"] = "limit"
	                                   curr["expression"] = parsingStack.Pop()
	                                   parsingQuery.limit = curr
	                                 }
	case 12:
		//line unql.y:70
		{ curr := make(map[string]interface{})
	                                   curr["op"] = "offset"
	                                   curr["expression"] = parsingStack.Pop()
	                                   parsingQuery.offset = curr
	                                 }
	case 15:
		//line unql.y:81
		{ curr := make(map[string]interface{})
	                             curr["op"] = "sort"
	                             curr["expression"] = parsingStack.Pop()
	                             curr["ascending"] = true
	                             parsingQuery.orderby = append(parsingQuery.orderby, curr)
	                           }
	case 16:
		//line unql.y:87
		{ curr := make(map[string]interface{})
	                             curr["op"] = "sort"
	                             curr["expression"] = parsingStack.Pop()
	                             curr["ascending"] = true
	                             parsingQuery.orderby = append(parsingQuery.orderby, curr)
	                           }
	case 17:
		//line unql.y:93
		{ curr := make(map[string]interface{})
	                              curr["op"] = "sort"
	                              curr["expression"] = parsingStack.Pop()
	                              curr["ascending"] = false
	                              parsingQuery.orderby = append(parsingQuery.orderby, curr)
	                           }
	case 18:
		//line unql.y:101
		{ logDebugGrammar("SELECT_COMPOUND") }
	case 24:
		//line unql.y:111
		{ logDebugGrammar("SELECT_CORE") }
	case 25:
		//line unql.y:114
		{ logDebugGrammar("SELECT GROUP")
	                                         parsingQuery.isAggregateQuery = true 
	                                         parsingQuery.groupby = parsingStack.Pop().([]interface{}) }
	case 26:
		//line unql.y:119
		{ logDebugGrammar("SELECT GROUP HAVING - EMPTY") }
	case 27:
		//line unql.y:120
		{ logDebugGrammar("SELECT GROUP HAVING - SELECT GROUP") }
	case 28:
		//line unql.y:121
		{ logDebugGrammar("SELECT GROUP HAVING - SELECT GROUP SELECT HAVING") }
	case 30:
		//line unql.y:129
		{ logDebugGrammar("SELECT WHERE - EMPTY") }
	case 31:
		//line unql.y:130
		{ logDebugGrammar("SELECT WHERE - EXPR")
	                               where_part := parsingStack.Pop()
	                               parsingQuery.where = where_part.(map[string]interface{}) }
	case 33:
		//line unql.y:136
		{ logDebugGrammar("SELECT_FROM") }
	case 36:
		//line unql.y:143
		{ ds := NewDataSource(yyS[yypt-0].s)
	                                   parsingQuery.AddDataSource(ds) 
	                                 }
	case 37:
		//line unql.y:146
		{ ds := NewDataSourceWithAs(yyS[yypt-2].s, yyS[yypt-0].s) 
	                                          parsingQuery.AddDataSource(ds) 
	                                        }
	case 39:
		//line unql.y:153
		{ logDebugGrammar("SELECT_SELECT") }
	case 40:
		//line unql.y:156
		{ logDebugGrammar("SELECT_SELECT_HEAD") }
	case 43:
		//line unql.y:161
		{ logDebugGrammar("SELECT SELECT TAIL - EMPTY") }
	case 44:
		//line unql.y:162
		{ logDebugGrammar("SELECT SELECT TAIL - EXPR")
	                            curr := make(map[string]interface{})
	                            curr["op"] = "select"
	                            curr["expression"] = parsingStack.Pop()
	                            parsingQuery.sel = curr
	                          }
	case 45:
		//line unql.y:168
		{ logDebugGrammar("SELECT SELECT TAIL - EXPR AS IDENTIFIER")
	                                          curr := make(map[string]interface{})
	                                          curr["op"] = "select"
	                                          curr["expression"] = parsingStack.Pop()
	                                          curr["as"] = yyS[yypt-0].s
	                                          parsingQuery.sel = curr
	                                        }
	case 46:
		//line unql.y:175
		{ logDebugGrammar("SELECT SELECT TAIL - AS IDENTIFIER")
	                               curr := make(map[string]interface{})
	                               curr["op"] = "select"
	                               curr["as"] = yyS[yypt-0].s
	                               parsingQuery.sel = curr
	                             }
	case 47:
		//line unql.y:183
		{ logDebugGrammar("EXPRESSION") }
	case 53:
		//line unql.y:191
		{  logDebugGrammar("EXPR - AND")
	                             
	                            curr := make(map[string]interface{})
	                            curr["right"] = parsingStack.Pop()
	                            curr["left"] = parsingStack.Pop()
	                            logDebugGrammar("LHS %v and RHS %v", curr["left"], curr["right"])
	                            curr["op"] = "&&"
	                            parsingStack.Push(curr)
	                         }
	case 54:
		//line unql.y:200
		{  logDebugGrammar("EXPR - OR")
	                            curr := make(map[string]interface{})
	                            curr["right"] = parsingStack.Pop()
	                            curr["left"] = parsingStack.Pop()
	                            logDebugGrammar("LHS %v and RHS %v", curr["left"], curr["right"])
	                            curr["op"] = "||"
	                            parsingStack.Push(curr)
	                         }
	case 55:
		//line unql.y:208
		{  logDebugGrammar("EXPR - EQ")
	                            curr := make(map[string]interface{})
	                            curr["right"] = parsingStack.Pop()
	                            curr["left"] = parsingStack.Pop()
	                            logDebugGrammar("LHS %v and RHS %v", curr["left"], curr["right"])
	                            curr["op"] = "=="
	                            parsingStack.Push(curr)
	                         }
	case 56:
		//line unql.y:216
		{  logDebugGrammar("EXPR - LT")
	                            curr := make(map[string]interface{})
	                            curr["right"] = parsingStack.Pop()
	                            curr["left"] = parsingStack.Pop()
	                            logDebugGrammar("LHS %v and RHS %v", curr["left"], curr["right"])
	                            curr["op"] = "<"
	                            parsingStack.Push(curr)
	                         }
	case 57:
		//line unql.y:224
		{  logDebugGrammar("EXPR - LTE")
	                            curr := make(map[string]interface{})
	                            curr["right"] = parsingStack.Pop()
	                            curr["left"] = parsingStack.Pop()
	                            logDebugGrammar("LHS %v and RHS %v", curr["left"], curr["right"])
	                            curr["op"] = "<="
	                            parsingStack.Push(curr)
	                         }
	case 58:
		//line unql.y:232
		{  logDebugGrammar("EXPR - GT")
	                            curr := make(map[string]interface{})
	                            right := parsingStack.Pop()
	                            left := parsingStack.Pop()
	                            curr["right"] = right
	                            curr["left"] = left
	                            logDebugGrammar("LHS %v and RHS %v", left, right)
	                            curr["op"] = ">"
	                            parsingStack.Push(curr)
	                         }
	case 59:
		//line unql.y:242
		{  logDebugGrammar("EXPR - GTE")
	                            curr := make(map[string]interface{})
	                            curr["right"] = parsingStack.Pop()
	                            curr["left"] = parsingStack.Pop()
	                            logDebugGrammar("LHS %v and RHS %v", curr["left"], curr["right"])
	                            curr["op"] = ">="
	                            parsingStack.Push(curr)
	                         }
	case 60:
		//line unql.y:250
		{  logDebugGrammar("EXPR - NE")
	                            curr := make(map[string]interface{})
	                            curr["right"] = parsingStack.Pop()
	                            curr["left"] = parsingStack.Pop()
	                            logDebugGrammar("LHS %v and RHS %v", curr["left"], curr["right"])
	                            curr["op"] = "!="
	                            parsingStack.Push(curr)
	                         }
	case 65:
		//line unql.y:267
		{ logDebugGrammar("SUFFIX_EXPR") }
	case 66:
		//line unql.y:270
		{ curr := make(map[string]interface{})
	                 curr["property"] = yyS[yypt-0].s
	                 parsingStack.Push(curr) }
	case 67:
		//line unql.y:273
		{ curr := make(map[string]interface{})
	                 curr["identifier"] = yyS[yypt-0].s
	                 parsingStack.Push(curr) }
	case 68:
		//line unql.y:276
		{ curr := make(map[string]interface{})
	                 curr["int"] = yyS[yypt-0].n
	                 parsingStack.Push(curr) }
	case 69:
		//line unql.y:279
		{ curr := make(map[string]interface{})
	                 curr["real"] = yyS[yypt-0].f
	                 parsingStack.Push(curr) }
	case 70:
		//line unql.y:282
		{ curr := make(map[string]interface{})
	                 curr["string"] = yyS[yypt-0].s
	                 parsingStack.Push(curr) }
	case 71:
		//line unql.y:285
		{ curr := make(map[string]interface{})
	                 curr["bool"] = true
	                 parsingStack.Push(curr) }
	case 72:
		//line unql.y:288
		{ curr := make(map[string]interface{})
	                 curr["bool"] = false
	                 parsingStack.Push(curr) }
	case 73:
		//line unql.y:291
		{ logDebugGrammar("ATOM - {}")
	                                              curr := make(map[string]interface{})
	                                              curr["object"] = parsingStack.Pop()
	                                              parsingStack.Push(curr)
	                                            }
	case 74:
		//line unql.y:296
		{ logDebugGrammar("ATOM - []")
	                                            curr := make(map[string]interface{})
	                                            curr["array"] = parsingStack.Pop()
	                                            parsingStack.Push(curr)
	                                          }
	case 75:
		//line unql.y:301
		{ logDebugGrammar("FUNCTION - $1.s")
	                                                      expression_list := parsingStack.Pop().([]interface{})
	                                                      function_map := parsingStack.Pop().(map[string]interface{})
	                                                      function_map["expression_list"] = expression_list
	                                                      parsingStack.Push(function_map)
	                                                    }
	case 78:
		//line unql.y:311
		{ logDebugGrammar("EXPRESSION_LIST - EXPRESSION")
	                                curr := make([]interface{},0)
	                               curr = append(curr, parsingStack.Pop())
	                               parsingStack.Push(curr)
	                             }
	case 79:
		//line unql.y:316
		{ logDebugGrammar("EXPRESSION_LIST - EXPRESSION COMMA EXPRESSION_LIST")
	                                               rest := parsingStack.Pop().([]interface{})
	                                               last := parsingStack.Pop().(interface{})
	                                               //for _,v := range last {
                                                 rest = append(rest,last)
	                                               //}
                                               parsingStack.Push(rest)
	                                             }
	case 81:
		//line unql.y:327
		{ last := parsingStack.Pop().(map[string]interface{})
	                                                                  rest := parsingStack.Pop().(map[string]interface{})
	                                                                  for k,v := range last {
	                                                                    rest[k] = v
	                                                                  }
	                                                                  parsingStack.Push(rest)
	                                                                }
	case 82:
		//line unql.y:336
		{ curr := make(map[string]interface{})
	                                                     curr[yyS[yypt-2].s] = parsingStack.Pop()
	                                                     parsingStack.Push(curr)
	                                                   }
	case 84:
		//line unql.y:345
		{ curr := make(map[string]interface{})
	                     curr["function"] = "min"
	                     parsingStack.Push(curr) }
	case 85:
		//line unql.y:348
		{ curr := make(map[string]interface{})
	                  curr["function"] = "max"
	                  parsingStack.Push(curr) }
	case 86:
		//line unql.y:351
		{ curr := make(map[string]interface{})
	                  curr["function"] = "avg"
	                  parsingStack.Push(curr) }
	case 87:
		//line unql.y:354
		{ curr := make(map[string]interface{})
	                    curr["function"] = "count"
	                    parsingStack.Push(curr) }
	}
	goto yystack /* stack new state and value */
}