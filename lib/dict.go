package lib

import (
	"regexp"
	"strings"
)

var dictionary = [][]string{
	// liste faite par chatgpt et copilot, donc il y a sûrement des erreurs et des doublons
	{"importer", "import"},
	{"paquet", "package"},
	{"publique", "public"},
	{"prive", "private"},
	{"rupture", "break"},
	{"boite", "case"},
	{"boitier", "case"},
	{"attraper", "catch"},
	{"continuer", "continue"},
	{"supprimer", "delete"},
	{"effectuer", "do"},
	{"sinon", "else"},
	{"finalement", "finally"},
	{"pour", "for"},
	{"fonction", "function"},
	{"si", "if"},
	{"nouveau", "new"},
	{"retourner", "return"},
	{"changer", "switch"},
	{"lancer", "throw"},
	{"essayer", "try"},
	{"typeDe", "typeof"},
	{"variable", "var"},
	{"vide", "void"},
	{"tantQue", "while"},
	{"avec", "with"},
	{"abstrait", "abstract"},
	{"booleen", "boolean"},
	{"octet", "byte"},
	{"caractere", "char"},
	{"classe", "class"},
	{"constante", "const"},
	{"debogueur", "debugger"},
	{"double", "double"},
	{"enumeration", "enum"},
	{"exporter", "export"},
	{"etendre", "extends"},
	{"final", "final"},
	{"flottant", "float"},
	{"implementer", "implements"},
	{"importer", "import"},
	{"entier", "int"},
	{"interface", "interface"},
	{"long", "long"},
	{"natif", "native"},
	{"paquetage", "package"},
	{"prive", "private"},
	{"protege", "protected"},
	{"public", "public"},
	{"court", "short"},
	{"statique", "static"},
	{"super", "super"},
	{"synchronise", "synchronized"},
	{"lance", "throws"},
	{"transitoire", "transient"},
	{"volatile", "volatile"},
	{"nul", "null"},
	{"vrai", "true"},
	{"faux", "false"},
	{"NaN", "NaN"},
	{"Infini", "Infinity"},
	{"indefini", "undefined"},
	{"arguments", "arguments"},
	{"eval", "eval"},
	{"ceci", "this"},
	{"prototype", "prototype"},
	{"constructeur", "constructor"},
	{"longueur", "length"},
	{"versChaine", "toString"},
	{"instance", "instance"},
	{"clonage", "clone"},
	{"finaliser", "finalize"},
	{"interface", "interface"},
	{"exception", "exception"},
	{"assertion", "assert"},
	{"superclasse", "superclass"},
	{"interface", "interface"},
	{"surcharge", "overload"},
	{"surcharge", "override"},
	{"this", "this"},
	{"super", "super"},
	{"statique", "static"},
	{"final", "final"},
	{"transitoire", "transient"},
	{"classeAnonyme", "anonymous class"},
	{"stream", "stream"}, // modif
	{"optionnel", "optional"},
	{"fonctionIdentite", "identity function"},
	{"referenceMethode", "method reference"},
	{"biPredicate", "BiPredicate"},
	{"tri", "sort"},
	{"filtrer", "filter"},
	{"mapper", "map"},
	{"reduire", "reduce"},
	{"collecter", "collect"},
	{"pourChaque", "forEach"},
	{"streamParallele", "parallelStream"},
	{"super", "super"},
	{"classe", "class"},
	{"heritage", "inheritance"},
	{"operateur", "operator"},
	{"metadonnees", "metadata"},
	{"enumeration", "enum"},
	{"defaut", "default"},
	{"constructeur", "constructor"},
	{"statique", "static"},
	{"scelle", "sealed"},
	{"nonNul", "nonnull"},
	{"non nul", "nonnull"},
	{"ouvrir", "open"},
	{"exports", "exports"},
	{"exige", "requires"},
	{"necessite", "requires"},
	{"transitif", "transitive"},
	{"transitionnel", "transitive"},
	{"utilise", "uses"},
	{"fournit", "provides"},
	{"avec", "with"},
	{"enregistrer", "records"},
	{"var", "var"},
	{"rendement", "yield"},
	// flemme d'ajouter + pr l'instant
}

func IsWordBetweenQuotes(word string, list string) bool {
	regex := `\"[^"]*` + list + `[^"]*\"`
	match, _ := regexp.MatchString(regex, strings.TrimSpace(word))
	return match
}

func GetTranslation(word string) string {
	for _, v := range dictionary {
		if strings.Contains(word, v[0]) && !IsWordBetweenQuotes(word, v[0]) {
			word = strings.ReplaceAll(word, v[0], v[1])
		} else if strings.Contains(word, v[0]) && IsWordBetweenQuotes(word, v[0]) {
			splited := strings.Split(word, "\"")
			word = strings.ReplaceAll(splited[0], v[0], v[1]) + "\"" + splited[1] + "\"" + strings.ReplaceAll(splited[2], v[0], v[1])
		}
	}
	return word
}
