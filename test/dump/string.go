package dump

var String = []string{
	"should return upper case first letters",
	"ğhould şeturn ipper çase öirst üetters",
	"Ğhould Şeturn İpper Çase Öirst Üetters",
	"Should Return LowLLer CASE ReplacE A<>!!mpty",
	"default_image.png",
}

var UpperCaseFirstLetterResponses = []string{
	"Should Return Upper Case First Letters",
	"Ğhould Şeturn İpper Çase Öirst Üetters",
	"Ğhould Şeturn İpper Çase Öirst Üetters",
	"Should Return Lowller Case Replace A<>!!Mpty",
	"Default_image.png",
}

var LowerCaseReplaceEmptyResponses = []string{
	"shouldreturnuppercasefirstletters",
	"ğhouldşeturnipperçaseöirstüetters",
	"ğhouldşeturnipperçaseöirstüetters",
	"shouldreturnlowllercasereplacea<>!!mpty",
	"default_image.png",
}

var HashImageNameResponses = []string{
	"should_return_upper_case_first_letters",
	"ğhould_şeturn_ipper_çase_öirst_üetters",
	"ğhould_şeturn_ipper_çase_öirst_üetters",
	"should_return_lowller_case_replace_a<>!!mpty",
	"default_image.png",
}
