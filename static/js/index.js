$(document).ready(function() {
    // var $ = function(selector) {
    //     return document.querySelector(selector)
    // }

    var bindEventAdd = function() {
        // 给 add button 绑定点击事件
        $('.todo-add').on('click', function() {
            let value = $('.todo-input').val()
            $('.todo-input').val("")
            let data = {
                notes: value,
            }
            handlePostRequest(data)
        })
    }

    var bindEventDelete = function() {
        // 通过事件委托的方式，将点击事件绑定到父元素上，当点击删除按钮时，通过事件对象获取到点击的具体按钮，
        // 再通过按钮的自定义属性获取到对应的 todo 标识符
        $('.todo-list').on('click', function(event) {
            // 判断点击的目标元素是否为删除按钮
            if (event.target.classList.contains('todo-delete')) {
                // 获取点击的按钮
                var deleteBtn = event.target;
                // 获取对应的 todo 标识符
                var todoId = deleteBtn.dataset.todoId;

                console.log("todoId = ", todoId);
                // $.ajax({})
            }
        })
    }

    var handlePostRequest = function(data) {
        // 向后台发送数据
        $.ajax({
            type: 'POST',
            url: "/todo",
            data: data,
            success: function(response) {
                console.log(response)
                // 当创建成功后，将调用 GET 请求，去获取 todo 列表更新画面
                handleGetRequest()
            },
            error: function(xhr, status, error) {
                console.log(error)
            },
        })
    }

    var handleGetRequest = function() {
        // 获取后台数据
        $.ajax({
            type: "GET",
            url: "/todo",
            success: function (response) {
                console.log(response);
                let template = ''
                for (let i = 0; i < response.length; i++) {
                    let d = response[i]
                    template += `
                        <div class="todo">
                            id: ${d.ID}
                            notes: ${d.Notes}
                            <button class="todo-delete" data-todo-id="${d.ID}">删除</button>
                        </div>
                    `

                }
                $('.todo-list').html(template)
            },
            error: function (xhr, status, error) {
                console.log(error);
            }
        });
    }

    var bindEvents = function() {
        // 添加 todo
        bindEventAdd()

        // 删除 todo
        bindEventDelete()
    }

    var __main = function() {
        // 绑定事件
        bindEvents()

        // 在网页渲染完成后，获取后台保存的数据
        handleGetRequest()
    }

    __main()
});