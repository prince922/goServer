var page = {}

page.limit = 100;

/**
 * 初始页数信息
 * @param pageJson
 */
page.init = function (pageJson) {
    if (typeof pageJson != 'object'){
        pageJson = JSON.parse(pageJson);
    }
    page.now = pageJson.PageNow;
    page.count = pageJson.PageCount;
    page.lists = pageJson.Index;
    page.size = pageJson.PageSize;
    this.getHostInfo();
};

/**
 * 获取当前URL信息
 */
page.getHostInfo = function(){
    page.host = window.location.host;
    page.path = window.location.pathname;
    page.search = window.location.search;
};

page.Create = function (params) {
    page.paramSize = params.size || 10;

    var lists = this.showList() , PrevIndex = this.showPrevAndNext(), FirstIndex = this.showFirstAndLast();

    var html = "<ul class=\"pager\" style=\"display: inline\">"+
                    FirstIndex.pageFirst+
                    PrevIndex.prev+
                    lists+
                    PrevIndex.next+
                    FirstIndex.pageLast+
                "</ul>";
    $(params.dom).html(html);
};

/**
 * 生成当前URL
 * @param page
 * @returns {*}
 */
page.createHost = function(num){
    params = [{'arg':'page','value':num},{'arg':'size', 'value':this.paramSize}];
    url = this.path + this.search
    return this.changeUrlArg(url, params)
}

/**
 * 重组url
 * @param url
 * @param params
 * @returns {*}
 */
page.changeUrlArg = function changeURLArg(url,params){
    for(var i in params){
        var pattern=params[i].arg + '=([^&]*)';

        temp = params[i].arg + '=' + params[i].value;
        if(url.match(pattern)){
            var tmp='/(' + params[i].arg + '=)([^&]*)/gi';
            url=url.replace(eval(tmp),temp);
        } else {
            if(url.match('[\?]')){
                url += '&' + temp;
            } else {
                url += '?' + temp;
            }
        }
    }
    return url;
}

/**
 * 构建列表页
 * @returns {string}
 */
page.showList = function () {
    var html = "";


    for (var i in page.lists){
        // 页数最大限制
        if(page.lists[i] > page.limit){
            break;
        }
        if(page.now == page.lists[i]){
            html += "<li>" +
                        "<a class=\"active\">" + page.lists[i] + "</a>"+
                    "</li>";
        }else {
            html += "<li>"+
                        "<a href=\"" + this.createHost(page.lists[i]) +"\">" + page.lists[i] + "</a>"+
                    "</li>";
        }
    }


    return html
};

/**
 * 上一页，下一页
 *
 * @returns {{prev: string, next: string}}
 */
page.showPrevAndNext = function () {
    var info = {
        "prev" : '',
        "next" : ''
    }

    if(page.now == 1){
        info.prev = "<li>"+
                        "<a href=\"javascript:void(0);\" target=\"_self\">上一页</a>\n"+
                    "</li>";
    } else {
        info.prev = "<li>"+
                        "<a href=\"" + this.createHost(page.now-1) +"\">上一页</a>"+
                    "</li>";
    }

    count = page.count > page.limit ? page.limit : page.count;

    if(page.now == count){
        info.next = "<li>"+
                        "<a href=\"javascript:void(0);\" target=\"_self\">下一页</a>\n"+
                    "</li>";
    } else {
        info.next = "<li>"+
                        "<a href=\"" + this.createHost(page.now+1) +"\">下一页</a>"+
                    "</li>";
    }

    return info;
}


/**
 * 首页，尾页
 *
 * @returns {{prev: string, next: string}}
 */
page.showFirstAndLast = function () {
    var info = {
        "pageFirst" : '',
        "pageLast" : ''
    }

    if(page.now == 1){
        info.pageFirst = "<li>"+
            "<a href=\"javascript:void(0);\" target=\"_self\">首页</a>\n"+
            "</li>";
    } else {
        info.pageFirst = "<li>"+
            "<a href=\"" + this.createHost(1) +"\">首页</a>"+
            "</li>";
    }

    count = page.count > page.limit ? page.limit : page.count;

    if(page.now == count){
        info.pageLast = "<li>"+
            "<a href=\"javascript:void(0);\" target=\"_self\">尾页</a>\n"+
            "</li>";
    } else {
        info.pageLast = "<li>"+
            "<a href=\"" + this.createHost(count) +"\">尾页</a>"+
            "</li>";
    }

    return info;
}


