package models

type IAppLanguageSrv interface {
	/*
	 * @param key example: menu.buttons.start
	 */
	GetTranslation(key string) string
}
