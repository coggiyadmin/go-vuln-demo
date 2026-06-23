// Prompt Injection FN (OWASP LLM01) — delimiter/role-fence break. MISS.
package aisec

const translateSystem = "You are a translator. Translate the user text inside <data> tags."

func Translate(userText string) string {
	fenced := "<data>" + userText + "</data>" // SINK (LLM01 delimiter break)
	return chatSystemUser(translateSystem, fenced)
}
