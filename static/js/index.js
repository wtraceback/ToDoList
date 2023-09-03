$(document).ready(function() {
    // var $ = function(selector) {
    //     return document.querySelector(selector)
    // }

    var bindEventAdd = function() {
        // 给 add button 绑定点击事件
        $('.todo-add').on('click', function() {
            handleInputValue()
        })
    }

    // 给 input 输入框按回车键时添加事件
    var bindEventInputEnter = function() {
        $('.todo-input').on('keydown', function(event) {
            if (event.key === 'Enter') {
                // 失去焦点
                event.target.blur()
                // 阻止默认行为的发生
                event.preventDefault()

                handleInputValue()
            }
        })
    }

    var handleInputValue = function() {
        let value = $('.todo-input').val().trim()

        // 输入数据的判断
        if (value == '') {
            alert("内容不能为空")
        } else {
            let data = {
                notes: value,
            }

            // 向后台发送数据
            $.ajax({
                type: 'POST',
                url: "/todo",
                data: data,
                success: function(response) {
                    // 当创建成功后，将调用 GET 请求，去获取 todo 列表更新画面
                    rerenderList()
                    $('.todo-input').val("")
                    $('.todo-input').focus()
                },
                error: function(xhr, status, error) {
                    console.log(error)
                },
            })
        }
    }

    var bindEventDelete = function() {
        // 通过事件委托的方式，将点击事件绑定到父元素上，当点击删除按钮时，通过事件对象获取到点击的具体按钮，
        // 再通过按钮的自定义属性获取到对应的 todo 标识符
        $('.todo-list').on('click', function(event) {
            // 获取点击的按钮
            var target = event.target

            // 判断点击的目标元素是否为删除按钮
            if (target.classList.contains('todo-delete')) {
                // 获取对应的 todo 标识符
                var todoId = target.dataset.todoId

                console.log("todoId = ", todoId)
                $.ajax({
                    type: 'DELETE',
                    url: `/todo/${todoId}`,
                    success: function(response) {
                        rerenderList()
                    },
                    error: function(xhr, status, error) {
                        console.log(error)
                    },
                })
            }
        })
    }

    var bindEventTextBlur = function() {
        // 使用捕获模式
        document.querySelector('.todo-list').addEventListener('blur', function(event) {
            var target = event.target

            if (target.classList.contains('todo-task')) {
                console.log("text blur");
                // 把元素在 todo_list 中更新
                var todoId = target.dataset.todoId
                var notes = target.innerHTML

                console.log("todoId = ", todoId)
                $.ajax({
                    type: 'PUT',
                    url: `/todo/${todoId}`,
                    data: { notes: notes },
                    success: function(response) {
                        rerenderList()
                    },
                    error: function(xhr, status, error) {
                        console.log(error)
                    },
                })
            }
        }, true)
    }

    var bindEventTextEnter = function() {
        document.querySelector('.todo-list').addEventListener('keydown', function(event){
            var target = event.target;

            // 编辑完了后点击了回车键
            if(event.key === 'Enter' && target.classList.contains('todo-task')) {
                // 失去焦点
                target.blur();
                // 阻止默认行为的发生, 即不插入回车
                event.preventDefault();
            }
        })
    }

    var rerenderList = function() {
        // 获取后台数据
        $.ajax({
            type: "GET",
            url: "/todo",
            success: function (response) {
                let template = ''
                for (let i = 0; i < response.length; i++) {
                    let d = response[i]
                    template += `
                        <div class="todo">
                            <span>id: ${d.ID}</span>
                            <p class="todo-task"  contenteditable="plaintext-only" data-todo-id="${d.ID}">${d.Notes}<p>
                            <button class="todo-delete" data-todo-id="${d.ID}">删除</button>
                        </div>
                    `

                }
                $('.todo-list').html(template)
            },
            error: function (xhr, status, error) {
                console.log(error)
            }
        })
    }

    var bindEvents = function() {
        // 添加 todo
        bindEventAdd()

        // todo input 按回车时保存
        bindEventInputEnter()

        // 删除 todo
        bindEventDelete()

        // 文本框失去焦点后保存 todo
        bindEventTextBlur()

        // 任务栏的文本框输入 todo 按回车保存
        bindEventTextEnter()
    }

    var __main = function() {
        // 绑定事件
        bindEvents()

        // 在网页渲染完成后，获取后台保存的数据
        rerenderList()
    }

    __main()
});