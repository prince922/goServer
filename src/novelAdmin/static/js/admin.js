/**
 * Created by Administrator on 2017/2/10.
 */
/**
 * 更改状态
 * @param select 下拉列表
 * @param url  提交网址
 * @param text  操作文字
 */
function changeStatus(select, url, text) {
    text = text ? text : '';
    var Checkbox = false;
    var ids = [];
    $("input[name='ids[]']").each(function () {
        if (this.checked == true) {
            ids.push(this.value)
            Checkbox = true;
        }
    });
    if (Checkbox) {
        if (select.value == '1') {
            if (!confirm('你确定' + text + '吗？')) {
                select.value = '';
                return;
            }
        } else if (select.value == '0') {
            if (!confirm('你确定取消' + text + '吗？')) {
                select.value = '';
                return;
            }
        } else if (select.value == '2') {
            if (!confirm('你确定' + text + '吗？')) {
                select.value = '';
                return;
            }
        } else {
            select.value = '';
            return;
        }
        $.ajax({
            type: 'post',
            url: url,
            data: {ids: ids.toString(), tag: select.value},
            success: function (data) {
                if (data.res == 1) {
                    alert('操作成功！');
                    location.reload();
                } else {
                    alert('操作失败！');
                }
            }
        });
    } else {
        select.value = '';
        alert("请选择要操作的内容!");
    }
}

/**
 * 游戏类别
 * @param select
 */
function changeCategoryAudit(select) {
    var Checkbox = false;
    var ids = [];
    $("input[name='ids[]']").each(function () {
        if (this.checked == true) {
            ids.push(this.value)
            Checkbox = true;
        }
    });
    if (Checkbox) {
        if (select.value == '1') {
            if (!confirm('你确定审核吗？')) {
                select.value = '';
                return;
            }
        } else if (select.value == '0') {
            if (!confirm('你确定取消审核吗？')) {
                select.value = '';
                return;
            }
        } else {
            select.value = '';
            return;
        }
        $.ajax({
            type: 'post',
            url: "set_category_audit.html",
            data: {ids: ids.toString(), tag: select.value},
            success: function (data) {
                if (data.res == 1) {
                    alert('操作成功！');
                    location.reload();
                } else {
                    alert('操作失败！');
                }
            }
        });
    } else {
        select.value = '';
        alert("请选择要操作的内容!");
    }
}
/**
 * 更改分类排序
 */
function changeCategorySort() {
    $flag = true;
    $('input[name^="sortIds"]').each(function () {
        if ($flag && !this.value) {
            $flag = false;
            alert('排序值不能为空');
            this.focus();
        }
    });
    if (!confirm('你确定排序吗？')) {
        return;
    }
    if ($flag) {
        document.form.submit();
    }
}
/**
 * 添加图片点击
 * @param id
 */
function addImage(id) {
    $('#imageFile').attr('container', id).click();
    $('#imageTag').val(id);
}

/**
 * 添加图片点击
 * @param id
 */
function addVideo(id) {
    $('#videoFile').attr('container', id).click();
    $('#videoTag').val(id);
}

/**
 * 添加图片点击
 * @param id
 */
function addImageChapter(id,book_id) {
    $('#imageFile').attr('container', id).click();
    $('#imageTag').val(id+'_'+book_id);
}
/**
 * 更换图片点击
 * @param el
 */
function changeImage(el) {
    if (confirm('你是要换一张么？')) {
        $(el).siblings(".hidden").val('');
        $(el).siblings(".addUpload").show();
        $(el).remove();
    }
}

/**
 * 上传文件
 * @param input
 */
function uploadFile(input,formName) {
    if (input.files.length > 0) {
        $("form[name='" + formName + "']").submit();
        startUploadPopup();
    }
    input.value = null;
}

/**
 * 上传成功
 * @param data
 */
function uploadSuccess(url,data) {
    closeUploadPopup();
    // $('.addUpload').hide();
    var container = $('#imageFile').attr('container');
    if (container == '.pics') {
        $(container).append('<li class="img_node"><input name="chapter_imgs[]" type="hidden" value="' + data + '"><img src="' + url + data + '" alt=""><span class="close" onclick="closePic(this)"></span><span class="close left" onclick="leftPic(this)"></span><span class="close right" onclick="rightPic(this)"></span></li>');
    } else {
        $(container).siblings(".addUpload").hide();
        $(container).val(data);
        $(container).parent().append('<a href="javascript:void(0);" class="uploadPics" onclick="changeImage(this)"><img src="' + url + data + '" style="'+$(container).data('style')+'"></a>');
    }


    if(container == '#video_thump'){
        $('#thump_type').val(2);
    }
}


/**
 * 上传视频成功
 * @param data
 */
function uploadVideoSuccess(url,data) {
    closeUploadPopup();

    // $('.addUpload').hide();
    var container = $('#videoFile').attr('container');
    $(container).val(data);
}

function substrUrl(url) {
    var index = url.indexOf('http://img.ai7.com/');
    if (index != -1) {
        return url.substring(index + 19);
    } else {
        return url;
    }
}
function uploadError(data) {
    closeUploadPopup();
    alert(data);
}
function startUploadPopup() {
    $("#uploadPopup p").text('正在上传');
    $("#uploadPopup,.uploadBackdrop").show();
}
function closeUploadPopup() {
    $("#uploadPopup,.uploadBackdrop").hide();
}
function closePic(el) {
    $(el).parent().remove()
}
function leftPic(el) {
    if($(el).parent().prev().hasClass('img_node'))
        $(el).parent().insertBefore($(el).parent().prev());
}
function rightPic(el) {
    if($(el).parent().next().hasClass('img_node'))
        $(el).parent().insertAfter($(el).parent().next());
}

//删除首页图片
function imgDel(ids, url) {
    if (!ids) {
        alert('请勾选需要操作的选项！');
        return;
    }
    if (!confirm('请确定要删除吗' + ids + '？')) {
        return;
    }
    $.ajax({
        type: 'POST',
        url: 'img_del.html',
        data: {ids: ids},
        success: function (data) {
            if (data.res == 1) {
                if (url) {
                    location.href = url;
                } else {
                    location.reload()
                }
            } else {
                alert('删除失败！');
            }
        }
    });
}
//批量删除图片
function imgDels() {
    var ids = [];
    $('input[name^="ids"]:checked').each(function () {
        ids.push(this.value);
    });
    imgDel(ids.toString());
}
//删除广告
function adDel(ids, url) {
    if (!ids) {
        alert('请勾选需要操作的选项！');
        return;
    }
    if (!confirm('请确定要删除吗' + ids + '？')) {
        return;
    }
    $.ajax({
        type: 'POST',
        url: 'ad_del.html',
        data: {ids: ids},
        success: function (data) {
            if (data.res == 1) {
                if (url) {
                    location.href = url;
                } else {
                    location.reload()
                }
            } else {
                alert('删除失败！');
            }
        }
    });
}
//批量删除图片
function adDels() {
    var ids = [];
    $('input[name^="ids"]:checked').each(function () {
        ids.push(this.value);
    });
    adDel(ids.toString());
}
/**
 * 选择下载点击
 */
function selectDownload(isMulti) {
    var multi = '';
    if (isMulti) {
        multi = '?multi=1';
    }
    layer.open({
        type: 2,
        title: '选择下载',
        area: ['740px', multi ? '610px' : '665px'],
        fixed: true, //不固定
        shadeClose: true,
        content: '/download/download_select.html' + multi
    });
}
function multiDownloadSelected(dataArr) {
    var len = dataArr.length;
    console.log(dataArr)
    for (var i = 0; i < len; i++) {
        var data = dataArr[i];
        $('#bind-downloads').append('<li style="margin-bottom:5px;"><input type="hidden" name="dids[]" value="' + data.id + '"><a class="button bg-main" href="/xiazai/' + data.id + '.html" target="_blank">' + data.name + '-' + data.xtype + '</a><span class="close" onclick="$(this).parent().remove()"></span></li>');
    }
}
/**
 * 下载列表点击选择
 * @param id
 * @param name
 */
function downloadSelected(id, name) {
    $('#did').val(id);
    $('#xzname').val(name);
    layer.closeAll();
}
/**
 * 书籍作者点击选择
 * @param id
 * @param name
 */
function book_author_selected(id, name) {
    $('#book_author_id').val(id);
    $('#book_author').val(name);
    layer.closeAll();
}
/**
 * 书籍作者
 */
function select_book_author() {
    layer.open({
        type: 2,
        title: '选择作者',
        area: ['740px', '665px'],
        fixed: false, //不固定
        shadeClose: true,
        content: '/book/author_select.html'
    });
}
/**
 * 书籍版权点击选择
 * @param id
 * @param name
 */
function book_copyright_selected(id, name) {
    $('#copyright_id').val(id);
    $('#copyright_name').val(name);
    layer.closeAll();
}
/**
 * 书籍版权
 */
function select_book_copyright() {
    layer.open({
        type: 2,
        title: '选择版权来源',
        area: ['740px', '665px'],
        fixed: false, //不固定
        shadeClose: true,
        content: '/book/copyright_select.html'
    });
}

/**
 * 选择开发商点击
 * @param type
 */
function selectDeveloper(type) {
    var content = '/game/developer_select.html';
    if (type == 'download') {
        content = '/download/developer_select.html';
    } else if (type == 'server') {
        content = '/server/developer_select.html';
    }
    layer.open({
        type: 2,
        title: '请选择开发商',
        area: ['740px', '605px'],
        fixed: false, //不固定
        shadeClose: true,
        content: content
    });
}
/**
 * 开发商列表点击选择
 * @param id
 * @param dev
 */
function setDeveloper(id, dev) {
    $('#devid').val(id);
    $('#developer').val(dev);
    layer.closeAll();
}
/**
 * 添加地址点击
 */
function addAddress() {
    layer.open({
        type: 2,
        title: '请选择下载地址',
        area: ['1200px', '610px'],
        fixed: false, //不固定
        shadeClose: true,
        content: '/download/download_address_select.html'
    });
}
/**
 * 选择下载地址回调
 * @param addrArr
 */
function selectAddressCallback(addrArr) {
    var len = addrArr.length;
    for (var i = 0; i < len; i++) {
        var addr = addrArr[i];
        $('#down_address').append('<li style="margin-bottom:5px;"> <input type="hidden" name="aids[]" value="' + addr.id + '"> <a class="button bg-main" href="' + addr.durl + '" target="_blank">' + addr.dname + '-' + addr.tname + '</a> <span class="close" onclick="$(this).parent().remove()"></span> </li>');
    }
}
/**
 * 下载地址选择窗，添加下载地址
 * @param obj
 */
function setDownloadAddress(obj) {
    $('#down_address').append('<li style="margin-bottom:5px;"> <input type="hidden" name="aids[]" value="' + obj.id + '"> <a class="button bg-main" href="' + obj.durl + '" target="_blank">' + obj.dname + '-' + obj.tname + '</a> <span class="close" onclick="$(this).parent().remove()"></span> </li>');
}
function openPackUrl(value, index) {
    if (value) {
        if (index == 1) {
            window.open('http://app.mi.com/details?id=' + value)
        } else if (index == 2) {
            window.open('http://android.myapp.com/myapp/detail.htm?apkName=' + value)
        }
    }
}
$(function () {
    if (typeof devUrl != 'undefined') {
        if (!$('#devid').val()) {
            $('#developer').val('');
        }
        var _suggestions = [];
        $('#developer').autocomplete({
            serviceUrl: devUrl,
            minChars: 0,
            onSearchComplete: function (query, suggestions) {
                _suggestions = suggestions;
            },
            onSelect: function (suggestion) {
                $('#devid').val(suggestion.data);
            },
        });
        $('#developer').blur(function () {
            if ($('#developer').val() != selectDevVal) {
                selectDevVal = '';
                $('#devid').val('');
                $('#developer').val('')
            }

            setTimeout(function () {
                if (_suggestions.length == 0) {
                    if (!confirm('没有匹配的开发商，是否重新选择')) {
                        $('#devid').val('');
                        $('#developer').val('')
                    } else {
                        $('#developer').focus();
                    }
                    return;
                }
                var val = $('#developer').val();
                if (!val) {
                    return;
                }
                //遍历_suggestions数组中是否包含当前文本框内容
                var flag = false;
                $.each(_suggestions, function (index, value) {
                    if (value.value == val) {
                        flag = true;
                    }
                })
                if (!flag) {
                    if (!confirm('没有匹配的开发商，是否重新选择')) {
                        $('#devid').val('');
                        $('#developer').val('')
                    } else {
                        $('#developer').focus();
                    }
                    return;
                }
            }, 200);
        });
    }
    if (typeof gameUrl != 'undefined') {
        if (!$('#gid').val()) {
            $('#gname').val('');
        }
        var _suggestions = [];
        $('#gname').autocomplete({
            serviceUrl: gameUrl,
            minChars: 0,
            onSearchComplete: function (query, suggestions) {
                _suggestions = suggestions;
            },
            onSelect: function (suggestion) {
                $('#gid').val(suggestion.data);
            }
        });
        $('#gname').blur(function () {
            setTimeout(function () {
                if (_suggestions.length == 0) {
                    if (!confirm('没有匹配的游戏，是否重新选择')) {
                        $('#gid').val('');
                        $('#gname').val('')
                    } else {
                        $('#gname').focus();
                    }
                    return;
                }
                var val = $('#gname').val();
                if (!val) {
                    return;
                }
                //遍历_suggestions数组中是否包含当前文本框内容
                var flag = false;
                $.each(_suggestions, function (index, value) {
                    if (value.value == val) {
                        flag = true;
                    }
                })
                if (!flag) {
                    if (!confirm('没有匹配的游戏，是否重新选择')) {
                        $('#gid').val('');
                        $('#gname').val('')
                    } else {
                        $('#gname').focus();
                    }
                    return;
                }
                typeof onPageSelect != 'undefined' && onPageSelect();
            }, 200);
        });
    }
    if (typeof downUrl != 'undefined') {
        if (!$('#did').val()) {
            $('#xzname').val('');
        }
        var _suggestions = [];
        $('#xzname').autocomplete({
            serviceUrl: downUrl,
            minChars: 0,
            onSearchComplete: function (query, suggestions) {
                _suggestions = suggestions;
            },
            onSelect: function (suggestion) {
                $('#did').val(suggestion.data);
            },
        });
        $('#xzname').blur(function () {
            setTimeout(function () {
                if (_suggestions.length == 0) {
                    if (!confirm('没有匹配的下载，是否重新选择')) {
                        $('#did').val('');
                        $('#xzname').val('')
                    } else {
                        $('#xzname').focus();
                    }
                    return;
                }
                var val = $('#xzname').val();
                if (!val) {
                    return;
                }
                //遍历_suggestions数组中是否包含当前文本框内容
                var flag = false;
                $.each(_suggestions, function (index, value) {
                    if (value.value == val) {
                        flag = true;
                    }
                });
                if (!flag) {
                    if (!confirm('没有匹配的下载，是否重新选择')) {
                        $('#did').val('');
                        $('#xzname').val('')
                    } else {
                        $('#xzname').focus();
                    }
                    return;
                }
            }, 500);
        });
    }
    if (typeof modul != 'undefined') {
        window.UEDITOR_CONFIG.initialFrameWidth = "100%";
        window.UEDITOR_CONFIG.initialFrameHeight = "300";
        window.UEDITOR_CONFIG.imageUrl = imageUrl;
        window.UEDITOR_CONFIG.imagePath = imagePath;
        window.UEDITOR_CONFIG.imageManagerUrl = imageManagerUrl;
        window.UEDITOR_CONFIG.imageManagerPath = imageManagerPath;
        if (modul == 'download') {
            $('#relativeButton').click(function () {
                var key = $('#rkey').val();
                if (!key) {
                    alert('请输入关键字查找！');
                    return;
                }
                $("#uploadPopup p").text('正在查找');
                $("#uploadPopup,.uploadBackdrop").show();
                $.post("/special/ajax_relate.html", {key: key},
                    function (data) {
                        $("#uploadPopup,.uploadBackdrop").hide();
                        if (data) {
                            $("#TempInfoList").empty().append(data);
                        } else {
                            alert('没有检索到相关信息！')
                        }
                    }
                );
            });
            $('#RAddButton').click(function () {
                var alloptions = $("#TempInfoList option");
                var so = $("#TempInfoList option:selected");
                var a = (so.get(so.length - 1).index == alloptions.length - 1) ? so.prev().attr("selected", true) : so.next().attr("selected", true);
                var isHave = false;
                $("#SelectInfoList").children().each(function () {
                    if (so.val() == $(this).val()) {
                        isHave = true;
                    }
                });
                if (!isHave) {
                    $("#SelectInfoList").append(so);
                } else {
                    so.remove();
                }
            });
            $('#RAddMoreButton').click(function () {
                $("#TempInfoList option").each(function () {
                    if ($("#SelectInfoList option[value=" + $(this).val() + "]").attr("selected")) {
                        $(this).remove()
                    }
                });
                var so = $("#TempInfoList option").attr("selected", true);
                $("#SelectInfoList").append(so);
            });
            $('#RDelButton').click(function () {
                var alloptions = $("#SelectInfoList option");
                var so = $("#SelectInfoList option:selected");
                var a = (so.get(so.length - 1).index == alloptions.length - 1) ? so.prev().attr("selected", true) : so.next().attr("selected", true);
                so.each(function () {
                    var stext = $(this).text();
                    var sval = $(this).val();
                    if (stext.indexOf('↓') != -1) {
                        $(this).text(stext.split('↓')[1]);
                        $(this).val(sval.split('↓')[0]);
                    }
                });
                so.remove();
                $("#TempInfoList").append(so);
            });
            $('#RDelMoreButton').click(function () {
                var so = $("#SelectInfoList option");
                so.each(function () {
                    var stext = $(this).text();
                    var sval = $(this).val();
                    if (stext.indexOf('↓') != -1) {
                        $(this).text(stext.split('↓')[1]);
                        $(this).val(sval.split('↓')[0]);
                    }
                });
                $("#TempInfoList").append(so);
            });
            UE.getEditor('descrip');
        } else if (modul == 'game') {
            UE.getEditor('gdescript');
        } else if (modul == 'article') {
            UE.getEditor('content');
        }
    }
    //排序功能
    if($('.sso').length > 0){
        var order = $("input[name='order']");
        var order_val = order.val();
        var direction = $("input[name='direction']");
        var direction_val = direction.val();
        var oos_html = '<i class="down"></i>';
        if(direction_val==1) oos_html = '<i class="up"></i>';
        $('.sso').each(function(){
            if($(this).attr('v')==order_val){
                $(this).append(oos_html);
            }
        }).bind('click',function(){
            var v = $(this).attr('v');
                order.val(v);
            if(v==order_val){
                if(direction_val=='1'){
                    direction.val('0');
                }else{
                    direction.val('1');
                }
            }else{
                direction.val('1');
            }
            //提交表单
            to_sso();
        });
    }
});

function addcreateid()
{
	// jQuery.getJSON('http://api.qirexiaoshuo.com/data/create?cname=渠道用户&callback=?', function(data){
	// 	if(data.result=="")
	// 	{
	// 		alert('新建ID失败，请重新创建');
	// 	}
	// 	else
	// 	{
	// 		document.getElementById("extendid").value=data.result;
	// 	}
	// });
	
}





