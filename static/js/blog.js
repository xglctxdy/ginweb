$(document).ready(function () {
    //注册
    $("#register-form").validate({
        rules: {
            username: {
                required: true,
                rangelength: [5, 10]
            },
            email:{
                required: true,
                rangelength: [10,20]
            },
            password: {
                required: true,
                rangelength: [5, 10]
            },
            repassword: {
                required: true,
                rangelength: [5, 10],
                equalTo: "#register-password"
            }
        },
        messages: {
            username: {
                required: "请输入用户名",
                rangelength: "用户名必须是5-10位"
            },
            email: {
                required :"请输入邮箱",
                rangelength: "邮箱必须是10-20位"
            },
            password: {
                required: "请输入密码",
                rangelength: "密码必须是5-10位"
            },
            repassword: {
                required: "请确认密码",
                rangelength: "密码必须是5-10位",
                equalTo: "两次输入的密码必须相等"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/register";
            $(form).ajaxSubmit({
                url: urlStr,
                type: "POST",
                dataType: "json",
                // 只要成功传输都会进入success
                success: function (data, status) {
                    alert("提示:" + data.message);
                    if (data.code == 0) {
                        // 注册成功跳转到登录页面
                        setTimeout(function () {
                            window.location.href = "/login";
                        }, 1000);
                    }
                    else {
                        // 注册失败刷新注册页
                        window.location.href = "/register";
                    }
                },
                // 由于其他问题导致传向后端时失败
                error: function (data, status) {
                    alert("发生错误:" + data.message + ":" + status);
                }
            });
        }
    });

    $("#login-form").validate({
        rules: {
            email: {
                required: true,
                rangelength: [10, 20]
            },
            password: {
                required: true,
                rangelength: [5, 10]
            }
        },
        messages: {
            username: {
                required: "请输入邮箱地址",
                rangelength: "邮箱地址必须是5-10位"
            },
            password: {
                required: "请输入密码",
                rangelength: "密码必须是5-10位"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/login"
            $(form).ajaxSubmit({
                url: urlStr,
                type: "POST",
                dataType: "json",
                success: function (data, status) {
                    alert("提示:" + data.message)
                    if (data.code == 0) {
                        // 成功登录，随即自动跳转到首页
                        setTimeout(function () {
                            window.location.href = "/"
                        }, 1000)
                    }
                    else {
                        // 登录失败刷新登录页面
                        window.location.href = "/login"
                    }
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            });
        }
    });

    //修改和添加文章的表单
    $("#write-art-form").validate({
        rules: {
            title: "required",
            content: {
                required: true,
                minlength: 2
            }
        },
        messages: {
            title: "请输入标题",
            content: {
                required: "请输入文章内容",
                minlength: "文章内容最少两个字符"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/question/add";
            //判断文章id确定提交的表单的服务器地址
            //若id大于零，说明是修改文章
            var artId = $("#write-question-id").val();
            // alert("artId:" + artId);
            if (artId > 0) {
                urlStr = "/question/update"
            }
            alert("urlStr:" + urlStr);
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert(":data:" + data.message);
                    setTimeout(function () {
                        window.location.href = "/"
                    }, 1000)
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            });
        }
    });


    //文件
    $("#album-upload-button").click(function () {
        var filedata = $("#album-upload-file").val();
        if (filedata.length <= 0) {
            alert("请选择文件!");
            return
        }
        //文件上传通过Formdata去储存文件的数据
        var data = new FormData()
        data.append("upload", $("#album-upload-file")[0].files[0]);
        alert(data)
        var urlStr = "/upload"
        $.ajax({
            url: urlStr,
            type: "post",
            dataType: "json",
            contentType: false,
            data: data,
            processData: false,
            success: function (data, status) {
                alert(":data:" + data.message);
                if (data.code == 1) {
                    setTimeout(function () {
                        window.location.href = "/album"
                    }, 1000)
                }
            },
            error: function (data, status) {
                alert("err:" + data.message + ":" + status)
            }
        })
    })
});
