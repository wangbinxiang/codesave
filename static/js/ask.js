function contentKeyUp() {
	$('#previewEditor').text($('#editor').val());
	$('#previewEditor').removeClass($('#previewEditor').attr('class'));
	$('pre code').each(function(i, e) {
		hljs.highlightBlock(e)
	});
}