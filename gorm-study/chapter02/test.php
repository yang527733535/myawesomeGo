3月27日 10:27


已读

旧的路径是传这个给你吗
已读
3月27日 11:37

项目的接口文档的地址发给我一下
已读

我这电脑没存
已读

ShowDoc
一个非常适合IT团队的在线API文档、技术文档工具。你可以使用Showdoc来编写在线API文档、技术文档、数据字典、在线手册

http://showdoc.stardustworld.cn/web/#/item/index
3月27日 14:01


已读

删除目录的
已读
3月27日 14:03


已读
3月27日 14:27



等會加一下这两个字段


3月27日 15:37


已读
3月30日 10:58

梁观锦和杨腾辉的聊天记录
梁观锦:[图片]
梁观锦:这个绑定不一致。
已读

得问兴宇

这个的人物头像和图片不一致
已读

ok
已读

掉接口太快了估计
3月31日 16:53

hello.exe
2.1 MB
已查收
打开文件夹
已读

chapter01.exe
14.1 MB
已查收
打开文件夹
已读
刚刚

17:41
    public function show(int $book_id, $role_name, $is_missing_art_status, $exitor_keys)
    {
        $select = ['id', 'name', 'is_leading', 'role_image', 'describe', 'is_missing_art_status', 'is_import', 'import_time', 'import_user_id', 'import_user_name', 'my_visual_id', 'photo', 'item_inherit_items', 'item_inherit_role_id'];
        $select = array_merge($select, $exitor_keys);

        $where = [
            ["book_id", $book_id],
            ['is_deleted', Roles::DELETED_OFF]
        ];

        if ($is_missing_art_status == 1) {
            $where[] = ['is_missing_art_status', 1];
        }

        if ($role_name) {
            $where[] = ['name', 'like', '%' . $role_name . '%'];
        }

        $role_list = Roles::select($select)->where($where)->orderBy('is_leading', 'desc')->orderBy('touch_time', 'desc')->get();

        foreach ($role_list as &$val) {
            $val->has_visual = empty($val->my_visual_id) ? 0 : 1;
            $val->has_photo  = empty($val->photo) ? 0 : 1;
            //unset($val->my_visual_id);
            unset($val->photo);
        }

        return $role_list;
    }
