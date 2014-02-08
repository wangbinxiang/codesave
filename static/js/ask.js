function contentKeyUp() {
	$('#previewEditor').text($('#editor').val());
	prettyPrint();
}