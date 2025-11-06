import React, { useState } from 'react';
import { Card, Form, Select, Input, Button, Table, Empty, Breadcrumb, Modal, InputNumber, Switch } from 'antd';
import { Link } from 'react-router-dom';

const PermissionSettings: React.FC = () => {
  const [status, setStatus] = useState<string | undefined>();
  const [keyword, setKeyword] = useState<string>('');
  const [openAdd, setOpenAdd] = useState(false);
  const [form] = Form.useForm();

  const columns = [
    { title: 'ID', dataIndex: 'id' },
    { title: '身份昵称', dataIndex: 'name' },
    { title: '状态', dataIndex: 'status' },
    { title: '操作', dataIndex: 'action' }
  ];

  return (
    <div>
      <Card>
        {/* 面包屑导航 */}
        <Breadcrumb style={{ marginBottom: 20 }}>
          <Breadcrumb.Item>
            <Link to="/">首页</Link>
          </Breadcrumb.Item>
          <Breadcrumb.Item>管理权限</Breadcrumb.Item>
          <Breadcrumb.Item>权限设置</Breadcrumb.Item>
        </Breadcrumb>
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

        <div style={{ marginTop: 16 }}>
          <Table
            columns={columns}
            dataSource={[]}
            pagination={false}
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