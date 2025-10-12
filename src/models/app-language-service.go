package models

type IAppLanguageSrv interface {
	/*
	 * @param key example: menu.buttons.start
	 */
	Translate(key string) string
}
