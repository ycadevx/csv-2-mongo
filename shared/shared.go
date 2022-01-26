package shared

var Config = configuration{
	FILEURL:  `/Users/YCA/go/src/github/csvToMongoDb\source\Document.csv`,
	SPERATOR: ';',
	ISHEADER: true,
	ROWCOUNT: 5,
}

type configuration struct {
	FILEURL  string
	MONGOURL string
	ISHEADER bool
	SPERATOR rune
	ROWCOUNT int
}
