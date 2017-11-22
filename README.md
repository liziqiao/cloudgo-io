#（直接复制了我的博客中的部分内容）

设计一个 web 小应用，展示静态文件服务、js 请求支持、模板输出、表单处理。

--------------------

##实现的功能

- **支持静态文件服务**
- **支持简单 js 访问**
- **提交表单，并输出一个表格**
- **对 ` /unknown`  给出开发中的提示，返回码  `5xx` **

--------------------

具体的代码见github:
https://github.com/liziqiao/cloudgo-io

-------------------


###支持静态文件服务

实现静态文件服务主要的就是先不要在assets文件夹中放index.html。
使用`go run main.go` 运行服务端，然后在浏览器访问localhost:8080/static/就可以看到下面的文件目录，点击相应的文件就可以看到文件内容

![这里写图片描述](http://img.blog.csdn.net/20171122163737181?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvc3lzdWx6cQ==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


###支持简单 js 访问

在之前的代码中的assets中添加index.html，再次访问localhost:8080/static/可以看到经过js代码处理过的页面：

![这里写图片描述](http://img.blog.csdn.net/20171122174658823?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvc3lzdWx6cQ==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

###提交表单，并输出一个表格

访问http://localhost:8080/register 会返回一个注册的表单如下：

![这里写图片描述](http://img.blog.csdn.net/20171122174925669?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvc3lzdWx6cQ==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

填写好表单就可以得到一个返回的表格：

![这里写图片描述](http://img.blog.csdn.net/20171122175015035?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvc3lzdWx6cQ==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


###对 ` /unknown`  给出开发中的提示，返回码  `5xx`

访问http://localhost:8080/api/unknown 就可以得到一个501 unknown返回：

![这里写图片描述](http://img.blog.csdn.net/20171122175214705?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvc3lzdWx6cQ==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
