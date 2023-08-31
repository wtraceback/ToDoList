(function() {
    // var $ = function(selector) {
    //     return document.querySelector(selector)
    // }

    var bindEventAdd = function() {
        // 给 add button 绑定点击事件
        $('.todo-add').on('click', function() {
            let value = $('.todo-input').val()
            let data = {
                notes: value,
            }
            handlePostRequest(data)
        })
    }

    var handlePostRequest = function(data) {
        // 向后台发送数据
        $.ajax({
            type: 'POST',
            url: "/todolist",
            data: data,
            success: function(response) {
                console.log(response)
            },
            error: function(xhr, status, error) {
                console.log(error)
            },
        })
    }

    var bindEvents = function() {
        // 添加 todo
        bindEventAdd()
    }

    var __main = function() {
        // 绑定事件
        bindEvents()
    }

    __main()
})();