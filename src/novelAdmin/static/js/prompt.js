/**
 * 修改tp消息输出方式
 */
$(function(){
    //ajax get请求
    $('.ajax-get').click(function(){
        var target;
        var that = this;
        if ( $(this).hasClass('confirm') ) {
            if(!confirm('确认要执行该操作吗?')){
                return false;
            }
        }
        if ( (target = $(this).attr('href')) || (target = $(this).attr('url')) ) {
            $.get(target).success(function(data){
                if (data.code==1) {
                    if (data.url) {
                        layer.msg(data.msg + ' 页面即将自动跳转~', {time: 2500, icon:6});
                    }else{
                        layer.msg(data.msg , {time: 2500, icon:6});
                    }
                    setTimeout(function(){
                        if (data.url) {
                            location.href=data.url;
                        }else{
                            location.reload();
                        }
                    },2500);
                }else{
                    layer.msg(data.msg , {time: 2500, icon:5});
                    setTimeout(function(){
                        if (data.url) {
                            location.href=data.url;
                        }
                    },2500);
                }
            },'json');
        }
        return false;
    });

    //ajax post submit请求
    $('.ajax-post').click(function(){
        var target,query,form;
        var target_form = $(this).attr('target-form');
        var that = this;
        var nead_confirm=false;
        if( ($(this).attr('type')=='submit') || (target = $(this).attr('href')) || (target = $(this).attr('url')) ){
            form = $('.'+target_form);
            if ($(this).attr('hide-data') === 'true'){//无数据时也可以使用的功能
                form = $('.hide-data');
                query = form.serialize();
            }else if (form.get(0)==undefined){
                return false;
            }else if ( form.get(0).nodeName=='FORM' ){
                if ( $(this).hasClass('confirm') ) {
                    if(!confirm('确认要执行该操作吗?')){
                        return false;
                    }
                }
                if($(this).attr('url') !== undefined){
                    target = $(this).attr('url');
                }else{
                    target = form.get(0).action;
                }
                query = form.serialize();
            }else if( form.get(0).nodeName=='INPUT' || form.get(0).nodeName=='SELECT' || form.get(0).nodeName=='TEXTAREA') {
                form.each(function(k,v){
                    if(v.type=='checkbox' && v.checked==true){
                        nead_confirm = true;
                    }
                });
                if ( nead_confirm && $(this).hasClass('confirm') ) {
                    if(!confirm('确认要执行该操作吗?')){
                        return false;
                    }
                }
                query = form.serialize();
            }else{
                if ( $(this).hasClass('confirm') ) {
                    if(!confirm('确认要执行该操作吗?')){
                        return false;
                    }
                }
                query = form.find('input,select,textarea').serialize();
            }
            $(that).addClass('disabled').attr('autocomplete','off').prop('disabled',true);
            $.post(target,query).success(function(data){
                if (data.code==1) {
                    if (data.url) {
                        layer.msg(data.msg + ' 页面即将自动跳转~', {time: 2500, icon:6});
                    }else{
                        layer.msg(data.msg , {time: 2500, icon:6});
                    }
                    setTimeout(function(){
                        $(that).removeClass('disabled').prop('disabled',false);
                        if (data.url) {
                            location.href=data.url;
                        }else{
                            location.reload();
                        }
                    },2500);
                }else{
                    layer.msg(data.msg , {time: 2500, icon:5});
                    setTimeout(function(){
                        $(that).removeClass('disabled').prop('disabled',false);
                        if (data.url) {
                            location.href=data.url;
                        }
                    },2500);
                }
            },'json');
        }
        return false;
    });
});

//Ajax 请求 json post
function AjaxJson(url, data, successFunc, errorFunc) {
    var data = data || "{}";
    successFunc = successFunc || null;
    errorFunc = errorFunc || null;
    if (!url || url === '#')
        return false;

    $.ajax({
        type: 'POST',
        url: url,
        data: data,
        dataType: 'json',
        success: function (data) {	 
            try {
                successFunc(data);
            } catch (e) {
            }
        },
        error: function (xhr, type) {
            try {
                errorFunc(xhr, type);
            } catch (e) {
            }
            console.log("页面动态加载不成功，请与管理员联系");
        },
    })
    //阻止冒泡
    return false;
}


//Ajax 请求 json post
function AjaxJsonP(url, data, successFunc, errorFunc) {
    var data = data || "{}";
    successFunc = successFunc || null;
    errorFunc = errorFunc || null;
    if (!url || url === '#')
        return false;

    $.ajax({
        type: 'POST',
        url: url,
        data: data,
        dataType: 'jsonp',
        success: function (data) {
            try {
                successFunc(data);
            } catch (e) {
            }
        },
        error: function (xhr, type) {
            try {
                errorFunc(xhr, type);
            } catch (e) {
            }
            console.log("页面动态加载不成功，请与管理员联系");
        },
    })
    //阻止冒泡
    return false;
}
