import React, { useState } from 'react';
import { Card, Form, Select, Input, Button, Table, Empty, Breadcrumb, Modal, InputNumber, Switch } from 'antd';
import { Link } from 'react-router-dom';

type Permission = {
  id: number;
  name: string;
  type: string; // 路径/类型
  sort: number;
  visible: boolean;
  children?: Permission[];
};

const PermissionSettings: React.FC = () => {
  const [status, setStatus] = useState<string | undefined>();
  const [keyword, setKeyword] = useState<string>('');
  const [openAdd, setOpenAdd] = useState(false);
  const [form] = Form.useForm();

  // Mock 数据：与上传图片页面的列表视觉一致（有层级，可展开）
  const [permissions, setPermissions] = useState<Permission[]>([
    {
      id: 1001,
      name: '主页',
      type: '/admin/index',
      sort: 127,
      visible: true,
      children: [
        { id: 1101, name: '仪表盘', type: '/admin/dashboard', sort: 30, visible: true },
        { id: 1102, name: '欢迎页', type: '/admin/welcome', sort: 20, visible: true },
      ],
    },
    { id: 1002, name: '用户', type: '/admin/user', sort: 125, visible: true },
    { id: 1003, name: '订单', type: '/admin/order', sort: 120, visible: true },
    {
      id: 1004,
      name: '商品',
      type: '/admin/product',
      sort: 115,
      visible: true,
      children: [
        { id: 1401, name: '商品列表', type: '/admin/product/list', sort: 20, visible: true },
        { id: 1402, name: '商品分类', type: '/admin/product/category', sort: 10, visible: true },
      ],
    },
    { id: 1005, name: '营销', type: '/admin/marketing', sort: 110, visible: true },
    { id: 1006, name: '分销', type: '/admin/agent', sort: 105, visible: true },
    { id: 1007, name: '客服', type: '/admin/kefu', sort: 104, visible: true },
    { id: 1008, name: '财务', type: '/admin/finance', sort: 90, visible: true },
    { id: 1009, name: '内容', type: '/admin/cms', sort: 85, visible: true },
    { id: 1010, name: '统计', type: '/admin/setting/pages', sort: 80, visible: true },
    { id: 1011, name: '应用', type: '/admin/app', sort: 70, visible: true },
    { id: 1012, name: '设置', type: '/admin/setting', sort: 1, visible: true },
    { id: 1013, name: '超级', type: '/admin/system', sort: 0, visible: false },
  ]);

  const updateById = (items: Permission[], id: number, updater: (it: Permission) => Permission): Permission[] => {
    return items.map((it) => {
      if (it.id === id) {
        return updater(it);
      }
      return it.children
        ? { ...it, children: updateById(it.children, id, updater) }
        : it;
    });
  };

  const columns = [
    { title: '权限名称', dataIndex: 'name', width: 200 },
    { title: '类型', dataIndex: 'type', width: 260 },
    { title: '排序', dataIndex: 'sort', width: 120 },
    {
      title: '是否显示',
      dataIndex: 'visible',
      width: 140,
      render: (_: any, record: Permission) => (
        <Switch
          checked={record.visible}
          onChange={(checked) => {
            setPermissions((prev) => updateById(prev, record.id, (it) => ({ ...it, visible: checked })));
          }}
        />
      ),
    },
    {
      title: '操作',
      key: 'action',
      width: 120,
      render: () => (
        <Button type="link" size="small" onClick={() => setOpenAdd(true)}>编辑</Button>
      ),
    },
  ];

  return (
    <div>
      <Card>
        {/* 面包屑导航 - 使用 items 避免弃用警告 */}
        <Breadcrumb
          style={{ marginBottom: 20 }}
          items={[
            { title: <Link to="/home">首页</Link> },
            { title: '管理权限' },
            { title: '权限设置' },
          ]}
        />
        <Form layout="inline" style={{ background: '#f7f8fa', padding: 16, borderRadius: 8 }}>
          <Form.Item label="状态">
            <Select
              style={{ width: 180 }}
              placeholder="请选择"
              value={status}
              onChange={setStatus}
              options={[{ value: 'enabled', label: '启用' }, { value: 'disabled', label: '禁用' }]}
              allowClear
            />
          </Form.Item>
          <Form.Item label="身份昵称">
            <Input
              style={{ width: 280 }}
              placeholder="请输入身份昵称"
              value={keyword}
              onChange={(e) => setKeyword(e.target.value)}
            />
          </Form.Item>
          <Form.Item>
            <Button type="primary">查询</Button>
          </Form.Item>
        </Form>

        <div style={{ marginTop: 12, display: 'flex', justifyContent: 'flex-start' }}>
          <Button type="primary" size="small" onClick={() => setOpenAdd(true)}>添加权限项</Button>
        </div>

        <div style={{ marginTop: 16 }} className="upload-like-box">
          <Table
            columns={columns as any}
            dataSource={permissions}
            pagination={false}
            size="small"
            indentSize={16}
            locale={{ emptyText: <Empty description="暂无数据" /> }}
            rowKey="id"
          />
        </div>
      </Card>
      {/* 添加权限项弹层 */}
      <Modal
        title="添加权限项"
        open={openAdd}
        width={800}
        destroyOnClose
        onCancel={() => setOpenAdd(false)}
        footer={[
          <Button key="cancel" onClick={() => setOpenAdd(false)}>取消</Button>,
          <Button key="ok" type="primary" onClick={() => {
            form.validateFields().then(() => {
              setOpenAdd(false);
            });
          }}>提交</Button>
        ]}
      >
        <Form form={form} labelCol={{ span: 5 }} wrapperCol={{ span: 19 }}>
          <Form.Item label="控制器名称" name="name" rules={[{ required: true, message: '请输入控制器名称' }]}> 
            <Input placeholder="请输入控制器名称" />
          </Form.Item>
          <Form.Item label="类型" name="type" rules={[{ required: true, message: '请输入类型路径' }]}> 
            <Input placeholder="例如：/admin/index" />
          </Form.Item>
          <Form.Item label="排序" name="sort" rules={[{ required: true, message: '请输入排序值' }]}> 
            <InputNumber style={{ width: '100%' }} min={0} placeholder="请输入排序" />
          </Form.Item>
          <Form.Item label="是否显示" name="visible" valuePropName="checked" initialValue={true}> 
            <Switch />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
};

export default PermissionSettings;