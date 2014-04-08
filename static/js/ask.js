function contentKeyUp() {
	$('#previewEditor').text($('#editor').val());
	prettyPrint();
}
//绑定添加tag事件
function bindTagList() {
	$('#tagList span').each(function(){
		$(this).bind('click', function(tag){
			var tag = {id: $(this).attr('data-tid'), name: $(this).text()};
			addTag(tag);
		})
	});
	$('#addTagList').on('click', 'a', function(event){
		event.preventDefault()
		$(this).parent().remove();
	});
}
//添加tag
function addTag(tag) {
	if($('#addTagList input[value="' + tag.id + '"]').length > 0 ) {
		return;
	}
	if ($('#addTagList li').length >= 5) {
		return
	};
	var tagLi = '<li><span class="label label-primary hand">' + tag.name + '</span><a href="#" class="tag-cancel" title="删除">&#10006;</a><input type="hidden" name="tags" value="' + tag.id + '" /></li>';
	$('#addTagList').append(tagLi);
}

function bindSubmitButton() {
	$('#submitButton').bind('click', function() {
		var isSubmit = true, tagsLength = 0;
		if($("#askForm").validationEngine("validate") != true){
			isSubmit = false;
		}
		tagsLength = $('input[name="tags"]').length;
		if (tagsLength == 0 || tagsLength > 5) {
			$('#tagsAlert').show();
			isSubmit = false;
		};
		if (isSubmit) {
			$("#askForm").submit();
		};
	});
}

function bindTagsAlertButton() {
	$('#tagsAlertButton').bind('click', function() {
		$('#tagsAlert').hide();
	});
}

function askInit(edit) {
	$('#askForm').validationEngine();
	bindSubmitButton();
	bindTagsAlertButton();
	tabIndent.renderAll();
	$('#editor').bind('keyup', contentKeyUp);
	bindTagList();
	if (edit) {
		contentKeyUp();
	};
}