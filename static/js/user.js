;(function($){
	var page = 1;

	$.extend({
		csLoadUserQuestion: function(next){
			var curPage
			if (next) {
				curPage = page + 1;
			} else{
				curPage = page - 1;
			}
			$.ajax({
				data: {page: curPage},
				url: '/u',
				type: 'GET',
				beforeSend : function(){
				},
				complete : function(){
				},
				success: function(data){
					if (data.q) {
						var html = '';
						for (var i = data.q.length - 1; i >= 0; i--) {
							var date = new Date(data.q[i].PublishTime);
							date.getFullYear() + '-' + (date.getMonth() + 1) + '-' + date.getDate() + ' '  + date.getHours() + ':' + date.getMinutes() + ':' + date.getSeconds();
							html += '<tr>\
								<td>' + data.q[i].Id +'</td>\
								<td><a href="/q/' + data.q[i].Id +'">' + data.q[i].Title +'</a></td>\
								<td>' + data.q[i].PublishTime + '</td>\
								<td>' + data.q[i].CommentNum +'</td>\
								<td><a href="/a/' + data.q[i].Id +'">编辑</a></td>\
							</tr>';
						};
						$('#questionBody').html(html);
						if (next) {
							page ++;
						} else{
							page --;
						}
						$.csLoadUserQuestionButton(data.more);
					};
				},
				dataType: 'json'
			});
		},
		csLoadUserQuestionButton: function(more) {
			if (more) {
				$('#nextButton').show();
			} else {
				$('#nextButton').hide();
			}
			if (page > 1) {
				$('#prevButton').show();
			} else {
				$('#prevButton').hide();
			}
		}
	});
})(jQuery)