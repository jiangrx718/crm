import React, { useState } from 'react';
import { Card, Form, Select, Input, Button, Table, Empty, Breadcrumb } from 'antd';
import { Link } from 'react-router-dom';

const RoleManagement: React.FC = () => {
  const [status, setStatus] = useState<string | undefined>();
  const [keyword, setKeyword] = useState<string>('');

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
          <Breadcrumb.Item>角色管理</Breadcrumb.Item>
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
          <Button type="primary" size="small">添加角色</Button>
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
    </div>
  );
};

export default RoleManagement;