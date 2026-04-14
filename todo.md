# todo

## [x] 添加list命令
用于显示当天的记录内容。
daydaynote list
索引 记录时间 记录内容

添加--date flag
用于显示具体日期中的记录内容
daydaynote list --date=2020/4/21

## [x] 添加delete命令
用于删除当天的记录内容
daydaynote delete 1

添加--date flag
用于删除具体日期中的具体索引的记录内容
daydaynote delete 1 --date=2020/4/21或者2020-4-21 两种日期格式

## [x] 添加edit命令
用于编辑当天的记录内容
daydaynote edit 1 content
1是索引，content是要编辑的内容

加上--date用于指定具体日期
