import React, { useState } from 'react';
import { Card, Form, Select, Input, Button, Table, Empty, Image, Breadcrumb, Switch, Modal, InputNumber, Upload, Radio } from 'antd';
import { Link } from 'react-router-dom';

type Cat = { id: number; name: string; icon?: string; status: 'show' | 'hide'; desc?: string; sort?: number; parentId?: number };

const initialData: Cat[] = [
  { id: 181, name: '院校介绍', icon: 'https://via.placeholder.com/40?text=A', status: 'show' },
  { id: 180, name: '潮流文化', icon: 'https://via.placeholder.com/40?text=B', status: 'show' },
  { id: 179, name: '品牌资讯', icon: 'https://via.placeholder.com/40?text=C', status: 'show' },
  { id: 182, name: '🎧分类', icon: 'https://via.placeholder.com/40?text=%F0%9F%8E%A7', status: 'show' },
];

const ArticleCategory: React.FC = () => {
  const [status, setStatus] = useState<string | undefined>();
  const [keyword, setKeyword] = useState<string>('');
  const [data, setData] = useState<Cat[]>(initialData);
  const [openAdd, setOpenAdd] = useState(false);
  const [form] = Form.useForm();
  const [openEdit, setOpenEdit] = useState(false);
  const [editForm] = Form.useForm();
  const [editing, setEditing] = useState<Cat | null>(null);

  const filtered = data.filter(item => {
    const byStatus = status ? (status === 'show' ? item.status === 'show' : item.status === 'hide') : true;
    const byKeyword = keyword ? item.name.includes(keyword) : true;
    return byStatus && byKeyword;
  });

  const columns = [
    { title: 'ID', dataIndex: 'id', width: 80 },
    { title: '分类名称', dataIndex: 'name' },
    { title: '分类图片', dataIndex: 'icon', render: (src: string) => {
      if (!src) return '-';
      if (src.startsWith('emoji:')) {
        const emoji = src.replace('emoji:', '');
        return <span style={{ fontSize: 24 }}>{emoji}</span>;
      }
      return <Image src={src} width={40} height={40} />;
    } },
    { title: '状态', dataIndex: 'status', width: 120, render: (_: any, record: Cat) => (
      <Switch
        checkedChildren="开启"
        unCheckedChildren="关闭"
        checked={record.status === 'show'}
        onChange={(checked) => setData(prev => prev.map(it => it.id === record.id ? { ...it, status: checked ? 'show' : 'hide' } : it))}
      />
    ) },
    { title: '操作', dataIndex: 'action', width: 200, render: (_: any, record: Cat) => (
      <div style={{ display: 'flex', gap: 8 }}>
        <Button type="link" onClick={() => onEdit(record)}>编辑</Button>
        <Button type="link" danger onClick={() => setData(prev => prev.filter(it => it.id !== record.id))}>删除</Button>
        <Button type="link">查看文章</Button>
      </div>
    ) },
  ];

  const toFileList = (url?: string) => (url ? [{ uid: '1', url, status: 'done', name: 'image' }] : []);

  const onEdit = (record: Cat) => {
    setEditing(record);
    editForm.setFieldsValue({
      parentId: record.parentId ?? 0,
      name: record.name,
      desc: record.desc,
      icon: toFileList(record.icon),
      sort: record.sort ?? 0,
      status: record.status,
    });
    setOpenEdit(true);
  };

  const onEditCancel = () => {
    setOpenEdit(false);
    editForm.resetFields();
    setEditing(null);
  };

  const onEditOk = async () => {
    const values = await editForm.validateFields();
    const file = values.icon?.[0];
    const iconUrl = file?.url || file?.thumbUrl || editing?.icon || undefined;
    setData(prev => prev.map(it => (it.id === (editing?.id ?? -1) ? {
      ...it,
      parentId: values.parentId,
      name: values.name,
      desc: values.desc,
      icon: iconUrl,
      sort: values.sort ?? 0,
      status: values.status,
    } : it)));
    onEditCancel();
  };

  const onAddOk = async () => {
    const values = await form.validateFields();
    const maxId = Math.max(0, ...data.map(it => it.id));
    const file = values.icon?.[0];
    const iconUrl = file?.url || file?.thumbUrl || 'https://via.placeholder.com/40?text=新';
    const newItem: Cat = {
      id: maxId + 1,
      name: values.name,
      icon: iconUrl,
      status: values.status,
      desc: values.desc,
      sort: values.sort ?? 0,
    };
    setData(prev => [newItem, ...prev]);
    form.resetFields();
    setOpenAdd(false);
  };

  return (
    <div>
      <Card>
        {/* 面包屑导航 */}
        <Breadcrumb style={{ marginBottom: 20 }}>
          <Breadcrumb.Item>
            <Link to="/home">首页</Link>
          </Breadcrumb.Item>
          <Breadcrumb.Item>内容管理</Breadcrumb.Item>
          <Breadcrumb.Item>文章分类</Breadcrumb.Item>
        </Breadcrumb>

        {/* 顶部筛选栏：与上传图布局一致 */}
        <Form layout="inline" style={{ background: '#f7f8fa', padding: 16, borderRadius: 8 }}>
          <Form.Item label="是否显示">
            <Select
              style={{ width: 180 }}
              placeholder="请选择"
              value={status}
              onChange={setStatus}
              options={[{ value: 'show', label: '显示' }, { value: 'hide', label: '隐藏' }]}
              allowClear
            />
          </Form.Item>
          <Form.Item label="分类名称">
            <Input
              style={{ width: 280 }}
              placeholder="请输入分类名称"
              value={keyword}
              onChange={(e) => setKeyword(e.target.value)}
            />
          </Form.Item>
          <Form.Item>
            <Button type="primary">查询</Button>
          </Form.Item>
        </Form>

        <div style={{ marginTop: 12, display: 'flex', justifyContent: 'flex-start' }}>
          <Button type="primary" size="small" onClick={() => setOpenAdd(true)}>添加文章分类</Button>
        </div>

        <div style={{ marginTop: 16 }}>
          <Table
            columns={columns}
            dataSource={filtered}
            pagination={false}
            locale={{ emptyText: <Empty description="暂无数据" /> }}
            rowKey="id"
          />
        </div>

        <Modal
          title="添加分类"
          open={openAdd}
          onOk={onAddOk}
          onCancel={() => { setOpenAdd(false); form.resetFields(); }}
          okText="确定"
          cancelText="取消"
          width={640}
          rootClassName="compact-modal"
          bodyStyle={{ padding: 12, maxHeight: '60vh', overflow: 'auto' }}
        >
          <Form
            form={form}
            layout="horizontal"
            labelCol={{ span: 6 }}
            wrapperCol={{ span: 18 }}
            requiredMark={true}
            initialValues={{ parentId: 0, status: 'show', sort: 0 }}
          >
            <Form.Item label="上级分类" name="parentId">
              <Select
                style={{ width: 240 }}
                options={[{ value: 0, label: '顶级分类' }, ...data.map(it => ({ value: it.id, label: it.name }))]}
              />
            </Form.Item>

            <Form.Item label="分类名称" name="name" rules={[{ required: true, message: '请输入分类名称' }]}> 
              <Input placeholder="请输入分类名称" />
            </Form.Item>

            <Form.Item label="分类简介" name="desc" rules={[{ required: true, message: '请输入分类简介' }]}> 
              <Input.TextArea placeholder="请输入分类简介" rows={3} />
            </Form.Item>

            <Form.Item label="分类图片" name="icon" valuePropName="fileList" getValueFromEvent={(e) => e?.fileList}>
              <Upload listType="picture-card" beforeUpload={() => false}>
                +
              </Upload>
            </Form.Item>

            <Form.Item label="排序" name="sort">
              <InputNumber min={0} style={{ width: 160 }} />
            </Form.Item>

            <Form.Item label="状态" name="status">
              <Radio.Group>
                <Radio value="show">显示</Radio>
                <Radio value="hide">隐藏</Radio>
              </Radio.Group>
            </Form.Item>
          </Form>
        </Modal>

        <Modal
          title="编辑分类"
          open={openEdit}
          onOk={onEditOk}
          onCancel={onEditCancel}
          okText="保存"
          cancelText="取消"
          width={640}
          rootClassName="compact-modal"
          bodyStyle={{ padding: 12, maxHeight: '60vh', overflow: 'auto' }}
        >
          <Form
            form={editForm}
            layout="horizontal"
            labelCol={{ span: 6 }}
            wrapperCol={{ span: 18 }}
            requiredMark={true}
            initialValues={{ parentId: 0, status: 'show', sort: 0 }}
          >
            <Form.Item label="上级分类" name="parentId">
              <Select
                style={{ width: 240 }}
                options={[{ value: 0, label: '顶级分类' }, ...data.map(it => ({ value: it.id, label: it.name }))]}
              />
            </Form.Item>

            <Form.Item label="分类名称" name="name" rules={[{ required: true, message: '请输入分类名称' }]}> 
              <Input placeholder="请输入分类名称" />
            </Form.Item>

            <Form.Item label="分类简介" name="desc" rules={[{ required: true, message: '请输入分类简介' }]}> 
              <Input.TextArea placeholder="请输入分类简介" rows={3} />
            </Form.Item>

            <Form.Item label="分类图片" name="icon" valuePropName="fileList" getValueFromEvent={(e) => e?.fileList}>
              <Upload listType="picture-card" beforeUpload={() => false}>
                +
              </Upload>
            </Form.Item>

            <Form.Item label="排序" name="sort">
              <InputNumber min={0} style={{ width: 160 }} />
            </Form.Item>

            <Form.Item label="状态" name="status">
              <Radio.Group>
                <Radio value="show">显示</Radio>
                <Radio value="hide">隐藏</Radio>
              </Radio.Group>
            </Form.Item>
          </Form>
        </Modal>
      </Card>
    </div>
  );
};

export default ArticleCategory;