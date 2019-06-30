环境变量
====================

通过环境变量来配置，将会自动读取当前工作目录的 ``.env`` 文件。

EMAIL_HOST

    邮件主机

EMAIL_PORT

    邮件端口

EMAIL_USER

    邮件主机上的用户名，默认为 ``EMAIL_FROM``。

EMAIL_PASSWORD

    邮件密码

EMAIL_FROM

    邮件来自地址。

EMAIL_TO

    邮件发往地址， 默认为 ``EMAIL_FROM``。

ROOM_LINKS

    空格分隔的房间详情页URL列表。

    例子: ``http://www.ziroom.com/z/vr/61819181.html http://m.ziroom.com/BJ/room?id=61819181``

POLL_INTERVAL

    默认值: 30 秒

    房间信息轮询间隔，使用 golang 时间长度字符串，例如: ``1h2m3s``。

DATABASE_DIR

    默认值: ``/tmp/ziroom-ob/db``

    数据库文件夹。
