自如 API
===============

使用自如手机版的 API: ``GET http://m.ziroom.com/wap/detail/room.json``

参数:

city_code

    城市码，通过码表对应。

id

    房间ID

相应值:

status

    类型: string

    操作状态，成功时应为 ``success``。

error_code

    类型: int

    错误码，成功时应为 ``0``。

error_message

    类型: string

    错误信息， 成功时应为空字符串。

data.id

    类型: string

    房间 ID , 应为查询时使用房间ID。

data.code

    类型: string

    房间编码，自如全站唯一。

data.house_code

    类型: string

    房屋编码，自如全站唯一。

data.city_code

    类型: string

    城市编码，应为查询时使用的城市编码。

data.status

    类型: string

    房间状态代码，通过码表对应。

data.name

    类型: string

    房间名称。

data.notice_word

    类型: string | undefined

    房间额外状态，空气检测中的房子值应为 ``空气检测中``
