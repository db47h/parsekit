package lexer_test

import "github.com/db47h/asm/token"

var tmplResultData = []struct {
	pos string
	t   token.Type
	s   string
}{
	{"1", 18, "\"\\n \\n\""},
	{"3", 9, "\"{{\""},
	{"3", 23, "<define>"},
	{"3", 16, "\" \""},
	{"3", 17, "\"root\""},
	{"3", 14, "\"}}\""},
	{"3", 18, "\"\\n<!DOCTYPE\"..."},
	{"7", 9, "\"{{\""},
	{"7", 7, "\".Title\""},
	{"7", 14, "\"}}\""},
	{"7", 18, "\"</title>\\n \"..."},
	{"10", 9, "\"{{\""},
	{"10", 7, "\".NotesEnab\"..."},
	{"10", 14, "\"}}\""},
	{"10", 18, "\";\\n    </sc\"..."},
	{"14", 9, "\"{{\""},
	{"14", 26, "<if>"},
	{"14", 16, "\" \""},
	{"14", 7, "\".NotesEnab\"..."},
	{"14", 14, "\"}}\""},
	{"14", 18, "\"\\n    <scri\"..."},
	{"16", 9, "\"{{\""},
	{"16", 7, "\".Sections\""},
	{"16", 14, "\"}}\""},
	{"16", 18, "\";\\n      va\"..."},
	{"17", 9, "\"{{\""},
	{"17", 7, "\".TitleNote\"..."},
	{"17", 14, "\"}}\""},
	{"17", 18, "\"\\n    </scr\"..."},
	{"20", 9, "\"{{\""},
	{"20", 25, "<end>"},
	{"20", 14, "\"}}\""},
	{"20", 18, "\"\\n\\n    <scr\"..."},
	{"46", 9, "\"{{\""},
	{"46", 7, "\".Title\""},
	{"46", 14, "\"}}\""},
	{"46", 18, "\"</h1>\\n    \"..."},
	{"47", 9, "\"{{\""},
	{"47", 30, "<with>"},
	{"47", 16, "\" \""},
	{"47", 7, "\".Subtitle\""},
	{"47", 14, "\"}}\""},
	{"47", 18, "\"<h3>\""},
	{"47", 9, "\"{{\""},
	{"47", 22, "<.>"},
	{"47", 14, "\"}}\""},
	{"47", 18, "\"</h3>\""},
	{"47", 9, "\"{{\""},
	{"47", 25, "<end>"},
	{"47", 14, "\"}}\""},
	{"47", 18, "\"\\n        \""},
	{"48", 9, "\"{{\""},
	{"48", 26, "<if>"},
	{"48", 16, "\" \""},
	{"48", 8, "\"not\""},
	{"48", 16, "\" \""},
	{"48", 7, "\".Time\""},
	{"48", 7, "\".IsZero\""},
	{"48", 14, "\"}}\""},
	{"48", 18, "\"<h3>\""},
	{"48", 9, "\"{{\""},
	{"48", 7, "\".Time\""},
	{"48", 7, "\".Format\""},
	{"48", 16, "\" \""},
	{"48", 17, "\"2 January 2006\""},
	{"48", 14, "\"}}\""},
	{"48", 18, "\"</h3>\""},
	{"48", 9, "\"{{\""},
	{"48", 25, "<end>"},
	{"48", 14, "\"}}\""},
	{"48", 18, "\"\\n        \""},
	{"49", 9, "\"{{\""},
	{"49", 28, "<range>"},
	{"49", 16, "\" \""},
	{"49", 7, "\".Authors\""},
	{"49", 14, "\"}}\""},
	{"49", 18, "\"\\n         \"..."},
	{"51", 9, "\"{{\""},
	{"51", 28, "<range>"},
	{"51", 16, "\" \""},
	{"51", 7, "\".TextElem\""},
	{"51", 14, "\"}}\""},
	{"51", 9, "\"{{\""},
	{"51", 8, "\"elem\""},
	{"51", 16, "\" \""},
	{"51", 19, "\"$\""},
	{"51", 7, "\".Template\""},
	{"51", 16, "\" \""},
	{"51", 22, "<.>"},
	{"51", 14, "\"}}\""},
	{"51", 9, "\"{{\""},
	{"51", 25, "<end>"},
	{"51", 14, "\"}}\""},
	{"51", 18, "\"\\n         \"..."},
	{"53", 9, "\"{{\""},
	{"53", 25, "<end>"},
	{"53", 14, "\"}}\""},
	{"53", 18, "\"\\n      </a\"..."},
	{"56", 9, "\"{{\""},
	{"56", 28, "<range>"},
	{"56", 16, "\" \""},
	{"56", 19, "\"$i\""},
	{"56", 2, "\",\""},
	{"56", 16, "\" \""},
	{"56", 19, "\"$s\""},
	{"56", 16, "\" \""},
	{"56", 5, "\":=\""},
	{"56", 16, "\" \""},
	{"56", 7, "\".Sections\""},
	{"56", 14, "\"}}\""},
	{"56", 18, "\"\\n  <!-- st\"..."},
	{"57", 9, "\"{{\""},
	{"57", 19, "\"$s\""},
	{"57", 7, "\".Number\""},
	{"57", 14, "\"}}\""},
	{"57", 18, "\" -->\\n     \"..."},
	{"58", 9, "\"{{\""},
	{"58", 19, "\"$s\""},
	{"58", 7, "\".HTMLAttri\"..."},
	{"58", 14, "\"}}\""},
	{"58", 18, "\">\\n      \""},
	{"59", 9, "\"{{\""},
	{"59", 26, "<if>"},
	{"59", 16, "\" \""},
	{"59", 19, "\"$s\""},
	{"59", 7, "\".Elem\""},
	{"59", 14, "\"}}\""},
	{"59", 18, "\"\\n        <\"..."},
	{"60", 9, "\"{{\""},
	{"60", 19, "\"$s\""},
	{"60", 7, "\".Title\""},
	{"60", 14, "\"}}\""},
	{"60", 18, "\"</h3>\\n    \"..."},
	{"61", 9, "\"{{\""},
	{"61", 28, "<range>"},
	{"61", 16, "\" \""},
	{"61", 19, "\"$s\""},
	{"61", 7, "\".Elem\""},
	{"61", 14, "\"}}\""},
	{"61", 9, "\"{{\""},
	{"61", 8, "\"elem\""},
	{"61", 16, "\" \""},
	{"61", 19, "\"$\""},
	{"61", 7, "\".Template\""},
	{"61", 16, "\" \""},
	{"61", 22, "<.>"},
	{"61", 14, "\"}}\""},
	{"61", 9, "\"{{\""},
	{"61", 25, "<end>"},
	{"61", 14, "\"}}\""},
	{"61", 18, "\"\\n      \""},
	{"62", 9, "\"{{\""},
	{"62", 24, "<else>"},
	{"62", 14, "\"}}\""},
	{"62", 18, "\"\\n        <\"..."},
	{"63", 9, "\"{{\""},
	{"63", 19, "\"$s\""},
	{"63", 7, "\".Title\""},
	{"63", 14, "\"}}\""},
	{"63", 18, "\"</h2>\\n    \"..."},
	{"64", 9, "\"{{\""},
	{"64", 25, "<end>"},
	{"64", 14, "\"}}\""},
	{"64", 18, "\"\\n      </a\"..."},
	{"66", 9, "\"{{\""},
	{"66", 19, "\"$i\""},
	{"66", 14, "\"}}\""},
	{"66", 18, "\" -->\\n  \""},
	{"67", 9, "\"{{\""},
	{"67", 25, "<end>"},
	{"67", 14, "\"}}\""},
	{"67", 18, "\"\\n\\n      <a\"..."},
	{"71", 9, "\"{{\""},
	{"71", 28, "<range>"},
	{"71", 16, "\" \""},
	{"71", 7, "\".Authors\""},
	{"71", 14, "\"}}\""},
	{"71", 18, "\"\\n         \"..."},
	{"73", 9, "\"{{\""},
	{"73", 28, "<range>"},
	{"73", 16, "\" \""},
	{"73", 7, "\".Elem\""},
	{"73", 14, "\"}}\""},
	{"73", 9, "\"{{\""},
	{"73", 8, "\"elem\""},
	{"73", 16, "\" \""},
	{"73", 19, "\"$\""},
	{"73", 7, "\".Template\""},
	{"73", 16, "\" \""},
	{"73", 22, "<.>"},
	{"73", 14, "\"}}\""},
	{"73", 9, "\"{{\""},
	{"73", 25, "<end>"},
	{"73", 14, "\"}}\""},
	{"73", 18, "\"\\n         \"..."},
	{"75", 9, "\"{{\""},
	{"75", 25, "<end>"},
	{"75", 14, "\"}}\""},
	{"75", 18, "\"\\n      </a\"..."},
	{"86", 9, "\"{{\""},
	{"86", 26, "<if>"},
	{"86", 16, "\" \""},
	{"86", 7, "\".PlayEnabl\"..."},
	{"86", 14, "\"}}\""},
	{"86", 18, "\"\\n    <scri\"..."},
	{"88", 9, "\"{{\""},
	{"88", 25, "<end>"},
	{"88", 14, "\"}}\""},
	{"88", 18, "\"\\n\\n    <scr\"..."},
	{"102", 9, "\"{{\""},
	{"102", 25, "<end>"},
	{"102", 14, "\"}}\""},
	{"102", 18, "\"\\n\\n\""},
	{"104", 9, "\"{{\""},
	{"104", 23, "<define>"},
	{"104", 16, "\" \""},
	{"104", 17, "\"newline\""},
	{"104", 14, "\"}}\""},
	{"104", 18, "\"\\n<br>\\n\""},
	{"106", 9, "\"{{\""},
	{"106", 25, "<end>"},
	{"106", 14, "\"}}\""},
	{"106", 18, "\"\\n\""},
	{"107", 6, "EOF"},
}
