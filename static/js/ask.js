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
	var tagLi = '<li><span class="label label-primary hand">' + tag.name + '</span><a href="#" class="tag-cancel" title="删除">&#10006;</a><input type="hidden" name="tags[]" value="' + tag.id + '" /></li>';
	$('#addTagList').append(tagLi);

}