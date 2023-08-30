(function() {
    var $ = function(selector) {
        return document.querySelector(selector)
    }

    var bindEventAdd = function() {
        // 给 add button 绑定点击事件
        $('.todo-add').addEventListener('click', function() {
            let value = $('.todo-input').value
            console.log("value = ", value);
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