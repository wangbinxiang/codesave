//1.初始化一个评论发布功能
;(function($){
	var page = 1;
	$.fn.extend({
		csNewComment: function(options) {
			var defaults = {
				mouseLocation: {},
				postUrl: '',
				qid: 0,
				isLogin: false,
				uid: 0
			};
			var settings = $.extend({}, defaults, options);
			var _this = this

			$('body').append('<div id="popover" class="popover fade right in" style="display: hidden;"><div class="arrow"></div><div id="popoverButton" class="popover-content text-primary btn-link">评论</div></div>');
			$('body').append('<div id="myModal"class="modal fade"><div class="modal-dialog"><div class="modal-content"><form id="commentForm"><div class="modal-header"><button type="button"class="close"data-dismiss="modal"aria-hidden="true">&times;</button><h4 class="modal-title">评论</h4></div><div class="modal-body"><textarea id="commentText"class="form-control validate[required,minSize[5],maxSize[255]]"rows="5"></textarea><p class="text-right text-muted">您还可以输入<b id="commentNum">255</b>个字符</p></div><div class="modal-footer"style="clear:both;"><button type="button"class="btn btn-default"data-dismiss="modal">取消</button><button type="button"id="commentSubmit"class="btn btn-primary">发表</button></div></form></div></div></div>');

			$('#commentForm').validationEngine();

			this.bind('dblclick', function(event){
		        if (document.selection && document.selection.empty) {//双击不选中
		            document.selection.empty(); 
		        } else if (window.getSelection) { 
		            var sel = window.getSelection(); 
		            sel.removeAllRanges(); 
		        }
		        settings.mouseLocation.pageX = event.pageX;
		        settings.mouseLocation.pageY = event.pageY;
		        $('#popover').css({top: event.pageY - 20, left: event.pageX + 5}).show();
		    }).bind('click', function(){
		      $('#popover').hide();
		    });

		    $('#popoverButton').bind('click', function(){
			    $('#myModal').modal();
			    $('#popover').hide();
			});

		    $('#commentText').keypress(function(e) { 
				if (e.ctrlKey && e.which == 13) {
					$('#commentSubmit').click();
					return false;
				}
			})

			$('#commentText').bind('keyup', _commentKey).bind('keydown', _commentKey);

			$('#commentSubmit').bind('click', function(){
				if ($("#commentForm").validationEngine("validate") == true) {
					var position = _this.position();
					var left = settings.mouseLocation.pageX - position.left;
					var top = settings.mouseLocation.pageY - position.top;
					var content = $('#commentText').val();
					
					var xsrf, xsrflist;
					xsrf = $.cookie("_xsrf");
					xsrflist = xsrf.split("|");

					var data = {Qid: settings.qid, Content: content, Left: left, Top: top, _xsrf: base64_decode(xsrflist[0])};
					$.ajax({
						data: data,
						url: settings.postUrl,
						type: 'POST',
						beforeSend : function(){
						},
						complete : function(){
						},
						success: function(data){
							if (data.result && data.id > 0) {
								$('#myModal').modal('hide');
								$.csShowComment(content, data.id, user_nickname, settings.mouseLocation.pageX, settings.mouseLocation.pageY);
								$('#commentText').val('');
								_commentKey();
							};
						},
						dataType: 'json'
					});
					
				};
		    });
		}
	});

	$.extend({
		csShowComment: function(comment, id, nickname, left, top) {
			insertComment(comment, id, nickname, left, top);
		},
		csLoadComment: function(next){
			var curPage
			if (next) {
				curPage = page + 1;
			} else{
				curPage = page - 1;
			}
			$.ajax({
				data: {qid: qid, page: curPage},
				url: '/q/c',
				type: 'GET',
				beforeSend : function(){
				},
				complete : function(){
				},
				success: function(data){
					if (data.c) {
						$('.comment').remove();
						for (var i = data.c.length - 1; i >= 0; i--) {
							$.csShowComment(data.c[i].Content, data.c[i].Id, data.c[i].Nickname, position.left + parseFloat(data.c[i].Left), position.top + parseFloat(data.c[i].Top));
						};
						$.csShowCommentBody(data.c);
						if (next) {
							page ++;
						} else{
							page --;
						}
						$.csShowPageButton(data.cMore);
					};
				},
				dataType: 'json'
			});
		},
		csShowPageButton: function(more) {
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
		},
		csShowCommentBody: function(comments) {
			var html = '';
			for (var i = comments.length - 1; i >= 0; i--) {
				html += '<tr><td>' + comments[i].Nickname + '</td><td>' +substr(comments[i].Content, 0, 30) + '</td></tr>';
			};
			$('#commentBody').html(html);
		}
	});
	
	function _commentKey(){
		var len = $.trim($('#commentText').val()).length;
		var num = 255 - len;
		$('#commentNum').text(num);
		if (num < 0) {
			$('#commentNum').addClass('text-danger');
		}else{
			$('#commentNum').removeClass('text-danger');
		}
	}
	function insertComment(comment, id, nickname, left, top){
		comment = htmlspecialchars($.trim(comment));
		var shortComment = substr(comment, 0, 10);
		var commentHtml = '<div id="comment' + id + '" class="comment popover fade right in" style="top: ' + top + 'px; left: ' + left + 'px; display: block;">\
		<div id="commentArrow' + id + '" class="arrow"></div>\
		<div id="commentTitle' + id + '" class="popover-content" style="border-bottom: 1px solid #eee;display:none;">\
		<a>' + nickname + '</a> <button type="button" class="close commentClose' + id + '" aria-hidden="true">&times;</button><button id="commentHidden' + id + '" type="button" class="close" aria-hidden="true">&minus;</button>\
		</div>\
		<div id="commentShort' + id + '" class="pull-left popover-content text-primary">' + shortComment + '...<button type="button" class="close commentClose' + id + '" aria-hidden="true">&times;</button>\
		</div>\
		<div id="commentTotal' + id + '" class="pull-left popover-content text-primary hidden">' + nl2br(comment) + '</div>\
		</div>';
		$('body').append(commentHtml);
		_bindComment(id);
	}

	function _bindComment(id){
		$('#commentShort' + id).bind('click', function(){
			$('#commentTitle' + id).fadeIn();
			var baseHeight = parseInt($('#test').css('height'));
			$('#comment' + id).animate({'max-width': "500px", 'min-width': '276px'}, 500);
			$('#commentShort' + id).addClass('hidden');
			$('#commentTotal' + id).removeClass('hidden');
			var newHeight = parseInt($('#comment' + id).css('height'));
				$('#commentArrow' + id).css({'top': (baseHeight/newHeight * 100 / 2 ) + '%'});
			});
			$('.commentClose' + id).bind('click', function(event){
				$('#comment' + id).hide();
				event.stopPropagation();
			});
			$('#commentHidden' + id).bind('click', function(){
			$('#commentTitle' + id).hide();
			$('#commentShort' + id).fadeIn().removeClass('hidden');
			$('#commentTotal' + id).addClass('hidden');
			$('#comment' + id).animate({'max-width': "276px", 'min-width': '0px'}, 500);
			$('#commentArrow' + id).css({'top': '50%'});
		});
	}
})(jQuery)
