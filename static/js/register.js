$.validationEngineLanguage.allRules.onlyLetterNumber = {"regex": /^([^x00-xff\s]|[0-9a-zA-Z_])+$/, "alertText": "* 只能输入英文字母，数字，汉字，下划线。"};
$.validationEngineLanguage.allRules.ajaxUserCall.url = '/r/verify';
function showRecaptcha(pubKey) {  
	Recaptcha.create(pubKey, 'recaptcha', {  
		theme: "red",  
		callback: Recaptcha.focus_response_field
	});  
}