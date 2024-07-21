# smart-farming

项目结构：
├───.idea
├───biz
│   ├───config
│   ├───dal // 数据库表单
│   │   ├───model
│   │   └───query
│   ├───handler // 处理器 (根据api添加文件夹，可以参考employees)
│   │   └───employees
│   └───routers // 路由层 (根据自己的api添加文件夹)
│       └───employees
├───cmd
│   ├───generate // go gen
│   └───smart-farming // 入口
└───pkg
└───constants // 全局配置