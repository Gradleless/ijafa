package lib

import (
	"fmt"
	"regexp"
	"strings"
)

var dictionary = [][]string{
	{"break", "rupture"},
	{"case", "boite"},
	{"case", "boitier"},
	{"catch", "attraper"},
	{"continue", "continuer"},
	{"delete", "supprimer"},
	{"do", "effectuer"},
	{"else", "sinon"},
	{"finally", "finalement"},
	{"for", "pour"},
	{"function", "fonction"},
	{"if", "si"},
	{"in", "dans"},
	{"instanceof", "instanceDe"},
	{"new", "nouveau"},
	{"return", "retourner"},
	{"switch", "changer"},
	{"throw", "lancer"},
	{"try", "essayer"},
	{"typeof", "typeDe"},
	{"var", "variable"},
	{"void", "vide"},
	{"while", "tantQue"},
	{"with", "avec"},
	{"abstract", "abstrait"},
	{"boolean", "booleen"},
	{"byte", "octet"},
	{"char", "caractere"},
	{"class", "classe"},
	{"const", "constante"},
	{"debugger", "debogueur"},
	{"double", "double"},
	{"enum", "enumeration"},
	{"export", "exporter"},
	{"extends", "etendre"},
	{"final", "final"},
	{"float", "flottant"},
	{"goto", "allerA"},
	{"implements", "implementer"},
	{"import", "importer"},
	{"int", "entier"},
	{"interface", "interface"},
	{"long", "long"},
	{"native", "natif"},
	{"package", "paquetage"},
	{"private", "prive"},
	{"protected", "protege"},
	{"public", "public"},
	{"short", "court"},
	{"static", "statique"},
	{"super", "super"},
	{"synchronized", "synchronise"},
	{"throws", "lance"},
	{"transient", "transitoire"},
	{"volatile", "volatile"},
	{"null", "nul"},
	{"true", "vrai"},
	{"false", "faux"},
	{"NaN", "NaN"},
	{"Infinity", "Infini"},
	{"undefined", "indefini"},
	{"arguments", "arguments"},
	{"eval", "eval"},
	{"this", "ceci"},
	{"prototype", "prototype"},
	{"constructor", "constructeur"},
	{"length", "longueur"},
	{"prototype", "prototype"},
	{"toString", "versChaine"},
	{"toSource", "versSource"},
	{"for", "pour"},
	{"function", "fonction"},
	{"if", "si"},
	{"in", "dans"},
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
			fmt.Println("word between quotes")
			splited := strings.Split(word, "\"")
			word = strings.ReplaceAll(splited[0], v[0], v[1]) + "\"" + splited[1] + "\"" + strings.ReplaceAll(splited[2], v[0], v[1])
		}
	}
	return word
}
