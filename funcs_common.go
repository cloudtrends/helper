package helper

import (
	//"log"
	//. "domolo.com/drevel"
	"strings"
)

func GenerateDynamicTemplateName(tmplt_name, time_stamp, web_ui_theme_dir string) (template_name string) {
	web_ui_theme_dir = strings.TrimSpace(web_ui_theme_dir)
	tmplt_name = strings.TrimSpace(tmplt_name)
	time_stamp = strings.TrimSpace(time_stamp)
	if 0 == len(web_ui_theme_dir) {
		//web_ui_theme_dir = "default"
	}
	template_name = tmplt_name + time_stamp + web_ui_theme_dir
	template_name = strings.Replace(template_name, "/", "", -1)
	return
}

func GetAdminsTemplateFile(ori_tmplt_file string) string {
	var template_file string
	var theme_dir string
	theme_dir = "default"
	template_file = theme_dir + "/admins/" + ori_tmplt_file
	return template_file
}
