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

    var handlePostRequest = function(data) {
        // 向后台发送数据
        $.ajax({
            type: 'POST',
            url: "/todolist",
            data: data,
            success: function(response) {
                console.log(response)
                // 当创建成功后，将调用 GET 请求，去获取 todolist 列表更新画面
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
            url: "/todolist",
            success: function (response) {
                console.log(response);
                let template = ''
                for (let i = 0; i < response.length; i++) {
                    let item = response[i]
                    template += `
                        <div class="item">
                            id: ${item.ID} notes: ${item.Notes}
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
    }

    var __main = function() {
        // 绑定事件
        bindEvents()

        // 在网页渲染完成后，获取后台保存的数据
        handleGetRequest()
    }

    __main()
});