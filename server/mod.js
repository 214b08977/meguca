
(function(){var $panel;function show_panel(){if($panel)
return;var $del=$('<input type=button value=Delete>').click(korosu);$panel=$('<div></div>').append($del).css({position:'fixed',bottom:0,right:0}).appendTo('body');}
function korosu(){var ids=[];$('header>input').each(function(){var $check=$(this);if($check.attr('checked')){var id=$check.parent().parent().attr('id');ids.push(parseInt(id));}});if(ids.length){ids.unshift(5,document.cookie);send(ids);}
else{var $button=$(this);var caption=_.bind($button.val,$button);caption('Nothing selected.');_.delay(caption,2000,'Delete');}}
function make_alloc_mod(text){var msg=this.make_alloc_vanilla(text);if($('#mod').attr('checked'))
msg.auth='Moderator';if(msg.auth)
msg.cookie=document.cookie;return msg;}
$(document).click(function(event){var $box=$(event.target);if($box.attr('type')=='checkbox'&&$box.parent('header').length)
show_panel();});$(document).ready(function(){$('h1').text('Moderation - '+$('h1').text());$('<input type=checkbox>').insertBefore('header>:first-child');$name.after(' <input type=checkbox id=mod>'+'<label for=mod>Moderator</label>');var pfp=PostForm.prototype;pfp.make_alloc_vanilla=pfp.make_alloc_request;pfp.make_alloc_request=make_alloc_mod;oneeSama.check=function(target){$('<input type=checkbox>').insertBefore(target.find('>header>:first-child'));};});})();
