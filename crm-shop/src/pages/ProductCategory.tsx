import React, { useState } from 'react';
import { Card, Form, Select, Input, Button, Table, Empty, Image, Tag, Breadcrumb } from 'antd';
import { Link } from 'react-router-dom';

const ProductCategory: React.FC = () => {
  const [categoryId, setCategoryId] = useState<string | undefined>();
  const [status, setStatus] = useState<string | undefined>();
  const [keyword, setKeyword] = useState<string>('');

  const columns = [
    { title: 'ID', dataIndex: 'id', width: 80 },
    { title: '分类名称', dataIndex: 'name' },
    { title: '分类图标', dataIndex: 'icon', render: (src: string) => src ? <Image src={src} width={32} height={32} /> : '-' },
    { title: '排序', dataIndex: 'sort', width: 100 },
    { title: '状态', dataIndex: 'status', width: 120, render: (s: string) => <Tag color={s === 'enabled' ? 'blue' : 'default'}>{s === 'enabled' ? '开启' : '关闭'}</Tag> },
    { title: '操作', dataIndex: 'action', width: 160, render: () => <div style={{ display: 'flex', gap: 8 }}><Button type="link">编辑</Button><Button type="link" danger>删除</Button></div> }
  ];

  return (
    <div>
      <Card>
        {/* 面包屑导航 */}
        <Breadcrumb style={{ marginBottom: 20 }}>
          <Breadcrumb.Item>
            <Link to="/">首页</Link>
          </Breadcrumb.Item>
          <Breadcrumb.Item>商品管理</Breadcrumb.Item>
          <Breadcrumb.Item>商品分类</Breadcrumb.Item>
        </Breadcrumb>

        <Form layout="inline" style={{ background: '#f7f8fa', padding: 16, borderRadius: 8 }}>
          <Form.Item label="商品分类">
            <Select
              style={{ width: 180 }}
              placeholder="请选择"
              value={categoryId}
              onChange={setCategoryId}
              options={[{ value: '1', label: '生活家居' }, { value: '2', label: '数码电器' }]}
              allowClear
            />
          </Form.Item>
          <Form.Item label="分类状态">
            <Select
              style={{ width: 180 }}
              placeholder="请选择"
              value={status}
              onChange={setStatus}
              options={[{ value: 'enabled', label: '启用' }, { value: 'disabled', label: '禁用' }]}
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
          <Button type="primary" size="small">添加分类</Button>
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

export default ProductCategory;