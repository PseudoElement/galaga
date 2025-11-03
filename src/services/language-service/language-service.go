package language_srv

import (
	"encoding/json"
	"os"
	"strings"

	consts "github.com/pseudoelement/galaga/src/constants"
	"github.com/pseudoelement/galaga/src/models"
)

type AppLanguageSrv struct {
	// langScheme LanguageScheme
	langScheme map[consts.Language]map[string]interface{}
	storage    models.IAppStorage
}

func NewAppLanguageSrv(storage models.IAppStorage) models.IAppLanguageSrv {
	appLangSrv := &AppLanguageSrv{
		langScheme: make(map[consts.Language]map[string]interface{}),
		storage:    storage,
	}
	appLangSrv.loadLanguages()

	return appLangSrv
}

/*
 * @param key example: menu.buttons.start
 */
func (ls *AppLanguageSrv) Translate(translationPath string) string {
	languageMap := ls.langScheme[ls.storage.Language()]

	var translation any = nil
	keys := strings.Split(translationPath, ".")
	for _, key := range keys {
		if translation == nil {
			translation = languageMap[key]
		} else {
			translation = translation.(map[string]interface{})[key]
		}
	}

	str, ok := translation.(string)
	if !ok {
		return "nil"
	}

	return str
}

func (ls *AppLanguageSrv) loadLanguages() {
	for langKey, _ := range consts.LANGUAGE_TO_TRANSLATE_MAP {
		ls.loadLanguageJSON(langKey)
	}
}

func (ls *AppLanguageSrv) loadLanguageJSON(lang consts.Language) {
	pwd, _ := os.Getwd()
	pathToFile := pwd + "/lang/" + consts.LANGUAGE_TO_TRANSLATE_MAP[lang] + ".json"
	fileBytes, err := os.ReadFile(pathToFile)
	if err != nil {
		panic(err)
	}

	ls.langScheme[lang] = make(map[string]interface{})
	languageMap := ls.langScheme[lang]
	err = json.Unmarshal(fileBytes, &languageMap)
	if err != nil {
		panic(err)
	}
}
